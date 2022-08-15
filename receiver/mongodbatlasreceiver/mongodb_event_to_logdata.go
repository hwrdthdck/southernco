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

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"encoding/json"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

const (
	// Number of log attributes to add to the plog.LogRecordSlice.
	totalLogAttributes = 11

	// Number of resource attributes to add to the plog.ResourceLogs.
	totalResourceAttributes = 4
)

// layout for the timestamp format in the plog.Logs structure
const layout = "2006-01-02T15:04:05.000-07:00"

// Severity mapping of the mongodb atlas logs
var severityMap = map[string]plog.SeverityNumber{
	"F":  plog.SeverityNumberFATAL,
	"E":  plog.SeverityNumberERROR,
	"W":  plog.SeverityNumberWARN,
	"I":  plog.SeverityNumberINFO,
	"D":  plog.SeverityNumberDEBUG,
	"D1": plog.SeverityNumberDEBUG,
	"D2": plog.SeverityNumberDEBUG2,
	"D3": plog.SeverityNumberDEBUG3,
	"D4": plog.SeverityNumberDEBUG4,
	"D5": plog.SeverityNumberDEBUG4,
}

// mongoAuditEventToLogRecord converts model.AuditLog event to plog.LogRecordSlice and adds the resource attributes.
func mongodbAuditEventToLogData(logger *zap.Logger, logs []model.AuditLog, pc ProjectContext, hostname, logName, clusterName string) plog.Logs {
	ld := plog.NewLogs()
	rl := ld.ResourceLogs().AppendEmpty()
	sl := rl.ScopeLogs().AppendEmpty()

	resourceAttrs := rl.Resource().Attributes()
	resourceAttrs.EnsureCapacity(totalResourceAttributes)

	// Attributes related to the object causing the event.
	resourceAttrs.InsertString("mongodb_atlas.org", pc.orgName)
	resourceAttrs.InsertString("mongodb_atlas.project", pc.Project.Name)
	resourceAttrs.InsertString("mongodb_atlas.cluster", clusterName)
	resourceAttrs.InsertString("mongodb_atlas.host.name", hostname)

	for _, log := range logs {
		lr := sl.LogRecords().AppendEmpty()
		data, err := json.Marshal(log)
		if err != nil {
			logger.Warn("failed to marshal", zap.Error(err))
		}
		t, err := time.Parse(layout, log.Timestamp.Date)
		if err != nil {
			logger.Warn("Time failed to parse correctly", zap.Error(err))
		}
		lr.SetTimestamp(pcommon.NewTimestampFromTime(t))
		lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
		// Insert Raw Log message into Body of LogRecord
		lr.Body().SetStringVal(string(data))
		// Since Audit Logs don't have a severity/level
		// Set the "SeverityNumber" and "SeverityText" to INFO
		lr.SetSeverityNumber(plog.SeverityNumberINFO)
		lr.SetSeverityText("INFO")
		attrs := lr.Attributes()
		attrs.EnsureCapacity(totalLogAttributes)
		if log.AuthType != "" {
			attrs.InsertString("authtype", log.AuthType)
		}
		attrs.InsertString("local.ip", log.Local.IP)
		attrs.InsertInt("local.port", int64(log.Local.Port))
		attrs.InsertString("remote.ip", log.Remote.IP)
		attrs.InsertInt("remote.port", int64(log.Remote.Port))
		attrs.InsertString("uuid.binary", log.ID.Binary)
		attrs.InsertString("uuid.type", log.ID.Type)
		attrs.InsertInt("result", int64(log.Result))
		attrs.InsertString("log_name", logName)
		if log.Param.User != "" {
			attrs.InsertString("param.user", log.Param.User)
			attrs.InsertString("param.database", log.Param.Database)
			attrs.InsertString("param.mechanism", log.Param.Mechanism)
		}
	}

	return ld
}

// mongoEventToLogRecord converts model.LogEntry event to plog.LogRecordSlice and adds the resource attributes.
func mongodbEventToLogData(logger *zap.Logger, logs []model.LogEntry, pc ProjectContext, hostname, logName, clusterName string) plog.Logs {
	ld := plog.NewLogs()
	rl := ld.ResourceLogs().AppendEmpty()
	sl := rl.ScopeLogs().AppendEmpty()

	resourceAttrs := rl.Resource().Attributes()
	resourceAttrs.EnsureCapacity(totalResourceAttributes)

	// Attributes related to the object causing the event.
	resourceAttrs.InsertString("mongodb_atlas.org", pc.orgName)
	resourceAttrs.InsertString("mongodb_atlas.project", pc.Project.Name)
	resourceAttrs.InsertString("mongodb_atlas.cluster", clusterName)
	resourceAttrs.InsertString("mongodb_atlas.host.name", hostname)

	for _, log := range logs {
		lr := sl.LogRecords().AppendEmpty()
		data, err := json.Marshal(log)
		if err != nil {
			logger.Warn("failed to marshal", zap.Error(err))
		}
		t, err := time.Parse(layout, log.Timestamp.Date)
		if err != nil {
			logger.Warn("Time failed to parse correctly", zap.Error(err))
		}
		lr.SetTimestamp(pcommon.NewTimestampFromTime(t))
		lr.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
		// Insert Raw Log message into Body of LogRecord
		lr.Body().SetStringVal(string(data))
		// Set the "SeverityNumber" and "SeverityText" if a known type of
		// severity is found.
		if severityNumber, ok := severityMap[log.Severity]; ok {
			lr.SetSeverityNumber(severityNumber)
			lr.SetSeverityText(log.Severity)
		} else {
			logger.Debug("unknown severity type", zap.String("type", log.Severity))
		}
		attrs := lr.Attributes()
		attrs.EnsureCapacity(totalLogAttributes)
		pcommon.NewMapFromRaw(log.Attributes).CopyTo(attrs)
		attrs.InsertString("message", log.Message)
		attrs.InsertString("component", log.Component)
		attrs.InsertString("context", log.Context)
		attrs.InsertInt("id", log.ID)
		attrs.InsertString("log_name", logName)
		attrs.InsertString("raw", string(data))
	}

	return ld
}
