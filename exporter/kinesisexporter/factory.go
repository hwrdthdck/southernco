// Copyright 2019 OpenTelemetry Authors
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

package kinesisexporter

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr   = "kinesis"
	otlpProto = "otlp_proto"
	// The default encoding scheme is set to otlpProto
	defaultEncoding = otlpProto
)

type creationFunc func(context.Context, component.ExporterCreateParams, configmodels.Exporter) (component.Exporter, error)

// NewFactory creates a factory for Kinesis exporter.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithTraces(createTraceExporter),
		exporterhelper.WithMetrics(createMetricsExporter))
}

func createDefaultConfig() configmodels.Exporter {
	return &Config{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		AWS: AWSConfig{
			Region:     "us-west-2",
			StreamName: "test-stream",
		},
		KPL: KPLConfig{
			BatchSize:            5242880,
			BatchCount:           1000,
			BacklogCount:         2000,
			FlushIntervalSeconds: 5,
			MaxConnections:       24,
		},
		Encoding: defaultEncoding,
	}
}

func createTraceExporter(
	ctx context.Context,
	params component.ExporterCreateParams,
	config configmodels.Exporter,
) (component.TracesExporter, error) {
	if err := validateParams(ctx, params, config); err != nil {
		return nil, err
	}

	c := config.(*Config)
	exp, err := newExporter(c, params.Logger)
	if err != nil {
		return nil, err
	}

	return exporterhelper.NewTraceExporter(
		c,
		params.Logger,
		exp.pushTraces,
		exporterhelper.WithStart(exp.start),
		exporterhelper.WithShutdown(exp.shutdown))
}

func createMetricsExporter(
	ctx context.Context,
	params component.ExporterCreateParams,
	config configmodels.Exporter,
) (component.MetricsExporter, error) {
	if err := validateParams(ctx, params, config); err != nil {
		return nil, err
	}
	c := config.(*Config)
	exp, err := newExporter(c, params.Logger)
	if err != nil {
		return nil, err
	}

	return exporterhelper.NewMetricsExporter(
		c,
		params.Logger,
		exp.pushMetrics,
		exporterhelper.WithStart(exp.start),
		exporterhelper.WithShutdown(exp.shutdown))
}

func validateParams(ctx context.Context, params component.ExporterCreateParams, config configmodels.Exporter) error {
	if config == nil {
		return errors.New("nil config")
	}

	if params.Logger == nil {
		return errors.New("nil logger")
	}

	if ctx == nil {
		return errors.New("nil context")
	}

	return nil
}
