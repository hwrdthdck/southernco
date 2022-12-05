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

package spanbatchfilterprocessor

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestUnmarshalDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NoError(t, config.UnmarshalProcessor(confmap.New(), cfg))
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
}

func TestUnmarshalConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NoError(t, config.UnmarshalProcessor(cm, cfg))
	assert.Equal(t,
		&Config{
			ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
			TokensPerBatch:    1000,
			MostPopular:       false,
		}, cfg)
}

func TestValidateConfig_ReverseFilter(t *testing.T) {
	cfg := &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentIDWithName(typeStr, "2")),
		TokensPerBatch:    10,
		MostPopular:       true,
	}
	assert.NoError(t, cfg.Validate())

}

func TestValidateConfig_InvalidTokens(t *testing.T) {
	cfg := &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentIDWithName(typeStr, "2")),
		TokensPerBatch:    -10,
	}
	assert.Error(t, cfg.Validate())
}
