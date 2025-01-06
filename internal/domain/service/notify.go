package service

import (
	"time"

	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/notify/api/notify/errors"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/repository"
	"github.com/limes-cloud/notify/internal/types"
)

type Notify struct {
	conf *conf.Config
	repo repository.Notify
}

func NewNotify(
	conf *conf.Config,
	repo repository.Notify,
) *Notify {
	return &Notify{
		conf: conf,
		repo: repo,
	}
}

// ListNotifyCategory 获取反馈建议分类列表
func (u *Notify) ListNotifyCategory(ctx kratosx.Context, req *types.ListNotifyCategoryRequest) ([]*entity.NotifyCategory, uint32, error) {
	list, total, err := u.repo.ListNotifyCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify category error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateNotifyCategory 创建反馈建议分类
func (u *Notify) CreateNotifyCategory(ctx kratosx.Context, req *entity.NotifyCategory) (uint32, error) {
	id, err := u.repo.CreateNotifyCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create notify category error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNotifyCategory 更新反馈建议分类
func (u *Notify) UpdateNotifyCategory(ctx kratosx.Context, req *entity.NotifyCategory) error {
	if err := u.repo.UpdateNotifyCategory(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update notify category error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotifyCategory 删除反馈建议分类
func (u *Notify) DeleteNotifyCategory(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteNotifyCategory(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete notify category error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListNotify 获取反馈建议列表
func (u *Notify) ListNotify(ctx kratosx.Context, req *types.ListNotifyRequest) ([]*entity.Notify, uint32, error) {
	list, total, err := u.repo.ListNotify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateNotify 创建反馈建议
func (u *Notify) CreateNotify(ctx kratosx.Context, notify *entity.Notify) (uint32, error) {
	id, err := u.repo.CreateNotify(ctx, notify)
	if err != nil {
		ctx.Logger().Warnw("msg", "create notify error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNotify 更新反馈建议
func (u *Notify) UpdateNotify(ctx kratosx.Context, notify *entity.Notify) error {
	if err := u.repo.UpdateNotify(ctx, notify); err != nil {
		ctx.Logger().Warnw("msg", "update notify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotify 删除反馈建议
func (u *Notify) DeleteNotify(ctx kratosx.Context, id uint32) error {
	// 获取当前的notify
	notify, err := u.repo.GetNotify(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}

	notify.DeletedAt = ktypes.NullInt64{
		Valid: true,
		Int64: time.Now().Unix(),
	}
	if err := u.repo.UpdateNotify(ctx, notify); err != nil {
		ctx.Logger().Warnw("msg", "update notify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}
