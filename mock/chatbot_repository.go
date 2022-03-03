// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/chatbot/repository/chatbot.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
)

// MockChatbotRepository is a mock of ChatbotRepository interface.
type MockChatbotRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatbotRepositoryMockRecorder
}

// MockChatbotRepositoryMockRecorder is the mock recorder for MockChatbotRepository.
type MockChatbotRepositoryMockRecorder struct {
	mock *MockChatbotRepository
}

// NewMockChatbotRepository creates a new mock instance.
func NewMockChatbotRepository(ctrl *gomock.Controller) *MockChatbotRepository {
	mock := &MockChatbotRepository{ctrl: ctrl}
	mock.recorder = &MockChatbotRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatbotRepository) EXPECT() *MockChatbotRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChatbotRepository) Create(ctx context.Context, bot *pb.Bot) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, bot)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockChatbotRepositoryMockRecorder) Create(ctx, bot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChatbotRepository)(nil).Create), ctx, bot)
}

// CreateGroup mocks base method.
func (m *MockChatbotRepository) CreateGroup(ctx context.Context, group *pb.Group) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, group)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockChatbotRepositoryMockRecorder) CreateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockChatbotRepository)(nil).CreateGroup), ctx, group)
}

// CreateGroupBot mocks base method.
func (m *MockChatbotRepository) CreateGroupBot(ctx context.Context, groupId int64, bot *pb.Bot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupBot", ctx, groupId, bot)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGroupBot indicates an expected call of CreateGroupBot.
func (mr *MockChatbotRepositoryMockRecorder) CreateGroupBot(ctx, groupId, bot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupBot", reflect.TypeOf((*MockChatbotRepository)(nil).CreateGroupBot), ctx, groupId, bot)
}

// CreateGroupTag mocks base method.
func (m *MockChatbotRepository) CreateGroupTag(ctx context.Context, tag *pb.GroupTag) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupTag", ctx, tag)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroupTag indicates an expected call of CreateGroupTag.
func (mr *MockChatbotRepositoryMockRecorder) CreateGroupTag(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupTag", reflect.TypeOf((*MockChatbotRepository)(nil).CreateGroupTag), ctx, tag)
}

// Delete mocks base method.
func (m *MockChatbotRepository) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChatbotRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChatbotRepository)(nil).Delete), ctx, id)
}

// DeleteGroup mocks base method.
func (m *MockChatbotRepository) DeleteGroup(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup.
func (mr *MockChatbotRepositoryMockRecorder) DeleteGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*MockChatbotRepository)(nil).DeleteGroup), ctx, id)
}

// DeleteGroupBot mocks base method.
func (m *MockChatbotRepository) DeleteGroupBot(ctx context.Context, groupId, botId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupBot", ctx, groupId, botId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupBot indicates an expected call of DeleteGroupBot.
func (mr *MockChatbotRepositoryMockRecorder) DeleteGroupBot(ctx, groupId, botId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupBot", reflect.TypeOf((*MockChatbotRepository)(nil).DeleteGroupBot), ctx, groupId, botId)
}

// DeleteGroupTag mocks base method.
func (m *MockChatbotRepository) DeleteGroupTag(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupTag", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupTag indicates an expected call of DeleteGroupTag.
func (mr *MockChatbotRepositoryMockRecorder) DeleteGroupTag(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupTag", reflect.TypeOf((*MockChatbotRepository)(nil).DeleteGroupTag), ctx, id)
}

// GetBotsByText mocks base method.
func (m *MockChatbotRepository) GetBotsByText(ctx context.Context, text []string) (map[string]*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBotsByText", ctx, text)
	ret0, _ := ret[0].(map[string]*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBotsByText indicates an expected call of GetBotsByText.
func (mr *MockChatbotRepositoryMockRecorder) GetBotsByText(ctx, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBotsByText", reflect.TypeOf((*MockChatbotRepository)(nil).GetBotsByText), ctx, text)
}

// GetBotsByUser mocks base method.
func (m *MockChatbotRepository) GetBotsByUser(ctx context.Context, userId int64) ([]*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBotsByUser", ctx, userId)
	ret0, _ := ret[0].([]*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBotsByUser indicates an expected call of GetBotsByUser.
func (mr *MockChatbotRepositoryMockRecorder) GetBotsByUser(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBotsByUser", reflect.TypeOf((*MockChatbotRepository)(nil).GetBotsByUser), ctx, userId)
}

// GetByID mocks base method.
func (m *MockChatbotRepository) GetByID(ctx context.Context, id int64) (pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockChatbotRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockChatbotRepository)(nil).GetByID), ctx, id)
}

// GetByIdentifier mocks base method.
func (m *MockChatbotRepository) GetByIdentifier(ctx context.Context, identifier string) (pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdentifier", ctx, identifier)
	ret0, _ := ret[0].(pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdentifier indicates an expected call of GetByIdentifier.
func (mr *MockChatbotRepositoryMockRecorder) GetByIdentifier(ctx, identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdentifier", reflect.TypeOf((*MockChatbotRepository)(nil).GetByIdentifier), ctx, identifier)
}

// GetByUUID mocks base method.
func (m *MockChatbotRepository) GetByUUID(ctx context.Context, uuid string) (pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", ctx, uuid)
	ret0, _ := ret[0].(pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUUID indicates an expected call of GetByUUID.
func (mr *MockChatbotRepositoryMockRecorder) GetByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockChatbotRepository)(nil).GetByUUID), ctx, uuid)
}

// GetGroup mocks base method.
func (m *MockChatbotRepository) GetGroup(ctx context.Context, id int64) (pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, id)
	ret0, _ := ret[0].(pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockChatbotRepositoryMockRecorder) GetGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroup), ctx, id)
}

