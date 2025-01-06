package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "github.com/limes-cloud/notify/api/notify/channel/v1"
	"github.com/limes-cloud/notify/api/notify/errors"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/service"
	"github.com/limes-cloud/notify/internal/infra/dbs"
	"github.com/limes-cloud/notify/internal/infra/sender"
	"github.com/limes-cloud/notify/internal/types"
)

type Channel struct {
	pb.UnimplementedChannelServer
	srv *service.Channel
}

func NewChannel(conf *conf.Config) *Channel {
	return &Channel{
		srv: service.NewChannel(conf, dbs.NewChannel(), sender.NewSender()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewChannel(c)
		pb.RegisterChannelHTTPServer(hs, srv)
		pb.RegisterChannelServer(gs, srv)
	})
}

// ListChannelType 获取发送渠道可用列表
func (ch *Channel) ListChannelType(_ context.Context, _ *pb.ListChannelTypeRequest) (*pb.ListChannelTypeReply, error) {
	tps := ch.srv.GetTypes()
	reply := pb.ListChannelTypeReply{}
	if err := valx.Transform(tps, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListChannel 获取发送渠道列表
func (ch *Channel) ListChannel(c context.Context, req *pb.ListChannelRequest) (*pb.ListChannelReply, error) {
	list, total, err := ch.srv.ListChannel(kratosx.MustContext(c), &types.ListChannelRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListChannelReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListChannelReply_Channel{
			Id:        item.Id,
			Type:      item.Type,
			Name:      item.Name,
			Status:    item.Status,
			Ak:        item.Ak,
			Sk:        item.Sk,
			Extra:     item.Extra,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateChannel 创建发送渠道
func (ch *Channel) CreateChannel(c context.Context, req *pb.CreateChannelRequest) (*pb.CreateChannelReply, error) {
	id, err := ch.srv.CreateChannel(kratosx.MustContext(c), &entity.Channel{
		Type:   req.Type,
		Name:   req.Name,
		Status: req.Status,
		Ak:     req.Ak,
		Sk:     req.Sk,
		Extra:  req.Extra,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateChannelReply{Id: id}, nil
}

// UpdateChannel 更新发送渠道
func (ch *Channel) UpdateChannel(c context.Context, req *pb.UpdateChannelRequest) (*pb.UpdateChannelReply, error) {
	if err := ch.srv.UpdateChannel(kratosx.MustContext(c), &entity.Channel{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		Type:      req.Type,
		Name:      req.Name,
		Status:    req.Status,
		Ak:        req.Ak,
		Sk:        req.Sk,
		Extra:     req.Extra,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateChannelReply{}, nil
}

// DeleteChannel 删除发送渠道
func (ch *Channel) DeleteChannel(c context.Context, req *pb.DeleteChannelRequest) (*pb.DeleteChannelReply, error) {
	if err := ch.srv.DeleteChannel(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteChannelReply{}, nil
}
