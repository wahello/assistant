// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/middle/repository/middle.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
)

// MockMiddleRepository is a mock of MiddleRepository interface.
type MockMiddleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMiddleRepositoryMockRecorder
}

// MockMiddleRepositoryMockRecorder is the mock recorder for MockMiddleRepository.
type MockMiddleRepositoryMockRecorder struct {
	mock *MockMiddleRepository
}

// NewMockMiddleRepository creates a new mock instance.
func NewMockMiddleRepository(ctrl *gomock.Controller) *MockMiddleRepository {
	mock := &MockMiddleRepository{ctrl: ctrl}
	mock.recorder = &MockMiddleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleRepository) EXPECT() *MockMiddleRepositoryMockRecorder {
	return m.recorder
}

// CollectMetadata mocks base method.
func (m *MockMiddleRepository) CollectMetadata(ctx context.Context, model []interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectMetadata", ctx, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// CollectMetadata indicates an expected call of CollectMetadata.
func (mr *MockMiddleRepositoryMockRecorder) CollectMetadata(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectMetadata", reflect.TypeOf((*MockMiddleRepository)(nil).CollectMetadata), ctx, model)
}

// CreateApp mocks base method.
func (m *MockMiddleRepository) CreateApp(ctx context.Context, app *pb.App) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApp", ctx, app)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApp indicates an expected call of CreateApp.
func (mr *MockMiddleRepositoryMockRecorder) CreateApp(ctx, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApp", reflect.TypeOf((*MockMiddleRepository)(nil).CreateApp), ctx, app)
}

// CreateCounter mocks base method.
func (m *MockMiddleRepository) CreateCounter(ctx context.Context, counter *pb.Counter) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCounter", ctx, counter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCounter indicates an expected call of CreateCounter.
func (mr *MockMiddleRepositoryMockRecorder) CreateCounter(ctx, counter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCounter", reflect.TypeOf((*MockMiddleRepository)(nil).CreateCounter), ctx, counter)
}

// CreateCredential mocks base method.
func (m *MockMiddleRepository) CreateCredential(ctx context.Context, credential *pb.Credential) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCredential", ctx, credential)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCredential indicates an expected call of CreateCredential.
func (mr *MockMiddleRepositoryMockRecorder) CreateCredential(ctx, credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCredential", reflect.TypeOf((*MockMiddleRepository)(nil).CreateCredential), ctx, credential)
}

// CreatePage mocks base method.
func (m *MockMiddleRepository) CreatePage(ctx context.Context, page *pb.Page) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePage", ctx, page)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePage indicates an expected call of CreatePage.
func (mr *MockMiddleRepositoryMockRecorder) CreatePage(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePage", reflect.TypeOf((*MockMiddleRepository)(nil).CreatePage), ctx, page)
}

// CreateSubscribe mocks base method.
func (m *MockMiddleRepository) CreateSubscribe(ctx context.Context, subscribe pb.Subscribe) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscribe", ctx, subscribe)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscribe indicates an expected call of CreateSubscribe.
func (mr *MockMiddleRepositoryMockRecorder) CreateSubscribe(ctx, subscribe interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscribe", reflect.TypeOf((*MockMiddleRepository)(nil).CreateSubscribe), ctx, subscribe)
}

// CreateUserSubscribe mocks base method.
func (m *MockMiddleRepository) CreateUserSubscribe(ctx context.Context, subscribe pb.UserSubscribe) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSubscribe", ctx, subscribe)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserSubscribe indicates an expected call of CreateUserSubscribe.
func (mr *MockMiddleRepositoryMockRecorder) CreateUserSubscribe(ctx, subscribe interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSubscribe", reflect.TypeOf((*MockMiddleRepository)(nil).CreateUserSubscribe), ctx, subscribe)
}

// DecreaseCounter mocks base method.
func (m *MockMiddleRepository) DecreaseCounter(ctx context.Context, id, amount int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseCounter", ctx, id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecreaseCounter indicates an expected call of DecreaseCounter.
func (mr *MockMiddleRepositoryMockRecorder) DecreaseCounter(ctx, id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseCounter", reflect.TypeOf((*MockMiddleRepository)(nil).DecreaseCounter), ctx, id, amount)
}

// GetAppByType mocks base method.
func (m *MockMiddleRepository) GetAppByType(ctx context.Context, userId int64, t string) (pb.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppByType", ctx, userId, t)
	ret0, _ := ret[0].(pb.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppByType indicates an expected call of GetAppByType.
func (mr *MockMiddleRepositoryMockRecorder) GetAppByType(ctx, userId, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppByType", reflect.TypeOf((*MockMiddleRepository)(nil).GetAppByType), ctx, userId, t)
}

// GetAvailableAppByType mocks base method.
func (m *MockMiddleRepository) GetAvailableAppByType(ctx context.Context, userId int64, t string) (pb.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableAppByType", ctx, userId, t)
	ret0, _ := ret[0].(pb.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableAppByType indicates an expected call of GetAvailableAppByType.
func (mr *MockMiddleRepositoryMockRecorder) GetAvailableAppByType(ctx, userId, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableAppByType", reflect.TypeOf((*MockMiddleRepository)(nil).GetAvailableAppByType), ctx, userId, t)
}

// GetCounter mocks base method.
func (m *MockMiddleRepository) GetCounter(ctx context.Context, id int64) (pb.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounter", ctx, id)
	ret0, _ := ret[0].(pb.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounter indicates an expected call of GetCounter.
func (mr *MockMiddleRepositoryMockRecorder) GetCounter(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounter", reflect.TypeOf((*MockMiddleRepository)(nil).GetCounter), ctx, id)
}

// GetCounterByFlag mocks base method.
func (m *MockMiddleRepository) GetCounterByFlag(ctx context.Context, userId int64, flag string) (pb.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounterByFlag", ctx, userId, flag)
	ret0, _ := ret[0].(pb.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounterByFlag indicates an expected call of GetCounterByFlag.
func (mr *MockMiddleRepositoryMockRecorder) GetCounterByFlag(ctx, userId, flag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounterByFlag", reflect.TypeOf((*MockMiddleRepository)(nil).GetCounterByFlag), ctx, userId, flag)
}

// GetCredentialByName mocks base method.
func (m *MockMiddleRepository) GetCredentialByName(ctx context.Context, userId int64, name string) (pb.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialByName", ctx, userId, name)
	ret0, _ := ret[0].(pb.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialByName indicates an expected call of GetCredentialByName.
func (mr *MockMiddleRepositoryMockRecorder) GetCredentialByName(ctx, userId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialByName", reflect.TypeOf((*MockMiddleRepository)(nil).GetCredentialByName), ctx, userId, name)
}

// GetCredentialByType mocks base method.
func (m *MockMiddleRepository) GetCredentialByType(ctx context.Context, userId int64, t string) (pb.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialByType", ctx, userId, t)
	ret0, _ := ret[0].(pb.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialByType indicates an expected call of GetCredentialByType.
func (mr *MockMiddleRepositoryMockRecorder) GetCredentialByType(ctx, userId, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialByType", reflect.TypeOf((*MockMiddleRepository)(nil).GetCredentialByType), ctx, userId, t)
}

// GetOrCreateModelTag mocks base method.
func (m *MockMiddleRepository) GetOrCreateModelTag(ctx context.Context, tag *pb.ModelTag) (pb.ModelTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateModelTag", ctx, tag)
	ret0, _ := ret[0].(pb.ModelTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreateModelTag indicates an expected call of GetOrCreateModelTag.
func (mr *MockMiddleRepositoryMockRecorder) GetOrCreateModelTag(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateModelTag", reflect.TypeOf((*MockMiddleRepository)(nil).GetOrCreateModelTag), ctx, tag)
}

// GetOrCreateTag mocks base method.
func (m *MockMiddleRepository) GetOrCreateTag(ctx context.Context, tag *pb.Tag) (pb.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateTag", ctx, tag)
	ret0, _ := ret[0].(pb.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreateTag indicates an expected call of GetOrCreateTag.
func (mr *MockMiddleRepositoryMockRecorder) GetOrCreateTag(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateTag", reflect.TypeOf((*MockMiddleRepository)(nil).GetOrCreateTag), ctx, tag)
}

// GetPageByUUID mocks base method.
func (m *MockMiddleRepository) GetPageByUUID(ctx context.Context, uuid string) (pb.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPageByUUID", ctx, uuid)
	ret0, _ := ret[0].(pb.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPageByUUID indicates an expected call of GetPageByUUID.
func (mr *MockMiddleRepositoryMockRecorder) GetPageByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPageByUUID", reflect.TypeOf((*MockMiddleRepository)(nil).GetPageByUUID), ctx, uuid)
}

// GetSubscribe mocks base method.
func (m *MockMiddleRepository) GetSubscribe(ctx context.Context, name string) (pb.Subscribe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribe", ctx, name)
	ret0, _ := ret[0].(pb.Subscribe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribe indicates an expected call of GetSubscribe.
func (mr *MockMiddleRepositoryMockRecorder) GetSubscribe(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribe", reflect.TypeOf((*MockMiddleRepository)(nil).GetSubscribe), ctx, name)
}

// GetUserSubscribe mocks base method.
func (m *MockMiddleRepository) GetUserSubscribe(ctx context.Context, userId, subscribeId int64) (pb.UserSubscribe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSubscribe", ctx, userId, subscribeId)
	ret0, _ := ret[0].(pb.UserSubscribe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscribe indicates an expected call of GetUserSubscribe.
func (mr *MockMiddleRepositoryMockRecorder) GetUserSubscribe(ctx, userId, subscribeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSubscribe", reflect.TypeOf((*MockMiddleRepository)(nil).GetUserSubscribe), ctx, userId, subscribeId)
}

// IncreaseCounter mocks base method.
func (m *MockMiddleRepository) IncreaseCounter(ctx context.Context, id, amount int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseCounter", ctx, id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseCounter indicates an expected call of IncreaseCounter.
func (mr *MockMiddleRepositoryMockRecorder) IncreaseCounter(ctx, id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseCounter", reflect.TypeOf((*MockMiddleRepository)(nil).IncreaseCounter), ctx, id, amount)
}

// ListApps mocks base method.
func (m *MockMiddleRepository) ListApps(ctx context.Context, userId int64) ([]*pb.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApps", ctx, userId)
	ret0, _ := ret[0].([]*pb.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApps indicates an expected call of ListApps.
func (mr *MockMiddleRepositoryMockRecorder) ListApps(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApps", reflect.TypeOf((*MockMiddleRepository)(nil).ListApps), ctx, userId)
}

// ListCounter mocks base method.
func (m *MockMiddleRepository) ListCounter(ctx context.Context, userId int64) ([]*pb.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCounter", ctx, userId)
	ret0, _ := ret[0].([]*pb.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCounter indicates an expected call of ListCounter.
func (mr *MockMiddleRepositoryMockRecorder) ListCounter(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCounter", reflect.TypeOf((*MockMiddleRepository)(nil).ListCounter), ctx, userId)
}

// ListCredentials mocks base method.
func (m *MockMiddleRepository) ListCredentials(ctx context.Context, userId int64) ([]*pb.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCredentials", ctx, userId)
	ret0, _ := ret[0].([]*pb.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCredentials indicates an expected call of ListCredentials.
func (mr *MockMiddleRepositoryMockRecorder) ListCredentials(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCredentials", reflect.TypeOf((*MockMiddleRepository)(nil).ListCredentials), ctx, userId)
}

// ListModelTagsByModel mocks base method.
func (m *MockMiddleRepository) ListModelTagsByModel(ctx context.Context, userId int64, model pb.ModelTag) ([]*pb.ModelTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelTagsByModel", ctx, userId, model)
	ret0, _ := ret[0].([]*pb.ModelTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModelTagsByModel indicates an expected call of ListModelTagsByModel.
func (mr *MockMiddleRepositoryMockRecorder) ListModelTagsByModel(ctx, userId, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelTagsByModel", reflect.TypeOf((*MockMiddleRepository)(nil).ListModelTagsByModel), ctx, userId, model)
}

// ListModelTagsByModelId mocks base method.
func (m *MockMiddleRepository) ListModelTagsByModelId(ctx context.Context, userId int64, modelId []int64) ([]*pb.ModelTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelTagsByModelId", ctx, userId, modelId)
	ret0, _ := ret[0].([]*pb.ModelTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModelTagsByModelId indicates an expected call of ListModelTagsByModelId.
func (mr *MockMiddleRepositoryMockRecorder) ListModelTagsByModelId(ctx, userId, modelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelTagsByModelId", reflect.TypeOf((*MockMiddleRepository)(nil).ListModelTagsByModelId), ctx, userId, modelId)
}

// ListModelTagsByTag mocks base method.
func (m *MockMiddleRepository) ListModelTagsByTag(ctx context.Context, userId int64, tag string) ([]*pb.ModelTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelTagsByTag", ctx, userId, tag)
	ret0, _ := ret[0].([]*pb.ModelTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModelTagsByTag indicates an expected call of ListModelTagsByTag.
func (mr *MockMiddleRepositoryMockRecorder) ListModelTagsByTag(ctx, userId, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelTagsByTag", reflect.TypeOf((*MockMiddleRepository)(nil).ListModelTagsByTag), ctx, userId, tag)
}

// ListSubscribe mocks base method.
func (m *MockMiddleRepository) ListSubscribe(ctx context.Context) ([]*pb.Subscribe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubscribe", ctx)
	ret0, _ := ret[0].([]*pb.Subscribe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscribe indicates an expected call of ListSubscribe.
func (mr *MockMiddleRepositoryMockRecorder) ListSubscribe(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscribe", reflect.TypeOf((*MockMiddleRepository)(nil).ListSubscribe), ctx)
}

// ListTags mocks base method.
func (m *MockMiddleRepository) ListTags(ctx context.Context, userId int64) ([]*pb.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", ctx, userId)
	ret0, _ := ret[0].([]*pb.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockMiddleRepositoryMockRecorder) ListTags(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockMiddleRepository)(nil).ListTags), ctx, userId)
}

// ListUserSubscribe mocks base method.
func (m *MockMiddleRepository) ListUserSubscribe(ctx context.Context, userId int64) ([]*pb.KV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserSubscribe", ctx, userId)
	ret0, _ := ret[0].([]*pb.KV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserSubscribe indicates an expected call of ListUserSubscribe.
func (mr *MockMiddleRepositoryMockRecorder) ListUserSubscribe(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserSubscribe", reflect.TypeOf((*MockMiddleRepository)(nil).ListUserSubscribe), ctx, userId)
}

// Search mocks base method.
func (m *MockMiddleRepository) Search(ctx context.Context, userId int64, filter [][]string) ([]*pb.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, userId, filter)
	ret0, _ := ret[0].([]*pb.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockMiddleRepositoryMockRecorder) Search(ctx, userId, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockMiddleRepository)(nil).Search), ctx, userId, filter)
}

// UpdateAppByID mocks base method.
func (m *MockMiddleRepository) UpdateAppByID(ctx context.Context, id int64, token, extra string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppByID", ctx, id, token, extra)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAppByID indicates an expected call of UpdateAppByID.
func (mr *MockMiddleRepositoryMockRecorder) UpdateAppByID(ctx, id, token, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppByID", reflect.TypeOf((*MockMiddleRepository)(nil).UpdateAppByID), ctx, id, token, extra)
}

// UpdateSubscribeStatus mocks base method.
func (m *MockMiddleRepository) UpdateSubscribeStatus(ctx context.Context, name string, status int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscribeStatus", ctx, name, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscribeStatus indicates an expected call of UpdateSubscribeStatus.
func (mr *MockMiddleRepositoryMockRecorder) UpdateSubscribeStatus(ctx, name, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscribeStatus", reflect.TypeOf((*MockMiddleRepository)(nil).UpdateSubscribeStatus), ctx, name, status)
}

// UpdateUserSubscribeStatus mocks base method.
func (m *MockMiddleRepository) UpdateUserSubscribeStatus(ctx context.Context, userId, subscribeId, status int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserSubscribeStatus", ctx, userId, subscribeId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserSubscribeStatus indicates an expected call of UpdateUserSubscribeStatus.
func (mr *MockMiddleRepositoryMockRecorder) UpdateUserSubscribeStatus(ctx, userId, subscribeId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSubscribeStatus", reflect.TypeOf((*MockMiddleRepository)(nil).UpdateUserSubscribeStatus), ctx, userId, subscribeId, status)
}
