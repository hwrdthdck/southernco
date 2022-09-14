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

package routingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/telemetryquerylanguage/contexts/tqllogs"
)

var _ component.LogsProcessor = (*logProcessor)(nil)

type logProcessor struct {
	logger *zap.Logger
	config *Config

	extractor extractor
	router    router[component.LogsExporter]
}

func newLogProcessor(logger *zap.Logger, config config.Processor) *logProcessor {
	cfg := rewriteRoutingEntriesToTQL(config.(*Config))

	return &logProcessor{
		logger: logger,
		config: cfg,
		router: newRouter[component.LogsExporter](
			cfg.Table,
			cfg.DefaultExporters,
			logger,
		),
		extractor: newExtractor(cfg.FromAttribute, logger),
	}
}

func (p *logProcessor) Start(_ context.Context, host component.Host) error {
	err := p.router.registerExporters(host.GetExporters()[config.LogsDataType])
	if err != nil {
		return err
	}
	return nil
}

func (p *logProcessor) ConsumeLogs(ctx context.Context, l plog.Logs) error {
	if p.config.FromAttribute == "" {
		err := p.route(ctx, l)
		if err != nil {
			return err
		}
		return nil
	}
	err := p.routeForContext(ctx, l)
	if err != nil {
		return err
	}
	return nil
}

func (p *logProcessor) route(ctx context.Context, l plog.Logs) error {
	// routingEntry is used to group plog.ResourceLogs that are routed to
	// the same set of exporters.
	// This way we're not ending up with all the logs split up which would cause
	// higher CPU usage.
	groups := map[string]struct {
		exporters []component.LogsExporter
		resLogs   plog.ResourceLogsSlice
	}{}
	var errs error

	for i := 0; i < l.ResourceLogs().Len(); i++ {
		rlogs := l.ResourceLogs().At(i)
		ltx := tqllogs.NewTransformContext(plog.LogRecord{}, pcommon.InstrumentationScope{}, rlogs.Resource())

		for key, route := range p.router.routes {
			exp := p.router.defaultExporters
			var gKey string
			if route.expression.Condition(ltx) {
				route.expression.Function(ltx)
				exp = route.exporters
				gKey = key
			}

			if g, ok := groups[gKey]; ok {
				rlogs.MoveTo(g.resLogs.AppendEmpty())
			} else {
				newResLogs := plog.NewResourceLogsSlice()
				rlogs.MoveTo(newResLogs.AppendEmpty())

				groups[gKey] = struct {
					exporters []component.LogsExporter
					resLogs   plog.ResourceLogsSlice
				}{
					exporters: exp,
					resLogs:   newResLogs,
				}
			}
		}
	}

	for _, g := range groups {
		l := plog.NewLogs()
		l.ResourceLogs().EnsureCapacity(g.resLogs.Len())
		g.resLogs.MoveAndAppendTo(l.ResourceLogs())

		for _, e := range g.exporters {
			errs = multierr.Append(errs, e.ConsumeLogs(ctx, l))
		}
	}
	return errs
}

func (p *logProcessor) routeForContext(ctx context.Context, l plog.Logs) error {
	value := p.extractor.extractFromContext(ctx)
	exporters := p.router.getExporters(value)

	var errs error
	for _, e := range exporters {
		errs = multierr.Append(errs, e.ConsumeLogs(ctx, l))
	}
	return errs
}

func (p *logProcessor) Shutdown(context.Context) error {
	return nil
}

func (p *logProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
