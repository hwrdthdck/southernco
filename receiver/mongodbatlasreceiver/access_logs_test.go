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

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/extension/experimental/storage"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap/zaptest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
)

var (
	authTrue  = true
	authFalse = false
)

func TestAccessLogToLogRecord(t *testing.T) {
	now := pcommon.NewTimestampFromTime(time.Time{})

	proj := &mongodbatlas.Project{
		ID:    testProjectID,
		OrgID: testOrgID,
		Name:  testProjectName,
		Links: []*mongodbatlas.Link{},
	}

	cluster := mongodbatlas.Cluster{
		GroupID: testProjectID,
		Name:    testClusterName,
	}

	inputLogs := []*mongodbatlas.AccessLogs{
		{
			GroupID:     testProjectID,
			Hostname:    "test-hostname.mongodb.net",
			ClusterName: testClusterName,
			IPAddress:   "192.168.1.1",
			AuthResult:  &authTrue,
			AuthSource:  "admin",
			LogLine:     "{\"t\":{\"$date\":\"2023-04-26T02:38:56.444+00:00\"},\"s\":\"I\",  \"c\":\"ACCESS\",   \"id\":20249,   \"ctx\":\"conn173\",\"msg\":\"Authentication failed\",\"attr\":{\"mechanism\":\"SCRAM-SHA-1\",\"speculative\":true,\"principalName\":\"mms-automation\",\"authenticationDatabase\":\"admin\",\"remote\":\"192.168.248.4:41052\",\"extraInfo\":{}}}",
			Timestamp:   "Wed Apr 26 02:38:56 GMT 2023",
			Username:    "test",
		},
		{
			GroupID:       testProjectID,
			Hostname:      "test-hostname.mongodb.net",
			ClusterName:   testClusterName,
			IPAddress:     "192.168.1.1",
			AuthResult:    &authFalse,
			FailureReason: "User not found",
			AuthSource:    "admin",
			LogLine:       "{\"s\":\"I\",  \"c\":\"ACCESS\",   \"id\":20249,   \"ctx\":\"conn173\",\"msg\":\"Authentication failed\",\"attr\":{\"mechanism\":\"SCRAM-SHA-1\",\"speculative\":true,\"principalName\":\"mms-automation\",\"authenticationDatabase\":\"admin\",\"remote\":\"192.168.248.4:41052\",\"extraInfo\":{},\"error\":\"UserNotFound: User \\\"mms-automation@admin\\\" not found\"}}",
			Timestamp:     "Wed Apr 26 02:38:56 GMT 2023",
			Username:      "test",
		},
	}

	expectedLogs := plog.NewLogs()
	rl := expectedLogs.ResourceLogs().AppendEmpty()

	assert.NoError(t, rl.Resource().Attributes().FromRaw(map[string]interface{}{
		"mongodbatlas.project.name": testProjectName,
		"mongodbatlas.project.id":   testProjectID,
		"mongodbatlas.org.id":       testOrgID,
		"mongodbatlas.cluster.name": testClusterName,
	}))

	records := rl.ScopeLogs().AppendEmpty().LogRecords()
	// First log is an example of a success, and tests that the timestamp works parsed from the log line
	lr := records.AppendEmpty()
	assert.NoError(t, lr.Attributes().FromRaw(map[string]interface{}{
		"event.domain": "mongodbatlas",
		"auth.result":  "success",
		"auth.source":  "admin",
		"username":     "test",
		"hostname":     "test-hostname.mongodb.net",
		"remote.ip":    "192.168.1.1",
	}))

	lr.SetObservedTimestamp(now)
	lr.SetTimestamp(pcommon.NewTimestampFromTime(time.Date(2023, time.April, 26, 02, 38, 56, 444000000, time.UTC)))
	lr.SetSeverityNumber(plog.SeverityNumberInfo)
	lr.SetSeverityText(plog.SeverityNumberInfo.String())

	var logBody map[string]interface{}
	assert.NoError(t, json.Unmarshal([]byte(inputLogs[0].LogLine), &logBody))
	assert.NoError(t, lr.Body().SetEmptyMap().FromRaw(logBody))

	// Second log is an example of a failure, and tests that the timestamp is missing from the log line
	lr = records.AppendEmpty()
	assert.NoError(t, lr.Attributes().FromRaw(map[string]interface{}{
		"event.domain":        "mongodbatlas",
		"auth.result":         "failure",
		"auth.failure_reason": "User not found",
		"auth.source":         "admin",
		"username":            "test",
		"hostname":            "test-hostname.mongodb.net",
		"remote.ip":           "192.168.1.1",
	}))

	lr.SetObservedTimestamp(now)
	// Second log does not have internal timestamp in ISO8601, it has external in unixDate format with less precision
	lr.SetTimestamp(pcommon.NewTimestampFromTime(time.Date(2023, time.April, 26, 02, 38, 56, 0, time.UTC)))
	lr.SetSeverityNumber(plog.SeverityNumberWarn)
	lr.SetSeverityText(plog.SeverityNumberWarn.String())

	logBody = map[string]interface{}{}
	assert.NoError(t, json.Unmarshal([]byte(inputLogs[1].LogLine), &logBody))
	assert.NoError(t, lr.Body().SetEmptyMap().FromRaw(logBody))

	logs := transformAccessLogs(now, inputLogs, proj, cluster, zaptest.NewLogger(t))

	require.NotNil(t, logs)
	require.NoError(t, plogtest.CompareLogs(expectedLogs, logs))
}

