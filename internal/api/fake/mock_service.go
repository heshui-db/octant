// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/api (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ForceUpdate mocks base method
func (m *MockService) ForceUpdate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUpdate")
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceUpdate indicates an expected call of ForceUpdate
func (mr *MockServiceMockRecorder) ForceUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUpdate", reflect.TypeOf((*MockService)(nil).ForceUpdate))
}

// Handler mocks base method
func (m *MockService) Handler(arg0 context.Context) (*mux.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler", arg0)
	ret0, _ := ret[0].(*mux.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handler indicates an expected call of Handler
func (mr *MockServiceMockRecorder) Handler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockService)(nil).Handler), arg0)
}
