package channel

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisQueue = "notify:priority:queue"
)

type Redis struct {
	r *redis.Client
}

// NewRedis 创建内存通知
func NewRedis(c *redis.Client) *Redis {
	return &Redis{
		r: c,
	}
}

// Pop 向优先队列中删除元素
func (m *Redis) Pop() (*Item, error) {
	res, err := m.r.BLPop(context.Background(), time.Minute, redisQueue).Result()
	if err != nil {
		return nil, err
	}

	if len(res) != 2 {
		return nil, errors.New("empty data")
	}

	item := &Item{}
	if err := json.Unmarshal([]byte(res[1]), &item); err != nil {
		return nil, errors.New("redis data unmarshal error")
	}
	return item, nil
}

// Push 向优先队列中添加元素
func (m *Redis) Push(ctx context.Context, data *Item) error {
	byteData, _ := json.Marshal(data)
	return m.r.RPush(ctx, redisQueue, string(byteData)).Err()
}

// Name 确认获取数据
func (m *Redis) Name() string {
	return "redis"
}
