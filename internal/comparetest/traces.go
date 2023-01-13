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

package comparetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/comparetest"

import (
	"fmt"
	"reflect"

	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/multierr"
)

// CompareTraces compares each part of two given Traces and returns
// an error if they don't match. The error describes what didn't match.
func CompareTraces(expected, actual ptrace.Traces, options ...TracesCompareOption) error {
	exp, act := ptrace.NewTraces(), ptrace.NewTraces()
	expected.CopyTo(exp)
	actual.CopyTo(act)

	for _, option := range options {
		option.applyOnTraces(exp, act)
	}

	expectedSpans, actualSpans := exp.ResourceSpans(), act.ResourceSpans()
	if expectedSpans.Len() != actualSpans.Len() {
		return fmt.Errorf("amount of ResourceSpans between Traces are not equal expected: %d, actual: %d",
			expectedSpans.Len(),
			actualSpans.Len())
	}

	// sort ResourceSpans
	expectedSpans.Sort(sortResourceSpans)
	actualSpans.Sort(sortResourceSpans)

	numResources := expectedSpans.Len()

	// Keep track of matching resources so that each can only be matched once
	matchingResources := make(map[ptrace.ResourceSpans]ptrace.ResourceSpans, numResources)

	var errs error
	for e := 0; e < numResources; e++ {
		er := expectedSpans.At(e)
		var foundMatch bool
		for a := 0; a < numResources; a++ {
			ar := actualSpans.At(a)
			if _, ok := matchingResources[ar]; ok {
				continue
			}
			if reflect.DeepEqual(er.Resource().Attributes().AsRaw(), ar.Resource().Attributes().AsRaw()) {
				foundMatch = true
				matchingResources[ar] = er
				break
			}
		}
		if !foundMatch {
			errs = multierr.Append(errs, fmt.Errorf("missing expected resource with attributes: %v", er.Resource().Attributes().AsRaw()))
		}
	}

	for i := 0; i < numResources; i++ {
		if _, ok := matchingResources[actualSpans.At(i)]; !ok {
			errs = multierr.Append(errs, fmt.Errorf("extra resource with attributes: %v", actualSpans.At(i).Resource().Attributes().AsRaw()))
		}
	}

	if errs != nil {
		return errs
	}

	for ar, er := range matchingResources {
		if err := CompareResourceSpans(er, ar); err != nil {
			return err
		}
	}

	return nil
}

// CompareResourceSpans compares each part of two given ResourceSpans and returns
// an error if they don't match. The error describes what didn't match.
func CompareResourceSpans(expected, actual ptrace.ResourceSpans) error {
	exp, act := ptrace.NewResourceSpans(), ptrace.NewResourceSpans()
	expected.CopyTo(exp)
	actual.CopyTo(act)

	eilms := exp.ScopeSpans()
	ailms := act.ScopeSpans()

	if eilms.Len() != ailms.Len() {
		return fmt.Errorf("number of instrumentation libraries does not match expected: %d, actual: %d", eilms.Len(),
			ailms.Len())
	}

	// sort InstrumentationLibrary
	eilms.Sort(sortSpansInstrumentationLibrary)
	ailms.Sort(sortSpansInstrumentationLibrary)

	for i := 0; i < eilms.Len(); i++ {
		eilm, ailm := eilms.At(i), ailms.At(i)
		eil, ail := eilm.Scope(), ailm.Scope()

		if eil.Name() != ail.Name() {
			return fmt.Errorf("instrumentation library Name does not match expected: %s, actual: %s", eil.Name(), ail.Name())
		}
		if eil.Version() != ail.Version() {
			return fmt.Errorf("instrumentation library Version does not match expected: %s, actual: %s", eil.Version(), ail.Version())
		}
		if err := CompareSpanSlices(eilm.Spans(), ailm.Spans()); err != nil {
			return err
		}
	}
	return nil
}

