package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/domain/entity"
)

type Template interface {
	// ListTemplate 获取通知模板列表
	ListTemplate(ctx kratosx.Context, nid uint32) ([]*entity.Template, error)

	// CreateTemplate 创建通知模板
	CreateTemplate(ctx kratosx.Context, req *entity.Template) (uint32, error)

	// UpdateTemplate 更新通知模板
	UpdateTemplate(ctx kratosx.Context, req *entity.Template) error

	// DeleteTemplate 删除通知模板
	DeleteTemplate(ctx kratosx.Context, id uint32) error
}
