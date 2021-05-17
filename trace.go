// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testdata

import (
	"time"

	"go.opentelemetry.io/collector/consumer/pdata"
	otlpcollectortrace "go.opentelemetry.io/collector/internal/data/protogen/collector/trace/v1"
	otlptrace "go.opentelemetry.io/collector/internal/data/protogen/trace/v1"
)

var (
	TestSpanStartTime      = time.Date(2020, 2, 11, 20, 26, 12, 321, time.UTC)
	TestSpanStartTimestamp = pdata.TimestampFromTime(TestSpanStartTime)

	TestSpanEventTime      = time.Date(2020, 2, 11, 20, 26, 13, 123, time.UTC)
	TestSpanEventTimestamp = pdata.TimestampFromTime(TestSpanEventTime)

	TestSpanEndTime      = time.Date(2020, 2, 11, 20, 26, 13, 789, time.UTC)
	TestSpanEndTimestamp = pdata.TimestampFromTime(TestSpanEndTime)
)

func GenerateTraceDataOneEmptyResourceSpans() pdata.Traces {
	td := pdata.NewTraces()
	td.ResourceSpans().AppendEmpty()
	return td
}

func generateTraceOtlpOneEmptyResourceSpans() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{},
		},
	}
}

func GenerateTraceDataNoLibraries() pdata.Traces {
	td := GenerateTraceDataOneEmptyResourceSpans()
	rs0 := td.ResourceSpans().At(0)
	initResource1(rs0.Resource())
	return td
}

func generateTraceOtlpNoLibraries() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource: generateOtlpResource1(),
			},
		},
	}
}

func GenerateTraceDataOneEmptyInstrumentationLibrary() pdata.Traces {
	td := GenerateTraceDataNoLibraries()
	td.ResourceSpans().At(0).InstrumentationLibrarySpans().AppendEmpty()
	return td
}

func generateTraceOtlpOneEmptyInstrumentationLibrary() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource: generateOtlpResource1(),
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{},
				},
			},
		},
	}
}

func GenerateTraceDataOneSpanNoResource() pdata.Traces {
	td := GenerateTraceDataOneEmptyResourceSpans()
	rs0 := td.ResourceSpans().At(0)
	fillSpanOne(rs0.InstrumentationLibrarySpans().AppendEmpty().Spans().AppendEmpty())
	return td
}

func generateTraceOtlpOneSpanNoResource() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							generateOtlpSpanOne(),
						},
					},
				},
			},
		},
	}
}

func GenerateTraceDataOneSpan() pdata.Traces {
	td := GenerateTraceDataOneEmptyInstrumentationLibrary()
	rs0ils0 := td.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0)
	fillSpanOne(rs0ils0.Spans().AppendEmpty())
	return td
}

func generateTraceOtlpOneSpan() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource: generateOtlpResource1(),
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							generateOtlpSpanOne(),
						},
					},
				},
			},
		},
	}
}

func GenerateTraceDataTwoSpansSameResource() pdata.Traces {
	td := GenerateTraceDataOneEmptyInstrumentationLibrary()
	rs0ils0 := td.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0)
	fillSpanOne(rs0ils0.Spans().AppendEmpty())
	fillSpanTwo(rs0ils0.Spans().AppendEmpty())
	return td
}

// generateTraceOtlpSameResourceTwoSpans returns the OTLP representation of the generateTraceOtlpSameResourceTwoSpans.
func generateTraceOtlpSameResourceTwoSpans() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource: generateOtlpResource1(),
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							generateOtlpSpanOne(),
							generateOtlpSpanTwo(),
						},
					},
				},
			},
		},
	}
}

func GenerateTraceDataTwoSpansSameResourceOneDifferent() pdata.Traces {
	td := pdata.NewTraces()
	rs0 := td.ResourceSpans().AppendEmpty()
	initResource1(rs0.Resource())
	rs0ils0 := rs0.InstrumentationLibrarySpans().AppendEmpty()
	fillSpanOne(rs0ils0.Spans().AppendEmpty())
	fillSpanTwo(rs0ils0.Spans().AppendEmpty())
	rs1 := td.ResourceSpans().AppendEmpty()
	initResource2(rs1.Resource())
	rs1ils0 := rs1.InstrumentationLibrarySpans().AppendEmpty()
	fillSpanThree(rs1ils0.Spans().AppendEmpty())
	return td
}

func GenerateTraceDataManySpansSameResource(spansCount int) pdata.Traces {
	td := GenerateTraceDataOneEmptyInstrumentationLibrary()
	rs0ils0 := td.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0)
	rs0ils0.Spans().Resize(spansCount)
	for i := 0; i < spansCount; i++ {
		fillSpanOne(rs0ils0.Spans().At(i))
	}
	return td
}

