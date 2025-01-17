// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/notify/template/notify_template.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotifyId uint32 `protobuf:"varint,1,opt,name=notifyId,proto3" json:"notifyId,omitempty"`
}

func (x *ListTemplateRequest) Reset() {
	*x = ListTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateRequest) ProtoMessage() {}

func (x *ListTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateRequest.ProtoReflect.Descriptor instead.
func (*ListTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{0}
}

func (x *ListTemplateRequest) GetNotifyId() uint32 {
	if x != nil {
		return x.NotifyId
	}
	return 0
}

type ListTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ListTemplateReply_Template `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListTemplateReply) Reset() {
	*x = ListTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateReply) ProtoMessage() {}

func (x *ListTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateReply.ProtoReflect.Descriptor instead.
func (*ListTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{1}
}

func (x *ListTemplateReply) GetList() []*ListTemplateReply_Template {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotifyId  uint32  `protobuf:"varint,1,opt,name=notifyId,proto3" json:"notifyId,omitempty"`
	ChannelId uint32  `protobuf:"varint,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Content   string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status    *bool   `protobuf:"varint,4,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Weight    *uint32 `protobuf:"varint,5,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTemplateRequest) GetNotifyId() uint32 {
	if x != nil {
		return x.NotifyId
	}
	return 0
}

func (x *CreateTemplateRequest) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *CreateTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateTemplateRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *CreateTemplateRequest) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type CreateTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTemplateReply) Reset() {
	*x = CreateTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateReply) ProtoMessage() {}

func (x *CreateTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateReply.ProtoReflect.Descriptor instead.
func (*CreateTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTemplateReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NotifyId  uint32  `protobuf:"varint,2,opt,name=notifyId,proto3" json:"notifyId,omitempty"`
	ChannelId uint32  `protobuf:"varint,3,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Content   string  `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status    *bool   `protobuf:"varint,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Weight    *uint32 `protobuf:"varint,6,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *UpdateTemplateRequest) Reset() {
	*x = UpdateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateRequest) ProtoMessage() {}

func (x *UpdateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTemplateRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTemplateRequest) GetNotifyId() uint32 {
	if x != nil {
		return x.NotifyId
	}
	return 0
}

func (x *UpdateTemplateRequest) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateTemplateRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *UpdateTemplateRequest) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type UpdateTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTemplateReply) Reset() {
	*x = UpdateTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateReply) ProtoMessage() {}

func (x *UpdateTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateReply.ProtoReflect.Descriptor instead.
func (*UpdateTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{5}
}

type DeleteTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTemplateRequest) Reset() {
	*x = DeleteTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateRequest) ProtoMessage() {}

func (x *DeleteTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTemplateRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTemplateReply) Reset() {
	*x = DeleteTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateReply) ProtoMessage() {}

func (x *DeleteTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateReply.ProtoReflect.Descriptor instead.
func (*DeleteTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{7}
}

type ListTemplateReply_Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListTemplateReply_Channel) Reset() {
	*x = ListTemplateReply_Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplateReply_Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateReply_Channel) ProtoMessage() {}

func (x *ListTemplateReply_Channel) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateReply_Channel.ProtoReflect.Descriptor instead.
func (*ListTemplateReply_Channel) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListTemplateReply_Channel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListTemplateReply_Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListTemplateReply_Channel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListTemplateReply_Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NotifyId  uint32                     `protobuf:"varint,2,opt,name=notifyId,proto3" json:"notifyId,omitempty"`
	ChannelId uint32                     `protobuf:"varint,3,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Content   string                     `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status    *bool                      `protobuf:"varint,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Weight    *uint32                    `protobuf:"varint,6,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	CreatedAt uint32                     `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt uint32                     `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Channel   *ListTemplateReply_Channel `protobuf:"bytes,9,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *ListTemplateReply_Template) Reset() {
	*x = ListTemplateReply_Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notify_template_notify_template_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplateReply_Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateReply_Template) ProtoMessage() {}

