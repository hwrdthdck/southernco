// Code generated by mockery v2.13.1. DO NOT EDIT.

package cluster

import (
	aerospike "github.com/aerospike/aerospike-client-go/v5"
	mock "github.com/stretchr/testify/mock"
)

// mockNodeFactoryFunc is an autogenerated mock type for the nodeFactoryFunc type
type mockNodeFactoryFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockNodeFactoryFunc) Execute(_a0 *aerospike.ClientPolicy, _a1 *aerospike.Host, _a2 bool) (Node, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 Node
	if rf, ok := ret.Get(0).(func(*aerospike.ClientPolicy, *aerospike.Host, bool) Node); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*aerospike.ClientPolicy, *aerospike.Host, bool) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockNodeFactoryFunc interface {
	mock.TestingT
	Cleanup(func())
}

// newMockNodeFactoryFunc creates a new instance of mockNodeFactoryFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockNodeFactoryFunc(t mockConstructorTestingTnewMockNodeFactoryFunc) *mockNodeFactoryFunc {
	mock := &mockNodeFactoryFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
