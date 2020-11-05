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

package translator

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/translator/conventions"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray"
)

const (
	// just a guess to avoid too many memory re-allocation
	initAttrCapacity = 15
)

// TODO: It might be nice to consolidate the `fromPdata` in x-ray exporter and
// `toPdata` in this receiver to a common package later

// ToTraces converts X-Ray segment (and its subsegments) to an OT ResourceSpans.
func ToTraces(rawSeg []byte) (*pdata.Traces, int, error) {
	var seg awsxray.Segment
	err := json.Unmarshal(rawSeg, &seg)
	if err != nil {
		// return 1 as total segment (&subsegments) count
		// because we can't parse the body the UDP packet.
		return nil, 1, err
	}
	count := totalSegmentsCount(seg)

	err = seg.Validate()
	if err != nil {
		return nil, count, err
	}

	traceData := pdata.NewTraces()
	rspanSlice := traceData.ResourceSpans()
	// ## allocate a new pdata.ResourceSpans for the segment document
	// (potentially with embedded subsegments)
	rspanSlice.Resize(1)      // initialize a new empty pdata.ResourceSpans
	rspan := rspanSlice.At(0) // retrieve the empty pdata.ResourceSpans we just created

	// ## initialize the fields in a ResourceSpans
	resource := rspan.Resource()
	resource.InitEmpty()
	// each segment (with its subsegments) is generated by one instrument
	// library so only allocate one `InstrumentationLibrarySpans` in the
	// `InstrumentationLibrarySpansSlice`.
	rspan.InstrumentationLibrarySpans().Resize(1)
	ils := rspan.InstrumentationLibrarySpans().At(0)
	ils.Spans().Resize(count)
	spans := ils.Spans()

	// populating global attributes shared among segment and embedded subsegment(s)
	populateResource(&seg, &resource)

	// recursively traverse segment and embedded subsegments
	// to populate the spans. We also need to pass in the
	// TraceID of the root segment in because embedded subsegments
	// do not have that information, but it's needed after we flatten
	// the embedded subsegment to generate independent child spans.
	_, _, err = segToSpans(seg, seg.TraceID, nil, &spans, 0)
	if err != nil {
		return nil, count, err
	}

	return &traceData, count, nil
}

func segToSpans(seg awsxray.Segment,
	traceID, parentID *string,
	spans *pdata.SpanSlice, startingIndex int) (int, *pdata.Span, error) {

	span := spans.At(startingIndex)

	err := populateSpan(&seg, traceID, parentID, &span)
	if err != nil {
		return 0, nil, err
	}

	startingIndexForSubsegment := 1 + startingIndex
	var populatedChildSpan *pdata.Span
	for _, s := range seg.Subsegments {
		startingIndexForSubsegment, populatedChildSpan, err = segToSpans(s,
			traceID, seg.ID,
			spans, startingIndexForSubsegment)
		if err != nil {
			return 0, nil, err
		}

		if seg.Cause != nil &&
			populatedChildSpan.Status().Code() != pdata.StatusCodeUnset {
			// if seg.Cause is not nil, then one of the subsegments must contain a
			// HTTP error code. Also, span.Status().Code() is already
			// set to `StatusCodeUnknownError` by `addCause()` in
			// `populateSpan()` above, so here we are just trying to figure out
			// whether we can get an even more specific error code.

			if span.Status().Code() == pdata.StatusCodeError {
				// update the error code to a possibly more specific code
				span.Status().SetCode(populatedChildSpan.Status().Code())
			}
		}
	}

	return startingIndexForSubsegment, &span, nil
}

