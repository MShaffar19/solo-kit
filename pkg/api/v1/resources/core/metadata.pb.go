// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metadata.proto

package core // import "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// Metadata contains general properties of resources for purposes of versioning, annotating, and namespacing.
type Metadata struct {
	//
	// Name of the resource.
	//
	// Names must be unique and follow the following syntax rules:
	//
	// One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Namespace is used for the namespacing of resources.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// An opaque value that represents the internal version of this object that can
	// be used by clients to determine when objects have changed.
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty" testdiff:"ignore"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. Some resources contain `selectors` which
	// can be linked with other resources by their labels
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata.
	Annotations          map[string]string `protobuf:"bytes,6,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_91abd6aad2c35ae2, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Metadata) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *Metadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Metadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "core.solo.io.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "core.solo.io.Metadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.solo.io.Metadata.LabelsEntry")
}
func (this *Metadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metadata)
	if !ok {
		that2, ok := that.(Metadata)
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
	if this.ResourceVersion != that1.ResourceVersion {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Annotations) != len(that1.Annotations) {
		return false
	}
	for i := range this.Annotations {
		if this.Annotations[i] != that1.Annotations[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor_metadata_91abd6aad2c35ae2) }

var fileDescriptor_metadata_91abd6aad2c35ae2 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x6e, 0x5b, 0x6c, 0x2a, 0x5a, 0x43, 0x85, 0x50, 0x44, 0xcb, 0x5e, 0xec, 0xc5,
	0x04, 0x15, 0x41, 0x2b, 0x88, 0x16, 0x3c, 0x08, 0x7a, 0xe9, 0xc1, 0x83, 0x17, 0x49, 0xb7, 0xe9,
	0x1a, 0xba, 0xcd, 0x2c, 0x49, 0xba, 0xd0, 0x9b, 0x8f, 0xe3, 0x53, 0x79, 0xf0, 0x11, 0x7c, 0x02,
	0x49, 0xb6, 0x4b, 0x17, 0xc1, 0x83, 0xa7, 0x99, 0x64, 0xbe, 0xff, 0xcf, 0xcc, 0x04, 0xed, 0x2c,
	0x84, 0xe5, 0x53, 0x6e, 0x39, 0xcd, 0x34, 0x58, 0xc0, 0xdb, 0x31, 0x68, 0x41, 0x0d, 0xa4, 0x40,
	0x25, 0xf4, 0xba, 0x09, 0x24, 0xe0, 0x0b, 0xcc, 0x65, 0x05, 0x13, 0xbd, 0x87, 0x68, 0xeb, 0x69,
	0x2d, 0xc3, 0x18, 0xd5, 0x15, 0x5f, 0x08, 0x52, 0xeb, 0x07, 0x83, 0xd6, 0xd8, 0xe7, 0xf8, 0x00,
	0xb5, 0x5c, 0x34, 0x19, 0x8f, 0x05, 0x09, 0x7d, 0x61, 0x73, 0x81, 0x6f, 0x51, 0x47, 0x0b, 0x03,
	0x4b, 0x1d, 0x8b, 0xd7, 0x5c, 0x68, 0x23, 0x41, 0x91, 0xba, 0x83, 0x46, 0xfb, 0xdf, 0x9f, 0x47,
	0x7b, 0x56, 0x18, 0x3b, 0x95, 0xb3, 0xd9, 0x30, 0x92, 0x89, 0x02, 0x2d, 0xa2, 0xf1, 0x6e, 0x89,
	0x3f, 0x17, 0x34, 0x1e, 0xa2, 0x66, 0xca, 0x27, 0x22, 0x35, 0xa4, 0xd1, 0x0f, 0x07, 0xed, 0xb3,
	0x88, 0x56, 0xbb, 0xa6, 0x65, 0x6f, 0xf4, 0xd1, 0x43, 0xf7, 0xca, 0xea, 0xd5, 0x78, 0xad, 0xc0,
	0x0f, 0xa8, 0xcd, 0x95, 0x02, 0xcb, 0xad, 0x04, 0x65, 0x48, 0xd3, 0x1b, 0x1c, 0xff, 0x61, 0x70,
	0xb7, 0x21, 0x0b, 0x97, 0xaa, 0xb6, 0x77, 0x85, 0xda, 0x95, 0x17, 0x70, 0x07, 0x85, 0x73, 0xb1,
	0x22, 0x81, 0x9f, 0xd7, 0xa5, 0xb8, 0x8b, 0x1a, 0x39, 0x4f, 0x97, 0xe5, 0x72, 0x8a, 0xc3, 0xb0,
	0x76, 0x19, 0xf4, 0x6e, 0x50, 0xe7, 0xb7, 0xf7, 0x7f, 0xf4, 0xa3, 0xeb, 0x8f, 0xaf, 0xc3, 0xe0,
	0xe5, 0x22, 0x91, 0xf6, 0x6d, 0x39, 0xa1, 0x31, 0x2c, 0x98, 0xeb, 0xfd, 0x44, 0x42, 0x11, 0xe7,
	0xd2, 0xb2, 0x6c, 0x9e, 0x30, 0x9e, 0x49, 0x96, 0x9f, 0xb2, 0x72, 0x7f, 0x86, 0xb9, 0x31, 0x27,
	0x4d, 0xff, 0x8d, 0xe7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x39, 0x9d, 0x52, 0xfc, 0x01,
	0x00, 0x00,
}
