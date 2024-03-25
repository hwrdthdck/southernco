// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingexporter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/exporter/exportertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/testutil"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateTracesExporter(t *testing.T) {
	endpoint := testutil.GetAvailableLocalAddress(t)
	tests := []struct {
		name             string
		config           *Config
		mustFailOnCreate bool
		mustFailOnStart  bool
	}{
		{
			name: "UseSecure",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					TLSSetting: configtls.ClientConfig{
						Insecure: false,
					},
				},
				NumStreams: 3,
			},
		},
		{
			name: "Keepalive",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					Keepalive: &configgrpc.KeepaliveClientConfig{
						Time:                30 * time.Second,
						Timeout:             25 * time.Second,
						PermitWithoutStream: true,
					},
				},
				NumStreams: 3,
			},
		},
		{
			name: "Compression",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    endpoint,
					Compression: "gzip",
				},
				NumStreams: 3,
			},
		},
		{
			name: "Headers",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					Headers: map[string]configopaque.String{
						"hdr1": "val1",
						"hdr2": "val2",
					},
				},
				NumStreams: 3,
			},
		},
		{
			name: "CompressionError",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    endpoint,
					Compression: "unknown compression",
				},
				NumStreams: 3,
			},
			mustFailOnCreate: false,
			mustFailOnStart:  true,
		},
		{
			name: "CaCert",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					TLSSetting: configtls.ClientConfig{
						TLSSetting: configtls.Config{
							CAFile: "testdata/test_cert.pem",
						},
						Insecure: false,
					},
				},
				NumStreams: 3,
			},
		},
		{
			name: "CertPemFileError",
			config: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					TLSSetting: configtls.ClientConfig{
						TLSSetting: configtls.Config{
							CAFile: "nosuchfile",
						},
					},
				},
				NumStreams: 3,
			},
			mustFailOnCreate: false,
			mustFailOnStart:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := exportertest.NewNopCreateSettings()
			tExporter, tErr := createLogsExporter(context.Background(), set, tt.config)
			checkErrorsAndStartAndShutdown(t, tExporter, tErr, tt.mustFailOnCreate, tt.mustFailOnStart)
			tExporter2, tErr2 := createMetricsExporter(context.Background(), set, tt.config)
			checkErrorsAndStartAndShutdown(t, tExporter2, tErr2, tt.mustFailOnCreate, tt.mustFailOnStart)
		})
	}
}

func checkErrorsAndStartAndShutdown(t *testing.T, exporter component.Component, err error, mustFailOnCreate, mustFailOnStart bool) {
	if mustFailOnCreate {
		assert.Error(t, err)
		return
	}
	assert.NoError(t, err)
	assert.NotNil(t, exporter)

	sErr := exporter.Start(context.Background(), componenttest.NewNopHost())
	if mustFailOnStart {
		require.Error(t, sErr)
		return
	}
	require.NoError(t, sErr)
	require.NoError(t, exporter.Shutdown(context.Background()))
}
