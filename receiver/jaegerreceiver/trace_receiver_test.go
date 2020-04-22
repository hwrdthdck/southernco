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

package jaegerreceiver

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"path"
	"testing"
	"time"

	"contrib.go.opencensus.io/exporter/jaeger"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/proto-gen/api_v2"
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/open-telemetry/opentelemetry-collector/component"
	"github.com/open-telemetry/opentelemetry-collector/component/componenttest"
	"github.com/open-telemetry/opentelemetry-collector/consumer/pdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter/exportertest"
	"github.com/open-telemetry/opentelemetry-collector/receiver"
	"github.com/open-telemetry/opentelemetry-collector/testutils"
	"github.com/open-telemetry/opentelemetry-collector/translator/conventions"
	tracetranslator "github.com/open-telemetry/opentelemetry-collector/translator/trace"
)

const jaegerReceiver = "jaeger_receiver_test"

func TestTraceSource(t *testing.T) {
	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, &Configuration{}, nil, params)
	assert.NoError(t, err, "should not have failed to create the Jaeger receiver")
	require.NotNil(t, jr)
}

func TestReception(t *testing.T) {
	// 1. Create the Jaeger receiver aka "server"
	config := &Configuration{
		CollectorHTTPPort: 14268, // that's the only one used by this test
	}
	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	defer jr.Shutdown(context.Background())
	assert.NoError(t, err, "should not have failed to create the Jaeger received")

	t.Log("Starting")

	require.NoError(t, jr.Start(context.Background(), componenttest.NewNopHost()))

	t.Log("Start")

	now := time.Unix(1542158650, 536343000).UTC()
	nowPlus10min := now.Add(10 * time.Minute)
	nowPlus10min2sec := now.Add(10 * time.Minute).Add(2 * time.Second)

	// 2. Then with a "live application", send spans to the Jaeger exporter.
	jexp, err := jaeger.NewExporter(jaeger.Options{
		Process: jaeger.Process{
			ServiceName: "issaTest",
			Tags: []jaeger.Tag{
				jaeger.BoolTag("bool", true),
				jaeger.StringTag("string", "yes"),
				jaeger.Int64Tag("int64", 1e7),
			},
		},
		CollectorEndpoint: fmt.Sprintf("http://localhost:%d/api/traces", config.CollectorHTTPPort),
	})
	assert.NoError(t, err, "should not have failed to create the Jaeger OpenCensus exporter")

	// 3. Now finally send some spans
	for _, sd := range traceFixture(now, nowPlus10min, nowPlus10min2sec) {
		jexp.ExportSpan(sd)
	}
	jexp.Flush()

	gotTraces := sink.AllTraces()
	assert.Equal(t, 1, len(gotTraces))
	want := expectedTraceData(now, nowPlus10min, nowPlus10min2sec)

	assert.EqualValues(t, want, gotTraces[0])
}

func TestPortsNotOpen(t *testing.T) {
	// an empty config should result in no open ports
	config := &Configuration{}

	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	assert.NoError(t, err, "should not have failed to create a new receiver")
	defer jr.Shutdown(context.Background())

	require.NoError(t, jr.Start(context.Background(), componenttest.NewNopHost()))

	// there is a race condition here that we're ignoring.
	//  this test may occasionally pass incorrectly, but it will not fail incorrectly
	//  TODO: consider adding a way for a receiver to asynchronously signal that is ready to receive spans to eliminate races/arbitrary waits
	l, err := net.Listen("tcp", "localhost:14250")
	assert.NoError(t, err, "should have been able to listen on 14250.  jaeger receiver incorrectly started grpc")
	if l != nil {
		l.Close()
	}

	l, err = net.Listen("tcp", "localhost:14268")
	assert.NoError(t, err, "should have been able to listen on 14268.  jaeger receiver incorrectly started thrift_http")
	if l != nil {
		l.Close()
	}
}

