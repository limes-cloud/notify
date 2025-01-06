package service

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/api/notify/errors"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/repository"
)

type Template struct {
	conf   *conf.Config
	repo   repository.Template
	notify repository.Notify
}

func NewTemplate(
	conf *conf.Config,
	repo repository.Template,
) *Template {
	return &Template{
		conf: conf,
		repo: repo,
	}
}

// ListTemplate 获取模板列表
func (u *Template) ListTemplate(ctx kratosx.Context, nid uint32) ([]*entity.Template, error) {
	list, err := u.repo.ListTemplate(ctx, nid)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify category error", "err", err.Error())
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreateTemplate 创建模板
func (u *Template) CreateTemplate(ctx kratosx.Context, req *entity.Template) (uint32, error) {
	// 获取通知信息
	// notify, err := u.notify.GetNotify(ctx, req.NotifyId)
	// if err != nil {
	//	return 0, errors.CreateError(err.Error())
	// }

	// 检查模板是否合法

	id, err := u.repo.CreateTemplate(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create notify category error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateTemplate 更新模板
func (u *Template) UpdateTemplate(ctx kratosx.Context, req *entity.Template) error {
	if err := u.repo.UpdateTemplate(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update notify category error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteTemplate 删除模板
func (u *Template) DeleteTemplate(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteTemplate(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete notify category error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
