// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: middle.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PageRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{0}
}
func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageRequest.Unmarshal(m, b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return xxx_messageInfo_PageRequest.Size(m)
}
func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PageRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type PageReply struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageReply) Reset()         { *m = PageReply{} }
func (m *PageReply) String() string { return proto.CompactTextString(m) }
func (*PageReply) ProtoMessage()    {}
func (*PageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{1}
}
func (m *PageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageReply.Unmarshal(m, b)
}
func (m *PageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageReply.Marshal(b, m, deterministic)
}
func (m *PageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageReply.Merge(m, src)
}
func (m *PageReply) XXX_Size() int {
	return xxx_messageInfo_PageReply.Size(m)
}
func (m *PageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PageReply.DiscardUnknown(m)
}

var xxx_messageInfo_PageReply proto.InternalMessageInfo

func (m *PageReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PageReply) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PageReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type AppRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Extra                string   `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppRequest) Reset()         { *m = AppRequest{} }
func (m *AppRequest) String() string { return proto.CompactTextString(m) }
func (*AppRequest) ProtoMessage()    {}
func (*AppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{2}
}
func (m *AppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppRequest.Unmarshal(m, b)
}
func (m *AppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppRequest.Marshal(b, m, deterministic)
}
func (m *AppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppRequest.Merge(m, src)
}
func (m *AppRequest) XXX_Size() int {
	return xxx_messageInfo_AppRequest.Size(m)
}
func (m *AppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppRequest proto.InternalMessageInfo

func (m *AppRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AppRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AppRequest) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

type AppReply struct {
	Apps                 []*App   `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppReply) Reset()         { *m = AppReply{} }
func (m *AppReply) String() string { return proto.CompactTextString(m) }
func (*AppReply) ProtoMessage()    {}
func (*AppReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{3}
}
func (m *AppReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppReply.Unmarshal(m, b)
}
func (m *AppReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppReply.Marshal(b, m, deterministic)
}
func (m *AppReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppReply.Merge(m, src)
}
func (m *AppReply) XXX_Size() int {
	return xxx_messageInfo_AppReply.Size(m)
}
func (m *AppReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AppReply.DiscardUnknown(m)
}

var xxx_messageInfo_AppReply proto.InternalMessageInfo

func (m *AppReply) GetApps() []*App {
	if m != nil {
		return m.Apps
	}
	return nil
}

type App struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	IsAuthorized         bool     `protobuf:"varint,2,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{4}
}
func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *App) GetIsAuthorized() bool {
	if m != nil {
		return m.IsAuthorized
	}
	return false
}

type CredentialReply struct {
	Items                []*KV    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialReply) Reset()         { *m = CredentialReply{} }
func (m *CredentialReply) String() string { return proto.CompactTextString(m) }
func (*CredentialReply) ProtoMessage()    {}
func (*CredentialReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{5}
}
func (m *CredentialReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialReply.Unmarshal(m, b)
}
func (m *CredentialReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialReply.Marshal(b, m, deterministic)
}
func (m *CredentialReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialReply.Merge(m, src)
}
func (m *CredentialReply) XXX_Size() int {
	return xxx_messageInfo_CredentialReply.Size(m)
}
func (m *CredentialReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialReply.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialReply proto.InternalMessageInfo

func (m *CredentialReply) GetItems() []*KV {
	if m != nil {
		return m.Items
	}
	return nil
}

type SettingReply struct {
	Items                []*KV    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingReply) Reset()         { *m = SettingReply{} }
func (m *SettingReply) String() string { return proto.CompactTextString(m) }
func (*SettingReply) ProtoMessage()    {}
func (*SettingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{6}
}
func (m *SettingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingReply.Unmarshal(m, b)
}
func (m *SettingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingReply.Marshal(b, m, deterministic)
}
func (m *SettingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingReply.Merge(m, src)
}
func (m *SettingReply) XXX_Size() int {
	return xxx_messageInfo_SettingReply.Size(m)
}
func (m *SettingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingReply.DiscardUnknown(m)
}

var xxx_messageInfo_SettingReply proto.InternalMessageInfo

func (m *SettingReply) GetItems() []*KV {
	if m != nil {
		return m.Items
	}
	return nil
}

type KVRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVRequest) Reset()         { *m = KVRequest{} }
func (m *KVRequest) String() string { return proto.CompactTextString(m) }
func (*KVRequest) ProtoMessage()    {}
func (*KVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{7}
}
func (m *KVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVRequest.Unmarshal(m, b)
}
func (m *KVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVRequest.Marshal(b, m, deterministic)
}
func (m *KVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVRequest.Merge(m, src)
}
func (m *KVRequest) XXX_Size() int {
	return xxx_messageInfo_KVRequest.Size(m)
}
func (m *KVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KVRequest proto.InternalMessageInfo

func (m *KVRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KVsRequest struct {
	Kvs                  []*KV    `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVsRequest) Reset()         { *m = KVsRequest{} }
