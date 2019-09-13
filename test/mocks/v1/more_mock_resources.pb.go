// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-kit/test/mocks/api/v1/more_mock_resources.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
//Description of the AnotherMockResource
type AnotherMockResource struct {
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	Status   core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	// comments that go above the basic field in our docs
	BasicField           string   `protobuf:"bytes,2,opt,name=basic_field,json=basicField,proto3" json:"basic_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnotherMockResource) Reset()         { *m = AnotherMockResource{} }
func (m *AnotherMockResource) String() string { return proto.CompactTextString(m) }
func (*AnotherMockResource) ProtoMessage()    {}
func (*AnotherMockResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3005c9d7690ad701, []int{0}
}
func (m *AnotherMockResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnotherMockResource.Unmarshal(m, b)
}
func (m *AnotherMockResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnotherMockResource.Marshal(b, m, deterministic)
}
func (m *AnotherMockResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnotherMockResource.Merge(m, src)
}
func (m *AnotherMockResource) XXX_Size() int {
	return xxx_messageInfo_AnotherMockResource.Size(m)
}
func (m *AnotherMockResource) XXX_DiscardUnknown() {
	xxx_messageInfo_AnotherMockResource.DiscardUnknown(m)
}

var xxx_messageInfo_AnotherMockResource proto.InternalMessageInfo

func (m *AnotherMockResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *AnotherMockResource) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *AnotherMockResource) GetBasicField() string {
	if m != nil {
		return m.BasicField
	}
	return ""
}

type ClusterResource struct {
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	Status   core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	// comments that go above the basic field in our docs
	BasicField           string   `protobuf:"bytes,2,opt,name=basic_field,json=basicField,proto3" json:"basic_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterResource) Reset()         { *m = ClusterResource{} }
func (m *ClusterResource) String() string { return proto.CompactTextString(m) }
func (*ClusterResource) ProtoMessage()    {}
func (*ClusterResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3005c9d7690ad701, []int{1}
}
func (m *ClusterResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResource.Unmarshal(m, b)
}
func (m *ClusterResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResource.Marshal(b, m, deterministic)
}
func (m *ClusterResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResource.Merge(m, src)
}
func (m *ClusterResource) XXX_Size() int {
	return xxx_messageInfo_ClusterResource.Size(m)
}
func (m *ClusterResource) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResource.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResource proto.InternalMessageInfo

func (m *ClusterResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *ClusterResource) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *ClusterResource) GetBasicField() string {
	if m != nil {
		return m.BasicField
	}
	return ""
}

func init() {
	proto.RegisterType((*AnotherMockResource)(nil), "testing.solo.io.AnotherMockResource")
	proto.RegisterType((*ClusterResource)(nil), "testing.solo.io.ClusterResource")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-kit/test/mocks/api/v1/more_mock_resources.proto", fileDescriptor_3005c9d7690ad701)
}

