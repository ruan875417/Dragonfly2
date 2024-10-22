// Code generated by MockGen. DO NOT EDIT.
// Source: resource.go
//
// Generated by this command:
//
//	mockgen -destination resource_mock.go -source resource.go -package persistentcache
//

// Package persistentcache is a generated GoMock package.
package persistentcache

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockResource is a mock of Resource interface.
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
}

// MockResourceMockRecorder is the mock recorder for MockResource.
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance.
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// HostManager mocks base method.
func (m *MockResource) HostManager() HostManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostManager")
	ret0, _ := ret[0].(HostManager)
	return ret0
}

// HostManager indicates an expected call of HostManager.
func (mr *MockResourceMockRecorder) HostManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostManager", reflect.TypeOf((*MockResource)(nil).HostManager))
}

// PeerManager mocks base method.
func (m *MockResource) PeerManager() PeerManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerManager")
	ret0, _ := ret[0].(PeerManager)
	return ret0
}

// PeerManager indicates an expected call of PeerManager.
func (mr *MockResourceMockRecorder) PeerManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerManager", reflect.TypeOf((*MockResource)(nil).PeerManager))
}

// TaskManager mocks base method.
func (m *MockResource) TaskManager() TaskManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskManager")
	ret0, _ := ret[0].(TaskManager)
	return ret0
}

// TaskManager indicates an expected call of TaskManager.
func (mr *MockResourceMockRecorder) TaskManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskManager", reflect.TypeOf((*MockResource)(nil).TaskManager))
}