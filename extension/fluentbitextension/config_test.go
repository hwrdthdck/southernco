// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fluentbitextension

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.NoError(t, err)

	factory := NewFactory()
	factories.Extensions[typeStr] = factory
	cfg, err := configtest.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"), factories)

	require.Nil(t, err)
	require.NotNil(t, cfg)

	ext0 := cfg.Extensions["fluentbit"]
	assert.Equal(t, factory.CreateDefaultConfig(), ext0)

	ext1 := cfg.Extensions["fluentbit/1"]
	assert.Equal(t,
		&Config{
			ExtensionSettings: configmodels.ExtensionSettings{
				TypeVal: "fluentbit",
				NameVal: "fluentbit/1",
			},
			ExecutablePath: "/usr/local/bin/fluent-bit",
		},
		ext1)

	assert.Equal(t, 1, len(cfg.Service.Extensions))
	assert.Equal(t, "fluentbit/1", cfg.Service.Extensions[0])
}
