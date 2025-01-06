package queue

import (
	"context"
	"encoding/json"
	"errors"

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
	res, err := m.r.BZPopMax(context.Background(), 0, redisQueue).Result()
	if err != nil {
		return nil, err
	}
	strData, ok := res.Member.(string)
	if !ok {
		return nil, errors.New("redis data type error")
	}

	item := &Item{}
	if err := json.Unmarshal([]byte(strData), &item); err != nil {
		return nil, errors.New("redis data unmarshal error")
	}

	return item, nil
}

// Push 向优先队列中添加元素
func (m *Redis) Push(data *Item) error {
	byteData, _ := json.Marshal(data)
	return m.r.ZAdd(context.Background(), redisQueue, &redis.Z{
		Score:  float64(data.Priority),
		Member: string(byteData),
	}).Err()
}

// Name 确认获取数据
func (m *Redis) Name() string {
	return "redis"
}
