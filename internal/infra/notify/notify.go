package notify

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/pool"

	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/infra/notify/queue"
)

type Notify struct {
	close      atomic.Bool
	logger     *log.Helper
	waitRunner pool.WaitRunner
	hook       func(ctx kratosx.Context, data any) error

	memory queue.Queue
	redis  queue.Queue
}

var (
	notifyIns  *Notify
	notifyOnce sync.Once
)

func NewNotify(conf *conf.Config) *Notify {
	notifyOnce.Do(func() {
		ctx := kratosx.MustContext(context.Background())

		// 根据配置创建指定的队列
		notifyIns = &Notify{
			waitRunner: ctx.WaitRunner(pool.WithWaitRunnerOption(10)),
			memory:     queue.NewMemory(),
			redis:      queue.NewRedis(ctx.Redis()),
			logger:     kratosx.MustContext(context.Background()).Logger(),
		}
	})
	return notifyIns
}

// Start 启动通知
func (n *Notify) Start(hook func(ctx kratosx.Context, item *queue.Item, sender func(item *queue.Item) error) error) {
	// 启动队列
	// memory队列
	go func() {
		n.handler(n.memory)
	}()

	// redis队列
	go func() {
		n.handler(n.redis)
	}()

	// 注册停止处理事件
	ctx := kratosx.MustContext(context.Background())
	ctx.RegisterAfterStop("notify", func() {
		notifyIns.Stop()
	})
}

// 处理队列消息
func (n *Notify) handler(queue queue.Queue) {
	for !n.close.Load() {
		// 取出数据
		item, err := queue.Pop()
		if err != nil {
			n.logger.Errorw("msg", "pop queue data error", "queue", queue.Name(), "error", err)
			continue
		}

		// 执行发送消息hook
		ctx := kratosx.MustContext(
			context.Background(),
			kratosx.WithTrace(item.ItemData.Trace, item.ItemData.Span),
		)
		n.waitRunner.AddTask(ctx, func() error {
			return n.hook(ctx, item)
		})
	}
}

// Stop 停止
func (n *Notify) Stop() {
	// 关闭取数开关
	n.close.Store(true)

	// 等待任务执行完成
	n.waitRunner.Wait()
}