func TestGRPCReception(t *testing.T) {
	// prepare
	config := &Configuration{
		CollectorGRPCPort: 14250, // that's the only one used by this test
	}
	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	assert.NoError(t, err, "should not have failed to create a new receiver")
	defer jr.Shutdown(context.Background())

	require.NoError(t, jr.Start(context.Background(), componenttest.NewNopHost()))
	t.Log("Start")

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", config.CollectorGRPCPort), grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	cl := api_v2.NewCollectorServiceClient(conn)

	now := time.Unix(1542158650, 536343000).UTC()
	d10min := 10 * time.Minute
	d2sec := 2 * time.Second
	nowPlus10min := now.Add(d10min)
	nowPlus10min2sec := now.Add(d10min).Add(d2sec)

	// test
	req := grpcFixture(now, d10min, d2sec)
	resp, err := cl.PostSpans(context.Background(), req, grpc.WaitForReady(true))

	// verify
	assert.NoError(t, err, "should not have failed to post spans")
	assert.NotNil(t, resp, "response should not have been nil")

	gotTraces := sink.AllTraces()
	assert.Equal(t, 1, len(gotTraces))
	want := expectedTraceData(now, nowPlus10min, nowPlus10min2sec)

	assert.Len(t, req.Batch.Spans, want.SpanCount(), "got a conflicting amount of spans")

	assert.EqualValues(t, want, gotTraces[0])
}

func TestGRPCReceptionWithTLS(t *testing.T) {
	// prepare
	grpcServerOptions := []grpc.ServerOption{}
	tlsCreds := receiver.TLSCredentials{
		CertFile: path.Join(".", "testdata", "certificate.pem"),
		KeyFile:  path.Join(".", "testdata", "key.pem"),
	}

	tlsOption, _ := tlsCreds.ToGrpcServerOption()

	grpcServerOptions = append(grpcServerOptions, tlsOption)

	port := testutils.GetAvailablePort(t)
	config := &Configuration{
		CollectorGRPCPort:    int(port),
		CollectorGRPCOptions: grpcServerOptions,
	}
	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	assert.NoError(t, err, "should not have failed to create a new receiver")
	defer jr.Shutdown(context.Background())

	require.NoError(t, jr.Start(context.Background(), componenttest.NewNopHost()))
	t.Log("Start")

	creds, err := credentials.NewClientTLSFromFile(path.Join(".", "testdata", "certificate.pem"), "opentelemetry.io")
	require.NoError(t, err)
	conn, err := grpc.Dial(jr.(*jReceiver).collectorGRPCAddr(), grpc.WithTransportCredentials(creds))
	require.NoError(t, err)
	defer conn.Close()

	cl := api_v2.NewCollectorServiceClient(conn)

	now := time.Now()
	d10min := 10 * time.Minute
	d2sec := 2 * time.Second
	nowPlus10min := now.Add(d10min)
	nowPlus10min2sec := now.Add(d10min).Add(d2sec)

	// test
	req := grpcFixture(now, d10min, d2sec)
	resp, err := cl.PostSpans(context.Background(), req, grpc.WaitForReady(true))

	// verify
	assert.NoError(t, err, "should not have failed to post spans")
	assert.NotNil(t, resp, "response should not have been nil")

	gotTraces := sink.AllTraces()
	assert.Equal(t, 1, len(gotTraces))
	want := expectedTraceData(now, nowPlus10min, nowPlus10min2sec)

	assert.Len(t, req.Batch.Spans, want.SpanCount(), "got a conflicting amount of spans")
	assert.EqualValues(t, want, gotTraces[0])
}

