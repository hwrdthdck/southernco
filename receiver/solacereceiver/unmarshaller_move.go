// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/google/uuid"
	move_v1 "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/model/move/v1"
)

type brokerTraceMoveUnmarshallerV1 struct {
	logger  *zap.Logger
	metrics *opencensusMetrics
}

// unmarshal implements tracesUnmarshaller.unmarshal
func (u *brokerTraceMoveUnmarshallerV1) unmarshal(message *inboundMessage) (ptrace.Traces, error) {
	spanData, err := u.unmarshalToSpanData(message)
	if err != nil {
		return ptrace.Traces{}, err
	}
	traces := ptrace.NewTraces()
	u.populateTraces(spanData, traces)
	return traces, nil
}

// unmarshalToSpanData will consume an solaceMessage and unmarshal it into a SpanData.
// Returns an error if one occurred.
func (u *brokerTraceMoveUnmarshallerV1) unmarshalToSpanData(message *inboundMessage) (*move_v1.SpanData, error) {
	var data = message.GetData()
	if len(data) == 0 {
		return nil, errEmptyPayload
	}
	var spanData move_v1.SpanData
	if err := proto.Unmarshal(data, &spanData); err != nil {
		return nil, err
	}
	return &spanData, nil
}

// populateTraces will create a new Span from the given traces and map the given SpanData to the span.
// This will set all required fields such as name version, trace and span ID, parent span ID (if applicable),
// timestamps, errors and states.
func (u *brokerTraceMoveUnmarshallerV1) populateTraces(spanData *move_v1.SpanData, traces ptrace.Traces) {
	// Append new resource span and map any attributes
	resourceSpan := traces.ResourceSpans().AppendEmpty()
	u.mapResourceSpanAttributes(resourceSpan.Resource().Attributes())
	instrLibrarySpans := resourceSpan.ScopeSpans().AppendEmpty()
	// Create a new span
	clientSpan := instrLibrarySpans.Spans().AppendEmpty()
	// map the tracing data for the span
	u.mapMoveSpanTracingInfo(spanData, clientSpan)
	// map the basic span data
	u.mapClientSpanData(spanData, clientSpan)
}

func (u *brokerTraceMoveUnmarshallerV1) mapResourceSpanAttributes(attrMap pcommon.Map) {
	routerName := "internal-" + uuid.New().String()
	version := "0.0." + uuid.New().String() // random uuid string as the version
	setResourceSpanAttributes(attrMap, routerName, version, nil)
}

func (u *brokerTraceMoveUnmarshallerV1) mapMoveSpanTracingInfo(spanData *move_v1.SpanData, span ptrace.Span) {

	// hard coded to internal span
	// SPAN_KIND_CONSUMER == 1
	span.SetKind(ptrace.SpanKindInternal)

	// map trace ID
	var traceID [16]byte
	copy(traceID[:16], spanData.TraceId)
	span.SetTraceID(traceID)
	// map span ID
	var spanID [8]byte
	copy(spanID[:8], spanData.SpanId)
	span.SetSpanID(spanID)
	// conditional parent-span-id
	if len(spanData.ParentSpanId) == 8 {
		var parentSpanID [8]byte
		copy(parentSpanID[:8], spanData.ParentSpanId)
		span.SetParentSpanID(parentSpanID)
	}
	// timestamps
	span.SetStartTimestamp(pcommon.Timestamp(spanData.GetStartTimeUnixNano()))
	span.SetEndTimestamp(pcommon.Timestamp(spanData.GetEndTimeUnixNano()))
}

func (u *brokerTraceMoveUnmarshallerV1) mapClientSpanData(moveSpan *move_v1.SpanData, span ptrace.Span) {
	const (
		sourceNameKey          = "messaging.source.name"
		sourceKindKey          = "messaging.solace.source.kind"
		destinationNameKey     = "messaging.destination.name"
		destinationKindKey     = "messaging.solace.destination.kind"
		moveOperationReasonKey = "messaging.solace.operation.reason"
	)
	const (
		spanOperation         = "move"
		moveNameSuffix        = " move"
		unknownEndpointName   = "(unknown)"
		anonymousEndpointName = "(anonymous)"
	)
	// Delete Info reasons
	const (
		ttlExpired              = "ttl_expired"
		rejectedNack            = "rejected_nack"
		maxRedeliveriesExceeded = "max_redeliveries_exceeded"
	)

	attributes := span.Attributes()
	attributes.PutStr(systemAttrKey, systemAttrValue)
	attributes.PutStr(operationAttrKey, spanOperation)

	// set source endpoint information
	// don't fatal out when we receive invalid endpoint name, instead just log and increment stats
	var sourceEndpointName string
	var sourceEndpointType = "(unknown)"
	switch casted := moveSpan.Source.(type) {
	case *move_v1.SpanData_SourceTopicEndpointName:
		if isAnonymousTopicEndpoint(casted.SourceTopicEndpointName) {
			sourceEndpointName = anonymousEndpointName
		} else {
			sourceEndpointName = casted.SourceTopicEndpointName
		}
		sourceEndpointType = "topic"
		attributes.PutStr(sourceNameKey, casted.SourceTopicEndpointName)
		attributes.PutStr(sourceKindKey, topicEndpointKind)
	case *move_v1.SpanData_SourceQueueName:
		if isAnonymousQueue(casted.SourceQueueName) {
			sourceEndpointName = anonymousEndpointName
		} else {
			sourceEndpointName = casted.SourceQueueName
		}
		sourceEndpointType = "queue"
		attributes.PutStr(sourceNameKey, casted.SourceQueueName)
		attributes.PutStr(sourceKindKey, queueKind)
	default:
		u.logger.Warn(fmt.Sprintf("Unknown source endpoint type %T", casted))
		u.metrics.recordRecoverableUnmarshallingError()
		sourceEndpointName = unknownEndpointName
	}
	span.SetName("(" + sourceEndpointType + ": \"" + sourceEndpointName + "\")" + moveNameSuffix)

	// set destination endpoint information
	// don't fatal out when we receive invalid endpoint name, instead just log and increment stats
	switch casted := moveSpan.Destination.(type) {
	case *move_v1.SpanData_DestinationTopicEndpointName:
		attributes.PutStr(destinationNameKey, casted.DestinationTopicEndpointName)
		attributes.PutStr(destinationKindKey, topicEndpointKind)
	case *move_v1.SpanData_DestinationQueueName:
		attributes.PutStr(destinationNameKey, casted.DestinationQueueName)
		attributes.PutStr(destinationKindKey, queueKind)
	default:
		u.logger.Warn(fmt.Sprintf("Unknown endpoint type %T", casted))
		u.metrics.recordRecoverableUnmarshallingError()
	}

	// do not fatal out when we don't have a valid move reason name
	// instead just log and increment stats
	switch casted := moveSpan.TypeInfo.(type) {
	// caused by expired ttl on message
	case *move_v1.SpanData_TtlExpiredInfo:
		attributes.PutStr(moveOperationReasonKey, ttlExpired)
	// caused by consumer N(ack)ing with Rejected outcome
	case *move_v1.SpanData_RejectedOutcomeInfo:
		attributes.PutStr(moveOperationReasonKey, rejectedNack)
	// caused by max redelivery reached/exceeded
	case *move_v1.SpanData_MaxRedeliveriesInfo:
		attributes.PutStr(moveOperationReasonKey, maxRedeliveriesExceeded)
	default:
		u.logger.Warn(fmt.Sprintf("Unknown move reason info type %T", casted))
		u.metrics.recordRecoverableUnmarshallingError()
	}
}
