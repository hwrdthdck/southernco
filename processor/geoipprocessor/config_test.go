// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor

import (
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/conventions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id       component.ID
		expected component.Config
	}{
		{
			id: component.NewID(metadata.Type),
			expected: &Config{
				Metadata: enabledAttributes(),
			},
		},
		{
			id: component.NewIDWithName(metadata.Type, "2"),
			expected: &Config{
				Metadata: []string{conventions.AttributeGeoCityName},
			},
		},
		{
			id: component.NewIDWithName(metadata.Type, "invalid_attribute"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
			require.NoError(t, err)

			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalConfig(sub, cfg))

			if tt.expected == nil {
				err = component.ValidateConfig(cfg)
				assert.Error(t, err)
				return
			}
			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}
