// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub

import (
	fmt "fmt"
	proto "github.com/catper/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type E int32

const (
	E_ZERO E = 0
)

var E_name = map[int32]string{
	0: "ZERO",
}

var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}

func (x *E) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(E_value, data, "E")
	if err != nil {
		return err
	}
	*x = E(value)
	return nil
}

func (E) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

type M_Subenum int32

const (
	M_M_ZERO M_Subenum = 0
)

var M_Subenum_name = map[int32]string{
	0: "M_ZERO",
}

var M_Subenum_value = map[string]int32{
	"M_ZERO": 0,
}

func (x M_Subenum) Enum() *M_Subenum {
	p := new(M_Subenum)
	*p = x
	return p
}

func (x M_Subenum) String() string {
	return proto.EnumName(M_Subenum_name, int32(x))
}

func (x *M_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Subenum_value, data, "M_Subenum")
	if err != nil {
		return err
	}
	*x = M_Subenum(value)
	return nil
}

func (M_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

type M_Submessage_Submessage_Subenum int32

const (
	M_Submessage_M_SUBMESSAGE_ZERO M_Submessage_Submessage_Subenum = 0
)

var M_Submessage_Submessage_Subenum_name = map[int32]string{
	0: "M_SUBMESSAGE_ZERO",
}

var M_Submessage_Submessage_Subenum_value = map[string]int32{
	"M_SUBMESSAGE_ZERO": 0,
}

func (x M_Submessage_Submessage_Subenum) Enum() *M_Submessage_Submessage_Subenum {
	p := new(M_Submessage_Submessage_Subenum)
	*p = x
	return p
}

func (x M_Submessage_Submessage_Subenum) String() string {
	return proto.EnumName(M_Submessage_Submessage_Subenum_name, int32(x))
}

func (x *M_Submessage_Submessage_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Submessage_Submessage_Subenum_value, data, "M_Submessage_Submessage_Subenum")
	if err != nil {
		return err
	}
	*x = M_Submessage_Submessage_Subenum(value)
	return nil
}

func (M_Submessage_Submessage_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 1, 0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2 *M2 `protobuf:"bytes,1,opt,name=m2" json:"m2,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*M_OneofInt32
	//	*M_OneofInt64
	OneofField           isM_OneofField `protobuf_oneof:"oneof_field"`
	Grouping             *M_Grouping    `protobuf:"group,4,opt,name=Grouping,json=grouping" json:"grouping,omitempty"`
	DefaultField         *string        `protobuf:"bytes,6,opt,name=default_field,json=defaultField,def=def" json:"default_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

const Default_M_DefaultField string = "def"

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

type isM_OneofField interface {
	isM_OneofField()
}

type M_OneofInt32 struct {
	OneofInt32 int32 `protobuf:"varint,2,opt,name=oneof_int32,json=oneofInt32,oneof"`
}

type M_OneofInt64 struct {
	OneofInt64 int64 `protobuf:"varint,3,opt,name=oneof_int64,json=oneofInt64,oneof"`
}

func (*M_OneofInt32) isM_OneofField() {}

func (*M_OneofInt64) isM_OneofField() {}

func (m *M) GetOneofField() isM_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *M) GetOneofInt32() int32 {
	if x, ok := m.GetOneofField().(*M_OneofInt32); ok {
		return x.OneofInt32
	}
	return 0
}

func (m *M) GetOneofInt64() int64 {
	if x, ok := m.GetOneofField().(*M_OneofInt64); ok {
		return x.OneofInt64
	}
	return 0
}

func (m *M) GetGrouping() *M_Grouping {
	if m != nil {
		return m.Grouping
	}
	return nil
}

func (m *M) GetDefaultField() string {
	if m != nil && m.DefaultField != nil {
		return *m.DefaultField
	}
	return Default_M_DefaultField
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_OneofInt32)(nil),
		(*M_OneofInt64)(nil),
	}
}

type M_Grouping struct {
	GroupField           *string  `protobuf:"bytes,5,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M_Grouping) Reset()         { *m = M_Grouping{} }
func (m *M_Grouping) String() string { return proto.CompactTextString(m) }
func (*M_Grouping) ProtoMessage()    {}
func (*M_Grouping) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

func (m *M_Grouping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Grouping.Unmarshal(m, b)
}
func (m *M_Grouping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Grouping.Marshal(b, m, deterministic)
}
func (m *M_Grouping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Grouping.Merge(m, src)
}
func (m *M_Grouping) XXX_Size() int {
	return xxx_messageInfo_M_Grouping.Size(m)
}
func (m *M_Grouping) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Grouping.DiscardUnknown(m)
}

var xxx_messageInfo_M_Grouping proto.InternalMessageInfo

func (m *M_Grouping) GetGroupField() string {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return ""
}

type M_Submessage struct {
	// Types that are valid to be assigned to SubmessageOneofField:
	//	*M_Submessage_SubmessageOneofInt32
	//	*M_Submessage_SubmessageOneofInt64
	SubmessageOneofField isM_Submessage_SubmessageOneofField `protobuf_oneof:"submessage_oneof_field"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *M_Submessage) Reset()         { *m = M_Submessage{} }