func expectedTraceData(t1, t2, t3 time.Time) pdata.Traces {
	traceID := pdata.TraceID(
		[]byte{0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF, 0x80})
	parentSpanID := pdata.SpanID([]byte{0x1F, 0x1E, 0x1D, 0x1C, 0x1B, 0x1A, 0x19, 0x18})
	childSpanID := pdata.SpanID([]byte{0xAF, 0xAE, 0xAD, 0xAC, 0xAB, 0xAA, 0xA9, 0xA8})

	traces := pdata.NewTraces()
	traces.ResourceSpans().Resize(1)
	rs := traces.ResourceSpans().At(0)
	rs.Resource().InitEmpty()
	rs.Resource().Attributes().InsertString(conventions.AttributeServiceName, "issaTest")
	rs.Resource().Attributes().InsertBool("bool", true)
	rs.Resource().Attributes().InsertString("string", "yes")
	rs.Resource().Attributes().InsertInt("int64", 10000000)
	rs.InstrumentationLibrarySpans().Resize(1)
	rs.InstrumentationLibrarySpans().At(0).Spans().Resize(2)

	span0 := rs.InstrumentationLibrarySpans().At(0).Spans().At(0)
	span0.SetSpanID(childSpanID)
	span0.SetParentSpanID(parentSpanID)
	span0.SetTraceID(traceID)
	span0.SetName("DBSearch")
	span0.SetStartTime(pdata.TimestampUnixNano(uint64(t1.UnixNano())))
	span0.SetEndTime(pdata.TimestampUnixNano(uint64(t2.UnixNano())))
	span0.Status().InitEmpty()
	span0.Status().SetCode(pdata.StatusCode(otlptrace.Status_NotFound))
	span0.Status().SetMessage("Stale indices")
	span0.Attributes().InsertBool("error", true)

	span1 := rs.InstrumentationLibrarySpans().At(0).Spans().At(1)
	span1.SetSpanID(parentSpanID)
	span1.SetTraceID(traceID)
	span1.SetName("ProxyFetch")
	span1.SetStartTime(pdata.TimestampUnixNano(uint64(t2.UnixNano())))
	span1.SetEndTime(pdata.TimestampUnixNano(uint64(t3.UnixNano())))
	span1.Status().InitEmpty()
	span1.Status().SetCode(pdata.StatusCode(otlptrace.Status_InternalError))
	span1.Status().SetMessage("Frontend crash")
	span1.Attributes().InsertBool("error", true)

	return traces
}

func traceFixture(t1, t2, t3 time.Time) []*trace.SpanData {
	traceID := trace.TraceID{0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF, 0x80}
	parentSpanID := trace.SpanID{0x1F, 0x1E, 0x1D, 0x1C, 0x1B, 0x1A, 0x19, 0x18}
	childSpanID := trace.SpanID{0xAF, 0xAE, 0xAD, 0xAC, 0xAB, 0xAA, 0xA9, 0xA8}

	return []*trace.SpanData{
		{
			SpanContext: trace.SpanContext{
				TraceID: traceID,
				SpanID:  childSpanID,
			},
			ParentSpanID: parentSpanID,
			Name:         "DBSearch",
			StartTime:    t1,
			EndTime:      t2,
			Status: trace.Status{
				Code:    trace.StatusCodeNotFound,
				Message: "Stale indices",
			},
			Links: []trace.Link{
				{
					TraceID: traceID,
					SpanID:  parentSpanID,
					Type:    trace.LinkTypeParent,
				},
			},
		},
		{
			SpanContext: trace.SpanContext{
				TraceID: traceID,
				SpanID:  parentSpanID,
			},
			Name:      "ProxyFetch",
			StartTime: t2,
			EndTime:   t3,
			Status: trace.Status{
				Code:    trace.StatusCodeInternal,
				Message: "Frontend crash",
			},
		},
	}
}

