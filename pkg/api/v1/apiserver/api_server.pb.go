// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api_server.proto

package apiserver

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//GRPC stuff
type ReadRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TypeUrl              string   `protobuf:"bytes,3,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{0}
}
func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ReadRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

type ReadResponse struct {
	Resource             *types.Any `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{1}
}
func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetResource() *types.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

type WriteRequest struct {
	Resource             *types.Any `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	OverwriteExisting    bool       `protobuf:"varint,2,opt,name=overwrite_existing,json=overwriteExisting,proto3" json:"overwrite_existing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{2}
}
func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetResource() *types.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteRequest) GetOverwriteExisting() bool {
	if m != nil {
		return m.OverwriteExisting
	}
	return false
}

type WriteResponse struct {
	Resource             *types.Any `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{3}
}
func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetResource() *types.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

type DeleteRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TypeUrl              string   `protobuf:"bytes,3,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	IgnoreNotExist       bool     `protobuf:"varint,4,opt,name=ignore_not_exist,json=ignoreNotExist,proto3" json:"ignore_not_exist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{4}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeleteRequest) GetIgnoreNotExist() bool {
	if m != nil {
		return m.IgnoreNotExist
	}
	return false
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{5}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TypeUrl              string   `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{6}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

type ListResponse struct {
	ResourceList         []*types.Any `protobuf:"bytes,1,rep,name=resource_list,json=resourceList,proto3" json:"resource_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{7}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetResourceList() []*types.Any {
	if m != nil {
		return m.ResourceList
	}
	return nil
}

type WatchRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TypeUrl              string   `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{8}
}
func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WatchRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

type WatchResponse struct {
	ResourceList         []*types.Any `protobuf:"bytes,1,rep,name=resource_list,json=resourceList,proto3" json:"resource_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{9}
}
func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetResourceList() []*types.Any {
	if m != nil {
		return m.ResourceList
	}
	return nil
}

type RegisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{10}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0f2054e1069c43, []int{11}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReadRequest)(nil), "apiserver.api.v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "apiserver.api.v1.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "apiserver.api.v1.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "apiserver.api.v1.WriteResponse")
	proto.RegisterType((*DeleteRequest)(nil), "apiserver.api.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "apiserver.api.v1.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "apiserver.api.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "apiserver.api.v1.ListResponse")
	proto.RegisterType((*WatchRequest)(nil), "apiserver.api.v1.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "apiserver.api.v1.WatchResponse")
	proto.RegisterType((*RegisterRequest)(nil), "apiserver.api.v1.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "apiserver.api.v1.RegisterResponse")
}

func init() { proto.RegisterFile("api_server.proto", fileDescriptor_bc0f2054e1069c43) }

var fileDescriptor_bc0f2054e1069c43 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xac, 0xdb, 0xb4, 0x24, 0x5f, 0x9c, 0x92, 0xae, 0x7a, 0x48, 0x2d, 0x28, 0x61, 0x4f, 0xbd,
	0xd4, 0x4e, 0x8b, 0x84, 0xc4, 0x8d, 0x20, 0xa0, 0xa2, 0xfc, 0x1c, 0x5c, 0xa1, 0x4a, 0xbd, 0x58,
	0x4e, 0xf8, 0x70, 0x57, 0x75, 0xbd, 0x66, 0x77, 0x1d, 0xc8, 0x0b, 0xf0, 0x2c, 0x3c, 0x13, 0x47,
	0x9e, 0x04, 0x79, 0xd7, 0x76, 0x9c, 0xca, 0x29, 0x82, 0x72, 0xb2, 0xfd, 0xcd, 0x68, 0x76, 0x66,
	0x3d, 0x1f, 0xf4, 0xc3, 0x94, 0x05, 0x12, 0xc5, 0x0c, 0x85, 0x9b, 0x0a, 0xae, 0x38, 0xc9, 0x27,
	0xc5, 0x20, 0x4c, 0x99, 0x3b, 0x3b, 0x72, 0x76, 0x23, 0x1e, 0x71, 0x0d, 0x7a, 0xf9, 0x9b, 0xe1,
	0x39, 0x7b, 0x11, 0xe7, 0x51, 0x8c, 0x9e, 0xfe, 0x9a, 0x64, 0x9f, 0xbd, 0x30, 0x99, 0x1b, 0x88,
	0x5e, 0x40, 0xd7, 0xc7, 0xf0, 0x93, 0x8f, 0x5f, 0x32, 0x94, 0x8a, 0x10, 0x68, 0x25, 0xe1, 0x35,
	0x0e, 0xac, 0xa1, 0x75, 0xd0, 0xf1, 0xf5, 0x3b, 0x79, 0x00, 0x9d, 0xfc, 0x29, 0xd3, 0x70, 0x8a,
	0x83, 0x75, 0x0d, 0x2c, 0x06, 0x64, 0x0f, 0xda, 0x6a, 0x9e, 0x62, 0x90, 0x89, 0x78, 0xb0, 0xa1,
	0xc1, 0x7b, 0xf9, 0xf7, 0x47, 0x11, 0xd3, 0xe7, 0x60, 0x1b, 0x6d, 0x99, 0xf2, 0x44, 0x22, 0x19,
	0x41, 0x5b, 0xa0, 0xe4, 0x99, 0x98, 0x9a, 0x03, 0xba, 0xc7, 0xbb, 0xae, 0x71, 0xe6, 0x96, 0xce,
	0xdc, 0x71, 0x32, 0xf7, 0x2b, 0x16, 0xe5, 0x60, 0x9f, 0x0b, 0xa6, 0xb0, 0xb4, 0xf7, 0xd7, 0x0a,
	0xe4, 0x10, 0x08, 0x9f, 0xa1, 0xf8, 0x9a, 0xab, 0x04, 0xf8, 0x8d, 0x49, 0xc5, 0x92, 0x48, 0xa7,
	0x68, 0xfb, 0x3b, 0x15, 0xf2, 0xaa, 0x00, 0xe8, 0x18, 0x7a, 0xc5, 0x81, 0xff, 0xec, 0xf9, 0xbb,
	0x05, 0xbd, 0x97, 0x18, 0xe3, 0xc2, 0xf5, 0xff, 0xbc, 0x54, 0x72, 0x00, 0x7d, 0x16, 0x25, 0x5c,
	0x60, 0x90, 0x70, 0x65, 0x12, 0x0d, 0x5a, 0x3a, 0xce, 0xb6, 0x99, 0x7f, 0xe0, 0x4a, 0xc7, 0xa1,
	0x7d, 0xd8, 0x2e, 0x7d, 0x98, 0x30, 0xf4, 0x35, 0x74, 0xdf, 0x31, 0xa9, 0x4a, 0x5f, 0x4b, 0x1e,
	0xac, 0xdb, 0x3c, 0xac, 0x2f, 0xff, 0xd8, 0x37, 0x60, 0x1b, 0x9d, 0xe2, 0x92, 0x9e, 0x41, 0xaf,
	0x8c, 0x1f, 0xc4, 0xb9, 0x21, 0x6b, 0xb8, 0xb1, 0xf2, 0xa6, 0xec, 0x92, 0x9a, 0x4b, 0xd0, 0x13,
	0xb0, 0xcf, 0x43, 0x35, 0xbd, 0xbc, 0xb3, 0xa7, 0x53, 0xe8, 0x15, 0x42, 0x77, 0x37, 0xb5, 0x03,
	0xf7, 0x7d, 0x8c, 0x98, 0x54, 0x28, 0x0a, 0x5f, 0x94, 0x40, 0x7f, 0x31, 0x32, 0x27, 0x1c, 0xff,
	0xdc, 0x80, 0xce, 0x38, 0x65, 0x67, 0x7a, 0x03, 0xc9, 0x19, 0xb4, 0x4b, 0x06, 0x79, 0xec, 0xde,
	0xdc, 0x4c, 0xf7, 0x86, 0xa0, 0x43, 0x6f, 0xa3, 0x14, 0xff, 0x6b, 0x8d, 0x9c, 0x40, 0x2b, 0x5f,
	0x21, 0xf2, 0xb0, 0x89, 0x5d, 0xad, 0xad, 0xb3, 0xbf, 0x0a, 0xae, 0x84, 0x4e, 0x61, 0x53, 0x17,
	0x9b, 0x34, 0x50, 0xeb, 0x2b, 0xe6, 0x3c, 0x5a, 0x89, 0x57, 0x5a, 0xef, 0x61, 0xcb, 0x14, 0x8b,
	0x34, 0x90, 0x97, 0xaa, 0xef, 0x0c, 0x57, 0x13, 0xea, 0x19, 0xf3, 0x5b, 0x6f, 0xca, 0x58, 0x6b,
	0x6b, 0x53, 0xc6, 0x7a, 0x09, 0xe9, 0x1a, 0x79, 0x0b, 0x9b, 0xba, 0x02, 0x8d, 0x19, 0x6b, 0x25,
	0xfb, 0xb3, 0xd4, 0xc8, 0x7a, 0xf1, 0xf4, 0xc7, 0xaf, 0x7d, 0xeb, 0x62, 0x14, 0x31, 0x75, 0x99,
	0x4d, 0xdc, 0x29, 0xbf, 0xf6, 0x24, 0x8f, 0xf9, 0x21, 0xe3, 0xe6, 0x79, 0xc5, 0x94, 0x97, 0x5e,
	0x45, 0x5e, 0x98, 0x32, 0x6f, 0x76, 0xe4, 0x55, 0x7a, 0x93, 0x2d, 0xdd, 0xab, 0x27, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x3d, 0xeb, 0x58, 0x16, 0xae, 0x05, 0x00, 0x00,
}

func (this *ReadRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReadRequest)
	if !ok {
		that2, ok := that.(ReadRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ReadResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReadResponse)
	if !ok {
		that2, ok := that.(ReadResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WriteRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WriteRequest)
	if !ok {
		that2, ok := that.(WriteRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if this.OverwriteExisting != that1.OverwriteExisting {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WriteResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WriteResponse)
	if !ok {
		that2, ok := that.(WriteResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeleteRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteRequest)
	if !ok {
		that2, ok := that.(DeleteRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.IgnoreNotExist != that1.IgnoreNotExist {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeleteResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteResponse)
	if !ok {
		that2, ok := that.(DeleteResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListRequest)
	if !ok {
		that2, ok := that.(ListRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListResponse)
	if !ok {
		that2, ok := that.(ListResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ResourceList) != len(that1.ResourceList) {
		return false
	}
	for i := range this.ResourceList {
		if !this.ResourceList[i].Equal(that1.ResourceList[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WatchRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WatchRequest)
	if !ok {
		that2, ok := that.(WatchRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WatchResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WatchResponse)
	if !ok {
		that2, ok := that.(WatchResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ResourceList) != len(that1.ResourceList) {
		return false
	}
	for i := range this.ResourceList {
		if !this.ResourceList[i].Equal(that1.ResourceList[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RegisterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisterRequest)
	if !ok {
		that2, ok := that.(RegisterRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RegisterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisterResponse)
	if !ok {
		that2, ok := that.(RegisterResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiServerClient is the client API for ApiServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiServerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ApiServer_WatchClient, error)
}

type apiServerClient struct {
	cc *grpc.ClientConn
}

func NewApiServerClient(cc *grpc.ClientConn) ApiServerClient {
	return &apiServerClient{cc}
}

func (c *apiServerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ApiServer_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApiServer_serviceDesc.Streams[0], "/apiserver.api.v1.ApiServer/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiServerWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApiServer_WatchClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type apiServerWatchClient struct {
	grpc.ClientStream
}

func (x *apiServerWatchClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApiServerServer is the server API for ApiServer service.
type ApiServerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Watch(*WatchRequest, ApiServer_WatchServer) error
}

// UnimplementedApiServerServer can be embedded to have forward compatible implementations.
type UnimplementedApiServerServer struct {
}

func (*UnimplementedApiServerServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedApiServerServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedApiServerServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedApiServerServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedApiServerServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedApiServerServer) Watch(req *WatchRequest, srv ApiServer_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterApiServerServer(s *grpc.Server, srv ApiServerServer) {
	s.RegisterService(&_ApiServer_serviceDesc, srv)
}

func _ApiServer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServerServer).Watch(m, &apiServerWatchServer{stream})
}

type ApiServer_WatchServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type apiServerWatchServer struct {
	grpc.ServerStream
}

func (x *apiServerWatchServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ApiServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.api.v1.ApiServer",
	HandlerType: (*ApiServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ApiServer_Register_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ApiServer_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ApiServer_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApiServer_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ApiServer_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ApiServer_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api_server.proto",
}