// CompareSpanSlices compares each part of two given SpanSlices and returns
// an error if they don't match. The error describes what didn't match.
func CompareSpanSlices(expected, actual ptrace.SpanSlice) error {
	exp, act := ptrace.NewSpanSlice(), ptrace.NewSpanSlice()
	expected.CopyTo(exp)
	actual.CopyTo(act)

	if exp.Len() != act.Len() {
		return fmt.Errorf("number of spans does not match expected: %d, actual: %d", exp.Len(), act.Len())
	}

	exp.Sort(sortSpanSlice)
	act.Sort(sortSpanSlice)

	numSpans := exp.Len()

	// Keep track of matching spans so that each span can only be matched once
	matchingSpans := make(map[ptrace.Span]ptrace.Span, numSpans)

	var errs error
	for e := 0; e < numSpans; e++ {
		elr := exp.At(e)
		var foundMatch bool
		for a := 0; a < numSpans; a++ {
			alr := act.At(a)
			if _, ok := matchingSpans[alr]; ok {
				continue
			}
			if reflect.DeepEqual(elr.Attributes().AsRaw(), alr.Attributes().AsRaw()) {
				foundMatch = true
				matchingSpans[alr] = elr
				break
			}
		}
		if !foundMatch {
			errs = multierr.Append(errs, fmt.Errorf("span missing expected resource with attributes: %v", elr.Attributes().AsRaw()))
		}
	}

	for i := 0; i < numSpans; i++ {
		if _, ok := matchingSpans[act.At(i)]; !ok {
			errs = multierr.Append(errs, fmt.Errorf("span has extra record with attributes: %v", act.At(i).Attributes().AsRaw()))
		}
	}

	if errs != nil {
		return errs
	}

	for alr, elr := range matchingSpans {
		if err := CompareSpans(alr, elr); err != nil {
			return multierr.Combine(fmt.Errorf("span with attributes: %v, does not match expected %v", alr.Attributes().AsRaw(), elr.Attributes().AsRaw()), err)
		}
	}
	return nil
}

// CompareSpans compares each part of two given Span and returns
// an error if they don't match. The error describes what didn't match.
func CompareSpans(expected, actual ptrace.Span) error {
	exp, act := ptrace.NewSpan(), ptrace.NewSpan()
	expected.CopyTo(exp)
	actual.CopyTo(act)

	if exp.TraceID() != act.TraceID() {
		return fmt.Errorf("span TraceID doesn't match expected: %d, actual: %d",
			exp.TraceID(),
			act.TraceID())
	}

	if exp.SpanID() != act.SpanID() {
		return fmt.Errorf("span SpanID doesn't match expected: %d, actual: %d",
			exp.SpanID(),
			act.SpanID())
	}

	if exp.TraceState().AsRaw() != act.TraceState().AsRaw() {
		return fmt.Errorf("span TraceState doesn't match expected: %s, actual: %s",
			exp.TraceState().AsRaw(),
			act.TraceState().AsRaw())
	}

	if exp.ParentSpanID() != act.ParentSpanID() {
		return fmt.Errorf("span ParentSpanID doesn't match expected: %d, actual: %d",
			exp.ParentSpanID(),
			act.ParentSpanID())
	}

	if exp.Name() != act.Name() {
		return fmt.Errorf("span Name doesn't match expected: %s, actual: %s",
			exp.Name(),
			act.Name())
	}

	if exp.Kind() != act.Kind() {
		return fmt.Errorf("span Kind doesn't match expected: %d, actual: %d",
			exp.Kind(),
			act.Kind())
	}

	if exp.StartTimestamp() != act.StartTimestamp() {
		return fmt.Errorf("span StartTimestamp doesn't match expected: %d, actual: %d",
			exp.StartTimestamp(),
			act.StartTimestamp())
	}

	if exp.EndTimestamp() != act.EndTimestamp() {
		return fmt.Errorf("span EndTimestamp doesn't match expected: %d, actual: %d",
			exp.EndTimestamp(),
			act.EndTimestamp())
	}

	if !reflect.DeepEqual(exp.Attributes().AsRaw(), act.Attributes().AsRaw()) {
		return fmt.Errorf("span Attributes doesn't match expected: %s, actual: %s",
			exp.Attributes().AsRaw(),
			act.Attributes().AsRaw())
	}

	if exp.DroppedAttributesCount() != act.DroppedAttributesCount() {
		return fmt.Errorf("span DroppedAttributesCount doesn't match expected: %d, actual: %d",
			exp.DroppedAttributesCount(),
			act.DroppedAttributesCount())
	}

	if !reflect.DeepEqual(exp.Events(), act.Events()) {
		return fmt.Errorf("span Events doesn't match expected: %v, actual: %v",
			exp.Events(),
			act.Events())
	}

	if exp.DroppedEventsCount() != act.DroppedEventsCount() {
		return fmt.Errorf("span DroppedEventsCount doesn't match expected: %d, actual: %d",
			exp.DroppedEventsCount(),
			act.DroppedEventsCount())
	}

	if !reflect.DeepEqual(exp.Links(), act.Links()) {
		return fmt.Errorf("span Links doesn't match expected: %v, actual: %v",
			exp.Links(),
			act.Links())
	}

	if exp.DroppedLinksCount() != act.DroppedLinksCount() {
		return fmt.Errorf("span DroppedLinksCount doesn't match expected: %d, actual: %d",
			exp.DroppedLinksCount(),
			actual.DroppedLinksCount())
	}

	if !reflect.DeepEqual(exp.Status(), act.Status()) {
		return fmt.Errorf("span Status doesn't match expected: %v, actual: %v",
			exp.Status(),
			act.Status())
	}

	return nil
}
