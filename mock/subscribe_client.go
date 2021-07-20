// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/subscribe.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockSubscribeSvcClient is a mock of SubscribeSvcClient interface.
type MockSubscribeSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribeSvcClientMockRecorder
}

// MockSubscribeSvcClientMockRecorder is the mock recorder for MockSubscribeSvcClient.
type MockSubscribeSvcClientMockRecorder struct {
	mock *MockSubscribeSvcClient
}

// NewMockSubscribeSvcClient creates a new mock instance.
func NewMockSubscribeSvcClient(ctrl *gomock.Controller) *MockSubscribeSvcClient {
	mock := &MockSubscribeSvcClient{ctrl: ctrl}
	mock.recorder = &MockSubscribeSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribeSvcClient) EXPECT() *MockSubscribeSvcClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSubscribeSvcClient) Close(ctx context.Context, in *pb.SubscribeRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Close", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockSubscribeSvcClientMockRecorder) Close(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSubscribeSvcClient)(nil).Close), varargs...)
}

// List mocks base method.
func (m *MockSubscribeSvcClient) List(ctx context.Context, in *pb.SubscribeRequest, opts ...grpc.CallOption) (*pb.SubscribeReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*pb.SubscribeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSubscribeSvcClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSubscribeSvcClient)(nil).List), varargs...)
}

// Open mocks base method.
func (m *MockSubscribeSvcClient) Open(ctx context.Context, in *pb.SubscribeRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Open", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockSubscribeSvcClientMockRecorder) Open(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockSubscribeSvcClient)(nil).Open), varargs...)
}

// Register mocks base method.
func (m *MockSubscribeSvcClient) Register(ctx context.Context, in *pb.SubscribeRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockSubscribeSvcClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSubscribeSvcClient)(nil).Register), varargs...)
}

// Status mocks base method.
func (m *MockSubscribeSvcClient) Status(ctx context.Context, in *pb.SubscribeRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockSubscribeSvcClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockSubscribeSvcClient)(nil).Status), varargs...)
}

// MockSubscribeSvcServer is a mock of SubscribeSvcServer interface.
type MockSubscribeSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribeSvcServerMockRecorder
}

// MockSubscribeSvcServerMockRecorder is the mock recorder for MockSubscribeSvcServer.
type MockSubscribeSvcServerMockRecorder struct {
	mock *MockSubscribeSvcServer
}

// NewMockSubscribeSvcServer creates a new mock instance.
func NewMockSubscribeSvcServer(ctrl *gomock.Controller) *MockSubscribeSvcServer {
	mock := &MockSubscribeSvcServer{ctrl: ctrl}
	mock.recorder = &MockSubscribeSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribeSvcServer) EXPECT() *MockSubscribeSvcServerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSubscribeSvcServer) Close(arg0 context.Context, arg1 *pb.SubscribeRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockSubscribeSvcServerMockRecorder) Close(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSubscribeSvcServer)(nil).Close), arg0, arg1)
}

// List mocks base method.
func (m *MockSubscribeSvcServer) List(arg0 context.Context, arg1 *pb.SubscribeRequest) (*pb.SubscribeReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*pb.SubscribeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSubscribeSvcServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSubscribeSvcServer)(nil).List), arg0, arg1)
}

// Open mocks base method.
func (m *MockSubscribeSvcServer) Open(arg0 context.Context, arg1 *pb.SubscribeRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockSubscribeSvcServerMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockSubscribeSvcServer)(nil).Open), arg0, arg1)
}

// Register mocks base method.
func (m *MockSubscribeSvcServer) Register(arg0 context.Context, arg1 *pb.SubscribeRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockSubscribeSvcServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSubscribeSvcServer)(nil).Register), arg0, arg1)
}

// Status mocks base method.
func (m *MockSubscribeSvcServer) Status(arg0 context.Context, arg1 *pb.SubscribeRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockSubscribeSvcServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockSubscribeSvcServer)(nil).Status), arg0, arg1)
}
