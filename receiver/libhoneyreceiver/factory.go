// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package libhoneyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/sharedcomponent"
)

const (
	httpPort = 8080
)

var defaultTracesURLPaths = []string{"/events", "/event", "/batch"}

// NewFactory creates a new OTLP receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithTraces(createTraces, metadata.TracesStability),
		receiver.WithLogs(createLogs, metadata.LogsStability),
	)
}

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	durationFieldsArr := []string{"duration_ms"}
	endpointStr := fmt.Sprintf("localhost:%d", httpPort)
	return &Config{
		HTTP: &HTTPConfig{
			ServerConfig: &confighttp.ServerConfig{
				Endpoint: endpointStr,
			},
			TracesURLPaths: defaultTracesURLPaths,
		},
		AuthApi: "",
		Resources: ResourcesConfig{
			ServiceName: "service.name",
		},
		Scopes: ScopesConfig{
			LibraryName:    "library.name",
			LibraryVersion: "library.version",
		},
		Attributes: AttributesConfig{
			TraceId:        "trace.trace_id",
			SpanId:         "trace.span_id",
			ParentId:       "trace.parent_id",
			Name:           "name",
			Error:          "error",
			SpanKind:       "span.kind",
			DurationFields: durationFieldsArr,
		},
	}
}

func createLogs(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	oCfg := cfg.(*Config)
	r, err := receivers.LoadOrStore(
		oCfg,
		func() (*libhoneyReceiver, error) {
			return newLibhoneyReceiver(oCfg, &set)
		},
	)
	if err != nil {
		return nil, err
	}

	r.Unwrap().registerLogConsumer(nextConsumer)
	return r, nil
}

// createTraces creates a trace receiver based on provided config.
func createTraces(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	oCfg := cfg.(*Config)
	r, err := receivers.LoadOrStore(
		oCfg,
		func() (*libhoneyReceiver, error) {
			return newLibhoneyReceiver(oCfg, &set)
		},
	)
	if err != nil {
		return nil, err
	}

	r.Unwrap().registerTraceConsumer(nextConsumer)
	return r, nil
}

// Used the same pattern as the OTLP receiver. Requires sharedcomponent.go from core collector repo
var receivers = sharedcomponent.NewMap[*Config, *libhoneyReceiver]()
