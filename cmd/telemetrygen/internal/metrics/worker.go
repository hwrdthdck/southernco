// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric/global"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type worker struct {
	running        *atomic.Bool    // pointer to shared flag that indicates it's time to stop the test
	numMetrics     int             // how many metrics the worker has to generate (only when duration==0)
	totalDuration  time.Duration   // how long to run the test for (overrides `numMetrics`)
	limitPerSecond rate.Limit      // how many metrics per second to generate
	wg             *sync.WaitGroup // notify when done
	logger         *zap.Logger
}

func (w worker) simulateMetrics() {
	limiter := rate.NewLimiter(w.limitPerSecond, 1)
	var i int
	meter := global.Meter("telemetrygen")

	index := int64(0)
	max := int64(1000)
	if w.limitPerSecond != rate.Inf {
		max = int64(w.limitPerSecond)
	}

	counter, _ := meter.Int64Counter("gen")
	for w.running.Load() {
		counter.Add(context.Background(), 1)
		if err := limiter.Wait(context.Background()); err != nil {
			w.logger.Fatal("limiter waited failed, retry", zap.Error(err))
		}

		index++
		if index > max {
			index = 0
		}
		i++
		if w.numMetrics != 0 {
			if i >= w.numMetrics {
				break
			}
		}
	}

	w.logger.Info("metrics generated", zap.Int("metrics", i))
	w.wg.Done()
}
