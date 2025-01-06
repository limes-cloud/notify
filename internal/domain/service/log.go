package service

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/api/notify/errors"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/repository"
	"github.com/limes-cloud/notify/internal/types"
)

type Log struct {
	conf *conf.Config
	repo repository.Log
}

func NewLog(
	conf *conf.Config,
	repo repository.Log,
) *Log {
	return &Log{
		conf: conf,
		repo: repo,
	}
}

// ListLog 获取日志列表
func (u *Log) ListLog(ctx kratosx.Context, req *types.ListLogRequest) ([]*entity.Log, uint32, error) {
	list, total, err := u.repo.ListLog(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify category error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}
