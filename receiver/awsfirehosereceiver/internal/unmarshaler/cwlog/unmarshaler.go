// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"
)

const (
	TypeStr         = "cwlogs"
	recordDelimiter = "\n"

	attributeAWSCloudWatchLogGroupName  = "aws.cloudwatch.log_group_name"
	attributeAWSCloudWatchLogStreamName = "aws.cloudwatch.log_stream_name"
)

var errInvalidRecords = errors.New("record format invalid")

// Unmarshaler for the CloudWatch Log JSON record format.
type Unmarshaler struct {
	logger *zap.Logger
}

var _ plog.Unmarshaler = (*Unmarshaler)(nil)

// NewUnmarshaler creates a new instance of the Unmarshaler.
func NewUnmarshaler(logger *zap.Logger) *Unmarshaler {
	return &Unmarshaler{logger}
}

// Unmarshal deserializes the records into cWLogs and uses the
// resourceLogsBuilder to group them into a single plog.Logs.
// Skips invalid cWLogs received in the record and
func (u Unmarshaler) UnmarshalLogs(compressedRecord []byte) (plog.Logs, error) {
	r, err := gzip.NewReader(bytes.NewReader(compressedRecord))
	if err != nil {
		return plog.Logs{}, fmt.Errorf("failed to decompress record: %w", err)
	}
	decoder := json.NewDecoder(r)

	// Multiple logs in each record separated by newline character
	logs := plog.NewLogs()
	for datumIndex := 0; ; datumIndex++ {
		var log cWLog
		if err := decoder.Decode(&log); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			u.logger.Error(
				"Unable to unmarshal input",
				zap.Error(err),
				zap.Int("datum_index", datumIndex),
			)
			continue
		}
		if !u.isValid(log) {
			u.logger.Error(
				"Invalid log",
				zap.Int("datum_index", datumIndex),
			)
			continue
		}

		rl := logs.ResourceLogs().AppendEmpty()
		resourceAttrs := rl.Resource().Attributes()
		resourceAttrs.PutStr(conventions.AttributeCloudAccountID, log.Owner)
		resourceAttrs.PutStr(attributeAWSCloudWatchLogGroupName, log.LogGroup)
		resourceAttrs.PutStr(attributeAWSCloudWatchLogStreamName, log.LogStream)

		logRecords := rl.ScopeLogs().AppendEmpty().LogRecords()
		for _, event := range log.LogEvents {
			logRecord := logRecords.AppendEmpty()
			// pcommon.Timestamp is a time specified as UNIX Epoch time in nanoseconds
			// but timestamp in cloudwatch logs are in milliseconds.
			logRecord.SetTimestamp(pcommon.Timestamp(event.Timestamp * int64(time.Millisecond)))
			logRecord.Body().SetStr(event.Message)
		}
	}
	if logs.ResourceLogs().Len() == 0 {
		return logs, errInvalidRecords
	}
	pdatautil.GroupByResourceLogs(logs.ResourceLogs())
	return logs, nil
}

// isValid validates that the cWLog has been unmarshalled correctly.
func (u Unmarshaler) isValid(log cWLog) bool {
	return log.Owner != "" && log.LogGroup != "" && log.LogStream != ""
}

// Type of the serialized messages.
func (u Unmarshaler) Type() string {
	return TypeStr
}
