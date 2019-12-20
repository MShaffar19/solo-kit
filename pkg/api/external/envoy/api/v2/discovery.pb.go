// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/external/envoy/api/v2/discovery.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// A DiscoveryRequest requests a set of versioned resources of the same type for
// a given Envoy node on some API.
type DiscoveryRequest struct {
	// The version_info provided in the request messages will be the version_info
	// received with the most recent successfully processed response or empty on
	// the first request. It is expected that no new request is sent after a
	// response is received until the Envoy instance is ready to ACK/NACK the new
	// configuration. ACK/NACK takes place by returning the new API config version
	// as applied or the previous API config version respectively. Each type_url
	// (see below) has an independent version associated with it.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The node making the request.
	Node *core.Node `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	// List of resources to subscribe to, e.g. list of cluster names or a route
	// configuration name. If this is empty, all resources for the API are
	// returned. LDS/CDS may have empty resource_names, which will cause all
	// resources for the Envoy instance to be returned. The LDS and CDS responses
	// will then imply a number of resources that need to be fetched via EDS/RDS,
	// which will be explicitly enumerated in resource_names.
	ResourceNames []string `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment". This is implicit
	// in requests made via singleton xDS APIs such as CDS, LDS, etc. but is
	// required for ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// nonce corresponding to DiscoveryResponse being ACK/NACKed. See above
	// discussion on version_info and the DiscoveryResponse nonce comment. This
	// may be empty only if 1) this is a non-persistent-stream xDS such as HTTP,
	// or 2) the client has not yet accepted an update in this xDS stream (unlike
	// delta, where it is populated only for new explicit ACKs).
	ResponseNonce string `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details* provides the Envoy
	// internal exception related to the failure. It is only intended for consumption during manual
	// debugging, the string provided is not guaranteed to be stable across Envoy versions.
	ErrorDetail          *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58de2eafe6af4e24, []int{0}
}
func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	// The version of the response data.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The response resources. These resources are typed and depend on the API being called.
	Resources []*types.Any `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// [#not-implemented-hide:]
	// Canary is used to support two Envoy command line flags:
	//
	// * --terminate-on-canary-transition-failure. When set, Envoy is able to
	//   terminate if it detects that configuration is stuck at canary. Consider
	//   this example sequence of updates:
	//   - Management server applies a canary config successfully.
	//   - Management server rolls back to a production config.
	//   - Envoy rejects the new production config.
	//   Since there is no sensible way to continue receiving configuration
	//   updates, Envoy will then terminate and apply production config from a
	//   clean slate.
	// * --dry-run-canary. When set, a canary response will never be applied, only
	//   validated via a dry run.
	Canary bool `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	// Type URL for resources. Identifies the xDS API when muxing over ADS.
	// Must be consistent with the type_url in the 'resources' repeated Any (if non-empty).
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// For gRPC based subscriptions, the nonce provides a way to explicitly ack a
	// specific DiscoveryResponse in a following DiscoveryRequest. Additional
	// messages may have been sent by Envoy to the management server for the
	// previous version on the stream prior to this DiscoveryResponse, that were
	// unprocessed at response send time. The nonce allows the management server
	// to ignore any further DiscoveryRequests for the previous version until a
	// DiscoveryRequest bearing the nonce. The nonce is optional and is not
	// required for non-stream based xDS implementations.
	Nonce string `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// [#not-implemented-hide:]
	// The control plane instance that sent the response.
	ControlPlane         *core.ControlPlane `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58de2eafe6af4e24, []int{1}
}
func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []*types.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *DiscoveryResponse) GetControlPlane() *core.ControlPlane {
	if m != nil {
		return m.ControlPlane
	}
	return nil
}

// DeltaDiscoveryRequest and DeltaDiscoveryResponse are used in a new gRPC
// endpoint for Delta xDS.
//
// With Delta xDS, the DeltaDiscoveryResponses do not need to include a full
// snapshot of the tracked resources. Instead, DeltaDiscoveryResponses are a
// diff to the state of a xDS client.
// In Delta XDS there are per-resource versions, which allow tracking state at
// the resource granularity.
// An xDS Delta session is always in the context of a gRPC bidirectional
// stream. This allows the xDS server to keep track of the state of xDS clients
// connected to it.
//
// In Delta xDS the nonce field is required and used to pair
// DeltaDiscoveryResponse to a DeltaDiscoveryRequest ACK or NACK.
// Optionally, a response message level system_version_info is present for
// debugging purposes only.
//
// DeltaDiscoveryRequest plays two independent roles. Any DeltaDiscoveryRequest
// can be either or both of: [1] informing the server of what resources the
// client has gained/lost interest in (using resource_names_subscribe and
// resource_names_unsubscribe), or [2] (N)ACKing an earlier resource update from
// the server (using response_nonce, with presence of error_detail making it a NACK).
// Additionally, the first message (for a given type_url) of a reconnected gRPC stream
// has a third role: informing the server of the resources (and their versions)
// that the client already possesses, using the initial_resource_versions field.
//
// As with state-of-the-world, when multiple resource types are multiplexed (ADS),
// all requests/acknowledgments/updates are logically walled off by type_url:
// a Cluster ACK exists in a completely separate world from a prior Route NACK.
// In particular, initial_resource_versions being sent at the "start" of every
// gRPC stream actually entails a message for each type_url, each with its own
// initial_resource_versions.
type DeltaDiscoveryRequest struct {
	// The node making the request.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment".
	TypeUrl string `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// DeltaDiscoveryRequests allow the client to add or remove individual
	// resources to the set of tracked resources in the context of a stream.
	// All resource names in the resource_names_subscribe list are added to the
	// set of tracked resources and all resource names in the resource_names_unsubscribe
	// list are removed from the set of tracked resources.
	//
	// *Unlike* state-of-the-world xDS, an empty resource_names_subscribe or
	// resource_names_unsubscribe list simply means that no resources are to be
	// added or removed to the resource list.
	// *Like* state-of-the-world xDS, the server must send updates for all tracked
	// resources, but can also send updates for resources the client has not subscribed to.
	//
	// NOTE: the server must respond with all resources listed in resource_names_subscribe,
	// even if it believes the client has the most recent version of them. The reason:
	// the client may have dropped them, but then regained interest before it had a chance
	// to send the unsubscribe message. See DeltaSubscriptionStateTest.RemoveThenAdd.
	//
	// These two fields can be set in any DeltaDiscoveryRequest, including ACKs
	// and initial_resource_versions.
	//
	// A list of Resource names to add to the list of tracked resources.
	ResourceNamesSubscribe []string `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	// A list of Resource names to remove from the list of tracked resources.
	ResourceNamesUnsubscribe []string `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	// Informs the server of the versions of the resources the xDS client knows of, to enable the
	// client to continue the same logical xDS session even in the face of gRPC stream reconnection.
	// It will not be populated: [1] in the very first stream of a session, since the client will
	// not yet have any resources,  [2] in any message after the first in a stream (for a given
	// type_url), since the server will already be correctly tracking the client's state.
	// (In ADS, the first message *of each type_url* of a reconnected stream populates this map.)
	// The map's keys are names of xDS resources known to the xDS client.
	// The map's values are opaque resource versions.
	InitialResourceVersions map[string]string `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// When the DeltaDiscoveryRequest is a ACK or NACK message in response
	// to a previous DeltaDiscoveryResponse, the response_nonce must be the
	// nonce in the DeltaDiscoveryResponse.
	// Otherwise (unlike in DiscoveryRequest) response_nonce must be omitted.
	ResponseNonce string `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details*
	// provides the Envoy internal exception related to the failure.
	ErrorDetail          *status.Status `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeltaDiscoveryRequest) Reset()         { *m = DeltaDiscoveryRequest{} }
func (m *DeltaDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryRequest) ProtoMessage()    {}
func (*DeltaDiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58de2eafe6af4e24, []int{2}
}
func (m *DeltaDiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryRequest.Unmarshal(m, b)
}
func (m *DeltaDiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryRequest.Merge(m, src)
}
func (m *DeltaDiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryRequest.Size(m)
}
func (m *DeltaDiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryRequest proto.InternalMessageInfo

func (m *DeltaDiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if m != nil {
		return m.ResourceNamesSubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if m != nil {
		return m.ResourceNamesUnsubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if m != nil {
		return m.InitialResourceVersions
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DeltaDiscoveryResponse struct {
	// The version of the response data (used for debugging).
	SystemVersionInfo string `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	// The response resources. These are typed resources, whose types must match
	// the type_url field.
	Resources []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// Type URL for resources. Identifies the xDS API when muxing over ADS.
	// Must be consistent with the type_url in the Any within 'resources' if 'resources' is non-empty.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Resources names of resources that have be deleted and to be removed from the xDS Client.
	// Removed resources for missing resources can be ignored.
	RemovedResources []string `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	// The nonce provides a way for DeltaDiscoveryRequests to uniquely
	// reference a DeltaDiscoveryResponse when (N)ACKing. The nonce is required.
	Nonce                string   `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeltaDiscoveryResponse) Reset()         { *m = DeltaDiscoveryResponse{} }
