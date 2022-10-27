// Copyright  The OpenTelemetry Authors
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

package k8sobjectsreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer/consumertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

func TestDefaultConfig(t *testing.T) {
	t.Parallel()
	cfg := createDefaultConfig()
	rCfg, ok := cfg.(*Config)
	require.True(t, ok)

	assert.Equal(t, &Config{
		ReceiverSettings: config.NewReceiverSettings(config.NewComponentID("k8sobjects")),
		APIConfig: k8sconfig.APIConfig{
			AuthType: k8sconfig.AuthTypeServiceAccount,
		},
	}, rCfg)
}

func TestFactoryType(t *testing.T) {
	t.Parallel()
	assert.Equal(t, config.Type("k8sobjects"), NewFactory().Type())
}

func TestCreateReceiver(t *testing.T) {
	rCfg := createDefaultConfig().(*Config)

	// Fails with bad K8s Config.
	r, err := createLogsReceiver(
		context.Background(), componenttest.NewNopReceiverCreateSettings(),
		rCfg, consumertest.NewNop(),
	)
	assert.Error(t, err)
	assert.Nil(t, r)

	// Override for test.
	rCfg.makeDynamicClient = newMockDynamicClient().getMockDynamicClient

	r, err = createLogsReceiver(
		context.Background(),
		componenttest.NewNopReceiverCreateSettings(),
		rCfg, consumertest.NewNop(),
	)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}
