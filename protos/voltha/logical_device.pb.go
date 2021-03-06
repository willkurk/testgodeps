// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/logical_device.proto

package voltha // import "github.com/opencord/voltha-protos/protos/voltha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/opencord/voltha-protos/protos/common"
import openflow_13 "github.com/opencord/voltha-protos/protos/openflow_13"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogicalPortId struct {
	// unique id of logical device
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// id of the port on the logical device
	PortId               string   `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogicalPortId) Reset()         { *m = LogicalPortId{} }
func (m *LogicalPortId) String() string { return proto.CompactTextString(m) }
func (*LogicalPortId) ProtoMessage()    {}
func (*LogicalPortId) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_a1c7cc8d1ed12585, []int{0}
}
func (m *LogicalPortId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPortId.Unmarshal(m, b)
}
func (m *LogicalPortId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPortId.Marshal(b, m, deterministic)
}
func (dst *LogicalPortId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPortId.Merge(dst, src)
}
func (m *LogicalPortId) XXX_Size() int {
	return xxx_messageInfo_LogicalPortId.Size(m)
}
func (m *LogicalPortId) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPortId.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPortId proto.InternalMessageInfo

func (m *LogicalPortId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalPortId) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

type LogicalPort struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OfpPort              *openflow_13.OfpPort `protobuf:"bytes,2,opt,name=ofp_port,json=ofpPort,proto3" json:"ofp_port,omitempty"`
	DeviceId             string               `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	DevicePortNo         uint32               `protobuf:"varint,4,opt,name=device_port_no,json=devicePortNo,proto3" json:"device_port_no,omitempty"`
	RootPort             bool                 `protobuf:"varint,5,opt,name=root_port,json=rootPort,proto3" json:"root_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogicalPort) Reset()         { *m = LogicalPort{} }
func (m *LogicalPort) String() string { return proto.CompactTextString(m) }
func (*LogicalPort) ProtoMessage()    {}
func (*LogicalPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_a1c7cc8d1ed12585, []int{1}
}
func (m *LogicalPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPort.Unmarshal(m, b)
}
func (m *LogicalPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPort.Marshal(b, m, deterministic)
}
func (dst *LogicalPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPort.Merge(dst, src)
}
func (m *LogicalPort) XXX_Size() int {
	return xxx_messageInfo_LogicalPort.Size(m)
}
func (m *LogicalPort) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPort.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPort proto.InternalMessageInfo

func (m *LogicalPort) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalPort) GetOfpPort() *openflow_13.OfpPort {
	if m != nil {
		return m.OfpPort
	}
	return nil
}

func (m *LogicalPort) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LogicalPort) GetDevicePortNo() uint32 {
	if m != nil {
		return m.DevicePortNo
	}
	return 0
}

func (m *LogicalPort) GetRootPort() bool {
	if m != nil {
		return m.RootPort
	}
	return false
}

type LogicalPorts struct {
	Items                []*LogicalPort `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LogicalPorts) Reset()         { *m = LogicalPorts{} }
func (m *LogicalPorts) String() string { return proto.CompactTextString(m) }
func (*LogicalPorts) ProtoMessage()    {}
func (*LogicalPorts) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_a1c7cc8d1ed12585, []int{2}
}
func (m *LogicalPorts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPorts.Unmarshal(m, b)
}
func (m *LogicalPorts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPorts.Marshal(b, m, deterministic)
}
func (dst *LogicalPorts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPorts.Merge(dst, src)
}
func (m *LogicalPorts) XXX_Size() int {
	return xxx_messageInfo_LogicalPorts.Size(m)
}
func (m *LogicalPorts) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPorts.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPorts proto.InternalMessageInfo

func (m *LogicalPorts) GetItems() []*LogicalPort {
	if m != nil {
		return m.Items
	}
	return nil
}