func (m *DeltaDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryResponse) ProtoMessage()    {}
func (*DeltaDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58de2eafe6af4e24, []int{3}
}
func (m *DeltaDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryResponse.Unmarshal(m, b)
}
func (m *DeltaDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryResponse.Merge(m, src)
}
func (m *DeltaDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryResponse.Size(m)
}
func (m *DeltaDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryResponse proto.InternalMessageInfo

func (m *DeltaDiscoveryResponse) GetSystemVersionInfo() string {
	if m != nil {
		return m.SystemVersionInfo
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetRemovedResources() []string {
	if m != nil {
		return m.RemovedResources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type Resource struct {
	// The resource's name, to distinguish it from others of the same type of resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// [#not-implemented-hide:]
	// The aliases are a list of other names that this resource can go by.
	Aliases []string `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// The resource level version. It allows xDS to track the state of individual
	// resources.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The resource being tracked.
	Resource             *types.Any `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_58de2eafe6af4e24, []int{4}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetResource() *types.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v2.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v2.DiscoveryResponse")
	proto.RegisterType((*DeltaDiscoveryRequest)(nil), "envoy.api.v2.DeltaDiscoveryRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.api.v2.DeltaDiscoveryRequest.InitialResourceVersionsEntry")
	proto.RegisterType((*DeltaDiscoveryResponse)(nil), "envoy.api.v2.DeltaDiscoveryResponse")
	proto.RegisterType((*Resource)(nil), "envoy.api.v2.Resource")
}

func init() {
	proto.RegisterFile("api/external/envoy/api/v2/discovery.proto", fileDescriptor_58de2eafe6af4e24)
}

var fileDescriptor_58de2eafe6af4e24 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xc1, 0x6e, 0xdb, 0x38,
	0x10, 0x85, 0x6c, 0xc7, 0xb1, 0x69, 0x27, 0x48, 0xb8, 0x59, 0x47, 0x31, 0x82, 0xac, 0xd7, 0xc0,
	0x02, 0x5e, 0x04, 0x2b, 0x2d, 0xbc, 0xbb, 0x40, 0xb6, 0xe8, 0xa1, 0x4d, 0xdd, 0x43, 0x7a, 0x08,
	0x02, 0x05, 0xc9, 0xa1, 0x17, 0x81, 0x96, 0x27, 0x2e, 0x11, 0x99, 0x54, 0x49, 0xca, 0xb0, 0x80,
	0x9e, 0x8a, 0x7e, 0x4c, 0x3f, 0xa1, 0x9f, 0xd2, 0x43, 0x4f, 0xf9, 0x87, 0x1e, 0x7a, 0x6a, 0x21,
	0x8a, 0xb2, 0xad, 0xc4, 0x0d, 0x7c, 0x32, 0x67, 0xe6, 0x71, 0x38, 0x33, 0xef, 0x8d, 0x85, 0xfe,
	0x24, 0x11, 0x75, 0x61, 0xa6, 0x40, 0x30, 0x12, 0xba, 0xc0, 0xa6, 0x3c, 0x71, 0x53, 0xd7, 0xb4,
	0xef, 0x8e, 0xa8, 0x0c, 0xf8, 0x14, 0x44, 0xe2, 0x44, 0x82, 0x2b, 0x8e, 0x9b, 0x3a, 0xea, 0x90,
	0x88, 0x3a, 0xd3, 0x7e, 0xfb, 0xb0, 0x80, 0x0d, 0xb8, 0x00, 0x77, 0x48, 0x24, 0x64, 0xd8, 0xf6,
	0xc1, 0x98, 0xf3, 0x71, 0x08, 0xae, 0xb6, 0x86, 0xf1, 0x8d, 0x4b, 0x98, 0x49, 0xd3, 0xde, 0x37,
	0x21, 0x11, 0x05, 0xae, 0x54, 0x44, 0xc5, 0xd2, 0x04, 0xf6, 0xc6, 0x7c, 0xcc, 0xf5, 0xd1, 0x4d,
	0x4f, 0xc6, 0x8b, 0x61, 0xa6, 0x32, 0x27, 0xcc, 0x54, 0xe6, 0xeb, 0xbe, 0x2f, 0xa1, 0x9d, 0x41,
	0x5e, 0x9d, 0x07, 0x6f, 0x63, 0x90, 0x0a, 0xff, 0x8e, 0x9a, 0x53, 0x10, 0x92, 0x72, 0xe6, 0x53,
	0x76, 0xc3, 0x6d, 0xab, 0x63, 0xf5, 0xea, 0x5e, 0xc3, 0xf8, 0xce, 0xd8, 0x0d, 0xc7, 0xc7, 0xa8,
	0xc2, 0xf8, 0x08, 0xec, 0x52, 0xc7, 0xea, 0x35, 0xfa, 0xfb, 0xce, 0x72, 0x43, 0x4e, 0xda, 0x82,
	0x73, 0xce, 0x47, 0xe0, 0x69, 0x10, 0xfe, 0x03, 0x6d, 0x0b, 0x90, 0x3c, 0x16, 0x01, 0xf8, 0x8c,
	0x4c, 0x40, 0xda, 0xe5, 0x4e, 0xb9, 0x57, 0xf7, 0xb6, 0x72, 0xef, 0x79, 0xea, 0xc4, 0x07, 0xa8,
	0xa6, 0x92, 0x08, 0xfc, 0x58, 0x84, 0x76, 0x45, 0x3f, 0xb9, 0x99, 0xda, 0x57, 0x22, 0x34, 0x19,
	0x22, 0xce, 0x24, 0xf8, 0x8c, 0xb3, 0x00, 0xec, 0x0d, 0x0d, 0xd8, 0xca, 0xbd, 0xe7, 0xa9, 0x13,
	0xff, 0x87, 0x9a, 0x20, 0x04, 0x17, 0xfe, 0x08, 0x14, 0xa1, 0xa1, 0x5d, 0xd5, 0xd5, 0x61, 0x27,
	0x9b, 0x93, 0x23, 0xa2, 0xc0, 0xb9, 0xd4, 0x73, 0xf2, 0x1a, 0x1a, 0x37, 0xd0, 0xb0, 0xee, 0x37,
	0x0b, 0xed, 0x2e, 0x0d, 0x21, 0xcb, 0xb8, 0xce, 0x14, 0xfa, 0xa8, 0x9e, 0xb7, 0x20, 0xed, 0x52,
	0xa7, 0xdc, 0x6b, 0xf4, 0xf7, 0xf2, 0xc7, 0x72, 0xbe, 0x9c, 0xe7, 0x2c, 0xf1, 0x16, 0x30, 0xdc,
	0x42, 0xd5, 0x80, 0x30, 0x22, 0x12, 0xbb, 0xdc, 0xb1, 0x7a, 0x35, 0xcf, 0x58, 0x8f, 0x75, 0xbf,
	0x87, 0x36, 0x96, 0x9b, 0xce, 0x0c, 0x3c, 0x40, 0x5b, 0x01, 0x67, 0x4a, 0xf0, 0xd0, 0x8f, 0x42,
	0xc2, 0xc0, 0x74, 0xfb, 0xdb, 0x0a, 0x2e, 0x5e, 0x64, 0xb8, 0x8b, 0x14, 0xe6, 0x35, 0x83, 0x25,
	0xab, 0xfb, 0xbd, 0x8c, 0x7e, 0x1d, 0x40, 0xa8, 0xc8, 0x03, 0x15, 0xe4, 0x14, 0x5b, 0xeb, 0x50,
	0xbc, 0x5c, 0x7d, 0xa9, 0x58, 0xfd, 0x09, 0xb2, 0x8b, 0xec, 0xfb, 0x32, 0x1e, 0xca, 0x40, 0xd0,
	0x21, 0x18, 0x1d, 0xb4, 0x0a, 0x3a, 0xb8, 0xcc, 0xa3, 0xf8, 0x29, 0x6a, 0xdf, 0xbb, 0x19, 0xb3,
	0xc5, 0xdd, 0x8a, 0xbe, 0x6b, 0x17, 0xee, 0x5e, 0x2d, 0xe2, 0xf8, 0x1d, 0x3a, 0xa0, 0x8c, 0x2a,
	0x4a, 0x42, 0x7f, 0x9e, 0xc5, 0x90, 0x27, 0xed, 0x0d, 0x4d, 0xd6, 0xb3, 0x62, 0x53, 0x2b, 0xe7,
	0xe0, 0x9c, 0x65, 0x49, 0x3c, 0x93, 0xe3, 0xda, 0xa4, 0x78, 0xc9, 0x94, 0x48, 0xbc, 0x7d, 0xba,
	0x3a, 0xba, 0x42, 0xb1, 0xd5, 0x75, 0x14, 0xbb, 0xb9, 0x96, 0x62, 0xdb, 0xaf, 0xd0, 0xe1, 0x63,
	0x65, 0xe1, 0x1d, 0x54, 0xbe, 0x85, 0xc4, 0x48, 0x36, 0x3d, 0xa6, 0x1a, 0x9a, 0x92, 0x30, 0x06,
	0xc3, 0x4e, 0x66, 0x3c, 0x29, 0x9d, 0x58, 0xdd, 0x2f, 0x16, 0x6a, 0xdd, 0xef, 0xdc, 0xac, 0x80,
	0x83, 0x7e, 0x91, 0x89, 0x54, 0x30, 0xf1, 0x57, 0x6c, 0xc2, 0x6e, 0x16, 0xba, 0x5e, 0xda, 0x87,
	0x7f, 0x1f, 0xee, 0x43, 0xab, 0x38, 0xe2, 0xbc, 0xdc, 0xe5, 0x8d, 0x78, 0x44, 0xf9, 0xc7, 0x68,
	0x57, 0xc0, 0x84, 0x4f, 0x61, 0xe4, 0x2f, 0x12, 0x57, 0x35, 0xf1, 0x3b, 0x26, 0xe0, 0xcd, 0xf3,
	0xac, 0x5c, 0x93, 0xee, 0x07, 0x0b, 0xd5, 0x72, 0x0c, 0xc6, 0xa8, 0x92, 0x0a, 0x49, 0xaf, 0x5e,
	0xdd, 0xd3, 0x67, 0x6c, 0xa3, 0x4d, 0x12, 0x52, 0x22, 0x41, 0x1a, 0x49, 0xe5, 0x66, 0x1a, 0x31,
	0x7d, 0x9b, 0x96, 0x73, 0x13, 0xff, 0x8d, 0x6a, 0x79, 0x3d, 0xe6, 0x2f, 0x70, 0xf5, 0xde, 0xcf,
	0x51, 0xa7, 0xf1, 0xa7, 0xaf, 0x15, 0xeb, 0xe3, 0xdd, 0x91, 0xf5, 0xf9, 0xee, 0xc8, 0x42, 0x6d,
	0xca, 0xb3, 0xb9, 0x44, 0x82, 0xcf, 0x92, 0xc2, 0x88, 0x4e, 0xb7, 0xe7, 0x3c, 0x5c, 0xa4, 0xa9,
	0x2e, 0xac, 0xd7, 0xff, 0x8f, 0xa9, 0x7a, 0x13, 0x0f, 0x9d, 0x80, 0x4f, 0x5c, 0xc9, 0x43, 0xfe,
	0x17, 0xe5, 0xd9, 0xef, 0x2d, 0x55, 0x6e, 0x74, 0x3b, 0x76, 0x7f, 0xfa, 0xe5, 0x19, 0x56, 0x75,
	0x39, 0xff, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x35, 0xaf, 0x15, 0x9d, 0x06, 0x00, 0x00,
}

func (this *DiscoveryRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryRequest)
	if !ok {
		that2, ok := that.(DiscoveryRequest)
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
	if this.VersionInfo != that1.VersionInfo {
		return false
	}
	if !this.Node.Equal(that1.Node) {
		return false
	}
	if len(this.ResourceNames) != len(that1.ResourceNames) {
		return false
	}
	for i := range this.ResourceNames {
		if this.ResourceNames[i] != that1.ResourceNames[i] {
			return false
		}
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.ResponseNonce != that1.ResponseNonce {
		return false
	}
	if !this.ErrorDetail.Equal(that1.ErrorDetail) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DiscoveryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryResponse)
	if !ok {
		that2, ok := that.(DiscoveryResponse)
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
	if this.VersionInfo != that1.VersionInfo {
		return false
	}
	if len(this.Resources) != len(that1.Resources) {
		return false
	}
	for i := range this.Resources {
		if !this.Resources[i].Equal(that1.Resources[i]) {
			return false
		}
	}
	if this.Canary != that1.Canary {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if !this.ControlPlane.Equal(that1.ControlPlane) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeltaDiscoveryRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeltaDiscoveryRequest)
	if !ok {
		that2, ok := that.(DeltaDiscoveryRequest)
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
	if !this.Node.Equal(that1.Node) {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if len(this.ResourceNamesSubscribe) != len(that1.ResourceNamesSubscribe) {
		return false
	}
	for i := range this.ResourceNamesSubscribe {
		if this.ResourceNamesSubscribe[i] != that1.ResourceNamesSubscribe[i] {
			return false
		}
	}
	if len(this.ResourceNamesUnsubscribe) != len(that1.ResourceNamesUnsubscribe) {
		return false
	}
	for i := range this.ResourceNamesUnsubscribe {
		if this.ResourceNamesUnsubscribe[i] != that1.ResourceNamesUnsubscribe[i] {
			return false
		}
	}
	if len(this.InitialResourceVersions) != len(that1.InitialResourceVersions) {
		return false
	}
	for i := range this.InitialResourceVersions {
		if this.InitialResourceVersions[i] != that1.InitialResourceVersions[i] {
			return false
		}
	}
	if this.ResponseNonce != that1.ResponseNonce {
		return false
	}
	if !this.ErrorDetail.Equal(that1.ErrorDetail) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeltaDiscoveryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeltaDiscoveryResponse)
	if !ok {
		that2, ok := that.(DeltaDiscoveryResponse)
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
	if this.SystemVersionInfo != that1.SystemVersionInfo {
		return false
	}
	if len(this.Resources) != len(that1.Resources) {
		return false
	}
	for i := range this.Resources {
		if !this.Resources[i].Equal(that1.Resources[i]) {
			return false
		}
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if len(this.RemovedResources) != len(that1.RemovedResources) {
		return false
	}
	for i := range this.RemovedResources {
		if this.RemovedResources[i] != that1.RemovedResources[i] {
			return false
		}
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Resource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Resource)
	if !ok {
		that2, ok := that.(Resource)
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
	if len(this.Aliases) != len(that1.Aliases) {
		return false
	}
	for i := range this.Aliases {
		if this.Aliases[i] != that1.Aliases[i] {
			return false
		}
	}
	if this.Version != that1.Version {
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
