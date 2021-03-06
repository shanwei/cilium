// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/attribute_context.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An attribute is a piece of metadata that describes an activity on a network.
// For example, the size of an HTTP request, or the status code of an HTTP response.
//
// Each attribute has a type and a name, which is logically defined as a proto message field
// of the `AttributeContext`. The `AttributeContext` is a collection of individual attributes
// supported by Envoy authorization system.
type AttributeContext struct {
	// The source of a network activity, such as starting a TCP connection.
	// In a multi hop network activity, the source represents the sender of the
	// last hop.
	Source *AttributeContext_Peer `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The destination of a network activity, such as accepting a TCP connection.
	// In a multi hop network activity, the destination represents the receiver of
	// the last hop.
	Destination *AttributeContext_Peer `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Represents a network request, such as an HTTP request.
	Request *AttributeContext_Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	// This is analogous to http_request.headers, however these contents will not be sent to the
	// upstream server. Context_extensions provide an extension mechanism for sending additional
	// information to the auth server without modifying the proto definition. It maps to the
	// internal opaque context in the filter chain.
	ContextExtensions    map[string]string `protobuf:"bytes,10,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttributeContext) Reset()         { *m = AttributeContext{} }
func (m *AttributeContext) String() string { return proto.CompactTextString(m) }
func (*AttributeContext) ProtoMessage()    {}
func (*AttributeContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_attribute_context_8dd39e33b53d559c, []int{0}
}
func (m *AttributeContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext.Unmarshal(m, b)
}
func (m *AttributeContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext.Marshal(b, m, deterministic)
}
func (dst *AttributeContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext.Merge(dst, src)
}
func (m *AttributeContext) XXX_Size() int {
	return xxx_messageInfo_AttributeContext.Size(m)
}
func (m *AttributeContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext proto.InternalMessageInfo

func (m *AttributeContext) GetSource() *AttributeContext_Peer {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttributeContext) GetDestination() *AttributeContext_Peer {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *AttributeContext) GetRequest() *AttributeContext_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AttributeContext) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

// This message defines attributes for a node that handles a network request.
// The node can be either a service or an application that sends, forwards,
// or receives the request. Service peers should fill in the `service`,
// `principal`, and `labels` as appropriate.
type AttributeContext_Peer struct {
	// The address of the peer, this is typically the IP address.
	// It can also be UDS path, or others.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The canonical service name of the peer.
	// It should be set to :ref:`the HTTP x-envoy-downstream-service-cluster
	// <config_http_conn_man_headers_downstream-service-cluster>`
	// If a more trusted source of the service name is available through mTLS/secure naming, it
	// should be used.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// The labels associated with the peer.
	// These could be pod labels for Kubernetes or tags for VMs.
	// The source of the labels could be an X.509 certificate or other configuration.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The authenticated identity of this peer.
	// For example, the identity associated with the workload such as a service account.
	// If an X.509 certificate is used to assert the identity this field should be sourced from
	// `Subject` or `Subject Alternative Names`. The primary identity should be the principal.
	// The principal format is issuer specific.
	//
	// Example:
	// *    SPIFFE format is `spiffe://trust-domain/path`
	// *    Google account format is `https://accounts.google.com/{userid}`
	Principal            string   `protobuf:"bytes,4,opt,name=principal,proto3" json:"principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttributeContext_Peer) Reset()         { *m = AttributeContext_Peer{} }
func (m *AttributeContext_Peer) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Peer) ProtoMessage()    {}
func (*AttributeContext_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_attribute_context_8dd39e33b53d559c, []int{0, 0}
}
func (m *AttributeContext_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Peer.Unmarshal(m, b)
}
func (m *AttributeContext_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Peer.Marshal(b, m, deterministic)
}
func (dst *AttributeContext_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Peer.Merge(dst, src)
}
func (m *AttributeContext_Peer) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Peer.Size(m)
}
func (m *AttributeContext_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Peer proto.InternalMessageInfo

func (m *AttributeContext_Peer) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AttributeContext_Peer) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AttributeContext_Peer) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *AttributeContext_Peer) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

// Represents a network request, such as an HTTP request.
type AttributeContext_Request struct {
	// The timestamp when the proxy receives the first byte of the request.
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Represents an HTTP request or an HTTP-like request.
	Http                 *AttributeContext_HttpRequest `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AttributeContext_Request) Reset()         { *m = AttributeContext_Request{} }
