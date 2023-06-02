// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	models "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/models"
)

// MockClient is an autogenerated mock type for the MockClient type
type MockClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *MockClient) Get(path string) ([]byte, error) {
	ret := _m.Called(path)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplications provides a mock function with given fields:
func (_m *MockClient) Applications() ([]models.Application, error) {
	ret := _m.Called()

	var r0 []models.Application
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Application, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Application); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Application)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClusterStats provides a mock function with given fields:
func (_m *MockClient) ClusterStats() (*models.ClusterProperties, error) {
	ret := _m.Called()

	var r0 *models.ClusterProperties
	var r1 error
	if rf, ok := ret.Get(0).(func() (*models.ClusterProperties, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *models.ClusterProperties); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ClusterProperties)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutorStats provides a mock function with given fields: appID
func (_m *MockClient) ExecutorStats(appID string) ([]models.Executor, error) {
	ret := _m.Called(appID)

	var r0 []models.Executor
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Executor, error)); ok {
		return rf(appID)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Executor); ok {
		r0 = rf(appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Executor)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobStats provides a mock function with given fields: appID
func (_m *MockClient) JobStats(appID string) ([]models.Job, error) {
	ret := _m.Called(appID)

	var r0 []models.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Job, error)); ok {
		return rf(appID)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Job); ok {
		r0 = rf(appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStageStats provides a mock function with given fields: appID
func (_m *MockClient) StageStats(appID string) ([]models.Stage, error) {
	ret := _m.Called(appID)

	var r0 []models.Stage
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Stage, error)); ok {
		return rf(appID)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Stage); ok {
		r0 = rf(appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Stage)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockClient interface {
	mock.TestingT
	Cleanup(func())
}

// newMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockClient(t mockConstructorTestingTnewMockClient) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
