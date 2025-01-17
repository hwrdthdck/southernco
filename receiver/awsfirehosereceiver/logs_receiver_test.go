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
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/unmarshalertest"
)

type logsRecordConsumer struct {
	result plog.Logs
}

var _ consumer.Logs = (*logsRecordConsumer)(nil)

func (rc *logsRecordConsumer) ConsumeLogs(_ context.Context, logs plog.Logs) error {
	rc.result = logs
	return nil
}

func (rc *logsRecordConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func TestLogsReceiver_Start(t *testing.T) {
	testCases := map[string]struct {
		encoding            string
		recordType          string
		wantUnmarshalerType plog.Unmarshaler
		wantErr             string
	}{
		"WithDefaultEncoding": {
			wantUnmarshalerType: &cwlog.Unmarshaler{},
		},
		"WithBuiltinEncoding": {
			encoding:            "cwlogs",
			wantUnmarshalerType: &cwlog.Unmarshaler{},
		},
		"WithExtensionEncoding": {
			encoding:            "otlp_logs",
			wantUnmarshalerType: plogUnmarshalerExtension{},
		},
		"WithDeprecatedRecordType": {
			recordType:          "otlp_logs",
			wantUnmarshalerType: plogUnmarshalerExtension{},
		},
		"WithUnknownEncoding": {
			encoding: "invalid",
			wantErr:  "unknown encoding extension \"invalid\"",
		},
		"WithNonLogUnmarshalerExtension": {
			encoding: "otlp_metrics",
			wantErr:  `extension "otlp_metrics" is not a logs unmarshaler`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			cfg := createDefaultConfig().(*Config)
			cfg.Encoding = testCase.encoding
			cfg.RecordType = testCase.recordType
			got, err := newLogsReceiver(
				cfg,
				receivertest.NewNopSettings(),
				consumertest.NewNop(),
			)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.IsType(t, &firehoseReceiver{}, got)
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
				got.(*firehoseReceiver).consumer.(*logsConsumer).unmarshaler,
			)
		})
	}
}

func TestLogsConsumer_Errors(t *testing.T) {
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
			lc := &logsConsumer{
				unmarshaler: unmarshalertest.NewErrLogs(testCase.unmarshalerErr),
				consumer:    consumertest.NewErr(testCase.consumerErr),
			}
			gotStatus, gotErr := lc.Consume(context.TODO(), [][]byte{{0}}, nil)
			require.Equal(t, testCase.wantStatus, gotStatus)
			require.Equal(t, testCase.wantErr, gotErr)
		})
	}
}

func TestLogsConsumer(t *testing.T) {
	t.Run("WithCommonAttributes", func(t *testing.T) {
		base := plog.NewLogs()
		base.ResourceLogs().AppendEmpty()
		rc := logsRecordConsumer{}
		lc := &logsConsumer{
			unmarshaler: unmarshalertest.NewWithLogs(base),
			consumer:    &rc,
		}
		gotStatus, gotErr := lc.Consume(context.TODO(), [][]byte{{0}}, map[string]string{
			"CommonAttributes": "Test",
		})
		require.Equal(t, http.StatusOK, gotStatus)
		require.NoError(t, gotErr)
		gotRms := rc.result.ResourceLogs()
		require.Equal(t, 1, gotRms.Len())
		gotRm := gotRms.At(0)
		require.Equal(t, 1, gotRm.Resource().Attributes().Len())
	})
	t.Run("WithMultipleRecords", func(t *testing.T) {
		// service0, scope0
		logs0, logRecords0, resource0, scope0 := newLogs()
		resource0.Attributes().PutStr("service.name", "service0")
		scope0.SetName("scope0")
		logRecords0.AppendEmpty().Body().SetStr("record0")
		logRecords0.AppendEmpty().Body().SetStr("record1")

		// service0, [scope0. scope1]
		// these scopes should be merged into the above resource logs
		logs1, logRecords1, resource1, scope0 := newLogs()
		resource1.Attributes().PutStr("service.name", "service0")
		scope0.SetName("scope0")
		logRecords1.AppendEmpty().Body().SetStr("record2")
		scopeLogs1 := logs1.ResourceLogs().At(0).ScopeLogs().AppendEmpty()
		scopeLogs1.Scope().SetName("scope1")
		scopeLogs1.LogRecords().AppendEmpty().Body().SetStr("record3")

		// service1, scope0
		logs2, logRecords2, resource2, scope2 := newLogs()
		resource2.Attributes().PutStr("service.name", "service1")
		scope2.SetName("scope0")
		logRecords2.AppendEmpty().Body().SetStr("record4")
		logRecords2.AppendEmpty().Body().SetStr("record5")

		logsRemaining := []plog.Logs{logs0, logs1, logs2}
		var unmarshaler unmarshalLogsFunc = func(data []byte) (plog.Logs, error) {
			logs := logsRemaining[0]
			logsRemaining = logsRemaining[1:]
			return logs, nil
		}

		rc := logsRecordConsumer{}
		lc := &logsConsumer{unmarshaler: unmarshaler, consumer: &rc}
		gotStatus, gotErr := lc.Consume(context.Background(), make([][]byte, len(logsRemaining)), nil)
		require.Equal(t, http.StatusOK, gotStatus)
		require.NoError(t, gotErr)
		assert.Equal(t, 6, rc.result.LogRecordCount())
		assert.Equal(t, 2, rc.result.ResourceLogs().Len()) // service0, service1

		type scopeRecords struct {
			scope   string
			records []string
		}

		type resourceScopeRecords struct {
			serviceName string
			scopes      []scopeRecords
		}

		var merged []resourceScopeRecords
		for i := 0; i < rc.result.ResourceLogs().Len(); i++ {
			resourceLogs := rc.result.ResourceLogs().At(i)
			serviceName, ok := resourceLogs.Resource().Attributes().Get("service.name")
			require.True(t, ok)

			var scopes []scopeRecords
			for i := 0; i < resourceLogs.ScopeLogs().Len(); i++ {
				scopeLogs := resourceLogs.ScopeLogs().At(i)
				scope := scopeRecords{scope: scopeLogs.Scope().Name()}
				for i := 0; i < scopeLogs.LogRecords().Len(); i++ {
					record := scopeLogs.LogRecords().At(i)
					scope.records = append(scope.records, record.Body().Str())
				}
				scopes = append(scopes, scope)
			}
			merged = append(merged, resourceScopeRecords{serviceName: serviceName.Str(), scopes: scopes})
		}
		assert.Equal(t, []resourceScopeRecords{{
			serviceName: "service0",
			scopes: []scopeRecords{{
				scope:   "scope0",
				records: []string{"record0", "record1", "record2"},
			}, {
				scope:   "scope1",
				records: []string{"record3"},
			}},
		}, {
			serviceName: "service1",
			scopes: []scopeRecords{{
				scope:   "scope0",
				records: []string{"record4", "record5"},
			}},
		}}, merged)
	})
}

func newLogs() (plog.Logs, plog.LogRecordSlice, pcommon.Resource, pcommon.InstrumentationScope) {
	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
	return logs, scopeLogs.LogRecords(), resourceLogs.Resource(), scopeLogs.Scope()
}

type unmarshalLogsFunc func([]byte) (plog.Logs, error)

func (f unmarshalLogsFunc) UnmarshalLogs(data []byte) (plog.Logs, error) {
	return f(data)
}
