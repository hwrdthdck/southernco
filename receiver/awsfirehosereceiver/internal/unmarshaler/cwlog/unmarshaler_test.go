// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlog

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestType(t *testing.T) {
	unmarshaler := NewUnmarshaler(zap.NewNop())
	require.Equal(t, TypeStr, unmarshaler.Type())
}

func TestUnmarshal(t *testing.T) {
	unmarshaler := NewUnmarshaler(zap.NewNop())
	testCases := map[string]struct {
		filename          string
		wantResourceCount int
		wantLogCount      int
		wantErr           error
	}{
		"WithMultipleRecords": {
			filename:          "multiple_records",
			wantResourceCount: 1,
			wantLogCount:      2,
		},
		"WithSingleRecord": {
			filename:          "single_record",
			wantResourceCount: 1,
			wantLogCount:      1,
		},
		"WithInvalidRecords": {
			filename: "invalid_records",
			wantErr:  errInvalidRecords,
		},
		"WithSomeInvalidRecords": {
			filename:          "some_invalid_records",
			wantResourceCount: 1,
			wantLogCount:      2,
		},
		"WithMultipleResources": {
			filename:          "multiple_resources",
			wantResourceCount: 3,
			wantLogCount:      6,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(".", "testdata", testCase.filename))
			require.NoError(t, err)

			var records [][]byte
			for _, record := range bytes.Split(data, []byte("\n")) {
				records = append(records, record)
			}

			got, err := unmarshaler.UnmarshalLogs(records)
			require.Equal(t, testCase.wantErr, err)
			require.NotNil(t, got)
			require.Equal(t, testCase.wantResourceCount, got.ResourceLogs().Len())
			require.Equal(t, testCase.wantLogCount, got.LogRecordCount())
		})
	}
}

func TestLogTimestamp(t *testing.T) {
	unmarshaler := NewUnmarshaler(zap.NewNop())
	record, err := os.ReadFile(filepath.Join(".", "testdata", "single_record"))
	require.NoError(t, err)

	records := [][]byte{record}

	got, err := unmarshaler.UnmarshalLogs(records)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, 1, got.ResourceLogs().Len())

	rm := got.ResourceLogs().At(0)
	require.Equal(t, 1, rm.ScopeLogs().Len())
	ilm := rm.ScopeLogs().At(0)
	ilm.LogRecords().At(0).Timestamp()
	expectedTimestamp := "2024-09-05 13:47:15.523 +0000 UTC"
	require.Equal(t, expectedTimestamp, ilm.LogRecords().At(0).Timestamp().String())
}
