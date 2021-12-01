// Copyright The OpenTelemetry Authors
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

package stanza // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/stanza"

import (
	"context"
	"sync"
	"time"

	"github.com/open-telemetry/opentelemetry-log-collection/entry"
	"github.com/open-telemetry/opentelemetry-log-collection/operator"
	"github.com/open-telemetry/opentelemetry-log-collection/operator/helper"
	"go.uber.org/zap"
)

// LogEmitter is a stanza operator that emits log entries to a channel
type LogEmitter struct {
	helper.OutputOperator
	logChan       chan []*entry.Entry
	stopOnce      sync.Once
	cancel        context.CancelFunc
	batchMux      sync.Mutex
	batch         []*entry.Entry
	wg            sync.WaitGroup
	maxBatchSize  uint
	flushInterval time.Duration
}

type LogEmitterOption func(*LogEmitter)

var (
	defaultFlushInterval      = 100 * time.Millisecond
	defaultMaxBatchSize  uint = 100
)

// LogEmitterWithMaxBatchSize returns an option that makes the LogEmitter use the specified max batch size
func LogEmitterWithMaxBatchSize(maxBatchSize uint) LogEmitterOption {
	return LogEmitterOption(func(le *LogEmitter) {
		le.maxBatchSize = maxBatchSize
		le.batch = make([]*entry.Entry, 0, maxBatchSize)
	})
}

// LogEmitterWithFlushInterval returns an option that makes the LogEmitter use the specified flush interval
func LogEmitterWithFlushInterval(flushInterval time.Duration) LogEmitterOption {
	return LogEmitterOption(func(le *LogEmitter) {
		le.flushInterval = flushInterval
	})
}

// LogEmitterWithLogger returns an option that makes the LogEmitter use the specified logger
func LogEmitterWithLogger(logger *zap.SugaredLogger) LogEmitterOption {
	return LogEmitterOption(func(le *LogEmitter) {
		le.OutputOperator.BasicOperator.SugaredLogger = logger
	})
}

// NewLogEmitter creates a new receiver output
func NewLogEmitter(opts ...LogEmitterOption) *LogEmitter {
	le := &LogEmitter{
		OutputOperator: helper.OutputOperator{
			BasicOperator: helper.BasicOperator{
				OperatorID:    "log_emitter",
				OperatorType:  "log_emitter",
				SugaredLogger: zap.NewNop().Sugar(),
			},
		},
		logChan:       make(chan []*entry.Entry),
		maxBatchSize:  defaultMaxBatchSize,
		batch:         make([]*entry.Entry, 0, defaultMaxBatchSize),
		flushInterval: defaultFlushInterval,
		cancel:        func() {},
	}

	for _, opt := range opts {
		opt(le)
	}

	return le
}

// Start starts the goroutine(s) required for this operator
func (e *LogEmitter) Start(_ operator.Persister) error {
	ctx, cancel := context.WithCancel(context.Background())
	e.cancel = cancel

	e.wg.Add(1)
	go e.flusher(ctx)
	return nil
}

// Stop will close the log channel and stop running goroutines
func (e *LogEmitter) Stop() error {
	e.stopOnce.Do(func() {
		close(e.logChan)
		e.cancel()

		e.wg.Wait()
	})

	return nil
}

// Process will emit an entry to the output channel
func (e *LogEmitter) Process(ctx context.Context, ent *entry.Entry) error {
	batchToFlush := e.appendEntry(ent)

	if batchToFlush != nil {
		e.flush(ctx, batchToFlush)
	}

	return nil
}

// appendEntry appends the entry to the current batch. If maxBatchSize is reached, a new batch will be made, and the old batch
// (which should be flushed) will be returned
func (e *LogEmitter) appendEntry(ent *entry.Entry) []*entry.Entry {
	e.batchMux.Lock()
	defer e.batchMux.Unlock()

	e.batch = append(e.batch, ent)
	if uint(len(e.batch)) >= e.maxBatchSize {
		// maxBatchSize has been exceeded; create a new batch, returning the old one.
		return e.makeNewBatchNoLock()
	}

	return nil
}

// flusher flushes the current batch every flush interval. Intended to be run as a goroutine
func (e *LogEmitter) flusher(ctx context.Context) {
	defer e.wg.Done()

	ticker := time.NewTicker(e.flushInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			batch := e.makeNewBatch()
			if batch != nil {
				e.flush(ctx, batch)
			}
		case <-ctx.Done():
			return
		}
	}
}

// flush flushes the provided batch to the log channel.
func (e *LogEmitter) flush(ctx context.Context, batch []*entry.Entry) {
	select {
	case e.logChan <- batch:
	case <-ctx.Done():
	}
}

// makeNewBatch replaces the current batch on the log emitter with a new batch, returning the old one
func (e *LogEmitter) makeNewBatch() []*entry.Entry {
	e.batchMux.Lock()
	defer e.batchMux.Unlock()

	if len(e.batch) == 0 {
		return nil
	}

	return e.makeNewBatchNoLock()
}

// makeNewBatchNoLock replaces the current batch on the log emitter with a new batch, returning the old one. It does not acquire the batchMux,
// so it should used in cases where the batchMux is already acquired
func (e *LogEmitter) makeNewBatchNoLock() []*entry.Entry {
	oldBatch := e.batch
	e.batch = make([]*entry.Entry, 0, e.maxBatchSize)

	return oldBatch
}