func grpcFixture(t1 time.Time, d1, d2 time.Duration) *api_v2.PostSpansRequest {
	traceID := model.TraceID{}
	traceID.Unmarshal([]byte{0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF, 0x80})
	parentSpanID := model.NewSpanID(binary.BigEndian.Uint64([]byte{0x1F, 0x1E, 0x1D, 0x1C, 0x1B, 0x1A, 0x19, 0x18}))
	childSpanID := model.NewSpanID(binary.BigEndian.Uint64([]byte{0xAF, 0xAE, 0xAD, 0xAC, 0xAB, 0xAA, 0xA9, 0xA8}))

	return &api_v2.PostSpansRequest{
		Batch: model.Batch{
			Process: &model.Process{
				ServiceName: "issaTest",
				Tags: []model.KeyValue{
					model.Bool("bool", true),
					model.String("string", "yes"),
					model.Int64("int64", 1e7),
				},
			},
			Spans: []*model.Span{
				{
					TraceID:       traceID,
					SpanID:        childSpanID,
					OperationName: "DBSearch",
					StartTime:     t1,
					Duration:      d1,
					Tags: []model.KeyValue{
						model.String(tracetranslator.TagStatusMsg, "Stale indices"),
						model.Int64(tracetranslator.TagStatusCode, trace.StatusCodeNotFound),
						model.Bool("error", true),
					},
					References: []model.SpanRef{
						{
							TraceID: traceID,
							SpanID:  parentSpanID,
							RefType: model.SpanRefType_CHILD_OF,
						},
					},
				},
				{
					TraceID:       traceID,
					SpanID:        parentSpanID,
					OperationName: "ProxyFetch",
					StartTime:     t1.Add(d1),
					Duration:      d2,
					Tags: []model.KeyValue{
						model.String(tracetranslator.TagStatusMsg, "Frontend crash"),
						model.Int64(tracetranslator.TagStatusCode, trace.StatusCodeInternal),
						model.Bool("error", true),
					},
				},
			},
		},
	}
}

func TestSampling(t *testing.T) {
	port := testutils.GetAvailablePort(t)
	config := &Configuration{
		CollectorGRPCPort:          int(port),
		RemoteSamplingStrategyFile: "testdata/strategies.json",
	}
	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	assert.NoError(t, err, "should not have failed to create a new receiver")
	defer jr.Shutdown(context.Background())

	require.NoError(t, jr.Start(context.Background(), componenttest.NewNopHost()))
	t.Log("Start")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.CollectorGRPCPort), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	cl := api_v2.NewSamplingManagerClient(conn)

	expected := &api_v2.SamplingStrategyResponse{
		StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC,
		ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
			SamplingRate: 0.8,
		},
		OperationSampling: &api_v2.PerOperationSamplingStrategies{
			DefaultSamplingProbability: 0.8,
			PerOperationStrategies: []*api_v2.OperationSamplingStrategy{
				{
					Operation: "op1",
					ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
						SamplingRate: 0.2,
					},
				},
				{
					Operation: "op2",
					ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
						SamplingRate: 0.4,
					},
				},
			},
		},
	}

	resp, err := cl.GetSamplingStrategy(context.Background(), &api_v2.SamplingStrategyParameters{
		ServiceName: "foo",
	})
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestSamplingFailsOnNotConfigured(t *testing.T) {
	port := testutils.GetAvailablePort(t)
	// prepare
	config := &Configuration{
		CollectorGRPCPort: int(port),
	}
	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	assert.NoError(t, err, "should not have failed to create a new receiver")
	defer jr.Shutdown(context.Background())

	require.NoError(t, jr.Start(context.Background(), componenttest.NewNopHost()))
	t.Log("Start")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.CollectorGRPCPort), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	cl := api_v2.NewSamplingManagerClient(conn)

	response, err := cl.GetSamplingStrategy(context.Background(), &api_v2.SamplingStrategyParameters{
		ServiceName: "nothing",
	})
	require.NoError(t, err)
	assert.Equal(t, 0.001, response.GetProbabilisticSampling().GetSamplingRate())
}

func TestSamplingFailsOnBadFile(t *testing.T) {
	port := testutils.GetAvailablePort(t)
	// prepare
	config := &Configuration{
		CollectorGRPCPort:          int(port),
		RemoteSamplingStrategyFile: "does-not-exist",
	}
	sink := new(exportertest.SinkTraceExporter)

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	jr, err := New(jaegerReceiver, config, sink, params)
	assert.NoError(t, err, "should not have failed to create a new receiver")
	defer jr.Shutdown(context.Background())
	assert.Error(t, jr.Start(context.Background(), componenttest.NewNopHost()))
}
