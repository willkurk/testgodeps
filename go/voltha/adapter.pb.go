// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/adapter.proto

package voltha // import "github.com/willkurk/testgodeps/go/voltha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import common "github.com/willkurk/testgodeps/go/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AdapterConfig struct {
	// Common adapter config attributes here
	LogLevel common.LogLevel_LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=voltha.LogLevel_LogLevel" json:"log_level,omitempty"`
	// Custom (vendor-specific) configuration attributes
	AdditionalConfig     *any.Any `protobuf:"bytes,64,opt,name=additional_config,json=additionalConfig,proto3" json:"additional_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdapterConfig) Reset()         { *m = AdapterConfig{} }
func (m *AdapterConfig) String() string { return proto.CompactTextString(m) }
func (*AdapterConfig) ProtoMessage()    {}
func (*AdapterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_adapter_783d2aead7374128, []int{0}
}
func (m *AdapterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdapterConfig.Unmarshal(m, b)
}
func (m *AdapterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdapterConfig.Marshal(b, m, deterministic)
}
func (dst *AdapterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdapterConfig.Merge(dst, src)
}
func (m *AdapterConfig) XXX_Size() int {
	return xxx_messageInfo_AdapterConfig.Size(m)
}
func (m *AdapterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdapterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdapterConfig proto.InternalMessageInfo

func (m *AdapterConfig) GetLogLevel() common.LogLevel_LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return common.LogLevel_DEBUG
}

func (m *AdapterConfig) GetAdditionalConfig() *any.Any {
	if m != nil {
		return m.AdditionalConfig
	}
	return nil
}

// Adapter (software plugin)
type Adapter struct {
	// Unique name of adapter, matching the python package name under
	// voltha/adapters.
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vendor  string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Adapter configuration
	Config *AdapterConfig `protobuf:"bytes,16,opt,name=config,proto3" json:"config,omitempty"`
	// Custom descriptors and custom configuration
	AdditionalDescription *any.Any `protobuf:"bytes,64,opt,name=additional_description,json=additionalDescription,proto3" json:"additional_description,omitempty"`
	LogicalDeviceIds      []string `protobuf:"bytes,4,rep,name=logical_device_ids,json=logicalDeviceIds,proto3" json:"logical_device_ids,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Adapter) Reset()         { *m = Adapter{} }
func (m *Adapter) String() string { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()    {}
func (*Adapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_adapter_783d2aead7374128, []int{1}
}
func (m *Adapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adapter.Unmarshal(m, b)
}
func (m *Adapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adapter.Marshal(b, m, deterministic)
}
func (dst *Adapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adapter.Merge(dst, src)
}
func (m *Adapter) XXX_Size() int {
	return xxx_messageInfo_Adapter.Size(m)
}
func (m *Adapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Adapter.DiscardUnknown(m)
}

var xxx_messageInfo_Adapter proto.InternalMessageInfo

func (m *Adapter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Adapter) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Adapter) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Adapter) GetConfig() *AdapterConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Adapter) GetAdditionalDescription() *any.Any {
	if m != nil {
		return m.AdditionalDescription
	}
	return nil
}

func (m *Adapter) GetLogicalDeviceIds() []string {
	if m != nil {
		return m.LogicalDeviceIds
	}
	return nil
}

