// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/notify/channel/notify_channel_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationChannelCreateChannel = "/notify.api.notify.channel.v1.Channel/CreateChannel"
const OperationChannelDeleteChannel = "/notify.api.notify.channel.v1.Channel/DeleteChannel"
const OperationChannelListChannel = "/notify.api.notify.channel.v1.Channel/ListChannel"
const OperationChannelListChannelType = "/notify.api.notify.channel.v1.Channel/ListChannelType"
const OperationChannelListOfficialTemplate = "/notify.api.notify.channel.v1.Channel/ListOfficialTemplate"
const OperationChannelUpdateChannel = "/notify.api.notify.channel.v1.Channel/UpdateChannel"

type ChannelHTTPServer interface {
	// CreateChannel CreateChannel 创建发送渠道
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelReply, error)
	// DeleteChannel DeleteChannel 删除发送渠道
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelReply, error)
	// ListChannel ListChannel 获取发送渠道列表
	ListChannel(context.Context, *ListChannelRequest) (*ListChannelReply, error)
	// ListChannelType ListChannelType 获取发送渠道可用列表
	ListChannelType(context.Context, *ListChannelTypeRequest) (*ListChannelTypeReply, error)
	// ListOfficialTemplate ListOfficialTemplate 获取公众号模板
	ListOfficialTemplate(context.Context, *ListOfficialTemplateRequest) (*ListOfficialTemplateReply, error)
	// UpdateChannel UpdateChannel 更新发送渠道
	UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelReply, error)
}

func RegisterChannelHTTPServer(s *http.Server, srv ChannelHTTPServer) {
	r := s.Route("/")
	r.GET("/notify/api/v1/channel/types", _Channel_ListChannelType0_HTTP_Handler(srv))
	r.GET("/notify/api/v1/channels", _Channel_ListChannel0_HTTP_Handler(srv))
	r.POST("/notify/api/v1/channel", _Channel_CreateChannel0_HTTP_Handler(srv))
	r.PUT("/notify/api/v1/channel", _Channel_UpdateChannel0_HTTP_Handler(srv))
	r.DELETE("/notify/api/v1/channel", _Channel_DeleteChannel0_HTTP_Handler(srv))
	r.GET("/notify/api/v1/channel/official_template", _Channel_ListOfficialTemplate0_HTTP_Handler(srv))
}

func _Channel_ListChannelType0_HTTP_Handler(srv ChannelHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListChannelTypeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChannelListChannelType)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListChannelType(ctx, req.(*ListChannelTypeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListChannelTypeReply)
		return ctx.Result(200, reply)
	}
}

func _Channel_ListChannel0_HTTP_Handler(srv ChannelHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListChannelRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChannelListChannel)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListChannel(ctx, req.(*ListChannelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListChannelReply)
		return ctx.Result(200, reply)
	}
}

func _Channel_CreateChannel0_HTTP_Handler(srv ChannelHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateChannelRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChannelCreateChannel)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateChannel(ctx, req.(*CreateChannelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateChannelReply)
		return ctx.Result(200, reply)
	}
}

func _Channel_UpdateChannel0_HTTP_Handler(srv ChannelHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateChannelRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChannelUpdateChannel)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateChannel(ctx, req.(*UpdateChannelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateChannelReply)
		return ctx.Result(200, reply)
	}
}

func _Channel_DeleteChannel0_HTTP_Handler(srv ChannelHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteChannelRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChannelDeleteChannel)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteChannel(ctx, req.(*DeleteChannelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteChannelReply)
		return ctx.Result(200, reply)
	}
}

func _Channel_ListOfficialTemplate0_HTTP_Handler(srv ChannelHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOfficialTemplateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChannelListOfficialTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListOfficialTemplate(ctx, req.(*ListOfficialTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOfficialTemplateReply)
		return ctx.Result(200, reply)
	}
}

type ChannelHTTPClient interface {
	CreateChannel(ctx context.Context, req *CreateChannelRequest, opts ...http.CallOption) (rsp *CreateChannelReply, err error)
	DeleteChannel(ctx context.Context, req *DeleteChannelRequest, opts ...http.CallOption) (rsp *DeleteChannelReply, err error)
	ListChannel(ctx context.Context, req *ListChannelRequest, opts ...http.CallOption) (rsp *ListChannelReply, err error)
	ListChannelType(ctx context.Context, req *ListChannelTypeRequest, opts ...http.CallOption) (rsp *ListChannelTypeReply, err error)
	ListOfficialTemplate(ctx context.Context, req *ListOfficialTemplateRequest, opts ...http.CallOption) (rsp *ListOfficialTemplateReply, err error)
	UpdateChannel(ctx context.Context, req *UpdateChannelRequest, opts ...http.CallOption) (rsp *UpdateChannelReply, err error)
}

type ChannelHTTPClientImpl struct {
	cc *http.Client
}

func NewChannelHTTPClient(client *http.Client) ChannelHTTPClient {
	return &ChannelHTTPClientImpl{client}
}

func (c *ChannelHTTPClientImpl) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...http.CallOption) (*CreateChannelReply, error) {
	var out CreateChannelReply
	pattern := "/notify/api/v1/channel"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChannelCreateChannel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChannelHTTPClientImpl) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...http.CallOption) (*DeleteChannelReply, error) {
	var out DeleteChannelReply
	pattern := "/notify/api/v1/channel"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChannelDeleteChannel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChannelHTTPClientImpl) ListChannel(ctx context.Context, in *ListChannelRequest, opts ...http.CallOption) (*ListChannelReply, error) {
	var out ListChannelReply
	pattern := "/notify/api/v1/channels"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChannelListChannel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChannelHTTPClientImpl) ListChannelType(ctx context.Context, in *ListChannelTypeRequest, opts ...http.CallOption) (*ListChannelTypeReply, error) {
	var out ListChannelTypeReply
	pattern := "/notify/api/v1/channel/types"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChannelListChannelType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChannelHTTPClientImpl) ListOfficialTemplate(ctx context.Context, in *ListOfficialTemplateRequest, opts ...http.CallOption) (*ListOfficialTemplateReply, error) {
	var out ListOfficialTemplateReply
	pattern := "/notify/api/v1/channel/official_template"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChannelListOfficialTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChannelHTTPClientImpl) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...http.CallOption) (*UpdateChannelReply, error) {
	var out UpdateChannelReply
	pattern := "/notify/api/v1/channel"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChannelUpdateChannel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
