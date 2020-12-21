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

package ecsobserver

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.ExampleComponents()
	assert.NoError(t, err)

	factory := NewFactory()
	factories.Extensions[typeStr] = factory
	cfg, err := configtest.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"), factories)

	require.Nil(t, err)
	require.NotNil(t, cfg)

	require.Len(t, cfg.Extensions, 2)

	ext0 := cfg.Extensions["ecs_observer"]
	assert.Equal(t, factory.CreateDefaultConfig(), ext0)

	ext1 := cfg.Extensions["ecs_observer/1"]
	assert.Equal(t,
		&Config{
			ExtensionSettings: configmodels.ExtensionSettings{
				TypeVal: "ecs_observer",
				NameVal: "ecs_observer/1",
			},
			RefreshInterval: 15 * time.Second,
			ClusterName:     "EC2-Testing",
			ClusterRegion:   "us-west-2",
			ResultFile:      "/opt/aws/amazon-cloudwatch-agent/etc/ecs_sd_targets.yaml",
			DockerLabel: &DockerLabelConfig{
				JobNameLabel:     "ECS_PROMETHEUS_JOB_NAME",
				MetricsPathLabel: "ECS_PROMETHEUS_METRICS_PATH",
				PortLabel:        "ECS_PROMETHEUS_EXPORTER_PORT_SUBSET_A",
			},
			TaskDefinitions: []*TaskDefinitionConfig{
				&TaskDefinitionConfig{
					JobName:           "task_def_1",
					MetricsPath:       "/stats/metrics",
					MetricsPorts:      "9901;9404;9406",
					TaskDefArnPattern: ".*:task-definition/bugbash-java-fargate-awsvpc-task-def-only:[0-9]+",
				},
				&TaskDefinitionConfig{
					ContainerNamePattern: "^bugbash-jar.*$",
					MetricsPorts:         "9902",
					TaskDefArnPattern:    ".*:task-definition/nginx:[0-9]+",
				},
			},
		},
		ext1,
	)
}
