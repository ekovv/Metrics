// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// GetCounter mocks base method.
func (m *MockStorage) GetCounter() map[string]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounter")
	ret0, _ := ret[0].(map[string]int)
	return ret0
}

// GetCounter indicates an expected call of GetCounter.
func (mr *MockStorageMockRecorder) GetCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounter", reflect.TypeOf((*MockStorage)(nil).GetCounter))
}

// GetGauge mocks base method.
func (m *MockStorage) GetGauge() map[string]float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGauge")
	ret0, _ := ret[0].(map[string]float64)
	return ret0
}

// GetGauge indicates an expected call of GetGauge.
func (mr *MockStorageMockRecorder) GetGauge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGauge", reflect.TypeOf((*MockStorage)(nil).GetGauge))
}

// SetCounter mocks base method.
func (m *MockStorage) SetCounter(metric string, value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCounter", metric, value)
}

// SetCounter indicates an expected call of SetCounter.
func (mr *MockStorageMockRecorder) SetCounter(metric, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCounter", reflect.TypeOf((*MockStorage)(nil).SetCounter), metric, value)
}

// SetGauge mocks base method.
func (m *MockStorage) SetGauge(metric string, value float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGauge", metric, value)
}

// SetGauge indicates an expected call of SetGauge.
func (mr *MockStorageMockRecorder) SetGauge(metric, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGauge", reflect.TypeOf((*MockStorage)(nil).SetGauge), metric, value)
}
