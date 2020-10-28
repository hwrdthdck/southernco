// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/translator/conventions"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

type mockMetadata struct {
	mock.Mock
}

func (m *mockMetadata) FQDNAvailable() bool {
	return m.MethodCalled("FQDNAvailable").Bool(0)
}

func (m *mockMetadata) FQDN(_ context.Context) (string, error) {
	args := m.MethodCalled("FQDN")
	return args.String(0), args.Error(1)
}

func (m *mockMetadata) Hostname() (string, error) {
	args := m.MethodCalled("Hostname")
	return args.String(0), args.Error(1)
}

func (m *mockMetadata) HostType() (string, error) {
	args := m.MethodCalled("HostType")
	return args.String(0), args.Error(1)
}

func TestNewDetector(t *testing.T) {
	d, err := NewDetector()
	require.NoError(t, err)
	assert.NotNil(t, d)
}

func TestDetectFQDNAvailable(t *testing.T) {
	md := &mockMetadata{}
	md.On("FQDNAvailable").Return(true)
	md.On("FQDN").Return("fqdn", nil)
	md.On("Hostname").Return("hostname", nil)
	md.On("HostType").Return("darwin/amd64", nil)

	detector := &Detector{provider: md}
	res, err := detector.Detect(context.Background())
	require.NoError(t, err)
	md.AssertExpectations(t)
	res.Attributes().Sort()

	expected := internal.NewResource(map[string]interface{}{
		conventions.AttributeHostName:     "fqdn",
		conventions.AttributeHostHostname: "hostname",
		conventions.AttributeHostType:     "darwin/amd64",
	})
	expected.Attributes().Sort()

	assert.Equal(t, expected, res)

}

func TestDetectFQDNNoAvailable(t *testing.T) {
	md := &mockMetadata{}
	md.On("FQDNAvailable").Return(false)
	md.On("Hostname").Return("hostname", nil)
	md.On("HostType").Return("darwin/amd64", nil)

	detector := &Detector{provider: md}
	res, err := detector.Detect(context.Background())
	require.NoError(t, err)
	res.Attributes().Sort()

	expected := internal.NewResource(map[string]interface{}{
		conventions.AttributeHostHostname: "hostname",
		conventions.AttributeHostType:     "darwin/amd64",
	})
	expected.Attributes().Sort()

	assert.Equal(t, expected, res)

}

func TestDetectError(t *testing.T) {
	// FQDN fails
	mdFQDN := &mockMetadata{}
	mdFQDN.On("Hostname").Return("", errors.New("err"))
	mdFQDN.On("HostType").Return("windows/arm64", nil)
	mdFQDN.On("FQDNAvailable").Return(true)
	mdFQDN.On("FQDN").Return("", errors.New("err"))

	detector := &Detector{provider: mdFQDN}
	res, err := detector.Detect(context.Background())
	assert.Error(t, err)
	assert.True(t, internal.IsEmptyResource(res))

	// Hostname fails
	mdHostname := &mockMetadata{}
	mdHostname.On("FQDNAvailable").Return(true)
	mdHostname.On("FQDN").Return("fqdn", nil)
	mdHostname.On("Hostname").Return("", errors.New("err"))
	mdHostname.On("HostType").Return("windows/arm64", nil)

	detector = &Detector{provider: mdHostname}
	res, err = detector.Detect(context.Background())
	assert.Error(t, err)
	assert.True(t, internal.IsEmptyResource(res))
}
