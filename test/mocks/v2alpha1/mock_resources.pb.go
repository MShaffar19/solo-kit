// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-kit/test/mocks/api/v2alpha1/mock_resources.proto

package v2alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//The best mock resource you ever done seen
type MockResource struct {
	Status   core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to WeStuckItInAOneof:
	//	*MockResource_SomeDumbField
	//	*MockResource_Data
	WeStuckItInAOneof isMockResource_WeStuckItInAOneof `protobuf_oneof:"we_stuck_it_in_a_oneof"`
	// Types that are valid to be assigned to TestOneofFields:
	//	*MockResource_OneofOne
	//	*MockResource_OneofTwo
	TestOneofFields      isMockResource_TestOneofFields `protobuf_oneof:"test_oneof_fields"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MockResource) Reset()         { *m = MockResource{} }
func (m *MockResource) String() string { return proto.CompactTextString(m) }
func (*MockResource) ProtoMessage()    {}
func (*MockResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc86c81bab68fcb, []int{0}
}
func (m *MockResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockResource.Unmarshal(m, b)
}
func (m *MockResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockResource.Marshal(b, m, deterministic)
}
func (m *MockResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockResource.Merge(m, src)
}
func (m *MockResource) XXX_Size() int {
	return xxx_messageInfo_MockResource.Size(m)
}
func (m *MockResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MockResource.DiscardUnknown(m)
}

var xxx_messageInfo_MockResource proto.InternalMessageInfo

type isMockResource_WeStuckItInAOneof interface {
	isMockResource_WeStuckItInAOneof()
	Equal(interface{}) bool
}
type isMockResource_TestOneofFields interface {
	isMockResource_TestOneofFields()
	Equal(interface{}) bool
}

type MockResource_SomeDumbField struct {
	SomeDumbField string `protobuf:"bytes,100,opt,name=some_dumb_field,json=someDumbField,proto3,oneof"`
}
type MockResource_Data struct {
	Data string `protobuf:"bytes,1,opt,name=data,json=data.json,proto3,oneof"`
}
type MockResource_OneofOne struct {
	OneofOne string `protobuf:"bytes,3,opt,name=oneof_one,json=oneofOne,proto3,oneof"`
}
type MockResource_OneofTwo struct {
	OneofTwo bool `protobuf:"varint,2,opt,name=oneof_two,json=oneofTwo,proto3,oneof"`
}

func (*MockResource_SomeDumbField) isMockResource_WeStuckItInAOneof() {}
func (*MockResource_Data) isMockResource_WeStuckItInAOneof()          {}
func (*MockResource_OneofOne) isMockResource_TestOneofFields()        {}
func (*MockResource_OneofTwo) isMockResource_TestOneofFields()        {}

func (m *MockResource) GetWeStuckItInAOneof() isMockResource_WeStuckItInAOneof {
	if m != nil {
		return m.WeStuckItInAOneof
	}
	return nil
}
func (m *MockResource) GetTestOneofFields() isMockResource_TestOneofFields {
	if m != nil {
		return m.TestOneofFields
	}
	return nil
}

func (m *MockResource) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MockResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MockResource) GetSomeDumbField() string {
	if x, ok := m.GetWeStuckItInAOneof().(*MockResource_SomeDumbField); ok {
		return x.SomeDumbField
	}
	return ""
}

func (m *MockResource) GetData() string {
	if x, ok := m.GetWeStuckItInAOneof().(*MockResource_Data); ok {
		return x.Data
	}
	return ""
}

func (m *MockResource) GetOneofOne() string {
	if x, ok := m.GetTestOneofFields().(*MockResource_OneofOne); ok {
		return x.OneofOne
	}
	return ""
}

func (m *MockResource) GetOneofTwo() bool {
	if x, ok := m.GetTestOneofFields().(*MockResource_OneofTwo); ok {
		return x.OneofTwo
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MockResource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MockResource_OneofMarshaler, _MockResource_OneofUnmarshaler, _MockResource_OneofSizer, []interface{}{
		(*MockResource_SomeDumbField)(nil),
		(*MockResource_Data)(nil),
		(*MockResource_OneofOne)(nil),
		(*MockResource_OneofTwo)(nil),
	}
}

func _MockResource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MockResource)
	// we_stuck_it_in_a_oneof
	switch x := m.WeStuckItInAOneof.(type) {
	case *MockResource_SomeDumbField:
		_ = b.EncodeVarint(100<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.SomeDumbField)
	case *MockResource_Data:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Data)
	case nil:
	default:
		return fmt.Errorf("MockResource.WeStuckItInAOneof has unexpected type %T", x)
	}
	// test_oneof_fields
	switch x := m.TestOneofFields.(type) {
	case *MockResource_OneofOne:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.OneofOne)
	case *MockResource_OneofTwo:
		t := uint64(0)
		if x.OneofTwo {
			t = 1
		}
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("MockResource.TestOneofFields has unexpected type %T", x)
	}
	return nil
}

func _MockResource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MockResource)
	switch tag {
	case 100: // we_stuck_it_in_a_oneof.some_dumb_field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.WeStuckItInAOneof = &MockResource_SomeDumbField{x}
		return true, err
	case 1: // we_stuck_it_in_a_oneof.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.WeStuckItInAOneof = &MockResource_Data{x}
		return true, err
	case 3: // test_oneof_fields.oneof_one
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.TestOneofFields = &MockResource_OneofOne{x}
		return true, err
	case 2: // test_oneof_fields.oneof_two
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestOneofFields = &MockResource_OneofTwo{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _MockResource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MockResource)
	// we_stuck_it_in_a_oneof
	switch x := m.WeStuckItInAOneof.(type) {
	case *MockResource_SomeDumbField:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.SomeDumbField)))
		n += len(x.SomeDumbField)
	case *MockResource_Data:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// test_oneof_fields
	switch x := m.TestOneofFields.(type) {
	case *MockResource_OneofOne:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OneofOne)))
		n += len(x.OneofOne)
	case *MockResource_OneofTwo:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*MockResource)(nil), "testing.solo.io.v2alpha1.MockResource")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-kit/test/mocks/api/v2alpha1/mock_resources.proto", fileDescriptor_bbc86c81bab68fcb)
}

var fileDescriptor_bbc86c81bab68fcb = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x8e, 0x9b, 0x40,
	0x10, 0xc6, 0x8d, 0x43, 0xfc, 0x67, 0x13, 0x2b, 0x0a, 0x8e, 0x12, 0x62, 0x29, 0xb1, 0x95, 0x8a,
	0x26, 0xbb, 0xc2, 0x91, 0x9b, 0xa4, 0x43, 0x51, 0x44, 0x63, 0x45, 0x22, 0xa9, 0xae, 0x59, 0x2d,
	0xb0, 0xc6, 0x7b, 0x18, 0xc6, 0x62, 0x97, 0x73, 0x7b, 0xf2, 0xd3, 0xdc, 0xa3, 0xdc, 0x53, 0xb8,
	0xb8, 0xf2, 0x3a, 0xbf, 0xc1, 0x89, 0x05, 0xec, 0xe2, 0xa4, 0x93, 0x2b, 0x76, 0xe6, 0xfb, 0x7e,
	0xfb, 0x69, 0xd8, 0x41, 0x7e, 0x22, 0xd4, 0xba, 0x0c, 0x71, 0x04, 0x19, 0x91, 0xb0, 0x81, 0xef,
	0x02, 0xea, 0x6f, 0x2a, 0x14, 0x51, 0x5c, 0x2a, 0x92, 0x41, 0x94, 0x4a, 0xc2, 0xb6, 0x82, 0xdc,
	0xcc, 0xd9, 0x66, 0xbb, 0x66, 0xae, 0x6e, 0xd1, 0x82, 0x4b, 0x28, 0x8b, 0x88, 0x4b, 0xbc, 0x2d,
	0x40, 0x81, 0x65, 0x57, 0x6e, 0x91, 0x27, 0xb8, 0xc2, 0xb1, 0x00, 0xdc, 0xda, 0x27, 0xee, 0x4b,
	0x19, 0xfa, 0x62, 0x97, 0x64, 0x5c, 0xb1, 0x98, 0x29, 0x56, 0x5f, 0x36, 0x21, 0x17, 0x20, 0x52,
	0x31, 0x55, 0x36, 0xe9, 0x17, 0x65, 0xb4, 0x75, 0x83, 0x7c, 0x48, 0x20, 0x01, 0x7d, 0x24, 0xd5,
	0xa9, 0xee, 0x7e, 0x3b, 0x74, 0xd1, 0xdb, 0x25, 0x44, 0x69, 0xd0, 0x8c, 0x67, 0x2d, 0x50, 0xaf,
	0x4e, 0xb2, 0x7b, 0x33, 0xc3, 0x79, 0x33, 0xff, 0x84, 0x23, 0x28, 0xf8, 0x79, 0x4a, 0x17, 0xff,
	0xd3, 0xb2, 0x67, 0xde, 0x1f, 0xa6, 0x9d, 0xa0, 0x31, 0x5b, 0xbf, 0xd0, 0xa0, 0x9d, 0xc9, 0xee,
	0x6b, 0xf0, 0xf3, 0x33, 0x70, 0xd9, 0x18, 0x1a, 0xf4, 0x04, 0x58, 0x18, 0xbd, 0x93, 0x90, 0x71,
	0x1a, 0x97, 0x59, 0x48, 0x57, 0x82, 0x6f, 0x62, 0x3b, 0x9e, 0x19, 0xce, 0xd0, 0x33, 0x6f, 0x8f,
	0xa6, 0xe1, 0x77, 0x82, 0x51, 0x25, 0xff, 0x2e, 0xb3, 0xf0, 0x4f, 0x25, 0x5a, 0x0e, 0x32, 0x75,
	0x90, 0xa1, 0x4d, 0xa3, 0xc7, 0xc3, 0x74, 0xa8, 0x7f, 0xe6, 0xb5, 0x84, 0xdc, 0xef, 0x04, 0xe7,
	0xc2, 0xfa, 0x82, 0x86, 0x90, 0x73, 0x58, 0x51, 0xc8, 0xb9, 0xfd, 0xaa, 0xb2, 0xfb, 0x46, 0x30,
	0xd0, 0xad, 0xbf, 0x39, 0x3f, 0xcb, 0x6a, 0x07, 0x76, 0x77, 0x66, 0x38, 0x83, 0x93, 0xfc, 0x7f,
	0x07, 0x3f, 0xc7, 0xfb, 0xa3, 0x69, 0xa2, 0x6e, 0x96, 0xee, 0x8f, 0x66, 0xdf, 0x7a, 0xad, 0x57,
	0xc3, 0xb3, 0xd1, 0xc7, 0x1d, 0xa7, 0x52, 0x95, 0x51, 0x4a, 0x85, 0xa2, 0x22, 0xa7, 0x8c, 0x6a,
	0xc2, 0x1b, 0xa3, 0xf7, 0xd5, 0x52, 0xd4, 0x55, 0x3d, 0x87, 0xf4, 0x16, 0x77, 0x0f, 0x5f, 0x8d,
	0x2b, 0x72, 0xe1, 0xde, 0xb5, 0x4b, 0x14, 0xf6, 0xf4, 0xf3, 0xfc, 0x78, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x53, 0x52, 0x12, 0xc3, 0xb1, 0x02, 0x00, 0x00,
}

func (this *MockResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource)
	if !ok {
		that2, ok := that.(MockResource)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.WeStuckItInAOneof == nil {
		if this.WeStuckItInAOneof != nil {
			return false
		}
	} else if this.WeStuckItInAOneof == nil {
		return false
	} else if !this.WeStuckItInAOneof.Equal(that1.WeStuckItInAOneof) {
		return false
	}
	if that1.TestOneofFields == nil {
		if this.TestOneofFields != nil {
			return false
		}
	} else if this.TestOneofFields == nil {
		return false
	} else if !this.TestOneofFields.Equal(that1.TestOneofFields) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MockResource_SomeDumbField) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_SomeDumbField)
	if !ok {
		that2, ok := that.(MockResource_SomeDumbField)
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
	if this.SomeDumbField != that1.SomeDumbField {
		return false
	}
	return true
}
func (this *MockResource_Data) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_Data)
	if !ok {
		that2, ok := that.(MockResource_Data)
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
	if this.Data != that1.Data {
		return false
	}
	return true
}
func (this *MockResource_OneofOne) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_OneofOne)
	if !ok {
		that2, ok := that.(MockResource_OneofOne)
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
	if this.OneofOne != that1.OneofOne {
		return false
	}
	return true
}
func (this *MockResource_OneofTwo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_OneofTwo)
	if !ok {
		that2, ok := that.(MockResource_OneofTwo)
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
	if this.OneofTwo != that1.OneofTwo {
		return false
	}
	return true
}
