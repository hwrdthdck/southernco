// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/models"
	mock "github.com/stretchr/testify/mock"
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
func (_m *MockClient) GetApplications() (*models.Applications, error) {
	ret := _m.Called()

	var r0 *models.Applications
	var r1 error
	if rf, ok := ret.Get(0).(func() (*models.Applications, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *models.Applications); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Applications)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterStats provides a mock function with given fields:
func (_m *MockClient) GetClusterStats() (*models.ClusterProperties, error) {
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
func (_m *MockClient) GetExecutorStats(appID string) (*models.Executors, error) {
	ret := _m.Called(appID)

	var r0 *models.Executors
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Executors, error)); ok {
		return rf(appID)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Executors); ok {
		r0 = rf(appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Executors)
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
func (_m *MockClient) GetJobStats(appID string) (*models.Jobs, error) {
	ret := _m.Called(appID)

	var r0 *models.Jobs
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Jobs, error)); ok {
		return rf(appID)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Jobs); ok {
		r0 = rf(appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Jobs)
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
func (_m *MockClient) GetStageStats(appID string) (*models.Stages, error) {
	ret := _m.Called(appID)

	var r0 *models.Stages
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Stages, error)); ok {
		return rf(appID)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Stages); ok {
		r0 = rf(appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Stages)
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
