// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/ponsim.proto

package voltha // import "github.com/opencord/voltha-protos/protos/voltha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import openflow_13 "github.com/opencord/voltha-protos/protos/openflow_13"

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

type PonSimOnuDeviceInfo struct {
	UniPort              int32    `protobuf:"varint,1,opt,name=uni_port,json=uniPort,proto3" json:"uni_port,omitempty"`
	SerialNumber         string   `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PonSimOnuDeviceInfo) Reset()         { *m = PonSimOnuDeviceInfo{} }
func (m *PonSimOnuDeviceInfo) String() string { return proto.CompactTextString(m) }
func (*PonSimOnuDeviceInfo) ProtoMessage()    {}
func (*PonSimOnuDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{0}
}
func (m *PonSimOnuDeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimOnuDeviceInfo.Unmarshal(m, b)
}
func (m *PonSimOnuDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimOnuDeviceInfo.Marshal(b, m, deterministic)
}
func (dst *PonSimOnuDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimOnuDeviceInfo.Merge(dst, src)
}
func (m *PonSimOnuDeviceInfo) XXX_Size() int {
	return xxx_messageInfo_PonSimOnuDeviceInfo.Size(m)
}
func (m *PonSimOnuDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimOnuDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimOnuDeviceInfo proto.InternalMessageInfo

func (m *PonSimOnuDeviceInfo) GetUniPort() int32 {
	if m != nil {
		return m.UniPort
	}
	return 0
}

func (m *PonSimOnuDeviceInfo) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

type PonSimDeviceInfo struct {
	NniPort              int32                  `protobuf:"varint,1,opt,name=nni_port,json=nniPort,proto3" json:"nni_port,omitempty"`
	Onus                 []*PonSimOnuDeviceInfo `protobuf:"bytes,2,rep,name=onus,proto3" json:"onus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PonSimDeviceInfo) Reset()         { *m = PonSimDeviceInfo{} }
func (m *PonSimDeviceInfo) String() string { return proto.CompactTextString(m) }
func (*PonSimDeviceInfo) ProtoMessage()    {}
func (*PonSimDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{1}
}
func (m *PonSimDeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimDeviceInfo.Unmarshal(m, b)
}
func (m *PonSimDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimDeviceInfo.Marshal(b, m, deterministic)
}
func (dst *PonSimDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimDeviceInfo.Merge(dst, src)
}
func (m *PonSimDeviceInfo) XXX_Size() int {
	return xxx_messageInfo_PonSimDeviceInfo.Size(m)
}
func (m *PonSimDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimDeviceInfo proto.InternalMessageInfo

func (m *PonSimDeviceInfo) GetNniPort() int32 {
	if m != nil {
		return m.NniPort
	}
	return 0
}

func (m *PonSimDeviceInfo) GetOnus() []*PonSimOnuDeviceInfo {
	if m != nil {
		return m.Onus
	}
	return nil
}

type FlowTable struct {
	Port                 int32                       `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Flows                []*openflow_13.OfpFlowStats `protobuf:"bytes,2,rep,name=flows,proto3" json:"flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FlowTable) Reset()         { *m = FlowTable{} }
func (m *FlowTable) String() string { return proto.CompactTextString(m) }
func (*FlowTable) ProtoMessage()    {}
func (*FlowTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{2}
}
func (m *FlowTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowTable.Unmarshal(m, b)
}
func (m *FlowTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowTable.Marshal(b, m, deterministic)
}
func (dst *FlowTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowTable.Merge(dst, src)
}
func (m *FlowTable) XXX_Size() int {
	return xxx_messageInfo_FlowTable.Size(m)
}
func (m *FlowTable) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowTable.DiscardUnknown(m)
}

var xxx_messageInfo_FlowTable proto.InternalMessageInfo

func (m *FlowTable) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FlowTable) GetFlows() []*openflow_13.OfpFlowStats {
	if m != nil {
		return m.Flows
	}
	return nil
}

type PonSimFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	OutPort              int32    `protobuf:"varint,3,opt,name=out_port,json=outPort,proto3" json:"out_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PonSimFrame) Reset()         { *m = PonSimFrame{} }
func (m *PonSimFrame) String() string { return proto.CompactTextString(m) }
func (*PonSimFrame) ProtoMessage()    {}
func (*PonSimFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{3}
}
func (m *PonSimFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimFrame.Unmarshal(m, b)
}
func (m *PonSimFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimFrame.Marshal(b, m, deterministic)
}
func (dst *PonSimFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimFrame.Merge(dst, src)
}
func (m *PonSimFrame) XXX_Size() int {
	return xxx_messageInfo_PonSimFrame.Size(m)
}
func (m *PonSimFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimFrame.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimFrame proto.InternalMessageInfo

func (m *PonSimFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PonSimFrame) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PonSimFrame) GetOutPort() int32 {
	if m != nil {
		return m.OutPort
	}
	return 0
}

type PonSimPacketCounter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PonSimPacketCounter) Reset()         { *m = PonSimPacketCounter{} }
func (m *PonSimPacketCounter) String() string { return proto.CompactTextString(m) }
func (*PonSimPacketCounter) ProtoMessage()    {}
func (*PonSimPacketCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{4}
}
func (m *PonSimPacketCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimPacketCounter.Unmarshal(m, b)
}
func (m *PonSimPacketCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimPacketCounter.Marshal(b, m, deterministic)
}
func (dst *PonSimPacketCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimPacketCounter.Merge(dst, src)
}
func (m *PonSimPacketCounter) XXX_Size() int {
	return xxx_messageInfo_PonSimPacketCounter.Size(m)
}
func (m *PonSimPacketCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimPacketCounter.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimPacketCounter proto.InternalMessageInfo

func (m *PonSimPacketCounter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PonSimPacketCounter) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PonSimPortMetrics struct {
	PortName             string                 `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	Packets              []*PonSimPacketCounter `protobuf:"bytes,2,rep,name=packets,proto3" json:"packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PonSimPortMetrics) Reset()         { *m = PonSimPortMetrics{} }
func (m *PonSimPortMetrics) String() string { return proto.CompactTextString(m) }
func (*PonSimPortMetrics) ProtoMessage()    {}
func (*PonSimPortMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{5}
}
func (m *PonSimPortMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimPortMetrics.Unmarshal(m, b)
}
func (m *PonSimPortMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimPortMetrics.Marshal(b, m, deterministic)
}
func (dst *PonSimPortMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimPortMetrics.Merge(dst, src)
}
func (m *PonSimPortMetrics) XXX_Size() int {
	return xxx_messageInfo_PonSimPortMetrics.Size(m)
}
func (m *PonSimPortMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimPortMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimPortMetrics proto.InternalMessageInfo

func (m *PonSimPortMetrics) GetPortName() string {
	if m != nil {
		return m.PortName
	}
	return ""
}

func (m *PonSimPortMetrics) GetPackets() []*PonSimPacketCounter {
	if m != nil {
		return m.Packets
	}
	return nil
}

type PonSimMetrics struct {
	Device               string               `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Metrics              []*PonSimPortMetrics `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PonSimMetrics) Reset()         { *m = PonSimMetrics{} }
func (m *PonSimMetrics) String() string { return proto.CompactTextString(m) }
func (*PonSimMetrics) ProtoMessage()    {}
func (*PonSimMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{6}
}
func (m *PonSimMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimMetrics.Unmarshal(m, b)
}
func (m *PonSimMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimMetrics.Marshal(b, m, deterministic)
}
func (dst *PonSimMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimMetrics.Merge(dst, src)
}
func (m *PonSimMetrics) XXX_Size() int {
	return xxx_messageInfo_PonSimMetrics.Size(m)
}
func (m *PonSimMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimMetrics proto.InternalMessageInfo

func (m *PonSimMetrics) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *PonSimMetrics) GetMetrics() []*PonSimPortMetrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type PonSimMetricsRequest struct {
	Port                 int32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PonSimMetricsRequest) Reset()         { *m = PonSimMetricsRequest{} }
func (m *PonSimMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*PonSimMetricsRequest) ProtoMessage()    {}
func (*PonSimMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ponsim_3e3faf884790af02, []int{7}
}
func (m *PonSimMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PonSimMetricsRequest.Unmarshal(m, b)
}
func (m *PonSimMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PonSimMetricsRequest.Marshal(b, m, deterministic)
}
func (dst *PonSimMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PonSimMetricsRequest.Merge(dst, src)
}
func (m *PonSimMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_PonSimMetricsRequest.Size(m)
}
func (m *PonSimMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PonSimMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PonSimMetricsRequest proto.InternalMessageInfo

func (m *PonSimMetricsRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*PonSimOnuDeviceInfo)(nil), "voltha.PonSimOnuDeviceInfo")
	proto.RegisterType((*PonSimDeviceInfo)(nil), "voltha.PonSimDeviceInfo")
	proto.RegisterType((*FlowTable)(nil), "voltha.FlowTable")
	proto.RegisterType((*PonSimFrame)(nil), "voltha.PonSimFrame")
	proto.RegisterType((*PonSimPacketCounter)(nil), "voltha.PonSimPacketCounter")
	proto.RegisterType((*PonSimPortMetrics)(nil), "voltha.PonSimPortMetrics")
	proto.RegisterType((*PonSimMetrics)(nil), "voltha.PonSimMetrics")
	proto.RegisterType((*PonSimMetricsRequest)(nil), "voltha.PonSimMetricsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PonSimClient is the client API for PonSim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PonSimClient interface {
	SendFrame(ctx context.Context, in *PonSimFrame, opts ...grpc.CallOption) (*empty.Empty, error)
	ReceiveFrames(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (PonSim_ReceiveFramesClient, error)
	GetDeviceInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PonSimDeviceInfo, error)
	UpdateFlowTable(ctx context.Context, in *FlowTable, opts ...grpc.CallOption) (*empty.Empty, error)
	GetStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PonSimMetrics, error)
}

type ponSimClient struct {
	cc *grpc.ClientConn
}

func NewPonSimClient(cc *grpc.ClientConn) PonSimClient {
	return &ponSimClient{cc}
}

func (c *ponSimClient) SendFrame(ctx context.Context, in *PonSimFrame, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/voltha.PonSim/SendFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ponSimClient) ReceiveFrames(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (PonSim_ReceiveFramesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PonSim_serviceDesc.Streams[0], "/voltha.PonSim/ReceiveFrames", opts...)
	if err != nil {
		return nil, err
	}
	x := &ponSimReceiveFramesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PonSim_ReceiveFramesClient interface {
	Recv() (*PonSimFrame, error)
	grpc.ClientStream
}

type ponSimReceiveFramesClient struct {
	grpc.ClientStream
}

func (x *ponSimReceiveFramesClient) Recv() (*PonSimFrame, error) {
	m := new(PonSimFrame)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ponSimClient) GetDeviceInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PonSimDeviceInfo, error) {
	out := new(PonSimDeviceInfo)
	err := c.cc.Invoke(ctx, "/voltha.PonSim/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ponSimClient) UpdateFlowTable(ctx context.Context, in *FlowTable, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/voltha.PonSim/UpdateFlowTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ponSimClient) GetStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PonSimMetrics, error) {
	out := new(PonSimMetrics)
	err := c.cc.Invoke(ctx, "/voltha.PonSim/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PonSimServer is the server API for PonSim service.
type PonSimServer interface {
	SendFrame(context.Context, *PonSimFrame) (*empty.Empty, error)
	ReceiveFrames(*empty.Empty, PonSim_ReceiveFramesServer) error
	GetDeviceInfo(context.Context, *empty.Empty) (*PonSimDeviceInfo, error)
	UpdateFlowTable(context.Context, *FlowTable) (*empty.Empty, error)
	GetStats(context.Context, *empty.Empty) (*PonSimMetrics, error)
}

func RegisterPonSimServer(s *grpc.Server, srv PonSimServer) {
	s.RegisterService(&_PonSim_serviceDesc, srv)
}

func _PonSim_SendFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PonSimFrame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PonSimServer).SendFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voltha.PonSim/SendFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PonSimServer).SendFrame(ctx, req.(*PonSimFrame))
	}
	return interceptor(ctx, in, info, handler)
}

func _PonSim_ReceiveFrames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PonSimServer).ReceiveFrames(m, &ponSimReceiveFramesServer{stream})
}

type PonSim_ReceiveFramesServer interface {
	Send(*PonSimFrame) error
	grpc.ServerStream
}

type ponSimReceiveFramesServer struct {
	grpc.ServerStream
}

func (x *ponSimReceiveFramesServer) Send(m *PonSimFrame) error {
	return x.ServerStream.SendMsg(m)
}

func _PonSim_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PonSimServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voltha.PonSim/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PonSimServer).GetDeviceInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PonSim_UpdateFlowTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PonSimServer).UpdateFlowTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voltha.PonSim/UpdateFlowTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PonSimServer).UpdateFlowTable(ctx, req.(*FlowTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _PonSim_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PonSimServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voltha.PonSim/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PonSimServer).GetStats(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PonSim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "voltha.PonSim",
	HandlerType: (*PonSimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFrame",
			Handler:    _PonSim_SendFrame_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _PonSim_GetDeviceInfo_Handler,
		},
		{
			MethodName: "UpdateFlowTable",
			Handler:    _PonSim_UpdateFlowTable_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _PonSim_GetStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveFrames",
			Handler:       _PonSim_ReceiveFrames_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "voltha_protos/ponsim.proto",
}

func init() { proto.RegisterFile("voltha_protos/ponsim.proto", fileDescriptor_ponsim_3e3faf884790af02) }

var fileDescriptor_ponsim_3e3faf884790af02 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf3, 0xa7, 0xf9, 0x37, 0x6d, 0x80, 0x6e, 0x4b, 0x95, 0x26, 0x07, 0xa2, 0xe5, 0x12,
	0x21, 0x61, 0x93, 0x46, 0x5c, 0x40, 0x02, 0x44, 0xa1, 0x15, 0x07, 0x4a, 0xb4, 0xa1, 0x17, 0x84,
	0xb0, 0x1c, 0x7b, 0x92, 0x5a, 0xd8, 0xbb, 0xc6, 0xde, 0x4d, 0xd5, 0x6f, 0xc8, 0xc7, 0x42, 0xde,
	0xb5, 0x49, 0x5c, 0x25, 0x3d, 0x65, 0x67, 0xfc, 0xf2, 0x7b, 0xf3, 0xd6, 0x93, 0x40, 0x7f, 0x25,
	0x42, 0x79, 0xe3, 0x3a, 0x71, 0x22, 0xa4, 0x48, 0xed, 0x58, 0xf0, 0x34, 0x88, 0x2c, 0x5d, 0x91,
	0xa6, 0x79, 0xd6, 0x1f, 0x2c, 0x85, 0x58, 0x86, 0x68, 0xeb, 0xee, 0x5c, 0x2d, 0x6c, 0x8c, 0x62,
	0x79, 0x67, 0x44, 0xfd, 0x67, 0x65, 0x80, 0x88, 0x91, 0x2f, 0x42, 0x71, 0xeb, 0x8c, 0x27, 0x46,
	0x40, 0xaf, 0xe1, 0x68, 0x2a, 0xf8, 0x2c, 0x88, 0xbe, 0x71, 0xf5, 0x09, 0x57, 0x81, 0x87, 0x5f,
	0xf8, 0x42, 0x90, 0x53, 0x68, 0x2b, 0x1e, 0x38, 0xb1, 0x48, 0x64, 0xaf, 0x3a, 0xac, 0x8e, 0x1a,
	0xac, 0xa5, 0x78, 0x30, 0x15, 0x89, 0x24, 0xcf, 0xa1, 0x9b, 0x62, 0x12, 0xb8, 0xa1, 0xc3, 0x55,
	0x34, 0xc7, 0xa4, 0x57, 0x1b, 0x56, 0x47, 0x1d, 0x76, 0x60, 0x9a, 0x57, 0xba, 0x47, 0x7f, 0xc1,
	0x13, 0x83, 0x2d, 0x33, 0xf9, 0x3d, 0x26, 0xcf, 0x99, 0x36, 0xec, 0x09, 0xae, 0xd2, 0x5e, 0x6d,
	0x58, 0x1f, 0xed, 0x9f, 0x0d, 0x2c, 0x33, 0xb5, 0xb5, 0x65, 0x32, 0xa6, 0x85, 0x94, 0x41, 0xe7,
	0x22, 0x14, 0xb7, 0xdf, 0xdd, 0x79, 0x88, 0x84, 0xc0, 0xde, 0x06, 0x54, 0x9f, 0xc9, 0x18, 0x1a,
	0x59, 0xd0, 0x35, 0x72, 0x33, 0xba, 0x58, 0xc4, 0x8e, 0x3e, 0xa7, 0xd2, 0x95, 0x29, 0x33, 0x4a,
	0xca, 0x60, 0xdf, 0x18, 0x5e, 0x24, 0x6e, 0x84, 0xe4, 0x11, 0xd4, 0x02, 0x5f, 0x33, 0x3b, 0xac,
	0x16, 0xf8, 0xa4, 0x07, 0xad, 0xd8, 0xbd, 0x0b, 0x85, 0xeb, 0xeb, 0xc4, 0x07, 0xac, 0x28, 0xb3,
	0x60, 0x42, 0x49, 0x13, 0xac, 0x6e, 0x82, 0x09, 0x25, 0xb3, 0x60, 0xf4, 0x7d, 0x71, 0xbd, 0x53,
	0xd7, 0xfb, 0x8d, 0xf2, 0x5c, 0x28, 0x2e, 0x31, 0xc9, 0x26, 0xe6, 0x6e, 0x84, 0x39, 0x5d, 0x9f,
	0xc9, 0x31, 0x34, 0x56, 0x6e, 0xa8, 0x50, 0xd3, 0xeb, 0xcc, 0x14, 0x74, 0x09, 0x87, 0x39, 0x40,
	0x24, 0xf2, 0x2b, 0xca, 0x24, 0xf0, 0x52, 0x32, 0x80, 0x4e, 0x66, 0xe6, 0x6c, 0x30, 0xda, 0x59,
	0xe3, 0x2a, 0xe3, 0xbc, 0xce, 0xe6, 0xcc, 0xcc, 0x76, 0x5c, 0x67, 0x69, 0x12, 0x56, 0x68, 0xe9,
	0x4f, 0xe8, 0x9a, 0xe7, 0x85, 0xc9, 0x09, 0x34, 0x7d, 0x7d, 0xed, 0xb9, 0x43, 0x5e, 0x91, 0x09,
	0xb4, 0x22, 0x23, 0xc9, 0xf9, 0xa7, 0xf7, 0xf8, 0xeb, 0x41, 0x59, 0xa1, 0xa4, 0x2f, 0xe0, 0xb8,
	0x44, 0x67, 0xf8, 0x47, 0x61, 0x2a, 0xb7, 0xbd, 0xba, 0xb3, 0xbf, 0x35, 0x68, 0x1a, 0x31, 0x79,
	0x03, 0x9d, 0x19, 0x72, 0xdf, 0xbc, 0x90, 0xa3, 0xb2, 0x8f, 0x6e, 0xf6, 0x4f, 0x2c, 0xb3, 0xfe,
	0x56, 0xb1, 0xfe, 0xd6, 0xe7, 0x6c, 0xfd, 0x69, 0x85, 0x7c, 0x80, 0x2e, 0x43, 0x0f, 0x83, 0x15,
	0x6a, 0x65, 0x4a, 0x76, 0x48, 0xfb, 0xdb, 0xb8, 0xb4, 0xf2, 0xaa, 0x4a, 0xce, 0xa1, 0x7b, 0x89,
	0x72, 0x63, 0x83, 0x77, 0x11, 0x7a, 0x65, 0xc2, 0xfa, 0x1b, 0xb4, 0x42, 0xde, 0xc1, 0xe3, 0xeb,
	0xd8, 0x77, 0x25, 0xae, 0xf7, 0xf5, 0xb0, 0x90, 0xff, 0x6f, 0x3d, 0x10, 0xe3, 0x2d, 0xb4, 0x2f,
	0x51, 0xce, 0xb2, 0x45, 0xdd, 0xe9, 0xff, 0xb4, 0xec, 0x9f, 0xdf, 0x31, 0xad, 0x7c, 0x1c, 0xff,
	0xb0, 0x97, 0x81, 0xbc, 0x51, 0x73, 0xcb, 0x13, 0x91, 0xfe, 0xf5, 0x7b, 0x22, 0xf1, 0x6d, 0xa3,
	0x7e, 0x59, 0xfc, 0xab, 0x98, 0x0f, 0xd3, 0x9c, 0x37, 0x75, 0x39, 0xf9, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x35, 0x5a, 0x27, 0x7b, 0x04, 0x00, 0x00,
}
