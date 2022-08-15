// Copyright  The OpenTelemetry Authors
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

package postgresqlreceiver

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/service/featuregate"
)

type configFunc func(hostname string) *Config

// cleanupFunc exists to allow integration test cases to clean any registries it had
// to modify in order to change behavior of the integration test. i.e. featuregates
type cleanupFunc func()

type testCase struct {
	name         string
	cfg          configFunc
	cleanup      cleanupFunc
	expectedFile string
}

func TestPostgreSQLIntegration(t *testing.T) {
	testCases := []testCase{
		{
			name: "single_db",
			cfg: func(hostname string) *Config {
				f := NewFactory()
				cfg := f.CreateDefaultConfig().(*Config)
				cfg.Endpoint = net.JoinHostPort(hostname, "15432")
				cfg.Databases = []string{"otel"}
				cfg.Username = "otel"
				cfg.Password = "otel"
				cfg.Insecure = true
				return cfg
			},
			expectedFile: filepath.Join("testdata", "integration", "expected_single_db.json"),
		},
		{
			name: "multi_db",
			cfg: func(hostname string) *Config {
				f := NewFactory()
				cfg := f.CreateDefaultConfig().(*Config)
				cfg.Endpoint = net.JoinHostPort(hostname, "15432")
				cfg.Databases = []string{"otel", "otel2"}
				cfg.Username = "otel"
				cfg.Password = "otel"
				cfg.Insecure = true
				return cfg
			},
			expectedFile: filepath.Join("testdata", "integration", "expected_multi_db.json"),
		},
		{
			name: "all_db",
			cfg: func(hostname string) *Config {
				f := NewFactory()
				cfg := f.CreateDefaultConfig().(*Config)
				cfg.Endpoint = net.JoinHostPort(hostname, "15432")
				cfg.Databases = []string{}
				cfg.Username = "otel"
				cfg.Password = "otel"
				cfg.Insecure = true
				return cfg
			},
			expectedFile: filepath.Join("testdata", "integration", "expected_all_db.json"),
		},
		{
			name: "with_resource_attributes",
			cfg: func(hostname string) *Config {
				require.NoError(t, featuregate.GetRegistry().Apply(map[string]bool{
					emitMetricsWithResourceAttributesFeatureGateID: true,
				}))
				f := NewFactory()
				cfg := f.CreateDefaultConfig().(*Config)
				cfg.Endpoint = net.JoinHostPort(hostname, "15432")
				cfg.Databases = []string{}
				cfg.Username = "otel"
				cfg.Password = "otel"
				cfg.Insecure = true
				return cfg
			},
			cleanup: func() {
				require.NoError(t, featuregate.GetRegistry().Apply(map[string]bool{
					emitMetricsWithResourceAttributesFeatureGateID: false,
				}))
			},
			expectedFile: filepath.Join("testdata", "integration", "expected_all_with_resource_attributes.json"),
		},

		{
			name: "query_metrics",
			cfg: func(hostname string) *Config {
				require.NoError(t, featuregate.GetRegistry().Apply(map[string]bool{
					emitMetricsWithResourceAttributesFeatureGateID: true,
				}))
				f := NewFactory()
				cfg := f.CreateDefaultConfig().(*Config)
				cfg.Endpoint = net.JoinHostPort(hostname, "15432")
				cfg.Databases = []string{}
				cfg.Username = "otel"
				cfg.Password = "otel"
				cfg.Insecure = true
				cfg.CollectQueries = true
				return cfg
			},
			cleanup: func() {
				require.NoError(t, featuregate.GetRegistry().Apply(map[string]bool{
					emitMetricsWithResourceAttributesFeatureGateID: false,
				}))
			},
			expectedFile: filepath.Join("testdata", "integration", "expected_resource_with_query.json"),
		},
	}

	container := getContainer(t, testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    filepath.Join("testdata", "integration"),
			Dockerfile: "Dockerfile.postgresql",
		},
		ExposedPorts: []string{"15432:5432"},
		WaitingFor: wait.ForListeningPort("5432").
			WithStartupTimeout(2 * time.Minute),
	})
	defer func() {
		require.NoError(t, container.Terminate(context.Background()))
	}()
	hostname, err := container.Host(context.Background())
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.cleanup != nil {
				defer tc.cleanup()
			}
			// expectedMetrics, err := golden.ReadMetrics(tc.expectedFile)
			// require.NoError(t, err)

			f := NewFactory()
			consumer := new(consumertest.MetricsSink)
			settings := componenttest.NewNopReceiverCreateSettings()
			rcvr, err := f.CreateMetricsReceiver(context.Background(), settings, tc.cfg(hostname), consumer)
			require.NoError(t, err, "failed creating metrics receiver")
			require.NoError(t, rcvr.Start(context.Background(), componenttest.NewNopHost()))
			require.Eventuallyf(t, func() bool {
				return consumer.DataPointCount() > 0
			}, 2*time.Minute, 1*time.Second, "failed to receive more than 0 metrics")

			actualMetrics := consumer.AllMetrics()[0]
			if tc.name == "query_metrics" {
				bytes, err := pmetric.NewJSONMarshaler().MarshalMetrics(actualMetrics)
				require.NoError(t, err)
				err = ioutil.WriteFile(fmt.Sprintf("%s.back.json", tc.expectedFile), bytes, 0644)
				require.NoError(t, err)
			}

			// require.NoError(t, scrapertest.CompareMetrics(expectedMetrics, actualMetrics, scrapertest.IgnoreMetricValues()))
		})
	}
}

func getContainer(t *testing.T, req testcontainers.ContainerRequest) testcontainers.Container {
	require.NoError(t, req.Validate())
	container, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	require.NoError(t, err)
	return container
}
