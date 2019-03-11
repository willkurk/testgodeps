// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/afrouter.proto

package afrouter // import "github.com/opencord/voltha-protos/protos/afrouter"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_afrouter_fb8b9c067a15bb16, []int{0}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Conn struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Pkg                  string   `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Svc                  string   `protobuf:"bytes,3,opt,name=svc,proto3" json:"svc,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,5,opt,name=backend,proto3" json:"backend,omitempty"`
	Connection           string   `protobuf:"bytes,6,opt,name=connection,proto3" json:"connection,omitempty"`
	Addr                 string   `protobuf:"bytes,7,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 uint64   `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conn) Reset()         { *m = Conn{} }
func (m *Conn) String() string { return proto.CompactTextString(m) }
func (*Conn) ProtoMessage()    {}
func (*Conn) Descriptor() ([]byte, []int) {
	return fileDescriptor_afrouter_fb8b9c067a15bb16, []int{1}
}
func (m *Conn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conn.Unmarshal(m, b)
}
func (m *Conn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conn.Marshal(b, m, deterministic)
}
func (dst *Conn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conn.Merge(dst, src)
}
func (m *Conn) XXX_Size() int {
	return xxx_messageInfo_Conn.Size(m)
}
func (m *Conn) XXX_DiscardUnknown() {
	xxx_messageInfo_Conn.DiscardUnknown(m)
}

var xxx_messageInfo_Conn proto.InternalMessageInfo

func (m *Conn) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Conn) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *Conn) GetSvc() string {
	if m != nil {
		return m.Svc
	}
	return ""
}

