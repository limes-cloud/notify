package sender

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"

	"github.com/limes-cloud/notify/internal/types"
)

type OfficialAccount struct {
	AK    string `json:"ak"`
	SK    string `json:"sk"`
	AppID string `json:"appId"`
}

func init() {
	register("official_account", "微信公众号", NewOfficialAccount)
}

// NewOfficialAccount 初始化公众号发送
func NewOfficialAccount(opts ...ISenderInitOptionFunc) (ISender, error) {
	options := &ISenderInitFuncOption{}
	for _, opt := range opts {
		opt(options)
	}

	if options.extra == "" {
		return nil, errors.New("extra info in empty")
	}
	oa := &OfficialAccount{AK: options.ak, SK: options.sk}
	if err := json.Unmarshal([]byte(options.extra), oa); err != nil {
		return nil, err
	}
	return oa, nil
}

// Send 发送邮件
func (oa *OfficialAccount) Send(user, _, body string) error {
	wc := wechat.NewWechat()
	cfg := &offConfig.Config{
		AppID:     oa.AK,
		AppSecret: oa.SK,
		Cache:     cache.NewMemory(),
	}

	var (
		msg  = types.OfficialTemplate{}
		data = map[string]*message.TemplateDataItem{}
	)
	if err := json.Unmarshal([]byte(body), &msg); err != nil {
		return fmt.Errorf("format error %s", err.Error())
	}
	for _, item := range msg.Fields {
		data[item.Keyword] = &message.TemplateDataItem{
			Value: item.Value,
			Color: item.Color,
		}
	}

	tpl := &message.TemplateMessage{
		ToUser:     user,
		TemplateID: msg.TemplateID,
		Data:       data,
		URL:        msg.URL,
	}
	tpl.MiniProgram.AppID = msg.MiniProgram.AppID
	tpl.MiniProgram.PagePath = msg.MiniProgram.PagePath
	_, err := wc.GetOfficialAccount(cfg).GetTemplate().Send(tpl)
	return err
}
