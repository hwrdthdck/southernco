// Copyright  The OpenTelemetry Authors
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

package sampling

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

func TestNewPercentageFilter_errorHandling(t *testing.T) {
	_, err := NewPercentageFilter(zap.NewNop(), -1)
	assert.EqualError(t, err, "expected a percentage between 0 and 1")

	_, err = NewPercentageFilter(zap.NewNop(), 1.5)
	assert.EqualError(t, err, "expected a percentage between 0 and 1")
}

func TestPercentageSampling(t *testing.T) {
	var empty = map[string]pdata.AttributeValue{}

	cases := []float32{0.01, 0.1, 0.33, 0.5, 0.66, 1}

	for _, percentage := range cases {
		t.Run(fmt.Sprintf("sample %.2f", percentage), func(t *testing.T) {
			trace := newTraceStringAttrs(empty, "example", "value")
			traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

			percentageFilter, err := NewPercentageFilter(zap.NewNop(), percentage)
			assert.NoError(t, err)

			sampled := 0.

			for i := 0; i < 100; i++ {
				decision, err := percentageFilter.Evaluate(traceID, trace)
				assert.NoError(t, err)

				if decision == Sampled {
					sampled++
				}
			}

			assert.InDelta(t, 100*percentage, sampled, 0.001, "Sampled traces")
			assert.InDelta(t, 100*(1-percentage), 100-sampled, 0.001, "Not sampled traces")
		})
	}
}

func TestPercentageSampling_ignoreAlreadySampledTraces(t *testing.T) {
	var empty = map[string]pdata.AttributeValue{}

	trace := newTraceStringAttrs(empty, "example", "value")
	traceID := pdata.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	percentageFilter, err := NewPercentageFilter(zap.NewNop(), 0.5)
	assert.NoError(t, err)

	// every two traces should get sampled, already sampled traces don't count
	for i := 0; i < 100; i++ {
		trace.Decisions = []Decision{NotSampled, NotSampled}
		decision, err := percentageFilter.Evaluate(traceID, trace)
		assert.NoError(t, err)
		assert.Equal(t, decision, Sampled)

		trace.Decisions = []Decision{NotSampled, NotSampled}
		decision, err = percentageFilter.Evaluate(traceID, trace)
		assert.NoError(t, err)
		assert.Equal(t, decision, NotSampled)

		// trace has been sampled, should be ignored
		trace.Decisions = []Decision{NotSampled, Sampled}
		decision, err = percentageFilter.Evaluate(traceID, trace)
		assert.NoError(t, err)
		assert.Equal(t, decision, NotSampled)
	}
}

func TestOnLateArrivingSpans_PercentageSampling(t *testing.T) {
	percentageFilter, err := NewPercentageFilter(zap.NewNop(), 0.1)
	assert.Nil(t, err)

	err = percentageFilter.OnLateArrivingSpans(NotSampled, nil)
	assert.Nil(t, err)
}
