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
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor"
	hostInfo "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/k8sapiserver"
)

// Factory for awscontainerinsightreceiver
const (
	// Key to invoke this receiver
	typeStr = "awscontainerinsightreceiver"

	// Default collection interval. Every 60s the receiver will collect metrics
	defaultCollectionInterval = 60 * time.Second

	// Default container orchestrator service is aws eks
	defaultContainerOrchestrator = "eks"
)

// NewFactory creates a factory for AWS container insight receiver
func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithMetrics(createMetricsReceiver))
}

// createDefaultConfig returns a default config for the receiver.
func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings:      config.NewReceiverSettings(config.NewID(typeStr)),
		CollectionInterval:    defaultCollectionInterval,
		ContainerOrchestrator: defaultContainerOrchestrator,
	}
}

// CreateMetricsReceiver creates an AWS Container Insight receiver.
func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateParams,
	baseCfg config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {

	rCfg := baseCfg.(*Config)
	logger := params.Logger
	hostInfo, err := hostInfo.NewInfo(rCfg.CollectionInterval, logger)
	// TODO: I will need to change the code here to let cadvisor and k8sapiserver return err as well
	if err != nil {
		logger.Warn("failed to initialize hostInfo", zap.Error(err))
	}
	cadvisor := cadvisor.New(rCfg.ContainerOrchestrator, hostInfo, logger)
	k8sapiserver := k8sapiserver.New(hostInfo, logger)
	return New(logger, rCfg, consumer, cadvisor, k8sapiserver)
}
