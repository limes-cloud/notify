package entity

import "github.com/limes-cloud/kratosx/types"

type Template struct {
	NotifyId  uint32   `json:"notifyId" gorm:"column:notify_id"`                   // 通知ID
	ChannelId uint32   `json:"channelId" gorm:"column:channel_id"`                 // 渠道ID
	Content   string   `json:"content" gorm:"column:content"`                      // 内容
	Status    *bool    `json:"status" gorm:"column:status"`                        // 状态
	Weight    *uint32  `json:"weight" gorm:"column:weight"`                        // 权重
	Channel   *Channel `json:"channel" gorm:"foreignKey:channel_id;references:id"` // 渠道
	types.BaseModel
}
