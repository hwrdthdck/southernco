// Copyright 2019, OpenTelemetry Authors
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

package stackdriverexporter

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/config/configtest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.ExampleComponents()
	assert.Nil(t, err)

	factory := NewFactory()
	factories.Exporters[configmodels.Type(typeStr)] = factory
	cfg, err := configtest.LoadConfigFile(
		t, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	assert.Equal(t, len(cfg.Exporters), 2)

	r0 := cfg.Exporters["stackdriver"]
	assert.Equal(t, r0, factory.CreateDefaultConfig())

	r1 := cfg.Exporters["stackdriver/customname"].(*Config)
	assert.Equal(t, r1,
		&Config{
			ExporterSettings:           configmodels.ExporterSettings{TypeVal: configmodels.Type(typeStr), NameVal: "stackdriver/customname"},
			ProjectID:                  "my-project",
			Prefix:                     "prefix",
			UserAgent:                  "opentelemetry-collector-contrib {{version}}",
			Endpoint:                   "test-endpoint",
			NumOfWorkers:               3,
			SkipCreateMetricDescriptor: true,
			UseInsecure:                true,
			TimeoutSettings: exporterhelper.TimeoutSettings{
				Timeout: 20 * time.Second,
			},
			ResourceMappings: []ResourceMapping{
				{
					SourceType: "source.resource1",
					TargetType: "target-resource1",
					LabelMappings: []LabelMapping{
						{
							SourceKey: "contrib.opencensus.io/exporter/stackdriver/project_id",
							TargetKey: "project_id",
							Optional:  true,
						},
						{
							SourceKey: "source.label1",
							TargetKey: "target_label_1",
							Optional:  false,
						},
					},
				},
				{
					SourceType: "source.resource2",
					TargetType: "target-resource2",
				},
			},
		})
}
