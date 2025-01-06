package sender

import (
	"github.com/limes-cloud/notify/internal/domain/entity"
)

type ISender interface {
	Send(email, title, body string) error
}

type ISenderInitFunc func(ak string, sk string, extra *string) (ISender, error)

type Item struct {
	Keyword string
	Name    string
	Func    ISenderInitFunc
}

type Sender struct {
	bucket map[string]*Item
	list   []*Item
}

var (
	ins = &Sender{bucket: make(map[string]*Item)}
)

func register(keyword string, name string, rf ISenderInitFunc) {
	ins.list = append(ins.list, &Item{
		Keyword: keyword,
		Name:    name,
		Func:    rf,
	})
}

func NewSender() *Sender {
	return ins
}

// GetTypes 获取可以发生的渠道类型
func (s Sender) GetTypes() []*entity.ChannelTyper {
	var list []*entity.ChannelTyper
	for _, item := range s.list {
		list = append(list, &entity.ChannelTyper{
			Keyword: item.Keyword,
			Name:    item.Name,
		})
	}
	return list
}
