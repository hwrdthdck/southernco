// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"github.com/IBM/sarama"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/traceutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil"
)

// keyableLogsMarshaler is an extension of the LogsMarshaler interface intended to provide partition key capabilities
// for log messages
type keyableLogsMarshaler interface {
	LogsMarshaler
	Key()
}

type pdataLogsMarshaler struct {
	marshaler plog.Marshaler
	encoding  string
	keyed     bool
}

// Key configures the pdataLogsMarshaler to set the message key on the kafka messages
func (p *pdataLogsMarshaler) Key() {
	p.keyed = true
}

func (p pdataLogsMarshaler) Marshal(ld plog.Logs, topic string) ([]*sarama.ProducerMessage, error) {
	var msgs []*sarama.ProducerMessage

	if p.keyed {
		logs := ld.ResourceLogs()

		for i := 0; i < logs.Len(); i++ {
			resourceLogs := logs.At(i)
			var hash = pdatautil.MapHash(resourceLogs.Resource().Attributes())

			newLogs := plog.NewLogs()
			resourceLogs.CopyTo(newLogs.ResourceLogs().AppendEmpty())

			bts, err := p.marshaler.MarshalLogs(newLogs)
			if err != nil {
				return nil, err
			}

			msgs = append(msgs, &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.ByteEncoder(bts),
				Key:   sarama.ByteEncoder(hash[:]),
			})
		}
	} else {
		bts, err := p.marshaler.MarshalLogs(ld)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(bts),
		})
	}

	return msgs, nil
}

func (p pdataLogsMarshaler) Encoding() string {
	return p.encoding
}

func newPdataLogsMarshaler(marshaler plog.Marshaler, encoding string) LogsMarshaler {
	return &pdataLogsMarshaler{
		marshaler: marshaler,
		encoding:  encoding,
	}
}

// KeyableMetricsMarshaler is an extension of the MetricsMarshaler interface intended to provide partition key capabilities
// for metrics messages
type KeyableMetricsMarshaler interface {
	MetricsMarshaler
	Key()
}

type pdataMetricsMarshaler struct {
	marshaler pmetric.Marshaler
	encoding  string
	keyed     bool
}

// Key configures the pdataMetricsMarshaler to set the message key on the kafka messages
func (p *pdataMetricsMarshaler) Key() {
	p.keyed = true
}

func (p pdataMetricsMarshaler) Marshal(ld pmetric.Metrics, topic string) ([]*sarama.ProducerMessage, error) {
	var msgs []*sarama.ProducerMessage
	if p.keyed {
		metrics := ld.ResourceMetrics()

		for i := 0; i < metrics.Len(); i++ {
			resourceMetrics := metrics.At(i)
			var hash = pdatautil.MapHash(resourceMetrics.Resource().Attributes())

			newMetrics := pmetric.NewMetrics()
			resourceMetrics.CopyTo(newMetrics.ResourceMetrics().AppendEmpty())

			bts, err := p.marshaler.MarshalMetrics(newMetrics)
			if err != nil {
				return nil, err
			}
			msgs = append(msgs, &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.ByteEncoder(bts),
				Key:   sarama.ByteEncoder(hash[:]),
			})
		}
	} else {
		bts, err := p.marshaler.MarshalMetrics(ld)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(bts),
		})
	}

	return msgs, nil
}

func (p pdataMetricsMarshaler) Encoding() string {
	return p.encoding
}

func newPdataMetricsMarshaler(marshaler pmetric.Marshaler, encoding string) MetricsMarshaler {
	return &pdataMetricsMarshaler{
		marshaler: marshaler,
		encoding:  encoding,
	}
}

// KeyableTracesMarshaler is an extension of the TracesMarshaler interface intended to provide partition key capabilities
// for trace messages
type KeyableTracesMarshaler interface {
	TracesMarshaler
	Key()
}

type pdataTracesMarshaler struct {
	marshaler ptrace.Marshaler
	encoding  string
	keyed     bool
}

func (p *pdataTracesMarshaler) Marshal(td ptrace.Traces, topic string) ([]*sarama.ProducerMessage, error) {
	var msgs []*sarama.ProducerMessage
	if p.keyed {
		for _, trace := range batchpersignal.SplitTraces(td) {
			bts, err := p.marshaler.MarshalTraces(trace)
			if err != nil {
				return nil, err
			}
			msgs = append(msgs, &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.ByteEncoder(bts),
				Key:   sarama.ByteEncoder(traceutil.TraceIDToHexOrEmptyString(trace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).TraceID())),
			})

		}
	} else {
		bts, err := p.marshaler.MarshalTraces(td)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(bts),
		})
	}

	return msgs, nil
}

func (p *pdataTracesMarshaler) Encoding() string {
	return p.encoding
}

// Key configures the pdataTracesMarshaler to set the message key on the kafka messages
func (p *pdataTracesMarshaler) Key() {
	p.keyed = true
}

func newPdataTracesMarshaler(marshaler ptrace.Marshaler, encoding string) TracesMarshaler {
	return &pdataTracesMarshaler{
		marshaler: marshaler,
		encoding:  encoding,
	}
}
