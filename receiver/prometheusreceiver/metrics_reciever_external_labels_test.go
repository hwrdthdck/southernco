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
	"context"
	"testing"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/model/pdata"
)

const targetExternalLabels = `
# HELP go_threads Number of OS threads created
# TYPE go_threads gauge
go_threads 19`

func TestExternalLabels(t *testing.T) {
	ctx := context.Background()
	targets := []*testData{
		{
			name: "target1",
			pages: []mockPrometheusResponse{
				{code: 200, data: targetExternalLabels},
			},
			validateFunc: verifyExternalLabels,
		},
	}

	mp, cfg, err := setupMockPrometheus(targets...)
	cfg.GlobalConfig.ExternalLabels = labels.FromStrings("key", "value")
	require.Nilf(t, err, "Failed to create Promtheus config: %v", err)
	defer mp.Close()

	cms := new(consumertest.MetricsSink)
	receiver := newPrometheusReceiver(logger, &Config{
		ReceiverSettings: config.NewReceiverSettings(config.NewID(typeStr)),
		PrometheusConfig: cfg}, cms)

	require.NoError(t, receiver.Start(ctx, componenttest.NewNopHost()), "Failed to invoke Start: %v", err)
	t.Cleanup(func() { require.NoError(t, receiver.Shutdown(ctx)) })

	mp.wg.Wait()
	metrics := cms.AllMetrics()

	// split and store results by target name
	results := make(map[string][]pdata.Metric)
	for _, md := range metrics {
		rms := md.ResourceMetrics()
		for i := 0; i < rms.Len(); i++ {
			rmi := rms.At(i)
			serviceNameAttr, ok := rmi.Resource().Attributes().Get("service.name")
			assert.True(t, ok, `expected "service.name" as a known attribute`)
			serviceName := serviceNameAttr.StringVal()
			results[serviceName] = append(results[serviceName], md)
		}
	}
	for _, target := range targets {
		target.validateFunc(t, target, results[target.name])
	}
}

func makeMetric(td *testData, metricL []pdata.Metric) pdata.ResourceMetrics {
	rms := pdata.NewResourceMetrics()
	td.resource.CopyTo(rms.Resource())
	ilm := rms.InstrumentationLibraryMetrics().AppendEmpty()
	destMetricL := ilm.Metrics()
	for _, metric := range metricL {
		destMetric := destMetricL.AppendMetric()
		metric.CopyTo(destMetric)
	}
	return rms
}

func verifyExternalLabels(t *testing.T, td *testData, mds []pdata.Metric) {
	verifyNumScrapeResults(t, td, mds)

        want := makeMetric(td, []pdata.Metric{
            internal.GaugeMetric("go_threads", k1v2, t1Ms, internal.DoublePoint(t1Ms, 19))
        })

	// TODO: Translate me.
	/*
		want := &agentmetricspb.ExportMetricsServiceRequest{
			Node:     td.node,
			Resource: td.resource,
		}
		doCompare("scrape-externalLabels", t, want, mds[0], []testExpectation{
			assertMetricPresent("go_threads",
				[]descriptorComparator{
					compareMetricType(metricspb.MetricDescriptor_GAUGE_DOUBLE),
					compareMetricLabelKeys([]string{"key"}),
				},
				[]seriesExpectation{
					{
						series: []seriesComparator{
							compareSeriesLabelValues([]string{"value"}),
						},
						points: []pointComparator{
							comparePointTimestamp(mds[0].Metrics[0].Timeseries[0].Points[0].Timestamp),
							compareDoubleVal(19),
						},
					},
				}),
		})
	*/
}
