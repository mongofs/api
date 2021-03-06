// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: im.proto

package grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{0}
}

type SendMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMessageReq) Reset() {
	*x = SendMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReq) ProtoMessage() {}

func (x *SendMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReq.ProtoReflect.Descriptor instead.
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendMessageReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMessageToMultipleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token []string `protobuf:"bytes,1,rep,name=token,proto3" json:"token,omitempty"`
	Data  []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMessageToMultipleReq) Reset() {
	*x = SendMessageToMultipleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageToMultipleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageToMultipleReq) ProtoMessage() {}

func (x *SendMessageToMultipleReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageToMultipleReq.ProtoReflect.Descriptor instead.
func (*SendMessageToMultipleReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageToMultipleReq) GetToken() []string {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *SendMessageToMultipleReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{3}
}

type OnlinesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *OnlinesReply) Reset() {
	*x = OnlinesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinesReply) ProtoMessage() {}

func (x *OnlinesReply) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinesReply.ProtoReflect.Descriptor instead.
func (*OnlinesReply) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{4}
}

func (x *OnlinesReply) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type BroadcastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BroadcastReq) Reset() {
	*x = BroadcastReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReq) ProtoMessage() {}

func (x *BroadcastReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReq.ProtoReflect.Descriptor instead.
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{5}
}

func (x *BroadcastReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BroadcastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *BroadcastReply) Reset() {
	*x = BroadcastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReply) ProtoMessage() {}

func (x *BroadcastReply) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReply.ProtoReflect.Descriptor instead.
func (*BroadcastReply) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{6}
}

func (x *BroadcastReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PushToClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid int64  `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Msg []byte `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PushToClient) Reset() {
	*x = PushToClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushToClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushToClient) ProtoMessage() {}

func (x *PushToClient) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushToClient.ProtoReflect.Descriptor instead.
func (*PushToClient) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{7}
}

func (x *PushToClient) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *PushToClient) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

// target1 ,content1
// target2 ,content2
type BroadcastByWTIReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BroadcastByWTIReq) Reset() {
	*x = BroadcastByWTIReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastByWTIReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastByWTIReq) ProtoMessage() {}

