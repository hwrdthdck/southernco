// Copyright 2018, OpenCensus Authors
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

package exporterparser

import (
	"context"
	"fmt"

	"contrib.go.opencensus.io/exporter/ocagent"
	"github.com/spf13/viper"

	agenttracepb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/trace/v1"
	"github.com/census-instrumentation/opencensus-service/data"
	"github.com/census-instrumentation/opencensus-service/exporter"
	"github.com/census-instrumentation/opencensus-service/internal"
	"github.com/census-instrumentation/opencensus-service/internal/compression"
	"github.com/census-instrumentation/opencensus-service/internal/compression/grpc"
)

type opencensusConfig struct {
	Endpoint    string            `mapstructure:"endpoint,omitempty"`
	Compression string            `mapstructure:"compression,omitempty"`
	Headers     map[string]string `mapstructure:"headers,omitempty"`
	// TODO: add insecure, service name options.
}

type ocagentExporter struct {
	exporter *ocagent.Exporter
}

var _ exporter.TraceExporter = (*ocagentExporter)(nil)

// OpenCensusTraceExportersFromViper unmarshals the viper and returns an exporter.TraceExporter targeting
// OpenCensus Agent/Collector according to the configuration settings.
func OpenCensusTraceExportersFromViper(v *viper.Viper) (tes []exporter.TraceExporter, mes []exporter.MetricsExporter, doneFns []func() error, err error) {
	var cfg struct {
		OpenCensus *opencensusConfig `mapstructure:"opencensus"`
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, nil, nil, err
	}
	ocac := cfg.OpenCensus
	if ocac == nil {
		return nil, nil, nil, nil
	}

	if ocac.Endpoint == "" {
		return nil, nil, nil, fmt.Errorf("OpenCensus config requires an Endpoint")
	}

	opts := []ocagent.ExporterOption{ocagent.WithAddress(ocac.Endpoint), ocagent.WithInsecure()}
	if ocac.Compression != "" {
		if compressionKey := grpc.GetGRPCCompressionKey(ocac.Compression); compressionKey != compression.Unsupported {
			opts = append(opts, ocagent.UseCompressor(compressionKey))
		} else {
			return nil, nil, nil, fmt.Errorf("Unsupported compression type: %s", ocac.Compression)
		}
	}
	if len(ocac.Headers) > 0 {
		opts = append(opts, ocagent.WithHeaders(ocac.Headers))
	}

	sde, serr := ocagent.NewExporter(opts...)
	if serr != nil {
		return nil, nil, nil, fmt.Errorf("Cannot configure OpenCensus Trace exporter: %v", serr)
	}

	oexp := &ocagentExporter{exporter: sde}
	tes = append(tes, oexp)

	// TODO: (@odeke-em, @songya23) implement ExportMetrics for OpenCensus.
	// mes = append(mes, oexp)
	doneFns = append(doneFns, func() error {
		sde.Flush()
		return nil
	})
	return tes, mes, doneFns, nil
}

func (sde *ocagentExporter) ExportSpans(ctx context.Context, td data.TraceData) error {
	err := sde.exporter.ExportTraceServiceRequest(
		&agenttracepb.ExportTraceServiceRequest{
			Spans: td.Spans,
			Node:  td.Node,
		},
	)
	if err != nil {
		return err
	}
	nSpansCounter := internal.NewExportedSpansRecorder("ocagent")
	nSpansCounter(ctx, td.Node, td.Spans)
	return nil
}
