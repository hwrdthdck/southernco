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

package kafkareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"

import (
	"bytes"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/relvacode/iso8601"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.13.0"
	"go.uber.org/zap"
)

const (
	// Constants for OpenTelemetry Specs
	receiverScopeName = "otelcol/" + typeStr

	// Constants for Azure Log Records
	azureCategory          = "azure.category"
	azureCorrelationID     = "azure.correlation.id"
	azureDuration          = "azure.duration"
	azureIdentity          = "azure.identity"
	azureOperationName     = "azure.operation.name"
	azureOperationVersion  = "azure.operation.version"
	azureProperties        = "azure.properties"
	azureResourceID        = "azure.resource.id"
	azureResultType        = "azure.result.type"
	azureResultSignature   = "azure.result.signature"
	azureResultDescription = "azure.result.description"
	azureTenantID          = "azure.tenant.id"
)

type azureResourceLogsUnmarshaler struct {
	settings receiver.CreateSettings
}

// azureRecords represents an array of Azure log records
// as exported via an Azure Event Hub
type azureRecords struct {
	Records []azureLogRecord `json:"records"`
}

// azureLogRecord represents a single Azure log following
// the common schema:
// https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/resource-logs-schema
type azureLogRecord struct {
	Time              string       `json:"time"`
	ResourceID        string       `json:"resourceId"`
	TenantID          *string      `json:"tenantId"`
	OperationName     string       `json:"operationName"`
	OperationVersion  *string      `json:"operationVersion"`
	Category          string       `json:"category"`
	ResultType        *string      `json:"resultType"`
	ResultSignature   *string      `json:"resultSignature"`
	ResultDescription *string      `json:"resultDescription"`
	DurationMs        *string      `json:"durationMs"`
	CallerIPAddress   *string      `json:"callerIpAddress"`
	CorrelationID     *string      `json:"correlationId"`
	Identity          *interface{} `json:"identity"`
	Level             *string      `json:"Level"`
	Location          *string      `json:"location"`
	Properties        *interface{} `json:"properties"`
}

func newAzureResourceLogsUnmarshaler() LogsUnmarshaler {
	return azureResourceLogsUnmarshaler{}
}

func (r azureResourceLogsUnmarshaler) Unmarshal(buf []byte) (plog.Logs, error) {
	l := plog.NewLogs()

	var azureLogs azureRecords
	decoder := jsoniter.NewDecoder(bytes.NewReader(buf))
	if err := decoder.Decode(&azureLogs); err != nil {
		return l, err
	}

	resourceLogs := l.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
	scopeLogs.Scope().SetName(receiverScopeName)
	scopeLogs.Scope().SetVersion(r.settings.BuildInfo.Version)
	logRecords := scopeLogs.LogRecords()

	resourceID := ""
	for _, azureLog := range azureLogs.Records {
		resourceID = azureLog.ResourceID
		nanos, err := asTimestamp(azureLog.Time)
		if err != nil {
			r.settings.Logger.Warn("Unable to convert timestamp from log", zap.String("timestamp", azureLog.Time))
			continue
		}

		lr := logRecords.AppendEmpty()

		lr.SetTimestamp(nanos)

		if azureLog.Level != nil {
			severity := asSeverity(*azureLog.Level)
			lr.SetSeverityNumber(severity)
			lr.SetSeverityText(*azureLog.Level)
		}

		if err := lr.Attributes().FromRaw(extractRawAttributes(azureLog)); err != nil {
			return l, err
		}

		// The Azure resource ID will be pulled into a common resource attribute.
		// This implementation assumes that a single log message from Azure will
		// contain ONLY logs from a single resource.
		if resourceID != "" {
			resourceLogs.Resource().Attributes().PutStr(azureResourceID, resourceID)
		}
	}

	return l, nil
}

func (r azureResourceLogsUnmarshaler) Encoding() string {
	return "azureresourcelogs"
}

func (r azureResourceLogsUnmarshaler) SetReceiverSettings(settings receiver.CreateSettings) {
	r.settings = settings
}

// asTimestamp will parse an ISO8601 string into an OpenTelemetry
// nanosecond timestamp. If the string cannot be parsed, it will
// return zero and the error.
func asTimestamp(s string) (pcommon.Timestamp, error) {
	t, err := iso8601.ParseString(s)
	if err != nil {
		return 0, err
	}
	return pcommon.Timestamp(t.UnixNano()), nil
}

// asSeverity converts the Azure log level to equivalent
// OpenTelemetry severity numbers. If the log level is not
// valid, then the 'Unspecified' value is returned.
func asSeverity(s string) plog.SeverityNumber {
	switch s {
	case "Informational":
		return plog.SeverityNumberInfo
	case "Warning":
		return plog.SeverityNumberWarn
	case "Error":
		return plog.SeverityNumberError
	case "Critical":
		return plog.SeverityNumberFatal
	default:
		return plog.SeverityNumberUnspecified
	}
}

func extractRawAttributes(log azureLogRecord) map[string]interface{} {
	var attrs = map[string]interface{}{}

	attrs[azureCategory] = log.Category
	setIf(attrs, azureCorrelationID, log.CorrelationID)
	if log.DurationMs != nil {
		duration, err := strconv.ParseInt(*log.DurationMs, 10, 64)
		if err == nil {
			attrs[azureDuration] = duration
		}
	}
	if log.Identity != nil {
		attrs[azureIdentity] = *log.Identity
	}
	attrs[azureOperationName] = log.OperationName
	setIf(attrs, azureOperationVersion, log.OperationVersion)
	if log.Properties != nil {
		attrs[azureProperties] = *log.Properties
	}
	setIf(attrs, azureResultDescription, log.ResultDescription)
	setIf(attrs, azureResultSignature, log.ResultSignature)
	setIf(attrs, azureResultType, log.ResultType)
	setIf(attrs, azureTenantID, log.TenantID)

	setIf(attrs, conventions.AttributeCloudRegion, log.Location)
	attrs[conventions.AttributeCloudProvider] = conventions.AttributeCloudProviderAzure

	setIf(attrs, conventions.AttributeNetSockPeerAddr, log.CallerIPAddress)
	return attrs
}

func setIf(attrs map[string]interface{}, key string, value *string) {
	if value != nil && *value != "" {
		attrs[key] = *value
	}
}
