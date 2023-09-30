// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"errors"
	"sort"

	"github.com/prometheus/prometheus/prompb"
)

// batchTimeSeries splits series into multiple batch write requests.
func batchTimeSeries(tsMap map[string]*prompb.TimeSeries, maxBatchByteSize int, m []prompb.MetricMetadata) ([]*prompb.WriteRequest, error) {
	if len(tsMap) == 0 {
		return nil, errors.New("invalid tsMap: cannot be empty map")
	}

	requests := make([]*prompb.WriteRequest, 0, len(tsMap))
	tsArray := make([]prompb.TimeSeries, 0, len(tsMap))
	sizeOfCurrentBatch := 0

	sizeOfM := 0
	for _, mi := range m {
		sizeOfM += mi.Size()
	}

	i := 0
	for _, v := range tsMap {
		sizeOfSeries := v.Size()

		if sizeOfCurrentBatch+sizeOfSeries+sizeOfM >= maxBatchByteSize {
			wrapped := convertTimeseriesToRequest(tsArray, m)
			requests = append(requests, wrapped)

			tsArray = make([]prompb.TimeSeries, 0, len(tsMap)-i)
			sizeOfCurrentBatch = 0
		}

		tsArray = append(tsArray, *v)
		sizeOfCurrentBatch += sizeOfSeries
		i++
	}

	if len(tsArray) != 0 {
		wrapped := convertTimeseriesToRequest(tsArray, m)
		requests = append(requests, wrapped)
	}

	return requests, nil
}

func convertTimeseriesToRequest(tsArray []prompb.TimeSeries, m []prompb.MetricMetadata) *prompb.WriteRequest {
	// the remote_write endpoint only requires the timeseries.
	// otlp defines it's own way to handle metric metadata
	return &prompb.WriteRequest{
		// Prometheus requires time series to be sorted by Timestamp to avoid out of order problems.
		// See:
		// * https://github.com/open-telemetry/wg-prometheus/issues/10
		// * https://github.com/open-telemetry/opentelemetry-collector/issues/2315
		Timeseries: orderBySampleTimestamp(tsArray),
		Metadata:   m,
	}
}

func orderBySampleTimestamp(tsArray []prompb.TimeSeries) []prompb.TimeSeries {
	for i := range tsArray {
		sL := tsArray[i].Samples
		sort.Slice(sL, func(i, j int) bool {
			return sL[i].Timestamp < sL[j].Timestamp
		})
	}
	return tsArray
}
