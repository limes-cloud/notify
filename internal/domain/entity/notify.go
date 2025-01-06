package entity

import "github.com/limes-cloud/kratosx/types"

type NotifyCategory struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	types.CreateModel
}

type Notify struct {
	CategoryId  uint32          `json:"categoryId" gorm:"column:category_id"`  // 通知类别
	Keyword     string          `json:"keyword" gorm:"column:keyword"`         // 关键字
	Name        string          `json:"name" gorm:"column:name"`               // 通知名称
	Title       string          `json:"title" gorm:"column:title"`             // 通知标题
	Status      *bool           `json:"status" gorm:"column:status"`           // 状态
	Variable    string          `json:"variable" gorm:"column:variable"`       // 变量
	SendMode    string          `json:"sendMode" gorm:"column:send_mode"`      // 发送模式 1:全部推送，2:优先推送
	IsTimely    *bool           `json:"isTimely" gorm:"column:is_timely"`      // 是否具有失效性
	Expire      *uint32         `json:"expire" gorm:"column:expire"`           // 失效时间,单位/s
	Cache       *uint32         `json:"cache" gorm:"column:cache"`             // 缓存时间,单位/s
	Description string          `json:"description" gorm:"column:description"` // 备注
	Category    *NotifyCategory `json:"category" gorm:"foreignKey:category_id;references:id"`
	types.DeleteModel
}
