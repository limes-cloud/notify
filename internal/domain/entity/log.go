package entity

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"strconv"

	ktypes "github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

const (
	tableCount = 32
)

type Log struct {
	NotifyId   uint32   `json:"notifyId" gorm:"column:notify_id"`                    // 通知ID
	ChannelId  uint32   `json:"channelId" gorm:"column:channel_id"`                  // 渠道ID
	User       string   `json:"user" gorm:"column:user"`                             // 用户
	Content    string   `json:"content" gorm:"column:content"`                       // 通知内容
	FromServer string   `json:"fromServer" gorm:"column:from_server"`                // 发送来源
	FromIp     string   `json:"fromIp" gorm:"column:from_ip"`                        // 发送IP
	Status     *bool    `json:"status" gorm:"column:status"`                         // 发送状态
	Reason     *string  `json:"reason" gorm:"column:reason"`                         // 发送失败原因
	Index      *uint32  `json:"index" gorm:"-"`                                      // 索引
	Channel    *Channel `json:"category" gorm:"foreignKey:channel_id;references:id"` // 渠道
	Notify     *Notify  `json:"notify" gorm:"foreignKey:notify_id;references:id"`    // 通知
	ktypes.BaseModel
}

// TableName 获取表名
func (s *Log) TableName() string {
	mod32 := func(cid string) string {
		m := md5.New()
		_, _ = m.Write([]byte(cid))
		hash := m.Sum(nil)
		hashInt := binary.BigEndian.Uint32(hash[:4]) % tableCount
		return strconv.Itoa(int(hashInt))
	}
	if s.User != "" {
		return "send_log_" + mod32(s.User)
	}
	if s.Index != nil {
		return "send_log_" + strconv.Itoa(int(*s.Index))
	}
	return "send_log_0"
}

// IsValid 校验数据是否合法
func (s *Log) IsValid() bool {
	return s.User != ""
}

// BeforeCreate 创建数据前校验
func (s *Log) BeforeCreate(tx *gorm.DB) error {
	if !s.IsValid() {
		return errors.New("can't save invalid data")
	}
	return nil
}

// BeforeUpdate 更新数据前校验
func (s *Log) BeforeUpdate(tx *gorm.DB) error {
	if !s.IsValid() {
		return errors.New("can't save invalid data")
	}
	return nil
}
