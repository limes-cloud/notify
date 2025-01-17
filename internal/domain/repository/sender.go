package repository

import (
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/infra/sender"
)

type Sender interface {
	// GetTypes 获取发送类型
	GetTypes() []*entity.ChannelTyper

	// GetSender 获取发送通路
	GetSender(name string) (sender.ISenderInitFunc, error)

	IsEmail(name string) bool
	IsOfficialAccount(name string) bool
}
