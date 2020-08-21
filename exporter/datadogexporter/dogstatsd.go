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

package datadogexporter

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

type dogStatsDExporter struct {
	logger *zap.Logger
	cfg    *Config
	client *statsd.Client
}

func newDogStatsDExporter(logger *zap.Logger, cfg *Config) (*dogStatsDExporter, error) {

	client, err := statsd.New(
		cfg.MetricsURL,
		statsd.WithNamespace("opentelemetry."),
		statsd.WithTags(cfg.Tags),
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to initialize DogStatsD client: %s", err)
	}

	return &dogStatsDExporter{logger, cfg, client}, nil
}

func (exp *dogStatsDExporter) PushMetricsData(_ context.Context, md pdata.Metrics) (int, error) {
	metrics, droppedTimeSeries, err := MapMetrics(exp, md)

	if err != nil {
		return droppedTimeSeries, err
	}

	for name, data := range metrics {
		for _, metric := range data {
			switch metric.GetType() {
			case Count:
				err = exp.client.Count(name, metric.GetValue().(int64), metric.GetTags(), metric.GetRate())
			case Gauge:
				err = exp.client.Gauge(name, metric.GetValue().(float64), metric.GetTags(), metric.GetRate())
			}

			if err != nil {
				return droppedTimeSeries, err
			}
		}
	}

	return droppedTimeSeries, nil
}

func (exp *dogStatsDExporter) GetLogger() *zap.Logger {
	return exp.logger
}

func (exp *dogStatsDExporter) GetConfig() *Config {
	return exp.cfg
}

func (exp *dogStatsDExporter) GetQueueSettings() exporterhelper.QueueSettings {
	return exporterhelper.CreateDefaultQueueSettings()
}

func (exp *dogStatsDExporter) GetRetrySettings() exporterhelper.RetrySettings {
	return exporterhelper.CreateDefaultRetrySettings()
}
