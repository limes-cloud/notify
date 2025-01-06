package queue

import (
	"container/heap"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type Memory struct {
	size atomic.Int64
	mux  sync.RWMutex
	data *MemoryPriorityQueue
}

// NewMemory 创建内存通知
func NewMemory() *Memory {
	return &Memory{
		data: &MemoryPriorityQueue{
			queue: make([]*Item, 0),
		},
		mux: sync.RWMutex{},
	}
}

// Pop 向优先队列中删除元素
func (m *Memory) Pop() (*Item, error) {
	// 不存在数据则空转，等待数据到来
	for m.size.Load() == 0 {
		time.Sleep(time.Millisecond * 10)
	}

	// 存在数据则获取数据
	m.mux.Lock()
	defer m.mux.Unlock()
	// 在并发的场景下，存在同时读到非0的场景
	// 再次获取当前队列值，不存在则释放锁，并重新获取数据
	if m.size.Load() == 0 {
		return nil, errors.New("queue is empty")
	}

	data, _ := m.data.Pop().(*Item)
	m.size.Add(-1)

	return data, nil
}

// Push 向优先队列中添加元素
func (m *Memory) Push(data *Item) error {
	m.size.Add(1)
	m.mux.Lock()
	defer m.mux.Unlock()
	heap.Push(m.data, data)
	return nil
}

// Name 确认队列名称
func (m *Memory) Name() string {
	return "memory"
}

// MemoryPriorityQueue 实现了 heap.Interface 并保存
type MemoryPriorityQueue struct {
	queue []*Item
}

// Len 返回优先队列中元素的个数
func (pq *MemoryPriorityQueue) Len() int { return len(pq.queue) }

// Less 根据优先级比较两个元素的大小，这里我们假设更高的优先级优先出队
func (pq *MemoryPriorityQueue) Less(i, j int) bool {
	return pq.queue[i].Priority > pq.queue[j].Priority
}

// Swap 交换两个元素的位置
func (pq *MemoryPriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

// Push 向优先队列中添加元素,实现heap.Push 接口
func (pq *MemoryPriorityQueue) Push(x any) {
	item, ok := x.(*Item)
	if !ok {
		return
	}
	pq.queue = append(pq.queue, item)
}

// Pop 从优先队列中移除最高优先级的元素,实现heap.Push 接口
func (pq *MemoryPriorityQueue) Pop() any {
	// 获取队尾元素,这里必须使用指针获取进行值复制，否则会导致数据逃逸，肯能导致内存泄漏
	// 泄漏原因主要是由于切片操作[:]不会创建新的切片，而是直接引用原数组的子集，但是修改了原结构的len值，导致语法上无法访问。
	// 但是实际上，当前的切片本身还是对该元素进行了引用。
	// 如果切片的容量未发送变化，则不会创建新的数组，则引用关系在没有新元素替换的情况下，将一直保持存在。
	if len(pq.queue) == 0 {
		return nil
	}
	data := *pq.queue[pq.Len()-1]
	pq.queue[pq.Len()-1] = (*Item)(nil) // 避免内存泄漏
	return data
}
