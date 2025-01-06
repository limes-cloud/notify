package queue

// todo kafka 消息队列
const (
	kafkaQueue = "notify:priority:queue"
)

type Conf struct {
	Brokers        []string
	Group          string
	Topic          string
	PlainMechanism struct {
		Username string `json:",optional"`
		Password string `json:",optional"`
	}
	Conns         int  `json:",default=1"`
	Consumers     int  `json:",default=8"`
	Processors    int  `json:",default=8"`
	MinBytes      int  `json:",default=10240"`    // 10K
	MaxBytes      int  `json:",default=10485760"` // 10M
	ForceCommit   bool `json:",default=true"`
	CommitInOrder bool `json:",default=false"`
}

//
// type Kafka struct {
//	r *kafka.Client
// }
//
// var (
//	kafkaIns  *Kafka
//	kafkaOnce sync.Once
// )
//
// // NewKafka 创建内存通知
// func NewKafka(c *kafka.Client) *Kafka {
//	kafkaOnce.Do(func() {
//		kafkaIns = &Kafka{
//			r: c,
//		}
//	})
//	return kafkaIns
// }
//
// // Pop 向优先队列中删除元素
// func (m *Kafka) Pop() (*Item, error) {
//	res, err := m.r.BZPopMax(context.Background(), 0, kafkaQueue).Result()
//	if err != nil {
//		return nil, err
//	}
//	strData, ok := res.Member.(string)
//	if !ok {
//		return nil, errors.New("kafka data type error")
//	}
//
//	item := &Item{}
//	if err := json.Unmarshal([]byte(strData), &item); err != nil {
//		return nil, errors.New("kafka data unmarshal error")
//	}
//
//	return item, nil
// }
//
// // Push 向优先队列中添加元素
// func (m *Kafka) Push(data *Item) error {
//	byteData, _ := json.Marshal(data)
//	return m.r.ZAdd(context.Background(), kafkaQueue, &kafka.Z{
//		Score:  float64(data.Priority),
//		Member: string(byteData),
//	}).Err()
// }
//
// // Commit 确认获取数据
// func (m *Kafka) Commit() error {
//	return nil
// }
