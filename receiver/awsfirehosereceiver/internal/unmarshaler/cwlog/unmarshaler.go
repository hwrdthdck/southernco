// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"

import (
	"bytes"
	"encoding/json"
	"errors"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler"
)

const (
	TypeStr         = "cwlogs"
	recordDelimiter = "\n"
)

var errInvalidRecords = errors.New("record format invalid")

// Unmarshaler for the CloudWatch Log JSON record format.
type Unmarshaler struct {
	logger *zap.Logger
}

var _ unmarshaler.LogsUnmarshaler = (*Unmarshaler)(nil)

// NewUnmarshaler creates a new instance of the Unmarshaler.
func NewUnmarshaler(logger *zap.Logger) *Unmarshaler {
	return &Unmarshaler{logger}
}

// UnmarshalLogs deserializes the records into CWLog and uses the
// ResourceLogsBuilder to group them into a single plog.Logs.
// Skips invalid CWLog received in the record.
func (u Unmarshaler) UnmarshalLogs(records [][]byte) (plog.Logs, error) {
	md := plog.NewLogs()
	builders := make(map[ResourceAttributes]*ResourceLogsBuilder)
	for recordIndex, record := range records {
		for datumIndex, datum := range bytes.Split(record, []byte(recordDelimiter)) {
			var log CWLog
			err := json.Unmarshal(datum, &log)
			if err != nil {
				u.logger.Error(
					"Unable to unmarshal input",
					zap.Error(err),
					zap.Int("datum_index", datumIndex),
					zap.Int("record_index", recordIndex),
				)
				continue
			}
			if !u.isValid(log) {
				u.logger.Error(
					"Invalid log",
					zap.Int("datum_index", datumIndex),
					zap.Int("record_index", recordIndex),
				)
				continue
			}
			attrs := ResourceAttributes{
				Owner:     log.Owner,
				LogGroup:  log.LogGroup,
				LogStream: log.LogStream,
			}
			lb, ok := builders[attrs]
			if !ok {
				lb = NewResourceLogsBuilder(md, attrs)
				builders[attrs] = lb
			}
			lb.AddLog(log)
		}
	}

	if len(builders) == 0 {
		return plog.NewLogs(), errInvalidRecords
	}

	return md, nil
}

// isValid validates that the CWLog has been unmarshalled correctly.
func (u Unmarshaler) isValid(log CWLog) bool {
	return log.Owner != "" && log.LogGroup != "" && log.LogStream != ""
}

// Type of the serialized messages.
func (u Unmarshaler) Type() string {
	return TypeStr
}
