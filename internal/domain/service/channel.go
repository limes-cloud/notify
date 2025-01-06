package service

import (
	"github.com/gogo/protobuf/proto"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/api/notify/errors"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/repository"
	"github.com/limes-cloud/notify/internal/types"
)

type Channel struct {
	conf   *conf.Config
	repo   repository.Channel
	sender repository.Sender
}

func NewChannel(
	conf *conf.Config,
	repo repository.Channel,
	sender repository.Sender,
) *Channel {
	return &Channel{
		conf:   conf,
		repo:   repo,
		sender: sender,
	}
}

const (
	PasswordChannel = "password"
	EmailChannel    = "email"
)

// GetTypes 获取可以发送的渠道类型
func (u *Channel) GetTypes() []*entity.ChannelTyper {
	return u.sender.GetTypes()
}

// ListChannel 获取登陆渠道列表
func (u *Channel) ListChannel(ctx kratosx.Context, req *types.ListChannelRequest) ([]*entity.Channel, uint32, error) {
	list, total, err := u.repo.ListChannel(ctx, req)
	if err != nil {
		ctx.Logger().Errorw("msg", "list channel error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateChannel 创建登陆渠道
func (u *Channel) CreateChannel(ctx kratosx.Context, channel *entity.Channel) (uint32, error) {
	channel.Status = proto.Bool(false)
	id, err := u.repo.CreateChannel(ctx, channel)
	if err != nil {
		ctx.Logger().Errorw("msg", "create channel error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateChannel 更新登陆渠道
func (u *Channel) UpdateChannel(ctx kratosx.Context, channel *entity.Channel) error {
	if err := u.repo.UpdateChannel(ctx, channel); err != nil {
		ctx.Logger().Errorw("msg", "update channel error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteChannel 删除登陆渠道
func (u *Channel) DeleteChannel(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteChannel(ctx, id); err != nil {
		ctx.Logger().Errorw("msg", "delete channel error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
