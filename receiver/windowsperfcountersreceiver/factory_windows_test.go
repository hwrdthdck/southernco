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

//go:build windows
// +build windows

package windowsperfcountersreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/consumertest"
)

func TestCreateMetricsReceiver(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	cfg.(*Config).PerfCounters = []PerfCounterConfig{
		{
			Object:   "object",
			Counters: []CounterConfig{{Name: "counter", Metric: "metric"}},
		},
	}

	cfg.(*Config).MetricMetaData = []MetricConfig{
		{
			MetricName:  "metric",
			Description: "desc",
			Unit:        "1",
			Gauge: GaugeMetric{
				ValueType: "double",
			},
		},
	}

	mReceiver, err := factory.CreateMetricsReceiver(context.Background(), creationParams, cfg, consumertest.NewNop())

	assert.NoError(t, err)
	assert.NotNil(t, mReceiver)
}
