// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TextRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextRequest) Reset()         { *m = TextRequest{} }
func (m *TextRequest) String() string { return proto.CompactTextString(m) }
func (*TextRequest) ProtoMessage()    {}
func (*TextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}
func (m *TextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextRequest.Unmarshal(m, b)
}
func (m *TextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextRequest.Marshal(b, m, deterministic)
}
func (m *TextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextRequest.Merge(m, src)
}
func (m *TextRequest) XXX_Size() int {
	return xxx_messageInfo_TextRequest.Size(m)
}
func (m *TextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TextRequest proto.InternalMessageInfo

func (m *TextRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type TextReply struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextReply) Reset()         { *m = TextReply{} }
func (m *TextReply) String() string { return proto.CompactTextString(m) }
func (*TextReply) ProtoMessage()    {}
func (*TextReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}
func (m *TextReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextReply.Unmarshal(m, b)
}
func (m *TextReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextReply.Marshal(b, m, deterministic)
}
func (m *TextReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextReply.Merge(m, src)
}
func (m *TextReply) XXX_Size() int {
	return xxx_messageInfo_TextReply.Size(m)
}
func (m *TextReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TextReply.DiscardUnknown(m)
}

var xxx_messageInfo_TextReply proto.InternalMessageInfo

func (m *TextReply) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type StateReply struct {
	State                bool     `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateReply) Reset()         { *m = StateReply{} }
func (m *StateReply) String() string { return proto.CompactTextString(m) }
func (*StateReply) ProtoMessage()    {}
func (*StateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}
func (m *StateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateReply.Unmarshal(m, b)
}
func (m *StateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateReply.Marshal(b, m, deterministic)
}
func (m *StateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateReply.Merge(m, src)
}
func (m *StateReply) XXX_Size() int {
	return xxx_messageInfo_StateReply.Size(m)
}
func (m *StateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StateReply.DiscardUnknown(m)
}

var xxx_messageInfo_StateReply proto.InternalMessageInfo

func (m *StateReply) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{3}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UuidRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UuidRequest) Reset()         { *m = UuidRequest{} }
func (m *UuidRequest) String() string { return proto.CompactTextString(m) }
func (*UuidRequest) ProtoMessage()    {}
func (*UuidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{4}
}
func (m *UuidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UuidRequest.Unmarshal(m, b)
}
func (m *UuidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UuidRequest.Marshal(b, m, deterministic)
}
func (m *UuidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UuidRequest.Merge(m, src)
}
func (m *UuidRequest) XXX_Size() int {
	return xxx_messageInfo_UuidRequest.Size(m)
}
func (m *UuidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UuidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UuidRequest proto.InternalMessageInfo

func (m *UuidRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type UuidReply struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UuidReply) Reset()         { *m = UuidReply{} }
func (m *UuidReply) String() string { return proto.CompactTextString(m) }
func (*UuidReply) ProtoMessage()    {}
func (*UuidReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{5}
}
func (m *UuidReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UuidReply.Unmarshal(m, b)
}
func (m *UuidReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UuidReply.Marshal(b, m, deterministic)
}
func (m *UuidReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UuidReply.Merge(m, src)
}
func (m *UuidReply) XXX_Size() int {
	return xxx_messageInfo_UuidReply.Size(m)
}
func (m *UuidReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UuidReply.DiscardUnknown(m)
}

var xxx_messageInfo_UuidReply proto.InternalMessageInfo

func (m *UuidReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type IdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{6}
}
func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type IdReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReply) Reset()         { *m = IdReply{} }
func (m *IdReply) String() string { return proto.CompactTextString(m) }
func (*IdReply) ProtoMessage()    {}
func (*IdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{7}
}
func (m *IdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReply.Unmarshal(m, b)
}
func (m *IdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReply.Marshal(b, m, deterministic)
}
func (m *IdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReply.Merge(m, src)
}
func (m *IdReply) XXX_Size() int {
	return xxx_messageInfo_IdReply.Size(m)
}
func (m *IdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReply.DiscardUnknown(m)
}

var xxx_messageInfo_IdReply proto.InternalMessageInfo

func (m *IdReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*TextRequest)(nil), "pb.TextRequest")
	proto.RegisterType((*TextReply)(nil), "pb.TextReply")
	proto.RegisterType((*StateReply)(nil), "pb.StateReply")
	proto.RegisterType((*KV)(nil), "pb.KV")
	proto.RegisterType((*UuidRequest)(nil), "pb.UuidRequest")
	proto.RegisterType((*UuidReply)(nil), "pb.UuidReply")
	proto.RegisterType((*IdRequest)(nil), "pb.IdRequest")
	proto.RegisterType((*IdReply)(nil), "pb.IdReply")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe4, 0xe2, 0x0e, 0x49,
	0xad, 0x28, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad,
	0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xe4, 0xb9, 0x38, 0x21, 0x4a,
	0x0a, 0x72, 0x2a, 0xb1, 0x2a, 0x50, 0xe2, 0xe2, 0x0a, 0x2e, 0x49, 0x2c, 0x49, 0x85, 0xa8, 0x10,
	0xe1, 0x62, 0x2d, 0x06, 0xf1, 0xc0, 0x4a, 0x38, 0x82, 0x20, 0x1c, 0x25, 0x1d, 0x2e, 0x26, 0xef,
	0x30, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0xa8, 0x66, 0x10, 0x13, 0xa4, 0xba, 0x2c, 0x31,
	0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x80, 0x5c, 0x15, 0x5a, 0x9a, 0x99, 0x82, 0xe4,
	0xaa, 0xd2, 0xd2, 0xcc, 0x14, 0x98, 0xa5, 0x20, 0x36, 0xc8, 0x55, 0x10, 0x25, 0x50, 0x57, 0x61,
	0x28, 0x90, 0xe6, 0xe2, 0xf4, 0x84, 0x9b, 0xc0, 0xc7, 0xc5, 0x04, 0x95, 0x66, 0x0e, 0x62, 0xca,
	0x4c, 0x51, 0x92, 0xe4, 0x62, 0xf7, 0x84, 0xea, 0x45, 0x93, 0x72, 0xe2, 0x88, 0x62, 0x4b, 0x2c,
	0xc8, 0xd4, 0x2f, 0x48, 0x4a, 0x62, 0x03, 0x07, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb0,
	0x31, 0xb6, 0x98, 0x34, 0x01, 0x00, 0x00,
}
