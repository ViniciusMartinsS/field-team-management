// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRetriever is a mock of UserRetriever interface.
type MockUserRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockUserRetrieverMockRecorder
}

// MockUserRetrieverMockRecorder is the mock recorder for MockUserRetriever.
type MockUserRetrieverMockRecorder struct {
	mock *MockUserRetriever
}

// NewMockUserRetriever creates a new mock instance.
func NewMockUserRetriever(ctrl *gomock.Controller) *MockUserRetriever {
	mock := &MockUserRetriever{ctrl: ctrl}
	mock.recorder = &MockUserRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRetriever) EXPECT() *MockUserRetrieverMockRecorder {
	return m.recorder
}

// ListByUserID mocks base method.
func (m *MockUserRetriever) ListByUserID(ctx context.Context, userID int) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", ctx, userID)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockUserRetrieverMockRecorder) ListByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockUserRetriever)(nil).ListByUserID), ctx, userID)
}