func (m *Conn) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Conn) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Conn) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Conn) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Conn) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Affinity struct {
	Router               string   `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Cluster              string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,4,opt,name=backend,proto3" json:"backend,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Affinity) Reset()         { *m = Affinity{} }
func (m *Affinity) String() string { return proto.CompactTextString(m) }
func (*Affinity) ProtoMessage()    {}
func (*Affinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_afrouter_fb8b9c067a15bb16, []int{2}
}
func (m *Affinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Affinity.Unmarshal(m, b)
}
func (m *Affinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Affinity.Marshal(b, m, deterministic)
}
func (dst *Affinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Affinity.Merge(dst, src)
}
func (m *Affinity) XXX_Size() int {
	return xxx_messageInfo_Affinity.Size(m)
}
func (m *Affinity) XXX_DiscardUnknown() {
	xxx_messageInfo_Affinity.DiscardUnknown(m)
}

var xxx_messageInfo_Affinity proto.InternalMessageInfo

func (m *Affinity) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Affinity) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Affinity) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Affinity) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Affinity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "afrouter.Result")
	proto.RegisterType((*Conn)(nil), "afrouter.Conn")
	proto.RegisterType((*Affinity)(nil), "afrouter.Affinity")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error)
	SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetAffinity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	SetConnection(context.Context, *Conn) (*Result, error)
	SetAffinity(context.Context, *Affinity) (*Result, error)
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_SetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetConnection(ctx, req.(*Conn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_SetAffinity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Affinity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetAffinity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetAffinity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetAffinity(ctx, req.(*Affinity))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "afrouter.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConnection",
			Handler:    _Configuration_SetConnection_Handler,
		},
		{
			MethodName: "SetAffinity",
			Handler:    _Configuration_SetAffinity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/afrouter.proto",
}

func init() {
	proto.RegisterFile("voltha_protos/afrouter.proto", fileDescriptor_afrouter_fb8b9c067a15bb16)
}

var fileDescriptor_afrouter_fb8b9c067a15bb16 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xed, 0xd6, 0x75, 0xdd, 0xc9, 0xc6, 0x08, 0x22, 0x41, 0x44, 0x46, 0x9f, 0xf6, 0xe2,
	0x86, 0x0e, 0xc1, 0x57, 0xdd, 0x37, 0xe8, 0xde, 0x7c, 0x91, 0x2e, 0x4d, 0xb7, 0xb2, 0x99, 0x94,
	0x4b, 0x5a, 0x10, 0xfd, 0x6c, 0x7e, 0x36, 0x49, 0xd2, 0x6c, 0x53, 0xf4, 0xed, 0x7e, 0xff, 0xeb,
	0xd1, 0xff, 0x3f, 0x77, 0x70, 0xdd, 0xc8, 0xbd, 0xde, 0x66, 0xaf, 0x15, 0x4a, 0x2d, 0xd5, 0x3c,
	0x2b, 0x50, 0xd6, 0x9a, 0xe3, 0xcc, 0x32, 0x89, 0x3d, 0x27, 0x8f, 0x10, 0xa5, 0x5c, 0xd5, 0x7b,
	0x4d, 0x28, 0xf4, 0x55, 0xcd, 0x18, 0x57, 0x8a, 0x06, 0x93, 0x60, 0x1a, 0xa7, 0x1e, 0xc9, 0x05,
	0xf4, 0x38, 0xa2, 0x44, 0xda, 0x99, 0x04, 0xd3, 0x41, 0xea, 0x20, 0xf9, 0x0a, 0x20, 0x5c, 0x4a,
	0x21, 0xc8, 0x25, 0x44, 0x8a, 0x63, 0xc3, 0xd1, 0xce, 0x0d, 0xd2, 0x96, 0xc8, 0x18, 0xba, 0xd5,
	0x6e, 0xd3, 0x0e, 0x99, 0xd2, 0x28, 0xaa, 0x61, 0xb4, 0xeb, 0x14, 0xd5, 0x30, 0xf3, 0x53, 0xb6,
	0xaf, 0x95, 0xe6, 0x48, 0x43, 0xab, 0x7a, 0x34, 0x9d, 0x75, 0xc6, 0x76, 0x5c, 0xe4, 0xb4, 0xe7,
	0x3a, 0x2d, 0x92, 0x1b, 0x00, 0x26, 0x85, 0xe0, 0x4c, 0x97, 0x52, 0xd0, 0xc8, 0x36, 0x4f, 0x14,
	0x42, 0x20, 0xcc, 0xf2, 0x1c, 0x69, 0xdf, 0x76, 0x6c, 0x6d, 0xb4, 0x4a, 0xa2, 0xa6, 0xf1, 0x24,
	0x98, 0x86, 0xa9, 0xad, 0x93, 0x4f, 0x88, 0x9f, 0x8a, 0xa2, 0x14, 0xa5, 0x7e, 0x37, 0x19, 0xdc,
	0x83, 0xf8, 0x0c, 0x8e, 0x4c, 0x74, 0x5b, 0xf9, 0xe8, 0x16, 0x4e, 0x5d, 0x77, 0xff, 0x75, 0x1d,
	0xfe, 0x74, 0x3d, 0x82, 0x4e, 0xe9, 0xa3, 0x74, 0xca, 0xfc, 0xfe, 0x03, 0x86, 0x4b, 0x29, 0x8a,
	0x72, 0x53, 0x63, 0x66, 0x6d, 0x2f, 0x60, 0xb8, 0xe2, 0x7a, 0x79, 0xcc, 0x31, 0x9a, 0x1d, 0xb6,
	0x66, 0xd4, 0xab, 0xf1, 0x91, 0xdd, 0xca, 0x92, 0x33, 0xf2, 0x00, 0xe7, 0x2b, 0xae, 0x0f, 0x31,
	0xc8, 0xf1, 0x13, 0xaf, 0xfd, 0x35, 0xf6, 0xbc, 0x78, 0xb9, 0xdb, 0x94, 0x7a, 0x5b, 0xaf, 0x67,
	0x4c, 0xbe, 0xcd, 0x65, 0xc5, 0x05, 0x93, 0x98, 0xcf, 0xdd, 0xcd, 0xdc, 0xb6, 0x37, 0xf3, 0xeb,
	0x74, 0xd6, 0x91, 0x15, 0x16, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x26, 0xcf, 0x45, 0x5b,
	0x02, 0x00, 0x00,
}
