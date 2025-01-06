package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/types"
)

type Notify interface {
	// ListNotifyCategory 获取通知分类列表
	ListNotifyCategory(ctx kratosx.Context, req *types.ListNotifyCategoryRequest) ([]*entity.NotifyCategory, uint32, error)

	// CreateNotifyCategory 创建通知分类
	CreateNotifyCategory(ctx kratosx.Context, req *entity.NotifyCategory) (uint32, error)

	// UpdateNotifyCategory 更新通知分类
	UpdateNotifyCategory(ctx kratosx.Context, req *entity.NotifyCategory) error

	// DeleteNotifyCategory 删除通知分类
	DeleteNotifyCategory(ctx kratosx.Context, id uint32) error

	// GetNotify 获取通知
	GetNotify(ctx kratosx.Context, id uint32) (*entity.Notify, error)

	// ListNotify 获取通知列表
	ListNotify(ctx kratosx.Context, req *types.ListNotifyRequest) ([]*entity.Notify, uint32, error)

	// CreateNotify 创建通知
	CreateNotify(ctx kratosx.Context, req *entity.Notify) (uint32, error)

	// UpdateNotify 更新通知
	UpdateNotify(ctx kratosx.Context, req *entity.Notify) error
}
