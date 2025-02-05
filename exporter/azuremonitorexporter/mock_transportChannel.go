// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v1.0.0. DO NOT EDIT.

package azuremonitorexporter

import (
	"time"

	contracts "github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	mock "github.com/stretchr/testify/mock"
)

// mockTransportChannel is an autogenerated mock type for the transportChannel type
type mockTransportChannel struct {
	mock.Mock
}

func (_m *mockTransportChannel) EndpointAddress() string {
	_m.Called()
	return _m.String()
}

func (_m *mockTransportChannel) Stop() {
	_m.Called()
}

func (_m *mockTransportChannel) IsThrottled() bool {
	_m.Called()
	return false
}

func (_m *mockTransportChannel) Close(retryTimeout ...time.Duration) <-chan struct{} {
	_m.Called()

	closedChan := make(chan struct{})
	close(closedChan)

	return closedChan
}

// Send provides a mock function with given fields: _a0
func (_m *mockTransportChannel) Send(_a0 *contracts.Envelope) {
	_m.Called(_a0)
}

func (_m *mockTransportChannel) Flush() {
	_m.Called()
}
