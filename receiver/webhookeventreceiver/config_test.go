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

package webhookeventreceiver

import (
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver/internal/metadata"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.uber.org/multierr"
)

// only one validate check so far
func TestValidateConfig(t *testing.T) {
	t.Parallel()

	var errs error
	errs = multierr.Append(errs, errMissingEndpointFromConfig)
	errs = multierr.Append(errs, errReadTimeoutExceedsMaxValue)
	errs = multierr.Append(errs, errWriteTimeoutExceedsMaxValue)

	tests := []struct {
		desc   string
		expect error
		conf   Config
	}{
		{
			desc:   "Missing valid endpoint",
			expect: errMissingEndpointFromConfig,
			conf: Config{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: "",
				},
			},
		},
		{
			desc:   "ReadTimeout exceeds maximum value",
			expect: errReadTimeoutExceedsMaxValue,
			conf: Config{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: "0.0.0.0:0",
				},
				ReadTimeout: "14s",
			},
		},
		{
			desc:   "WriteTimeout exceeds maximum value",
			expect: errWriteTimeoutExceedsMaxValue,
			conf: Config{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: "0.0.0.0:0",
				},
				WriteTimeout: "14s",
			},
		},
		{
			desc:   "Multiple invalid configs",
			expect: errs,
			conf: Config{
				HTTPServerSettings: confighttp.HTTPServerSettings{
					Endpoint: "",
				},
				WriteTimeout: "14s",
				ReadTimeout:  "15s",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.conf.Validate()
			require.Error(t, err)
			require.Contains(t, err.Error(), test.expect.Error())
		})
	}
}

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	// LoadConf includes the TypeStr which NewFactory does not set
	id := component.NewIDWithName(metadata.Type, "valid_config")
	cmNoStr, err := cm.Sub(id.String())
	require.NoError(t, err)

	expect := &Config{
		HTTPServerSettings: confighttp.HTTPServerSettings{
			Endpoint: "0.0.0.0:8080",
		},
		ReadTimeout:  "500ms",
		WriteTimeout: "500ms",
		Path:         "some/path",
		HealthPath:   "health/path",
	}

	// create expected config
	factory := NewFactory()
	conf := factory.CreateDefaultConfig()
	require.NoError(t, component.UnmarshalConfig(cmNoStr, conf))
	require.NoError(t, component.ValidateConfig(conf))

	require.Equal(t, expect, conf)
}
