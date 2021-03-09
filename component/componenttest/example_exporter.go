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

package componenttest

import (
	"context"

	"github.com/spf13/viper"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/pdata"
)

// ExampleExporter is for testing purposes. We are defining an example config and factory
// for "exampleexporter" exporter type.
type ExampleExporter struct {
	configmodels.ExporterSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
	ExtraInt                      int32                    `mapstructure:"extra_int"`
	ExtraSetting                  string                   `mapstructure:"extra"`
	ExtraMapSetting               map[string]string        `mapstructure:"extra_map"`
	ExtraListSetting              []string                 `mapstructure:"extra_list"`
}

// ExampleExporterFactory is factory for ExampleExporter.
type ExampleExporterFactory struct {
}

// Type gets the type of the Exporter config created by this factory.
func (f *ExampleExporterFactory) Type() configmodels.Type {
	return "exampleexporter"
}

// CreateDefaultConfig creates the default configuration for the Exporter.
func (f *ExampleExporterFactory) CreateDefaultConfig() configmodels.Exporter {
	return &ExampleExporter{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: f.Type(),
			NameVal: string(f.Type()),
		},
		ExtraSetting:     "some export string",
		ExtraMapSetting:  nil,
		ExtraListSetting: nil,
	}
}

// Unmarshal implements the custom unmarshalers.
func (f *ExampleExporterFactory) Unmarshal(componentViperSection *viper.Viper, intoCfg interface{}) error {
	return componentViperSection.UnmarshalExact(intoCfg)
}

// CreateTracesExporter creates a trace exporter based on this config.
func (f *ExampleExporterFactory) CreateTracesExporter(
	_ context.Context,
	_ component.ExporterCreateParams,
	_ configmodels.Exporter,
) (component.TracesExporter, error) {
	return &ExampleExporterConsumer{}, nil
}

// CreateMetricsExporter creates a metrics exporter based on this config.
func (f *ExampleExporterFactory) CreateMetricsExporter(
	_ context.Context,
	_ component.ExporterCreateParams,
	_ configmodels.Exporter,
) (component.MetricsExporter, error) {
	return &ExampleExporterConsumer{}, nil
}

func (f *ExampleExporterFactory) CreateLogsExporter(
	_ context.Context,
	_ component.ExporterCreateParams,
	_ configmodels.Exporter,
) (component.LogsExporter, error) {
	return &ExampleExporterConsumer{}, nil
}

// ExampleExporterConsumer stores consumed traces and metrics for testing purposes.
type ExampleExporterConsumer struct {
	Traces           []pdata.Traces
	Metrics          []pdata.Metrics
	Logs             []pdata.Logs
	ExporterStarted  bool
	ExporterShutdown bool
}

// Start tells the exporter to start. The exporter may prepare for exporting
// by connecting to the endpoint. Host parameter can be used for communicating
// with the host after Start() has already returned.
func (exp *ExampleExporterConsumer) Start(_ context.Context, _ component.Host) error {
	exp.ExporterStarted = true
	return nil
}

// ConsumeTraces receives pdata.Traces for processing by the TracesConsumer.
func (exp *ExampleExporterConsumer) ConsumeTraces(_ context.Context, td pdata.Traces) error {
	exp.Traces = append(exp.Traces, td)
	return nil
}

// ConsumeMetrics receives pdata.Metrics for processing by the MetricsConsumer.
func (exp *ExampleExporterConsumer) ConsumeMetrics(_ context.Context, md pdata.Metrics) error {
	exp.Metrics = append(exp.Metrics, md)
	return nil
}

func (exp *ExampleExporterConsumer) ConsumeLogs(_ context.Context, ld pdata.Logs) error {
	exp.Logs = append(exp.Logs, ld)
	return nil
}

// Shutdown is invoked during shutdown.
func (exp *ExampleExporterConsumer) Shutdown(context.Context) error {
	exp.ExporterShutdown = true
	return nil
}