func (m *KVsRequest) String() string { return proto.CompactTextString(m) }
func (*KVsRequest) ProtoMessage()    {}
func (*KVsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{8}
}
func (m *KVsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVsRequest.Unmarshal(m, b)
}
func (m *KVsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVsRequest.Marshal(b, m, deterministic)
}
func (m *KVsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVsRequest.Merge(m, src)
}
func (m *KVsRequest) XXX_Size() int {
	return xxx_messageInfo_KVsRequest.Size(m)
}
func (m *KVsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KVsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KVsRequest proto.InternalMessageInfo

func (m *KVsRequest) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
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
	return fileDescriptor_492fc6d32fb115aa, []int{9}
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

func init() {
	proto.RegisterType((*PageRequest)(nil), "pb.PageRequest")
	proto.RegisterType((*PageReply)(nil), "pb.PageReply")
	proto.RegisterType((*AppRequest)(nil), "pb.AppRequest")
	proto.RegisterType((*AppReply)(nil), "pb.AppReply")
	proto.RegisterType((*App)(nil), "pb.App")
	proto.RegisterType((*CredentialReply)(nil), "pb.CredentialReply")
	proto.RegisterType((*SettingReply)(nil), "pb.SettingReply")
	proto.RegisterType((*KVRequest)(nil), "pb.KVRequest")
	proto.RegisterType((*KVsRequest)(nil), "pb.KVsRequest")
	proto.RegisterType((*KV)(nil), "pb.KV")
}

func init() { proto.RegisterFile("middle.proto", fileDescriptor_492fc6d32fb115aa) }

var fileDescriptor_492fc6d32fb115aa = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x05, 0xec, 0x10, 0x98, 0xe0, 0x04, 0x6d, 0x7b, 0xb0, 0xd2, 0x1e, 0xa2, 0x3d, 0x24, 0x91,
	0x1a, 0x81, 0x02, 0x97, 0xde, 0x2a, 0x37, 0x07, 0x0e, 0x28, 0x4a, 0x03, 0x15, 0x87, 0x9e, 0xba,
	0xae, 0x47, 0xe9, 0x0a, 0x63, 0x6f, 0xed, 0x71, 0x14, 0xfa, 0x5f, 0xfa, 0x5f, 0xab, 0xdd, 0xb5,
	0x81, 0xc6, 0x6d, 0xc4, 0xa1, 0xb7, 0x9d, 0xe7, 0x99, 0xf7, 0xe6, 0xe3, 0xc9, 0xd0, 0x5b, 0xc9,
	0x28, 0x8a, 0x71, 0xa0, 0xb2, 0x94, 0x52, 0xd6, 0x52, 0xe1, 0x29, 0x84, 0x22, 0x2f, 0x63, 0x7e,
	0x0f, 0x47, 0x9f, 0xc4, 0x03, 0xce, 0xf0, 0x47, 0x81, 0x39, 0x31, 0x06, 0x6e, 0x51, 0xc8, 0xc8,
	0x6f, 0x9e, 0x35, 0x2f, 0xbb, 0x33, 0xf3, 0x66, 0xaf, 0xe1, 0x80, 0x24, 0xc5, 0xe8, 0xb7, 0x0c,
	0x68, 0x03, 0xe6, 0xc3, 0xe1, 0xb7, 0x34, 0x21, 0x4c, 0xc8, 0x77, 0x0c, 0x5e, 0x85, 0xfc, 0x0e,
	0xba, 0x96, 0x52, 0xc5, 0xeb, 0xff, 0x42, 0xf8, 0x15, 0x20, 0x50, 0x6a, 0xa7, 0xc5, 0x44, 0xac,
	0xb0, 0x62, 0xd4, 0x6f, 0x8d, 0xd1, 0x5a, 0x55, 0x84, 0xe6, 0x6d, 0x54, 0xd2, 0x25, 0x26, 0x25,
	0x9b, 0x0d, 0x34, 0x8a, 0x4f, 0x94, 0x09, 0xdf, 0xb5, 0xa8, 0x09, 0xf8, 0x05, 0x74, 0x8c, 0x82,
	0xee, 0xf8, 0x0d, 0xb8, 0x42, 0xa9, 0xdc, 0x6f, 0x9e, 0x39, 0x97, 0x47, 0xa3, 0xc3, 0x81, 0x0a,
	0x07, 0xfa, 0x9b, 0x01, 0xf9, 0x07, 0x70, 0x02, 0xa5, 0xb6, 0x13, 0x34, 0x77, 0x27, 0xe0, 0xd0,
	0x93, 0x79, 0x50, 0xd0, 0xf7, 0x34, 0x93, 0x3f, 0x31, 0x32, 0xdd, 0x74, 0x66, 0x7f, 0x60, 0x7c,
	0x08, 0x27, 0x37, 0x19, 0x46, 0x98, 0x90, 0x14, 0xb1, 0x15, 0x7c, 0x0b, 0x07, 0x92, 0x70, 0x55,
	0x29, 0xb6, 0xb5, 0xe2, 0x74, 0x31, 0xb3, 0x20, 0xbf, 0x82, 0xde, 0x1c, 0x89, 0x64, 0xf2, 0xb0,
	0x4f, 0xf6, 0x18, 0xba, 0xd3, 0x45, 0xb5, 0xa9, 0x3e, 0x38, 0x4b, 0x5c, 0x97, 0x3d, 0xea, 0xa7,
	0xee, 0xfb, 0x51, 0xc4, 0xc5, 0x66, 0xf3, 0x26, 0xe0, 0xe7, 0x00, 0xd3, 0x45, 0x5e, 0x55, 0xf9,
	0xe0, 0x2c, 0x1f, 0x9f, 0xd3, 0x6b, 0x88, 0x5f, 0x41, 0x6b, 0xba, 0xd8, 0x97, 0x75, 0xf4, 0xcb,
	0x85, 0xf6, 0xad, 0xb1, 0x1e, 0x1b, 0x00, 0xdc, 0x64, 0x28, 0x08, 0xb5, 0x2f, 0xd8, 0x89, 0xe6,
	0xdc, 0x31, 0xdd, 0xa9, 0xa7, 0x81, 0xcf, 0xf8, 0x44, 0x66, 0x42, 0xde, 0x60, 0xef, 0xe0, 0x70,
	0x82, 0xf4, 0x42, 0xf2, 0xc6, 0x5f, 0xbc, 0xc1, 0xce, 0xa1, 0x75, 0x9f, 0xd9, 0x3c, 0xcb, 0xf1,
	0x0f, 0xd2, 0x0b, 0x70, 0x03, 0xa5, 0xf2, 0x7a, 0x66, 0xaf, 0x3a, 0x71, 0x99, 0x78, 0x0d, 0xde,
	0x9c, 0xd2, 0x0c, 0x03, 0xa5, 0xee, 0xf4, 0xe5, 0xd8, 0xf1, 0x26, 0xc1, 0x16, 0x98, 0x78, 0x4e,
	0x82, 0x36, 0x3d, 0xbc, 0x87, 0xe3, 0x09, 0xd2, 0xf6, 0xb0, 0x7f, 0x51, 0x79, 0xa5, 0x81, 0x67,
	0xa7, 0xe7, 0x0d, 0x36, 0x86, 0xbe, 0x5d, 0xcd, 0xf6, 0x93, 0xd5, 0xdb, 0x5e, 0xa4, 0x3e, 0xca,
	0x35, 0xc0, 0x04, 0xa9, 0xb4, 0x45, 0x5d, 0xaa, 0x6f, 0xfa, 0xdb, 0x31, 0x0d, 0x6f, 0xb0, 0x21,
	0x78, 0x56, 0xa7, 0xaa, 0xf2, 0xca, 0xcb, 0xbe, 0x7c, 0x83, 0x5b, 0x4c, 0x8a, 0x3d, 0x76, 0x3b,
	0x02, 0xaf, 0xf2, 0xb8, 0x20, 0x99, 0x26, 0xf5, 0x92, 0xda, 0xce, 0x3e, 0x76, 0xbe, 0xb4, 0x85,
	0x92, 0x43, 0x15, 0x86, 0x6d, 0xf3, 0x2b, 0x1a, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x37, 0xc4,
	0x09, 0xde, 0xaa, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MiddleClient is the client API for Middle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MiddleClient interface {
	CreatePage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*TextReply, error)
	GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageReply, error)
	Qr(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error)
	Apps(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*AppReply, error)
	StoreAppOAuth(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*StateReply, error)
	GetCredentials(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*CredentialReply, error)
	CreateCredential(ctx context.Context, in *KVsRequest, opts ...grpc.CallOption) (*TextReply, error)
	GetSetting(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*SettingReply, error)
	CreateSetting(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*TextReply, error)
	GetMenu(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error)
	Authorization(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*StateReply, error)
}

