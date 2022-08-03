// Copyright OpenTelemetry Authors
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

package azureblobexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azureblobexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type traceExporter struct {
	blobClient      BlobClient
	logger          *zap.Logger
	tracesMarshaler ptrace.Marshaler
}

func (ex *traceExporter) onTraceData(context context.Context, traceData ptrace.Traces) error {
	buf, err := ex.tracesMarshaler.MarshalTraces(traceData)
	if err != nil {
		ex.logger.Error(err.Error())
		return err
	}

	err = ex.blobClient.UploadData(context, buf, config.TracesDataType)
	if err != nil {
		ex.logger.Error(err.Error())
	}

	return err
}

// Returns a new instance of the trace exporter
func newTracesExporter(config *Config, blobClient BlobClient, set component.ExporterCreateSettings) (component.TracesExporter, error) {
	exporter := &traceExporter{
		blobClient:      blobClient,
		logger:          set.Logger,
		tracesMarshaler: ptrace.NewJSONMarshaler(),
	}

	return exporterhelper.NewTracesExporter(config, set, exporter.onTraceData)
}
