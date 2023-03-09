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

package countconnector

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/collector/consumer/consumertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
)

func TestTracesToMetrics(t *testing.T) {
	testCases := []struct {
		name string
		cfg  *Config
	}{
		{
			name: "zero_conditions",
			cfg: &Config{
				Spans:      defaultSpansConfig(),
				SpanEvents: defaultSpanEventsConfig(),
			},
		},
		{
			name: "one_condition",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.if": {
						MetricInfo: MetricInfo{
							Description: "Span count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.if": {
						MetricInfo: MetricInfo{
							Description: "Span event count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_conditions",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.if": {
						MetricInfo: MetricInfo{
							Description: "Span count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["span.optional"] != nil`,
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.if": {
						MetricInfo: MetricInfo{
							Description: "Span event count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["event.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_metrics",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.all": {
						MetricInfo: MetricInfo{
							Description: "All spans count",
						},
					},
					"span.count.if": {
						MetricInfo: MetricInfo{
							Description: "Span count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["span.optional"] != nil`,
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.all": {
						MetricInfo: MetricInfo{
							Description: "All span events count",
						},
					},
					"spanevent.count.if": {
						MetricInfo: MetricInfo{
							Description: "Span event count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["event.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "one_attribute",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span count by attribute",
						},
						Attributes: []AttributeConfig{
							{
								Key: "span.required",
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span event count by attribute",
						},
						Attributes: []AttributeConfig{
							{
								Key: "event.required",
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_attributes",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span count by attributes",
						},
						Attributes: []AttributeConfig{
							{
								Key: "span.required",
							},
							{
								Key: "span.optional",
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span event count by attributes",
						},
						Attributes: []AttributeConfig{
							{
								Key: "event.required",
							},
							{
								Key: "event.optional",
							},
						},
					},
				},
			},
		},
		{
			name: "default_attribute_value",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span count by attribute with default",
						},
						Attributes: []AttributeConfig{
							{
								Key: "span.required",
							},
							{
								Key:          "span.optional",
								DefaultValue: "other",
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span event count by attribute with default",
						},
						Attributes: []AttributeConfig{
							{
								Key: "event.required",
							},
							{
								Key:          "event.optional",
								DefaultValue: "other",
							},
						},
					},
				},
			},
		},
		{
			name: "condition_and_attribute",
			cfg: &Config{
				Spans: map[string]MetricInfoWithAttributes{
					"span.count.if.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
						Attributes: []AttributeConfig{
							{
								Key: "span.required",
							},
						},
					},
				},
				SpanEvents: map[string]MetricInfoWithAttributes{
					"spanevent.count.if.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Span event count by attribute if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
						Attributes: []AttributeConfig{
							{
								Key: "event.required",
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, tc.cfg.Validate())
			factory := NewFactory()
			sink := &consumertest.MetricsSink{}
			conn, err := factory.CreateTracesToMetrics(context.Background(),
				connectortest.NewNopCreateSettings(), tc.cfg, sink)
			require.NoError(t, err)
			require.NotNil(t, conn)
			assert.False(t, conn.Capabilities().MutatesData)

			require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
			defer func() {
				assert.NoError(t, conn.Shutdown(context.Background()))
			}()

			testSpans, err := golden.ReadTraces(filepath.Join("testdata", "traces", "input.json"))
			assert.NoError(t, err)
			assert.NoError(t, conn.ConsumeTraces(context.Background(), testSpans))

			allMetrics := sink.AllMetrics()
			assert.Equal(t, 1, len(allMetrics))

			// golden.WriteMetrics(t, filepath.Join("testdata", "traces", tc.name+".json"), allMetrics[0])
			expected, err := golden.ReadMetrics(filepath.Join("testdata", "traces", tc.name+".json"))
			assert.NoError(t, err)
			assert.NoError(t, pmetrictest.CompareMetrics(expected, allMetrics[0],
				pmetrictest.IgnoreTimestamp(),
				pmetrictest.IgnoreResourceMetricsOrder(),
				pmetrictest.IgnoreMetricsOrder(),
				pmetrictest.IgnoreMetricDataPointsOrder()))
		})
	}
}

func TestMetricsToMetrics(t *testing.T) {
	testCases := []struct {
		name string
		cfg  *Config
	}{
		{
			name: "zero_conditions",
			cfg: &Config{
				Metrics:    defaultMetricsConfig(),
				DataPoints: defaultDataPointsConfig(),
			},
		},
		{
			name: "one_condition",
			cfg: &Config{
				Metrics: map[string]MetricInfo{
					"metric.count.if": {
						Description: "Metric count if ...",
						Conditions: []string{
							`resource.attributes["resource.optional"] != nil`,
						},
					},
				},
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.if": {
						MetricInfo: MetricInfo{
							Description: "Data point count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_conditions",
			cfg: &Config{
				Metrics: map[string]MetricInfo{
					"metric.count.if": {
						Description: "Metric count if ...",
						Conditions: []string{
							`resource.attributes["resource.optional"] != nil`,
							`type == METRIC_DATA_TYPE_HISTOGRAM`,
						},
					},
				},
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.if": {
						MetricInfo: MetricInfo{
							Description: "Data point count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["datapoint.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_metrics",
			cfg: &Config{
				Metrics: map[string]MetricInfo{
					"metric.count.all": {
						Description: "All metrics count",
					},
					"metric.count.if": {
						Description: "Metric count if ...",
						Conditions: []string{
							`resource.attributes["resource.optional"] != nil`,
							`type == METRIC_DATA_TYPE_HISTOGRAM`,
						},
					},
				},
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.all": {
						MetricInfo: MetricInfo{
							Description: "All data points count",
						},
					},
					"datapoint.count.if": {
						MetricInfo: MetricInfo{
							Description: "Data point count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["datapoint.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "one_attribute",
			cfg: &Config{
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Data point count by attribute",
						},
						Attributes: []AttributeConfig{
							{
								Key: "datapoint.required",
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_attributes",
			cfg: &Config{
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Data point count by attributes",
						},
						Attributes: []AttributeConfig{
							{
								Key: "datapoint.required",
							},
							{
								Key: "datapoint.optional",
							},
						},
					},
				},
			},
		},
		{
			name: "default_attribute_value",
			cfg: &Config{
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Data point count by attribute with default",
						},
						Attributes: []AttributeConfig{
							{
								Key: "datapoint.required",
							},
							{
								Key:          "datapoint.optional",
								DefaultValue: "other",
							},
						},
					},
				},
			},
		},
		{
			name: "condition_and_attribute",
			cfg: &Config{
				DataPoints: map[string]MetricInfoWithAttributes{
					"datapoint.count.if.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Data point count by attribute if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
						Attributes: []AttributeConfig{
							{
								Key: "datapoint.required",
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, tc.cfg.Validate())
			factory := NewFactory()
			sink := &consumertest.MetricsSink{}
			conn, err := factory.CreateMetricsToMetrics(context.Background(),
				connectortest.NewNopCreateSettings(), tc.cfg, sink)
			require.NoError(t, err)
			require.NotNil(t, conn)
			assert.False(t, conn.Capabilities().MutatesData)

			require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
			defer func() {
				assert.NoError(t, conn.Shutdown(context.Background()))
			}()

			testMetrics, err := golden.ReadMetrics(filepath.Join("testdata", "metrics", "input.json"))
			assert.NoError(t, err)
			assert.NoError(t, conn.ConsumeMetrics(context.Background(), testMetrics))

			allMetrics := sink.AllMetrics()
			assert.Equal(t, 1, len(allMetrics))

			// golden.WriteMetrics(t, filepath.Join("testdata", "metrics", tc.name+".json"), allMetrics[0])
			expected, err := golden.ReadMetrics(filepath.Join("testdata", "metrics", tc.name+".json"))
			assert.NoError(t, err)
			assert.NoError(t, pmetrictest.CompareMetrics(expected, allMetrics[0],
				pmetrictest.IgnoreTimestamp(),
				pmetrictest.IgnoreResourceMetricsOrder(),
				pmetrictest.IgnoreMetricsOrder(),
				pmetrictest.IgnoreMetricDataPointsOrder()))
		})
	}
}

func TestLogsToMetrics(t *testing.T) {
	testCases := []struct {
		name string
		cfg  *Config
	}{
		{
			name: "zero_conditions",
			cfg:  &Config{Logs: defaultLogsConfig()},
		},
		{
			name: "one_condition",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"count.if": {
						MetricInfo: MetricInfo{
							Description: "Count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_conditions",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"count.if": {
						MetricInfo: MetricInfo{
							Description: "Count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
								`attributes["log.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_metrics",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"count.all": {
						MetricInfo: MetricInfo{
							Description: "All logs count",
						},
					},
					"count.if": {
						MetricInfo: MetricInfo{
							Description: "Count if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
					},
				},
			},
		},
		{
			name: "one_attribute",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"log.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Log count by attribute",
						},
						Attributes: []AttributeConfig{
							{
								Key: "log.required",
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_attributes",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"log.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Log count by attributes",
						},
						Attributes: []AttributeConfig{
							{
								Key: "log.required",
							},
							{
								Key: "log.optional",
							},
						},
					},
				},
			},
		},
		{
			name: "default_attribute_value",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"log.count.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Log count by attribute with default",
						},
						Attributes: []AttributeConfig{
							{
								Key: "log.required",
							},
							{
								Key:          "log.optional",
								DefaultValue: "other",
							},
						},
					},
				},
			},
		},
		{
			name: "condition_and_attribute",
			cfg: &Config{
				Logs: map[string]MetricInfoWithAttributes{
					"log.count.if.by_attr": {
						MetricInfo: MetricInfo{
							Description: "Log count by attribute if ...",
							Conditions: []string{
								`resource.attributes["resource.optional"] != nil`,
							},
						},
						Attributes: []AttributeConfig{
							{
								Key: "log.required",
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, tc.cfg.Validate())
			factory := NewFactory()
			sink := &consumertest.MetricsSink{}
			conn, err := factory.CreateLogsToMetrics(context.Background(),
				connectortest.NewNopCreateSettings(), tc.cfg, sink)
			require.NoError(t, err)
			require.NotNil(t, conn)
			assert.False(t, conn.Capabilities().MutatesData)

			require.NoError(t, conn.Start(context.Background(), componenttest.NewNopHost()))
			defer func() {
				assert.NoError(t, conn.Shutdown(context.Background()))
			}()

			testLogs, err := golden.ReadLogs(filepath.Join("testdata", "logs", "input.json"))
			assert.NoError(t, err)
			assert.NoError(t, conn.ConsumeLogs(context.Background(), testLogs))

			allMetrics := sink.AllMetrics()
			assert.Equal(t, 1, len(allMetrics))

			// golden.WriteMetrics(t, filepath.Join("testdata", "logs", tc.name+".json"), allMetrics[0])
			expected, err := golden.ReadMetrics(filepath.Join("testdata", "logs", tc.name+".json"))
			assert.NoError(t, err)
			assert.NoError(t, pmetrictest.CompareMetrics(expected, allMetrics[0],
				pmetrictest.IgnoreTimestamp(),
				pmetrictest.IgnoreResourceMetricsOrder(),
				pmetrictest.IgnoreMetricsOrder(),
				pmetrictest.IgnoreMetricDataPointsOrder()))
		})
	}
}
