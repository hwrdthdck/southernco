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

//go:build windows
// +build windows

package iisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"

import (
	"context"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type iisReceiver struct {
	params   component.ReceiverCreateSettings
	config   *Config
	consumer consumer.Metrics
	watchers []winperfcounters.PerfCounterWatcher
}

// new returns a iisReceiver
func newIisReceiver(params component.ReceiverCreateSettings, cfg *Config, consumer consumer.Metrics) *iisReceiver {
	return &iisReceiver{params: params, config: cfg, consumer: consumer}
}

// Start creates and starts the prometheus receiver.
func (rcvr iisReceiver) Start(ctx context.Context, host component.Host) error {
	watchers := []winperfcounters.PerfCounterWatcher{}
	for _, objCfg := range getScraperCfgs() {
		objWatchers, err := objCfg.BuildPaths()
		if err != nil {
			rcvr.params.Logger.Warn("some performance counters could not be initialized", zap.Error(err))
		}
		for _, objWatcher := range objWatchers {
			watchers = append(watchers, objWatcher)
		}
	}
	rcvr.watchers = watchers

	return nil
}

func (rcvr *iisReceiver) scrape(ctx context.Context) (pdata.Metrics, error) {
	var errs error

	watchedMetrics := []winperfcounters.CounterValue{}
	for _, watcher := range rcvr.watchers {
		counterValue, err := watcher.ScrapeData()
		if err != nil {
			errs = multierr.Append(errs, err)
			continue
		}
		watchedMetrics = append(watchedMetrics, counterValue)
	}

	metricBuilder := metadata.NewMetricsBuilder(rcvr.config.metricSettings)
	now := pdata.NewTimestampFromTime(time.Now())

	for _, watchedMetric := range watchedMetrics {
		metricBuilder.RecordAny(now, watchedMetric.Value, watchedMetric.MetricRep.Name, watchedMetric.MetricRep.Attributes)
	}

	return metricBuilder.Emit(), errs
}

// Shutdown stops the underlying Prometheus receiver.
func (rcvr iisReceiver) Shutdown(ctx context.Context) error {
	var errs error
	for _, watcher := range rcvr.watchers {
		err := watcher.Close()
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}
