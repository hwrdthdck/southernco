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

package spanmetricsprocessor

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/translator/conventions"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor/metricsdimensions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor/mocks"
)

const (
	stringAttrName = "stringAttrName"
	intAttrName    = "intAttrName"
	doubleAttrName = "doubleAttrName"
	boolAttrName   = "boolAttrName"
	nullAttrName   = "nullAttrName"
	mapAttrName    = "mapAttrName"
	arrayAttrName  = "arrayAttrName"

	sampleLatency         = 11
	sampleLatencyDuration = sampleLatency * time.Millisecond
)

// metricID represents the minimum attributes that uniquely identifies a metric in our tests.
type metricID struct {
	service   string
	operation string
	kind      string
}

type metricDataPoint interface {
	LabelsMap() pdata.StringMap
}

type serviceSpans struct {
	serviceName string
	spans       []span
}

type span struct {
	operation  string
	kind       pdata.SpanKind
	statusCode pdata.StatusCode
}

func TestProcessorStart(t *testing.T) {
	// Create otlp exporters.
	otlpConfig, mexp, texp := newOTLPExporters(t)

	for _, tc := range []struct {
		name            string
		exporter        component.Exporter
		metricsExporter string
		wantErrorMsg    string
	}{
		{"export to active otlp metrics exporter", mexp, "otlp", ""},
		{"unable to find configured exporter in active exporter list", mexp, "prometheus", "failed to find metrics exporter: 'prometheus'; please configure metrics_exporter from one of: [otlp]"},
		{"export to active otlp traces exporter should error", texp, "otlp", "the exporter \"otlp\" isn't a metrics exporter"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare
			exporters := map[configmodels.DataType]map[configmodels.Exporter]component.Exporter{
				configmodels.MetricsDataType: {
					otlpConfig: tc.exporter,
				},
			}
			mhost := &mocks.Host{}
			mhost.On("GetExporters").Return(exporters)

			// Create spanmetrics processor
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig().(*Config)
			cfg.MetricsExporter = tc.metricsExporter

			procCreationParams := component.ProcessorCreateParams{Logger: zap.NewNop()}
			traceProcessor, err := factory.CreateTracesProcessor(context.Background(), procCreationParams, cfg, consumertest.NewTracesNop())
			require.NoError(t, err)

			// Test
			smp := traceProcessor.(*processorImp)
			err = smp.Start(context.Background(), mhost)

			// Verify
			if tc.wantErrorMsg != "" {
				assert.EqualError(t, err, tc.wantErrorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProcessorShutdown(t *testing.T) {
	// Prepare
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)

	// Test
	next := new(consumertest.TracesSink)
	p := newProcessor(zap.NewNop(), cfg, next)
	err := p.Shutdown(context.Background())

	// Verify
	assert.NoError(t, err)
}

func TestProcessorCapabilities(t *testing.T) {
	// Prepare
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)

	// Test
	next := new(consumertest.TracesSink)
	p := newProcessor(zap.NewNop(), cfg, next)
	caps := p.GetCapabilities()

	// Verify
	assert.NotNil(t, p)
	assert.Equal(t, false, caps.MutatesConsumedData)
}

func TestProcessorConsumeTracesErrors(t *testing.T) {
	for _, tc := range []struct {
		name                              string
		numMarshalCallsBeforeRaisingErr   int
		marshalErr                        error
		numUnmarshalCallsBeforeRaisingErr int
		unmarshalErr                      error
		consumeMetricsErr                 error
		consumeTracesErr                  error
		wantErrMsg                        string
	}{
		{
			name:                            "aggregateMetrics error",
			numMarshalCallsBeforeRaisingErr: 0,
			marshalErr:                      fmt.Errorf("marshal error"),
			wantErrMsg:                      "marshal error",
		},
		{
			name:                              "collectCallMetrics error",
			numUnmarshalCallsBeforeRaisingErr: 0,
			unmarshalErr:                      fmt.Errorf("unmarshal error"),
			wantErrMsg:                        "unmarshal error",
		},
		{
			name:                              "collectCallMetrics error",
			numUnmarshalCallsBeforeRaisingErr: 1,
			unmarshalErr:                      fmt.Errorf("unmarshal error"),
			wantErrMsg:                        "unmarshal error",
		},
		{
			name:              "metricsExporter error",
			consumeMetricsErr: fmt.Errorf("metricsExporter error"),
			wantErrMsg:        "metricsExporter error",
		},
		{
			name:             "nextConsumer error",
			consumeTracesErr: fmt.Errorf("nextConsumer error"),
			wantErrMsg:       "nextConsumer error",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare
			mexp := &mocks.MetricsExporter{}
			mexp.On("ConsumeMetrics", mock.Anything, mock.Anything).Return(tc.consumeMetricsErr)

			tcon := &mocks.TracesConsumer{}
			tcon.On("ConsumeTraces", mock.Anything, mock.Anything).Return(tc.consumeTracesErr)

			// Mock json.Marshal
			jsonSerder := &mocks.JsonSerder{}
			if tc.marshalErr != nil {
				for callNum := 0; callNum < tc.numMarshalCallsBeforeRaisingErr; callNum++ {
					jsonSerder.On("Marshal", mock.Anything).Return([]byte(""), nil).Once()
				}
			}
			jsonSerder.On("Marshal", mock.Anything).Return([]byte(""), tc.marshalErr)

			// Mock json.Unmarshal
			if tc.unmarshalErr != nil {
				for callNum := 0; callNum < tc.numUnmarshalCallsBeforeRaisingErr; callNum++ {
					jsonSerder.On("Unmarshal", mock.Anything, mock.Anything).Return(nil).Once()
				}
			}
			jsonSerder.On("Unmarshal", mock.Anything, mock.Anything).Return(tc.unmarshalErr)

			p := newProcessorImp(mexp, tcon, jsonSerder, nil)

			traces := buildSampleTrace()

			// Test
			ctx := metadata.NewIncomingContext(context.Background(), nil)
			err := p.ConsumeTraces(ctx, traces)

			// Verify
			assert.EqualError(t, err, tc.wantErrMsg)
		})
	}
}

func TestProcessorConsumeTraces(t *testing.T) {
	// Prepare
	mexp := &mocks.MetricsExporter{}
	tcon := &mocks.TracesConsumer{}

	mexp.On("ConsumeMetrics", mock.Anything, mock.MatchedBy(func(input pdata.Metrics) bool {
		return verifyConsumeMetricsInput(input, t)
	})).Return(nil)
	tcon.On("ConsumeTraces", mock.Anything, mock.Anything).Return(nil)

	defaultNullValue := "defaultNullValue"
	p := newProcessorImp(mexp, tcon, &JSONSerde{}, &defaultNullValue)

	traces := buildSampleTrace()

	// Test
	ctx := metadata.NewIncomingContext(context.Background(), nil)
	err := p.ConsumeTraces(ctx, traces)

	// Verify
	assert.NoError(t, err)
}

func BenchmarkProcessorConsumeTraces(b *testing.B) {
	// Prepare
	mexp := &mocks.MetricsExporter{}
	tcon := &mocks.TracesConsumer{}

	mexp.On("ConsumeMetrics", mock.Anything, mock.Anything).Return(nil)
	tcon.On("ConsumeTraces", mock.Anything, mock.Anything).Return(nil)

	defaultNullValue := "defaultNullValue"
	p := newProcessorImp(mexp, tcon, &JSONSerde{}, &defaultNullValue)

	traces := buildSampleTrace()

	// Test
	ctx := metadata.NewIncomingContext(context.Background(), nil)
	for n := 0; n < b.N; n++ {
		p.ConsumeTraces(ctx, traces)
	}
}

func TestDimensionKeyCache(t *testing.T) {
	// Prepare
	mexp := &mocks.MetricsExporter{}
	tcon := &mocks.TracesConsumer{}

	mexp.On("ConsumeMetrics", mock.Anything, mock.Anything).Return(nil)
	tcon.On("ConsumeTraces", mock.Anything, mock.Anything).Return(nil)

	defaultNullValue := "defaultNullValue"
	p := newProcessorImp(mexp, tcon, &JSONSerde{}, &defaultNullValue)

	traces := buildSampleTrace()

	// Test
	ctx := metadata.NewIncomingContext(context.Background(), nil)

	// Send 2 traces with the same span attributes should reuse a cached dimension key.
	assert.True(t, p.metricsKeyCache.Empty())
	p.ConsumeTraces(ctx, traces)
	p.ConsumeTraces(ctx, traces)
	assert.False(t, p.metricsKeyCache.Empty())

	d := p.metricsKeyCache.InsertDimensions([]metricsdimensions.DimensionKeyValue{
		{Key: serviceNameKey, Value: "service-a"},
		{Key: operationKey, Value: "/ping"},
		{Key: spanKindKey, Value: "SPAN_KIND_SERVER"},
		{Key: statusCodeKey, Value: "STATUS_CODE_OK"},
		{Key: stringAttrName, Value: "stringAttrValue"},
		{Key: intAttrName, Value: "99"},
		{Key: doubleAttrName, Value: "99.99"},
		{Key: boolAttrName, Value: "true"},
		{Key: nullAttrName},
	}...)
	assert.True(t, d.HasCachedMetricKey())
}

func newProcessorImp(mexp *mocks.MetricsExporter, tcon *mocks.TracesConsumer, jsonSerder JSONSerder, defaultNullValue *string) *processorImp {
	return &processorImp{
		logger:          zap.NewNop(),
		metricsExporter: mexp,
		nextConsumer:    tcon,

		callSum:             make(map[string]int64),
		latencySum:          make(map[string]float64),
		latencyCount:        make(map[string]uint64),
		latencyBucketCounts: make(map[string][]uint64),
		latencyBounds:       defaultLatencyHistogramBucketsMs,
		dimensions: []Dimension{
			// Set nil defaults to force a lookup for the attribute in the span.
			{stringAttrName, nil},
			{intAttrName, nil},
			{doubleAttrName, nil},
			{boolAttrName, nil},
			{mapAttrName, nil},
			{arrayAttrName, nil},
			{nullAttrName, defaultNullValue},
		},
		jsonSerder:      jsonSerder,
		metricsKeyCache: metricsdimensions.NewCache(),
	}
}

// verifyConsumeMetricsInput verifies the input of the ConsumeMetrics call from this processor.
// This is the best point to verify the computed metrics from spans are as expected.
func verifyConsumeMetricsInput(input pdata.Metrics, t *testing.T) bool {
	require.Equal(t, 6, input.MetricCount(),
		"Should be 3 for each of call count and latency. Each group of 3 metrics is made of: "+
			"service-a (server kind) -> service-a (client kind) -> service-b (service kind)",
	)

	rm := input.ResourceMetrics()
	require.Equal(t, 1, rm.Len())

	ilm := rm.At(0).InstrumentationLibraryMetrics()
	require.Equal(t, 1, ilm.Len())
	assert.Equal(t, "spanmetricsprocessor", ilm.At(0).InstrumentationLibrary().Name())

	m := ilm.At(0).Metrics()
	require.Equal(t, 6, m.Len())

	seenMetricIDs := make(map[metricID]bool)
	mi := 0
	// The first 3 metrics are for call counts.
	for ; mi < 3; mi++ {
		assert.Equal(t, "calls", m.At(mi).Name())

		data := m.At(mi).IntSum()
		assert.Equal(t, pdata.AggregationTemporalityCumulative, data.AggregationTemporality())

		dps := data.DataPoints()
		require.Equal(t, 1, dps.Len())

		dp := dps.At(0)
		assert.Equal(t, int64(1), dp.Value(), "There should only be one metric per Service/operation/kind combination")
		assert.NotZero(t, dp.Timestamp(), "Timestamp should be set")

		verifyMetricLabels(dp, t, seenMetricIDs)
	}

	seenMetricIDs = make(map[metricID]bool)
	// The remaining metrics are for latency.
	for ; mi < m.Len(); mi++ {
		assert.Equal(t, "latency", m.At(mi).Name())

		data := m.At(mi).IntHistogram()
		assert.Equal(t, pdata.AggregationTemporalityCumulative, data.AggregationTemporality())

		dps := data.DataPoints()
		require.Equal(t, 1, dps.Len())

		dp := dps.At(0)
		assert.Equal(t, int64(sampleLatency), dp.Sum(), "Should be a single 11ms latency measurement")
		assert.NotZero(t, dp.Timestamp(), "Timestamp should be set")

		// Verify bucket counts. Firstly, find the bucket index where the 11ms latency should belong in.
		var foundLatencyIndex int
		for foundLatencyIndex = 0; foundLatencyIndex < len(dp.ExplicitBounds()); foundLatencyIndex++ {
			if dp.ExplicitBounds()[foundLatencyIndex] > sampleLatency {
				break
			}
		}

		// Then verify that all histogram buckets are empty except for the bucket with the 11ms latency.
		var wantBucketCount uint64
		for bi := 0; bi < len(dp.BucketCounts()); bi++ {
			wantBucketCount = 0
			if bi == foundLatencyIndex {
				wantBucketCount = 1
			}
			assert.Equal(t, wantBucketCount, dp.BucketCounts()[bi])
		}
		verifyMetricLabels(dp, t, seenMetricIDs)
	}
	return true
}

func verifyMetricLabels(dp metricDataPoint, t *testing.T, seenMetricIDs map[metricID]bool) {
	mID := metricID{}
	dp.LabelsMap().ForEach(func(k string, v string) {
		switch k {
		case serviceNameKey:
			mID.service = v
		case operationKey:
			mID.operation = v
		case spanKindKey:
			mID.kind = v
		case stringAttrName:
			assert.Equal(t, "stringAttrValue", v)
		case intAttrName:
			assert.Equal(t, "99", v)
		case doubleAttrName:
			assert.Equal(t, "99.99", v)
		case boolAttrName:
			assert.Equal(t, "true", v)
		case nullAttrName:
			assert.Empty(t, v)
		}
	})
	// Service/operation/kind should be a unique metric.
	assert.False(t, seenMetricIDs[mID])
	seenMetricIDs[mID] = true
}

// buildSampleTrace builds the following trace:
//   service-a/ping (server) ->
//     service-a/ping (client) ->
//       service-b/ping (server)
func buildSampleTrace() pdata.Traces {
	traces := pdata.NewTraces()

	serviceASpans := buildServiceSpans(
		serviceSpans{
			serviceName: "service-a",
			spans: []span{
				{
					operation:  "/ping",
					kind:       pdata.SpanKindSERVER,
					statusCode: pdata.StatusCodeOk,
				},
				{
					operation:  "/ping",
					kind:       pdata.SpanKindCLIENT,
					statusCode: pdata.StatusCodeOk,
				},
			},
		})
	serviceBSpans := buildServiceSpans(
		serviceSpans{
			serviceName: "service-b",
			spans: []span{
				{
					operation:  "/ping",
					kind:       pdata.SpanKindSERVER,
					statusCode: pdata.StatusCodeError,
				},
			},
		})
	ignoredSpans := buildServiceSpans(serviceSpans{})
	traces.ResourceSpans().Append(serviceASpans)
	traces.ResourceSpans().Append(serviceBSpans)
	traces.ResourceSpans().Append(ignoredSpans)
	return traces
}

func buildServiceSpans(serviceSpans serviceSpans) pdata.ResourceSpans {
	spans := pdata.NewResourceSpans()

	if serviceSpans.serviceName != "" {
		spans.Resource().Attributes().
			InsertString(conventions.AttributeServiceName, serviceSpans.serviceName)
	}

	ils := pdata.NewInstrumentationLibrarySpans()
	for _, span := range serviceSpans.spans {
		ils.Spans().Append(buildSpan(span))
	}
	spans.InstrumentationLibrarySpans().Append(ils)
	return spans
}

func buildSpan(span span) pdata.Span {
	s := pdata.NewSpan()
	s.SetName(span.operation)
	s.SetKind(span.kind)
	s.Status().SetCode(span.statusCode)
	now := time.Now()
	s.SetStartTime(pdata.TimestampUnixNano(now.UnixNano()))
	s.SetEndTime(pdata.TimestampUnixNano(now.Add(sampleLatencyDuration).UnixNano()))
	s.Attributes().InsertString(stringAttrName, "stringAttrValue")
	s.Attributes().InsertInt(intAttrName, 99)
	s.Attributes().InsertDouble(doubleAttrName, 99.99)
	s.Attributes().InsertBool(boolAttrName, true)
	s.Attributes().InsertNull(nullAttrName)
	s.Attributes().Insert(mapAttrName, pdata.NewAttributeValueMap())
	s.Attributes().Insert(arrayAttrName, pdata.NewAttributeValueArray())
	return s
}

func newOTLPExporters(t *testing.T) (*otlpexporter.Config, component.MetricsExporter, component.TracesExporter) {
	otlpExpFactory := otlpexporter.NewFactory()
	otlpConfig := &otlpexporter.Config{
		ExporterSettings: configmodels.ExporterSettings{
			NameVal: "otlp",
			TypeVal: "otlp",
		},
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Endpoint: "example.com:1234",
		},
	}
	expCreationParams := component.ExporterCreateParams{Logger: zap.NewNop()}
	mexp, err := otlpExpFactory.CreateMetricsExporter(context.Background(), expCreationParams, otlpConfig)
	require.NoError(t, err)
	texp, err := otlpExpFactory.CreateTracesExporter(context.Background(), expCreationParams, otlpConfig)
	require.NoError(t, err)
	return otlpConfig, mexp, texp
}
