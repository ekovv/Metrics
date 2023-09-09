// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Inc provides a mock function with given fields: name, value
func (_m *Repository) Inc(name string, value float64) {
	_m.Called(name, value)
}

// Set provides a mock function with given fields: name, value
func (_m *Repository) Set(name string, value float64) {
	_m.Called(name, value)
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}