// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/okr_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockOkrSvcClient is a mock of OkrSvcClient interface.
type MockOkrSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockOkrSvcClientMockRecorder
}

// MockOkrSvcClientMockRecorder is the mock recorder for MockOkrSvcClient.
type MockOkrSvcClientMockRecorder struct {
	mock *MockOkrSvcClient
}

// NewMockOkrSvcClient creates a new mock instance.
func NewMockOkrSvcClient(ctrl *gomock.Controller) *MockOkrSvcClient {
	mock := &MockOkrSvcClient{ctrl: ctrl}
	mock.recorder = &MockOkrSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOkrSvcClient) EXPECT() *MockOkrSvcClientMockRecorder {
	return m.recorder
}

// CreateKeyResult mocks base method.
func (m *MockOkrSvcClient) CreateKeyResult(ctx context.Context, in *pb.KeyResultRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateKeyResult", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKeyResult indicates an expected call of CreateKeyResult.
func (mr *MockOkrSvcClientMockRecorder) CreateKeyResult(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeyResult", reflect.TypeOf((*MockOkrSvcClient)(nil).CreateKeyResult), varargs...)
}

// CreateKeyResultValue mocks base method.
func (m *MockOkrSvcClient) CreateKeyResultValue(ctx context.Context, in *pb.KeyResultValueRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateKeyResultValue", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKeyResultValue indicates an expected call of CreateKeyResultValue.
func (mr *MockOkrSvcClientMockRecorder) CreateKeyResultValue(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeyResultValue", reflect.TypeOf((*MockOkrSvcClient)(nil).CreateKeyResultValue), varargs...)
}

// CreateObjective mocks base method.
func (m *MockOkrSvcClient) CreateObjective(ctx context.Context, in *pb.ObjectiveRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateObjective", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjective indicates an expected call of CreateObjective.
func (mr *MockOkrSvcClientMockRecorder) CreateObjective(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjective", reflect.TypeOf((*MockOkrSvcClient)(nil).CreateObjective), varargs...)
}

// DeleteKeyResult mocks base method.
func (m *MockOkrSvcClient) DeleteKeyResult(ctx context.Context, in *pb.KeyResultRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteKeyResult", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKeyResult indicates an expected call of DeleteKeyResult.
func (mr *MockOkrSvcClientMockRecorder) DeleteKeyResult(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKeyResult", reflect.TypeOf((*MockOkrSvcClient)(nil).DeleteKeyResult), varargs...)
}

// DeleteObjective mocks base method.
func (m *MockOkrSvcClient) DeleteObjective(ctx context.Context, in *pb.ObjectiveRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjective", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjective indicates an expected call of DeleteObjective.
func (mr *MockOkrSvcClientMockRecorder) DeleteObjective(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjective", reflect.TypeOf((*MockOkrSvcClient)(nil).DeleteObjective), varargs...)
}

// GetKeyResult mocks base method.
func (m *MockOkrSvcClient) GetKeyResult(ctx context.Context, in *pb.KeyResultRequest, opts ...grpc.CallOption) (*pb.KeyResultReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyResult", varargs...)
	ret0, _ := ret[0].(*pb.KeyResultReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyResult indicates an expected call of GetKeyResult.
func (mr *MockOkrSvcClientMockRecorder) GetKeyResult(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyResult", reflect.TypeOf((*MockOkrSvcClient)(nil).GetKeyResult), varargs...)
}

// GetKeyResults mocks base method.
func (m *MockOkrSvcClient) GetKeyResults(ctx context.Context, in *pb.KeyResultRequest, opts ...grpc.CallOption) (*pb.KeyResultsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyResults", varargs...)
	ret0, _ := ret[0].(*pb.KeyResultsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyResults indicates an expected call of GetKeyResults.
func (mr *MockOkrSvcClientMockRecorder) GetKeyResults(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyResults", reflect.TypeOf((*MockOkrSvcClient)(nil).GetKeyResults), varargs...)
}

// GetObjective mocks base method.
func (m *MockOkrSvcClient) GetObjective(ctx context.Context, in *pb.ObjectiveRequest, opts ...grpc.CallOption) (*pb.ObjectiveReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjective", varargs...)
	ret0, _ := ret[0].(*pb.ObjectiveReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjective indicates an expected call of GetObjective.
func (mr *MockOkrSvcClientMockRecorder) GetObjective(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjective", reflect.TypeOf((*MockOkrSvcClient)(nil).GetObjective), varargs...)
}

// GetObjectives mocks base method.
func (m *MockOkrSvcClient) GetObjectives(ctx context.Context, in *pb.ObjectiveRequest, opts ...grpc.CallOption) (*pb.ObjectivesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectives", varargs...)
	ret0, _ := ret[0].(*pb.ObjectivesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectives indicates an expected call of GetObjectives.
func (mr *MockOkrSvcClientMockRecorder) GetObjectives(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectives", reflect.TypeOf((*MockOkrSvcClient)(nil).GetObjectives), varargs...)
}

// MockOkrSvcServer is a mock of OkrSvcServer interface.
type MockOkrSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockOkrSvcServerMockRecorder
}

// MockOkrSvcServerMockRecorder is the mock recorder for MockOkrSvcServer.
type MockOkrSvcServerMockRecorder struct {
	mock *MockOkrSvcServer
}

// NewMockOkrSvcServer creates a new mock instance.
func NewMockOkrSvcServer(ctrl *gomock.Controller) *MockOkrSvcServer {
	mock := &MockOkrSvcServer{ctrl: ctrl}
	mock.recorder = &MockOkrSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOkrSvcServer) EXPECT() *MockOkrSvcServerMockRecorder {
	return m.recorder
}

// CreateKeyResult mocks base method.
func (m *MockOkrSvcServer) CreateKeyResult(arg0 context.Context, arg1 *pb.KeyResultRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKeyResult", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKeyResult indicates an expected call of CreateKeyResult.
func (mr *MockOkrSvcServerMockRecorder) CreateKeyResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeyResult", reflect.TypeOf((*MockOkrSvcServer)(nil).CreateKeyResult), arg0, arg1)
}

// CreateKeyResultValue mocks base method.
func (m *MockOkrSvcServer) CreateKeyResultValue(arg0 context.Context, arg1 *pb.KeyResultValueRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKeyResultValue", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKeyResultValue indicates an expected call of CreateKeyResultValue.
func (mr *MockOkrSvcServerMockRecorder) CreateKeyResultValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeyResultValue", reflect.TypeOf((*MockOkrSvcServer)(nil).CreateKeyResultValue), arg0, arg1)
}

// CreateObjective mocks base method.
func (m *MockOkrSvcServer) CreateObjective(arg0 context.Context, arg1 *pb.ObjectiveRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObjective", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjective indicates an expected call of CreateObjective.
func (mr *MockOkrSvcServerMockRecorder) CreateObjective(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjective", reflect.TypeOf((*MockOkrSvcServer)(nil).CreateObjective), arg0, arg1)
}

// DeleteKeyResult mocks base method.
func (m *MockOkrSvcServer) DeleteKeyResult(arg0 context.Context, arg1 *pb.KeyResultRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKeyResult", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKeyResult indicates an expected call of DeleteKeyResult.
func (mr *MockOkrSvcServerMockRecorder) DeleteKeyResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKeyResult", reflect.TypeOf((*MockOkrSvcServer)(nil).DeleteKeyResult), arg0, arg1)
}

// DeleteObjective mocks base method.
func (m *MockOkrSvcServer) DeleteObjective(arg0 context.Context, arg1 *pb.ObjectiveRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjective", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjective indicates an expected call of DeleteObjective.
func (mr *MockOkrSvcServerMockRecorder) DeleteObjective(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjective", reflect.TypeOf((*MockOkrSvcServer)(nil).DeleteObjective), arg0, arg1)
}

// GetKeyResult mocks base method.
func (m *MockOkrSvcServer) GetKeyResult(arg0 context.Context, arg1 *pb.KeyResultRequest) (*pb.KeyResultReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyResult", arg0, arg1)
	ret0, _ := ret[0].(*pb.KeyResultReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyResult indicates an expected call of GetKeyResult.
func (mr *MockOkrSvcServerMockRecorder) GetKeyResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyResult", reflect.TypeOf((*MockOkrSvcServer)(nil).GetKeyResult), arg0, arg1)
}

// GetKeyResults mocks base method.
func (m *MockOkrSvcServer) GetKeyResults(arg0 context.Context, arg1 *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyResults", arg0, arg1)
	ret0, _ := ret[0].(*pb.KeyResultsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyResults indicates an expected call of GetKeyResults.
func (mr *MockOkrSvcServerMockRecorder) GetKeyResults(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyResults", reflect.TypeOf((*MockOkrSvcServer)(nil).GetKeyResults), arg0, arg1)
}

// GetObjective mocks base method.
func (m *MockOkrSvcServer) GetObjective(arg0 context.Context, arg1 *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjective", arg0, arg1)
	ret0, _ := ret[0].(*pb.ObjectiveReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjective indicates an expected call of GetObjective.
func (mr *MockOkrSvcServerMockRecorder) GetObjective(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjective", reflect.TypeOf((*MockOkrSvcServer)(nil).GetObjective), arg0, arg1)
}

// GetObjectives mocks base method.
func (m *MockOkrSvcServer) GetObjectives(arg0 context.Context, arg1 *pb.ObjectiveRequest) (*pb.ObjectivesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectives", arg0, arg1)
	ret0, _ := ret[0].(*pb.ObjectivesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectives indicates an expected call of GetObjectives.
func (mr *MockOkrSvcServerMockRecorder) GetObjectives(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectives", reflect.TypeOf((*MockOkrSvcServer)(nil).GetObjectives), arg0, arg1)
}

// MockUnsafeOkrSvcServer is a mock of UnsafeOkrSvcServer interface.
type MockUnsafeOkrSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeOkrSvcServerMockRecorder
}

// MockUnsafeOkrSvcServerMockRecorder is the mock recorder for MockUnsafeOkrSvcServer.
type MockUnsafeOkrSvcServerMockRecorder struct {
	mock *MockUnsafeOkrSvcServer
}

// NewMockUnsafeOkrSvcServer creates a new mock instance.
func NewMockUnsafeOkrSvcServer(ctrl *gomock.Controller) *MockUnsafeOkrSvcServer {
	mock := &MockUnsafeOkrSvcServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeOkrSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeOkrSvcServer) EXPECT() *MockUnsafeOkrSvcServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedOkrSvcServer mocks base method.
func (m *MockUnsafeOkrSvcServer) mustEmbedUnimplementedOkrSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedOkrSvcServer")
}

// mustEmbedUnimplementedOkrSvcServer indicates an expected call of mustEmbedUnimplementedOkrSvcServer.
func (mr *MockUnsafeOkrSvcServerMockRecorder) mustEmbedUnimplementedOkrSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedOkrSvcServer", reflect.TypeOf((*MockUnsafeOkrSvcServer)(nil).mustEmbedUnimplementedOkrSvcServer))
}