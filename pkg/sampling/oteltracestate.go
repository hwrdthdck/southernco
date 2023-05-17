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

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"

import (
	"fmt"
	"strings"
)

type otelTraceState struct {
	tvalueString string
	tvalueParsed Threshold
	baseTraceState
}

type otelTraceStateParser struct{}

func (wp otelTraceStateParser) parseField(concrete *otelTraceState, key, input string) error {
	switch {
	case key == "t":
		value, err := stripKey(key, input)
		if err != nil {
			return err
		}

		prob, _, err := TvalueToProbabilityAndAdjustedCount(value)
		if err != nil {
			return fmt.Errorf("otel tracestate t-value: %w", err)
		}

		th, err := ProbabilityToThreshold(prob)
		if err != nil {
			return fmt.Errorf("otel tracestate t-value: %w", err)
		}

		concrete.tvalueString = input
		concrete.tvalueParsed = th

		return nil
	}

	return baseTraceStateParser{}.parseField(&concrete.baseTraceState, key, input)
}

func (otts otelTraceState) serialize() string {
	var sb strings.Builder

	if otts.hasTValue() {
		_, _ = sb.WriteString(otts.tvalueString)
	}

	otelSyntax.serialize(&otts.baseTraceState, &sb)

	return sb.String()
}

func (otts otelTraceState) hasTValue() bool {
	return otts.tvalueString != ""
}
