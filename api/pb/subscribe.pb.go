// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subscribe.proto

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

type SubscribeRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{0}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SubscribeReply struct {
	Text                 []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeReply) Reset()         { *m = SubscribeReply{} }
func (m *SubscribeReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeReply) ProtoMessage()    {}
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{1}
}
func (m *SubscribeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReply.Unmarshal(m, b)
}
func (m *SubscribeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReply.Marshal(b, m, deterministic)
}
func (m *SubscribeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReply.Merge(m, src)
}
func (m *SubscribeReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeReply.Size(m)
}
func (m *SubscribeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReply proto.InternalMessageInfo

func (m *SubscribeReply) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

type State struct {
	State                bool     `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{2}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "pb.SubscribeRequest")
	proto.RegisterType((*SubscribeReply)(nil), "pb.SubscribeReply")
	proto.RegisterType((*State)(nil), "pb.State")
}

func init() { proto.RegisterFile("subscribe.proto", fileDescriptor_38d2980c9543da44) }

var fileDescriptor_38d2980c9543da44 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0x4d, 0x2a,
	0x4e, 0x2e, 0xca, 0x4c, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0x52, 0xe3, 0x12, 0x08, 0x86, 0x09, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71,
	0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x2a,
	0x5c, 0x7c, 0x48, 0xea, 0x0a, 0x72, 0x2a, 0x91, 0x54, 0x31, 0xc3, 0x55, 0xc9, 0x72, 0xb1, 0x06,
	0x97, 0x24, 0x96, 0xa4, 0x0a, 0x89, 0x70, 0xb1, 0x16, 0x83, 0x18, 0x60, 0x33, 0x38, 0x82, 0x20,
	0x1c, 0xa3, 0x69, 0x8c, 0x5c, 0x9c, 0x70, 0x53, 0x84, 0x8c, 0xb8, 0x58, 0x7c, 0x32, 0x8b, 0x4b,
	0x84, 0x44, 0xf4, 0x0a, 0x92, 0xf4, 0xd0, 0x1d, 0x21, 0x25, 0x84, 0x26, 0x5a, 0x90, 0x53, 0xa9,
	0xc4, 0x20, 0xa4, 0xc9, 0xc5, 0xe2, 0x5f, 0x90, 0x9a, 0x87, 0x43, 0x0f, 0x27, 0x58, 0x14, 0x64,
	0x95, 0x12, 0x83, 0x90, 0x16, 0x17, 0xab, 0x73, 0x4e, 0x7e, 0x71, 0x2a, 0x11, 0x6a, 0x9d, 0x38,
	0xa2, 0xd8, 0x12, 0x0b, 0x32, 0xf5, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x41, 0x63, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0x44, 0xf9, 0xe8, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscribeClient is the client API for Subscribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscribeClient interface {
	List(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error)
	Open(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*State, error)
	Close(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*State, error)
}

type subscribeClient struct {
	cc *grpc.ClientConn
}

func NewSubscribeClient(cc *grpc.ClientConn) SubscribeClient {
	return &subscribeClient{cc}
}

func (c *subscribeClient) List(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error) {
	out := new(SubscribeReply)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Open(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Close(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeServer is the server API for Subscribe service.
type SubscribeServer interface {
	List(context.Context, *SubscribeRequest) (*SubscribeReply, error)
	Open(context.Context, *SubscribeRequest) (*State, error)
	Close(context.Context, *SubscribeRequest) (*State, error)
}

// UnimplementedSubscribeServer can be embedded to have forward compatible implementations.
type UnimplementedSubscribeServer struct {
}

func (*UnimplementedSubscribeServer) List(ctx context.Context, req *SubscribeRequest) (*SubscribeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedSubscribeServer) Open(ctx context.Context, req *SubscribeRequest) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (*UnimplementedSubscribeServer) Close(ctx context.Context, req *SubscribeRequest) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterSubscribeServer(s *grpc.Server, srv SubscribeServer) {
	s.RegisterService(&_Subscribe_serviceDesc, srv)
}

func _Subscribe_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).List(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Open(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Close(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Subscribe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Subscribe",
	HandlerType: (*SubscribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Subscribe_List_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _Subscribe_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Subscribe_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscribe.proto",
}