// Copyright  OpenTelemetry Authors
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

package awscontainerinsightreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/testbed/testbed"
	"go.uber.org/zap"
)

func TestNewFactory(t *testing.T) {
	c := NewFactory()
	assert.NotNil(t, c)
}

func TestCreateMetricsReceiver(t *testing.T) {
	metricsReceiver, _ := createMetricsReceiver(
		context.Background(),
		component.ReceiverCreateSettings{Logger: zap.NewNop()},
		createDefaultConfig(),
		&testbed.MockMetricConsumer{},
	)

	require.NotNil(t, metricsReceiver)
}