func TestAccessLogsRetrieval(t *testing.T) {
	cases := []struct {
		name             string
		config           func() *Config
		setup            func(rcvr *accessLogsReceiver)
		expectedLogCount int
		validateEntries  func(*testing.T, []plog.Logs)
	}{
		{
			name: "basic",
			config: func() *Config {
				return &Config{
					ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(typeStr),
					Granularity:               defaultGranularity,
					RetrySettings:             exporterhelper.NewDefaultRetrySettings(),
					AccessLogs: &AccessLogsConfig{
						Projects: []ProjectConfig{
							{
								Name: testProjectName,
							},
						},
						PollInterval: 1 * time.Second,
					},
				}
			},
			setup: func(rcvr *accessLogsReceiver) {
				rcvr.client = simpleAccessLogClient()
			},
			expectedLogCount: 1,
			validateEntries: func(t *testing.T, logs []plog.Logs) {
				l := logs[0]
				expectedStringAttributes := map[string]string{
					"event.domain": "mongodbatlas",
					"auth.result":  "success",
					"auth.source":  "admin",
					"username":     "test",
					"hostname":     "test-hostname.mongodb.net",
					"remote.ip":    "192.168.1.1",
				}
				validateAttributes(t, expectedStringAttributes, l)
				expectedResourceAttributes := map[string]string{
					"mongodbatlas.cluster.name": testClusterName,
					"mongodbatlas.project.name": testProjectName,
					"mongodbatlas.project.id":   testProjectID,
					"mongodbatlas.org.id":       testOrgID,
				}

				ra := l.ResourceLogs().At(0).Resource().Attributes()
				for k, v := range expectedResourceAttributes {
					value, ok := ra.Get(k)
					require.True(t, ok)
					require.Equal(t, v, value.AsString())
				}
			},
		},
		{
			name: "multiple page read all",
			config: func() *Config {
				return &Config{
					ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(typeStr),
					Granularity:               defaultGranularity,
					RetrySettings:             exporterhelper.NewDefaultRetrySettings(),
					AccessLogs: &AccessLogsConfig{
						Projects: []ProjectConfig{
							{
								Name: testProjectName,
							},
						},
						PollInterval: 2 * time.Second,
						PageSize:     2,
					},
				}
			},
			setup: func(rcvr *accessLogsReceiver) {
				rcvr.client = repeatedRequestAccessLogClient()
			},
			expectedLogCount: 3,
			validateEntries: func(t *testing.T, logs []plog.Logs) {
				require.Equal(t, 1, logs[0].ResourceLogs().Len())
				require.Equal(t, 1, logs[0].ResourceLogs().At(0).ScopeLogs().Len())
				require.Equal(t, 2, logs[0].ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().Len())

				require.Equal(t, 1, logs[1].ResourceLogs().Len())
				require.Equal(t, 1, logs[1].ResourceLogs().At(0).ScopeLogs().Len())
				require.Equal(t, 1, logs[1].ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().Len())
			},
		},
		{
			name: "multiple page break early based on timestamp",
			config: func() *Config {
				return &Config{
					ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(typeStr),
					Granularity:               defaultGranularity,
					RetrySettings:             exporterhelper.NewDefaultRetrySettings(),
					AccessLogs: &AccessLogsConfig{
						Projects: []ProjectConfig{
							{
								Name: testProjectName,
							},
						},
						PollInterval: 1 * time.Second,
						PageSize:     2,
					},
				}
			},
			setup: func(rcvr *accessLogsReceiver) {
				rcvr.client = maxSizeButOldDataAccessLogsClient()
			},
			expectedLogCount: 2,
			validateEntries: func(t *testing.T, logs []plog.Logs) {
				require.Equal(t, 1, logs[0].ResourceLogs().Len())
				require.Equal(t, 1, logs[0].ResourceLogs().At(0).ScopeLogs().Len())
				require.Equal(t, 2, logs[0].ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().Len())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			logSink := &consumertest.LogsSink{}
			rcvr := newAccessLogsReceiver(receivertest.NewNopCreateSettings(), tc.config(), logSink)
			tc.setup(rcvr)

			err := rcvr.Start(context.Background(), componenttest.NewNopHost(), storage.NewNopClient())
			require.NoError(t, err)

			require.Eventually(t, func() bool {
				return logSink.LogRecordCount() >= tc.expectedLogCount
			}, 20*time.Second, 10*time.Millisecond)

			require.NoError(t, rcvr.Shutdown(context.Background()))
			tc.validateEntries(t, logSink.AllLogs())
		})
	}
}

