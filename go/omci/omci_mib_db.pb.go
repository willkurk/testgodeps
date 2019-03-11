// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/omci_mib_db.proto

package omci // import "github.com/willkurk/testgodeps/go/omci"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/willkurk/testgodeps/go/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OpenOmciEventType_OpenOmciEventType int32

const (
	OpenOmciEventType_state_change OpenOmciEventType_OpenOmciEventType = 0
)

var OpenOmciEventType_OpenOmciEventType_name = map[int32]string{
	0: "state_change",
}
var OpenOmciEventType_OpenOmciEventType_value = map[string]int32{
	"state_change": 0,
}

func (x OpenOmciEventType_OpenOmciEventType) String() string {
	return proto.EnumName(OpenOmciEventType_OpenOmciEventType_name, int32(x))
}
func (OpenOmciEventType_OpenOmciEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{6, 0}
}

type MibAttributeData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MibAttributeData) Reset()         { *m = MibAttributeData{} }
func (m *MibAttributeData) String() string { return proto.CompactTextString(m) }
func (*MibAttributeData) ProtoMessage()    {}
func (*MibAttributeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{0}
}
func (m *MibAttributeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibAttributeData.Unmarshal(m, b)
}
func (m *MibAttributeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibAttributeData.Marshal(b, m, deterministic)
}
func (dst *MibAttributeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibAttributeData.Merge(dst, src)
}
func (m *MibAttributeData) XXX_Size() int {
	return xxx_messageInfo_MibAttributeData.Size(m)
}
func (m *MibAttributeData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibAttributeData.DiscardUnknown(m)
}

var xxx_messageInfo_MibAttributeData proto.InternalMessageInfo

