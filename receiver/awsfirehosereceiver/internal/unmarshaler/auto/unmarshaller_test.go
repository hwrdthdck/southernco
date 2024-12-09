package auto

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog/compression"
)

func TestType(t *testing.T) {
	unmarshaler := NewUnmarshaler(zap.NewExample())
	require.Equal(t, TypeStr, unmarshaler.Type())
}

func TestUnmarshalMetrics_JSON(t *testing.T) {
	t.Parallel()

	unmarshaler := NewUnmarshaler(zap.NewNop())
	testCases := map[string]struct {
		dir                  string
		filename             string
		metricResourceCount  int
		metricCount          int
		metricDataPointCount int
		err                  error
	}{
		"cwmetric:WithMultipleRecords": {
			dir:                  "cwmetricstream",
			filename:             "multiple_records",
			metricResourceCount:  6,
			metricCount:          33,
			metricDataPointCount: 127,
		},
		"cwmetric:WithSingleRecord": {
			dir:                  "cwmetricstream",
			filename:             "single_record",
			metricResourceCount:  1,
			metricCount:          1,
			metricDataPointCount: 1,
		},
		"cwmetric:WithInvalidRecords": {
			dir:      "cwmetricstream",
			filename: "invalid_records",
			err:      errInvalidRecords,
		},
		"cwmetric:WithSomeInvalidRecords": {
			dir:                  "cwmetricstream",
			filename:             "some_invalid_records",
			metricResourceCount:  5,
			metricCount:          35,
			metricDataPointCount: 88,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			record, err := os.ReadFile(filepath.Join("..", testCase.dir, "testdata", testCase.filename))
			require.NoError(t, err)

			compressedRecord, err := compression.Zip(record)
			require.NoError(t, err)
			records := [][]byte{compressedRecord}

			metrics, err := unmarshaler.UnmarshalMetrics("application/json", records)
			require.Equal(t, testCase.err, err)

			require.Equal(t, testCase.metricResourceCount, metrics.ResourceMetrics().Len())
			require.Equal(t, testCase.metricDataPointCount, metrics.DataPointCount())
			require.Equal(t, testCase.metricCount, metrics.MetricCount())
		})
	}
}

// Unmarshall cloudwatch metrics and logs
func TestUnmarshalLogs_JSON(t *testing.T) {
	t.Parallel()

	unmarshaler := NewUnmarshaler(zap.NewNop())
	testCases := map[string]struct {
		dir              string
		filename         string
		logResourceCount int
		logRecordCount   int
		err              error
	}{
		"cwlog:WithMultipleRecords": {
			dir:              "cwlog",
			filename:         "multiple_records",
			logResourceCount: 1,
			logRecordCount:   2,
		},
		"cwlog:WithSingleRecord": {
			dir:              "cwlog",
			filename:         "single_record",
			logResourceCount: 1,
			logRecordCount:   1,
		},
		"cwlog:WithInvalidRecords": {
			dir:      "cwlog",
			filename: "invalid_records",
			err:      errInvalidRecords,
		},
		"cwlog:WithSomeInvalidRecords": {
			dir:              "cwlog",
			filename:         "some_invalid_records",
			logResourceCount: 1,
			logRecordCount:   2,
		},
		"cwlog:WithMultipleResources": {
			dir:              "cwlog",
			filename:         "multiple_resources",
			logResourceCount: 3,
			logRecordCount:   6,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			record, err := os.ReadFile(filepath.Join("..", testCase.dir, "testdata", testCase.filename))
			require.NoError(t, err)

			compressedRecord, err := compression.Zip(record)
			require.NoError(t, err)
			records := [][]byte{compressedRecord}

			logs, err := unmarshaler.UnmarshalLogs("application/json", records)
			require.Equal(t, testCase.err, err)

			require.Equal(t, testCase.logResourceCount, logs.ResourceLogs().Len())
			require.Equal(t, testCase.logRecordCount, logs.LogRecordCount())
		})
	}
}
