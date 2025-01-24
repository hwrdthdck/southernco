// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"testing"

	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
)

func BenchmarkLogsConsumer_cwlogs(b *testing.B) {
	// numLogGroups is the maximum number of unique log groups
	// to use across the generated logs, using a random generator.
	const numLogGroups = 10

	// numRecords is the number of records in the Firehose envelope.
	for _, numRecords := range []int{10, 100} {
		// numLogs is the number of CoudWatch log records within a Firehose record.
		for _, numLogs := range []int{1, 10} {
			b.Run(fmt.Sprintf("%dresources_%drecords_%dlogs", numLogGroups, numRecords, numLogs), func(b *testing.B) {
				lc := &logsConsumer{
					unmarshaler: cwlog.NewUnmarshaler(zap.NewNop()),
					consumer:    consumertest.NewNop(),
				}
				records := make([][]byte, numRecords)
				for i := range records {
					records[i] = makeCloudWatchLogRecord(numLogs, numLogGroups)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					code, err := lc.Consume(context.Background(), records, nil)
					if err != nil {
						b.Fatal(err)
					}
					if code != http.StatusOK {
						b.Fatalf("expected status code 200, got %d", code)
					}
				}
			})
		}
	}
}

func BenchmarkMetricsConsumer_cwmetrics(b *testing.B) {
	// numStreams is the maximum number of unique metric streams
	// to use across the generated metrics, using a random generator.
	const numStreams = 10

	// numRecords is the number of records in the Firehose envelope.
	for _, numRecords := range []int{10, 100} {
		// numMetrics is the number of CoudWatch metrics within a Firehose record.
		for _, numMetrics := range []int{1, 10} {
			b.Run(fmt.Sprintf("%dresources_%drecords_%dmetrics", numStreams, numRecords, numMetrics), func(b *testing.B) {
				mc := &metricsConsumer{
					unmarshaler: cwmetricstream.NewUnmarshaler(zap.NewNop()),
					consumer:    consumertest.NewNop(),
				}
				records := make([][]byte, numRecords)
				for i := range records {
					records[i] = makeCloudWatchMetricRecord(numMetrics, numStreams)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					code, err := mc.Consume(context.Background(), records, nil)
					if err != nil {
						b.Fatal(err)
					}
					if code != http.StatusOK {
						b.Fatalf("expected status code 200, got %d", code)
					}
				}
			})
		}
	}
}

func makeCloudWatchLogRecord(numLogs, numLogGroups int) []byte {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	for i := 0; i < numLogs; i++ {
		group := rand.IntN(numLogGroups)
		fmt.Fprintf(w,
			`{"messageType":"DATA_MESSAGE","owner":"123","logGroup":"group_%d","logStream":"stream","logEvents":[{"id":"the_id","timestamp":1725594035523,"message":"message %d"}]}`,
			group, i,
		)
		fmt.Fprintln(w)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func makeCloudWatchMetricRecord(numMetrics, numStreams int) []byte {
	var buf bytes.Buffer
	for i := 0; i < numMetrics; i++ {
		stream := rand.IntN(numStreams)
		fmt.Fprintf(&buf,
			`{"metric_stream_name":"stream_%d","account_id":"1234567890","region":"us-east-1","namespace":"AWS/NATGateway","metric_name":"metric_%d","dimensions":{"NatGatewayId":"nat-01a4160dfb995b990"},"timestamp":1643916720000,"value":{"max":0.0,"min":0.0,"sum":0.0,"count":2.0},"unit":"Count"}`,
			stream, i,
		)
		fmt.Fprintln(&buf)
	}
	return buf.Bytes()
}
