// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"context"
	"testing"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	semconv "go.opentelemetry.io/collector/semconv/v1.25.0"
)

func TestStartTraces(t *testing.T) {
	config := createDefaultConfig().(*Config)
	config.MySQLEndpoint = "127.0.0.1:9030"
	config.Username = "admin"
	config.Password = "admin"
	config.Database = "otel2"
	config.Table.Traces = "traces"
	config.HistoryDays = 3

	exporter, err := newTracesExporter(nil, config)
	ctx := context.Background()
	if err != nil {
		t.Error(err)
		return
	}
	defer exporter.shutdown(ctx)

	err = exporter.start(ctx, nil)
	if err != nil {
		t.Error(err)
		return
	}
}

var traces = []*dTrace{
	{
		ServiceName:    "service_name",
		Timestamp:      time.Now().Format(timeFormat),
		TraceID:        "trace_id",
		EndTime:        time.Now().Format(timeFormat),
		Duration:       1000,
		SpanAttributes: map[string]any{"k": "v", "a": 1},
		Events: []*dEvent{
			{
				Timestamp:  time.Now().Format(timeFormat),
				Name:       "event_name",
				Attributes: map[string]any{"k": "v", "a": 1},
			},
		},
		Links: []*dLink{
			{
				TraceID:    "trace_id",
				SpanID:     "span_id",
				TraceState: "trace_state",
				Attributes: map[string]any{"k": "v", "a": 1},
			},
		},
		ResourceAttributes: map[string]any{"k": "v", "a": 1},
	},
	{
		ServiceName:    "service_name",
		Timestamp:      time.Now().Format(timeFormat),
		TraceID:        "trace_id",
		EndTime:        time.Now().Format(timeFormat),
		Duration:       1000,
		SpanAttributes: map[string]any{"k": "v", "a": 1},
		Events: []*dEvent{
			{
				Timestamp:  time.Now().Format(timeFormat),
				Name:       "event_name",
				Attributes: map[string]any{"k": "v", "a": 1},
			},
		},
		Links: []*dLink{
			{
				TraceID:    "trace_id",
				SpanID:     "span_id",
				TraceState: "trace_state",
				Attributes: map[string]any{"k": "v", "a": 1},
			},
		},
		ResourceAttributes: map[string]any{"k": "v", "a": 1},
	},
}

func TestPushTraceDataInternal(t *testing.T) {
	config := createDefaultConfig().(*Config)
	config.Endpoint = "http://127.0.0.1:8030"
	config.MySQLEndpoint = "127.0.0.1:9030"
	config.Username = "admin"
	config.Password = "admin"
	config.Database = "otel2"
	config.Table.Traces = "traces"
	config.HistoryDays = 3

	exporter, err := newTracesExporter(nil, config)
	ctx := context.Background()
	if err != nil {
		t.Error(err)
		return
	}
	defer exporter.shutdown(ctx)

	err = exporter.start(ctx, nil)
	if err != nil {
		t.Error(err)
		return
	}

	// no timezone
	err = exporter.pushTraceDataInternal(ctx, traces)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestPushTraceData(t *testing.T) {
	config := createDefaultConfig().(*Config)
	config.Endpoint = "http://127.0.0.1:8030"
	config.MySQLEndpoint = "127.0.0.1:9030"
	config.Username = "admin"
	config.Password = "admin"
	config.Database = "otel2"
	config.Table.Traces = "traces"
	config.HistoryDays = 3
	config.TimeZone = "America/New_York"

	err := config.Validate()
	if err != nil {
		t.Error(err)
		return
	}

	exporter, err := newTracesExporter(nil, config)
	ctx := context.Background()
	if err != nil {
		t.Error(err)
		return
	}
	defer exporter.shutdown(ctx)

	err = exporter.start(ctx, nil)
	if err != nil {
		t.Error(err)
		return
	}

	err = exporter.pushTraceData(ctx, simpleTraces(10))
	if err != nil {
		t.Error(err)
		return
	}
}

func simpleTraces(count int) ptrace.Traces {
	traces := ptrace.NewTraces()
	rs := traces.ResourceSpans().AppendEmpty()
	rs.SetSchemaUrl("https://opentelemetry.io/schemas/1.4.0")
	rs.Resource().SetDroppedAttributesCount(10)
	rs.Resource().Attributes().PutStr("service.name", "test-service")
	ss := rs.ScopeSpans().AppendEmpty()
	ss.Scope().SetName("io.opentelemetry.contrib.doris")
	ss.Scope().SetVersion("1.0.0")
	ss.SetSchemaUrl("https://opentelemetry.io/schemas/1.7.0")
	ss.Scope().SetDroppedAttributesCount(20)
	ss.Scope().Attributes().PutStr("lib", "doris")
	timestamp := time.Now()
	for i := 0; i < count; i++ {
		s := ss.Spans().AppendEmpty()
		s.SetTraceID([16]byte{1, 2, 3, byte(i)})
		s.SetSpanID([8]byte{1, 2, 3, byte(i)})
		s.TraceState().FromRaw("trace state")
		s.SetParentSpanID([8]byte{1, 2, 4, byte(i)})
		s.SetName("call db")
		s.SetKind(ptrace.SpanKindInternal)
		s.SetStartTimestamp(pcommon.NewTimestampFromTime(timestamp))
		s.SetEndTimestamp(pcommon.NewTimestampFromTime(timestamp.Add(time.Minute)))
		s.Attributes().PutStr(semconv.AttributeServiceName, "v")
		s.Status().SetMessage("error")
		s.Status().SetCode(ptrace.StatusCodeError)
		event := s.Events().AppendEmpty()
		event.SetName("event1")
		event.SetTimestamp(pcommon.NewTimestampFromTime(timestamp))
		event.Attributes().PutStr("level", "info")
		link := s.Links().AppendEmpty()
		link.SetTraceID([16]byte{1, 2, 5, byte(i)})
		link.SetSpanID([8]byte{1, 2, 5, byte(i)})
		link.TraceState().FromRaw("error")
		link.Attributes().PutStr("k", "v")
	}
	return traces
}
