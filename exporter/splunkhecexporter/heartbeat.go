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

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	metricsPrefix              = "otelcol_exporter_splunkhec_"
	defaultHBSentMetricsName   = metricsPrefix + "heartbeat_sent"
	defaultHBFailedMetricsName = metricsPrefix + "heartbeat_failed"
)

type heartbeater struct {
	config     *Config
	pushLogFn  func(ctx context.Context, ld plog.Logs) error
	hbRunOnce  sync.Once
	hbDoneChan chan struct{}

	// Observability
	heartbeatSuccessTotal *stats.Int64Measure
	heartbeatErrorTotal   *stats.Int64Measure
	tagMutators           []tag.Mutator
}

func getMetricsName(overrides map[string]string, metricName string) string {
	if name, ok := overrides[metricName]; ok {
		return name
	}
	return metricName
}

func newHeartbeater(config *Config, pushLogFn func(ctx context.Context, ld plog.Logs) error) *heartbeater {
	interval := config.Heartbeat.Interval
	if interval == 0 {
		return nil
	}

	var heartbeatSent, heartbeatFailed *stats.Int64Measure
	var tagMutators []tag.Mutator
	if config.Telemetry.Enabled {
		overrides := config.Telemetry.OverrideMetricsNames
		extraAttributes := config.Telemetry.ExtraAttributes
		var tags []tag.Key
		tagMutators = []tag.Mutator{}
		for key, val := range extraAttributes {
			newTag, _ := tag.NewKey(key)
			tags = append(tags, newTag)
			tagMutators = append(tagMutators, tag.Insert(newTag, val))
		}

		heartbeatSent = stats.Int64(
			getMetricsName(overrides, defaultHBSentMetricsName),
			"number of heartbeats sent",
			stats.UnitDimensionless)

		heartbeatSentView := &view.View{
			Name:        heartbeatSent.Name(),
			Description: heartbeatSent.Description(),
			TagKeys:     tags,
			Measure:     heartbeatSent,
			Aggregation: view.Sum(),
		}

		heartbeatFailed = stats.Int64(
			getMetricsName(overrides, defaultHBFailedMetricsName),
			"number of heartbeats failed",
			stats.UnitDimensionless)

		heartbeatFailedView := &view.View{
			Name:        heartbeatFailed.Name(),
			Description: heartbeatFailed.Description(),
			TagKeys:     tags,
			Measure:     heartbeatFailed,
			Aggregation: view.Sum(),
		}

		if err := view.Register(heartbeatSentView, heartbeatFailedView); err != nil {
			return nil
		}
	}

	return &heartbeater{
		config:                config,
		pushLogFn:             pushLogFn,
		hbDoneChan:            make(chan struct{}),
		heartbeatSuccessTotal: heartbeatSent,
		heartbeatErrorTotal:   heartbeatFailed,
		tagMutators:           tagMutators,
	}
}

func (h *heartbeater) shutdown() {
	close(h.hbDoneChan)
}

func (h *heartbeater) initHeartbeat(buildInfo component.BuildInfo) {
	interval := h.config.Heartbeat.Interval
	if interval == 0 {
		return
	}

	h.hbRunOnce.Do(func() {
		heartbeatLog := h.generateHeartbeatLog(buildInfo)
		go func() {
			ticker := time.NewTicker(interval)
			for {
				select {
				case <-h.hbDoneChan:
					return
				case <-ticker.C:
					err := h.pushLogFn(context.Background(), heartbeatLog)
					h.observe(err)
				}
			}
		}()
	})
}

// there is only use case for open census metrics recording for now. Extend to use open telemetry in the future.
func (h *heartbeater) observe(err error) {
	if !h.config.Telemetry.Enabled {
		return
	}

	var counter *stats.Int64Measure
	if err == nil {
		counter = h.heartbeatSuccessTotal
	} else {
		counter = h.heartbeatErrorTotal
	}
	_ = stats.RecordWithTags(context.Background(), h.tagMutators, counter.M(1))
}

func (h *heartbeater) generateHeartbeatLog(buildInfo component.BuildInfo) plog.Logs {
	host, err := os.Hostname()
	if err != nil {
		host = "unknownhost"
	}

	ret := plog.NewLogs()
	resourceLogs := ret.ResourceLogs().AppendEmpty()

	resourceAttrs := resourceLogs.Resource().Attributes()
	resourceAttrs.PutStr(h.config.HecToOtelAttrs.Index, "_internal")
	resourceAttrs.PutStr(h.config.HecToOtelAttrs.Source, "otelcol")
	resourceAttrs.PutStr(h.config.HecToOtelAttrs.SourceType, "heartbeat")
	resourceAttrs.PutStr(h.config.HecToOtelAttrs.Host, host)

	logRecord := resourceLogs.ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
	logRecord.Body().SetStr(fmt.Sprintf(
		"HeartbeatInfo version=%s description=%s os=%s arch=%s",
		buildInfo.Version,
		buildInfo.Description,
		runtime.GOOS,
		runtime.GOARCH))
	return ret
}
