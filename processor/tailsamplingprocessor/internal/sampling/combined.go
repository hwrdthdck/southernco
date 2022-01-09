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

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

type Combined struct {
	// the subpolicy evaluators
	subpolicies []PolicyEvaluator
	logger      *zap.Logger
}

func NewCombined(
	logger *zap.Logger,
	subpolicies []PolicyEvaluator,
) PolicyEvaluator {

	return &Combined{
		subpolicies: subpolicies,
		logger:      logger,
	}
}

func (c *Combined) Evaluate(traceID pdata.TraceID, trace *TraceData) (Decision, error) {
	for _, sub := range c.subpolicies {
		decision, err := sub.Evaluate(traceID, trace)
		if err != nil {
			return Unspecified, err
		}
		if decision == NotSampled {
			return NotSampled, nil
		}

	}
	return Sampled, nil
}

func (c *Combined) OnLateArrivingSpans(Decision, []*pdata.Span) error {
	c.logger.Debug("Spans are arriving late, decision is already made!!!")
	return nil
}

func (c *Combined) OnDroppedSpans(pdata.TraceID, *TraceData) (Decision, error) {
	return Sampled, nil
}
