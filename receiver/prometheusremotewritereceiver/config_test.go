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

package prometheusremotewritereceiver

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	tests := []struct {
		id       config.ComponentID
		expected config.Receiver
	}{
		{
			id:       config.NewComponentIDWithName(typeStr, "defaults"),
			expected: createDefaultConfig(),
		},
		{
			id: config.NewComponentIDWithName(typeStr, ""),
			expected: &Config{
				ReceiverSettings: config.NewReceiverSettings(config.NewComponentID("prometheusremotewrite")),
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint:           "0.0.0.0:19291",
					TLSSetting:         (*configtls.TLSServerSetting)(nil),
					CORS:               (*confighttp.CORSSettings)(nil),
					Auth:               (*configauth.Authentication)(nil),
					MaxRequestBodySize: 0,
					IncludeMetadata:    false},
				TimeThreshold: 24,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, config.UnmarshalReceiver(sub, cfg))
			assert.NoError(t, cfg.Validate())
			assert.Equal(t, tt.expected, cfg)
		})
	}
}
