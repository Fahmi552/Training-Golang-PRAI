// Code generated by MockGen. DO NOT EDIT.
// Source: session-7-db-pg-gorm/service/user_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	entity "Training/session-10-crud-user-grpcgateway/entity"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserServiceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserService)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockIUserService) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserServiceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserService)(nil).DeleteUser), ctx, id)
}

// GetAllUsers mocks base method.
func (m *MockIUserService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserServiceMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserService)(nil).GetAllUsers), ctx)
}

// GetUserByEmail mocks base method.
func (m *MockIUserService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIUserServiceMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIUserService)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockIUserService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockIUserServiceMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockIUserService)(nil).GetUserByID), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockIUserService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, user)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserServiceMockRecorder) UpdateUser(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserService)(nil).UpdateUser), ctx, id, user)
}

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepository)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockIUserRepository) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserRepositoryMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserRepository)(nil).DeleteUser), ctx, id)
}

// GetAllUsers mocks base method.
func (m *MockIUserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserRepositoryMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetAllUsers), ctx)
}

// GetUserByEmail mocks base method.
func (m *MockIUserRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIUserRepositoryMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockIUserRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockIUserRepositoryMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByID), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockIUserRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, user)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserRepositoryMockRecorder) UpdateUser(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepository)(nil).UpdateUser), ctx, id, user)
}
