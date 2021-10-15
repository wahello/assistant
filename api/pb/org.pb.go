// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: org.proto

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

type Objective struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"user_id"
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" db:"name"`
	// @inject_tag: db:"tag_id"
	TagId int64 `protobuf:"varint,4,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty" db:"tag_id"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Objective) Reset()         { *m = Objective{} }
func (m *Objective) String() string { return proto.CompactTextString(m) }
func (*Objective) ProtoMessage()    {}
func (*Objective) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{0}
}
func (m *Objective) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Objective.Unmarshal(m, b)
}
func (m *Objective) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Objective.Marshal(b, m, deterministic)
}
func (m *Objective) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Objective.Merge(m, src)
}
func (m *Objective) XXX_Size() int {
	return xxx_messageInfo_Objective.Size(m)
}
func (m *Objective) XXX_DiscardUnknown() {
	xxx_messageInfo_Objective.DiscardUnknown(m)
}

var xxx_messageInfo_Objective proto.InternalMessageInfo

func (m *Objective) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Objective) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Objective) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Objective) GetTagId() int64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *Objective) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Objective) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type KeyResult struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"user_id"
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" db:"name"`
	// @inject_tag: db:"objective_id"
	ObjectiveId int64 `protobuf:"varint,4,opt,name=objective_id,json=objectiveId,proto3" json:"objective_id,omitempty" db:"objective_id"`
	// @inject_tag: db:"tag_id"
	TagId int64 `protobuf:"varint,5,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty" db:"tag_id"`
	// @inject_tag: db:"complete"
	Complete bool `protobuf:"varint,6,opt,name=complete,proto3" json:"complete,omitempty" db:"complete"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *KeyResult) Reset()         { *m = KeyResult{} }
func (m *KeyResult) String() string { return proto.CompactTextString(m) }
func (*KeyResult) ProtoMessage()    {}
func (*KeyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{1}
}
func (m *KeyResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyResult.Unmarshal(m, b)
}
func (m *KeyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyResult.Marshal(b, m, deterministic)
}
func (m *KeyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyResult.Merge(m, src)
}
func (m *KeyResult) XXX_Size() int {
	return xxx_messageInfo_KeyResult.Size(m)
}
func (m *KeyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyResult.DiscardUnknown(m)
}

var xxx_messageInfo_KeyResult proto.InternalMessageInfo

func (m *KeyResult) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KeyResult) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *KeyResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyResult) GetObjectiveId() int64 {
	if m != nil {
		return m.ObjectiveId
	}
	return 0
}

func (m *KeyResult) GetTagId() int64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *KeyResult) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *KeyResult) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *KeyResult) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ObjectiveRequest struct {
	Objective *Objective `protobuf:"bytes,1,opt,name=objective,proto3" json:"objective,omitempty"`
	Tag       string     `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (m *ObjectiveRequest) Reset()         { *m = ObjectiveRequest{} }
