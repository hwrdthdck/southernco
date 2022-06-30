// Copyright 2021 OpenTelemetry Authors
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

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type S3Exporter struct {
	config     config.Exporter
	dataWriter DataWriter
	logger     *zap.Logger
	marshaler  Marshaler
}

func NewS3Exporter(config config.Exporter,
	params component.ExporterCreateSettings) (*S3Exporter, error) {

	if config == nil {
		return nil, errors.New("s3 exporter config is nil")
	}

	logger := params.Logger
	expConfig := config.(*Config)
	expConfig.logger = logger

	validateConfig := expConfig.Validate()

	if validateConfig != nil {
		return nil, validateConfig
	}

	marshaler, err := NewMarshaler(expConfig.MarshalerName, logger)
	if err != nil {
		return nil, errors.New("unknown marshaler")
	}

	s3Exporter := &S3Exporter{
		config:     config,
		dataWriter: &S3Writer{},
		logger:     logger,
		marshaler:  marshaler,
	}
	return s3Exporter, nil
}

func (e *S3Exporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (e *S3Exporter) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (e *S3Exporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	buf, err := e.marshaler.MarshalLogs(logs)

	if err != nil {
		return err
	}
	expConfig := e.config.(*Config)

	return e.dataWriter.WriteBuffer(ctx, buf, expConfig, "logs", e.marshaler.Format())
}

func (e *S3Exporter) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	buf, err := e.marshaler.MarshalTraces(traces)
	if err != nil {
		return err
	}
	expConfig := e.config.(*Config)

	return e.dataWriter.WriteBuffer(ctx, buf, expConfig, "traces", e.marshaler.Format())
}

func (e *S3Exporter) Shutdown(context.Context) error {
	return nil
}
