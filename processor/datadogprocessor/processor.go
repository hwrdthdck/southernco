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

package datadogprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogprocessor"

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-agent/pkg/otlp/model/translator"
	"github.com/DataDog/datadog-agent/pkg/trace/pb"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type processor struct {
	logger       *zap.Logger
	nextConsumer consumer.Traces
	cfg          *Config

	// metricsExporter specifies the metrics exporter used to exporter APM Stats
	// as metrics through
	metricsExporter component.MetricsExporter

	// agent specifies the agent used to ingest traces and output APM Stats.
	// It is implemented by the traceagent structure; replaced in tests.
	agent ingester

	// translator specifies the translator used to transform APM Stats Payloads
	// from the agent to OTLP Metrics.
	translator *translator.Translator

	// in specifies the channel through which the agent will output Stats Payloads
	// resulting from ingested traces.
	in chan pb.StatsPayload

	// exit specifies the exit channel, which will be closed upon shutdown.
	exit chan struct{}
}

func newProcessor(ctx context.Context, logger *zap.Logger, config component.Config, nextConsumer consumer.Traces) (*processor, error) {
	cfg := config.(*Config)
	in := make(chan pb.StatsPayload, 100)
	trans, err := translator.New(logger)
	if err != nil {
		return nil, err
	}
	return &processor{
		logger:       logger,
		nextConsumer: nextConsumer,
		agent:        newAgent(ctx, in),
		translator:   trans,
		in:           in,
		cfg:          cfg,
		exit:         make(chan struct{}),
	}, nil
}

// Start implements the component.Component interface.
func (p *processor) Start(ctx context.Context, host component.Host) error {
	for k, exp := range host.GetExporters()[component.DataTypeMetrics] {
		mexp, ok := exp.(component.MetricsExporter)
		if !ok {
			return fmt.Errorf("the exporter %q isn't a metrics exporter", k.String())
		}
		if k.String() == p.cfg.MetricsExporter {
			p.metricsExporter = mexp
			break
		}
	}
	if p.metricsExporter == nil {
		return fmt.Errorf("failed to find metrics exporter %q; please specify a valid datadog::metrics_exporter", p.cfg.MetricsExporter)
	}
	p.agent.Start()
	go p.run()
	p.logger.Debug("Started datadogprocessor", zap.String("metrics_exporter", p.cfg.MetricsExporter))
	return nil
}

// Shutdown implements the component.Component interface.
func (p *processor) Shutdown(context.Context) error {
	p.logger.Info("Shutting down datadogprocessor")
	p.agent.Stop()
	p.exit <- struct{}{} // signal exit
	<-p.exit             // wait for close
	return nil
}

// Capabilities implements the consumer interface.
func (p *processor) Capabilities() consumer.Capabilities {
	// A resource attribute is added to traces to specify that stats have already
	// been computed for them; thus, we end up mutating the data:
	return consumer.Capabilities{MutatesData: true}
}

// ConsumeTraces implements consumer.Traces.
func (p *processor) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	p.logger.Debug("Received traces.", zap.Int("spans", traces.SpanCount()))
	p.agent.Ingest(ctx, traces)
	return p.nextConsumer.ConsumeTraces(ctx, traces)
}

// run awaits incoming stats resulting from the agent's ingestion, converts them
// to metrics and flushes them using the configured metrics exporter.
func (p *processor) run() {
	defer close(p.exit)
loop:
	for {
		select {
		case stats := <-p.in:
			if len(stats.Stats) == 0 {
				continue
			}
			mx := p.translator.StatsPayloadToMetrics(stats)
			ctx := context.TODO()
			p.logger.Debug("Exporting APM Stats metrics.", zap.Int("count", mx.MetricCount()))
			if err := p.metricsExporter.ConsumeMetrics(ctx, mx); err != nil {
				p.logger.Error("Error exporting metrics.", zap.Error(err))
			}
		case <-p.exit:
			break loop
		}
	}
}
