// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/schema.proto

package schema // import "github.com/opencord/voltha-protos/protos/schema"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Contains the name and content of a *.proto file
type ProtoFile struct {
	FileName             string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Proto                string   `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	Descriptor_          []byte   `protobuf:"bytes,3,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoFile) Reset()         { *m = ProtoFile{} }
func (m *ProtoFile) String() string { return proto.CompactTextString(m) }
func (*ProtoFile) ProtoMessage()    {}
func (*ProtoFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_1eb43ba0f7460ca6, []int{0}
}
func (m *ProtoFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoFile.Unmarshal(m, b)
}
func (m *ProtoFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoFile.Marshal(b, m, deterministic)
}
func (dst *ProtoFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoFile.Merge(dst, src)
}
func (m *ProtoFile) XXX_Size() int {
	return xxx_messageInfo_ProtoFile.Size(m)
}
func (m *ProtoFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoFile.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoFile proto.InternalMessageInfo

func (m *ProtoFile) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ProtoFile) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *ProtoFile) GetDescriptor_() []byte {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

// Proto files and compiled descriptors for this interface
type Schemas struct {
	// Proto files
	Protos []*ProtoFile `protobuf:"bytes,1,rep,name=protos,proto3" json:"protos,omitempty"`
	// Proto file name from which swagger.json shall be generated
	SwaggerFrom string `protobuf:"bytes,2,opt,name=swagger_from,json=swaggerFrom,proto3" json:"swagger_from,omitempty"`
	// Proto file name from which yang schemas shall be generated
	YangFrom             string   `protobuf:"bytes,3,opt,name=yang_from,json=yangFrom,proto3" json:"yang_from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Schemas) Reset()         { *m = Schemas{} }
func (m *Schemas) String() string { return proto.CompactTextString(m) }
func (*Schemas) ProtoMessage()    {}
func (*Schemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_1eb43ba0f7460ca6, []int{1}
}
func (m *Schemas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schemas.Unmarshal(m, b)
}
func (m *Schemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schemas.Marshal(b, m, deterministic)
}
func (dst *Schemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schemas.Merge(dst, src)
}
func (m *Schemas) XXX_Size() int {
	return xxx_messageInfo_Schemas.Size(m)
}
func (m *Schemas) XXX_DiscardUnknown() {
	xxx_messageInfo_Schemas.DiscardUnknown(m)
}

var xxx_messageInfo_Schemas proto.InternalMessageInfo

func (m *Schemas) GetProtos() []*ProtoFile {
	if m != nil {
		return m.Protos
	}
	return nil
}

func (m *Schemas) GetSwaggerFrom() string {
	if m != nil {
		return m.SwaggerFrom
	}
	return ""
}

func (m *Schemas) GetYangFrom() string {
	if m != nil {
		return m.YangFrom
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoFile)(nil), "schema.ProtoFile")
	proto.RegisterType((*Schemas)(nil), "schema.Schemas")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchemaServiceClient interface {
	// Return active grpc schemas
	GetSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Schemas, error)
}

type schemaServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchemaServiceClient(cc *grpc.ClientConn) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) GetSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Schemas, error) {
	out := new(Schemas)
	err := c.cc.Invoke(ctx, "/schema.SchemaService/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
type SchemaServiceServer interface {
	// Return active grpc schemas
	GetSchema(context.Context, *empty.Empty) (*Schemas, error)
}

func RegisterSchemaServiceServer(s *grpc.Server, srv SchemaServiceServer) {
	s.RegisterService(&_SchemaService_serviceDesc, srv)
}

func _SchemaService_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.SchemaService/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchema(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchemaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchema",
			Handler:    _SchemaService_GetSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/schema.proto",
}

func init() { proto.RegisterFile("voltha_protos/schema.proto", fileDescriptor_schema_1eb43ba0f7460ca6) }

var fileDescriptor_schema_1eb43ba0f7460ca6 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x2d, 0xaf, 0x7d, 0x99, 0x56, 0x8a, 0x83, 0x48, 0x48, 0x45, 0x6a, 0x57, 0x75, 0x61,
	0x06, 0xeb, 0x1f, 0x08, 0xad, 0x3b, 0x91, 0x16, 0x5c, 0xb8, 0xb0, 0x4c, 0xd3, 0xdb, 0x64, 0x20,
	0x93, 0x1b, 0x26, 0xd3, 0x4a, 0xb7, 0xfe, 0x82, 0x9f, 0xe6, 0x2f, 0xf8, 0x21, 0x32, 0xb9, 0x53,
	0xd1, 0xdd, 0x9c, 0x73, 0xee, 0x70, 0xcf, 0x3d, 0x87, 0xc5, 0x7b, 0x2c, 0x6c, 0x2e, 0x57, 0x95,
	0x41, 0x8b, 0xb5, 0xa8, 0xd3, 0x1c, 0xb4, 0x4c, 0x1a, 0xc4, 0x3b, 0x84, 0xe2, 0x8b, 0x0c, 0x31,
	0x2b, 0x40, 0xc8, 0x4a, 0x09, 0x59, 0x96, 0x68, 0xa5, 0x55, 0x58, 0xd6, 0x34, 0x15, 0x0f, 0xbd,
	0xda, 0xa0, 0xf5, 0x6e, 0x2b, 0x40, 0x57, 0xf6, 0x40, 0xe2, 0xf8, 0x95, 0x85, 0x4f, 0xee, 0x31,
	0x57, 0x05, 0xf0, 0x21, 0x0b, 0xb7, 0xaa, 0x80, 0x55, 0x29, 0x35, 0x44, 0xc1, 0x28, 0x98, 0x84,
	0x8b, 0xff, 0x8e, 0x78, 0x94, 0x1a, 0xf8, 0x19, 0xfb, 0xd7, 0x7c, 0x89, 0x5a, 0x8d, 0x40, 0x80,
	0x5f, 0x32, 0xb6, 0x81, 0x3a, 0x35, 0xaa, 0xb2, 0x68, 0xa2, 0xf6, 0x28, 0x98, 0xf4, 0x17, 0xbf,
	0x98, 0xb1, 0x65, 0xdd, 0x65, 0x63, 0xb2, 0xe6, 0xd7, 0xac, 0x43, 0x47, 0x44, 0xc1, 0xa8, 0x3d,
	0xe9, 0x4d, 0x4f, 0x13, 0x7f, 0xcc, 0x8f, 0x81, 0x85, 0x1f, 0xe0, 0x57, 0xac, 0x5f, 0xbf, 0xc9,
	0x2c, 0x03, 0xb3, 0xda, 0x1a, 0xd4, 0x7e, 0x65, 0xcf, 0x73, 0x73, 0x83, 0xda, 0x79, 0x3d, 0xc8,
	0x32, 0x23, 0xbd, 0x4d, 0x5e, 0x1d, 0xe1, 0xc4, 0xe9, 0x33, 0x3b, 0xa1, 0xad, 0x4b, 0x30, 0x7b,
	0x95, 0x02, 0x9f, 0xb1, 0xf0, 0x01, 0x2c, 0x71, 0xfc, 0x3c, 0xa1, 0x44, 0x92, 0x63, 0x22, 0xc9,
	0xcc, 0x25, 0x12, 0x0f, 0x8e, 0x86, 0xbc, 0xe3, 0xf1, 0xe0, 0xfd, 0xf3, 0xeb, 0xa3, 0x15, 0xf2,
	0xae, 0x8f, 0xfd, 0xfe, 0xf6, 0x45, 0x64, 0xca, 0xe6, 0xbb, 0x75, 0x92, 0xa2, 0x16, 0x58, 0x41,
	0x99, 0xa2, 0xd9, 0x08, 0xaa, 0xe8, 0xc6, 0x57, 0xf4, 0xa7, 0xa9, 0x35, 0x9d, 0x74, 0xf7, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x60, 0x96, 0xf7, 0xf4, 0xc8, 0x01, 0x00, 0x00,
}