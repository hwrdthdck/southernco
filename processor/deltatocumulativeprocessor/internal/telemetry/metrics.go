package telemetry

import (
	"context"
	"errors"
	"reflect"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func New(set component.TelemetrySettings) (Metrics, error) {
	m := Metrics{
		tracked: func() int { return 0 },
	}

	trackedCb := metadata.WithDeltatocumulativeStreamsTrackedCallback(func() int64 {
		return int64(m.tracked())
	})

	telb, err := metadata.NewTelemetryBuilder(set, trackedCb)
	if err != nil {
		return Metrics{}, err
	}
	m.TelemetryBuilder = *telb

	return m, nil
}

type Metrics struct {
	metadata.TelemetryBuilder

	tracked func() int
}

func (m Metrics) Datapoints() Counter {
	return Counter{Int64Counter: m.DeltatocumulativeDatapoints}
}

func (m *Metrics) WithTracked(streams func() int) {
	m.tracked = streams
}

func Error(msg string) attribute.KeyValue {
	return attribute.String("error", msg)
}

func Cause(err error) attribute.KeyValue {
	for {
		uw := errors.Unwrap(err)
		if uw == nil {
			break
		}
		err = uw
	}

	return Error(reflect.TypeOf(err).String())
}

type Counter struct{ metric.Int64Counter }

func (c Counter) Inc(ctx context.Context, attrs ...attribute.KeyValue) {
	c.Add(ctx, 1, metric.WithAttributes(attrs...))
}
