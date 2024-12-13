// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package simplespan // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/simplespan"

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/eventtime"
)

type FieldMapConfig struct {
	Resources  ResourcesConfig  `mapstructure:"resources"`
	Scopes     ScopesConfig     `mapstructure:"scopes"`
	Attributes AttributesConfig `mapstructure:"attributes"`
}

type ResourcesConfig struct {
	ServiceName string `mapstructure:"service_name"`
}

type ScopesConfig struct {
	LibraryName    string `mapstructure:"library_name"`
	LibraryVersion string `mapstructure:"library_version"`
}

type AttributesConfig struct {
	TraceID        string   `mapstructure:"trace_id"`
	ParentID       string   `mapstructure:"parent_id"`
	SpanID         string   `mapstructure:"span_id"`
	Name           string   `mapstructure:"name"`
	Error          string   `mapstructure:"error"`
	SpanKind       string   `mapstructure:"spankind"`
	DurationFields []string `mapstructure:"durationFields"`
}

type SimpleSpan struct {
	Samplerate       int                    `json:"samplerate" msgpack:"samplerate"`
	MsgPackTimestamp *time.Time             `msgpack:"time"`
	Time             string                 `json:"time"` // should not be trusted. use MsgPackTimestamp
	Data             map[string]interface{} `json:"data" msgpack:"data"`
}

// Overrides unmarshall to make sure the MsgPackTimestamp is set
func (s *SimpleSpan) UnmarshalJSON(j []byte) error {
	type _simpleSpan SimpleSpan
	tstr := eventtime.GetEventTimeDefaultString()
	tzero := time.Time{}
	tmp := _simpleSpan{Time: "none", MsgPackTimestamp: &tzero, Samplerate: 1}

	err := json.Unmarshal(j, &tmp)
	if err != nil {
		return err
	}
	if tmp.MsgPackTimestamp.IsZero() && tmp.Time == "none" {
		// neither timestamp was set. give it right now.
		tmp.Time = tstr
		tnow := time.Now()
		tmp.MsgPackTimestamp = &tnow
	}
	if tmp.MsgPackTimestamp.IsZero() {
		propertime := eventtime.GetEventTime(tmp.Time)
		tmp.MsgPackTimestamp = &propertime
	}

	*s = SimpleSpan(tmp)
	return nil
}

func (s *SimpleSpan) DebugString() string {
	return fmt.Sprintf("%#v", s)
}

// returns log until we add the trace parser
func (s *SimpleSpan) SignalType() (string, error) {
	return "log", nil
}

func (s *SimpleSpan) GetService(fields FieldMapConfig, seen *ServiceHistory, dataset string) (string, error) {
	if serviceName, ok := s.Data[fields.Resources.ServiceName]; ok {
		seen.NameCount[serviceName.(string)] += 1
		return serviceName.(string), nil
	}
	return dataset, errors.New("no service.name found in event")
}

func (s *SimpleSpan) GetScope(fields FieldMapConfig, seen *ScopeHistory, serviceName string) (string, error) {
	if scopeLibraryName, ok := s.Data[fields.Scopes.LibraryName]; ok {
		scopeKey := serviceName + scopeLibraryName.(string)
		if _, ok := seen.Scope[scopeKey]; ok {
			// if we've seen it, we don't expect it to be different right away so we'll just return it.
			return scopeKey, nil
		}
		// otherwise, we need to make a new found scope
		scopeLibraryVersion := "unset"
		if scopeLibVer, ok := s.Data[fields.Scopes.LibraryVersion]; ok {
			scopeLibraryVersion = scopeLibVer.(string)
		}
		newScope := SimpleScope{
			ServiceName:    serviceName, // we only set the service name once. If the same library comes from multiple services in the same batch, we're in trouble.
			LibraryName:    scopeLibraryName.(string),
			LibraryVersion: scopeLibraryVersion,
			ScopeSpans:     ptrace.NewSpanSlice(),
			ScopeLogs:      plog.NewLogRecordSlice(),
		}
		seen.Scope[scopeKey] = newScope
		return scopeKey, nil
	}
	return "libhoney.receiver", errors.New("library name not found")
}

type SimpleScope struct {
	ServiceName    string
	LibraryName    string
	LibraryVersion string
	ScopeSpans     ptrace.SpanSlice
	ScopeLogs      plog.LogRecordSlice
}

type ScopeHistory struct {
	Scope map[string]SimpleScope // key here is service.name+library.name
}
type ServiceHistory struct {
	NameCount map[string]int
}

func (s *SimpleSpan) ToPLogRecord(newLog *plog.LogRecord, already_used_fields *[]string, logger zap.Logger) error {
	time_ns := s.MsgPackTimestamp.UnixNano()
	logger.Debug("processing log with", zap.Int64("timestamp", time_ns))
	newLog.SetTimestamp(pcommon.Timestamp(time_ns))

	if logSevCode, ok := s.Data["severity_code"]; ok {
		logSevInt := int32(logSevCode.(int64))
		newLog.SetSeverityNumber(plog.SeverityNumber(logSevInt))
	}

	if logSevText, ok := s.Data["severity_text"]; ok {
		newLog.SetSeverityText(logSevText.(string))
	}

	if logFlags, ok := s.Data["flags"]; ok {
		logFlagsUint := uint32(logFlags.(uint64))
		newLog.SetFlags(plog.LogRecordFlags(logFlagsUint))
	}

	// undoing this is gonna be complicated: https://github.com/honeycombio/husky/blob/91c0498333cd9f5eed1fdb8544ca486db7dea565/otlp/logs.go#L61
	if logBody, ok := s.Data["body"]; ok {
		newLog.Body().SetStr(logBody.(string))
	}

	newLog.Attributes().PutInt("SampleRate", int64(s.Samplerate))

	logFieldsAlready := []string{"severity_text", "severity_code", "flags", "body"}
	for k, v := range s.Data {
		if slices.Contains(*already_used_fields, k) {
			continue
		}
		if slices.Contains(logFieldsAlready, k) {
			continue
		}
		switch v := v.(type) {
		case string:
			newLog.Attributes().PutStr(k, v)
		case int:
			newLog.Attributes().PutInt(k, int64(v))
		case int64, int16, int32:
			intv := v.(int64)
			newLog.Attributes().PutInt(k, intv)
		case float64:
			newLog.Attributes().PutDouble(k, v)
		case bool:
			newLog.Attributes().PutBool(k, v)
		default:
			logger.Warn("Span data type issue", zap.Int64("timestamp", time_ns), zap.String("key", k))
		}
	}
	return nil
}