var fileDescriptor_3005c9d7690ad701 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0x87, 0x3d, 0xa8, 0x44, 0x8e, 0x01, 0x53, 0x09, 0x69, 0x18, 0x84, 0x60, 0x4c, 0xd0, 0x84,
	0x5e, 0xc0, 0xc5, 0xb8, 0x81, 0x89, 0x4e, 0x2c, 0xb8, 0xb9, 0x90, 0xe3, 0x38, 0xcb, 0x85, 0xb6,
	0x2f, 0xb9, 0xbb, 0x36, 0xb2, 0xf2, 0x69, 0xfc, 0x0c, 0x4e, 0x6e, 0xfa, 0x01, 0x9c, 0x1d, 0xdc,
	0x1c, 0xf9, 0x06, 0xa6, 0xff, 0x88, 0x44, 0x63, 0x70, 0x72, 0xea, 0xb5, 0xbf, 0xf7, 0xe9, 0x7b,
	0xcf, 0xbd, 0x87, 0xaf, 0x1d, 0xa1, 0xa7, 0xc1, 0xd8, 0x66, 0xe0, 0x11, 0x05, 0x2e, 0xb4, 0x05,
	0x24, 0xcf, 0x99, 0xd0, 0x44, 0x73, 0xa5, 0x89, 0x07, 0x6c, 0xa6, 0x08, 0x9d, 0x0b, 0x12, 0x76,
	0x88, 0x07, 0x92, 0x8f, 0xa2, 0x2f, 0x23, 0xc9, 0x15, 0x04, 0x92, 0x71, 0x65, 0xcf, 0x25, 0x68,
	0x30, 0xcb, 0x51, 0xb1, 0xf0, 0x1d, 0x3b, 0xa2, 0x6d, 0x01, 0xb5, 0x8a, 0x03, 0x0e, 0xc4, 0x19,
	0x89, 0x56, 0x49, 0x59, 0xad, 0xf3, 0x5b, 0xbf, 0xac, 0x09, 0xd7, 0x74, 0x42, 0x35, 0x4d, 0x11,
	0xb2, 0x05, 0xa2, 0x34, 0xd5, 0x81, 0xfa, 0x43, 0x8f, 0xec, 0x3d, 0x45, 0x7a, 0x5f, 0x10, 0xee,
	0x87, 0xb0, 0x98, 0x4b, 0xb8, 0x5f, 0x90, 0x38, 0x64, 0x6d, 0x87, 0xfb, 0xed, 0x90, 0xba, 0x62,
	0x42, 0x35, 0x27, 0xdf, 0x16, 0xc9, 0x2f, 0x9a, 0xaf, 0x08, 0x1f, 0xf4, 0x7c, 0xd0, 0x53, 0x2e,
	0x07, 0xc0, 0x66, 0xc3, 0xf4, 0x7c, 0xcc, 0x73, 0xbc, 0x97, 0x09, 0x59, 0xa8, 0x81, 0x5a, 0xa5,
	0x6e, 0xd5, 0x66, 0x20, 0x79, 0x76, 0x50, 0xf6, 0x20, 0x4d, 0xfb, 0xc6, 0xcb, 0x5b, 0x7d, 0x67,
	0xb8, 0xae, 0x36, 0xbb, 0xb8, 0x90, 0x78, 0x59, 0x85, 0x98, 0xab, 0x6c, 0x72, 0x37, 0x71, 0x96,
	0x52, 0x69, 0xa5, 0x79, 0x8a, 0x4b, 0x63, 0xaa, 0x04, 0x1b, 0xdd, 0x09, 0xee, 0x4e, 0xac, 0x5c,
	0x03, 0xb5, 0x8a, 0xfd, 0xe2, 0xe3, 0xc7, 0x53, 0xde, 0x90, 0xb9, 0x06, 0x1a, 0xe2, 0x38, 0xbd,
	0x8a, 0xc2, 0x8b, 0xa3, 0xe5, 0xca, 0xd8, 0xc5, 0x79, 0xea, 0xc9, 0xe5, 0xca, 0xa8, 0x9a, 0x15,
	0x9a, 0xec, 0x3e, 0x1a, 0xef, 0x7a, 0xba, 0xcd, 0x67, 0x84, 0xcb, 0x97, 0x6e, 0xa0, 0x34, 0x97,
	0xff, 0xa4, 0x54, 0xff, 0x41, 0x69, 0xc3, 0xe3, 0x38, 0xf1, 0x60, 0x6e, 0xe4, 0x61, 0x9a, 0xfb,
	0x2c, 0xd9, 0xee, 0xda, 0x61, 0xb9, 0x32, 0x72, 0x16, 0xea, 0x93, 0x87, 0xf7, 0x43, 0x74, 0x7b,
	0xb2, 0xe5, 0x85, 0x0f, 0x3b, 0xe3, 0x42, 0x3c, 0xd8, 0xb3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x83, 0xd3, 0x8d, 0x6b, 0x24, 0x03, 0x00, 0x00,
}

func (this *AnotherMockResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AnotherMockResource)
	if !ok {
		that2, ok := that.(AnotherMockResource)
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
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if this.BasicField != that1.BasicField {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ClusterResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterResource)
	if !ok {
		that2, ok := that.(ClusterResource)
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
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if this.BasicField != that1.BasicField {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
