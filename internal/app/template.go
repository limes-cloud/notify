package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "github.com/limes-cloud/notify/api/notify/template/v1"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/service"
	"github.com/limes-cloud/notify/internal/infra/dbs"
)

type Template struct {
	pb.UnimplementedTemplateServer
	srv *service.Template
}

func NewTemplate(conf *conf.Config) *Template {
	return &Template{
		srv: service.NewTemplate(
			conf,
			dbs.NewTemplate(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewTemplate(c)
		pb.RegisterTemplateHTTPServer(hs, srv)
		pb.RegisterTemplateServer(gs, srv)
	})
}

// ListTemplate 获取反馈建议分类列表
func (fb *Template) ListTemplate(c context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	list, err := fb.srv.ListTemplate(kratosx.MustContext(c), req.NotifyId)
	if err != nil {
		return nil, err
	}
	reply := pb.ListTemplateReply{}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListTemplateReply_Template{
			Id:        item.Id,
			NotifyId:  item.NotifyId,
			ChannelId: item.ChannelId,
			Content:   item.Content,
			Status:    item.Status,
			Weight:    item.Weight,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
			Channel: &pb.ListTemplateReply_Channel{
				Id:   item.Channel.Id,
				Name: item.Channel.Name,
			},
		})
	}
	return &reply, nil
}

// CreateTemplate 创建反馈建议分类
func (fb *Template) CreateTemplate(c context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateReply, error) {
	id, err := fb.srv.CreateTemplate(kratosx.MustContext(c), &entity.Template{
		NotifyId:  req.NotifyId,
		ChannelId: req.ChannelId,
		Content:   req.Content,
		Status:    req.Status,
		Weight:    req.Weight,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTemplateReply{Id: id}, nil
}

// UpdateTemplate 更新反馈建议分类
func (fb *Template) UpdateTemplate(c context.Context, req *pb.UpdateTemplateRequest) (*pb.UpdateTemplateReply, error) {
	if err := fb.srv.UpdateTemplate(kratosx.MustContext(c), &entity.Template{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		NotifyId:  req.NotifyId,
		ChannelId: req.ChannelId,
		Content:   req.Content,
		Status:    req.Status,
		Weight:    req.Weight,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateTemplateReply{}, nil
}

// DeleteTemplate 删除反馈建议分类
func (fb *Template) DeleteTemplate(c context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateReply, error) {
	if err := fb.srv.DeleteTemplate(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteTemplateReply{}, nil
}
