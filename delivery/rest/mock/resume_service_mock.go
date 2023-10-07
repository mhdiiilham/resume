// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mhdiiilham/resume/delivery/rest (interfaces: ResumeService)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/mhdiiilham/resume/entity"
)

// MockResumeService is a mock of ResumeService interface.
type MockResumeService struct {
	ctrl     *gomock.Controller
	recorder *MockResumeServiceMockRecorder
}

// MockResumeServiceMockRecorder is the mock recorder for MockResumeService.
type MockResumeServiceMockRecorder struct {
	mock *MockResumeService
}

// NewMockResumeService creates a new mock instance.
func NewMockResumeService(ctrl *gomock.Controller) *MockResumeService {
	mock := &MockResumeService{ctrl: ctrl}
	mock.recorder = &MockResumeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResumeService) EXPECT() *MockResumeServiceMockRecorder {
	return m.recorder
}

// GetBasic mocks base method.
func (m *MockResumeService) GetBasic() entity.Basic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBasic")
	ret0, _ := ret[0].(entity.Basic)
	return ret0
}

// GetBasic indicates an expected call of GetBasic.
func (mr *MockResumeServiceMockRecorder) GetBasic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBasic", reflect.TypeOf((*MockResumeService)(nil).GetBasic))
}