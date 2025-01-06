package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"

	pb "github.com/limes-cloud/notify/api/notify/log/v1"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/service"
	"github.com/limes-cloud/notify/internal/infra/dbs"
	"github.com/limes-cloud/notify/internal/types"
)

type Log struct {
	pb.UnimplementedLogServer
	srv *service.Log
}

func NewLog(conf *conf.Config) *Log {
	return &Log{
		srv: service.NewLog(conf, dbs.NewLog()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewLog(c)
		pb.RegisterLogHTTPServer(hs, srv)
		pb.RegisterLogServer(gs, srv)
	})
}

// ListLog 获取发送渠道列表
func (ch *Log) ListLog(c context.Context, req *pb.ListLogRequest) (*pb.ListLogReply, error) {
	list, total, err := ch.srv.ListLog(kratosx.MustContext(c), &types.ListLogRequest{
		Page:      req.Page,
		PageSize:  req.PageSize,
		User:      req.User,
		Index:     req.Index,
		NotifyId:  req.NotifyId,
		ChannelId: req.ChannelId,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListLogReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListLogReply_Log{
			Id:         item.Id,
			NotifyId:   item.NotifyId,
			ChannelId:  item.ChannelId,
			User:       item.User,
			Content:    item.Content,
			FromServer: item.FromServer,
			FromIp:     item.FromIp,
			Status:     item.Status,
			Reason:     item.Reason,
			CreatedAt:  uint32(item.CreatedAt),
			UpdatedAt:  uint32(item.UpdatedAt),
			Channel: &pb.ListLogReply_Channel{
				Id:   item.Channel.Id,
				Name: item.Channel.Name,
			},
			Notify: &pb.ListLogReply_Notify{
				Id:   item.Notify.Id,
				Name: item.Notify.Name,
			},
		})
	}
	return &reply, nil
}
