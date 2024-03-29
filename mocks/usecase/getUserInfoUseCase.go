// Code generated by MockGen. DO NOT EDIT.
// Source: src/usecase/getUserInfoUseCase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/oyuno-hito/gin-helloworld/src/repository"
)

// MockUserInfoUseCase is a mock of UserInfoUseCase interface.
type MockUserInfoUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserInfoUseCaseMockRecorder
}

// MockUserInfoUseCaseMockRecorder is the mock recorder for MockUserInfoUseCase.
type MockUserInfoUseCaseMockRecorder struct {
	mock *MockUserInfoUseCase
}

// NewMockUserInfoUseCase creates a new mock instance.
func NewMockUserInfoUseCase(ctrl *gomock.Controller) *MockUserInfoUseCase {
	mock := &MockUserInfoUseCase{ctrl: ctrl}
	mock.recorder = &MockUserInfoUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInfoUseCase) EXPECT() *MockUserInfoUseCaseMockRecorder {
	return m.recorder
}

// GetUserInfoUseCase mocks base method.
func (m *MockUserInfoUseCase) GetUserInfoUseCase(id int) (*repository.UserRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoUseCase", id)
	ret0, _ := ret[0].(*repository.UserRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoUseCase indicates an expected call of GetUserInfoUseCase.
func (mr *MockUserInfoUseCaseMockRecorder) GetUserInfoUseCase(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoUseCase", reflect.TypeOf((*MockUserInfoUseCase)(nil).GetUserInfoUseCase), id)
}
