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

package skywalkingexporter

import (
	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/translator/conventions/v1.5.0"
	common "skywalking.apache.org/repo/goapi/collect/common/v3"
	logpb "skywalking.apache.org/repo/goapi/collect/logging/v3"
	"strconv"
)

const (
	spanIDField            = "spanID"
	severityNumber         = "severityNumber"
	severityText           = "severityText"
	name                   = "name"
	flags                  = "flags"
	instrumentationName    = "otlp.name"
	instrumentationVersion = "otlp.version"
	defaultServiceName     = "otel-collector"
)

func logDataToLogRecode(ld pdata.Logs) []*logpb.LogData {
	lds := make([]*logpb.LogData, 0)
	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rl := rls.At(i)
		ills := rl.InstrumentationLibraryLogs()
		resource := rl.Resource()
		for j := 0; j < ills.Len(); j++ {
			ils := ills.At(j)
			logs := ils.Logs()
			for k := 0; k < logs.Len(); k++ {
				logData := &logpb.LogData{}
				logData.Tags = &logpb.LogTags{}
				resourceToLogData(resource, logData)
				instrumentationLibraryToLogContents(ils.InstrumentationLibrary(), logData)
				mapLogRecordToLogService(logs.At(k), logData)
				lds = append(lds, logData)
			}
		}
	}
	return lds
}

func resourceToLogData(resource pdata.Resource, logData *logpb.LogData) {
	attrs := resource.Attributes()

	if serviceName, ok := attrs.Get(conventions.AttributeServiceName); ok {
		logData.Service = pdata.AttributeValueToString(serviceName)
	} else {
		logData.Service = defaultServiceName
	}

	if serviceInstanceID, ok := attrs.Get(conventions.AttributeServiceInstanceID); ok {
		logData.ServiceInstance = pdata.AttributeValueToString(serviceInstanceID)
	}

	attrs.Range(func(k string, v pdata.AttributeValue) bool {
		if k == conventions.AttributeServiceName || k == conventions.AttributeServiceInstanceID {
			return true
		}
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   k,
			Value: pdata.AttributeValueToString(v),
		})
		return true
	})
}

func instrumentationLibraryToLogContents(instrumentationLibrary pdata.InstrumentationLibrary, logData *logpb.LogData) {
	if nameValue := instrumentationLibrary.Name(); nameValue != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   instrumentationName,
			Value: nameValue,
		})
	}
	if version := instrumentationLibrary.Version(); version != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   instrumentationVersion,
			Value: version,
		})
	}
}

func mapLogRecordToLogService(lr pdata.LogRecord, logData *logpb.LogData) {
	if lr.Body().Type() == pdata.AttributeValueTypeNull {
		return
	}

	if timestamp := lr.Timestamp(); timestamp > 0 {
		logData.Timestamp = lr.Timestamp().AsTime().UnixMilli()
	}

	if sn := strconv.FormatInt(int64(lr.SeverityNumber()), 10); sn != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   severityNumber,
			Value: sn,
		})
	}

	if st := lr.SeverityText(); st != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   severityText,
			Value: st,
		})
	}

	if ln := lr.Name(); ln != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   name,
			Value: ln,
		})
	}

	lr.Attributes().Range(func(k string, v pdata.AttributeValue) bool {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   k,
			Value: pdata.AttributeValueToString(v),
		})
		return true
	})

	logData.Body = &logpb.LogDataBody{
		Type: "body-type",
		Content: &logpb.LogDataBody_Text{
			Text: &logpb.TextLog{
				Text: pdata.AttributeValueToString(lr.Body()),
			}},
	}

	if flag := strconv.FormatUint(uint64(lr.Flags()), 16); flag != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   flags,
			Value: flag,
		})
	}

	if traceId := lr.TraceID().HexString(); traceId != "" {
		logData.TraceContext = &logpb.TraceContext{TraceId: traceId}
	}

	if spanId := lr.SpanID().HexString(); spanId != "" {
		logData.Tags.Data = append(logData.Tags.Data, &common.KeyStringValuePair{
			Key:   spanIDField,
			Value: spanId,
		})
	}
}