func (m *AttributeContext_Request) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Request) ProtoMessage()    {}
func (*AttributeContext_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_attribute_context_8dd39e33b53d559c, []int{0, 1}
}
func (m *AttributeContext_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Request.Unmarshal(m, b)
}
func (m *AttributeContext_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Request.Marshal(b, m, deterministic)
}
func (dst *AttributeContext_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Request.Merge(dst, src)
}
func (m *AttributeContext_Request) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Request.Size(m)
}
func (m *AttributeContext_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Request.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Request proto.InternalMessageInfo

func (m *AttributeContext_Request) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *AttributeContext_Request) GetHttp() *AttributeContext_HttpRequest {
	if m != nil {
		return m.Http
	}
	return nil
}

// This message defines attributes for an HTTP request.
// HTTP/1.x, HTTP/2, gRPC are all considered as HTTP requests.
type AttributeContext_HttpRequest struct {
	// The unique ID for a request, which can be propagated to downstream
	// systems. The ID should have low probability of collision
	// within a single day for a specific service.
	// For HTTP requests, it should be X-Request-ID or equivalent.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The HTTP request method, such as `GET`, `POST`.
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// The HTTP request headers. If multiple headers share the same key, they
	// must be merged according to the HTTP spec. All header keys must be
	// lowercased, because HTTP header keys are case-insensitive.
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The HTTP URL path.
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// The HTTP request `Host` or 'Authority` header value.
	Host string `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	// The HTTP URL scheme, such as `http` and `https`.
	Scheme string `protobuf:"bytes,6,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// The HTTP URL query in the format of `name1=value`&name2=value2`, as it
	// appears in the first line of the HTTP request. No decoding is performed.
	Query string `protobuf:"bytes,7,opt,name=query,proto3" json:"query,omitempty"`
	// The HTTP URL fragment, excluding leading `#`. No URL decoding is performed.
	Fragment string `protobuf:"bytes,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	// The HTTP request size in bytes. If unknown, it must be -1.
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// The network protocol used with the request, such as
	// "http/1.1", "spdy/3", "h2", "h2c"
	Protocol             string   `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttributeContext_HttpRequest) Reset()         { *m = AttributeContext_HttpRequest{} }
func (m *AttributeContext_HttpRequest) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_HttpRequest) ProtoMessage()    {}
func (*AttributeContext_HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_attribute_context_8dd39e33b53d559c, []int{0, 2}
}
func (m *AttributeContext_HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_HttpRequest.Unmarshal(m, b)
}
func (m *AttributeContext_HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_HttpRequest.Marshal(b, m, deterministic)
}
func (dst *AttributeContext_HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_HttpRequest.Merge(dst, src)
}
func (m *AttributeContext_HttpRequest) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_HttpRequest.Size(m)
}
func (m *AttributeContext_HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_HttpRequest proto.InternalMessageInfo

func (m *AttributeContext_HttpRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AttributeContext_HttpRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AttributeContext_HttpRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*AttributeContext)(nil), "envoy.service.auth.v2alpha.AttributeContext")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v2alpha.AttributeContext.ContextExtensionsEntry")
	proto.RegisterType((*AttributeContext_Peer)(nil), "envoy.service.auth.v2alpha.AttributeContext.Peer")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v2alpha.AttributeContext.Peer.LabelsEntry")
	proto.RegisterType((*AttributeContext_Request)(nil), "envoy.service.auth.v2alpha.AttributeContext.Request")
	proto.RegisterType((*AttributeContext_HttpRequest)(nil), "envoy.service.auth.v2alpha.AttributeContext.HttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v2alpha.AttributeContext.HttpRequest.HeadersEntry")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/attribute_context.proto", fileDescriptor_attribute_context_8dd39e33b53d559c)
}

