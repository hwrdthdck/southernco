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

package sqlqueryreceiver

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func TestCreateReceiver(t *testing.T) {
	f := createReceiverFunc(fakeDBConnect, mkFakeClient)
	ctx := context.Background()
	receiver, err := f(
		ctx,
		component.ReceiverCreateSettings{
			TelemetrySettings: component.TelemetrySettings{
				TracerProvider: trace.NewNoopTracerProvider(),
			},
		},
		createDefaultConfig(),
		consumertest.NewNop(),
	)
	require.NoError(t, err)
	err = receiver.Start(ctx, componenttest.NewNopHost())
	require.NoError(t, err)
	err = receiver.Shutdown(ctx)
	require.NoError(t, err)
}

func fakeDBConnect(string, string) (*sql.DB, error) {
	return nil, nil
}

func mkFakeClient(db *sql.DB, s string, logger *zap.Logger) dbClient {
	return &fakeDBClient{responses: [][]metricRow{{{"foo": "111"}}}}
}
