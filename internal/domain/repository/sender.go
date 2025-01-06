package repository

import (
	"github.com/limes-cloud/notify/internal/domain/entity"
)

type Sender interface {
	// GetTypes 获取发送类型
	GetTypes() []*entity.ChannelTyper
}