type LogicalDevice struct {
	// unique id of logical device
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// unique datapath id for the logical device (used by the SDN controller)
	DatapathId uint64 `protobuf:"varint,2,opt,name=datapath_id,json=datapathId,proto3" json:"datapath_id,omitempty"`
	// device description
	Desc *openflow_13.OfpDesc `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// device features
	SwitchFeatures *openflow_13.OfpSwitchFeatures `protobuf:"bytes,4,opt,name=switch_features,json=switchFeatures,proto3" json:"switch_features,omitempty"`
	// name of the root device anchoring logical device
	RootDeviceId string `protobuf:"bytes,5,opt,name=root_device_id,json=rootDeviceId,proto3" json:"root_device_id,omitempty"`
	// logical device ports
	Ports []*LogicalPort `protobuf:"bytes,128,rep,name=ports,proto3" json:"ports,omitempty"`
	// flows configured on the logical device
	Flows *openflow_13.Flows `protobuf:"bytes,129,opt,name=flows,proto3" json:"flows,omitempty"`
	// flow groups configured on the logical device
	FlowGroups           *openflow_13.FlowGroups `protobuf:"bytes,130,opt,name=flow_groups,json=flowGroups,proto3" json:"flow_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LogicalDevice) Reset()         { *m = LogicalDevice{} }
func (m *LogicalDevice) String() string { return proto.CompactTextString(m) }
func (*LogicalDevice) ProtoMessage()    {}
func (*LogicalDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_a1c7cc8d1ed12585, []int{3}
}
func (m *LogicalDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalDevice.Unmarshal(m, b)
}
func (m *LogicalDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalDevice.Marshal(b, m, deterministic)
}
func (dst *LogicalDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalDevice.Merge(dst, src)
}
func (m *LogicalDevice) XXX_Size() int {
	return xxx_messageInfo_LogicalDevice.Size(m)
}
func (m *LogicalDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalDevice.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalDevice proto.InternalMessageInfo

func (m *LogicalDevice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalDevice) GetDatapathId() uint64 {
	if m != nil {
		return m.DatapathId
	}
	return 0
}

func (m *LogicalDevice) GetDesc() *openflow_13.OfpDesc {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *LogicalDevice) GetSwitchFeatures() *openflow_13.OfpSwitchFeatures {
	if m != nil {
		return m.SwitchFeatures
	}
	return nil
}

func (m *LogicalDevice) GetRootDeviceId() string {
	if m != nil {
		return m.RootDeviceId
	}
	return ""
}

func (m *LogicalDevice) GetPorts() []*LogicalPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *LogicalDevice) GetFlows() *openflow_13.Flows {
	if m != nil {
		return m.Flows
	}
	return nil
}

func (m *LogicalDevice) GetFlowGroups() *openflow_13.FlowGroups {
	if m != nil {
		return m.FlowGroups
	}
	return nil
}

type LogicalDevices struct {
	Items                []*LogicalDevice `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LogicalDevices) Reset()         { *m = LogicalDevices{} }
func (m *LogicalDevices) String() string { return proto.CompactTextString(m) }
func (*LogicalDevices) ProtoMessage()    {}
func (*LogicalDevices) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_a1c7cc8d1ed12585, []int{4}
}
func (m *LogicalDevices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalDevices.Unmarshal(m, b)
}
func (m *LogicalDevices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalDevices.Marshal(b, m, deterministic)
}
func (dst *LogicalDevices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalDevices.Merge(dst, src)
}
func (m *LogicalDevices) XXX_Size() int {
	return xxx_messageInfo_LogicalDevices.Size(m)
}
func (m *LogicalDevices) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalDevices.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalDevices proto.InternalMessageInfo

func (m *LogicalDevices) GetItems() []*LogicalDevice {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*LogicalPortId)(nil), "voltha.LogicalPortId")
	proto.RegisterType((*LogicalPort)(nil), "voltha.LogicalPort")
	proto.RegisterType((*LogicalPorts)(nil), "voltha.LogicalPorts")
	proto.RegisterType((*LogicalDevice)(nil), "voltha.LogicalDevice")
	proto.RegisterType((*LogicalDevices)(nil), "voltha.LogicalDevices")
}

func init() {
	proto.RegisterFile("voltha_protos/logical_device.proto", fileDescriptor_logical_device_a1c7cc8d1ed12585)
}

var fileDescriptor_logical_device_a1c7cc8d1ed12585 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x6a, 0xdb, 0x4c,
	0x14, 0xfe, 0x65, 0x5b, 0x8e, 0x7d, 0xe4, 0xf8, 0x87, 0x29, 0x21, 0x22, 0x6d, 0x89, 0x11, 0x5d,
	0x38, 0x94, 0x4a, 0x89, 0x43, 0xa1, 0x5d, 0x14, 0x8a, 0x09, 0x29, 0x86, 0x52, 0xca, 0x2c, 0xbb,
	0x11, 0x13, 0x8d, 0x24, 0x0f, 0xc8, 0x3a, 0x42, 0x33, 0x4e, 0xb6, 0x6d, 0x37, 0x7d, 0x9c, 0x3e,
	0x49, 0x5f, 0xa2, 0x0f, 0xd1, 0x75, 0x99, 0x19, 0x29, 0xb5, 0xe2, 0x76, 0x25, 0xf4, 0x5d, 0xce,
	0xe5, 0x3b, 0x0c, 0x04, 0xb7, 0x58, 0xa8, 0x35, 0x8b, 0xab, 0x1a, 0x15, 0xca, 0xa8, 0xc0, 0x5c,
	0x24, 0xac, 0x88, 0x79, 0x7a, 0x2b, 0x92, 0x34, 0x34, 0x28, 0x19, 0x5a, 0xcd, 0xc9, 0x93, 0x1c,
	0x31, 0x2f, 0xd2, 0x88, 0x55, 0x22, 0x62, 0x65, 0x89, 0x8a, 0x29, 0x81, 0xa5, 0xb4, 0xaa, 0x13,
	0xbf, 0x5b, 0x69, 0x93, 0x2a, 0xd6, 0x30, 0xa7, 0x5d, 0x06, 0xab, 0xb4, 0xcc, 0x0a, 0xbc, 0x8b,
	0x2f, 0x2e, 0xad, 0x20, 0x78, 0x05, 0x87, 0xef, 0x6d, 0xe3, 0x8f, 0x58, 0xab, 0x15, 0x27, 0x53,
	0xe8, 0x09, 0xee, 0x3b, 0x33, 0x67, 0x3e, 0xa6, 0x3d, 0xc1, 0xc9, 0x31, 0x1c, 0x54, 0x58, 0xab,
	0x58, 0x70, 0xbf, 0x67, 0xc0, 0x61, 0x65, 0x84, 0xc1, 0x77, 0x07, 0xbc, 0x1d, 0xeb, 0x9e, 0xf1,
	0x1c, 0x46, 0x98, 0x55, 0xb1, 0x56, 0x1b, 0xa7, 0xb7, 0x38, 0x0a, 0x77, 0xfb, 0xb7, 0x24, 0x3d,
	0xc0, 0xac, 0x32, 0x15, 0x1e, 0xc3, 0xd8, 0x2e, 0xaf, 0x9b, 0xf5, 0x4d, 0xa1, 0x91, 0x05, 0x56,
	0x9c, 0x3c, 0x83, 0x69, 0x43, 0x9a, 0x71, 0x4a, 0xf4, 0x07, 0x33, 0x67, 0x7e, 0x48, 0x27, 0x16,
	0xd5, 0x05, 0x3e, 0xa0, 0x2e, 0x51, 0x23, 0x2a, 0xdb, 0xd5, 0x9d, 0x39, 0xf3, 0x11, 0x1d, 0x69,
	0x40, 0xd3, 0xc1, 0x6b, 0x98, 0xec, 0x0c, 0x2c, 0xc9, 0x19, 0xb8, 0x42, 0xa5, 0x1b, 0xe9, 0x3b,
	0xb3, 0xfe, 0xdc, 0x5b, 0x3c, 0x0a, 0x6d, 0x58, 0xe1, 0x8e, 0x88, 0x5a, 0x45, 0xf0, 0xad, 0x7f,
	0x9f, 0xd3, 0x95, 0xe9, 0xb7, 0xb7, 0xee, 0x29, 0x78, 0x9c, 0x29, 0x56, 0x31, 0xb5, 0x6e, 0xb3,
	0x1a, 0x50, 0x68, 0xa1, 0x15, 0x27, 0x67, 0x30, 0xe0, 0xa9, 0x4c, 0xcc, 0x62, 0x7f, 0xcb, 0x42,
	0x93, 0xd4, 0x48, 0xc8, 0x0a, 0xfe, 0x97, 0x77, 0x42, 0x25, 0xeb, 0x38, 0x4b, 0x99, 0xda, 0xd6,
	0xa9, 0x34, 0xcb, 0x7a, 0x8b, 0xd9, 0x9e, 0xeb, 0x81, 0x8e, 0x4e, 0x2d, 0x70, 0xdd, 0xfc, 0xeb,
	0xd8, 0x4c, 0x20, 0x7f, 0x82, 0x75, 0xcd, 0xc8, 0x13, 0x8d, 0x5e, 0xb5, 0xe1, 0xbe, 0x04, 0x57,
	0x27, 0x26, 0xfd, 0xcf, 0xff, 0x8e, 0x62, 0x39, 0xfe, 0xf9, 0xeb, 0xc7, 0xd3, 0x81, 0x5e, 0x9b,
	0x5a, 0x35, 0x39, 0x07, 0x57, 0xcf, 0x22, 0xfd, 0x2f, 0x8e, 0x19, 0x8f, 0x74, 0xc6, 0xbb, 0xd6,
	0xd4, 0xd2, 0xd5, 0xae, 0xff, 0xa8, 0x15, 0x92, 0xb7, 0xe0, 0x19, 0x3a, 0xaf, 0x71, 0x5b, 0x49,
	0xff, 0xab, 0xf5, 0x1d, 0xef, 0xf9, 0xde, 0x19, 0xbe, 0x35, 0x43, 0x76, 0x0f, 0x05, 0x6f, 0x60,
	0xda, 0x39, 0x84, 0x24, 0xcf, 0xbb, 0x67, 0x3c, 0x7a, 0x30, 0xbb, 0x95, 0x35, 0x87, 0x5c, 0x5e,
	0x7c, 0x8a, 0x72, 0xa1, 0xd6, 0xdb, 0x9b, 0x30, 0xc1, 0x8d, 0x79, 0x0f, 0x09, 0xd6, 0x3c, 0xb2,
	0x96, 0x17, 0xcd, 0x33, 0x69, 0x3e, 0x16, 0xbc, 0x19, 0x9a, 0xdf, 0xcb, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0x5b, 0xa2, 0xe7, 0xb0, 0x03, 0x00, 0x00,
}
