// Copyright The OpenTelemetry Authors
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

package tracking

import (
	"context"
	"reflect"
	"testing"
	"time"

	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

func TestMetricTracker_Convert(t *testing.T) {
	miSum := MetricIdentity{
		Resource:               pdata.NewResource(),
		InstrumentationLibrary: pdata.NewInstrumentationLibrary(),
		MetricDataType:         pdata.MetricDataTypeSum,
		MetricIsMonotonic:      true,
		MetricName:             "",
		MetricUnit:             "",
		Attributes:             pdata.NewAttributeMap(),
	}
	miIntSum := miSum
	miIntSum.MetricValueType = pdata.MetricValueTypeInt
	miSum.MetricValueType = pdata.MetricValueTypeDouble

	m := NewMetricTracker(context.Background(), zap.NewNop(), 0)

	tests := []struct {
		name    string
		value   ValuePoint
		wantOut DeltaValue
	}{
		{
			name: "Initial Value recorded",
			value: ValuePoint{
				ObservedTimestamp: 10,
				FloatValue:        100.0,
				IntValue:          100,
			},
			wantOut: DeltaValue{
				StartTimestamp: 10,
				FloatValue:     100.0,
				IntValue:       100,
			},
		},
		{
			name: "Higher Value Recorded",
			value: ValuePoint{
				ObservedTimestamp: 50,
				FloatValue:        225.0,
				IntValue:          225,
			},
			wantOut: DeltaValue{
				StartTimestamp: 10,
				FloatValue:     125.0,
				IntValue:       125,
			},
		},
		{
			name: "Lower Value Recorded - No Previous Offset",
			value: ValuePoint{
				ObservedTimestamp: 100,
				FloatValue:        75.0,
				IntValue:          75,
			},
			wantOut: DeltaValue{
				StartTimestamp: 50,
				FloatValue:     75.0,
				IntValue:       75,
			},
		},
		{
			name: "Record delta above first recorded value",
			value: ValuePoint{
				ObservedTimestamp: 150,
				FloatValue:        300.0,
				IntValue:          300,
			},
			wantOut: DeltaValue{
				StartTimestamp: 100,
				FloatValue:     225.0,
				IntValue:       225,
			},
		},
		{
			name: "Lower Value Recorded - Previous Offset Recorded",
			value: ValuePoint{
				ObservedTimestamp: 200,
				FloatValue:        25.0,
				IntValue:          25,
			},
			wantOut: DeltaValue{
				StartTimestamp: 150,
				FloatValue:     25.0,
				IntValue:       25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floatPoint := MetricPoint{
				Identity: miSum,
				Value:    tt.value,
			}
			intPoint := MetricPoint{
				Identity: miIntSum,
				Value:    tt.value,
			}

			if gotOut, valid := m.Convert(floatPoint); !valid || !reflect.DeepEqual(gotOut.StartTimestamp, tt.wantOut.StartTimestamp) || !reflect.DeepEqual(gotOut.FloatValue, tt.wantOut.FloatValue) {
				t.Errorf("MetricTracker.Convert(MetricDataTypeSum) = %v, want %v", gotOut, tt.wantOut)
			}

			if gotOut, valid := m.Convert(intPoint); !valid || !reflect.DeepEqual(gotOut.StartTimestamp, tt.wantOut.StartTimestamp) || !reflect.DeepEqual(gotOut.IntValue, tt.wantOut.IntValue) {
				t.Errorf("MetricTracker.Convert(MetricDataTypeIntSum) = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}

	t.Run("Invalid metric identity", func(t *testing.T) {
		invalidID := miIntSum
		invalidID.MetricDataType = pdata.MetricDataTypeGauge
		_, valid := m.Convert(MetricPoint{
			Identity: invalidID,
			Value: ValuePoint{
				ObservedTimestamp: 0,
				FloatValue:        100.0,
				IntValue:          100,
			},
		})
		if valid {
			t.Error("Expected invalid for non cumulative metric")
		}
	})
}

func Test_metricTracker_removeStale(t *testing.T) {
	currentTime := pdata.Timestamp(100)
	freshPoint := ValuePoint{
		ObservedTimestamp: currentTime,
	}
	stalePoint := ValuePoint{
		ObservedTimestamp: currentTime - 1,
	}

	type fields struct {
		MaxStale time.Duration
		States   map[string]*State
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut map[string]*State
	}{
		{
			name: "Removes stale entry, leaves fresh entry",
			fields: fields{
				MaxStale: 0, // This logic isn't tested here
				States: map[string]*State{
					"stale": {
						PrevPoint: stalePoint,
					},
					"fresh": {
						PrevPoint: freshPoint,
					},
				},
			},
			wantOut: map[string]*State{
				"fresh": {
					PrevPoint: freshPoint,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &metricTracker{
				logger:   zap.NewNop(),
				maxStale: tt.fields.MaxStale,
			}
			for k, v := range tt.fields.States {
				tr.states.Store(k, v)
			}
			tr.removeStale(currentTime)

			gotOut := make(map[string]*State)
			tr.states.Range(func(key, value interface{}) bool {
				gotOut[key.(string)] = value.(*State)
				return true
			})

			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("MetricTracker.removeStale() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_metricTracker_sweeper(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sweepEvent := make(chan pdata.Timestamp)
	closed := false

	onSweep := func(staleBefore pdata.Timestamp) {
		sweepEvent <- staleBefore
	}

	tr := &metricTracker{
		logger:   zap.NewNop(),
		maxStale: 1 * time.Millisecond,
	}

	start := time.Now()
	go func() {
		tr.sweeper(ctx, onSweep)
		closed = true
		close(sweepEvent)
	}()

	for i := 1; i <= 2; i++ {
		staleBefore := <-sweepEvent
		tickTime := time.Since(start) + tr.maxStale*time.Duration(i)
		if closed {
			t.Fatalf("Sweeper returned prematurely.")
		}

		if tickTime < tr.maxStale {
			t.Errorf("Sweeper tick time is too fast. (%v, want %v)", tickTime, tr.maxStale)
		}

		staleTime := staleBefore.AsTime()
		if time.Since(staleTime) < tr.maxStale {
			t.Errorf("Sweeper called with invalid staleBefore value = %v", staleTime)
		}
	}
	cancel()
	<-sweepEvent
	if !closed {
		t.Errorf("Sweeper did not terminate.")
	}
}