func (m *MibAttributeData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MibAttributeData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MibInstanceData struct {
	InstanceId           uint32              `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Created              string              `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Modified             string              `protobuf:"bytes,3,opt,name=modified,proto3" json:"modified,omitempty"`
	Attributes           []*MibAttributeData `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MibInstanceData) Reset()         { *m = MibInstanceData{} }
func (m *MibInstanceData) String() string { return proto.CompactTextString(m) }
func (*MibInstanceData) ProtoMessage()    {}
func (*MibInstanceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{1}
}
func (m *MibInstanceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibInstanceData.Unmarshal(m, b)
}
func (m *MibInstanceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibInstanceData.Marshal(b, m, deterministic)
}
func (dst *MibInstanceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibInstanceData.Merge(dst, src)
}
func (m *MibInstanceData) XXX_Size() int {
	return xxx_messageInfo_MibInstanceData.Size(m)
}
func (m *MibInstanceData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibInstanceData.DiscardUnknown(m)
}

var xxx_messageInfo_MibInstanceData proto.InternalMessageInfo

func (m *MibInstanceData) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *MibInstanceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *MibInstanceData) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

func (m *MibInstanceData) GetAttributes() []*MibAttributeData {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type MibClassData struct {
	ClassId              uint32             `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Instances            []*MibInstanceData `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MibClassData) Reset()         { *m = MibClassData{} }
func (m *MibClassData) String() string { return proto.CompactTextString(m) }
func (*MibClassData) ProtoMessage()    {}
func (*MibClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{2}
}
func (m *MibClassData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibClassData.Unmarshal(m, b)
}
func (m *MibClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibClassData.Marshal(b, m, deterministic)
}
func (dst *MibClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibClassData.Merge(dst, src)
}
func (m *MibClassData) XXX_Size() int {
	return xxx_messageInfo_MibClassData.Size(m)
}
func (m *MibClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibClassData.DiscardUnknown(m)
}

var xxx_messageInfo_MibClassData proto.InternalMessageInfo

func (m *MibClassData) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *MibClassData) GetInstances() []*MibInstanceData {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ManagedEntity struct {
	ClassId              uint32   `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagedEntity) Reset()         { *m = ManagedEntity{} }
func (m *ManagedEntity) String() string { return proto.CompactTextString(m) }
func (*ManagedEntity) ProtoMessage()    {}
func (*ManagedEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{3}
}
func (m *ManagedEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedEntity.Unmarshal(m, b)
}
func (m *ManagedEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedEntity.Marshal(b, m, deterministic)
}
func (dst *ManagedEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedEntity.Merge(dst, src)
}
func (m *ManagedEntity) XXX_Size() int {
	return xxx_messageInfo_ManagedEntity.Size(m)
}
func (m *ManagedEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedEntity proto.InternalMessageInfo

func (m *ManagedEntity) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *ManagedEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MessageType struct {
	MessageType          uint32   `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageType) Reset()         { *m = MessageType{} }
func (m *MessageType) String() string { return proto.CompactTextString(m) }
func (*MessageType) ProtoMessage()    {}
func (*MessageType) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{4}
}
func (m *MessageType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageType.Unmarshal(m, b)
}
func (m *MessageType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageType.Marshal(b, m, deterministic)
}
func (dst *MessageType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageType.Merge(dst, src)
}
func (m *MessageType) XXX_Size() int {
	return xxx_messageInfo_MessageType.Size(m)
}
func (m *MessageType) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageType.DiscardUnknown(m)
}

var xxx_messageInfo_MessageType proto.InternalMessageInfo

func (m *MessageType) GetMessageType() uint32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

type MibDeviceData struct {
	DeviceId             string           `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Created              string           `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	LastSyncTime         string           `protobuf:"bytes,3,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	MibDataSync          uint32           `protobuf:"varint,4,opt,name=mib_data_sync,json=mibDataSync,proto3" json:"mib_data_sync,omitempty"`
	Version              uint32           `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Classes              []*MibClassData  `protobuf:"bytes,6,rep,name=classes,proto3" json:"classes,omitempty"`
	ManagedEntities      []*ManagedEntity `protobuf:"bytes,7,rep,name=managed_entities,json=managedEntities,proto3" json:"managed_entities,omitempty"`
	MessageTypes         []*MessageType   `protobuf:"bytes,8,rep,name=message_types,json=messageTypes,proto3" json:"message_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MibDeviceData) Reset()         { *m = MibDeviceData{} }
func (m *MibDeviceData) String() string { return proto.CompactTextString(m) }
func (*MibDeviceData) ProtoMessage()    {}
func (*MibDeviceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{5}
}
func (m *MibDeviceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibDeviceData.Unmarshal(m, b)
}
func (m *MibDeviceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibDeviceData.Marshal(b, m, deterministic)
}
func (dst *MibDeviceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibDeviceData.Merge(dst, src)
}
func (m *MibDeviceData) XXX_Size() int {
	return xxx_messageInfo_MibDeviceData.Size(m)
}
func (m *MibDeviceData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibDeviceData.DiscardUnknown(m)
}

var xxx_messageInfo_MibDeviceData proto.InternalMessageInfo

func (m *MibDeviceData) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *MibDeviceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *MibDeviceData) GetLastSyncTime() string {
	if m != nil {
		return m.LastSyncTime
	}
	return ""
}

func (m *MibDeviceData) GetMibDataSync() uint32 {
	if m != nil {
		return m.MibDataSync
	}
	return 0
}

func (m *MibDeviceData) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MibDeviceData) GetClasses() []*MibClassData {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *MibDeviceData) GetManagedEntities() []*ManagedEntity {
	if m != nil {
		return m.ManagedEntities
	}
	return nil
}

func (m *MibDeviceData) GetMessageTypes() []*MessageType {
	if m != nil {
		return m.MessageTypes
	}
	return nil
}

type OpenOmciEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenOmciEventType) Reset()         { *m = OpenOmciEventType{} }
func (m *OpenOmciEventType) String() string { return proto.CompactTextString(m) }
func (*OpenOmciEventType) ProtoMessage()    {}
func (*OpenOmciEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{6}
}
func (m *OpenOmciEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenOmciEventType.Unmarshal(m, b)
}
func (m *OpenOmciEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenOmciEventType.Marshal(b, m, deterministic)
}
func (dst *OpenOmciEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenOmciEventType.Merge(dst, src)
}
func (m *OpenOmciEventType) XXX_Size() int {
	return xxx_messageInfo_OpenOmciEventType.Size(m)
}
func (m *OpenOmciEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenOmciEventType.DiscardUnknown(m)
}

var xxx_messageInfo_OpenOmciEventType proto.InternalMessageInfo

