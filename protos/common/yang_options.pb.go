// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/yang_options.proto

package common // import "github.com/opencord/voltha-protos/protos/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageParserOption int32

const (
	// Move any enclosing child enum/message definition to the same level
	// as the parent (this message) in the yang generated file
	MessageParserOption_MOVE_TO_PARENT_LEVEL MessageParserOption = 0
	// Create both a grouping and a container for this message.  The container
	// name will be the message name.  The grouping name will be the message
	// name prefixed with "grouping_"
	MessageParserOption_CREATE_BOTH_GROUPING_AND_CONTAINER MessageParserOption = 1
)

var MessageParserOption_name = map[int32]string{
	0: "MOVE_TO_PARENT_LEVEL",
	1: "CREATE_BOTH_GROUPING_AND_CONTAINER",
}
var MessageParserOption_value = map[string]int32{
	"MOVE_TO_PARENT_LEVEL":               0,
	"CREATE_BOTH_GROUPING_AND_CONTAINER": 1,
}

func (x MessageParserOption) String() string {
	return proto.EnumName(MessageParserOption_name, int32(x))
}
func (MessageParserOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_yang_options_985f6d8f51c2d464, []int{0}
}

type InlineNode struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InlineNode) Reset()         { *m = InlineNode{} }
func (m *InlineNode) String() string { return proto.CompactTextString(m) }
func (*InlineNode) ProtoMessage()    {}
func (*InlineNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_yang_options_985f6d8f51c2d464, []int{0}
}
func (m *InlineNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InlineNode.Unmarshal(m, b)
}
func (m *InlineNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InlineNode.Marshal(b, m, deterministic)
}
func (dst *InlineNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InlineNode.Merge(dst, src)
}
func (m *InlineNode) XXX_Size() int {
	return xxx_messageInfo_InlineNode.Size(m)
}
func (m *InlineNode) XXX_DiscardUnknown() {
	xxx_messageInfo_InlineNode.DiscardUnknown(m)
}

var xxx_messageInfo_InlineNode proto.InternalMessageInfo

