// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsdreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateReceiver(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.NetAddr.Endpoint = "localhost:0" // Endpoint is required, not going to be used here.

	params := receivertest.NewNopCreateSettings()
	tReceiver, err := createMetricsReceiver(context.Background(), params, cfg, consumertest.NewNop())
	assert.NoError(t, err)
	assert.NotNil(t, tReceiver, "receiver creation failed")
}

func TestCreateMetricsReceiverWithNilConsumer(t *testing.T) {
	receiver, err := createMetricsReceiver(
		context.Background(),
		receivertest.NewNopCreateSettings(),
		createDefaultConfig(),
		nil,
	)

	assert.Error(t, err, "nil consumer")
	assert.Nil(t, receiver)
}
