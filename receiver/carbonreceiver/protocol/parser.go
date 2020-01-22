// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Parser abstracts the type of parsing being done by the receiver.
type Parser interface {
	// Parse receives the string with plaintext data, aka line, in the Carbon
	// format and transforms it to the collector metric format.
	Parse(line string) (*metricspb.Metric, error)
}

// Below a few helper functions useful to different parsers.
func buildMetricForSinglePoint(
	metricName string,
	metricType metricspb.MetricDescriptor_Type,
	labelKeys []*metricspb.LabelKey,
	labelValues []*metricspb.LabelValue,
	point *metricspb.Point,
) *metricspb.Metric {
	return &metricspb.Metric{
		MetricDescriptor: &metricspb.MetricDescriptor{
			Name:      metricName,
			Type:      metricType,
			LabelKeys: labelKeys,
		},
		Timeseries: []*metricspb.TimeSeries{
			{
				// TODO: StartTimestamp can be set if each cumulative time series are
				//  	tracked but right now it is not clear if it brings benefits.
				//		Perhaps as an option so cost is "pay for play".
				LabelValues: labelValues,
				Points:      []*metricspb.Point{point},
			},
		},
	}
}

func convertUnixSec(sec int64) *timestamp.Timestamp {
	if sec == 0 {
		return nil
	}

	ts := &timestamp.Timestamp{
		Seconds: sec,
	}
	return ts
}