func (m *InlineNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InlineNode) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type RpcReturnDef struct {
	// The gRPC methods return message types.  NETCONF expects an actual
	// attribute as defined in the YANG schema.  The xnl_tag will be used
	// as the top most tag when translating a gRPC response into an xml
	// response
	XmlTag string `protobuf:"bytes,1,opt,name=xml_tag,json=xmlTag,proto3" json:"xml_tag,omitempty"`
	// When the gRPC response is a list of items, we need to differentiate
	// between a YANG schema attribute whose name is "items" and when "items"
	// is used only to indicate a list of items is being returned.  The default
	// behavior assumes a list is returned when "items" is present in
	// the response.  This option will therefore be used when the attribute
	// name in the YANG schema is 'items'
	ListItemsName        string   `protobuf:"bytes,2,opt,name=list_items_name,json=listItemsName,proto3" json:"list_items_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcReturnDef) Reset()         { *m = RpcReturnDef{} }
func (m *RpcReturnDef) String() string { return proto.CompactTextString(m) }
func (*RpcReturnDef) ProtoMessage()    {}
func (*RpcReturnDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_yang_options_985f6d8f51c2d464, []int{1}
}
func (m *RpcReturnDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcReturnDef.Unmarshal(m, b)
}
func (m *RpcReturnDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcReturnDef.Marshal(b, m, deterministic)
}
func (dst *RpcReturnDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcReturnDef.Merge(dst, src)
}
func (m *RpcReturnDef) XXX_Size() int {
	return xxx_messageInfo_RpcReturnDef.Size(m)
}
func (m *RpcReturnDef) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcReturnDef.DiscardUnknown(m)
}

var xxx_messageInfo_RpcReturnDef proto.InternalMessageInfo

func (m *RpcReturnDef) GetXmlTag() string {
	if m != nil {
		return m.XmlTag
	}
	return ""
}

func (m *RpcReturnDef) GetListItemsName() string {
	if m != nil {
		return m.ListItemsName
	}
	return ""
}

var E_YangChildRule = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*MessageParserOption)(nil),
	Field:         7761774,
	Name:          "voltha.yang_child_rule",
	Tag:           "varint,7761774,opt,name=yang_child_rule,json=yangChildRule,enum=voltha.MessageParserOption",
	Filename:      "voltha_protos/yang_options.proto",
}

var E_YangMessageRule = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*MessageParserOption)(nil),
	Field:         7761775,
	Name:          "voltha.yang_message_rule",
	Tag:           "varint,7761775,opt,name=yang_message_rule,json=yangMessageRule,enum=voltha.MessageParserOption",
	Filename:      "voltha_protos/yang_options.proto",
}

var E_YangInlineNode = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*InlineNode)(nil),
	Field:         7761776,
	Name:          "voltha.yang_inline_node",
	Tag:           "bytes,7761776,opt,name=yang_inline_node,json=yangInlineNode",
	Filename:      "voltha_protos/yang_options.proto",
}

var E_YangXmlTag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*RpcReturnDef)(nil),
	Field:         7761777,
	Name:          "voltha.yang_xml_tag",
	Tag:           "bytes,7761777,opt,name=yang_xml_tag,json=yangXmlTag",
	Filename:      "voltha_protos/yang_options.proto",
}

func init() {
	proto.RegisterType((*InlineNode)(nil), "voltha.InlineNode")
	proto.RegisterType((*RpcReturnDef)(nil), "voltha.RpcReturnDef")
	proto.RegisterEnum("voltha.MessageParserOption", MessageParserOption_name, MessageParserOption_value)
	proto.RegisterExtension(E_YangChildRule)
	proto.RegisterExtension(E_YangMessageRule)
	proto.RegisterExtension(E_YangInlineNode)
	proto.RegisterExtension(E_YangXmlTag)
}

func init() {
	proto.RegisterFile("voltha_protos/yang_options.proto", fileDescriptor_yang_options_985f6d8f51c2d464)
}

var fileDescriptor_yang_options_985f6d8f51c2d464 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x69, 0x41, 0x45, 0x98, 0xad, 0x2b, 0x66, 0x12, 0x15, 0x08, 0xa8, 0xfa, 0x30, 0x4d,
	0x48, 0x24, 0x30, 0xde, 0xfa, 0xd6, 0x75, 0x61, 0x54, 0xda, 0x92, 0xca, 0x0a, 0xe3, 0xf2, 0x80,
	0x95, 0x26, 0x67, 0x89, 0x85, 0x2f, 0x51, 0xec, 0xa0, 0xed, 0xa7, 0xf2, 0xc2, 0x4f, 0xe0, 0xf2,
	0x0f, 0x50, 0xec, 0x84, 0x21, 0x6d, 0x0f, 0x7d, 0x4a, 0x7c, 0x72, 0xfc, 0x7d, 0x39, 0x17, 0x34,
	0xf9, 0xa6, 0xb8, 0x29, 0x12, 0x5a, 0x56, 0xca, 0x28, 0xed, 0x5f, 0x26, 0x32, 0xa7, 0xaa, 0x34,
	0x4c, 0x49, 0xed, 0xd9, 0x18, 0x1e, 0xb8, 0x8c, 0xc7, 0x93, 0x5c, 0xa9, 0x9c, 0x83, 0x6f, 0xa3,
	0xeb, 0xfa, 0xdc, 0xcf, 0x40, 0xa7, 0x15, 0x2b, 0x8d, 0xaa, 0x5c, 0xe6, 0xf4, 0x15, 0x42, 0x4b,
	0xc9, 0x99, 0x84, 0x50, 0x65, 0x80, 0x87, 0xa8, 0xcf, 0xb2, 0x71, 0x6f, 0xd2, 0xdb, 0xbf, 0x47,
	0xfa, 0x2c, 0xc3, 0x18, 0xdd, 0x31, 0x97, 0x25, 0x8c, 0xfb, 0x36, 0x62, 0xdf, 0xa7, 0x11, 0xda,
	0x22, 0x65, 0x4a, 0xc0, 0xd4, 0x95, 0x3c, 0x82, 0x73, 0xfc, 0x08, 0xdd, 0xbd, 0x10, 0x9c, 0x9a,
	0x24, 0x6f, 0x2f, 0x0e, 0x2e, 0x04, 0x8f, 0x93, 0x1c, 0xef, 0xa1, 0x1d, 0xce, 0xb4, 0xa1, 0xcc,
	0x80, 0xd0, 0x54, 0x26, 0xa2, 0xe3, 0x6c, 0x37, 0xe1, 0x65, 0x13, 0x0d, 0x13, 0x01, 0x2f, 0x3e,
	0xa0, 0x87, 0xa7, 0xa0, 0x75, 0x92, 0xc3, 0x2a, 0xa9, 0x34, 0x54, 0x91, 0x2d, 0x05, 0x8f, 0xd1,
	0xee, 0x69, 0x74, 0x16, 0xd0, 0x38, 0xa2, 0xab, 0x39, 0x09, 0xc2, 0x98, 0x9e, 0x04, 0x67, 0xc1,
	0xc9, 0xe8, 0x16, 0xde, 0x43, 0xd3, 0x05, 0x09, 0xe6, 0x71, 0x40, 0x0f, 0xa3, 0xf8, 0x1d, 0x3d,
	0x26, 0xd1, 0xfb, 0xd5, 0x32, 0x3c, 0xa6, 0xf3, 0xf0, 0x88, 0x2e, 0xa2, 0x30, 0x9e, 0x2f, 0xc3,
	0x80, 0x8c, 0x7a, 0xb3, 0x1c, 0xed, 0xd8, 0xde, 0xa4, 0x05, 0xe3, 0x19, 0xad, 0x6a, 0x0e, 0xf8,
	0xb9, 0xe7, 0x3a, 0xe2, 0x75, 0x1d, 0xf1, 0x5a, 0xb5, 0x93, 0xea, 0xf1, 0xcf, 0x1f, 0xdf, 0x6f,
	0x4f, 0x7a, 0xfb, 0xc3, 0x83, 0x27, 0x9e, 0xeb, 0xa1, 0x77, 0xc3, 0xbf, 0x91, 0xed, 0x86, 0xbb,
	0x68, 0xb0, 0xa4, 0xe6, 0x30, 0xfb, 0x8a, 0x1e, 0x58, 0x91, 0x70, 0xa9, 0x1b, 0xaa, 0x7e, 0x6d,
	0xa4, 0xb2, 0x25, 0xb4, 0x1f, 0xac, 0xec, 0x0b, 0x1a, 0x59, 0x19, 0xb3, 0x63, 0xa3, 0xb2, 0x99,
	0xdb, 0xd3, 0x6b, 0xae, 0xb7, 0x0c, 0x78, 0xd6, 0x99, 0x7e, 0x3b, 0xd3, 0xfd, 0x03, 0xdc, 0x99,
	0xae, 0x66, 0x4e, 0x86, 0x0d, 0xed, 0xea, 0x3c, 0xfb, 0x84, 0xb6, 0x2c, 0xbf, 0x1d, 0x2a, 0x7e,
	0x76, 0x43, 0x1d, 0xa6, 0x50, 0xff, 0xe0, 0x7f, 0x3a, 0xf8, 0x6e, 0x07, 0xff, 0x7f, 0x3d, 0x08,
	0x6a, 0x60, 0x1f, 0xed, 0x46, 0x1c, 0xbe, 0xfe, 0xec, 0xe7, 0xcc, 0x14, 0xf5, 0xda, 0x4b, 0x95,
	0xf0, 0x55, 0x09, 0x32, 0x55, 0x55, 0xe6, 0xbb, 0x6b, 0x2f, 0xdb, 0x75, 0x6e, 0x1f, 0xa9, 0x12,
	0x42, 0xc9, 0xf5, 0xc0, 0x1e, 0xdf, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xae, 0xa5, 0x75,
	0xf4, 0x02, 0x00, 0x00,
}
