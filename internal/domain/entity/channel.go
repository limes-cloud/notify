package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Channel struct {
	Type   string  `json:"type" gorm:"column:type"`
	Name   string  `json:"name" gorm:"column:name"`
	Ak     *string `json:"ak" gorm:"column:ak"`
	Sk     *string `json:"sk" gorm:"column:sk"`
	Extra  *string `json:"extra" gorm:"column:extra"`
	Status *bool   `json:"status" gorm:"column:status"`
	types.BaseModel
}

type ChannelTyper struct {
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
}

func (ch *Channel) GetAkString() string {
	if ch.Ak == nil {
		return ""
	}
	return *ch.Ak
}

func (ch *Channel) GetSkString() string {
	if ch.Sk == nil {
		return ""
	}
	return *ch.Sk
}

func (ch *Channel) GetExtraString() string {
	if ch.Extra == nil {
		return ""
	}
	return *ch.Extra
}
