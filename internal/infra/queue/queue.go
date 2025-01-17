package queue

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/pool"

	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/infra/queue/channel"
	"github.com/limes-cloud/notify/internal/types"
)

type Queue struct {
	close      atomic.Bool
	logger     *log.Helper
	waitRunner pool.WaitRunner
	hook       func(ctx kratosx.Context, req *types.SendNotifyRequest) error
	redis      channel.Queue
}

var (
	notifyIns  *Queue
	notifyOnce sync.Once
)

func NewQueue(conf *conf.Config) *Queue {
	notifyOnce.Do(func() {
		ctx := kratosx.MustContext(context.Background())

		// 根据配置创建指定的队列
		notifyIns = &Queue{
			waitRunner: ctx.WaitRunner(pool.WithWaitRunnerOption(10)),
			redis:      channel.NewRedis(ctx.Redis()),
			logger:     kratosx.MustContext(context.Background()).Logger(),
		}
	})
	return notifyIns
}

func (n *Queue) SetHook(hook func(ctx kratosx.Context, req *types.SendNotifyRequest) error) {
	n.hook = hook
}

// Watch 启动通知
func (n *Queue) Watch(hook func(ctx kratosx.Context, req *types.SendNotifyRequest) error) {
	n.SetHook(hook)

	// 启动队列
	go func() {
		n.handler(n.redis)
	}()

	// 注册停止处理事件
	ctx := kratosx.MustContext(context.Background())
	ctx.RegisterAfterStop("notify", func() {
		notifyIns.stop()
	})
}

func (n *Queue) PushRedis(ctx kratosx.Context, req *types.SendNotifyRequest) error {
	b, _ := json.Marshal(req)
	return n.redis.Push(ctx, &channel.Item{
		Data: string(b),
	})
}

// 处理队列消息
func (n *Queue) handler(queue channel.Queue) {
	for !n.close.Load() {
		// 取出数据
		item, err := queue.Pop()
		if err != nil {
			n.logger.Errorw("msg", "pop queue data error", "queue", queue.Name(), "error", err)
			continue
		}

		// 处理消息
		var req types.SendNotifyRequest
		if err := json.Unmarshal([]byte(item.Data), &req); err != nil {
			n.logger.Errorw("msg", "send data format error", "queue", queue.Name(), "error", err, "data", item.Data)
			continue
		}

		// 检查数据
		if req.User == nil || req.FromServer == "" || len(req.Variable) == 0 || req.Timestamp == 0 {
			n.logger.Errorw("msg", "send data format error", "queue", queue.Name(), "error", err, "data", item.Data)
			continue
		}

		// 执行发送消息hook
		ctx := kratosx.MustContext(
			context.Background(),
			kratosx.WithTrace(item.Trace, item.Span),
		)
		n.waitRunner.AddTask(ctx, func() error {
			return n.hook(ctx, &req)
		})
	}
}

// stop 停止
func (n *Queue) stop() {
	// 关闭取数开关
	n.close.Store(true)

	// 等待任务执行完成
	n.waitRunner.Wait()
}