func (x *ListTemplateReply_Template) ProtoReflect() protoreflect.Message {
	mi := &file_api_notify_template_notify_template_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateReply_Template.ProtoReflect.Descriptor instead.
func (*ListTemplateReply_Template) Descriptor() ([]byte, []int) {
	return file_api_notify_template_notify_template_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ListTemplateReply_Template) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListTemplateReply_Template) GetNotifyId() uint32 {
	if x != nil {
		return x.NotifyId
	}
	return 0
}

func (x *ListTemplateReply_Template) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ListTemplateReply_Template) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ListTemplateReply_Template) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *ListTemplateReply_Template) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *ListTemplateReply_Template) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListTemplateReply_Template) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ListTemplateReply_Template) GetChannel() *ListTemplateReply_Channel {
	if x != nil {
		return x.Channel
	}
	return nil
}

var File_api_notify_template_notify_template_proto protoreflect.FileDescriptor

var file_api_notify_template_notify_template_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64, 0x22,
	0xf6, 0x03, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x1a, 0x41, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xce, 0x02, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x52,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd4, 0x01,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x30, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x36, 0x0a, 0x1d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_notify_template_notify_template_proto_rawDescOnce sync.Once
	file_api_notify_template_notify_template_proto_rawDescData = file_api_notify_template_notify_template_proto_rawDesc
)

func file_api_notify_template_notify_template_proto_rawDescGZIP() []byte {
	file_api_notify_template_notify_template_proto_rawDescOnce.Do(func() {
		file_api_notify_template_notify_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_notify_template_notify_template_proto_rawDescData)
	})
	return file_api_notify_template_notify_template_proto_rawDescData
}

var file_api_notify_template_notify_template_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_notify_template_notify_template_proto_goTypes = []interface{}{
	(*ListTemplateRequest)(nil),        // 0: notify.api.notify.template.v1.ListTemplateRequest
	(*ListTemplateReply)(nil),          // 1: notify.api.notify.template.v1.ListTemplateReply
	(*CreateTemplateRequest)(nil),      // 2: notify.api.notify.template.v1.CreateTemplateRequest
	(*CreateTemplateReply)(nil),        // 3: notify.api.notify.template.v1.CreateTemplateReply
	(*UpdateTemplateRequest)(nil),      // 4: notify.api.notify.template.v1.UpdateTemplateRequest
	(*UpdateTemplateReply)(nil),        // 5: notify.api.notify.template.v1.UpdateTemplateReply
	(*DeleteTemplateRequest)(nil),      // 6: notify.api.notify.template.v1.DeleteTemplateRequest
	(*DeleteTemplateReply)(nil),        // 7: notify.api.notify.template.v1.DeleteTemplateReply
	(*ListTemplateReply_Channel)(nil),  // 8: notify.api.notify.template.v1.ListTemplateReply.Channel
	(*ListTemplateReply_Template)(nil), // 9: notify.api.notify.template.v1.ListTemplateReply.Template
}
var file_api_notify_template_notify_template_proto_depIdxs = []int32{
	9, // 0: notify.api.notify.template.v1.ListTemplateReply.list:type_name -> notify.api.notify.template.v1.ListTemplateReply.Template
	8, // 1: notify.api.notify.template.v1.ListTemplateReply.Template.channel:type_name -> notify.api.notify.template.v1.ListTemplateReply.Channel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_notify_template_notify_template_proto_init() }
func file_api_notify_template_notify_template_proto_init() {
	if File_api_notify_template_notify_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_notify_template_notify_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplateReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplateReply_Channel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_notify_template_notify_template_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplateReply_Template); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_notify_template_notify_template_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_api_notify_template_notify_template_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_api_notify_template_notify_template_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_notify_template_notify_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_notify_template_notify_template_proto_goTypes,
		DependencyIndexes: file_api_notify_template_notify_template_proto_depIdxs,
		MessageInfos:      file_api_notify_template_notify_template_proto_msgTypes,
	}.Build()
	File_api_notify_template_notify_template_proto = out.File
	file_api_notify_template_notify_template_proto_rawDesc = nil
	file_api_notify_template_notify_template_proto_goTypes = nil
	file_api_notify_template_notify_template_proto_depIdxs = nil
}
