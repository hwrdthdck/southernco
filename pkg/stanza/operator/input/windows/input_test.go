// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"golang.org/x/sys/windows"
)

// MockPersister is a mock implementation of the Persister interface.
type MockPersister struct {
	mock.Mock
}

// Get retrieves a value from the mock persister.
func (m *MockPersister) Get(ctx context.Context, key string) ([]byte, error) {
	args := m.Called(ctx, key)
	if args.Get(0) != nil {
		return args.Get(0).([]byte), args.Error(1)
	}
	return nil, args.Error(1)
}

// Set sets a value in the mock persister.
func (m *MockPersister) Set(ctx context.Context, key string, value []byte) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

// Delete deletes a value from the mock persister.
func (m *MockPersister) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

// TestInputStart_LocalSubscriptionError ensures the input correctly handles local subscription errors.
func TestInputStart_LocalSubscriptionError(t *testing.T) {
	persister := new(MockPersister)
	persister.On("Get", mock.Anything, "test-channel").Return(nil, nil)

	input := NewInput(zap.NewNop(), nil)
	input.channel = "test-channel"
	input.startAt = "beginning"
	input.pollInterval = 1 * time.Second

	err := input.Start(persister)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "The specified channel could not be found")
}

// TestInputStart_RemoteSubscriptionError ensures the input correctly handles remote subscription errors.
func TestInputStart_RemoteSubscriptionError(t *testing.T) {
	persister := new(MockPersister)
	persister.On("Get", mock.Anything, "test-channel").Return(nil, nil)

	input := NewInput(zap.NewNop(), func() error { return nil })
	input.channel = "test-channel"
	input.startAt = "beginning"
	input.pollInterval = 1 * time.Second
	input.remote = RemoteConfig{
		Server: "remote-server",
	}

	err := input.Start(persister)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "The specified channel could not be found")
}

// TestInputStart_RemoteSessionError ensures the input correctly handles remote session errors.
func TestInputStart_RemoteSessionError(t *testing.T) {
	persister := new(MockPersister)
	persister.On("Get", mock.Anything, "test-channel").Return(nil, nil)

	input := NewInput(zap.NewNop(), func() error {
		return errors.New("remote session error")
	})
	input.channel = "test-channel"
	input.startAt = "beginning"
	input.pollInterval = 1 * time.Second
	input.remote = RemoteConfig{
		Server: "remote-server",
	}

	err := input.Start(persister)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to start remote session: remote session error")
}

// TestInputStart_RemoteAccessDeniedError ensures the input correctly handles remote access denied errors.
func TestInputStart_RemoteAccessDeniedError(t *testing.T) {
	persister := new(MockPersister)
	persister.On("Get", mock.Anything, "test-channel").Return(nil, nil)

	evtSubscribeFunc = func(session uintptr, signalEvent windows.Handle, channelPath *uint16, query *uint16, bookmark uintptr, context uintptr, callback uintptr, flags uint32) (uintptr, error) {
		return 0, windows.ERROR_ACCESS_DENIED
	}

	input := NewInput(zap.NewNop(), func() error { return nil })
	input.channel = "test-channel"
	input.startAt = "beginning"
	input.pollInterval = 1 * time.Second
	input.remote = RemoteConfig{
		Server: "remote-server",
	}

	err := input.Start(persister)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open subscription for remote server")
	assert.Contains(t, err.Error(), "Access is denied")

	// Restore original evtSubscribeFunc
	evtSubscribeFunc = evtSubscribe
}
