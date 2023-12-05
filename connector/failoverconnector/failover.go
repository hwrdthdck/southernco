// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector/internal/state"
)

type consumerProvider[C any] func(...component.ID) (C, error)

type failoverRouter[C any] struct {
	consumerProvider consumerProvider[C]
	cfg              *Config
	pS               *state.PipelineSelector
	consumers        []C
	rS               *state.RetryState
}

var (
	errNoValidPipeline = errors.New("All provided pipelines return errors")
	errConsumer        = errors.New("Error registering consumer")
)

func newFailoverRouter[C any](provider consumerProvider[C], cfg *Config) *failoverRouter[C] {
	return &failoverRouter[C]{
		consumerProvider: provider,
		cfg:              cfg,
		pS:               state.NewPipelineSelector(len(cfg.PipelinePriority), cfg.MaxRetries),
		rS:               &state.RetryState{},
	}
}

func (f *failoverRouter[C]) getCurrentConsumer() (C, int, bool) {
	// if currentIndex incremented passed bounds of pipeline list
	var nilConsumer C
	if f.pS.CurrentIndex >= len(f.cfg.PipelinePriority) {
		return nilConsumer, -1, false
	}
	idx := f.pS.GetCurrentIndex()
	return f.consumers[idx], idx, true
}

func (f *failoverRouter[C]) registerConsumers() error {
	consumers := make([]C, 0)
	for _, pipelines := range f.cfg.PipelinePriority {
		newConsumer, err := f.consumerProvider(pipelines...)
		if err != nil {
			return errConsumer
		}
		consumers = append(consumers, newConsumer)
	}
	f.consumers = consumers
	return nil
}

func (f *failoverRouter[C]) handlePipelineError(idx int) {
	// avoids race condition in case of consumeSIGNAL invocations
	// where index was updated during execution
	if idx != f.pS.GetCurrentIndex() {
		return
	}
	doRetry := f.pS.IndexIsStable(idx)
	// UpdatePipelineIndex either increments the pipeline to the next priority
	// or returns it to the stable
	f.pS.UpdatePipelineIndex(idx)
	// if the currentIndex is not the stableIndex, that means the currentIndex is a higher
	// priority index that was set during a retry, in which case we don't want to start a
	// new retry goroutine
	if !doRetry {
		return
	}
	// kill existing retry goroutine if error is from a stable pipeline that failed for the first time
	ctx, cancel := context.WithCancel(context.Background())
	f.rS.InvokeCancel()
	f.rS.UpdateCancelFunc(cancel)
	f.enableRetry(ctx)
}

func (f *failoverRouter[C]) enableRetry(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(f.cfg.RetryInterval)
		defer ticker.Stop()

		var cancelFunc context.CancelFunc
		// checkContinueRetry checks that any higher priority levels have retries remaining
		// (have not exceeded their maxRetries)
		for f.checkContinueRetry(f.pS.StableIndex) {
			select {
			case <-ticker.C:
				// When the nextRetry interval starts we kill the existing iteration through
				// the higher priority pipelines if still in progress
				if cancelFunc != nil {
					cancelFunc()
				}
				cancelFunc = f.handleRetry(ctx, f.pS.StableIndex)
			case <-ctx.Done():
				return
			}
		}
	}()
}

// handleRetry is responsible for launching goroutine and returning cancelFunc for context to be called if new
// interval starts in the middle of the execution
func (f *failoverRouter[C]) handleRetry(parentCtx context.Context, stableIndex int) context.CancelFunc {
	retryCtx, cancelFunc := context.WithCancel(parentCtx)
	go f.retryHighPriorityPipelines(retryCtx, stableIndex)
	return cancelFunc
}

// retryHighPriorityPipelines responsible for single iteration through all higher priority pipelines
func (f *failoverRouter[C]) retryHighPriorityPipelines(ctx context.Context, stableIndex int) {
	ticker := time.NewTicker(f.cfg.RetryGap)

	defer ticker.Stop()

	for i := 0; i < stableIndex; i++ {
		// if stableIndex was updated to a higher priority level during the execution of the goroutine
		// will return to avoid overwriting higher priority level with lower one
		if stableIndex > f.pS.GetStableIndex() {
			return
		}
		// checks that max retries were not used for this index
		if f.pS.MaxRetriesUsed(i) {
			continue
		}
		select {
		// return when context is cancelled by parent goroutine
		case <-ctx.Done():
			return
		case <-ticker.C:
			// when ticker triggers currentIndex is updated
			f.pS.SetToRetryIndex(i)
		}
	}
}

// checkStopRetry checks if retry should be suspended if all higher priority levels have exceeded their max retries
func (f *failoverRouter[C]) checkContinueRetry(index int) bool {
	for i := 0; i < index; i++ {
		if f.pS.PipelineRetries[i] < f.cfg.MaxRetries {
			return true
		}
	}
	return false
}

// reportStable reports back to the failoverRouter that the current priority level that was called by Consume.SIGNAL was
// stable
func (f *failoverRouter[C]) reportStable(idx int) {
	// is stableIndex is already the known stableIndex return
	if f.pS.IndexIsStable(idx) {
		return
	}
	// if the stableIndex is a retried index, the update the stable index to the retried index
	// NOTE retry will not stop due to potential higher priority index still available
	f.pS.SetNewStableIndex(idx)
}
