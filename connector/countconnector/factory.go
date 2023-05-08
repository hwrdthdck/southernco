// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package countconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

const (
	typeStr   = "count"
	stability = component.StabilityLevelDevelopment
)

// NewFactory returns a ConnectorFactory.
func NewFactory() connector.Factory {
	return connector.NewFactory(
		typeStr,
		createDefaultConfig,
		connector.WithTracesToMetrics(createTracesToMetrics, component.StabilityLevelDevelopment),
		connector.WithMetricsToMetrics(createMetricsToMetrics, component.StabilityLevelDevelopment),
		connector.WithLogsToMetrics(createLogsToMetrics, component.StabilityLevelDevelopment),
	)
}

// createDefaultConfig creates the default configuration.
func createDefaultConfig() component.Config {
	return &Config{}
}

// createTracesToMetrics creates a traces to metrics connector based on provided config.
func createTracesToMetrics(
	_ context.Context,
	set connector.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Traces, error) {
	c := cfg.(*Config)

	spanMetricDefs := make(map[string]metricDef[ottlspan.TransformContext], len(c.Spans))
	for name, info := range c.Spans {
		md := metricDef[ottlspan.TransformContext]{
			desc:  info.Description,
			attrs: info.Attributes,
		}
		if len(info.Conditions) > 0 {
			// Error checked in Config.Validate()
			condition, _ := filterottl.NewBoolExprForSpan(info.Conditions, filterottl.StandardSpanFuncs(), ottl.PropagateError, set.TelemetrySettings)
			md.condition = condition
		}
		spanMetricDefs[name] = md
	}

	spanEventMetricDefs := make(map[string]metricDef[ottlspanevent.TransformContext], len(c.SpanEvents))
	for name, info := range c.SpanEvents {
		md := metricDef[ottlspanevent.TransformContext]{
			desc:  info.Description,
			attrs: info.Attributes,
		}
		if len(info.Conditions) > 0 {
			// Error checked in Config.Validate()
			condition, _ := filterottl.NewBoolExprForSpanEvent(info.Conditions, filterottl.StandardSpanEventFuncs(), ottl.PropagateError, set.TelemetrySettings)
			md.condition = condition
		}
		spanEventMetricDefs[name] = md
	}

	return &count{
		metricsConsumer:      nextConsumer,
		spansMetricDefs:      spanMetricDefs,
		spanEventsMetricDefs: spanEventMetricDefs,
	}, nil
}

// createMetricsToMetrics creates a metricds to metrics connector based on provided config.
func createMetricsToMetrics(
	_ context.Context,
	set connector.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Metrics, error) {
	c := cfg.(*Config)

	metricMetricDefs := make(map[string]metricDef[ottlmetric.TransformContext], len(c.Metrics))
	for name, info := range c.Metrics {
		md := metricDef[ottlmetric.TransformContext]{
			desc: info.Description,
		}
		if len(info.Conditions) > 0 {
			// Error checked in Config.Validate()
			condition, _ := filterottl.NewBoolExprForMetric(info.Conditions, filterottl.StandardMetricFuncs(), ottl.PropagateError, set.TelemetrySettings)
			md.condition = condition
		}
		metricMetricDefs[name] = md
	}

	dataPointMetricDefs := make(map[string]metricDef[ottldatapoint.TransformContext], len(c.DataPoints))
	for name, info := range c.DataPoints {
		md := metricDef[ottldatapoint.TransformContext]{
			desc:  info.Description,
			attrs: info.Attributes,
		}
		if len(info.Conditions) > 0 {
			// Error checked in Config.Validate()
			condition, _ := filterottl.NewBoolExprForDataPoint(info.Conditions, filterottl.StandardDataPointFuncs(), ottl.PropagateError, set.TelemetrySettings)
			md.condition = condition
		}
		dataPointMetricDefs[name] = md
	}

	return &count{
		metricsConsumer:      nextConsumer,
		metricsMetricDefs:    metricMetricDefs,
		dataPointsMetricDefs: dataPointMetricDefs,
	}, nil
}

// createLogsToMetrics creates a logs to metrics connector based on provided config.
func createLogsToMetrics(
	_ context.Context,
	set connector.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Logs, error) {
	c := cfg.(*Config)

	metricDefs := make(map[string]metricDef[ottllog.TransformContext], len(c.Logs))
	for name, info := range c.Logs {
		md := metricDef[ottllog.TransformContext]{
			desc:  info.Description,
			attrs: info.Attributes,
		}
		if len(info.Conditions) > 0 {
			// Error checked in Config.Validate()
			condition, _ := filterottl.NewBoolExprForLog(info.Conditions, filterottl.StandardLogFuncs(), ottl.PropagateError, set.TelemetrySettings)
			md.condition = condition
		}
		metricDefs[name] = md
	}

	return &count{
		metricsConsumer: nextConsumer,
		logsMetricDefs:  metricDefs,
	}, nil
}

type metricDef[K any] struct {
	condition expr.BoolExpr[K]
	desc      string
	attrs     []AttributeConfig
}
