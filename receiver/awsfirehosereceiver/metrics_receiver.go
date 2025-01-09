// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	expmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/otlpmetricstream"
)

const defaultMetricsEncoding = cwmetricstream.TypeStr

// The metricsConsumer implements the firehoseConsumer
// to use a metrics consumer and unmarshaler.
type metricsConsumer struct {
	config   *Config
	settings receiver.Settings
	// consumer passes the translated metrics on to the
	// next consumer.
	consumer consumer.Metrics
	// unmarshaler is the configured pmetric.Unmarshaler
	// to use when processing the records.
	unmarshaler pmetric.Unmarshaler
}

var _ firehoseConsumer = (*metricsConsumer)(nil)

// newMetricsReceiver creates a new instance of the receiver
// with a metricsConsumer.
func newMetricsReceiver(
	config *Config,
	set receiver.Settings,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	c := &metricsConsumer{
		config:   config,
		settings: set,
		consumer: nextConsumer,
	}
	return &firehoseReceiver{
		settings: set,
		config:   config,
		consumer: c,
	}, nil
}

func (c *metricsConsumer) Start(_ context.Context, host component.Host) error {
	encoding := c.config.Encoding
	if encoding == "" {
		encoding = c.config.RecordType
		if encoding == "" {
			encoding = defaultMetricsEncoding
		}
	}
	switch encoding {
	case cwmetricstream.TypeStr:
		// TODO: make cwmetrics an encoding extension
		c.unmarshaler = cwmetricstream.NewUnmarshaler(c.settings.Logger)
	case otlpmetricstream.TypeStr:
		// TODO: make otlp_v1 an encoding extension
		c.unmarshaler = otlpmetricstream.NewUnmarshaler(c.settings.Logger)
	default:
		unmarshaler, err := loadEncodingExtension[pmetric.Unmarshaler](host, encoding, "metrics")
		if err != nil {
			return err
		}
		c.unmarshaler = unmarshaler
	}
	return nil
}

// Consume uses the configured unmarshaler to deserialize the records into a
// single pmetric.Metrics. If there are common attributes available, then it will
// attach those to each of the pcommon.Resources. It will send the final result
// to the next consumer.
func (c *metricsConsumer) Consume(ctx context.Context, records [][]byte, commonAttributes map[string]string) (int, error) {
	metrics := pmetric.NewMetrics()
	for _, record := range records {
		recordMetrics, err := c.unmarshaler.UnmarshalMetrics(record)
		if err != nil {
			return http.StatusBadRequest, err
		}
		metrics = expmetrics.Merge(metrics, recordMetrics)
	}

	if commonAttributes != nil {
		for i := 0; i < metrics.ResourceMetrics().Len(); i++ {
			rm := metrics.ResourceMetrics().At(i)
			for k, v := range commonAttributes {
				if _, found := rm.Resource().Attributes().Get(k); !found {
					rm.Resource().Attributes().PutStr(k, v)
				}
			}
		}
	}

	if err := c.consumer.ConsumeMetrics(ctx, metrics); err != nil {
		if consumererror.IsPermanent(err) {
			return http.StatusBadRequest, err
		}
		return http.StatusServiceUnavailable, err
	}
	return http.StatusOK, nil
}
