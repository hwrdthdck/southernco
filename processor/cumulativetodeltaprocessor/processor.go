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

package cumulativetodeltaprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

type cumulativeToDeltaProcessor struct {
	metrics []string
	logger  *zap.Logger
}

func newCumulativeToDeltaProcessor(config *Config, logger *zap.Logger) *cumulativeToDeltaProcessor {
	return &cumulativeToDeltaProcessor{
		metrics: config.Metrics,
		logger:  logger,
	}
}

// Start is invoked during service startup.
func (mgp *cumulativeToDeltaProcessor) Start(context.Context, component.Host) error {
	return nil
}

// processMetrics implements the ProcessMetricsFunc type.
func (mgp *cumulativeToDeltaProcessor) processMetrics(_ context.Context, md pdata.Metrics) (pdata.Metrics, error) {
	return md, nil
}

// Shutdown is invoked during service shutdown.
func (mgp *cumulativeToDeltaProcessor) Shutdown(context.Context) error {
	return nil
}
