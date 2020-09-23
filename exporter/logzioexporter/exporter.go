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

package logzioexporter

import (
	"errors"

	"github.com/logzio/jaeger-logzio/store"
	"go.opentelemetry.io/collector/component"
)

// logzioExporter implements an OpenTelemetry trace exporter that exports all spans to Logz.io
type logzioExporter struct {
	accountToken string
	writer       *store.LogzioSpanWriter
}

func newLogzioExporter(config *Config, params component.ExporterCreateParams) (*logzioExporter, error) {

	if config == nil {
		return nil, errors.New("exporter config can't be null")
	}
	return &logzioExporter{}, nil
}

func newLogzioTraceExporter(config *Config, params component.ExporterCreateParams) (component.TraceExporter, error) {
	return nil, nil
}

