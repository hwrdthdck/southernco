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

package prometheusreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"

import (
	"context"
	"errors"

	_ "github.com/prometheus/prometheus/discovery/install" // init() of this package registers service discovery impl.
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

// This file implements config for Prometheus receiver.

const (
	typeStr   = "prometheus"
	stability = component.StabilityLevelBeta
)

var errRenamingDisallowed = errors.New("metric renaming using metric_relabel_configs is disallowed")

// NewFactory creates a new Prometheus receiver factory.
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}

func createDefaultConfig() component.ReceiverConfig {
	return &Config{
		ReceiverSettings: config.NewReceiverSettings(component.NewID(typeStr)),
	}
}

func createMetricsReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	cfg component.ReceiverConfig,
	nextConsumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	return newPrometheusReceiver(set, cfg.(*Config), nextConsumer), nil
}
