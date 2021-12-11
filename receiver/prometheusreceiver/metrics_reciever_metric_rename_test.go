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

package prometheusreceiver

import (
	"testing"

	"github.com/prometheus/common/model"
	promcfg "github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/pkg/relabel"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/model/pdata"
)

var renamingLabel = `
# HELP http_go_threads Number of OS threads created
# TYPE http_go_threads gauge
http_go_threads 19

# HELP http_connected_total connected clients
# TYPE http_connected_total counter
http_connected_total{url="localhost",status="ok"} 15.0

# HELP redis_http_requests_total Redis connected clients
# TYPE redis_http_requests_total counter
redis_http_requests_total{method="post",port="6380"} 10.0
redis_http_requests_total{job="sample-app",statusCode="200"} 12.0

# HELP rpc_duration_total RPC clients
# TYPE rpc_duration_total counter
rpc_duration_total{monitor="codeLab",host="local"} 100.0
rpc_duration_total{address="localhost:9090/metrics",contentType="application/json"} 120.0
`

func TestLabelRenaming(t *testing.T) {
	targets := []*testData{
		{
			name: "target1",
			pages: []mockPrometheusResponse{
				{code: 200, data: renamingLabel},
			},
			validateFunc: verifyRenameLabel,
		},
	}

	testComponent(t, targets, false, "", func(cfg *promcfg.Config) {
		for _, scrapeConfig := range cfg.ScrapeConfigs {
			scrapeConfig.MetricRelabelConfigs = []*relabel.Config{
				{
					// this config should add new label {foo="bar"} to all metrics'
					Regex:       relabel.MustNewRegexp("(.*)"),
					Action:      relabel.Replace,
					TargetLabel: "foo",
					Replacement: "bar",
				},
				{
					// this config should create new label {id="target1/metrics"}
					// using the value from capture group of matched regex
					SourceLabels: model.LabelNames{"address"},
					Regex:        relabel.MustNewRegexp(".*/(.*)"),
					Action:       relabel.Replace,
					TargetLabel:  "id",
					Replacement:  "$1",
				},
				{
					// this config creates a new label for metrics that has matched regex label.
					// They key of this new label will be as given in 'replacement'
					// and value will be of the matched regex label value.
					Regex:       relabel.MustNewRegexp("method(.*)"),
					Action:      relabel.LabelMap,
					Replacement: "bar$1",
				},
				{
					// this config should drop the matched regex label
					Regex:  relabel.MustNewRegexp("(url.*)"),
					Action: relabel.LabelDrop,
				},
			}
		}
	})

}

func verifyRenameLabel(t *testing.T, td *testData, resourceMetrics []*pdata.ResourceMetrics) {
	verifyNumValidScrapeResults(t, td, resourceMetrics)
	m1 := resourceMetrics[0]

	// m1 has 4 metrics + 5 internal scraper metrics
	assert.Equal(t, 9, metricsCount(m1))

	wantAttributes := td.attributes

	metrics1 := m1.InstrumentationLibraryMetrics().At(0).Metrics()
	ts1 := getTS(metrics1)
	e1 := []testExpectation{
		assertMetricPresent("http_go_threads",
			compareMetricType(pdata.MetricDataTypeGauge),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(19),
						compareAttributes(map[string]string{"foo": "bar"}),
					},
				},
			}),
		assertMetricPresent("http_connected_total",
			compareMetricType(pdata.MetricDataTypeSum),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(15),
						compareAttributes(map[string]string{"foo": "bar", "status": "ok"}),
					},
				},
			}),
		assertMetricPresent("redis_http_requests_total",
			compareMetricType(pdata.MetricDataTypeSum),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(10),
						compareAttributes(map[string]string{"method": "post", "port": "6380", "bar": "post", "foo": "bar"}),
					},
				},
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(12),
						// since honor_label bool in config is true by default,
						// Prometheus reserved keywords like "job" and "instance" should be prefixed by "exported_"
						compareAttributes(map[string]string{"exported_job": "sample-app", "statusCode": "200", "foo": "bar"}),
					},
				},
			}),
		assertMetricPresent("rpc_duration_total",
			compareMetricType(pdata.MetricDataTypeSum),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(100),
						compareAttributes(map[string]string{"monitor": "codeLab", "host": "local", "foo": "bar"}),
					},
				},
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(120),
						compareAttributes(map[string]string{"address": "localhost:9090/metrics",
							"contentType": "application/json", "id": "metrics", "foo": "bar"}),
					},
				},
			}),
	}
	doCompare(t, "scrape-labelRename-1", wantAttributes, m1, e1)
}

func TestLabelRenamingKeepAction(t *testing.T) {
	targets := []*testData{
		{
			name: "target1",
			pages: []mockPrometheusResponse{
				{code: 200, data: renamingLabel},
			},
			validateFunc: verifyRenameLabelKeepAction,
		},
	}

	testComponent(t, targets, false, "", func(cfg *promcfg.Config) {
		for _, scrapeConfig := range cfg.ScrapeConfigs {
			scrapeConfig.MetricRelabelConfigs = []*relabel.Config{
				{
					// this config should keep only metric that matches the regex metric name, and drop the rest
					Regex: relabel.MustNewRegexp("__name__|__scheme__|__address__|" +
						"__metrics_path__|__scrape_interval__|instance|job|(m.*)"),
					Action: relabel.LabelKeep,
				},
			}
		}
	})

}

func verifyRenameLabelKeepAction(t *testing.T, td *testData, resourceMetrics []*pdata.ResourceMetrics) {
	verifyNumValidScrapeResults(t, td, resourceMetrics)
	m1 := resourceMetrics[0]

	// m1 has 4 metrics + 5 internal scraper metrics
	assert.Equal(t, 9, metricsCount(m1))

	wantAttributes := td.attributes

	metrics1 := m1.InstrumentationLibraryMetrics().At(0).Metrics()
	ts1 := getTS(metrics1)
	e1 := []testExpectation{
		assertMetricPresent("http_go_threads",
			compareMetricType(pdata.MetricDataTypeGauge),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(19),
						assertAttributesAbsent(),
					},
				},
			}),
		assertMetricPresent("http_connected_total",
			compareMetricType(pdata.MetricDataTypeSum),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(15),
						assertAttributesAbsent(),
					},
				},
			}),
		assertMetricPresent(" Redis connected clients",
			compareMetricType(pdata.MetricDataTypeSum),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(10),
						compareAttributes(map[string]string{"method": "post"}),
					},
				},
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(12),
						assertAttributesAbsent(),
					},
				},
			}),
		assertMetricPresent("RPC clients",
			compareMetricType(pdata.MetricDataTypeSum),
			[]dataPointExpectation{
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(100),
						compareAttributes(map[string]string{"monitor": "codeLab"}),
					},
				},
				{
					numberPointComparator: []numberPointComparator{
						compareTimestamp(ts1),
						compareDoubleValue(120),
						assertAttributesAbsent(),
					},
				},
			}),
	}
	doCompare(t, "scrape-LabelRenameKeepAction-1", wantAttributes, m1, e1)
}
