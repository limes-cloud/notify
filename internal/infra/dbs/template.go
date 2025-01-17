package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/domain/entity"
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
func (r Template) ListTemplate(ctx kratosx.Context, nid uint32) ([]*entity.Template, error) {
	var (
		list []*entity.Template
		fs   = []string{"*"}
	)

	db := ctx.DB().Model(entity.Template{}).
		Preload("Channel", "status=true").
		Select(fs).Where("notify_id = ?", nid)
	db = db.Order("weight desc, id asc")
	return list, db.Find(&list).Error
}

// CreateTemplate 创建数据
func (r Template) CreateTemplate(ctx kratosx.Context, fc *entity.Template) (uint32, error) {
	return fc.Id, ctx.DB().Create(fc).Error
}

// GetTemplate 获取指定的数据
func (r Template) GetTemplate(ctx kratosx.Context, id uint32) (*entity.Template, error) {
	var (
		fc = entity.Template{}
		fs = []string{"*"}
	)

	return &fc, ctx.DB().Select(fs).First(&fc, id).Error
}

// UpdateTemplate 更新数据
func (r Template) UpdateTemplate(ctx kratosx.Context, fc *entity.Template) error {
	return ctx.DB().Updates(fc).Error
}

// DeleteTemplate 删除数据
func (r Template) DeleteTemplate(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Template{}).Error
}
