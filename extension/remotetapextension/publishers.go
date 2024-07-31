// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapextension

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension/internal/marshaler"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/remotetap"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// Publisher formats incoming data from components into text and broadcasts it to observers.
type Publisher struct {
	logger           *zap.Logger
	callbackManager  *CallbackManager
	logsMarshaler    plog.Marshaler
	metricsMarshaler pmetric.Marshaler
	tracesMarshaler  ptrace.Marshaler
}

func NewPublisher(logger *zap.Logger, cm *CallbackManager) *Publisher {
	return &Publisher{
		logger:           logger,
		callbackManager:  cm,
		metricsMarshaler: marshaler.NewTextMetricsMarshaler(),
		logsMarshaler:    marshaler.NewTextLogsMarshaler(),
		tracesMarshaler:  marshaler.NewTextTracesMarshaler(),
	}
}

func (p *Publisher) PublishMetrics(componentID remotetap.ComponentID, md pmetric.Metrics) {
	data, err := p.metricsMarshaler.MarshalMetrics(md)
	if err != nil {
		p.logger.Warn("could not marshal metrics to text", zap.String("componentID", string(componentID)))
		return
	}
	p.callbackManager.Broadcast(componentID, string(data))
}

func (p *Publisher) PublishTraces(componentID remotetap.ComponentID, td ptrace.Traces) {
	data, err := p.tracesMarshaler.MarshalTraces(td)
	if err != nil {
		p.logger.Warn("could not marshal traces to text", zap.String("componentID", string(componentID)))
		return
	}
	p.callbackManager.Broadcast(componentID, string(data))
}

func (p *Publisher) PublishLogs(componentID remotetap.ComponentID, ld plog.Logs) {
	data, err := p.logsMarshaler.MarshalLogs(ld)
	if err != nil {
		p.logger.Warn("could not marshal logs to text", zap.String("componentID", string(componentID)))
		return
	}
	p.callbackManager.Broadcast(componentID, string(data))
}

func (p *Publisher) PublishData(componentID remotetap.ComponentID, data string) {
	p.callbackManager.Broadcast(componentID, data)
}
