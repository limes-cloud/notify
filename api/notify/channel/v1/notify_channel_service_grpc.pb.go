// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/notify/channel/notify_channel_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Channel_ListChannelType_FullMethodName      = "/notify.api.notify.channel.v1.Channel/ListChannelType"
	Channel_ListChannel_FullMethodName          = "/notify.api.notify.channel.v1.Channel/ListChannel"
	Channel_CreateChannel_FullMethodName        = "/notify.api.notify.channel.v1.Channel/CreateChannel"
	Channel_UpdateChannel_FullMethodName        = "/notify.api.notify.channel.v1.Channel/UpdateChannel"
	Channel_DeleteChannel_FullMethodName        = "/notify.api.notify.channel.v1.Channel/DeleteChannel"
	Channel_ListOfficialTemplate_FullMethodName = "/notify.api.notify.channel.v1.Channel/ListOfficialTemplate"
)

// ChannelClient is the client API for Channel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelClient interface {
	// ListChannelType 获取发送渠道可用列表
	ListChannelType(ctx context.Context, in *ListChannelTypeRequest, opts ...grpc.CallOption) (*ListChannelTypeReply, error)
	// ListChannel 获取发送渠道列表
	ListChannel(ctx context.Context, in *ListChannelRequest, opts ...grpc.CallOption) (*ListChannelReply, error)
	// CreateChannel 创建发送渠道
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelReply, error)
	// UpdateChannel 更新发送渠道
	UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelReply, error)
	// DeleteChannel 删除发送渠道
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelReply, error)
	// ListOfficialTemplate 获取公众号模板
	ListOfficialTemplate(ctx context.Context, in *ListOfficialTemplateRequest, opts ...grpc.CallOption) (*ListOfficialTemplateReply, error)
}

type channelClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelClient(cc grpc.ClientConnInterface) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) ListChannelType(ctx context.Context, in *ListChannelTypeRequest, opts ...grpc.CallOption) (*ListChannelTypeReply, error) {
	out := new(ListChannelTypeReply)
	err := c.cc.Invoke(ctx, Channel_ListChannelType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ListChannel(ctx context.Context, in *ListChannelRequest, opts ...grpc.CallOption) (*ListChannelReply, error) {
	out := new(ListChannelReply)
	err := c.cc.Invoke(ctx, Channel_ListChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelReply, error) {
	out := new(CreateChannelReply)
	err := c.cc.Invoke(ctx, Channel_CreateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelReply, error) {
	out := new(UpdateChannelReply)
	err := c.cc.Invoke(ctx, Channel_UpdateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelReply, error) {
	out := new(DeleteChannelReply)
	err := c.cc.Invoke(ctx, Channel_DeleteChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ListOfficialTemplate(ctx context.Context, in *ListOfficialTemplateRequest, opts ...grpc.CallOption) (*ListOfficialTemplateReply, error) {
	out := new(ListOfficialTemplateReply)
	err := c.cc.Invoke(ctx, Channel_ListOfficialTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
// All implementations must embed UnimplementedChannelServer
// for forward compatibility
type ChannelServer interface {
	// ListChannelType 获取发送渠道可用列表
	ListChannelType(context.Context, *ListChannelTypeRequest) (*ListChannelTypeReply, error)
	// ListChannel 获取发送渠道列表
	ListChannel(context.Context, *ListChannelRequest) (*ListChannelReply, error)
	// CreateChannel 创建发送渠道
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelReply, error)
	// UpdateChannel 更新发送渠道
	UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelReply, error)
	// DeleteChannel 删除发送渠道
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelReply, error)
	// ListOfficialTemplate 获取公众号模板
	ListOfficialTemplate(context.Context, *ListOfficialTemplateRequest) (*ListOfficialTemplateReply, error)
	mustEmbedUnimplementedChannelServer()
}

// UnimplementedChannelServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (UnimplementedChannelServer) ListChannelType(context.Context, *ListChannelTypeRequest) (*ListChannelTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannelType not implemented")
}
func (UnimplementedChannelServer) ListChannel(context.Context, *ListChannelRequest) (*ListChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannel not implemented")
}
func (UnimplementedChannelServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedChannelServer) UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}
func (UnimplementedChannelServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedChannelServer) ListOfficialTemplate(context.Context, *ListOfficialTemplateRequest) (*ListOfficialTemplateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfficialTemplate not implemented")
}
func (UnimplementedChannelServer) mustEmbedUnimplementedChannelServer() {}

// UnsafeChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServer will
// result in compilation errors.
type UnsafeChannelServer interface {
	mustEmbedUnimplementedChannelServer()
}

func RegisterChannelServer(s grpc.ServiceRegistrar, srv ChannelServer) {
	s.RegisterService(&Channel_ServiceDesc, srv)
}

func _Channel_ListChannelType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ListChannelType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Channel_ListChannelType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ListChannelType(ctx, req.(*ListChannelTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ListChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ListChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Channel_ListChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ListChannel(ctx, req.(*ListChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Channel_CreateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Channel_UpdateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).UpdateChannel(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Channel_DeleteChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ListOfficialTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfficialTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ListOfficialTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Channel_ListOfficialTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ListOfficialTemplate(ctx, req.(*ListOfficialTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Channel_ServiceDesc is the grpc.ServiceDesc for Channel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Channel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notify.api.notify.channel.v1.Channel",
	HandlerType: (*ChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListChannelType",
			Handler:    _Channel_ListChannelType_Handler,
		},
		{
			MethodName: "ListChannel",
			Handler:    _Channel_ListChannel_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _Channel_CreateChannel_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _Channel_UpdateChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _Channel_DeleteChannel_Handler,
		},
		{
			MethodName: "ListOfficialTemplate",
			Handler:    _Channel_ListOfficialTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/notify/channel/notify_channel_service.proto",
}
