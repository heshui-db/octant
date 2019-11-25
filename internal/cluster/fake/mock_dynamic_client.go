// Code generated by MockGen. DO NOT EDIT.
// Source: ../../vendor/k8s.io/client-go/dynamic/interface.go

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	dynamic "k8s.io/client-go/dynamic"
	reflect "reflect"
)

// MockDynamicInterface is a mock of Interface interface
type MockDynamicInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicInterfaceMockRecorder
}

// MockDynamicInterfaceMockRecorder is the mock recorder for MockDynamicInterface
type MockDynamicInterfaceMockRecorder struct {
	mock *MockDynamicInterface
}

// NewMockDynamicInterface creates a new mock instance
func NewMockDynamicInterface(ctrl *gomock.Controller) *MockDynamicInterface {
	mock := &MockDynamicInterface{ctrl: ctrl}
	mock.recorder = &MockDynamicInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamicInterface) EXPECT() *MockDynamicInterfaceMockRecorder {
	return m.recorder
}

// Resource mocks base method
func (m *MockDynamicInterface) Resource(resource schema.GroupVersionResource) dynamic.NamespaceableResourceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resource", resource)
	ret0, _ := ret[0].(dynamic.NamespaceableResourceInterface)
	return ret0
}

// Resource indicates an expected call of Resource
func (mr *MockDynamicInterfaceMockRecorder) Resource(resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resource", reflect.TypeOf((*MockDynamicInterface)(nil).Resource), resource)
}

// MockResourceInterface is a mock of ResourceInterface interface
type MockResourceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceInterfaceMockRecorder
}

// MockResourceInterfaceMockRecorder is the mock recorder for MockResourceInterface
type MockResourceInterfaceMockRecorder struct {
	mock *MockResourceInterface
}

// NewMockResourceInterface creates a new mock instance
func NewMockResourceInterface(ctrl *gomock.Controller) *MockResourceInterface {
	mock := &MockResourceInterface{ctrl: ctrl}
	mock.recorder = &MockResourceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceInterface) EXPECT() *MockResourceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockResourceInterface) Create(obj *unstructured.Unstructured, options v1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockResourceInterfaceMockRecorder) Create(obj, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResourceInterface)(nil).Create), varargs...)
}

// Update mocks base method
func (m *MockResourceInterface) Update(obj *unstructured.Unstructured, options v1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockResourceInterfaceMockRecorder) Update(obj, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceInterface)(nil).Update), varargs...)
}

// UpdateStatus mocks base method
func (m *MockResourceInterface) UpdateStatus(obj *unstructured.Unstructured, options v1.UpdateOptions) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", obj, options)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockResourceInterfaceMockRecorder) UpdateStatus(obj, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockResourceInterface)(nil).UpdateStatus), obj, options)
}

// Delete mocks base method
func (m *MockResourceInterface) Delete(name string, options *v1.DeleteOptions, subresources ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockResourceInterfaceMockRecorder) Delete(name, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourceInterface)(nil).Delete), varargs...)
}

// DeleteCollection mocks base method
func (m *MockResourceInterface) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", options, listOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockResourceInterfaceMockRecorder) DeleteCollection(options, listOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockResourceInterface)(nil).DeleteCollection), options, listOptions)
}

// Get mocks base method
func (m *MockResourceInterface) Get(name string, options v1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockResourceInterfaceMockRecorder) Get(name, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceInterface)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockResourceInterface) List(opts v1.ListOptions) (*unstructured.UnstructuredList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", opts)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockResourceInterfaceMockRecorder) List(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceInterface)(nil).List), opts)
}

// Watch mocks base method
func (m *MockResourceInterface) Watch(opts v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", opts)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockResourceInterfaceMockRecorder) Watch(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockResourceInterface)(nil).Watch), opts)
}

// Patch mocks base method
func (m *MockResourceInterface) Patch(name string, pt types.PatchType, data []byte, options v1.PatchOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, pt, data, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockResourceInterfaceMockRecorder) Patch(name, pt, data, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, pt, data, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockResourceInterface)(nil).Patch), varargs...)
}

// MockNamespaceableResourceInterface is a mock of NamespaceableResourceInterface interface
type MockNamespaceableResourceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceableResourceInterfaceMockRecorder
}

// MockNamespaceableResourceInterfaceMockRecorder is the mock recorder for MockNamespaceableResourceInterface
type MockNamespaceableResourceInterfaceMockRecorder struct {
	mock *MockNamespaceableResourceInterface
}

// NewMockNamespaceableResourceInterface creates a new mock instance
func NewMockNamespaceableResourceInterface(ctrl *gomock.Controller) *MockNamespaceableResourceInterface {
	mock := &MockNamespaceableResourceInterface{ctrl: ctrl}
	mock.recorder = &MockNamespaceableResourceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceableResourceInterface) EXPECT() *MockNamespaceableResourceInterfaceMockRecorder {
	return m.recorder
}

// Namespace mocks base method
func (m *MockNamespaceableResourceInterface) Namespace(arg0 string) dynamic.ResourceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace", arg0)
	ret0, _ := ret[0].(dynamic.ResourceInterface)
	return ret0
}

// Namespace indicates an expected call of Namespace
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Namespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Namespace), arg0)
}

// Create mocks base method
func (m *MockNamespaceableResourceInterface) Create(obj *unstructured.Unstructured, options v1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Create(obj, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Create), varargs...)
}

// Update mocks base method
func (m *MockNamespaceableResourceInterface) Update(obj *unstructured.Unstructured, options v1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Update(obj, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Update), varargs...)
}

// UpdateStatus mocks base method
func (m *MockNamespaceableResourceInterface) UpdateStatus(obj *unstructured.Unstructured, options v1.UpdateOptions) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", obj, options)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockNamespaceableResourceInterfaceMockRecorder) UpdateStatus(obj, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).UpdateStatus), obj, options)
}

// Delete mocks base method
func (m *MockNamespaceableResourceInterface) Delete(name string, options *v1.DeleteOptions, subresources ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Delete(name, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Delete), varargs...)
}

// DeleteCollection mocks base method
func (m *MockNamespaceableResourceInterface) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", options, listOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockNamespaceableResourceInterfaceMockRecorder) DeleteCollection(options, listOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).DeleteCollection), options, listOptions)
}

// Get mocks base method
func (m *MockNamespaceableResourceInterface) Get(name string, options v1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Get(name, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockNamespaceableResourceInterface) List(opts v1.ListOptions) (*unstructured.UnstructuredList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", opts)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockNamespaceableResourceInterfaceMockRecorder) List(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).List), opts)
}

// Watch mocks base method
func (m *MockNamespaceableResourceInterface) Watch(opts v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", opts)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Watch(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Watch), opts)
}

// Patch mocks base method
func (m *MockNamespaceableResourceInterface) Patch(name string, pt types.PatchType, data []byte, options v1.PatchOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, pt, data, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Patch(name, pt, data, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, pt, data, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Patch), varargs...)
}