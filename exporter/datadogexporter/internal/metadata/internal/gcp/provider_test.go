// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcp

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testShortHostname = "hostname"
	testCloudAccount  = "projectID"
	testHostname      = testShortHostname + ".c." + testCloudAccount + ".internal"
	testBadHostname   = "badhostname"
)

var (
	testGCPIntegrationHostname    = fmt.Sprintf("%s.%s", testShortHostname, testCloudAccount)
	testGCPIntegrationBadHostname = fmt.Sprintf("%s.%s", testBadHostname, testCloudAccount)
)

var _ gcpDetector = (*mockDetector)(nil)

type mockDetector struct {
	projectID    string
	instanceName string
}

func (m *mockDetector) CloudPlatform() gcp.Platform {
	return gcp.GCE
}

func (m *mockDetector) ProjectID() (string, error) {
	return m.projectID, nil
}

func (m *mockDetector) GCEHostName() (string, error) {
	return m.instanceName, nil
}

func TestProvider(t *testing.T) {
	tests := []struct {
		name         string
		projectID    string
		instanceName string
		hostname     string
	}{
		{
			name:         "good hostname",
			projectID:    testCloudAccount,
			instanceName: testHostname,
			hostname:     testGCPIntegrationHostname,
		},
		{
			name:         "bad hostname",
			projectID:    testCloudAccount,
			instanceName: testBadHostname,
			hostname:     testGCPIntegrationBadHostname,
		},
	}

	for _, testInstance := range tests {
		t.Run(testInstance.name, func(t *testing.T) {
			provider := &Provider{detector: &mockDetector{
				projectID:    testInstance.projectID,
				instanceName: testInstance.instanceName,
			}}

			hostname, err := provider.Hostname(context.Background())
			require.NoError(t, err)
			assert.Equal(t, testInstance.hostname, hostname)
		})
	}
}
