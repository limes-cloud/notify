package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/types"
)

type Log struct {
}

var (
	logIns  *Log
	logOnce sync.Once
)

func NewLog() *Log {
	logOnce.Do(func() {
		logIns = &Log{}
	})
	return logIns
}

// ListLog 获取列表
func (r Log) ListLog(ctx kratosx.Context, req *types.ListLogRequest) ([]*entity.Log, uint32, error) {
	var (
		m     = entity.Log{}
		list  []*entity.Log
		total int64
	)
	db := ctx.DB().Model(entity.Log{}).Preload("Channel").Preload("Notify")

	if req.User != nil {
		m.User = *req.User
		db = db.Where("user = ?", *req.User)
	}
	if req.NotifyId != nil {
		db = db.Where("notify_id = ?", *req.NotifyId)
	}
	if req.ChannelId != nil {
		db = db.Where("channel_id = ?", *req.ChannelId)
	}
	if req.Index != nil {
		m.Index = req.Index
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateLog 创建数据
func (r Log) CreateLog(ctx kratosx.Context, log *entity.Log) (uint32, error) {
	return log.Id, ctx.DB().Create(log).Error
}

// GetLog 获取指定的数据
func (r Log) GetLog(ctx kratosx.Context, id uint32) (*entity.Log, error) {
	var (
		log = entity.Log{}
		fs  = []string{"*"}
	)

	return &log, ctx.DB().Select(fs).First(&log, id).Error
}

// UpdateLog 更新数据
func (r Log) UpdateLog(ctx kratosx.Context, log *entity.Log) error {
	return ctx.DB().Updates(log).Error
}

// DeleteLog 删除数据
func (r Log) DeleteLog(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Log{}).Error
}
