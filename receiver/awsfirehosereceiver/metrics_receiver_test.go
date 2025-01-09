// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/unmarshalertest"
)

type recordConsumer struct {
	result pmetric.Metrics
}

var _ consumer.Metrics = (*recordConsumer)(nil)

func (rc *recordConsumer) ConsumeMetrics(_ context.Context, metrics pmetric.Metrics) error {
	rc.result = metrics
	return nil
}

func (rc *recordConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func TestMetricsReceiver_Start(t *testing.T) {
	testCases := map[string]struct {
		encoding            string
		recordType          string
		wantUnmarshalerType pmetric.Unmarshaler
		wantErr             string
	}{
		"WithDefaultEncoding": {
			wantUnmarshalerType: &cwmetricstream.Unmarshaler{},
		},
		"WithBuiltinEncoding": {
			encoding:            "cwmetrics",
			wantUnmarshalerType: &cwmetricstream.Unmarshaler{},
		},
		"WithExtensionEncoding": {
			encoding:            "otlp_metrics",
			wantUnmarshalerType: pmetricUnmarshalerExtension{},
		},
		"WithDeprecatedRecordType": {
			recordType:          "otlp_metrics",
			wantUnmarshalerType: pmetricUnmarshalerExtension{},
		},
		"WithUnknownEncoding": {
			encoding: "invalid",
			wantErr:  `unknown encoding extension "invalid"`,
		},
		"WithNonLogUnmarshalerExtension": {
			encoding: "otlp_logs",
			wantErr:  `extension "otlp_logs" is not a metrics unmarshaler`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			cfg := createDefaultConfig().(*Config)
			cfg.Encoding = testCase.encoding
			cfg.RecordType = testCase.recordType
			got, err := newMetricsReceiver(
				cfg,
				receivertest.NewNopSettings(),
				consumertest.NewNop(),
			)
			require.NoError(t, err)
			require.NotNil(t, got)
			t.Cleanup(func() {
				require.NoError(t, got.Shutdown(context.Background()))
			})

			host := hostWithExtensions{
				extensions: map[component.ID]component.Component{
					component.MustNewID("otlp_logs"):    plogUnmarshalerExtension{},
					component.MustNewID("otlp_metrics"): pmetricUnmarshalerExtension{},
				},
			}

			err = got.Start(context.Background(), host)
			if testCase.wantErr != "" {
				require.EqualError(t, err, testCase.wantErr)
			} else {
				require.NoError(t, err)
			}

			assert.IsType(t,
				testCase.wantUnmarshalerType,
				got.(*firehoseReceiver).consumer.(*metricsConsumer).unmarshaler,
			)
		})
	}
}

func TestMetricsConsumer(t *testing.T) {
	testErr := errors.New("test error")
	testCases := map[string]struct {
		unmarshalerErr error
		consumerErr    error
		wantStatus     int
		wantErr        error
	}{
		"WithUnmarshalerError": {
			unmarshalerErr: testErr,
			wantStatus:     http.StatusBadRequest,
			wantErr:        testErr,
		},
		"WithConsumerErrorPermanent": {
			consumerErr: consumererror.NewPermanent(testErr),
			wantStatus:  http.StatusBadRequest,
			wantErr:     consumererror.NewPermanent(testErr),
		},
		"WithConsumerError": {
			consumerErr: testErr,
			wantStatus:  http.StatusServiceUnavailable,
			wantErr:     testErr,
		},
		"WithNoError": {
			wantStatus: http.StatusOK,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			mc := &metricsConsumer{
				unmarshaler: unmarshalertest.NewErrMetrics(testCase.unmarshalerErr),
				consumer:    consumertest.NewErr(testCase.consumerErr),
			}
			gotStatus, gotErr := mc.Consume(context.Background(), [][]byte{{0}}, nil)
			require.Equal(t, testCase.wantStatus, gotStatus)
			require.Equal(t, testCase.wantErr, gotErr)
		})
	}

	t.Run("WithCommonAttributes", func(t *testing.T) {
		base := pmetric.NewMetrics()
		base.ResourceMetrics().AppendEmpty()
		rc := recordConsumer{}
		mc := &metricsConsumer{
			unmarshaler: unmarshalertest.NewWithMetrics(base),
			consumer:    &rc,
		}
		gotStatus, gotErr := mc.Consume(context.Background(), [][]byte{{0}}, map[string]string{
			"CommonAttributes": "Test",
		})
		require.Equal(t, http.StatusOK, gotStatus)
		require.NoError(t, gotErr)
		gotRms := rc.result.ResourceMetrics()
		require.Equal(t, 1, gotRms.Len())
		gotRm := gotRms.At(0)
		require.Equal(t, 1, gotRm.Resource().Attributes().Len())
	})
}
