package app

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gogo/protobuf/proto"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "github.com/limes-cloud/notify/api/notify/notify/v1"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/service"
	"github.com/limes-cloud/notify/internal/infra/dbs"
	"github.com/limes-cloud/notify/internal/infra/queue"
	"github.com/limes-cloud/notify/internal/infra/sender"
	"github.com/limes-cloud/notify/internal/types"
)

type Notify struct {
	pb.UnimplementedNotifyServer
	srv *service.Notify
}

func NewNotify(conf *conf.Config) *Notify {
	return &Notify{
		srv: service.NewNotify(
			conf,
			dbs.NewNotify(),
			dbs.NewTemplate(),
			sender.NewSender(),
			dbs.NewLog(),
			queue.NewQueue(conf),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewNotify(c)
		pb.RegisterNotifyHTTPServer(hs, srv)
		pb.RegisterNotifyServer(gs, srv)
	})
}

// ListNotifyCategory 获取通知分类列表
func (fb *Notify) ListNotifyCategory(c context.Context, req *pb.ListNotifyCategoryRequest) (*pb.ListNotifyCategoryReply, error) {
	list, total, err := fb.srv.ListNotifyCategory(kratosx.MustContext(c), &types.ListNotifyCategoryRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListNotifyCategoryReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListNotifyCategoryReply_NotifyCategory{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
		})
	}
	return &reply, nil
}

// CreateNotifyCategory 创建通知分类
func (fb *Notify) CreateNotifyCategory(c context.Context, req *pb.CreateNotifyCategoryRequest) (*pb.CreateNotifyCategoryReply, error) {
	id, err := fb.srv.CreateNotifyCategory(kratosx.MustContext(c), &entity.NotifyCategory{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotifyCategoryReply{Id: id}, nil
}

// UpdateNotifyCategory 更新通知分类
func (fb *Notify) UpdateNotifyCategory(c context.Context, req *pb.UpdateNotifyCategoryRequest) (*pb.UpdateNotifyCategoryReply, error) {
	if err := fb.srv.UpdateNotifyCategory(kratosx.MustContext(c), &entity.NotifyCategory{
		CreateModel: ktypes.CreateModel{Id: req.Id},
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateNotifyCategoryReply{}, nil
}

// DeleteNotifyCategory 删除通知分类
func (fb *Notify) DeleteNotifyCategory(c context.Context, req *pb.DeleteNotifyCategoryRequest) (*pb.DeleteNotifyCategoryReply, error) {
	if err := fb.srv.DeleteNotifyCategory(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteNotifyCategoryReply{}, nil
}

// GetNotify 获取指定通知
func (fb *Notify) GetNotify(c context.Context, req *pb.GetNotifyRequest) (*pb.GetNotifyReply, error) {
	res, err := fb.srv.GetNotify(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetNotifyReply{
		Id:          res.Id,
		CategoryId:  res.CategoryId,
		Keyword:     res.Keyword,
		Name:        res.Name,
		Title:       res.Title,
		Status:      res.Status,
		Variable:    strings.Split(res.Variable, ","),
		Priority:    res.Priority,
		SendMode:    res.SendMode,
		Expire:      res.Expire,
		Cache:       res.Cache,
		Description: res.Description,
		CreatedAt:   uint32(res.CreatedAt),
		UpdatedAt:   uint32(res.UpdatedAt),
	}, nil
}

// ListNotify 获取通知列表
func (fb *Notify) ListNotify(c context.Context, req *pb.ListNotifyRequest) (*pb.ListNotifyReply, error) {
	list, total, err := fb.srv.ListNotify(kratosx.MustContext(c), &types.ListNotifyRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		CategoryId: req.CategoryId,
		Status:     req.Status,
		Keyword:    req.Keyword,
		Name:       req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListNotifyReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListNotifyReply_Notify{
			Id:          item.Id,
			CategoryId:  item.CategoryId,
			Keyword:     item.Keyword,
			Name:        item.Name,
			Title:       item.Title,
			Status:      item.Status,
			Variable:    strings.Split(item.Variable, ","),
			Priority:    item.Priority,
			SendMode:    item.SendMode,
			Expire:      item.Expire,
			Cache:       item.Cache,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
			Category: &pb.ListNotifyReply_Notify_Category{
				Id:   item.Category.Id,
				Name: item.Category.Name,
			},
		})
	}
	return &reply, nil
}

// CreateNotify 创建通知
func (fb *Notify) CreateNotify(c context.Context, req *pb.CreateNotifyRequest) (*pb.CreateNotifyReply, error) {
	id, err := fb.srv.CreateNotify(kratosx.MustContext(c), &entity.Notify{
		CategoryId:  req.CategoryId,
		Keyword:     req.Keyword,
		Name:        req.Name,
		Title:       req.Title,
		Status:      proto.Bool(false),
		Variable:    strings.Join(req.Variable, ","),
		Priority:    req.Priority,
		SendMode:    req.SendMode,
		Expire:      req.Expire,
		Cache:       req.Cache,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotifyReply{Id: id}, nil
}

// DeleteNotify 删除通知
func (fb *Notify) DeleteNotify(c context.Context, req *pb.DeleteNotifyRequest) (*pb.DeleteNotifyReply, error) {
	if err := fb.srv.DeleteNotify(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteNotifyReply{}, nil
}

// UpdateNotify 更新通知
func (fb *Notify) UpdateNotify(c context.Context, req *pb.UpdateNotifyRequest) (*pb.UpdateNotifyReply, error) {
	if err := fb.srv.UpdateNotify(kratosx.MustContext(c), &entity.Notify{
		DeleteModel: ktypes.DeleteModel{Id: req.Id},
		CategoryId:  req.CategoryId,
		Keyword:     req.Keyword,
		Name:        req.Name,
		Title:       req.Title,
		Status:      req.Status,
		Priority:    req.Priority,
		Variable:    strings.Join(req.Variable, ","),
		SendMode:    req.SendMode,
		Expire:      req.Expire,
		Cache:       req.Cache,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateNotifyReply{}, nil
}

// SendNotify 发送通知
func (fb *Notify) SendNotify(c context.Context, req *pb.SendNotifyRequest) (*pb.SendNotifyReply, error) {
	ctx := kratosx.MustContext(c)
	if err := fb.srv.SendNotify(ctx, &types.SendNotifyRequest{
		Notify:   req.Notify,
		Variable: req.Variable,
		User: &types.SendNotifyUser{
			Email:          req.User.GetEmail(),
			OfficialOpenid: req.User.GetOfficialOpenid(),
		},
		Timestamp:  req.Timestamp,
		FromServer: req.FromServer,
		IP:         ctx.ClientIP(),
	}); err != nil {
		return nil, err
	}
	return &pb.SendNotifyReply{}, nil
}

// AsyncSendNotify 异步发送通知
func (fb *Notify) AsyncSendNotify(c context.Context, req *pb.SendNotifyRequest) (*pb.SendNotifyReply, error) {
	ctx := kratosx.MustContext(c)
	if err := fb.srv.AsyncSendNotify(kratosx.MustContext(c), &types.SendNotifyRequest{
		Notify:   req.Notify,
		Variable: req.Variable,
		User: &types.SendNotifyUser{
			Email:          req.User.GetEmail(),
			OfficialOpenid: req.User.GetOfficialOpenid(),
		},
		Timestamp:  req.Timestamp,
		FromServer: req.FromServer,
		IP:         ctx.ClientIP(),
	}); err != nil {
		return nil, err
	}
	return &pb.SendNotifyReply{}, nil
}
