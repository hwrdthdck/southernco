// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/unmarshalertest"
)

type metricsRecordConsumer struct {
	result pmetric.Metrics
}

var _ consumer.Metrics = (*metricsRecordConsumer)(nil)

func (rc *metricsRecordConsumer) ConsumeMetrics(_ context.Context, metrics pmetric.Metrics) error {
	rc.result = metrics
	return nil
}

func (rc *metricsRecordConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func TestMetricsReceiver_Start(t *testing.T) {
	testCases := map[string]struct {
		encoding            string
		recordType          string
		wantUnmarshalerType pmetric.Unmarshaler
		wantErr             string
	}{
		"WithDefaultEncoding": {
			wantUnmarshalerType: &cwmetricstream.Unmarshaler{},
		},
		"WithBuiltinEncoding": {
			encoding:            "cwmetrics",
			wantUnmarshalerType: &cwmetricstream.Unmarshaler{},
		},
		"WithExtensionEncoding": {
			encoding:            "otlp_metrics",
			wantUnmarshalerType: pmetricUnmarshalerExtension{},
		},
		"WithDeprecatedRecordType": {
			recordType:          "otlp_metrics",
			wantUnmarshalerType: pmetricUnmarshalerExtension{},
		},
		"WithUnknownEncoding": {
			encoding: "invalid",
			wantErr:  `unknown encoding extension "invalid"`,
		},
		"WithNonLogUnmarshalerExtension": {
			encoding: "otlp_logs",
			wantErr:  `extension "otlp_logs" is not a metrics unmarshaler`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			cfg := createDefaultConfig().(*Config)
			cfg.Encoding = testCase.encoding
			cfg.RecordType = testCase.recordType
			got, err := newMetricsReceiver(
				cfg,
				receivertest.NewNopSettings(),
				consumertest.NewNop(),
			)
			require.NoError(t, err)
			require.NotNil(t, got)
			t.Cleanup(func() {
				require.NoError(t, got.Shutdown(context.Background()))
			})

			host := hostWithExtensions{
				extensions: map[component.ID]component.Component{
					component.MustNewID("otlp_logs"):    plogUnmarshalerExtension{},
					component.MustNewID("otlp_metrics"): pmetricUnmarshalerExtension{},
				},
			}

			err = got.Start(context.Background(), host)
			if testCase.wantErr != "" {
				require.EqualError(t, err, testCase.wantErr)
			} else {
				require.NoError(t, err)
			}

			assert.IsType(t,
				testCase.wantUnmarshalerType,
				got.(*firehoseReceiver).consumer.(*metricsConsumer).unmarshaler,
			)
		})
	}
}

func TestMetricsConsumer_Errors(t *testing.T) {
	testErr := errors.New("test error")
	testCases := map[string]struct {
		unmarshalerErr error
		consumerErr    error
		wantStatus     int
		wantErr        error
	}{
		"WithUnmarshalerError": {
			unmarshalerErr: testErr,
			wantStatus:     http.StatusBadRequest,
			wantErr:        testErr,
		},
		"WithConsumerErrorPermanent": {
			consumerErr: consumererror.NewPermanent(testErr),
			wantStatus:  http.StatusBadRequest,
			wantErr:     consumererror.NewPermanent(testErr),
		},
		"WithConsumerError": {
			consumerErr: testErr,
			wantStatus:  http.StatusServiceUnavailable,
			wantErr:     testErr,
		},
		"WithNoError": {
			wantStatus: http.StatusOK,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			mc := &metricsConsumer{
				unmarshaler: unmarshalertest.NewErrMetrics(testCase.unmarshalerErr),
				consumer:    consumertest.NewErr(testCase.consumerErr),
			}
			gotStatus, gotErr := mc.Consume(context.Background(), [][]byte{{0}}, nil)
			require.Equal(t, testCase.wantStatus, gotStatus)
			require.Equal(t, testCase.wantErr, gotErr)
		})
	}
}

