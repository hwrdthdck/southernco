// Copyright 2020, OpenTelemetry Authors
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
	"errors"
	"strings"

	// TODO: don't use the opencensus-proto package???
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
)

// DogStatsDParser supports the Parse method for parsing StatsD messages with Tags.
type DogStatsDParser struct{}

// Parse returns an OTLP metric representation of the input StatsD string.
func (p *DogStatsDParser) Parse(line string) (*metricspb.Metric, error) {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return nil, errors.New("not enough statsd message parts")
	}

	return &metricspb.Metric{
		MetricDescriptor: &metricspb.MetricDescriptor{
			Name: parts[0],
		},
	}, nil
}