func testClientBase() *mockAccessLogsClient {
	ac := &mockAccessLogsClient{}
	ac.On("GetProject", mock.Anything, mock.Anything).Return(&mongodbatlas.Project{
		ID:    testProjectID,
		OrgID: testOrgID,
		Name:  testProjectName,
		Links: []*mongodbatlas.Link{},
	}, nil)
	ac.On("GetClusters", mock.Anything, testProjectID).Return(
		[]mongodbatlas.Cluster{
			{
				GroupID: testProjectID,
				Name:    testClusterName,
			},
		},
		nil)
	return ac
}

func simpleAccessLogClient() accessLogClient {
	ac := testClientBase()
	ac.On("GetAccessLogs", mock.Anything, testProjectID, testClusterName, mock.Anything).Return(
		[]*mongodbatlas.AccessLogs{
			{
				GroupID:     testProjectID,
				Hostname:    "test-hostname.mongodb.net",
				ClusterName: testClusterName,
				IPAddress:   "192.168.1.1",
				AuthResult:  &authTrue,
				AuthSource:  "admin",
				LogLine:     "{\"t\":{\"$date\":\"2023-04-26T02:38:56.444+00:00\"},\"s\":\"I\",  \"c\":\"ACCESS\",   \"id\":20249,   \"ctx\":\"conn173\",\"msg\":\"Authentication failed\",\"attr\":{\"mechanism\":\"SCRAM-SHA-1\",\"speculative\":true,\"principalName\":\"mms-automation\",\"authenticationDatabase\":\"admin\",\"remote\":\"192.168.248.4:41052\",\"extraInfo\":{},\"error\":\"UserNotFound: User \\\"mms-automation@admin\\\" not found\"}}",
				Timestamp:   "Wed Apr 26 02:38:56 GMT 2023",
				Username:    "test",
			},
		},
		nil)
	return ac
}

