// Copyright  OpenTelemetry Authors
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

package observiqexporter

import (
	"crypto/md5" //nolint:gosec // Not used for crypto, just for generating ID based on contents of log
	"encoding/hex"
	"encoding/json"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type preEncodedJSON aliases []byte, represents JSON that has already been encoded
// Does not support unmarshalling
type preEncodedJSON []byte

type observIQLogBatch struct {
	Logs []*observIQLog `json:"logs"`
}

type observIQLog struct {
	ID    string         `json:"id"`
	Size  int            `json:"size"`
	Entry preEncodedJSON `json:"entry"`
}

type observIQAgentInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
type observIQLogEntry struct {
	Timestamp string                 `json:"@timestamp"`
	Severity  string                 `json:"severity,omitempty"`
	EntryType string                 `json:"type,omitempty"`
	Message   string                 `json:"message,omitempty"`
	Resource  interface{}            `json:"resource,omitempty"`
	Agent     *observIQAgentInfo     `json:"agent,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Body      map[string]interface{} `json:"body,omitempty"`
}

// Convert pdata.Logs to observIQLogBatch
func logdataToObservIQFormat(ld pdata.Logs, agentID string, agentName string, buildInfo component.BuildInfo) (*observIQLogBatch, []error) {
	var rls = ld.ResourceLogs()
	var sliceOut = make([]*observIQLog, 0, ld.LogRecordCount())
	var errorsOut = make([]error, 0)

	for i := 0; i < rls.Len(); i++ {
		rl := rls.At(i)
		res := rl.Resource()
		resMap := attributeMapToBaseType(res.Attributes())
		ills := rl.InstrumentationLibraryLogs()
		for j := 0; j < ills.Len(); j++ {
			ill := ills.At(j)
			logs := ill.Logs()
			for k := 0; k < logs.Len(); k++ {
				oiqLogEntry := resourceAndInstrmentationLogToEntry(resMap, logs.At(k), agentID, agentName, buildInfo)

				jsonOIQLogEntry, err := json.Marshal(oiqLogEntry)

				if err != nil {
					//Skip this log, keep record of error
					errorsOut = append(errorsOut, consumererror.Permanent(err))
					continue
				}

				//md5 sum of the message is ID
				md5Sum := md5.Sum(jsonOIQLogEntry) //nolint:gosec // Not used for crypto, just for generating ID based on contents of log
				md5SumAsHex := hex.EncodeToString(md5Sum[:])

				sliceOut = append(sliceOut, &observIQLog{
					ID:    md5SumAsHex,
					Size:  len(jsonOIQLogEntry),
					Entry: preEncodedJSON(jsonOIQLogEntry),
				})
			}
		}
	}

	return &observIQLogBatch{Logs: sliceOut}, errorsOut
}

// Output timestamp format, an ISO8601 compliant timesetamp with millisecond precision
const timestampFieldOutputLayout = "2006-01-02T15:04:05.000Z07:00"

func resourceAndInstrmentationLogToEntry(resMap interface{}, log pdata.LogRecord, agentID string, agentName string, buildInfo component.BuildInfo) *observIQLogEntry {
	msg := messageFromRecord(log)

	return &observIQLogEntry{
		Timestamp: timestampFromRecord(log),
		Severity:  severityFromRecord(log),
		Resource:  resMap,
		Message:   msg,
		Data:      attributeMapToBaseType(log.Attributes()),
		Body:      bodyFromRecord(log),
		Agent:     &observIQAgentInfo{Name: agentName, ID: agentID, Version: buildInfo.Version},
	}
}

func timestampFromRecord(log pdata.LogRecord) string {
	if log.Timestamp() == 0 {
		return timeNow().UTC().Format(timestampFieldOutputLayout)
	}
	return log.Timestamp().AsTime().UTC().Format(timestampFieldOutputLayout)
}

func messageFromRecord(log pdata.LogRecord) string {
	if log.Body().Type() == pdata.AttributeValueTypeString {
		return log.Body().StringVal()
	}

	return ""
}

// If Body is a map, it is suitable to be used on the observIQ log entry as "body"
func bodyFromRecord(log pdata.LogRecord) map[string]interface{} {
	if log.Body().Type() == pdata.AttributeValueTypeMap {
		return attributeMapToBaseType(log.Body().MapVal())
	}
	return nil
}

//Mappings from opentelemetry severity number to observIQ severity string
var severityNumberToObservIQName = map[int32]string{
	0:  "default",
	1:  "trace",
	2:  "trace",
	3:  "trace",
	4:  "trace",
	5:  "debug",
	6:  "debug",
	7:  "debug",
	8:  "debug",
	9:  "info",
	10: "info",
	11: "info",
	12: "info",
	13: "warn",
	14: "warn",
	15: "warn",
	16: "warn",
	17: "error",
	18: "error",
	19: "error",
	20: "error",
	21: "fatal",
	22: "fatal",
	23: "fatal",
	24: "fatal",
}

/*
	Get severity from the a log record.
	We prefer the severity number, and map it to a string
	representing the opentelemetry defined severity.
	If there is no severity number, we use "default"
*/
func severityFromRecord(log pdata.LogRecord) string {
	var sevAsInt32 = int32(log.SeverityNumber())
	if sevAsInt32 < int32(len(severityNumberToObservIQName)) && sevAsInt32 >= 0 {
		return severityNumberToObservIQName[sevAsInt32]
	}
	return "default"
}

/*
	Transform AttributeMap to native Go map, skipping keys with nil values, and replacing dots in keys with _
*/
func attributeMapToBaseType(m pdata.AttributeMap) map[string]interface{} {
	mapOut := make(map[string]interface{}, m.Len())
	m.Range(func(k string, v pdata.AttributeValue) bool {
		val := attributeValueToBaseType(v)
		if val != nil {
			dedotedKey := strings.ReplaceAll(k, ".", "_")
			mapOut[dedotedKey] = val
		}
		return true
	})
	return mapOut
}

/*
	attrib is the attribute value to convert to it's native Go type - skips nils in arrays/maps
*/
func attributeValueToBaseType(attrib pdata.AttributeValue) interface{} {
	switch attrib.Type() {
	case pdata.AttributeValueTypeString:
		return attrib.StringVal()
	case pdata.AttributeValueTypeBool:
		return attrib.BoolVal()
	case pdata.AttributeValueTypeInt:
		return attrib.IntVal()
	case pdata.AttributeValueTypeDouble:
		return attrib.DoubleVal()
	case pdata.AttributeValueTypeMap:
		attribMap := attrib.MapVal()
		return attributeMapToBaseType(attribMap)
	case pdata.AttributeValueTypeArray:
		arrayVal := attrib.ArrayVal()
		slice := make([]interface{}, 0, arrayVal.Len())
		for i := 0; i < arrayVal.Len(); i++ {
			val := attributeValueToBaseType(arrayVal.At(i))
			if val != nil {
				slice = append(slice, val)
			}
		}
		return slice
	case pdata.AttributeValueTypeNull:
		return nil
	}
	return nil
}

// The marshaled JSON is just the []byte that this type aliases.
func (p preEncodedJSON) MarshalJSON() ([]byte, error) {
	return p, nil
}