func generateTraceOtlpTwoSpansSameResourceOneDifferent() *otlpcollectortrace.ExportTraceServiceRequest {
	return &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				Resource: generateOtlpResource1(),
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							generateOtlpSpanOne(),
							generateOtlpSpanTwo(),
						},
					},
				},
			},
			{
				Resource: generateOtlpResource2(),
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							generateOtlpSpanThree(),
						},
					},
				},
			},
		},
	}
}

func fillSpanOne(span pdata.Span) {
	span.SetName("operationA")
	span.SetStartTimestamp(TestSpanStartTimestamp)
	span.SetEndTimestamp(TestSpanEndTimestamp)
	span.SetDroppedAttributesCount(1)
	evs := span.Events()
	ev0 := evs.AppendEmpty()
	ev0.SetTimestamp(TestSpanEventTimestamp)
	ev0.SetName("event-with-attr")
	initSpanEventAttributes(ev0.Attributes())
	ev0.SetDroppedAttributesCount(2)
	ev1 := evs.AppendEmpty()
	ev1.SetTimestamp(TestSpanEventTimestamp)
	ev1.SetName("event")
	ev1.SetDroppedAttributesCount(2)
	span.SetDroppedEventsCount(1)
	status := span.Status()
	status.SetCode(pdata.StatusCodeError)
	status.SetMessage("status-cancelled")
}

func generateOtlpSpanOne() *otlptrace.Span {
	return &otlptrace.Span{
		Name:                   "operationA",
		StartTimeUnixNano:      uint64(TestSpanStartTimestamp),
		EndTimeUnixNano:        uint64(TestSpanEndTimestamp),
		DroppedAttributesCount: 1,
		Events: []*otlptrace.Span_Event{
			{
				Name:                   "event-with-attr",
				TimeUnixNano:           uint64(TestSpanEventTimestamp),
				Attributes:             generateOtlpSpanEventAttributes(),
				DroppedAttributesCount: 2,
			},
			{
				Name:                   "event",
				TimeUnixNano:           uint64(TestSpanEventTimestamp),
				DroppedAttributesCount: 2,
			},
		},
		DroppedEventsCount: 1,
		Status: otlptrace.Status{
			Code:           otlptrace.Status_STATUS_CODE_ERROR,
			DeprecatedCode: otlptrace.Status_DEPRECATED_STATUS_CODE_UNKNOWN_ERROR,
			Message:        "status-cancelled",
		},
	}
}

func fillSpanTwo(span pdata.Span) {
	span.SetName("operationB")
	span.SetStartTimestamp(TestSpanStartTimestamp)
	span.SetEndTimestamp(TestSpanEndTimestamp)
	link0 := span.Links().AppendEmpty()
	initSpanLinkAttributes(link0.Attributes())
	link0.SetDroppedAttributesCount(4)
	link1 := span.Links().AppendEmpty()
	link1.SetDroppedAttributesCount(4)
	span.SetDroppedLinksCount(3)
}

func generateOtlpSpanTwo() *otlptrace.Span {
	return &otlptrace.Span{
		Name:              "operationB",
		StartTimeUnixNano: uint64(TestSpanStartTimestamp),
		EndTimeUnixNano:   uint64(TestSpanEndTimestamp),
		Links: []*otlptrace.Span_Link{
			{
				Attributes:             generateOtlpSpanLinkAttributes(),
				DroppedAttributesCount: 4,
			},
			{
				DroppedAttributesCount: 4,
			},
		},
		DroppedLinksCount: 3,
	}
}

func fillSpanThree(span pdata.Span) {
	span.SetName("operationC")
	span.SetStartTimestamp(TestSpanStartTimestamp)
	span.SetEndTimestamp(TestSpanEndTimestamp)
	initSpanAttributes(span.Attributes())
	span.SetDroppedAttributesCount(5)
}

func generateOtlpSpanThree() *otlptrace.Span {
	return &otlptrace.Span{
		Name:                   "operationC",
		StartTimeUnixNano:      uint64(TestSpanStartTimestamp),
		EndTimeUnixNano:        uint64(TestSpanEndTimestamp),
		Attributes:             generateOtlpSpanAttributes(),
		DroppedAttributesCount: 5,
	}
}

func GenerateTracesManySpansSameResource(spansCount int) pdata.Traces {
	td := GenerateTraceDataOneEmptyInstrumentationLibrary()
	rs0ilm0 := td.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0)
	rs0ilm0.Spans().Resize(spansCount)
	for i := 0; i < spansCount; i++ {
		fillSpanOne(rs0ilm0.Spans().At(i))
	}
	return td
}
