package official

import (
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"

	"github.com/limes-cloud/notify/internal/types"
)

type Template struct {
}

var (
	templateIns  *Template
	templateOnce sync.Once
)

func NewTemplate() *Template {
	templateOnce.Do(func() {
		templateIns = &Template{}
	})
	return templateIns
}

// ListTemplate 获取列表
func (r Template) ListTemplate(ctx kratosx.Context, req *types.OfficialAccount) ([]*types.OfficialTemplate, error) {
	wc := wechat.NewWechat()
	cfg := &offConfig.Config{
		AppID:     req.AppID,
		AppSecret: req.AppSecret,
		Cache:     cache.NewMemory(),
	}
	tpl := wc.GetOfficialAccount(cfg).GetTemplate()
	list, err := tpl.List()
	if err != nil {
		return nil, err
	}

	handleContent := func(content string) []*types.OfficialTemplateFiled {
		content = strings.ReplaceAll(content, "{{", "")
		content = strings.ReplaceAll(content, "}}", "")
		content = strings.ReplaceAll(content, ".DATA", "")

		var list []*types.OfficialTemplateFiled
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			arr := strings.Split(line, ":")
			if len(arr) != 2 {
				continue
			}
			list = append(list, &types.OfficialTemplateFiled{
				Keyword: arr[1],
				Name:    arr[0],
				Value:   "",
				Color:   "#000000",
			})
		}
		return list
	}

	var resList []*types.OfficialTemplate
	for _, item := range list {
		if item.Content == "{{content.DATA}}" {
			continue
		}
		resList = append(resList, &types.OfficialTemplate{
			TemplateID: item.TemplateID,
			Title:      item.Title,
			Fields:     handleContent(item.Content),
			// PrimaryIndustry: item.PrimaryIndustry,
			// DeputyIndustry:  item.DeputyIndustry,
			// Content:         item.Content,
			// Example:         item.Example,
		})
	}
	return resList, err
}
