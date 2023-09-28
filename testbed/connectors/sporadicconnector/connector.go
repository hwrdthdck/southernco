package sporadicconnector

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type baseConnector struct {
	component.StartFunc
	component.ShutdownFunc
}

// Capabilities implements the consumer interface.
func (c *baseConnector) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

type metricsConnector struct {
	baseConnector

	config          *Config
	metricsConsumer consumer.Metrics
}

func newMetricsConnector(logger *zap.Logger, config component.Config) *metricsConnector {
	cfg := config.(*Config)
	return &metricsConnector{
		config:        cfg,
		baseConnector: baseConnector{},
	}
}

func (mc *metricsConnector) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	switch mc.config.Decision {
	case 1:
		if err := randomNonPermanentError(); err != nil {
			return err
		}
	case 2:
		if err := randomPermanentError(); err != nil {
			return err
		}
	}
	return mc.metricsConsumer.ConsumeMetrics(ctx, md)
}

type logsConnector struct {
	baseConnector

	config       *Config
	logsConsumer consumer.Logs
}

func newLogsConnector(logger *zap.Logger, config component.Config) *logsConnector {
	cfg := config.(*Config)
	return &logsConnector{
		config:        cfg,
		baseConnector: baseConnector{},
	}
}

func (lc *logsConnector) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	switch lc.config.Decision {
	case 1:
		if err := randomNonPermanentError(); err != nil {
			return err
		}
	case 2:
		if err := randomPermanentError(); err != nil {
			return err
		}
	}
	return lc.logsConsumer.ConsumeLogs(ctx, ld)
}

type traceConnector struct {
	baseConnector

	config         *Config
	tracesConsumer consumer.Traces
}

func newTracesConnector(logger *zap.Logger, config component.Config) *traceConnector {
	cfg := config.(*Config)

	return &traceConnector{
		config:        cfg,
		baseConnector: baseConnector{},
	}
}

func (tc *traceConnector) ConsumeTraces(ctx context.Context, tr ptrace.Traces) error {
	switch tc.config.Decision {
	case 1:
		if err := randomNonPermanentError(); err != nil {
			return err
		}
	case 2:
		if err := randomPermanentError(); err != nil {
			return err
		}
	}
	return tc.tracesConsumer.ConsumeTraces(ctx, tr)
}
