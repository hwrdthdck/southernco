// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sematextexporter

import (
	"path/filepath"
	"testing"
	"time"
	"os"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sematextexporter/internal/metadata"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	content, err := os.ReadFile(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	// Step 3: Replace placeholders with dynamically generated UUIDs
	contentStr := strings.ReplaceAll(string(content), "<METRICS_APP_TOKEN>", metricsAppToken)
	contentStr = strings.ReplaceAll(contentStr, "<LOGS_APP_TOKEN>", logsAppToken)

	// Step 4: Write the updated content to a temporary file for testing
	tmpConfigPath := filepath.Join("testdata", "config_tmp.yaml")
	err = os.WriteFile(tmpConfigPath, []byte(contentStr), 0644) // Replace ioutil.WriteFile with os.WriteFile
	require.NoError(t, err)
	defer os.Remove(tmpConfigPath) // Clean up after the test

	// Step 5: Load the modified configuration file
	cm, err := confmaptest.LoadConf(tmpConfigPath)
	require.NoError(t, err)
	tests := []struct {
		id       component.ID
		expected component.Config
	}{
		{
			id:       component.NewIDWithName(metadata.Type, "default-config"),
			expected: createDefaultConfig(),
		},
		{
			id: component.NewIDWithName(metadata.Type, "override-config"),
			expected: &Config{
				ClientConfig: confighttp.ClientConfig{
					Timeout: 500 * time.Millisecond,
					Headers: map[string]configopaque.String{"User-Agent": "OpenTelemetry -> Sematext"},
				},
				MetricsConfig: MetricsConfig{
					MetricsEndpoint: usMetricsEndpoint,
					QueueSettings: exporterhelper.QueueConfig{
						Enabled:      true,
						NumConsumers: 3,
						QueueSize:    10,
					},
					AppToken:        metricsAppToken,
					MetricsSchema:   "telegraf-prometheus-v2",
					PayloadMaxLines: 72,
					PayloadMaxBytes: 27,
				},
				LogsConfig: LogsConfig{
					AppToken:     logsAppToken,
					LogsEndpoint: usLogsEndpoint,
				},

				BackOffConfig: configretry.BackOffConfig{
					Enabled:             true,
					InitialInterval:     1 * time.Second,
					MaxInterval:         3 * time.Second,
					MaxElapsedTime:      10 * time.Second,
					RandomizationFactor: backoff.DefaultRandomizationFactor,
					Multiplier:          backoff.DefaultMultiplier,
				},
				Region: usRegion,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, sub.Unmarshal(cfg))

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}
func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      *Config
		expectError bool
	}{
		{
			name: "Valid configuration 1",
			config: &Config{
				Region: usRegion,
				MetricsConfig: MetricsConfig{
					AppToken: metricsAppToken,
				},
				LogsConfig: LogsConfig{
					AppToken: logsAppToken,
				},
			},
			expectError: false,
		},
		{
			name: "Valid configuration 2",
			config: &Config{
				Region: euRegion,
				MetricsConfig: MetricsConfig{
					AppToken: metricsAppToken,
				},
				LogsConfig: LogsConfig{
					AppToken: logsAppToken,
				},
			},
			expectError: false,
		},
		{
			name: "Invalid region",
			config: &Config{
				Region: "ASIA",
				MetricsConfig: MetricsConfig{
					AppToken: metricsAppToken,
				},
				LogsConfig: LogsConfig{
					AppToken: logsAppToken,
				},
			},
			expectError: true,
		},
		{
			name: "Invalid metrics AppToken",
			config: &Config{
				Region: usRegion,
				MetricsConfig: MetricsConfig{
					AppToken: "short-token",
				},
				LogsConfig: LogsConfig{
					AppToken: logsAppToken,
				},
			},
			expectError: true,
		},
		{
			name: "Invalid logs AppToken",
			config: &Config{
				Region: euRegion,
				MetricsConfig: MetricsConfig{
					AppToken: metricsAppToken,
				},
				LogsConfig: LogsConfig{
					AppToken: "short-token",
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.expectError {
				assert.Error(t, err, "Expected an error for invalid configuration")
			} else {
				assert.NoError(t, err, "Expected no error for valid configuration")
			}
		})
	}
}
