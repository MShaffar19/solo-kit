// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metadata.proto

package core

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//*
// Metadata contains general properties of resources for purposes of versioning, annotating, and namespacing.
type Metadata struct {
	//
	//Name of the resource.
	//
	//Names must be unique and follow the following syntax rules:
	//
	//One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Namespace is used for the namespacing of resources.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Cluster indicates the cluster this resource belongs to
	// Cluster is only applicable in certain contexts, e.g. Kubernetes
	// An empty string here refers to the local cluster
	Cluster string `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// An opaque value that represents the internal version of this object that can
	// be used by clients to determine when objects have changed.
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty" testdiff:"ignore"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. Some resources contain `selectors` which
	// can be linked with other resources by their labels
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata.
	Annotations          map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
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

func (m *Metadata) GetCluster() string {
	if m != nil {
		return m.Cluster
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

func init() { proto.RegisterFile("metadata.proto", fileDescriptor_56d9f74966f40d04) }

var fileDescriptor_56d9f74966f40d04 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xd9, 0x6e, 0xff, 0xd8, 0x54, 0xb4, 0x86, 0x0a, 0xa1, 0x88, 0x96, 0xbd, 0xd8, 0x8b,
	0x09, 0x2a, 0x82, 0x56, 0x10, 0x2d, 0x78, 0x10, 0xf4, 0xd2, 0x83, 0x07, 0x2f, 0x92, 0x6e, 0xd3,
	0x35, 0x74, 0x9b, 0x59, 0x92, 0x6c, 0xa1, 0xaf, 0xe2, 0x13, 0xf8, 0x54, 0x1e, 0x7c, 0x04, 0x9f,
	0x40, 0x92, 0xed, 0xd2, 0x22, 0x78, 0xf0, 0x34, 0x33, 0x3b, 0xbf, 0xef, 0xdb, 0x99, 0x09, 0xda,
	0x99, 0x0b, 0xcb, 0x27, 0xdc, 0x72, 0x9a, 0x69, 0xb0, 0x80, 0xb7, 0x63, 0xd0, 0x82, 0x1a, 0x48,
	0x81, 0x4a, 0xe8, 0x76, 0x12, 0x48, 0xc0, 0x37, 0x98, 0xcb, 0x0a, 0x26, 0x7a, 0x0f, 0xd1, 0xd6,
	0xd3, 0x4a, 0x86, 0x31, 0xaa, 0x2a, 0x3e, 0x17, 0xa4, 0xd2, 0x0b, 0xfa, 0xcd, 0x91, 0xcf, 0xf1,
	0x01, 0x6a, 0xba, 0x68, 0x32, 0x1e, 0x0b, 0x12, 0xfa, 0xc6, 0xfa, 0x03, 0x26, 0xa8, 0x11, 0xa7,
	0xb9, 0xb1, 0x42, 0x93, 0x86, 0xef, 0x95, 0x25, 0xbe, 0x45, 0x6d, 0x2d, 0x0c, 0xe4, 0x3a, 0x16,
	0xaf, 0x0b, 0xa1, 0x8d, 0x04, 0x45, 0xaa, 0x0e, 0x19, 0xee, 0x7f, 0x7f, 0x1e, 0xed, 0x59, 0x61,
	0xec, 0x44, 0x4e, 0xa7, 0x83, 0x48, 0x26, 0x0a, 0xb4, 0x88, 0x46, 0xbb, 0x25, 0xfe, 0x5c, 0xd0,
	0x78, 0x80, 0xea, 0x29, 0x1f, 0x8b, 0xd4, 0x90, 0x5a, 0x2f, 0xec, 0xb7, 0xce, 0x22, 0xba, 0xb9,
	0x0f, 0x2d, 0xa7, 0xa6, 0x8f, 0x1e, 0xba, 0x57, 0x56, 0x2f, 0x47, 0x2b, 0x05, 0x7e, 0x40, 0x2d,
	0xae, 0x14, 0x58, 0x6e, 0x25, 0x28, 0x43, 0xea, 0xde, 0xe0, 0xf8, 0x0f, 0x83, 0xbb, 0x35, 0x59,
	0xb8, 0x6c, 0x6a, 0xbb, 0x57, 0xa8, 0xb5, 0xf1, 0x07, 0xdc, 0x46, 0xe1, 0x4c, 0x2c, 0x49, 0xe0,
	0xb7, 0x75, 0x29, 0xee, 0xa0, 0xda, 0x82, 0xa7, 0x79, 0x79, 0xb6, 0xa2, 0x18, 0x54, 0x2e, 0x83,
	0xee, 0x0d, 0x6a, 0xff, 0xf6, 0xfe, 0x8f, 0x7e, 0x78, 0xfd, 0xf1, 0x75, 0x18, 0xbc, 0x5c, 0x24,
	0xd2, 0xbe, 0xe5, 0x63, 0x1a, 0xc3, 0x9c, 0xb9, 0xd9, 0x4f, 0x24, 0x14, 0x71, 0x26, 0x2d, 0xcb,
	0x66, 0x09, 0xe3, 0x99, 0x64, 0x8b, 0x53, 0x56, 0xde, 0xcf, 0x30, 0xb7, 0xe6, 0xb8, 0xee, 0x1f,
	0xf8, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x1f, 0x9c, 0x5c, 0x16, 0x02, 0x00, 0x00,
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
	if this.Cluster != that1.Cluster {
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
