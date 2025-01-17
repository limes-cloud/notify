package sender

import (
	"errors"

	"github.com/limes-cloud/notify/internal/domain/entity"
)

type ISender interface {
	Send(user, title, body string) error
}

type ISenderInitFuncOption struct {
	ak    string
	sk    string
	extra string
}

type ISenderInitOptionFunc func(*ISenderInitFuncOption)

func WithAK(ak string) ISenderInitOptionFunc {
	return func(o *ISenderInitFuncOption) {
		o.ak = ak
	}
}

func WithSK(sk string) ISenderInitOptionFunc {
	return func(o *ISenderInitFuncOption) {
		o.sk = sk
	}
}

func WithExtra(extra string) ISenderInitOptionFunc {
	return func(o *ISenderInitFuncOption) {
		o.extra = extra
	}
}

type ISenderInitFunc func(opts ...ISenderInitOptionFunc) (ISender, error)

type Item struct {
	Keyword string
	Name    string
	Func    ISenderInitFunc
}

type Sender struct {
	list   []*Item
	bucket map[string]*Item
}

var (
	ins = &Sender{bucket: make(map[string]*Item)}
)

func register(keyword string, name string, fn ISenderInitFunc) {
	ni := &Item{
		Keyword: keyword,
		Name:    name,
		Func:    fn,
	}
	ins.bucket[keyword] = ni
	ins.list = append(ins.list, ni)
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

// GetSender 获取发送器
func (s Sender) GetSender(name string) (ISenderInitFunc, error) {
	if ins.bucket[name] == nil {
		return nil, errors.New(name + " sender not exist")
	}
	return ins.bucket[name].Func, nil
}

func (s Sender) IsEmail(name string) bool {
	return name == "email"
}

func (s Sender) IsOfficialAccount(name string) bool {
	return name == "official_account"
}
