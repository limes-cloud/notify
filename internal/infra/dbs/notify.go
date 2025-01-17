package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/types"
)

type Notify struct {
}

var (
	feedbackIns  *Notify
	feedbackOnce sync.Once
)

func NewNotify() *Notify {
	feedbackOnce.Do(func() {
		feedbackIns = &Notify{}
	})
	return feedbackIns
}

// ListNotifyCategory 获取列表
func (r Notify) ListNotifyCategory(ctx kratosx.Context, req *types.ListNotifyCategoryRequest) ([]*entity.NotifyCategory, uint32, error) {
	var (
		list  []*entity.NotifyCategory
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.NotifyCategory{}).Select(fs)

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateNotifyCategory 创建数据
func (r Notify) CreateNotifyCategory(ctx kratosx.Context, fc *entity.NotifyCategory) (uint32, error) {
	return fc.Id, ctx.DB().Create(fc).Error
}

// GetNotifyCategory 获取指定的数据
func (r Notify) GetNotifyCategory(ctx kratosx.Context, id uint32) (*entity.NotifyCategory, error) {
	var (
		fc = entity.NotifyCategory{}
		fs = []string{"*"}
	)

	return &fc, ctx.DB().Select(fs).First(&fc, id).Error
}

// UpdateNotifyCategory 更新数据
func (r Notify) UpdateNotifyCategory(ctx kratosx.Context, fc *entity.NotifyCategory) error {
	return ctx.DB().Updates(fc).Error
}

// DeleteNotifyCategory 删除数据
func (r Notify) DeleteNotifyCategory(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.NotifyCategory{}).Error
}

// GetNotify 获取指定的数据
func (r Notify) GetNotify(ctx kratosx.Context, id uint32) (*entity.Notify, error) {
	var (
		feedback = entity.Notify{}
		fs       = []string{"*"}
	)
	return &feedback, ctx.DB().Select(fs).First(&feedback, id).Error
}

// GetNotifyByKeyword 获取指定的数据
func (r Notify) GetNotifyByKeyword(ctx kratosx.Context, keyword string) (*entity.Notify, error) {
	var (
		feedback = entity.Notify{}
		fs       = []string{"*"}
	)
	return &feedback, ctx.DB().Select(fs).First(&feedback, "keyword=?", keyword).Error
}

// ListNotify 获取列表
func (r Notify) ListNotify(ctx kratosx.Context, req *types.ListNotifyRequest) ([]*entity.Notify, uint32, error) {
	var (
		list  []*entity.Notify
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Notify{}).Select(fs)
	db = db.Preload("Category")

	if req.CategoryId != nil {
		db = db.Where("category_id = ?", *req.CategoryId)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Name != nil {
		db = db.Where("name like ?", *req.Name+"%")
	}
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Priority != nil {
		db = db.Where("priority = ?", *req.Priority)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateNotify 创建数据
func (r Notify) CreateNotify(ctx kratosx.Context, feedback *entity.Notify) (uint32, error) {
	return feedback.Id, ctx.DB().Create(feedback).Error
}

// UpdateNotify 更新数据
func (r Notify) UpdateNotify(ctx kratosx.Context, feedback *entity.Notify) error {
	return ctx.DB().Updates(feedback).Error
}
