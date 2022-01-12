// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: id.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Node struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"ip"
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty" db:"ip"`
	// @inject_tag: db:"port"
	Port int64 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty" db:"port"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Node) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Node) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Node) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GetGlobalIdRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int64  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *GetGlobalIdRequest) Reset()         { *m = GetGlobalIdRequest{} }
func (m *GetGlobalIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetGlobalIdRequest) ProtoMessage()    {}
func (*GetGlobalIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{1}
}
func (m *GetGlobalIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGlobalIdRequest.Unmarshal(m, b)
}
func (m *GetGlobalIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGlobalIdRequest.Marshal(b, m, deterministic)
}
func (m *GetGlobalIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGlobalIdRequest.Merge(m, src)
}
func (m *GetGlobalIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetGlobalIdRequest.Size(m)
}
func (m *GetGlobalIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGlobalIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGlobalIdRequest proto.InternalMessageInfo

func (m *GetGlobalIdRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *GetGlobalIdRequest) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type GetGlobalIdReply struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetGlobalIdReply) Reset()         { *m = GetGlobalIdReply{} }
func (m *GetGlobalIdReply) String() string { return proto.CompactTextString(m) }
func (*GetGlobalIdReply) ProtoMessage()    {}
func (*GetGlobalIdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{2}
}
func (m *GetGlobalIdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGlobalIdReply.Unmarshal(m, b)
}
func (m *GetGlobalIdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGlobalIdReply.Marshal(b, m, deterministic)
}
func (m *GetGlobalIdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGlobalIdReply.Merge(m, src)
}
func (m *GetGlobalIdReply) XXX_Size() int {
	return xxx_messageInfo_GetGlobalIdReply.Size(m)
}
func (m *GetGlobalIdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGlobalIdReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGlobalIdReply proto.InternalMessageInfo

func (m *GetGlobalIdReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "pb.Node")
	proto.RegisterType((*GetGlobalIdRequest)(nil), "pb.GetGlobalIdRequest")
	proto.RegisterType((*GetGlobalIdReply)(nil), "pb.GetGlobalIdReply")
}

func init() { proto.RegisterFile("id.proto", fileDescriptor_4b3ad0c1fc883139) }

var fileDescriptor_4b3ad0c1fc883139 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x4c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x73,
	0xf5, 0x41, 0x2c, 0x88, 0x8c, 0x52, 0x05, 0x17, 0x8b, 0x5f, 0x7e, 0x4a, 0xaa, 0x10, 0x1f, 0x17,
	0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0x66, 0x0a, 0x98, 0x5f, 0x20,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x94, 0x59, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x90, 0x5f,
	0x54, 0x22, 0xc1, 0x0c, 0x56, 0x01, 0x66, 0x0b, 0xc9, 0x72, 0x71, 0x25, 0x17, 0xa5, 0x26, 0x96,
	0xa4, 0xa6, 0xc4, 0x27, 0x96, 0x48, 0xb0, 0x80, 0x65, 0x38, 0xa1, 0x22, 0x8e, 0x60, 0xe9, 0xd2,
	0x82, 0x14, 0x98, 0x34, 0x2b, 0x44, 0x1a, 0x2a, 0xe2, 0x58, 0xa2, 0x64, 0xc1, 0x25, 0xe4, 0x9e,
	0x5a, 0xe2, 0x9e, 0x93, 0x9f, 0x94, 0x98, 0xe3, 0x99, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x02, 0xb5, 0x97, 0x11, 0xc3, 0x5e, 0x26, 0x84, 0xbd, 0x4a, 0x4a, 0x5c, 0x02, 0x28, 0x3a, 0x0b,
	0x72, 0x2a, 0xd1, 0xdd, 0x6f, 0xe4, 0xc6, 0xc5, 0xea, 0x99, 0x12, 0x5c, 0x96, 0x2c, 0x64, 0xcb,
	0xc5, 0x8d, 0xa4, 0x58, 0x48, 0x4c, 0xaf, 0x20, 0x49, 0x0f, 0xd3, 0x5e, 0x29, 0x11, 0x0c, 0xf1,
	0x82, 0x9c, 0x4a, 0x25, 0x06, 0x27, 0x99, 0x13, 0x0f, 0xe5, 0x18, 0x2f, 0x3c, 0x94, 0x63, 0x98,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd8, 0x12, 0x0b,
	0x32, 0xf5, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x81, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x2a, 0x06, 0xe4, 0x6a, 0x01, 0x00, 0x00,
}
