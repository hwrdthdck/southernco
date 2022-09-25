// Copyright  The OpenTelemetry Authors
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

package logs

import (
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
)

func TestTransform(t *testing.T) {
	traceID := [16]byte{0x08, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x0, 0x0, 0x0, 0x0, 0x0a}
	var spanID [8]byte
	copy(spanID[:], traceID[8:])
	ddTr := traceIDToUint64(traceID)
	ddSp := spanIDToUint64(spanID)

	type args struct {
		lr       plog.LogRecord
		res      pcommon.Resource
		sendBody bool
	}
	tests := []struct {
		name string
		args args
		want datadogV2.HTTPLogItem
	}{
		{
			name: "log_with_attribute",
			args: args{
				lr: func() plog.LogRecord {
					l := plog.NewLogRecord()
					l.Attributes().PutString("app", "test")
					l.SetSeverityNumber(5)
					return l
				}(),
				res: pcommon.NewResource(),
			},
			want: datadogV2.HTTPLogItem{
				Message: *datadog.PtrString(""),
				AdditionalProperties: map[string]string{
					"app":              "test",
					"status":           "debug",
					otelSeverityNumber: "5",
				},
			},
		},
		{
			name: "log_and_resource_with_attribute",
			args: args{
				lr: func() plog.LogRecord {
					l := plog.NewLogRecord()
					l.Attributes().PutString("app", "test")
					l.SetSeverityNumber(5)
					return l
				}(),
				res: func() pcommon.Resource {
					r := pcommon.NewResource()
					r.Attributes().PutString(conventions.AttributeServiceName, "otlp_col")
					return r
				}(),
			},
			want: datadogV2.HTTPLogItem{
				Ddtags:  datadog.PtrString("service:otlp_col"),
				Message: *datadog.PtrString(""),
				Service: datadog.PtrString("otlp_col"),
				AdditionalProperties: map[string]string{
					"app":              "test",
					"status":           "debug",
					otelSeverityNumber: "5",
				},
			},
		},
		{
			name: "log_with_service_attribute",
			args: args{
				lr: func() plog.LogRecord {
					l := plog.NewLogRecord()
					l.Attributes().PutString("app", "test")
					l.Attributes().PutString(conventions.AttributeServiceName, "otlp_col")
					l.SetSeverityNumber(5)
					return l
				}(),
				res: func() pcommon.Resource {
					r := pcommon.NewResource()
					return r
				}(),
			},
			want: datadogV2.HTTPLogItem{
				Message: *datadog.PtrString(""),
				Service: datadog.PtrString("otlp_col"),
				AdditionalProperties: map[string]string{
					"app":              "test",
					"status":           "debug",
					otelSeverityNumber: "5",
					"service.name":     "otlp_col",
				},
			},
		},
		{
			name: "log_with_trace",
			args: args{
				lr: func() plog.LogRecord {
					l := plog.NewLogRecord()
					l.Attributes().PutString("app", "test")
					l.SetSpanID(spanID)
					l.SetTraceID(traceID)
					l.Attributes().PutString(conventions.AttributeServiceName, "otlp_col")
					l.SetSeverityNumber(5)
					return l
				}(),
				res: func() pcommon.Resource {
					r := pcommon.NewResource()
					return r
				}(),
			},
			want: datadogV2.HTTPLogItem{
				Message: *datadog.PtrString(""),
				Service: datadog.PtrString("otlp_col"),
				AdditionalProperties: map[string]string{
					"app":              "test",
					"status":           "debug",
					otelSeverityNumber: "5",
					otelSpanID:         fmt.Sprintf("%x", string(spanID[:])),
					otelTraceID:        fmt.Sprintf("%x", string(traceID[:])),
					ddSpanID:           fmt.Sprintf("%d", ddSp),
					ddTraceID:          fmt.Sprintf("%d", ddTr),
					"service.name":     "otlp_col",
				},
			},
		},
		{
			// here SeverityText should take precedence for log status
			name: "log_with_severity_text_and_severity_number",
			args: args{
				lr: func() plog.LogRecord {
					l := plog.NewLogRecord()
					l.Attributes().PutString("app", "test")
					l.SetSpanID(spanID)
					l.SetTraceID(traceID)
					l.Attributes().PutString(conventions.AttributeServiceName, "otlp_col")
					l.SetSeverityText("alert")
					l.SetSeverityNumber(5)
					return l
				}(),
				res: func() pcommon.Resource {
					r := pcommon.NewResource()
					return r
				}(),
			},
			want: datadogV2.HTTPLogItem{
				Message: *datadog.PtrString(""),
				Service: datadog.PtrString("otlp_col"),
				AdditionalProperties: map[string]string{
					"app":              "test",
					"status":           "alert",
					otelSeverityText:   "alert",
					otelSeverityNumber: "5",
					otelSpanID:         fmt.Sprintf("%x", string(spanID[:])),
					otelTraceID:        fmt.Sprintf("%x", string(traceID[:])),
					ddSpanID:           fmt.Sprintf("%d", ddSp),
					ddTraceID:          fmt.Sprintf("%d", ddTr),
					"service.name":     "otlp_col",
				},
			},
		},
		{
			name: "log_with_message_from_body",
			args: args{
				lr: func() plog.LogRecord {
					l := plog.NewLogRecord()
					l.Attributes().PutString("app", "test")
					l.SetSpanID(spanID)
					l.SetTraceID(traceID)
					l.Attributes().PutString(conventions.AttributeServiceName, "otlp_col")
					l.SetSeverityNumber(13)
					l.Body().SetStringVal("This is log")
					return l
				}(),
				res: func() pcommon.Resource {
					r := pcommon.NewResource()
					return r
				}(),
				sendBody: true,
			},
			want: datadogV2.HTTPLogItem{
				Message: *datadog.PtrString(""),
				Service: datadog.PtrString("otlp_col"),
				AdditionalProperties: map[string]string{
					"message":          "This is log",
					"app":              "test",
					"status":           "warn",
					otelSeverityNumber: "13",
					otelSpanID:         fmt.Sprintf("%x", string(spanID[:])),
					otelTraceID:        fmt.Sprintf("%x", string(traceID[:])),
					ddSpanID:           fmt.Sprintf("%d", ddSp),
					ddTraceID:          fmt.Sprintf("%d", ddTr),
					"service.name":     "otlp_col",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Transform(tt.args.lr, tt.args.res, tt.args.sendBody)

			gs, err := got.MarshalJSON()
			if err != nil {
				t.Fatal(err)
				return
			}
			ws, err := tt.want.MarshalJSON()
			if err != nil {
				t.Fatal(err)
				return
			}
			if !assert.JSONEq(t, string(ws), string(gs)) {
				t.Errorf("Transform() = %v, want %v", string(gs), string(ws))
			}
		})
	}
}

func Test_deriveStatus(t *testing.T) {
	type args struct {
		severity plog.SeverityNumber
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "trace3",
			args: args{
				severity: 3,
			},
			want: logLevelTrace,
		},
		{
			name: "trace4",
			args: args{
				severity: 4,
			},
			want: logLevelTrace,
		},
		{
			name: "debug5",
			args: args{
				severity: 5,
			},
			want: logLevelDebug,
		},
		{
			name: "debug7",
			args: args{
				severity: 7,
			},
			want: logLevelDebug,
		},
		{
			name: "debug8",
			args: args{
				severity: 8,
			},
			want: logLevelDebug,
		},
		{
			name: "info9",
			args: args{
				severity: 9,
			},
			want: logLevelInfo,
		},
		{
			name: "info12",
			args: args{
				severity: 12,
			},
			want: logLevelInfo,
		},
		{
			name: "warn13",
			args: args{
				severity: 13,
			},
			want: logLevelWarn,
		},
		{
			name: "warn16",
			args: args{
				severity: 16,
			},
			want: logLevelWarn,
		},
		{
			name: "error17",
			args: args{
				severity: 17,
			},
			want: logLevelError,
		},
		{
			name: "error20",
			args: args{
				severity: 20,
			},
			want: logLevelError,
		},
		{
			name: "fatal21",
			args: args{
				severity: 21,
			},
			want: logLevelFatal,
		},
		{
			name: "fatal24",
			args: args{
				severity: 24,
			},
			want: logLevelFatal,
		},
		{
			name: "undefined",
			args: args{
				severity: 50,
			},
			want: logLevelError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, derviveStatusFromSeverityNumber(tt.args.severity), "derviveDdStatusFromSeverityNumber(%v)", tt.args.severity)
		})
	}
}
