// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jmxmetricsreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configerror"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/testbed/testbed"
	"go.uber.org/zap"
)

func TestWithInvalidConfig(t *testing.T) {
	f := NewFactory()
	assert.Equal(t, configmodels.Type("jmx_metrics"), f.Type())

	cfg := f.CreateDefaultConfig()
	require.NotNil(t, cfg)

	r, err := f.CreateMetricsReceiver(
		context.Background(),
		component.ReceiverCreateParams{Logger: zap.NewNop()},
		cfg,
		&testbed.MockMetricConsumer{},
	)
	require.Error(t, err)
	assert.Equal(t, "jmx_metrics missing required fields: [`service_url` `groovy_script`]", err.Error())
	require.Nil(t, r)
}

func TestWithValidConfig(t *testing.T) {
	f := NewFactory()
	assert.Equal(t, configmodels.Type("jmx_metrics"), f.Type())

	cfg := f.CreateDefaultConfig()
	cfg.(*config).ServiceURL = "myserviceurl"
	cfg.(*config).GroovyScript = "mygroovyscriptpath"

	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	consumer := &testbed.MockMetricConsumer{}
	r, err := f.CreateMetricsReceiver(context.Background(), params, cfg, consumer)
	require.NoError(t, err)
	require.NotNil(t, r)
	receiver := r.(*jmxMetricsReceiver)
	assert.Same(t, receiver.logger, params.Logger)
	assert.Same(t, receiver.config, cfg)
	assert.Same(t, receiver.consumer, consumer)
}

func TestTracerReceiverNotSupported(t *testing.T) {
	f := NewFactory()
	assert.Equal(t, configmodels.Type("jmx_metrics"), f.Type())

	r, err := f.CreateTraceReceiver(
		context.Background(),
		component.ReceiverCreateParams{Logger: zap.NewNop()},
		f.CreateDefaultConfig(),
		&testbed.MockTraceConsumer{},
	)

	require.Error(t, err)
	require.Equal(t, configerror.ErrDataTypeIsNotSupported, err)
	require.Nil(t, r)
}