var fileDescriptor_attribute_context_8dd39e33b53d559c = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x51, 0x6b, 0xdb, 0x3c,
	0x14, 0x25, 0x4e, 0x1a, 0xd7, 0x37, 0x1f, 0x1f, 0x9d, 0x18, 0x45, 0x98, 0x41, 0xcb, 0x9e, 0xfa,
	0x24, 0x33, 0xaf, 0x0f, 0x5d, 0x61, 0x0f, 0xa5, 0x2b, 0x74, 0x50, 0xc6, 0xf0, 0xb6, 0xe7, 0xa2,
	0xd8, 0xb7, 0xb1, 0x98, 0x63, 0xb9, 0x92, 0x6c, 0x9a, 0xfd, 0x80, 0xed, 0x9f, 0xed, 0x17, 0xed,
	0x07, 0x0c, 0x4b, 0x72, 0x16, 0xca, 0xf6, 0xe0, 0x3e, 0x45, 0xe7, 0xe6, 0x9e, 0xa3, 0x73, 0x8f,
	0xae, 0x21, 0xc5, 0xba, 0x93, 0x9b, 0x44, 0xa3, 0xea, 0x44, 0x8e, 0x09, 0x6f, 0x4d, 0x99, 0x74,
	0x29, 0xaf, 0x9a, 0x92, 0x27, 0xdc, 0x18, 0x25, 0x96, 0xad, 0xc1, 0xdb, 0x5c, 0xd6, 0x06, 0x1f,
	0x0c, 0x6b, 0x94, 0x34, 0x92, 0xc4, 0x96, 0xc3, 0x3c, 0x87, 0xf5, 0x1c, 0xe6, 0x39, 0xf1, 0x91,
	0xd3, 0xe3, 0x8d, 0x48, 0xba, 0x34, 0xc9, 0xa5, 0xc2, 0x84, 0x17, 0x85, 0x42, 0xad, 0x1d, 0x39,
	0x3e, 0x5a, 0x49, 0xb9, 0xaa, 0x30, 0xb1, 0x68, 0xd9, 0xde, 0x25, 0x46, 0xac, 0x51, 0x1b, 0xbe,
	0x6e, 0x5c, 0xc3, 0xcb, 0x9f, 0x11, 0x1c, 0x5c, 0x0c, 0x37, 0x5f, 0xba, 0x8b, 0xc9, 0x7b, 0x98,
	0x6b, 0xd9, 0xaa, 0x1c, 0xe9, 0xe4, 0x78, 0x72, 0xb2, 0x48, 0x5f, 0xb1, 0x7f, 0x7b, 0x60, 0x8f,
	0xd9, 0xec, 0x23, 0xa2, 0xca, 0xbc, 0x00, 0xf9, 0x04, 0x8b, 0x02, 0xb5, 0x11, 0x35, 0x37, 0x42,
	0xd6, 0x34, 0x78, 0xaa, 0xde, 0xae, 0x0a, 0xf9, 0x00, 0xa1, 0xc2, 0xfb, 0x16, 0xb5, 0xa1, 0x33,
	0x2b, 0x78, 0x3a, 0x4a, 0x30, 0x73, 0xdc, 0x6c, 0x10, 0x21, 0x0a, 0x88, 0xcf, 0xfc, 0x16, 0x1f,
	0x0c, 0xd6, 0x5a, 0xc8, 0x5a, 0x53, 0x38, 0x9e, 0x9e, 0x2c, 0xd2, 0xcb, 0x51, 0xd2, 0xfe, 0xf7,
	0x6a, 0xab, 0x72, 0x55, 0x1b, 0xb5, 0xc9, 0x9e, 0xe5, 0x8f, 0xeb, 0xf1, 0xf7, 0x00, 0x66, 0xfd,
	0x64, 0xe4, 0x14, 0x42, 0xff, 0x66, 0x3e, 0xed, 0xd8, 0xdf, 0xc8, 0x1b, 0xc1, 0xba, 0x94, 0xf5,
	0xaf, 0xca, 0x2e, 0x5c, 0x47, 0x36, 0xb4, 0x12, 0x0a, 0xa1, 0x77, 0x64, 0x33, 0x8d, 0xb2, 0x01,
	0x92, 0x2f, 0x30, 0xaf, 0xf8, 0x12, 0x2b, 0x4d, 0xa7, 0x76, 0x80, 0xb7, 0xa3, 0xc3, 0x66, 0x37,
	0x96, 0xef, 0xac, 0x7b, 0x31, 0xf2, 0x02, 0xa2, 0x46, 0x89, 0x3a, 0x17, 0x0d, 0xaf, 0x6c, 0xea,
	0x51, 0xf6, 0xa7, 0x10, 0xbf, 0x81, 0xc5, 0x0e, 0x89, 0x1c, 0xc0, 0xf4, 0x2b, 0x6e, 0xec, 0x3c,
	0x51, 0xd6, 0x1f, 0xc9, 0x73, 0xd8, 0xeb, 0x78, 0xd5, 0x0e, 0x6e, 0x1d, 0x38, 0x0f, 0xce, 0x26,
	0xf1, 0x8f, 0x09, 0x84, 0xfe, 0x45, 0x08, 0x83, 0x59, 0xbf, 0xa0, 0xdb, 0x20, 0xdc, 0xf6, 0xb2,
	0x61, 0x7b, 0xd9, 0xe7, 0x61, 0x7b, 0x33, 0xdb, 0x47, 0x6e, 0x60, 0x56, 0x1a, 0xd3, 0xf8, 0xb5,
	0x3a, 0x1b, 0x35, 0xe9, 0xb5, 0x31, 0xcd, 0xb0, 0x09, 0x56, 0x25, 0xfe, 0x15, 0xc0, 0x62, 0xa7,
	0x4a, 0xfe, 0x87, 0x40, 0x14, 0x7e, 0x88, 0x40, 0x14, 0xe4, 0x10, 0xe6, 0x6b, 0x34, 0xa5, 0x2c,
	0xfc, 0x10, 0x1e, 0x91, 0x5b, 0x08, 0x4b, 0xe4, 0x05, 0xaa, 0x21, 0xf2, 0xab, 0xa7, 0x1a, 0x61,
	0xd7, 0x4e, 0xc7, 0x45, 0x3f, 0xa8, 0x12, 0x02, 0xb3, 0x86, 0x9b, 0xd2, 0xc7, 0x6e, 0xcf, 0x7d,
	0xad, 0x94, 0xda, 0xd0, 0x3d, 0x57, 0xeb, 0xcf, 0xbd, 0x41, 0x9d, 0x97, 0xb8, 0x46, 0x3a, 0x77,
	0x06, 0x1d, 0xea, 0xc3, 0xbf, 0x6f, 0x51, 0x6d, 0x68, 0xe8, 0xc2, 0xb7, 0x80, 0xc4, 0xb0, 0x7f,
	0xa7, 0xf8, 0x6a, 0x8d, 0xb5, 0xa1, 0xfb, 0xf6, 0x8f, 0x2d, 0xee, 0xd5, 0xb5, 0xf8, 0x86, 0x34,
	0x3a, 0x9e, 0x9c, 0x4c, 0x33, 0x7b, 0xee, 0xfb, 0xed, 0x43, 0xe4, 0xb2, 0xa2, 0xe0, 0xfa, 0x07,
	0x1c, 0x9f, 0xc3, 0x7f, 0xbb, 0xd6, 0x47, 0x2d, 0xc0, 0x3b, 0x38, 0xfc, 0xfb, 0x67, 0x33, 0x46,
	0x65, 0x39, 0xb7, 0x5e, 0x5e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x74, 0x5c, 0xa4, 0xe7, 0x63,
	0x05, 0x00, 0x00,
}
