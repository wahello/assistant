// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/chatbot.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockChatbotSvcClient is a mock of ChatbotSvcClient interface.
type MockChatbotSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatbotSvcClientMockRecorder
}

// MockChatbotSvcClientMockRecorder is the mock recorder for MockChatbotSvcClient.
type MockChatbotSvcClientMockRecorder struct {
	mock *MockChatbotSvcClient
}

// NewMockChatbotSvcClient creates a new mock instance.
func NewMockChatbotSvcClient(ctrl *gomock.Controller) *MockChatbotSvcClient {
	mock := &MockChatbotSvcClient{ctrl: ctrl}
	mock.recorder = &MockChatbotSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatbotSvcClient) EXPECT() *MockChatbotSvcClientMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockChatbotSvcClient) Handle(ctx context.Context, in *pb.ChatbotRequest, opts ...grpc.CallOption) (*pb.ChatbotReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(*pb.ChatbotReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockChatbotSvcClientMockRecorder) Handle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockChatbotSvcClient)(nil).Handle), varargs...)
}

// MockChatbotSvcServer is a mock of ChatbotSvcServer interface.
type MockChatbotSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatbotSvcServerMockRecorder
}

// MockChatbotSvcServerMockRecorder is the mock recorder for MockChatbotSvcServer.
type MockChatbotSvcServerMockRecorder struct {
	mock *MockChatbotSvcServer
}

// NewMockChatbotSvcServer creates a new mock instance.
func NewMockChatbotSvcServer(ctrl *gomock.Controller) *MockChatbotSvcServer {
	mock := &MockChatbotSvcServer{ctrl: ctrl}
	mock.recorder = &MockChatbotSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatbotSvcServer) EXPECT() *MockChatbotSvcServerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockChatbotSvcServer) Handle(arg0 context.Context, arg1 *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(*pb.ChatbotReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockChatbotSvcServerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockChatbotSvcServer)(nil).Handle), arg0, arg1)
}
