// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	aerospike "github.com/aerospike/aerospike-client-go/v6"
	mock "github.com/stretchr/testify/mock"
)

// Asclient is an autogenerated mock type for the asclient type
type Asclient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Asclient) Close() {
	_m.Called()
}

// GetNodes provides a mock function with given fields:
func (_m *Asclient) GetNodes() []*aerospike.Node {
	ret := _m.Called()

	var r0 []*aerospike.Node
	if rf, ok := ret.Get(0).(func() []*aerospike.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*aerospike.Node)
		}
	}

	return r0
}

type mockConstructorTestingTNewAsclient interface {
	mock.TestingT
	Cleanup(func())
}

// NewAsclient creates a new instance of Asclient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAsclient(t mockConstructorTestingTNewAsclient) *Asclient {
	mock := &Asclient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