func (x *BroadcastByWTIReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastByWTIReq.ProtoReflect.Descriptor instead.
func (*BroadcastByWTIReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{8}
}

func (x *BroadcastByWTIReq) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// wti ???????????????????????????target????????????
type WTIDistributeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]*WTIDistribute `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WTIDistributeReply) Reset() {
	*x = WTIDistributeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WTIDistributeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WTIDistributeReply) ProtoMessage() {}

func (x *WTIDistributeReply) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WTIDistributeReply.ProtoReflect.Descriptor instead.
func (*WTIDistributeReply) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{9}
}

func (x *WTIDistributeReply) GetData() map[string]*WTIDistribute {
	if x != nil {
		return x.Data
	}
	return nil
}

type WTIDistribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag        string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`                // ????????????
	Number     int64  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`         // ????????????
	CreateTime int64  `protobuf:"varint,3,opt,name=createTime,proto3" json:"createTime,omitempty"` // ????????????
}

func (x *WTIDistribute) Reset() {
	*x = WTIDistribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WTIDistribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WTIDistribute) ProtoMessage() {}

func (x *WTIDistribute) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WTIDistribute.ProtoReflect.Descriptor instead.
func (*WTIDistribute) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{10}
}

func (x *WTIDistribute) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *WTIDistribute) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *WTIDistribute) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_im_proto protoreflect.FileDescriptor

var file_im_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6d, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3a, 0x0a,
	0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x18, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x26, 0x0a, 0x0c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x24, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x42, 0x79, 0x57, 0x54, 0x49, 0x52, 0x65, 0x71, 0x12,
	0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x42, 0x79, 0x57, 0x54, 0x49, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xa2, 0x01, 0x0a, 0x12, 0x57, 0x54, 0x49, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x57, 0x54, 0x49, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x50, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x57, 0x54, 0x49, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a, 0x0d, 0x57, 0x54, 0x49, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x32, 0xc9, 0x03, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x28, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x07, 0x4f, 0x6e, 0x6c, 0x69, 0x65,
	0x6e, 0x73, 0x12, 0x0f, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x69, 0x6d, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x57, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x69, 0x6d, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18,
	0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x0c, 0x57, 0x54, 0x49, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x42, 0x79, 0x57,
	0x54, 0x49, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x3e, 0x0a, 0x0d, 0x57, 0x54, 0x49, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x0f, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1c, 0x2e, 0x69, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x57, 0x54, 0x49,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_proto_rawDescOnce sync.Once
	file_im_proto_rawDescData = file_im_proto_rawDesc
)

func file_im_proto_rawDescGZIP() []byte {
	file_im_proto_rawDescOnce.Do(func() {
		file_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_proto_rawDescData)
	})
	return file_im_proto_rawDescData
}

var file_im_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_im_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: im.basic.Empty
	(*SendMessageReq)(nil),           // 1: im.basic.SendMessageReq
	(*SendMessageToMultipleReq)(nil), // 2: im.basic.SendMessageToMultipleReq
	(*SendMessageReply)(nil),         // 3: im.basic.SendMessageReply
	(*OnlinesReply)(nil),             // 4: im.basic.OnlinesReply
	(*BroadcastReq)(nil),             // 5: im.basic.BroadcastReq
	(*BroadcastReply)(nil),           // 6: im.basic.BroadcastReply
	(*PushToClient)(nil),             // 7: im.basic.PushToClient
	(*BroadcastByWTIReq)(nil),        // 8: im.basic.BroadcastByWTIReq
	(*WTIDistributeReply)(nil),       // 9: im.basic.WTIDistributeReply
	(*WTIDistribute)(nil),            // 10: im.basic.WTIDistribute
	nil,                              // 11: im.basic.BroadcastByWTIReq.DataEntry
	nil,                              // 12: im.basic.WTIDistributeReply.DataEntry
}
var file_im_proto_depIdxs = []int32{
	11, // 0: im.basic.BroadcastByWTIReq.data:type_name -> im.basic.BroadcastByWTIReq.DataEntry
	12, // 1: im.basic.WTIDistributeReply.data:type_name -> im.basic.WTIDistributeReply.DataEntry
	10, // 2: im.basic.WTIDistributeReply.DataEntry.value:type_name -> im.basic.WTIDistribute
	0,  // 3: im.basic.Basic.Ping:input_type -> im.basic.Empty
	0,  // 4: im.basic.Basic.Onliens:input_type -> im.basic.Empty
	1,  // 5: im.basic.Basic.SendMessage:input_type -> im.basic.SendMessageReq
	2,  // 6: im.basic.Basic.SendMessageToMultiple:input_type -> im.basic.SendMessageToMultipleReq
	5,  // 7: im.basic.Basic.Broadcast:input_type -> im.basic.BroadcastReq
	8,  // 8: im.basic.Basic.WTIBroadcast:input_type -> im.basic.BroadcastByWTIReq
	0,  // 9: im.basic.Basic.WTIDistribute:input_type -> im.basic.Empty
	0,  // 10: im.basic.Basic.Ping:output_type -> im.basic.Empty
	4,  // 11: im.basic.Basic.Onliens:output_type -> im.basic.OnlinesReply
	3,  // 12: im.basic.Basic.SendMessage:output_type -> im.basic.SendMessageReply
	3,  // 13: im.basic.Basic.SendMessageToMultiple:output_type -> im.basic.SendMessageReply
	6,  // 14: im.basic.Basic.Broadcast:output_type -> im.basic.BroadcastReply
	6,  // 15: im.basic.Basic.WTIBroadcast:output_type -> im.basic.BroadcastReply
	9,  // 16: im.basic.Basic.WTIDistribute:output_type -> im.basic.WTIDistributeReply
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_im_proto_init() }
func file_im_proto_init() {
	if File_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReq); i {
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
		file_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageToMultipleReq); i {
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
		file_im_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReply); i {
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
		file_im_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinesReply); i {
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
		file_im_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReq); i {
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
		file_im_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReply); i {
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
		file_im_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushToClient); i {
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
		file_im_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastByWTIReq); i {
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
		file_im_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WTIDistributeReply); i {
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
		file_im_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WTIDistribute); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_proto_goTypes,
		DependencyIndexes: file_im_proto_depIdxs,
		MessageInfos:      file_im_proto_msgTypes,
	}.Build()
	File_im_proto = out.File
	file_im_proto_rawDesc = nil
	file_im_proto_goTypes = nil
	file_im_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BasicClient is the client API for Basic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasicClient interface {
	// Ping Service
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// ????????????????????????
	Onliens(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OnlinesReply, error)
	// ???????????????????????????
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageReply, error)
	// ???????????????????????????
	SendMessageToMultiple(ctx context.Context, in *SendMessageToMultipleReq, opts ...grpc.CallOption) (*SendMessageReply, error)
	// ??????
	Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*BroadcastReply, error)
	// ???????????? v1.01
	// ??????WTI????????????????????????????????????????????????
	WTIBroadcast(ctx context.Context, in *BroadcastByWTIReq, opts ...grpc.CallOption) (*BroadcastReply, error)
	// ???????????????tag?????????????????????
	WTIDistribute(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WTIDistributeReply, error)
}

type basicClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicClient(cc grpc.ClientConnInterface) BasicClient {
	return &basicClient{cc}
}

func (c *basicClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) Onliens(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OnlinesReply, error) {
	out := new(OnlinesReply)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/Onliens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageReply, error) {
	out := new(SendMessageReply)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) SendMessageToMultiple(ctx context.Context, in *SendMessageToMultipleReq, opts ...grpc.CallOption) (*SendMessageReply, error) {
	out := new(SendMessageReply)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/SendMessageToMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*BroadcastReply, error) {
	out := new(BroadcastReply)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) WTIBroadcast(ctx context.Context, in *BroadcastByWTIReq, opts ...grpc.CallOption) (*BroadcastReply, error) {
	out := new(BroadcastReply)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/WTIBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) WTIDistribute(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WTIDistributeReply, error) {
	out := new(WTIDistributeReply)
	err := c.cc.Invoke(ctx, "/im.basic.Basic/WTIDistribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicServer is the server API for Basic service.
type BasicServer interface {
	// Ping Service
	Ping(context.Context, *Empty) (*Empty, error)
	// ????????????????????????
	Onliens(context.Context, *Empty) (*OnlinesReply, error)
	// ???????????????????????????
	SendMessage(context.Context, *SendMessageReq) (*SendMessageReply, error)
	// ???????????????????????????
	SendMessageToMultiple(context.Context, *SendMessageToMultipleReq) (*SendMessageReply, error)
	// ??????
	Broadcast(context.Context, *BroadcastReq) (*BroadcastReply, error)
	// ???????????? v1.01
	// ??????WTI????????????????????????????????????????????????
	WTIBroadcast(context.Context, *BroadcastByWTIReq) (*BroadcastReply, error)
	// ???????????????tag?????????????????????
	WTIDistribute(context.Context, *Empty) (*WTIDistributeReply, error)
}

// UnimplementedBasicServer can be embedded to have forward compatible implementations.
type UnimplementedBasicServer struct {
}

func (*UnimplementedBasicServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedBasicServer) Onliens(context.Context, *Empty) (*OnlinesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Onliens not implemented")
}
func (*UnimplementedBasicServer) SendMessage(context.Context, *SendMessageReq) (*SendMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedBasicServer) SendMessageToMultiple(context.Context, *SendMessageToMultipleReq) (*SendMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToMultiple not implemented")
}
func (*UnimplementedBasicServer) Broadcast(context.Context, *BroadcastReq) (*BroadcastReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (*UnimplementedBasicServer) WTIBroadcast(context.Context, *BroadcastByWTIReq) (*BroadcastReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WTIBroadcast not implemented")
}
func (*UnimplementedBasicServer) WTIDistribute(context.Context, *Empty) (*WTIDistributeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WTIDistribute not implemented")
}

func RegisterBasicServer(s *grpc.Server, srv BasicServer) {
	s.RegisterService(&_Basic_serviceDesc, srv)
}

func _Basic_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_Onliens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).Onliens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/Onliens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).Onliens(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_SendMessageToMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToMultipleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).SendMessageToMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/SendMessageToMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).SendMessageToMultiple(ctx, req.(*SendMessageToMultipleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).Broadcast(ctx, req.(*BroadcastReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_WTIBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastByWTIReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).WTIBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/WTIBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).WTIBroadcast(ctx, req.(*BroadcastByWTIReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_WTIDistribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).WTIDistribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.basic.Basic/WTIDistribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).WTIDistribute(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Basic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "im.basic.Basic",
	HandlerType: (*BasicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Basic_Ping_Handler,
		},
		{
			MethodName: "Onliens",
			Handler:    _Basic_Onliens_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Basic_SendMessage_Handler,
		},
		{
			MethodName: "SendMessageToMultiple",
			Handler:    _Basic_SendMessageToMultiple_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Basic_Broadcast_Handler,
		},
		{
			MethodName: "WTIBroadcast",
			Handler:    _Basic_WTIBroadcast_Handler,
		},
		{
			MethodName: "WTIDistribute",
			Handler:    _Basic_WTIDistribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im.proto",
}