// GetGroupBotSetting mocks base method.
func (m *MockChatbotRepository) GetGroupBotSetting(ctx context.Context, groupId, botId int64) ([]*pb.KV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupBotSetting", ctx, groupId, botId)
	ret0, _ := ret[0].([]*pb.KV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupBotSetting indicates an expected call of GetGroupBotSetting.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupBotSetting(ctx, groupId, botId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupBotSetting", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupBotSetting), ctx, groupId, botId)
}

// GetGroupByName mocks base method.
func (m *MockChatbotRepository) GetGroupByName(ctx context.Context, userId int64, name string) (pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupByName", ctx, userId, name)
	ret0, _ := ret[0].(pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupByName indicates an expected call of GetGroupByName.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupByName(ctx, userId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupByName", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupByName), ctx, userId, name)
}

// GetGroupBySequence mocks base method.
func (m *MockChatbotRepository) GetGroupBySequence(ctx context.Context, userId, sequence int64) (pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupBySequence", ctx, userId, sequence)
	ret0, _ := ret[0].(pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupBySequence indicates an expected call of GetGroupBySequence.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupBySequence(ctx, userId, sequence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupBySequence", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupBySequence), ctx, userId, sequence)
}

// GetGroupByUUID mocks base method.
func (m *MockChatbotRepository) GetGroupByUUID(ctx context.Context, uuid string) (pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupByUUID", ctx, uuid)
	ret0, _ := ret[0].(pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupByUUID indicates an expected call of GetGroupByUUID.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupByUUID", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupByUUID), ctx, uuid)
}

// GetGroupSetting mocks base method.
func (m *MockChatbotRepository) GetGroupSetting(ctx context.Context, groupId int64) ([]*pb.KV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupSetting", ctx, groupId)
	ret0, _ := ret[0].([]*pb.KV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupSetting indicates an expected call of GetGroupSetting.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupSetting(ctx, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupSetting", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupSetting), ctx, groupId)
}

// List mocks base method.
func (m *MockChatbotRepository) List(ctx context.Context) ([]*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockChatbotRepositoryMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockChatbotRepository)(nil).List), ctx)
}

// ListGroup mocks base method.
func (m *MockChatbotRepository) ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroup", ctx, userId)
	ret0, _ := ret[0].([]*pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroup indicates an expected call of ListGroup.
func (mr *MockChatbotRepositoryMockRecorder) ListGroup(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroup", reflect.TypeOf((*MockChatbotRepository)(nil).ListGroup), ctx, userId)
}

// ListGroupBot mocks base method.
func (m *MockChatbotRepository) ListGroupBot(ctx context.Context, groupId int64) ([]*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupBot", ctx, groupId)
	ret0, _ := ret[0].([]*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupBot indicates an expected call of ListGroupBot.
func (mr *MockChatbotRepositoryMockRecorder) ListGroupBot(ctx, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupBot", reflect.TypeOf((*MockChatbotRepository)(nil).ListGroupBot), ctx, groupId)
}

// ListGroupTag mocks base method.
func (m *MockChatbotRepository) ListGroupTag(ctx context.Context, groupId int64) ([]*pb.GroupTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupTag", ctx, groupId)
	ret0, _ := ret[0].([]*pb.GroupTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupTag indicates an expected call of ListGroupTag.
func (mr *MockChatbotRepositoryMockRecorder) ListGroupTag(ctx, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupTag", reflect.TypeOf((*MockChatbotRepository)(nil).ListGroupTag), ctx, groupId)
}

// Update mocks base method.
func (m *MockChatbotRepository) Update(ctx context.Context, bot *pb.Bot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bot)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockChatbotRepositoryMockRecorder) Update(ctx, bot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChatbotRepository)(nil).Update), ctx, bot)
}

// UpdateGroup mocks base method.
func (m *MockChatbotRepository) UpdateGroup(ctx context.Context, group *pb.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", ctx, group)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroup indicates an expected call of UpdateGroup.
func (mr *MockChatbotRepositoryMockRecorder) UpdateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*MockChatbotRepository)(nil).UpdateGroup), ctx, group)
}

// UpdateGroupBotSetting mocks base method.
func (m *MockChatbotRepository) UpdateGroupBotSetting(ctx context.Context, groupId, botId int64, kvs []*pb.KV) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupBotSetting", ctx, groupId, botId, kvs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroupBotSetting indicates an expected call of UpdateGroupBotSetting.
func (mr *MockChatbotRepositoryMockRecorder) UpdateGroupBotSetting(ctx, groupId, botId, kvs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupBotSetting", reflect.TypeOf((*MockChatbotRepository)(nil).UpdateGroupBotSetting), ctx, groupId, botId, kvs)
}

// UpdateGroupSetting mocks base method.
func (m *MockChatbotRepository) UpdateGroupSetting(ctx context.Context, groupId int64, kvs []*pb.KV) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupSetting", ctx, groupId, kvs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroupSetting indicates an expected call of UpdateGroupSetting.
func (mr *MockChatbotRepositoryMockRecorder) UpdateGroupSetting(ctx, groupId, kvs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupSetting", reflect.TypeOf((*MockChatbotRepository)(nil).UpdateGroupSetting), ctx, groupId, kvs)
}
