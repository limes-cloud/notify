package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/types"
)

type Official interface {
	// ListTemplate 获取通知模板列表
	ListTemplate(ctx kratosx.Context, req *types.OfficialAccount) ([]*types.OfficialTemplate, error)
}
