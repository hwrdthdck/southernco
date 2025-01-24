// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/otlpmetricstream"
)

const defaultMetricsEncoding = cwmetricstream.TypeStr

// The metricsConsumer implements the firehoseConsumer
// to use a metrics consumer and unmarshaler.
type metricsConsumer struct {
	config   *Config
	settings receiver.Settings
	// consumer passes the translated metrics on to the
	// next consumer.
	consumer consumer.Metrics
	// unmarshaler is the configured pmetric.Unmarshaler
	// to use when processing the records.
	unmarshaler pmetric.Unmarshaler
}

var _ firehoseConsumer = (*metricsConsumer)(nil)

// newMetricsReceiver creates a new instance of the receiver
// with a metricsConsumer.
func newMetricsReceiver(
	config *Config,
	set receiver.Settings,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	c := &metricsConsumer{
		config:   config,
		settings: set,
		consumer: nextConsumer,
	}
	return &firehoseReceiver{
		settings: set,
		config:   config,
		consumer: c,
	}, nil
}

func (c *metricsConsumer) Start(_ context.Context, host component.Host) error {
	encoding := c.config.Encoding
	if encoding == "" {
		encoding = c.config.RecordType
		if encoding == "" {
			encoding = defaultMetricsEncoding
		}
	}
	switch encoding {
	case cwmetricstream.TypeStr:
		// TODO: make cwmetrics an encoding extension
		c.unmarshaler = cwmetricstream.NewUnmarshaler(c.settings.Logger)
	case otlpmetricstream.TypeStr:
		// TODO: make otlp_v1 an encoding extension
		c.unmarshaler = otlpmetricstream.NewUnmarshaler(c.settings.Logger)
	default:
		unmarshaler, err := loadEncodingExtension[pmetric.Unmarshaler](host, encoding, "metrics")
		if err != nil {
			return err
		}
		c.unmarshaler = unmarshaler
	}
	return nil
}

// Consume uses the configured unmarshaler to deserialize the records into a
// single pmetric.Metrics. If there are common attributes available, then it will
// attach those to each of the pcommon.Resources. It will send the final result
// to the next consumer.
func (c *metricsConsumer) Consume(ctx context.Context, records [][]byte, commonAttributes map[string]string) (int, error) {
	merged := metricsMerger{resources: make(map[identity.Resource]resourceMetricsMerger)}
	for _, record := range records {
		metrics, err := c.unmarshaler.UnmarshalMetrics(record)
		if err != nil {
			return http.StatusBadRequest, err
		}
		if err := merged.merge(metrics); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	metrics := pmetric.NewMetrics()
	for _, rm := range merged.resources {
		resourceMetrics := metrics.ResourceMetrics().AppendEmpty()
		rm.resource.MoveTo(resourceMetrics.Resource())
		resourceMetrics.SetSchemaUrl(rm.schemaURL)
		for _, sm := range rm.scopes {
			scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
			sm.scope.MoveTo(scopeMetrics.Scope())
			scopeMetrics.SetSchemaUrl(sm.schemaURL)
			for _, metric := range sm.metrics {
				metric.MoveTo(scopeMetrics.Metrics().AppendEmpty())
			}
		}
	}

	if commonAttributes != nil {
		for i := 0; i < metrics.ResourceMetrics().Len(); i++ {
			rm := metrics.ResourceMetrics().At(i)
			for k, v := range commonAttributes {
				if _, found := rm.Resource().Attributes().Get(k); !found {
					rm.Resource().Attributes().PutStr(k, v)
				}
			}
		}
	}

	if err := c.consumer.ConsumeMetrics(ctx, metrics); err != nil {
		if consumererror.IsPermanent(err) {
			return http.StatusBadRequest, err
		}
		return http.StatusServiceUnavailable, err
	}
	return http.StatusOK, nil
}

type metricsMerger struct {
	resources map[identity.Resource]resourceMetricsMerger
}

type resourceMetricsMerger struct {
	resource  pcommon.Resource
	schemaURL string
	scopes    map[identity.Scope]scopeMetricsMerger
}

type scopeMetricsMerger struct {
	scope     pcommon.InstrumentationScope
	schemaURL string
	metrics   map[identity.Metric]pmetric.Metric
}

func (m *metricsMerger) merge(metrics pmetric.Metrics) error {
	for i := 0; i < metrics.ResourceMetrics().Len(); i++ {
		resourceMetrics := metrics.ResourceMetrics().At(i)
		resourceIdentity := identity.OfResource(resourceMetrics.Resource())
		rm, ok := m.resources[resourceIdentity]
		if !ok {
			rm = resourceMetricsMerger{
				resource:  resourceMetrics.Resource(),
				schemaURL: resourceMetrics.SchemaUrl(),
				scopes:    make(map[identity.Scope]scopeMetricsMerger),
			}
			m.resources[resourceIdentity] = rm
		}
		if err := rm.merge(resourceMetrics.ScopeMetrics()); err != nil {
			return err
		}
	}
	return nil
}

func (rm *resourceMetricsMerger) merge(scopeMetricsSlice pmetric.ScopeMetricsSlice) error {
	for i := 0; i < scopeMetricsSlice.Len(); i++ {
		scopeMetrics := scopeMetricsSlice.At(i)
		scopeIdentity := identity.OfScope(identity.Resource{}, scopeMetrics.Scope())
		sm, ok := rm.scopes[scopeIdentity]
		if !ok {
			sm = scopeMetricsMerger{
				scope:     scopeMetrics.Scope(),
				schemaURL: scopeMetrics.SchemaUrl(),
				metrics:   make(map[identity.Metric]pmetric.Metric),
			}
			rm.scopes[scopeIdentity] = sm
		}
		if err := sm.merge(scopeMetrics.Metrics()); err != nil {
			return err
		}
	}
	return nil
}

func (sm *scopeMetricsMerger) merge(metricSlice pmetric.MetricSlice) error {
	for i := 0; i < metricSlice.Len(); i++ {
		metric := metricSlice.At(i)
		metricIdentity := identity.OfMetric(identity.Scope{}, metric)
		existingMetric, ok := sm.metrics[metricIdentity]
		if !ok {
			sm.metrics[metricIdentity] = metric
		} else {
			if err := moveAppendDataPoints(metric, existingMetric); err != nil {
				return err
			}
		}
	}
	return nil
}

func moveAppendDataPoints(from, to pmetric.Metric) error {
	switch from.Type() {
	case pmetric.MetricTypeSummary:
		from.Summary().DataPoints().MoveAndAppendTo(to.Summary().DataPoints())
	case pmetric.MetricTypeSum:
		from.Sum().DataPoints().MoveAndAppendTo(to.Sum().DataPoints())
	case pmetric.MetricTypeGauge:
		from.Gauge().DataPoints().MoveAndAppendTo(to.Gauge().DataPoints())
	case pmetric.MetricTypeHistogram:
		from.Histogram().DataPoints().MoveAndAppendTo(to.Histogram().DataPoints())
	case pmetric.MetricTypeExponentialHistogram:
		from.ExponentialHistogram().DataPoints().MoveAndAppendTo(to.ExponentialHistogram().DataPoints())
	default:
		return fmt.Errorf("unhandled metric type %q", from.Type())
	}
	return nil
}
