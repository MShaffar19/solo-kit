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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

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
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A sequence number representing a specific generation of the desired state.
	// Currently only populated for resources backed by Kubernetes
	Generation           int64    `protobuf:"varint,8,opt,name=generation,proto3" json:"generation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *Metadata) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func init() {
	proto.RegisterType((*Metadata)(nil), "core.solo.io.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "core.solo.io.Metadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.solo.io.Metadata.LabelsEntry")
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor_56d9f74966f40d04) }

var fileDescriptor_56d9f74966f40d04 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0x4d, 0x7f, 0xa7, 0x97, 0x7b, 0x7b, 0x87, 0x0a, 0x43, 0x91, 0x5a, 0xb2, 0xb1, 0x1b,
	0x13, 0x54, 0x04, 0xad, 0x20, 0x5a, 0x70, 0x21, 0xe8, 0x26, 0x0b, 0x17, 0x6e, 0x64, 0x9a, 0x9e,
	0xc6, 0xa1, 0xe9, 0x9c, 0x30, 0x33, 0x29, 0xf4, 0x8d, 0x7c, 0x06, 0x1f, 0xc6, 0x85, 0x8f, 0xe0,
	0x13, 0x48, 0x26, 0x0d, 0x0d, 0x82, 0x0b, 0x57, 0xe7, 0xef, 0xfb, 0xbe, 0x39, 0xe7, 0x63, 0xc8,
	0xdf, 0x15, 0x18, 0x3e, 0xe7, 0x86, 0xfb, 0xa9, 0x42, 0x83, 0xf4, 0x4f, 0x84, 0x0a, 0x7c, 0x8d,
	0x09, 0xfa, 0x02, 0x07, 0xfd, 0x18, 0x63, 0xb4, 0x83, 0x20, 0xcf, 0x0a, 0x8c, 0xf7, 0xe6, 0x92,
	0xf6, 0xc3, 0x96, 0x46, 0x29, 0xa9, 0x4b, 0xbe, 0x02, 0x56, 0x1b, 0x39, 0xe3, 0x4e, 0x68, 0x73,
	0xba, 0x4f, 0x3a, 0x79, 0xd4, 0x29, 0x8f, 0x80, 0xb9, 0x76, 0xb0, 0x6b, 0x50, 0x46, 0x5a, 0x51,
	0x92, 0x69, 0x03, 0x8a, 0xb5, 0xec, 0xac, 0x2c, 0xe9, 0x35, 0xe9, 0x29, 0xd0, 0x98, 0xa9, 0x08,
	0x9e, 0xd7, 0xa0, 0xb4, 0x40, 0xc9, 0xea, 0x39, 0x64, 0xba, 0xf7, 0xf9, 0x7e, 0xf0, 0xdf, 0x80,
	0x36, 0x73, 0xb1, 0x58, 0x4c, 0x3c, 0x11, 0x4b, 0x54, 0xe0, 0x85, 0xff, 0x4a, 0xf8, 0x63, 0x81,
	0xa6, 0x13, 0xd2, 0x4c, 0xf8, 0x0c, 0x12, 0xcd, 0x1a, 0x23, 0x77, 0xdc, 0x3d, 0xf1, 0xfc, 0xea,
	0x3d, 0x7e, 0xb9, 0xb5, 0x7f, 0x6f, 0x41, 0xb7, 0xd2, 0xa8, 0x4d, 0xb8, 0x65, 0xd0, 0x3b, 0xd2,
	0xe5, 0x52, 0xa2, 0xe1, 0x46, 0xa0, 0xd4, 0xac, 0x69, 0x05, 0x0e, 0x7f, 0x10, 0xb8, 0xd9, 0x21,
	0x0b, 0x95, 0x2a, 0x97, 0x0e, 0x09, 0x89, 0x41, 0x82, 0xb2, 0x25, 0x6b, 0x8f, 0x9c, 0xb1, 0x1b,
	0x56, 0x3a, 0x83, 0x0b, 0xd2, 0xad, 0x6c, 0x40, 0x7b, 0xc4, 0x5d, 0xc2, 0x86, 0x39, 0xd6, 0x8d,
	0x3c, 0xa5, 0x7d, 0xd2, 0x58, 0xf3, 0x24, 0x2b, 0x6d, 0x2d, 0x8a, 0x49, 0xed, 0xdc, 0x19, 0x5c,
	0x91, 0xde, 0xf7, 0xb7, 0x7f, 0xc3, 0x9f, 0x5e, 0xbe, 0x7e, 0x0c, 0x9d, 0xa7, 0xb3, 0x58, 0x98,
	0x97, 0x6c, 0xe6, 0x47, 0xb8, 0x0a, 0xf2, 0xdb, 0x8e, 0x04, 0x16, 0x71, 0x29, 0x4c, 0x90, 0x2e,
	0xe3, 0x80, 0xa7, 0x22, 0x58, 0x1f, 0x07, 0xa5, 0xbf, 0x3a, 0xc8, 0x6d, 0x98, 0x35, 0xed, 0x07,
	0x38, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x98, 0x9c, 0x46, 0x6d, 0x36, 0x02, 0x00, 0x00,
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
	if this.Generation != that1.Generation {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
