// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsxrayexporter

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
	"github.com/open-telemetry/opentelemetry-collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

// NewTraceExporter creates an exporter.TraceExporter that just drops the
// received data and logs debugging messages.
func NewTraceExporter(config configmodels.Exporter, logger *zap.Logger) (exporter.TraceExporter, error) {
	typeLog := zap.String("type", config.Type())
	nameLog := zap.String("name", config.Name())
	return exporterhelper.NewTraceExporter(
		config,
		func(ctx context.Context, td consumerdata.TraceData) (int, error) {
			logger.Info("TraceExporter", typeLog, nameLog, zap.Int("#spans", len(td.Spans)))
			// TODO: Add ability to record the received data
			return 0, nil
		},
		exporterhelper.WithTracing(true),
		exporterhelper.WithMetrics(false),
		exporterhelper.WithShutdown(logger.Sync),
	)
}
