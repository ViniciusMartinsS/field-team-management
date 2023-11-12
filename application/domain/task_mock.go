// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskUsecase is a mock of TaskUsecase interface.
type MockTaskUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUsecaseMockRecorder
}

// MockTaskUsecaseMockRecorder is the mock recorder for MockTaskUsecase.
type MockTaskUsecaseMockRecorder struct {
	mock *MockTaskUsecase
}

// NewMockTaskUsecase creates a new mock instance.
func NewMockTaskUsecase(ctrl *gomock.Controller) *MockTaskUsecase {
	mock := &MockTaskUsecase{ctrl: ctrl}
	mock.recorder = &MockTaskUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUsecase) EXPECT() *MockTaskUsecaseMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTaskUsecase) Add(ctx context.Context, task Task, user User) (Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, task, user)
	ret0, _ := ret[0].(Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTaskUsecaseMockRecorder) Add(ctx, task, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTaskUsecase)(nil).Add), ctx, task, user)
}

// ListByUser mocks base method.
func (m *MockTaskUsecase) ListByUser(ctx context.Context, user User) ([]Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUser", ctx, user)
	ret0, _ := ret[0].([]Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUser indicates an expected call of ListByUser.
func (mr *MockTaskUsecaseMockRecorder) ListByUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUser", reflect.TypeOf((*MockTaskUsecase)(nil).ListByUser), ctx, user)
}

// Remove mocks base method.
func (m *MockTaskUsecase) Remove(ctx context.Context, id int, user User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, id, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTaskUsecaseMockRecorder) Remove(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTaskUsecase)(nil).Remove), ctx, id, user)
}

// Update mocks base method.
func (m *MockTaskUsecase) Update(ctx context.Context, task Task) (Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, task)
	ret0, _ := ret[0].(Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskUsecaseMockRecorder) Update(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskUsecase)(nil).Update), ctx, task)
}

// MockTaskCreator is a mock of TaskCreator interface.
type MockTaskCreator struct {
	ctrl     *gomock.Controller
	recorder *MockTaskCreatorMockRecorder
}

// MockTaskCreatorMockRecorder is the mock recorder for MockTaskCreator.
type MockTaskCreatorMockRecorder struct {
	mock *MockTaskCreator
}

// NewMockTaskCreator creates a new mock instance.
func NewMockTaskCreator(ctrl *gomock.Controller) *MockTaskCreator {
	mock := &MockTaskCreator{ctrl: ctrl}
	mock.recorder = &MockTaskCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskCreator) EXPECT() *MockTaskCreatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTaskCreator) Add(ctx context.Context, task Task) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, task)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTaskCreatorMockRecorder) Add(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTaskCreator)(nil).Add), ctx, task)
}

// MockTaskRetriever is a mock of TaskRetriever interface.
type MockTaskRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRetrieverMockRecorder
}

// MockTaskRetrieverMockRecorder is the mock recorder for MockTaskRetriever.
type MockTaskRetrieverMockRecorder struct {
	mock *MockTaskRetriever
}

// NewMockTaskRetriever creates a new mock instance.
func NewMockTaskRetriever(ctrl *gomock.Controller) *MockTaskRetriever {
	mock := &MockTaskRetriever{ctrl: ctrl}
	mock.recorder = &MockTaskRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRetriever) EXPECT() *MockTaskRetrieverMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockTaskRetriever) List(ctx context.Context) ([]Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTaskRetrieverMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTaskRetriever)(nil).List), ctx)
}

// ListByIDAndUserID mocks base method.
func (m *MockTaskRetriever) ListByIDAndUserID(ctx context.Context, id, userID int) (Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDAndUserID", ctx, id, userID)
	ret0, _ := ret[0].(Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDAndUserID indicates an expected call of ListByIDAndUserID.
func (mr *MockTaskRetrieverMockRecorder) ListByIDAndUserID(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDAndUserID", reflect.TypeOf((*MockTaskRetriever)(nil).ListByIDAndUserID), ctx, id, userID)
}

// ListByUserID mocks base method.
func (m *MockTaskRetriever) ListByUserID(ctx context.Context, userID int) ([]Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", ctx, userID)
	ret0, _ := ret[0].([]Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockTaskRetrieverMockRecorder) ListByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockTaskRetriever)(nil).ListByUserID), ctx, userID)
}

// MockTaskUpdater is a mock of TaskUpdater interface.
type MockTaskUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUpdaterMockRecorder
}

// MockTaskUpdaterMockRecorder is the mock recorder for MockTaskUpdater.
type MockTaskUpdaterMockRecorder struct {
	mock *MockTaskUpdater
}

// NewMockTaskUpdater creates a new mock instance.
func NewMockTaskUpdater(ctrl *gomock.Controller) *MockTaskUpdater {
	mock := &MockTaskUpdater{ctrl: ctrl}
	mock.recorder = &MockTaskUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUpdater) EXPECT() *MockTaskUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockTaskUpdater) Update(ctx context.Context, task Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTaskUpdaterMockRecorder) Update(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskUpdater)(nil).Update), ctx, task)
}

// MockTaskRemover is a mock of TaskRemover interface.
type MockTaskRemover struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRemoverMockRecorder
}

// MockTaskRemoverMockRecorder is the mock recorder for MockTaskRemover.
type MockTaskRemoverMockRecorder struct {
	mock *MockTaskRemover
}

// NewMockTaskRemover creates a new mock instance.
func NewMockTaskRemover(ctrl *gomock.Controller) *MockTaskRemover {
	mock := &MockTaskRemover{ctrl: ctrl}
	mock.recorder = &MockTaskRemoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRemover) EXPECT() *MockTaskRemoverMockRecorder {
	return m.recorder
}

// Remove mocks base method.
func (m *MockTaskRemover) Remove(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTaskRemoverMockRecorder) Remove(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTaskRemover)(nil).Remove), ctx, id)
}

// MockSummaryEncryptor is a mock of SummaryEncryptor interface.
type MockSummaryEncryptor struct {
	ctrl     *gomock.Controller
	recorder *MockSummaryEncryptorMockRecorder
}

// MockSummaryEncryptorMockRecorder is the mock recorder for MockSummaryEncryptor.
type MockSummaryEncryptorMockRecorder struct {
	mock *MockSummaryEncryptor
}

// NewMockSummaryEncryptor creates a new mock instance.
func NewMockSummaryEncryptor(ctrl *gomock.Controller) *MockSummaryEncryptor {
	mock := &MockSummaryEncryptor{ctrl: ctrl}
	mock.recorder = &MockSummaryEncryptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSummaryEncryptor) EXPECT() *MockSummaryEncryptorMockRecorder {
	return m.recorder
}

// Decrypt mocks base method.
func (m *MockSummaryEncryptor) Decrypt(value string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", value)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockSummaryEncryptorMockRecorder) Decrypt(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockSummaryEncryptor)(nil).Decrypt), value)
}

// Encrypt mocks base method.
func (m *MockSummaryEncryptor) Encrypt(value string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", value)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockSummaryEncryptorMockRecorder) Encrypt(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockSummaryEncryptor)(nil).Encrypt), value)
}
