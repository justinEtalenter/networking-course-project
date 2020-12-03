// Code generated by MockGen. DO NOT EDIT.
// Source: yak/backend/pkg/repositories (interfaces: ObjectPerms)

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"
	models "yak/backend/pkg/models"

	gomock "github.com/golang/mock/gomock"
)

// MockObjectPerms is a mock of ObjectPerms interface.
type MockObjectPerms struct {
	ctrl     *gomock.Controller
	recorder *MockObjectPermsMockRecorder
}

// MockObjectPermsMockRecorder is the mock recorder for MockObjectPerms.
type MockObjectPermsMockRecorder struct {
	mock *MockObjectPerms
}

// NewMockObjectPerms creates a new mock instance.
func NewMockObjectPerms(ctrl *gomock.Controller) *MockObjectPerms {
	mock := &MockObjectPerms{ctrl: ctrl}
	mock.recorder = &MockObjectPermsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectPerms) EXPECT() *MockObjectPermsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockObjectPerms) Create(arg0, arg1, arg2 int, arg3 *models.Permission) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockObjectPermsMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockObjectPerms)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockObjectPerms) Delete(arg0, arg1, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockObjectPermsMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockObjectPerms)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Get mocks base method.
func (m *MockObjectPerms) Get(arg0, arg1, arg2 int) (*models.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockObjectPermsMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockObjectPerms)(nil).Get), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockObjectPerms) Update(arg0, arg1, arg2, arg3 int, arg4 *models.UpdatePermission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockObjectPermsMockRecorder) Update(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockObjectPerms)(nil).Update), arg0, arg1, arg2, arg3, arg4)
}