func TestMetricsConsumer(t *testing.T) {
	t.Run("WithCommonAttributes", func(t *testing.T) {
		base := pmetric.NewMetrics()
		base.ResourceMetrics().AppendEmpty()
		rc := metricsRecordConsumer{}
		mc := &metricsConsumer{
			unmarshaler: unmarshalertest.NewWithMetrics(base),
			consumer:    &rc,
		}
		gotStatus, gotErr := mc.Consume(context.Background(), [][]byte{{0}}, map[string]string{
			"CommonAttributes": "Test",
		})
		require.Equal(t, http.StatusOK, gotStatus)
		require.NoError(t, gotErr)
		gotRms := rc.result.ResourceMetrics()
		require.Equal(t, 1, gotRms.Len())
		gotRm := gotRms.At(0)
		require.Equal(t, 1, gotRm.Resource().Attributes().Len())
	})
	t.Run("WithMultipleRecords", func(t *testing.T) {
		// service0, scope0
		metrics0, metricSlice0 := newMetrics("service0", "scope0")
		sum0unit0 := addDeltaSumMetric(metricSlice0, "sum0", "unit0")
		addSumDataPoint(sum0unit0, 1)

		// service0, [scope0. scope1]
		// these scopes should be merged into the above resource metrics
		metrics1, metricSlice1 := newMetrics("service0", "scope0")
		sum0unit0 = addDeltaSumMetric(metricSlice1, "sum0", "unit0")
		addSumDataPoint(sum0unit0, 2)
		sum0unit1 := addDeltaSumMetric(metricSlice1, "sum0", "unit1")
		addSumDataPoint(sum0unit1, 3)
		scopeMetrics1 := metrics1.ResourceMetrics().At(0).ScopeMetrics().AppendEmpty()
		newScope("scope1").MoveTo(scopeMetrics1.Scope())
		sum0unit0 = addDeltaSumMetric(scopeMetrics1.Metrics(), "sum0", "unit0")
		addSumDataPoint(sum0unit0, 4)

		// service1, scope0
		metrics2, metricSlice2 := newMetrics("service1", "scope0")
		sum0unit0 = addDeltaSumMetric(metricSlice2, "sum0", "unit0")
		addSumDataPoint(sum0unit0, 5)
		sum1unit0 := addDeltaSumMetric(metricSlice2, "sum1", "unit0")
		addSumDataPoint(sum1unit0, 6)

		metricsRemaining := []pmetric.Metrics{metrics0, metrics1, metrics2}
		var unmarshaler unmarshalMetricsFunc = func([]byte) (pmetric.Metrics, error) {
			metrics := metricsRemaining[0]
			metricsRemaining = metricsRemaining[1:]
			return metrics, nil
		}

		rc := metricsRecordConsumer{}
		lc := &metricsConsumer{unmarshaler: unmarshaler, consumer: &rc}
		gotStatus, gotErr := lc.Consume(context.Background(), make([][]byte, len(metricsRemaining)), nil)
		require.Equal(t, http.StatusOK, gotStatus)
		require.NoError(t, gotErr)
		assert.Equal(t, 6, rc.result.DataPointCount())        // data points are not aggregated
		assert.Equal(t, 2, rc.result.ResourceMetrics().Len()) // service0, service1

		type metric struct {
			name       string
			unit       string
			datapoints []int64
		}

		type scopeMetrics struct {
			scope   string
			metrics []metric
		}

		type resourceScopeMetrics struct {
			serviceName string
			scopes      []scopeMetrics
		}

		var merged []resourceScopeMetrics
		for i := 0; i < rc.result.ResourceMetrics().Len(); i++ {
			resourceMetrics := rc.result.ResourceMetrics().At(i)
			serviceName, ok := resourceMetrics.Resource().Attributes().Get("service.name")
			require.True(t, ok)

			var scopes []scopeMetrics
			for i := 0; i < resourceMetrics.ScopeMetrics().Len(); i++ {
				var scope scopeMetrics
				scopeMetrics := resourceMetrics.ScopeMetrics().At(i)
				scope.scope = scopeMetrics.Scope().Name()
				for i := 0; i < scopeMetrics.Metrics().Len(); i++ {
					mi := scopeMetrics.Metrics().At(i)
					m := metric{name: mi.Name(), unit: mi.Unit()}
					for i := 0; i < mi.Sum().DataPoints().Len(); i++ {
						dp := mi.Sum().DataPoints().At(i)
						m.datapoints = append(m.datapoints, dp.IntValue())
					}
					scope.metrics = append(scope.metrics, m)
				}
				scopes = append(scopes, scope)
			}
			merged = append(merged, resourceScopeMetrics{serviceName: serviceName.Str(), scopes: scopes})
		}
		assert.Equal(t, []resourceScopeMetrics{{
			serviceName: "service0",
			scopes: []scopeMetrics{{
				scope: "scope0",
				metrics: []metric{
					{name: "sum0", unit: "unit0", datapoints: []int64{1, 2}},
					{name: "sum0", unit: "unit1", datapoints: []int64{3}},
				},
			}, {
				scope: "scope1",
				metrics: []metric{
					{name: "sum0", unit: "unit0", datapoints: []int64{4}},
				},
			}},
		}, {
			serviceName: "service1",
			scopes: []scopeMetrics{{
				scope: "scope0",
				metrics: []metric{
					{name: "sum0", unit: "unit0", datapoints: []int64{5}},
					{name: "sum1", unit: "unit0", datapoints: []int64{6}},
				},
			}},
		}}, merged)
	})
}

func newMetrics(serviceName, scopeName string) (pmetric.Metrics, pmetric.MetricSlice) {
	metrics := pmetric.NewMetrics()
	resourceMetrics := metrics.ResourceMetrics().AppendEmpty()
	scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
	newResource(serviceName).MoveTo(resourceMetrics.Resource())
	newScope(scopeName).MoveTo(scopeMetrics.Scope())
	return metrics, scopeMetrics.Metrics()
}

func addDeltaSumMetric(metrics pmetric.MetricSlice, name, unit string) pmetric.Sum {
	m := metrics.AppendEmpty()
	m.SetName(name)
	m.SetUnit(unit)
	sum := m.SetEmptySum()
	sum.SetAggregationTemporality(pmetric.AggregationTemporalityDelta)
	return sum
}

func addSumDataPoint(sum pmetric.Sum, value int64) {
	dp := sum.DataPoints().AppendEmpty()
	dp.SetIntValue(value)
}

type unmarshalMetricsFunc func([]byte) (pmetric.Metrics, error)

func (f unmarshalMetricsFunc) UnmarshalMetrics(data []byte) (pmetric.Metrics, error) {
	return f(data)
}