type OpenOmciEvent struct {
	Type                 OpenOmciEventType_OpenOmciEventType `protobuf:"varint,1,opt,name=type,proto3,enum=omci.OpenOmciEventType_OpenOmciEventType" json:"type,omitempty"`
	Data                 string                              `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *OpenOmciEvent) Reset()         { *m = OpenOmciEvent{} }
func (m *OpenOmciEvent) String() string { return proto.CompactTextString(m) }
func (*OpenOmciEvent) ProtoMessage()    {}
func (*OpenOmciEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_omci_mib_db_b2d797f82b5a5919, []int{7}
}
func (m *OpenOmciEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenOmciEvent.Unmarshal(m, b)
}
func (m *OpenOmciEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenOmciEvent.Marshal(b, m, deterministic)
}
func (dst *OpenOmciEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenOmciEvent.Merge(dst, src)
}
func (m *OpenOmciEvent) XXX_Size() int {
	return xxx_messageInfo_OpenOmciEvent.Size(m)
}
func (m *OpenOmciEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenOmciEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OpenOmciEvent proto.InternalMessageInfo

func (m *OpenOmciEvent) GetType() OpenOmciEventType_OpenOmciEventType {
	if m != nil {
		return m.Type
	}
	return OpenOmciEventType_state_change
}

func (m *OpenOmciEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*MibAttributeData)(nil), "omci.MibAttributeData")
	proto.RegisterType((*MibInstanceData)(nil), "omci.MibInstanceData")
	proto.RegisterType((*MibClassData)(nil), "omci.MibClassData")
	proto.RegisterType((*ManagedEntity)(nil), "omci.ManagedEntity")
	proto.RegisterType((*MessageType)(nil), "omci.MessageType")
	proto.RegisterType((*MibDeviceData)(nil), "omci.MibDeviceData")
	proto.RegisterType((*OpenOmciEventType)(nil), "omci.OpenOmciEventType")
	proto.RegisterType((*OpenOmciEvent)(nil), "omci.OpenOmciEvent")
	proto.RegisterEnum("omci.OpenOmciEventType_OpenOmciEventType", OpenOmciEventType_OpenOmciEventType_name, OpenOmciEventType_OpenOmciEventType_value)
}

func init() {
	proto.RegisterFile("voltha_protos/omci_mib_db.proto", fileDescriptor_omci_mib_db_b2d797f82b5a5919)
}

var fileDescriptor_omci_mib_db_b2d797f82b5a5919 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0xfd, 0xec, 0x28, 0x89, 0x33, 0xb2, 0x12, 0x67, 0xbf, 0xb6, 0x6c, 0x03, 0xa1, 0x41, 0xb4,
	0xc1, 0xbd, 0xb1, 0x21, 0x85, 0x16, 0x0a, 0x29, 0xd4, 0x49, 0x28, 0x81, 0x8a, 0x80, 0x9a, 0xab,
	0xde, 0x88, 0x95, 0x34, 0x95, 0x97, 0x68, 0x25, 0xe3, 0x5d, 0xbb, 0xe8, 0xbd, 0xfa, 0x1a, 0x79,
	0x89, 0x5e, 0xe5, 0x09, 0x72, 0x5d, 0x76, 0xf5, 0x63, 0xb9, 0x2e, 0xa5, 0x77, 0x3e, 0x67, 0x66,
	0xce, 0xcc, 0x9e, 0x63, 0x04, 0x2f, 0x96, 0x79, 0xaa, 0xa6, 0x2c, 0x98, 0xcd, 0x73, 0x95, 0xcb,
	0x71, 0x2e, 0x22, 0x1e, 0x08, 0x1e, 0x06, 0x71, 0x38, 0x32, 0x14, 0xb1, 0x34, 0x75, 0x44, 0xd7,
	0xdb, 0x04, 0x2a, 0x56, 0xd6, 0xdd, 0x0b, 0x18, 0x78, 0x3c, 0xfc, 0xa8, 0xd4, 0x9c, 0x87, 0x0b,
	0x85, 0x97, 0x4c, 0x31, 0xf2, 0x1c, 0xac, 0x8c, 0x09, 0xa4, 0x9d, 0x93, 0xce, 0x70, 0x6f, 0xb2,
	0xfd, 0xf0, 0x78, 0x7f, 0xdc, 0xf1, 0x0d, 0x45, 0x9e, 0xc0, 0xf6, 0x92, 0xa5, 0x0b, 0xa4, 0x5d,
	0x5d, 0xf3, 0x4b, 0xe0, 0xfe, 0xe8, 0xc0, 0x81, 0xc7, 0xc3, 0xeb, 0x4c, 0x2a, 0x96, 0x45, 0xa5,
	0xc8, 0x29, 0xd8, 0xbc, 0xc2, 0x01, 0x8f, 0x8d, 0x96, 0x53, 0x6b, 0x41, 0x5d, 0xb9, 0x8e, 0x09,
	0x85, 0xdd, 0x68, 0x8e, 0x4c, 0x61, 0x5c, 0x69, 0xd6, 0x90, 0x1c, 0x41, 0x4f, 0xe4, 0x31, 0xff,
	0xc6, 0x31, 0xa6, 0x5b, 0xa6, 0xd4, 0x60, 0x72, 0x01, 0xc0, 0xea, 0x9b, 0x25, 0xb5, 0x4e, 0xb6,
	0x86, 0xf6, 0xd9, 0xb3, 0x91, 0x7e, 0xeb, 0xe8, 0xf7, 0xe7, 0x4c, 0xec, 0x9f, 0x8f, 0xf7, 0xc7,
	0x3b, 0xe5, 0x9b, 0xfc, 0xd6, 0x98, 0x5b, 0x40, 0xdf, 0xe3, 0xe1, 0x45, 0xca, 0xa4, 0x34, 0x27,
	0x9f, 0x40, 0x2f, 0xd2, 0x60, 0xe3, 0xde, 0x5d, 0x43, 0x5f, 0xc7, 0xe4, 0x13, 0xec, 0xd5, 0xa7,
	0x4b, 0xda, 0x35, 0x5b, 0x9f, 0x36, 0x5b, 0xdb, 0xcf, 0x9f, 0x10, 0xbd, 0xd4, 0x59, 0xf3, 0xc0,
	0x5f, 0xcd, 0xba, 0x9f, 0xc1, 0xf1, 0x58, 0xc6, 0x12, 0x8c, 0xaf, 0x32, 0xc5, 0x55, 0xf1, 0x0f,
	0xbb, 0xeb, 0x54, 0xba, 0x1b, 0xa9, 0xb8, 0xef, 0xc0, 0xf6, 0x50, 0x4a, 0x96, 0xe0, 0x6d, 0x31,
	0x43, 0x32, 0x84, 0xbe, 0x28, 0x61, 0xa0, 0x8a, 0x19, 0xae, 0xeb, 0xd9, 0x62, 0xd5, 0xe9, 0x3e,
	0x74, 0xc1, 0xf1, 0x78, 0x78, 0x89, 0x4b, 0x5e, 0xc5, 0xe6, 0xc2, 0x5e, 0x6c, 0x50, 0x7d, 0x48,
	0xb3, 0xaa, 0x57, 0xf2, 0x7f, 0x8d, 0xec, 0x25, 0xec, 0xa7, 0x4c, 0xaa, 0x40, 0x16, 0x59, 0x14,
	0x28, 0x2e, 0xb0, 0x0a, 0xae, 0xaf, 0xd9, 0x2f, 0x45, 0x16, 0xdd, 0x72, 0x81, 0xc4, 0x05, 0xc7,
	0xfc, 0x47, 0x99, 0x62, 0xa6, 0x93, 0x5a, 0xfa, 0x40, 0xdf, 0x16, 0x3c, 0xd4, 0x37, 0xe8, 0x3e,
	0xbd, 0x63, 0x89, 0x73, 0xc9, 0xf3, 0x8c, 0x6e, 0x9b, 0x6a, 0x0d, 0xc9, 0x39, 0x94, 0x96, 0xa0,
	0xa4, 0x3b, 0x26, 0x01, 0xd2, 0x24, 0xd0, 0x44, 0x39, 0x39, 0xd0, 0xf6, 0xc3, 0xca, 0x53, 0xbf,
	0x9e, 0x21, 0x1f, 0x60, 0x20, 0x4a, 0xe7, 0x03, 0xd4, 0xd6, 0x73, 0x94, 0x74, 0xd7, 0xe8, 0xfc,
	0x5f, 0xe9, 0xb4, 0x73, 0xf1, 0x0f, 0x44, 0x0b, 0x72, 0x94, 0xe4, 0x2d, 0x38, 0x6d, 0x73, 0x25,
	0xed, 0x99, 0xe1, 0xc3, 0x6a, 0x78, 0x65, 0xae, 0xdf, 0x6f, 0x39, 0x2d, 0xdd, 0xf7, 0x70, 0x78,
	0x33, 0xc3, 0xec, 0x46, 0x44, 0xfc, 0x6a, 0x89, 0x99, 0x32, 0xfe, 0xbf, 0xfa, 0x03, 0x49, 0x06,
	0xd0, 0x97, 0x8a, 0x29, 0x0c, 0xa2, 0x29, 0xcb, 0x12, 0x1c, 0xfc, 0xe7, 0x86, 0xe0, 0xac, 0xb5,
	0x91, 0x73, 0xb0, 0x9a, 0x64, 0xf7, 0xcf, 0x5e, 0x97, 0xbb, 0x37, 0x94, 0x36, 0x19, 0xdf, 0x8c,
	0x11, 0x02, 0x96, 0x36, 0xbf, 0x4a, 0xcf, 0xfc, 0x9e, 0x0c, 0xbf, 0x9e, 0x26, 0x5c, 0x4d, 0x17,
	0xe1, 0x28, 0xca, 0xc5, 0xf8, 0x3b, 0x4f, 0xd3, 0xbb, 0xc5, 0xfc, 0x6e, 0xac, 0x50, 0xaa, 0x24,
	0x8f, 0x71, 0x26, 0xc7, 0x49, 0x6e, 0xbe, 0x2f, 0xe1, 0x8e, 0xf9, 0x72, 0xbc, 0xf9, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x0e, 0xde, 0x9c, 0x7c, 0x04, 0x00, 0x00,
}