func populateSpan(
	seg *awsxray.Segment,
	traceID, parentID *string,
	span *pdata.Span) error {

	span.Status().InitEmpty() // by default this sets the code to `Status_Unset`
	attrs := span.Attributes()
	attrs.InitEmptyWithCapacity(initAttrCapacity)

	err := addNameAndNamespace(seg, span)
	if err != nil {
		return err
	}

	// decode trace id
	var traceIDBytes [16]byte
	if seg.TraceID == nil {
		// if seg.TraceID is nil, then `seg` must be an embedded subsegment.
		traceIDBytes, err = decodeXRayTraceID(traceID)
		if err != nil {
			return err
		}

	} else {
		traceIDBytes, err = decodeXRayTraceID(seg.TraceID)
		if err != nil {
			return err
		}

	}

	// decode parent id
	var parentIDBytes [8]byte
	if parentID != nil {
		parentIDBytes, err = decodeXRaySpanID(parentID)
		if err != nil {
			return err
		}
	} else if seg.ParentID != nil {
		parentIDBytes, err = decodeXRaySpanID(seg.ParentID)
		if err != nil {
			return err
		}
	}

	// decode span id
	spanIDBytes, err := decodeXRaySpanID(seg.ID)
	if err != nil {
		return err
	}

	span.SetTraceID(pdata.NewTraceID(traceIDBytes))
	span.SetSpanID(pdata.NewSpanID(spanIDBytes))

	if parentIDBytes != [8]byte{} {
		span.SetParentSpanID(pdata.NewSpanID(parentIDBytes))
	} else {
		span.SetKind(pdata.SpanKindSERVER)
	}

	addStartTime(seg.StartTime, span)
	addEndTime(seg.EndTime, span)
	addBool(seg.InProgress, awsxray.AWSXRayInProgressAttribute, &attrs)
	addString(seg.User, conventions.AttributeEnduserID, &attrs)

	addHTTP(seg, span)
	addCause(seg, span)
	addAWSToSpan(seg.AWS, &attrs)
	err = addSQLToSpan(seg.SQL, &attrs)
	if err != nil {
		return err
	}

	addBool(seg.Traced, awsxray.AWSXRayTracedAttribute, &attrs)

	addAnnotations(seg.Annotations, &attrs)
	addMetadata(seg.Metadata, &attrs)

	return nil
}

func populateResource(seg *awsxray.Segment, rs *pdata.Resource) {
	// allocate a new attribute map within the Resource in the pdata.ResourceSpans allocated above
	attrs := rs.Attributes()
	attrs.InitEmptyWithCapacity(initAttrCapacity)

	addAWSToResource(seg.AWS, &attrs)
	addSdkToResource(seg, &attrs)
	if seg.Service != nil {
		addString(
			seg.Service.Version,
			conventions.AttributeServiceVersion,
			&attrs)
	}

	addString(seg.ResourceARN, awsxray.AWSXRayResourceARNAttribute, &attrs)
}

func totalSegmentsCount(seg awsxray.Segment) int {
	subsegmentCount := 0
	for _, s := range seg.Subsegments {
		subsegmentCount += totalSegmentsCount(s)
	}

	return 1 + subsegmentCount
}

/*
decodeXRayTraceID decodes the traceid from xraysdk
one example of xray format: "1-5f84c7a1-e7d1852db8c4fd35d88bf49a"
decodeXRayTraceID transfers it to "5f84c7a1e7d1852db8c4fd35d88bf49a" and decode it from hex
*/
func decodeXRayTraceID(traceID *string) ([16]byte, error) {
	tid := [16]byte{}

	if traceID == nil {
		return tid, errors.New("traceID is null")
	}
	if len(*traceID) < 35 {
		return tid, errors.New("traceID length is wrong")
	}
	traceIDtoBeDecoded := (*traceID)[2:10] + (*traceID)[11:]

	_, err := hex.Decode(tid[:], []byte(traceIDtoBeDecoded))
	return tid, err
}

// decodeXRaySpanID decodes the spanid from xraysdk
func decodeXRaySpanID(spanID *string) ([8]byte, error) {
	sid := [8]byte{}
	if spanID == nil {
		return sid, errors.New("spanid is null")
	}
	if len(*spanID) != 16 {
		return sid, errors.New("spanID length is wrong")
	}
	_, err := hex.Decode(sid[:], []byte(*spanID))
	return sid, err
}