type middleClient struct {
	cc *grpc.ClientConn
}

func NewMiddleClient(cc *grpc.ClientConn) MiddleClient {
	return &middleClient{cc}
}

func (c *middleClient) CreatePage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/CreatePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageReply, error) {
	out := new(PageReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Qr(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/Qr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Apps(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*AppReply, error) {
	out := new(AppReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/Apps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) StoreAppOAuth(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/StoreAppOAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) GetCredentials(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*CredentialReply, error) {
	out := new(CredentialReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/GetCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) CreateCredential(ctx context.Context, in *KVsRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/CreateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) GetSetting(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*SettingReply, error) {
	out := new(SettingReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/GetSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) CreateSetting(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/CreateSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) GetMenu(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/GetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Authorization(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/Authorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddleServer is the server API for Middle service.
type MiddleServer interface {
	CreatePage(context.Context, *PageRequest) (*TextReply, error)
	GetPage(context.Context, *PageRequest) (*PageReply, error)
	Qr(context.Context, *TextRequest) (*TextReply, error)
	Apps(context.Context, *TextRequest) (*AppReply, error)
	StoreAppOAuth(context.Context, *AppRequest) (*StateReply, error)
	GetCredentials(context.Context, *TextRequest) (*CredentialReply, error)
	CreateCredential(context.Context, *KVsRequest) (*TextReply, error)
	GetSetting(context.Context, *TextRequest) (*SettingReply, error)
	CreateSetting(context.Context, *KVRequest) (*TextReply, error)
	GetMenu(context.Context, *TextRequest) (*TextReply, error)
	Authorization(context.Context, *TextRequest) (*StateReply, error)
}

// UnimplementedMiddleServer can be embedded to have forward compatible implementations.
type UnimplementedMiddleServer struct {
}

func (*UnimplementedMiddleServer) CreatePage(ctx context.Context, req *PageRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}
func (*UnimplementedMiddleServer) GetPage(ctx context.Context, req *PageRequest) (*PageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (*UnimplementedMiddleServer) Qr(ctx context.Context, req *TextRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Qr not implemented")
}
func (*UnimplementedMiddleServer) Apps(ctx context.Context, req *TextRequest) (*AppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apps not implemented")
}
func (*UnimplementedMiddleServer) StoreAppOAuth(ctx context.Context, req *AppRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAppOAuth not implemented")
}
func (*UnimplementedMiddleServer) GetCredentials(ctx context.Context, req *TextRequest) (*CredentialReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentials not implemented")
}
func (*UnimplementedMiddleServer) CreateCredential(ctx context.Context, req *KVsRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredential not implemented")
}
func (*UnimplementedMiddleServer) GetSetting(ctx context.Context, req *TextRequest) (*SettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetting not implemented")
}
func (*UnimplementedMiddleServer) CreateSetting(ctx context.Context, req *KVRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSetting not implemented")
}
func (*UnimplementedMiddleServer) GetMenu(ctx context.Context, req *TextRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (*UnimplementedMiddleServer) Authorization(ctx context.Context, req *TextRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorization not implemented")
}

func RegisterMiddleServer(s *grpc.Server, srv MiddleServer) {
	s.RegisterService(&_Middle_serviceDesc, srv)
}

func _Middle_CreatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).CreatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/CreatePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).CreatePage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).GetPage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Qr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Qr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Qr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Qr(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Apps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Apps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Apps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Apps(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_StoreAppOAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).StoreAppOAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/StoreAppOAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).StoreAppOAuth(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_GetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).GetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/GetCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).GetCredentials(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_CreateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).CreateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/CreateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).CreateCredential(ctx, req.(*KVsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_GetSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).GetSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/GetSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).GetSetting(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_CreateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).CreateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/CreateSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).CreateSetting(ctx, req.(*KVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_GetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).GetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/GetMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).GetMenu(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Authorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Authorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Authorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Authorization(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Middle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Middle",
	HandlerType: (*MiddleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePage",
			Handler:    _Middle_CreatePage_Handler,
		},
		{
			MethodName: "GetPage",
			Handler:    _Middle_GetPage_Handler,
		},
		{
			MethodName: "Qr",
			Handler:    _Middle_Qr_Handler,
		},
		{
			MethodName: "Apps",
			Handler:    _Middle_Apps_Handler,
		},
		{
			MethodName: "StoreAppOAuth",
			Handler:    _Middle_StoreAppOAuth_Handler,
		},
		{
			MethodName: "GetCredentials",
			Handler:    _Middle_GetCredentials_Handler,
		},
		{
			MethodName: "CreateCredential",
			Handler:    _Middle_CreateCredential_Handler,
		},
		{
			MethodName: "GetSetting",
			Handler:    _Middle_GetSetting_Handler,
		},
		{
			MethodName: "CreateSetting",
			Handler:    _Middle_CreateSetting_Handler,
		},
		{
			MethodName: "GetMenu",
			Handler:    _Middle_GetMenu_Handler,
		},
		{
			MethodName: "Authorization",
			Handler:    _Middle_Authorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middle.proto",
}
