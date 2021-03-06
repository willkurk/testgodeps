// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/afrouter.proto

package afrouter // import "github.com/willkurk/testgodeps/go/afrouter"

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
	return fileDescriptor_afrouter_24c5b5962011003a, []int{0}
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
	return fileDescriptor_afrouter_24c5b5962011003a, []int{1}
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
	return fileDescriptor_afrouter_24c5b5962011003a, []int{2}
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
	proto.RegisterFile("voltha_protos/afrouter.proto", fileDescriptor_afrouter_24c5b5962011003a)
}

var fileDescriptor_afrouter_24c5b5962011003a = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4b, 0xeb, 0x30,
	0x18, 0xc6, 0x4f, 0xb7, 0xae, 0xeb, 0xde, 0xc3, 0xc6, 0x08, 0x87, 0x43, 0x10, 0x91, 0xd1, 0xab,
	0x21, 0xb2, 0x82, 0x43, 0xf0, 0x56, 0xf7, 0x0d, 0xba, 0x3b, 0x6f, 0xa4, 0x4b, 0xd3, 0x2e, 0xb4,
	0x26, 0x25, 0x7f, 0x2a, 0xa2, 0x9f, 0xcd, 0xcf, 0x26, 0x49, 0x97, 0x6d, 0x82, 0xde, 0xbd, 0xbf,
	0xe7, 0xe9, 0x4b, 0x9f, 0x27, 0x09, 0x5c, 0x76, 0xa2, 0xd1, 0xfb, 0xfc, 0xb9, 0x95, 0x42, 0x0b,
	0x95, 0xe6, 0xa5, 0x14, 0x46, 0x53, 0xb9, 0x72, 0x8c, 0x62, 0xcf, 0xc9, 0x3d, 0x44, 0x19, 0x55,
	0xa6, 0xd1, 0x08, 0xc3, 0x58, 0x19, 0x42, 0xa8, 0x52, 0x38, 0x58, 0x04, 0xcb, 0x38, 0xf3, 0x88,
	0xfe, 0xc1, 0x88, 0x4a, 0x29, 0x24, 0x1e, 0x2c, 0x82, 0xe5, 0x24, 0xeb, 0x21, 0xf9, 0x0c, 0x20,
	0xdc, 0x08, 0xce, 0xd1, 0x7f, 0x88, 0x14, 0x95, 0x1d, 0x95, 0x6e, 0x6f, 0x92, 0x1d, 0x08, 0xcd,
	0x61, 0xd8, 0xd6, 0xd5, 0x61, 0xc9, 0x8e, 0x56, 0x51, 0x1d, 0xc1, 0xc3, 0x5e, 0x51, 0x1d, 0xb1,
	0x3f, 0x25, 0x8d, 0x51, 0x9a, 0x4a, 0x1c, 0x3a, 0xd5, 0xa3, 0x75, 0x76, 0x39, 0xa9, 0x29, 0x2f,
	0xf0, 0xa8, 0x77, 0x0e, 0x88, 0xae, 0x00, 0x88, 0xe0, 0x9c, 0x12, 0xcd, 0x04, 0xc7, 0x91, 0x33,
	0xcf, 0x14, 0x84, 0x20, 0xcc, 0x8b, 0x42, 0xe2, 0xb1, 0x73, 0xdc, 0x6c, 0xb5, 0x56, 0x48, 0x8d,
	0xe3, 0x45, 0xb0, 0x0c, 0x33, 0x37, 0x27, 0x1f, 0x10, 0x3f, 0x94, 0x25, 0xe3, 0x4c, 0xbf, 0xd9,
	0x0e, 0xfd, 0x81, 0xf8, 0x0e, 0x3d, 0xd9, 0xea, 0x6e, 0xf2, 0xd5, 0x1d, 0x9c, 0xa7, 0x1e, 0xfe,
	0x9a, 0x3a, 0xfc, 0x9e, 0x7a, 0x06, 0x03, 0xe6, 0xab, 0x0c, 0x58, 0x71, 0xfb, 0x0e, 0xd3, 0x8d,
	0xe0, 0x25, 0xab, 0x8c, 0xcc, 0x5d, 0xec, 0x35, 0x4c, 0xb7, 0x54, 0x6f, 0x4e, 0x3d, 0x66, 0xab,
	0xe3, 0xad, 0x59, 0xf5, 0x62, 0x7e, 0xe2, 0xfe, 0xca, 0x92, 0x3f, 0xe8, 0x0e, 0xfe, 0x6e, 0xa9,
	0x3e, 0xd6, 0x40, 0xa7, 0x4f, 0xbc, 0xf6, 0xd3, 0xda, 0xe3, 0xcd, 0xd3, 0x75, 0xc5, 0xf4, 0xde,
	0xec, 0x56, 0x44, 0xbc, 0xa4, 0xaf, 0xac, 0x69, 0x6a, 0x23, 0xeb, 0x54, 0x53, 0xa5, 0x2b, 0x51,
	0xd0, 0x56, 0xa5, 0x95, 0x38, 0xbe, 0x99, 0x5d, 0xe4, 0x1e, 0xcd, 0xfa, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x10, 0x39, 0x31, 0x15, 0x54, 0x02, 0x00, 0x00,
}