func repeatedRequestAccessLogClient() accessLogClient {
	currentTime := time.Now().UTC()
	ac := testClientBase()
	ac.On("GetAccessLogs", mock.Anything, testProjectID, testClusterName, mock.Anything).Return(
		[]*mongodbatlas.AccessLogs{
			{
				GroupID:     testProjectID,
				Hostname:    "test-hostname.mongodb.net",
				ClusterName: testClusterName,
				IPAddress:   "192.168.1.1",
				AuthResult:  &authTrue,
				AuthSource:  "admin",
				LogLine:     fmt.Sprintf("{\"t\":{\"$date\":\"%s\"}}", currentTime.Add(1000*time.Millisecond).Format(time.RFC3339)),
				Username:    "test",
			},
			{
				GroupID:     testProjectID,
				Hostname:    "test-hostname.mongodb.net",
				ClusterName: testClusterName,
				IPAddress:   "192.168.1.1",
				AuthResult:  &authTrue,
				AuthSource:  "admin",
				LogLine:     fmt.Sprintf("{\"t\":{\"$date\":\"%s\"}}", currentTime.Add(900*time.Millisecond).Format(time.RFC3339)),
				Username:    "test",
			},
		},
		nil).Once()

	ac.On("GetAccessLogs", mock.Anything, testProjectID, testClusterName, mock.Anything).Return(
		[]*mongodbatlas.AccessLogs{
			{
				GroupID:     testProjectID,
				Hostname:    "test-hostname.mongodb.net",
				ClusterName: testClusterName,
				IPAddress:   "192.168.1.1",
				AuthResult:  &authTrue,
				AuthSource:  "admin",
				LogLine:     fmt.Sprintf("{\"t\":{\"$date\":\"%s\"}}", currentTime.Add(800*time.Millisecond).Format(time.RFC3339)),
				Username:    "test",
			},
		},
		nil).Once()
	return ac
}

func maxSizeButOldDataAccessLogsClient() accessLogClient {
	currentTime := time.Now().UTC()
	ac := testClientBase()
	ac.On("GetAccessLogs", mock.Anything, testProjectID, testClusterName, mock.Anything).Return(
		[]*mongodbatlas.AccessLogs{
			{
				GroupID:     testProjectID,
				Hostname:    "test-hostname.mongodb.net",
				ClusterName: testClusterName,
				IPAddress:   "192.168.1.1",
				AuthResult:  &authTrue,
				AuthSource:  "admin",
				LogLine:     fmt.Sprintf("{\"t\":{\"$date\":\"%s\"}}", currentTime.Add(500*time.Millisecond).Format(time.RFC3339)),
				Username:    "test",
			},
			{
				GroupID:     testProjectID,
				Hostname:    "test-hostname.mongodb.net",
				ClusterName: testClusterName,
				IPAddress:   "192.168.1.1",
				AuthResult:  &authTrue,
				AuthSource:  "admin",
				LogLine:     fmt.Sprintf("{\"t\":{\"$date\":\"%s\"}}", currentTime.Add(-100*time.Millisecond).Format(time.RFC3339)),
				Username:    "test",
			},
		},
		nil).Once()
	return ac
}

type mockAccessLogsClient struct {
	mock.Mock
}

func (mac *mockAccessLogsClient) GetProject(ctx context.Context, pID string) (*mongodbatlas.Project, error) {
	args := mac.Called(ctx, pID)
	return args.Get(0).(*mongodbatlas.Project), args.Error(1)
}

func (mac *mockAccessLogsClient) GetClusters(ctx context.Context, groupID string) ([]mongodbatlas.Cluster, error) {
	args := mac.Called(ctx, groupID)
	return args.Get(0).([]mongodbatlas.Cluster), args.Error(1)
}

func (mac *mockAccessLogsClient) GetAccessLogs(ctx context.Context, groupID string, clusterName string, opts *internal.GetAccessLogsOptions) (ret []*mongodbatlas.AccessLogs, err error) {
	args := mac.Called(ctx, groupID, clusterName, opts)
	return args.Get(0).([]*mongodbatlas.AccessLogs), args.Error(1)
}
