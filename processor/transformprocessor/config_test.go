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

package transformprocessor

import (
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/telemetryquerylanguage/tqlconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/service/servicetest"
)

func TestLoadingConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.NoError(t, err)

	factory := NewFactory()
	factories.Processors[typeStr] = factory
	cfg, err := servicetest.LoadConfigAndValidate(filepath.Join("testdata", "config.yaml"), factories)
	assert.NoError(t, err)
	require.NotNil(t, cfg)

	p0 := cfg.Processors[config.NewComponentID(typeStr)]
	assert.Equal(t, p0, &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
		Config: tqlconfig.Config{
			Traces: tqlconfig.SignalConfig{
				Queries: []string{
					`set(name, "bear") where attributes["http.path"] == "/animal"`,
					`keep_keys(attributes, "http.method", "http.path")`,
				},
			},
			Metrics: tqlconfig.SignalConfig{
				Queries: []string{
					`set(metric.name, "bear") where attributes["http.path"] == "/animal"`,
					`keep_keys(attributes, "http.method", "http.path")`,
				},
			},
			Logs: tqlconfig.SignalConfig{
				Queries: []string{
					`set(body, "bear") where attributes["http.path"] == "/animal"`,
					`keep_keys(attributes, "http.method", "http.path")`,
				},
			},
		},
	})
}

func TestLoadInvalidConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.NoError(t, err)

	factory := NewFactory()
	factories.Processors[typeStr] = factory

	cfg, err := servicetest.LoadConfigAndValidate(filepath.Join("testdata", "invalid_config_bad_syntax_trace.yaml"), factories)
	assert.Error(t, err)
	assert.NotNil(t, cfg)

	cfg, err = servicetest.LoadConfigAndValidate(filepath.Join("testdata", "invalid_config_unknown_function_trace.yaml"), factories)
	assert.Error(t, err)
	assert.NotNil(t, cfg)

	cfg, err = servicetest.LoadConfigAndValidate(filepath.Join("testdata", "invalid_config_bad_syntax_metric.yaml"), factories)
	assert.Error(t, err)
	assert.NotNil(t, cfg)

	cfg, err = servicetest.LoadConfigAndValidate(filepath.Join("testdata", "invalid_config_unknown_function_metric.yaml"), factories)
	assert.Error(t, err)
	assert.NotNil(t, cfg)

	cfg, err = servicetest.LoadConfigAndValidate(filepath.Join("testdata", "invalid_config_bad_syntax_log.yaml"), factories)
	assert.Error(t, err)
	assert.NotNil(t, cfg)

	cfg, err = servicetest.LoadConfigAndValidate(filepath.Join("testdata", "invalid_config_unknown_function_log.yaml"), factories)
	assert.Error(t, err)
	assert.NotNil(t, cfg)
}
