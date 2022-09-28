// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: demo.proto

package pb

import (
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

//model
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Test) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Test) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

//req 、resp
type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Page     int64  `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize int64  `protobuf:"varint,4,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *ListReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ListReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestInfo   []*Test `protobuf:"bytes,1,rep,name=testInfo,proto3" json:"testInfo,omitempty"`
	TotalCount int64   `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{2}
}

func (x *ListResp) GetTestInfo() []*Test {
	if x != nil {
		return x.TestInfo
	}
	return nil
}

func (x *ListResp) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{3}
}

func (x *AddReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *AddReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type AddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestInfo *Test `protobuf:"bytes,1,opt,name=testInfo,proto3" json:"testInfo,omitempty"`
}

func (x *AddResp) Reset() {
	*x = AddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResp) ProtoMessage() {}

func (x *AddResp) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResp.ProtoReflect.Descriptor instead.
func (*AddResp) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{4}
}

func (x *AddResp) GetTestInfo() *Test {
	if x != nil {
		return x.TestInfo
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UpdateReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestInfo *Test `protobuf:"bytes,1,opt,name=testInfo,proto3" json:"testInfo,omitempty"`
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResp) GetTestInfo() *Test {
	if x != nil {
		return x.TestInfo
	}
	return nil
}

type DelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelReq) Reset() {
	*x = DelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelReq) ProtoMessage() {}

func (x *DelReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelReq.ProtoReflect.Descriptor instead.
func (*DelReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{7}
}

func (x *DelReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DelResp) Reset() {
	*x = DelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelResp) ProtoMessage() {}

func (x *DelResp) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelResp.ProtoReflect.Descriptor instead.
func (*DelResp) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{8}
}

func (x *DelResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x4a, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x07,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x50, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a,
	0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4f, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x32, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x08, 0x74,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x18, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x92, 0x01, 0x0a, 0x04, 0x64, 0x65, 0x6d,
	0x6f, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a,
	0x03, 0x64, 0x65, 0x6c, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_demo_proto_goTypes = []interface{}{
	(*Test)(nil),       // 0: pb.Test
	(*ListReq)(nil),    // 1: pb.ListReq
	(*ListResp)(nil),   // 2: pb.ListResp
	(*AddReq)(nil),     // 3: pb.AddReq
	(*AddResp)(nil),    // 4: pb.AddResp
	(*UpdateReq)(nil),  // 5: pb.UpdateReq
	(*UpdateResp)(nil), // 6: pb.UpdateResp
	(*DelReq)(nil),     // 7: pb.delReq
	(*DelResp)(nil),    // 8: pb.delResp
}
var file_demo_proto_depIdxs = []int32{
	0, // 0: pb.ListResp.testInfo:type_name -> pb.Test
	0, // 1: pb.AddResp.testInfo:type_name -> pb.Test
	0, // 2: pb.UpdateResp.testInfo:type_name -> pb.Test
	1, // 3: pb.demo.list:input_type -> pb.ListReq
	3, // 4: pb.demo.add:input_type -> pb.AddReq
	5, // 5: pb.demo.updata:input_type -> pb.UpdateReq
	7, // 6: pb.demo.del:input_type -> pb.delReq
	2, // 7: pb.demo.list:output_type -> pb.ListResp
	4, // 8: pb.demo.add:output_type -> pb.AddResp
	6, // 9: pb.demo.updata:output_type -> pb.UpdateResp
	8, // 10: pb.demo.del:output_type -> pb.delResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResp); i {
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
		file_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResp); i {
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
		file_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
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
		file_demo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelReq); i {
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
		file_demo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelResp); i {
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
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}