type Adapters struct {
	Items                []*Adapter `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Adapters) Reset()         { *m = Adapters{} }
func (m *Adapters) String() string { return proto.CompactTextString(m) }
func (*Adapters) ProtoMessage()    {}
func (*Adapters) Descriptor() ([]byte, []int) {
	return fileDescriptor_adapter_783d2aead7374128, []int{2}
}
func (m *Adapters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adapters.Unmarshal(m, b)
}
func (m *Adapters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adapters.Marshal(b, m, deterministic)
}
func (dst *Adapters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adapters.Merge(dst, src)
}
func (m *Adapters) XXX_Size() int {
	return xxx_messageInfo_Adapters.Size(m)
}
func (m *Adapters) XXX_DiscardUnknown() {
	xxx_messageInfo_Adapters.DiscardUnknown(m)
}

var xxx_messageInfo_Adapters proto.InternalMessageInfo

func (m *Adapters) GetItems() []*Adapter {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*AdapterConfig)(nil), "voltha.AdapterConfig")
	proto.RegisterType((*Adapter)(nil), "voltha.Adapter")
	proto.RegisterType((*Adapters)(nil), "voltha.Adapters")
}

func init() {
	proto.RegisterFile("voltha_protos/adapter.proto", fileDescriptor_adapter_783d2aead7374128)
}

var fileDescriptor_adapter_783d2aead7374128 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x95, 0x50, 0x02, 0x18, 0xb5, 0xa5, 0x56, 0xa9, 0x02, 0x15, 0x6a, 0x84, 0x54, 0x29,
	0xaa, 0xda, 0x44, 0xa5, 0x52, 0xcf, 0x85, 0x72, 0xa9, 0xc4, 0x29, 0xc7, 0x5e, 0xa2, 0x10, 0x1b,
	0x63, 0xe1, 0x64, 0xa2, 0xd8, 0xa4, 0xe2, 0x15, 0x7a, 0xeb, 0x83, 0xf5, 0x3d, 0xf6, 0x09, 0xf6,
	0xbc, 0xc2, 0x76, 0x16, 0xd8, 0xc3, 0xde, 0xec, 0xf9, 0xe6, 0xf7, 0x7c, 0x63, 0xf4, 0xbe, 0x01,
	0xa1, 0xf6, 0x59, 0x5a, 0xd5, 0xa0, 0x40, 0xc6, 0x19, 0xc9, 0x2a, 0x45, 0xeb, 0x48, 0x5f, 0xb1,
	0x67, 0xe0, 0x74, 0xc2, 0x00, 0x98, 0xa0, 0xb1, 0xae, 0x6e, 0x8f, 0xbb, 0x38, 0x2b, 0x4f, 0xa6,
	0x65, 0x3a, 0xbd, 0xcd, 0xe7, 0x50, 0x14, 0x50, 0x5a, 0xe6, 0xdf, 0xb2, 0x82, 0xaa, 0xcc, 0x90,
	0xf9, 0x5f, 0x07, 0xbd, 0x5c, 0x9a, 0x51, 0x3f, 0xa1, 0xdc, 0x71, 0x86, 0xbf, 0xa3, 0x81, 0x00,
	0x96, 0x0a, 0xda, 0x50, 0xe1, 0x3b, 0x81, 0x13, 0xbe, 0x5a, 0x4c, 0x22, 0x93, 0x8f, 0x36, 0xc0,
	0x36, 0xe7, 0xfa, 0xe3, 0x21, 0xe9, 0x0b, 0x7b, 0xc2, 0x4b, 0xf4, 0x26, 0x23, 0x84, 0x2b, 0x0e,
	0x65, 0x26, 0xd2, 0x5c, 0x3f, 0xe6, 0xff, 0x08, 0x9c, 0x70, 0xb8, 0x78, 0x1b, 0x19, 0xed, 0xa8,
	0xd5, 0x8e, 0x96, 0xe5, 0x29, 0x19, 0x5d, 0xda, 0xcd, 0xe8, 0xf9, 0x3f, 0x17, 0xf5, 0xac, 0x0c,
	0x1e, 0x23, 0x97, 0x13, 0x3d, 0x7f, 0xb0, 0xea, 0xde, 0xdd, 0xff, 0x9f, 0x39, 0x89, 0xcb, 0x09,
	0x9e, 0x21, 0xaf, 0xa1, 0x25, 0x81, 0xda, 0x77, 0xaf, 0x91, 0x2d, 0xe2, 0x0f, 0xa8, 0xd7, 0xd0,
	0x5a, 0x72, 0x28, 0xfd, 0xce, 0x35, 0x6f, 0xab, 0xf8, 0x0b, 0xf2, 0xac, 0xda, 0x48, 0xab, 0x8d,
	0xdb, 0xd5, 0x6e, 0x3e, 0x21, 0xb1, 0x4d, 0x38, 0x41, 0xef, 0xae, 0x96, 0x22, 0x54, 0xe6, 0x35,
	0xaf, 0xce, 0xb7, 0xe7, 0x36, 0x6b, 0x87, 0x8e, 0x2f, 0xd1, 0xf5, 0x25, 0x89, 0x3f, 0x23, 0x2c,
	0x80, 0xf1, 0x5c, 0x3f, 0xd8, 0xf0, 0x9c, 0xa6, 0x9c, 0x48, 0xff, 0x45, 0xd0, 0x09, 0x07, 0xc9,
	0xc8, 0x92, 0xb5, 0x06, 0xbf, 0x88, 0x9c, 0x7f, 0x45, 0x7d, 0xab, 0x26, 0xf1, 0x47, 0xd4, 0xe5,
	0x8a, 0x16, 0xd2, 0x77, 0x82, 0x4e, 0x38, 0x5c, 0xbc, 0x7e, 0xe2, 0x9e, 0x18, 0xba, 0xfa, 0xf4,
	0x3b, 0x64, 0x5c, 0xed, 0x8f, 0xdb, 0x28, 0x87, 0x22, 0xfe, 0xc3, 0x85, 0x38, 0x1c, 0xeb, 0x43,
	0xac, 0xa8, 0x54, 0x0c, 0x08, 0xad, 0x64, 0xcc, 0x20, 0x36, 0xd1, 0xad, 0xa7, 0xc5, 0xbf, 0x3d,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0xdb, 0x22, 0x31, 0x7e, 0x02, 0x00, 0x00,
}
