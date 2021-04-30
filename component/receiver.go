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

package component

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

// Receiver allows the collector to receive metrics, traces and logs.
type Receiver interface {
	Component
}

// A TracesReceiver receives traces.
// Its purpose is to translate data from any format to the collector's internal trace format.
// TracesReceiver feeds a consumer.Traces with data.
//
// For example it could be Zipkin data source which translates Zipkin spans into pdata.Traces.
type TracesReceiver interface {
	Receiver
}

// A MetricsReceiver receives metrics.
// Its purpose is to translate data from any format to the collector's internal metrics format.
// MetricsReceiver feeds a consumer.Metrics with data.
//
// For example it could be Prometheus data source which translates Prometheus metrics into pdata.Metrics.
type MetricsReceiver interface {
	Receiver
}

// A LogsReceiver receives logs.
// Its purpose is to translate data from any format to the collector's internal logs data format.
// LogsReceiver feeds a consumer.Logs with data.
//
// For example a LogsReceiver can read syslogs and convert them into pdata.Logs.
type LogsReceiver interface {
	Receiver
}

// ReceiverCreateParams configures Receiver creators.
type ReceiverCreateParams struct {
	// Logger that the factory can use during creation and can pass to the created
	// component to be used later as well.
	Logger *zap.Logger

	// BuildInfo can be used by components for informational purposes
	BuildInfo BuildInfo
}

// ReceiverFactory can create TracesReceiver and MetricsReceiver. This is the
// new preferred factory type to create receivers.
type ReceiverFactory interface {
	Factory

	// CreateDefaultConfig creates the default configuration for the Receiver.
	// This method can be called multiple times depending on the pipeline
	// configuration and should not cause side-effects that prevent the creation
	// of multiple instances of the Receiver.
	// The object returned by this method needs to pass the checks implemented by
	// 'configcheck.ValidateConfig'. It is recommended to have such check in the
	// tests of any implementation of the Factory interface.
	CreateDefaultConfig() config.Receiver

	// CreateTracesReceiver creates a trace receiver based on this config.
	// If the receiver type does not support tracing or if the config is not valid
	// error will be returned instead.
	CreateTracesReceiver(ctx context.Context, params ReceiverCreateParams,
		cfg config.Receiver, nextConsumer consumer.Traces) (TracesReceiver, error)

	// CreateMetricsReceiver creates a metrics receiver based on this config.
	// If the receiver type does not support metrics or if the config is not valid
	// error will be returned instead.
	CreateMetricsReceiver(ctx context.Context, params ReceiverCreateParams,
		cfg config.Receiver, nextConsumer consumer.Metrics) (MetricsReceiver, error)

	// CreateLogsReceiver creates a log receiver based on this config.
	// If the receiver type does not support the data type or if the config is not valid
	// error will be returned instead.
	CreateLogsReceiver(ctx context.Context, params ReceiverCreateParams,
		cfg config.Receiver, nextConsumer consumer.Logs) (LogsReceiver, error)
}