func (m *M_Submessage) String() string { return proto.CompactTextString(m) }
func (*M_Submessage) ProtoMessage()    {}
func (*M_Submessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 1}
}

func (m *M_Submessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Submessage.Unmarshal(m, b)
}
func (m *M_Submessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Submessage.Marshal(b, m, deterministic)
}
func (m *M_Submessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Submessage.Merge(m, src)
}
func (m *M_Submessage) XXX_Size() int {
	return xxx_messageInfo_M_Submessage.Size(m)
}
func (m *M_Submessage) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Submessage.DiscardUnknown(m)
}

var xxx_messageInfo_M_Submessage proto.InternalMessageInfo

type isM_Submessage_SubmessageOneofField interface {
	isM_Submessage_SubmessageOneofField()
}

type M_Submessage_SubmessageOneofInt32 struct {
	SubmessageOneofInt32 int32 `protobuf:"varint,1,opt,name=submessage_oneof_int32,json=submessageOneofInt32,oneof"`
}

type M_Submessage_SubmessageOneofInt64 struct {
	SubmessageOneofInt64 int64 `protobuf:"varint,2,opt,name=submessage_oneof_int64,json=submessageOneofInt64,oneof"`
}

func (*M_Submessage_SubmessageOneofInt32) isM_Submessage_SubmessageOneofField() {}

func (*M_Submessage_SubmessageOneofInt64) isM_Submessage_SubmessageOneofField() {}

func (m *M_Submessage) GetSubmessageOneofField() isM_Submessage_SubmessageOneofField {
	if m != nil {
		return m.SubmessageOneofField
	}
	return nil
}

func (m *M_Submessage) GetSubmessageOneofInt32() int32 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt32); ok {
		return x.SubmessageOneofInt32
	}
	return 0
}

func (m *M_Submessage) GetSubmessageOneofInt64() int64 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt64); ok {
		return x.SubmessageOneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M_Submessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_Submessage_SubmessageOneofInt32)(nil),
		(*M_Submessage_SubmessageOneofInt64)(nil),
	}
}

var E_ExtensionField = &proto.ExtensionDesc{
	ExtendedType:  (*M2)(nil),
	ExtensionType: (*string)(nil),
	Field:         1,
	Name:          "goproto.test.import_public.sub.extension_field",
	Tag:           "bytes,1,opt,name=extension_field",
	Filename:      "import_public/sub/a.proto",
}

func init() {
	proto.RegisterEnum("goproto.test.import_public.sub.E", E_name, E_value)
	proto.RegisterEnum("goproto.test.import_public.sub.M_Subenum", M_Subenum_name, M_Subenum_value)
	proto.RegisterEnum("goproto.test.import_public.sub.M_Submessage_Submessage_Subenum", M_Submessage_Submessage_Subenum_name, M_Submessage_Submessage_Subenum_value)
	proto.RegisterType((*M)(nil), "goproto.test.import_public.sub.M")
	proto.RegisterType((*M_Grouping)(nil), "goproto.test.import_public.sub.M.Grouping")
	proto.RegisterType((*M_Submessage)(nil), "goproto.test.import_public.sub.M.Submessage")
	proto.RegisterExtension(E_ExtensionField)
}

func init() {
	proto.RegisterFile("import_public/sub/a.proto", fileDescriptor_382f7805394b5c4e)
}