func (m *ObjectiveRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectiveRequest) ProtoMessage()    {}
func (*ObjectiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{2}
}
func (m *ObjectiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectiveRequest.Unmarshal(m, b)
}
func (m *ObjectiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectiveRequest.Marshal(b, m, deterministic)
}
func (m *ObjectiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectiveRequest.Merge(m, src)
}
func (m *ObjectiveRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectiveRequest.Size(m)
}
func (m *ObjectiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectiveRequest proto.InternalMessageInfo

func (m *ObjectiveRequest) GetObjective() *Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

func (m *ObjectiveRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ObjectiveReply struct {
	Objective *Objective `protobuf:"bytes,1,opt,name=objective,proto3" json:"objective,omitempty"`
}

func (m *ObjectiveReply) Reset()         { *m = ObjectiveReply{} }
func (m *ObjectiveReply) String() string { return proto.CompactTextString(m) }
func (*ObjectiveReply) ProtoMessage()    {}
func (*ObjectiveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{3}
}
func (m *ObjectiveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectiveReply.Unmarshal(m, b)
}
func (m *ObjectiveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectiveReply.Marshal(b, m, deterministic)
}
func (m *ObjectiveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectiveReply.Merge(m, src)
}
func (m *ObjectiveReply) XXX_Size() int {
	return xxx_messageInfo_ObjectiveReply.Size(m)
}
func (m *ObjectiveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectiveReply.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectiveReply proto.InternalMessageInfo

func (m *ObjectiveReply) GetObjective() *Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

type ObjectivesReply struct {
	Objective []*Objective `protobuf:"bytes,1,rep,name=objective,proto3" json:"objective,omitempty"`
}

func (m *ObjectivesReply) Reset()         { *m = ObjectivesReply{} }
func (m *ObjectivesReply) String() string { return proto.CompactTextString(m) }
func (*ObjectivesReply) ProtoMessage()    {}
func (*ObjectivesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{4}
}
func (m *ObjectivesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectivesReply.Unmarshal(m, b)
}
func (m *ObjectivesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectivesReply.Marshal(b, m, deterministic)
}
func (m *ObjectivesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectivesReply.Merge(m, src)
}
func (m *ObjectivesReply) XXX_Size() int {
	return xxx_messageInfo_ObjectivesReply.Size(m)
}
func (m *ObjectivesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectivesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectivesReply proto.InternalMessageInfo

func (m *ObjectivesReply) GetObjective() []*Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

type KeyResultRequest struct {
	KeyResult *KeyResult `protobuf:"bytes,1,opt,name=key_result,json=keyResult,proto3" json:"key_result,omitempty"`
	Tag       string     `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (m *KeyResultRequest) Reset()         { *m = KeyResultRequest{} }
func (m *KeyResultRequest) String() string { return proto.CompactTextString(m) }
func (*KeyResultRequest) ProtoMessage()    {}
func (*KeyResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{5}
}
func (m *KeyResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyResultRequest.Unmarshal(m, b)
}
func (m *KeyResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyResultRequest.Marshal(b, m, deterministic)
}
func (m *KeyResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyResultRequest.Merge(m, src)
}
func (m *KeyResultRequest) XXX_Size() int {
	return xxx_messageInfo_KeyResultRequest.Size(m)
}
func (m *KeyResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyResultRequest proto.InternalMessageInfo

func (m *KeyResultRequest) GetKeyResult() *KeyResult {
	if m != nil {
		return m.KeyResult
	}
	return nil
}

func (m *KeyResultRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type KeyResultReply struct {
	KeyResult *KeyResult `protobuf:"bytes,1,opt,name=key_result,json=keyResult,proto3" json:"key_result,omitempty"`
}

func (m *KeyResultReply) Reset()         { *m = KeyResultReply{} }
func (m *KeyResultReply) String() string { return proto.CompactTextString(m) }
func (*KeyResultReply) ProtoMessage()    {}
func (*KeyResultReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{6}
}
func (m *KeyResultReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyResultReply.Unmarshal(m, b)
}
func (m *KeyResultReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyResultReply.Marshal(b, m, deterministic)
}
func (m *KeyResultReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyResultReply.Merge(m, src)
}
func (m *KeyResultReply) XXX_Size() int {
	return xxx_messageInfo_KeyResultReply.Size(m)
}
func (m *KeyResultReply) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyResultReply.DiscardUnknown(m)
}

var xxx_messageInfo_KeyResultReply proto.InternalMessageInfo

func (m *KeyResultReply) GetKeyResult() *KeyResult {
	if m != nil {
		return m.KeyResult
	}
	return nil
}

type KeyResultsReply struct {
	Result []*KeyResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (m *KeyResultsReply) Reset()         { *m = KeyResultsReply{} }
func (m *KeyResultsReply) String() string { return proto.CompactTextString(m) }
func (*KeyResultsReply) ProtoMessage()    {}
func (*KeyResultsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb462779e28924f, []int{7}
}
func (m *KeyResultsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyResultsReply.Unmarshal(m, b)
}
func (m *KeyResultsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyResultsReply.Marshal(b, m, deterministic)
}
func (m *KeyResultsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyResultsReply.Merge(m, src)
}
func (m *KeyResultsReply) XXX_Size() int {
	return xxx_messageInfo_KeyResultsReply.Size(m)
}
func (m *KeyResultsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyResultsReply.DiscardUnknown(m)
}

var xxx_messageInfo_KeyResultsReply proto.InternalMessageInfo

func (m *KeyResultsReply) GetResult() []*KeyResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Objective)(nil), "pb.Objective")
	proto.RegisterType((*KeyResult)(nil), "pb.KeyResult")
	proto.RegisterType((*ObjectiveRequest)(nil), "pb.ObjectiveRequest")
	proto.RegisterType((*ObjectiveReply)(nil), "pb.ObjectiveReply")
	proto.RegisterType((*ObjectivesReply)(nil), "pb.ObjectivesReply")
	proto.RegisterType((*KeyResultRequest)(nil), "pb.KeyResultRequest")
	proto.RegisterType((*KeyResultReply)(nil), "pb.KeyResultReply")
	proto.RegisterType((*KeyResultsReply)(nil), "pb.KeyResultsReply")
}

func init() { proto.RegisterFile("org.proto", fileDescriptor_ccb462779e28924f) }

var fileDescriptor_ccb462779e28924f = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0xb3, 0xcd, 0x36, 0xf3, 0xda, 0x26, 0x61, 0x1a, 0x31, 0x04, 0x5d, 0xea, 0x82,
	0x50, 0x50, 0x52, 0xa8, 0x17, 0x2b, 0x5a, 0xa8, 0x0a, 0x12, 0x3c, 0x14, 0xb7, 0x37, 0x2f, 0x61,
	0x36, 0xf3, 0x58, 0xd6, 0xa6, 0xee, 0xb8, 0x3b, 0x29, 0xe4, 0x5b, 0xf8, 0x05, 0xfc, 0x3e, 0x1e,
	0x7b, 0xf4, 0xe0, 0xa5, 0xe4, 0xe2, 0xc7, 0x90, 0x99, 0x4d, 0x66, 0x37, 0x89, 0x6d, 0x58, 0xbc,
	0xcd, 0xfc, 0xdf, 0xfc, 0x1f, 0xff, 0xdf, 0x7b, 0xd9, 0x00, 0x49, 0xd2, 0xa8, 0x2f, 0xd2, 0x44,
	0x26, 0xb4, 0x26, 0xc2, 0x1e, 0x84, 0x2c, 0xc3, 0xfc, 0xde, 0xeb, 0x44, 0x49, 0x94, 0xe8, 0xe3,
	0x91, 0x3a, 0xe5, 0xaa, 0xff, 0xc3, 0x06, 0x72, 0x1e, 0x7e, 0xc1, 0x91, 0x8c, 0xaf, 0x91, 0x36,
	0xa1, 0x16, 0xf3, 0xae, 0x7d, 0x60, 0x1f, 0x3a, 0x41, 0x2d, 0xe6, 0xf4, 0x21, 0x6c, 0x4f, 0x32,
	0x4c, 0x87, 0x31, 0xef, 0xd6, 0xb4, 0xe8, 0xaa, 0xeb, 0x80, 0x53, 0x0a, 0x5b, 0x5f, 0xd9, 0x15,
	0x76, 0x9d, 0x03, 0xfb, 0x90, 0x04, 0xfa, 0x4c, 0x1f, 0x80, 0x2b, 0x59, 0xa4, 0xde, 0x6e, 0xe9,
	0xb7, 0x75, 0xc9, 0xa2, 0x01, 0xa7, 0x8f, 0x01, 0x46, 0x29, 0x32, 0x89, 0x7c, 0xc8, 0x64, 0xb7,
	0xae, 0x4b, 0x64, 0xae, 0x9c, 0x49, 0x55, 0x9e, 0x08, 0xbe, 0x28, 0xbb, 0x79, 0x79, 0xae, 0x9c,
	0x49, 0xff, 0xb7, 0x0d, 0xe4, 0x23, 0x4e, 0x03, 0xcc, 0x26, 0x63, 0xf9, 0x7f, 0xf9, 0x9e, 0xc0,
	0x6e, 0xb2, 0x20, 0x2d, 0x52, 0xee, 0x18, 0x6d, 0xc0, 0x4b, 0x08, 0xf5, 0x32, 0x42, 0x0f, 0x1a,
	0xa3, 0xe4, 0x4a, 0x8c, 0x51, 0xa2, 0x4e, 0xd8, 0x08, 0xcc, 0x7d, 0x05, 0x6f, 0xfb, 0x7e, 0xbc,
	0xc6, 0x2a, 0xde, 0x27, 0x68, 0x9b, 0xe9, 0x07, 0xf8, 0x6d, 0x82, 0x99, 0xa4, 0xcf, 0x80, 0x98,
	0x4c, 0x9a, 0x75, 0xe7, 0x78, 0xaf, 0x2f, 0xc2, 0x7e, 0xf1, 0xb0, 0xa8, 0xd3, 0x36, 0x38, 0x92,
	0x45, 0x9a, 0x9e, 0x04, 0xea, 0xe8, 0xbf, 0x81, 0x66, 0xa9, 0xa5, 0x18, 0x4f, 0x2b, 0x35, 0xf4,
	0x4f, 0xa1, 0x65, 0xf4, 0xec, 0x9f, 0x7e, 0xe7, 0x5e, 0x7f, 0x00, 0x6d, 0xb3, 0xaf, 0x05, 0xd1,
	0x73, 0x80, 0x4b, 0x9c, 0x0e, 0x53, 0x2d, 0x96, 0x13, 0x14, 0x2f, 0xc9, 0xa5, 0x59, 0xf2, 0x3a,
	0xd2, 0x29, 0x34, 0x4b, 0x3d, 0x55, 0xa4, 0x4a, 0x1d, 0xfd, 0x97, 0xd0, 0x32, 0xfa, 0x9c, 0xe9,
	0x29, 0xb8, 0xc6, 0xec, 0xac, 0x9b, 0xe7, 0xc5, 0xe3, 0x3f, 0x0e, 0xb8, 0xe7, 0x69, 0x74, 0x71,
	0x3d, 0xa2, 0x27, 0xd0, 0x7a, 0xa7, 0xd7, 0x5a, 0x7c, 0x2e, 0x9d, 0xe5, 0x29, 0xe4, 0xb4, 0xbd,
	0xa6, 0x52, 0x2f, 0x24, 0x93, 0xf9, 0xf8, 0x7d, 0x8b, 0xbe, 0x82, 0xdd, 0x0f, 0x28, 0x37, 0xf9,
	0xe8, 0x8a, 0x9a, 0x7b, 0x5f, 0xc3, 0x5e, 0xd9, 0x9b, 0xdd, 0x61, 0xde, 0x5f, 0x52, 0xb3, 0x85,
	0xfb, 0x04, 0x5a, 0xef, 0x51, 0xfd, 0x4e, 0xab, 0x87, 0x36, 0xbc, 0xc5, 0xe7, 0xd7, 0x59, 0x1e,
	0xd2, 0x06, 0xde, 0x4d, 0x3e, 0xba, 0xa2, 0x96, 0x79, 0x8b, 0x75, 0xdd, 0x61, 0xde, 0x5f, 0x52,
	0xd7, 0x79, 0x2b, 0x87, 0x7e, 0xfb, 0xe8, 0xe7, 0xad, 0x67, 0xdf, 0xdc, 0x7a, 0xd6, 0xf7, 0x99,
	0x67, 0xdd, 0xcc, 0x3c, 0xeb, 0xd7, 0xcc, 0xb3, 0x3e, 0xbb, 0x4c, 0xc4, 0x47, 0x22, 0x0c, 0x5d,
	0xfd, 0x77, 0xf9, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xc7, 0xeb, 0xde, 0x61, 0x05,
	0x00, 0x00,
}
