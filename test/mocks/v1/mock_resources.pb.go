//
//Syntax Comments
//Syntax Comments a

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-kit/test/mocks/api/v1/mock_resources.proto

//
//package Comments
//package Comments a

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

//
//Mock resources for goofin off
type MockResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        *core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Metadata      *core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Data          string         `protobuf:"bytes,1,opt,name=data,json=data.json,proto3" json:"data,omitempty"`
	SomeDumbField string         `protobuf:"bytes,100,opt,name=some_dumb_field,json=someDumbField,proto3" json:"some_dumb_field,omitempty"`
	// Types that are assignable to TestOneofFields:
	//	*MockResource_OneofOne
	//	*MockResource_OneofTwo
	TestOneofFields isMockResource_TestOneofFields `protobuf_oneof:"test_oneof_fields"`
}

func (x *MockResource) Reset() {
	*x = MockResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockResource) ProtoMessage() {}

func (x *MockResource) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockResource.ProtoReflect.Descriptor instead.
func (*MockResource) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescGZIP(), []int{0}
}

func (x *MockResource) GetStatus() *core.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MockResource) GetMetadata() *core.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MockResource) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *MockResource) GetSomeDumbField() string {
	if x != nil {
		return x.SomeDumbField
	}
	return ""
}

func (m *MockResource) GetTestOneofFields() isMockResource_TestOneofFields {
	if m != nil {
		return m.TestOneofFields
	}
	return nil
}

func (x *MockResource) GetOneofOne() string {
	if x, ok := x.GetTestOneofFields().(*MockResource_OneofOne); ok {
		return x.OneofOne
	}
	return ""
}

func (x *MockResource) GetOneofTwo() bool {
	if x, ok := x.GetTestOneofFields().(*MockResource_OneofTwo); ok {
		return x.OneofTwo
	}
	return false
}

type isMockResource_TestOneofFields interface {
	isMockResource_TestOneofFields()
}

type MockResource_OneofOne struct {
	OneofOne string `protobuf:"bytes,3,opt,name=oneof_one,json=oneofOne,proto3,oneof"`
}

type MockResource_OneofTwo struct {
	OneofTwo bool `protobuf:"varint,2,opt,name=oneof_two,json=oneofTwo,proto3,oneof"`
}

func (*MockResource_OneofOne) isMockResource_TestOneofFields() {}

func (*MockResource_OneofTwo) isMockResource_TestOneofFields() {}

type FakeResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    uint32         `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Metadata *core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *FakeResource) Reset() {
	*x = FakeResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FakeResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FakeResource) ProtoMessage() {}

func (x *FakeResource) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FakeResource.ProtoReflect.Descriptor instead.
func (*FakeResource) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescGZIP(), []int{1}
}

func (x *FakeResource) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FakeResource) GetMetadata() *core.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

//
//@solo-kit:xds-service=MockXdsResourceDiscoveryService
//@solo-kit:resource.no_references
type MockXdsResourceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @solo-kit:resource.name
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *MockXdsResourceConfig) Reset() {
	*x = MockXdsResourceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockXdsResourceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockXdsResourceConfig) ProtoMessage() {}

func (x *MockXdsResourceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockXdsResourceConfig.ProtoReflect.Descriptor instead.
func (*MockXdsResourceConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescGZIP(), []int{2}
}

func (x *MockXdsResourceConfig) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b,
	0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xb8, 0xf5,
	0x04, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x6f, 0x6d, 0x65, 0x5f,
	0x64, 0x75, 0x6d, 0x62, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xb8, 0xf5, 0x04, 0x01, 0x52, 0x0d, 0x73, 0x6f, 0x6d, 0x65, 0x44, 0x75, 0x6d, 0x62,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x74, 0x77,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x54, 0x77, 0x6f, 0x3a, 0x13, 0x82, 0xf1, 0x04, 0x04, 0x0a, 0x02, 0x6d, 0x6b, 0x82, 0xf1, 0x04,
	0x07, 0x12, 0x05, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x6d, 0x0a,
	0x0c, 0x46, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x13, 0x82, 0xf1, 0x04, 0x04, 0x0a, 0x02, 0x66,
	0x6b, 0x82, 0xf1, 0x04, 0x07, 0x12, 0x05, 0x66, 0x61, 0x6b, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x15,
	0x4d, 0x6f, 0x63, 0x6b, 0x58, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0x86, 0x03,
	0x0a, 0x1f, 0x4d, 0x6f, 0x63, 0x6b, 0x58, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x6f, 0x63, 0x6b, 0x58,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x6d, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x4d, 0x6f, 0x63, 0x6b, 0x58, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x8d, 0x01, 0x0a, 0x1a, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x4d, 0x6f, 0x63, 0x6b, 0x58, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23,
	0x2f, 0x76, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x6d, 0x6f,
	0x63, 0x6b, 0x78, 0x64, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x33, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x73,
	0x2f, 0x76, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescData = file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDesc
)

func file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDescData
}

var file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_goTypes = []interface{}{
	(*MockResource)(nil),              // 0: testing.solo.io.MockResource
	(*FakeResource)(nil),              // 1: testing.solo.io.FakeResource
	(*MockXdsResourceConfig)(nil),     // 2: testing.solo.io.MockXdsResourceConfig
	(*core.Status)(nil),               // 3: core.solo.io.Status
	(*core.Metadata)(nil),             // 4: core.solo.io.Metadata
	(*v2.DiscoveryRequest)(nil),       // 5: envoy.api.v2.DiscoveryRequest
	(*v2.DeltaDiscoveryRequest)(nil),  // 6: envoy.api.v2.DeltaDiscoveryRequest
	(*v2.DiscoveryResponse)(nil),      // 7: envoy.api.v2.DiscoveryResponse
	(*v2.DeltaDiscoveryResponse)(nil), // 8: envoy.api.v2.DeltaDiscoveryResponse
}
var file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_depIdxs = []int32{
	3, // 0: testing.solo.io.MockResource.status:type_name -> core.solo.io.Status
	4, // 1: testing.solo.io.MockResource.metadata:type_name -> core.solo.io.Metadata
	4, // 2: testing.solo.io.FakeResource.metadata:type_name -> core.solo.io.Metadata
	5, // 3: testing.solo.io.MockXdsResourceDiscoveryService.StreamMockXdsResourceConfig:input_type -> envoy.api.v2.DiscoveryRequest
	6, // 4: testing.solo.io.MockXdsResourceDiscoveryService.DeltaMockXdsResourceConfig:input_type -> envoy.api.v2.DeltaDiscoveryRequest
	5, // 5: testing.solo.io.MockXdsResourceDiscoveryService.FetchMockXdsResourceConfig:input_type -> envoy.api.v2.DiscoveryRequest
	7, // 6: testing.solo.io.MockXdsResourceDiscoveryService.StreamMockXdsResourceConfig:output_type -> envoy.api.v2.DiscoveryResponse
	8, // 7: testing.solo.io.MockXdsResourceDiscoveryService.DeltaMockXdsResourceConfig:output_type -> envoy.api.v2.DeltaDiscoveryResponse
	7, // 8: testing.solo.io.MockXdsResourceDiscoveryService.FetchMockXdsResourceConfig:output_type -> envoy.api.v2.DiscoveryResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_init() }
func file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_init() {
	if File_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockResource); i {
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
		file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FakeResource); i {
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
		file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockXdsResourceConfig); i {
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
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MockResource_OneofOne)(nil),
		(*MockResource_OneofTwo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto = out.File
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_rawDesc = nil
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_goTypes = nil
	file_github_com_solo_io_solo_kit_test_mocks_api_v1_mock_resources_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MockXdsResourceDiscoveryServiceClient is the client API for MockXdsResourceDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MockXdsResourceDiscoveryServiceClient interface {
	StreamMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigClient, error)
	DeltaMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigClient, error)
	FetchMockXdsResourceConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type mockXdsResourceDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMockXdsResourceDiscoveryServiceClient(cc grpc.ClientConnInterface) MockXdsResourceDiscoveryServiceClient {
	return &mockXdsResourceDiscoveryServiceClient{cc}
}

func (c *mockXdsResourceDiscoveryServiceClient) StreamMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MockXdsResourceDiscoveryService_serviceDesc.Streams[0], "/testing.solo.io.MockXdsResourceDiscoveryService/StreamMockXdsResourceConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient{stream}
	return x, nil
}

type MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient struct {
	grpc.ClientStream
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mockXdsResourceDiscoveryServiceClient) DeltaMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MockXdsResourceDiscoveryService_serviceDesc.Streams[1], "/testing.solo.io.MockXdsResourceDiscoveryService/DeltaMockXdsResourceConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient{stream}
	return x, nil
}

type MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient struct {
	grpc.ClientStream
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mockXdsResourceDiscoveryServiceClient) FetchMockXdsResourceConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/testing.solo.io.MockXdsResourceDiscoveryService/FetchMockXdsResourceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MockXdsResourceDiscoveryServiceServer is the server API for MockXdsResourceDiscoveryService service.
type MockXdsResourceDiscoveryServiceServer interface {
	StreamMockXdsResourceConfig(MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigServer) error
	DeltaMockXdsResourceConfig(MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigServer) error
	FetchMockXdsResourceConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedMockXdsResourceDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMockXdsResourceDiscoveryServiceServer struct {
}

func (*UnimplementedMockXdsResourceDiscoveryServiceServer) StreamMockXdsResourceConfig(MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMockXdsResourceConfig not implemented")
}
func (*UnimplementedMockXdsResourceDiscoveryServiceServer) DeltaMockXdsResourceConfig(MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaMockXdsResourceConfig not implemented")
}
func (*UnimplementedMockXdsResourceDiscoveryServiceServer) FetchMockXdsResourceConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMockXdsResourceConfig not implemented")
}

func RegisterMockXdsResourceDiscoveryServiceServer(s *grpc.Server, srv MockXdsResourceDiscoveryServiceServer) {
	s.RegisterService(&_MockXdsResourceDiscoveryService_serviceDesc, srv)
}

func _MockXdsResourceDiscoveryService_StreamMockXdsResourceConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MockXdsResourceDiscoveryServiceServer).StreamMockXdsResourceConfig(&mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer{stream})
}

type MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer struct {
	grpc.ServerStream
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MockXdsResourceDiscoveryServiceServer).DeltaMockXdsResourceConfig(&mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer{stream})
}

type MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer struct {
	grpc.ServerStream
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MockXdsResourceDiscoveryService_FetchMockXdsResourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockXdsResourceDiscoveryServiceServer).FetchMockXdsResourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testing.solo.io.MockXdsResourceDiscoveryService/FetchMockXdsResourceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockXdsResourceDiscoveryServiceServer).FetchMockXdsResourceConfig(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MockXdsResourceDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testing.solo.io.MockXdsResourceDiscoveryService",
	HandlerType: (*MockXdsResourceDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMockXdsResourceConfig",
			Handler:    _MockXdsResourceDiscoveryService_FetchMockXdsResourceConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMockXdsResourceConfig",
			Handler:       _MockXdsResourceDiscoveryService_StreamMockXdsResourceConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaMockXdsResourceConfig",
			Handler:       _MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/solo-io/solo-kit/test/mocks/api/v1/mock_resources.proto",
}