var fileDescriptor_382f7805394b5c4e = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x49, 0x5a, 0xd2, 0x09, 0xe1, 0xcf, 0x8a, 0x22, 0xd3, 0x03, 0x98, 0x9c, 0xac,
	0x56, 0x5d, 0x4b, 0x26, 0xf2, 0xa1, 0x37, 0x82, 0xdc, 0x82, 0x90, 0x55, 0xc9, 0x16, 0x97, 0x5e,
	0x2c, 0x6f, 0xbc, 0x5e, 0x2c, 0xd9, 0xbb, 0x56, 0xbc, 0x2b, 0xf1, 0x08, 0xbc, 0x17, 0x2f, 0x86,
	0xbc, 0xb6, 0x93, 0x46, 0x04, 0xe8, 0x6d, 0x3d, 0xf3, 0xfd, 0xbe, 0xd1, 0x7c, 0x1e, 0x78, 0x53,
	0x54, 0xb5, 0xdc, 0xa8, 0xa4, 0xd6, 0xb4, 0x2c, 0xd6, 0x6e, 0xa3, 0xa9, 0x9b, 0x92, 0x7a, 0x23,
	0x95, 0xc4, 0x6f, 0xb9, 0x34, 0x0f, 0xa2, 0x58, 0xa3, 0xc8, 0x9e, 0x8e, 0x34, 0x9a, 0x9e, 0x1f,
	0x40, 0x69, 0x87, 0x2e, 0x7e, 0x4e, 0x00, 0x85, 0xd8, 0x83, 0x51, 0xe5, 0x59, 0xc8, 0x46, 0xce,
	0xcc, 0x5b, 0x90, 0x7f, 0xbb, 0x91, 0xd0, 0x8b, 0x46, 0x95, 0x87, 0xdf, 0xc3, 0x4c, 0x0a, 0x26,
	0xf3, 0xa4, 0x10, 0xea, 0x83, 0x67, 0x8d, 0x6c, 0xe4, 0x1c, 0x7f, 0x3e, 0x8a, 0xc0, 0x14, 0xbf,
	0xb4, 0xb5, 0x3d, 0x89, 0xbf, 0xb4, 0xc6, 0x36, 0x72, 0xc6, 0x0f, 0x25, 0xfe, 0x12, 0xdf, 0xc0,
	0x94, 0x6f, 0xa4, 0xae, 0x0b, 0xc1, 0xad, 0x89, 0x8d, 0x1c, 0xf0, 0x2e, 0xfe, 0x3b, 0x9f, 0xdc,
	0xf6, 0x44, 0xb4, 0x65, 0xb1, 0x03, 0xf3, 0x8c, 0xe5, 0xa9, 0x2e, 0x55, 0x92, 0x17, 0xac, 0xcc,
	0xac, 0x13, 0x1b, 0x39, 0xa7, 0xd7, 0xe3, 0x8c, 0xe5, 0xd1, 0xd3, 0xbe, 0x73, 0xd3, 0x36, 0xce,
	0x2f, 0x61, 0x3a, 0xf0, 0xf8, 0x1d, 0xcc, 0x8c, 0x43, 0xcf, 0x1c, 0xb7, 0x4c, 0x04, 0xa6, 0xd4,
	0x89, 0x7f, 0x21, 0x80, 0x58, 0xd3, 0x8a, 0x35, 0x4d, 0xca, 0x19, 0xf6, 0xe1, 0x75, 0xb3, 0xfd,
	0x4a, 0x1e, 0xae, 0x8f, 0xfa, 0xf5, 0x5f, 0xed, 0xfa, 0x77, 0xbb, 0x20, 0xfe, 0xc2, 0xf9, 0x4b,
	0x13, 0xdb, 0xf8, 0x30, 0xe7, 0x2f, 0x17, 0x97, 0x80, 0x77, 0xd3, 0x93, 0x58, 0x53, 0x26, 0x74,
	0x85, 0xcf, 0xe0, 0x65, 0x98, 0xc4, 0xdf, 0x56, 0x61, 0x10, 0xc7, 0x1f, 0x6f, 0x83, 0xe4, 0x3e,
	0x88, 0xee, 0x5e, 0x1c, 0xad, 0xac, 0x03, 0x43, 0xcc, 0x5e, 0x8b, 0x33, 0x78, 0x32, 0xb0, 0x00,
	0x27, 0xe1, 0x00, 0xcc, 0x87, 0xdf, 0x63, 0x54, 0x17, 0x73, 0x40, 0x01, 0x9e, 0xc2, 0xa4, 0xeb,
	0x5e, 0x7f, 0x85, 0xe7, 0xec, 0x87, 0x62, 0xa2, 0x29, 0xa4, 0xe8, 0x14, 0xf8, 0x11, 0xa7, 0x61,
	0x82, 0x38, 0x8d, 0x9e, 0x6d, 0x51, 0x93, 0xe3, 0x2a, 0xb8, 0xff, 0xc4, 0x0b, 0xf5, 0x5d, 0x53,
	0xb2, 0x96, 0x95, 0xcb, 0x65, 0x99, 0x0a, 0xee, 0x1a, 0x2b, 0xaa, 0xf3, 0xee, 0xb1, 0xbe, 0xe2,
	0x4c, 0x5c, 0x71, 0xe9, 0xb6, 0xde, 0x59, 0xaa, 0x52, 0xf7, 0x8f, 0xab, 0xfd, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x13, 0x4f, 0x31, 0x07, 0x04, 0x03, 0x00, 0x00,
}
