package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/types"
)

type Log interface {
	// ListLog 获取通知分类列表
	ListLog(ctx kratosx.Context, req *types.ListLogRequest) ([]*entity.Log, uint32, error)

	// CreateLog 创建通知分类
	CreateLog(ctx kratosx.Context, req *entity.Log) (uint32, error)

	// UpdateLog 更新通知分类
	UpdateLog(ctx kratosx.Context, req *entity.Log) error

	// DeleteLog 删除通知分类
	DeleteLog(ctx kratosx.Context, id uint32) error
}
