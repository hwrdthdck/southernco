// Copyright OpenTelemetry Authors
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

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"context"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricExporter struct {
	config           *Config
	transportChannel transportChannel
	logger           *zap.Logger
	packer           *metricPacker
}

func (exporter *metricExporter) onMetricData(context context.Context, metricData pmetric.Metrics) error {
	resourceMetrics := metricData.ResourceMetrics()

	for i := 0; i < resourceMetrics.Len(); i++ {
		scopeMetrics := resourceMetrics.At(i).ScopeMetrics()
		resource := resourceMetrics.At(i).Resource()
		for j := 0; j < scopeMetrics.Len(); j++ {
			metrics := scopeMetrics.At(j).Metrics()
			scope := scopeMetrics.At(j).Scope()
			for k := 0; k < metrics.Len(); k++ {
				for _, envelope := range exporter.packer.MetricToEnvelopes(metrics.At(k), resource, scope) {
					envelope.IKey = string(exporter.config.InstrumentationKey)
					exporter.transportChannel.Send(envelope)
				}
			}
		}
	}

	return nil
}

// Returns a new instance of the metric exporter
func newMetricsExporter(config *Config, transportChannel transportChannel, set exporter.CreateSettings) (exporter.Metrics, error) {
	exporter := &metricExporter{
		config:           config,
		transportChannel: transportChannel,
		logger:           set.Logger,
		packer:           newMetricPacker(set.Logger),
	}

	return exporterhelper.NewMetricsExporter(context.TODO(), set, config, exporter.onMetricData)
}
