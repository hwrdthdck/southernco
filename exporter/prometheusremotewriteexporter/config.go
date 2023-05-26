// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"
)

// Config defines configuration for Remote Write exporter.
type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	exporterhelper.RetrySettings   `mapstructure:"retry_on_failure"`

	// prefix attached to each exported metric name
	// See: https://prometheus.io/docs/practices/naming/#metric-names
	Namespace string `mapstructure:"namespace"`

	// QueueConfig allows users to fine tune the queues
	// that handle outgoing requests.
	RemoteWriteQueue RemoteWriteQueue `mapstructure:"remote_write_queue"`

	// ExternalLabels defines a map of label keys and values that are allowed to start with reserved prefix "__"
	ExternalLabels map[string]string `mapstructure:"external_labels"`

	HTTPClientSettings confighttp.HTTPClientSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.

	// ResourceToTelemetrySettings is the option for converting resource attributes to telemetry attributes.
	// "Enabled" - A boolean field to enable/disable this option. Default is `false`.
	// If enabled, all the resource attributes will be converted to metric labels by default.
	ResourceToTelemetrySettings resourcetotelemetry.Settings `mapstructure:"resource_to_telemetry_conversion"`
	WAL                         *WALConfig                   `mapstructure:"wal"`

	// TargetInfo allows customizing the target_info metric
	TargetInfo *TargetInfo `mapstructure:"target_info,omitempty"`

	// CreatedMetric allows customizing creation of _created metrics
	CreatedMetric *CreatedMetric `mapstructure:"export_created_metric,omitempty"`
}

type CreatedMetric struct {
	// Enabled if true the _created metrics could be exported
	Enabled bool `mapstructure:"enabled"`
}

type TargetInfo struct {
	// Enabled if false the target_info metric is not generated by the exporter
	Enabled bool `mapstructure:"enabled"`
}

// RemoteWriteQueue allows to configure the remote write queue.
type RemoteWriteQueue struct {
	// Enabled if false the queue is not enabled, the export requests
	// are executed synchronously.
	Enabled bool `mapstructure:"enabled"`

	// QueueSize is the maximum number of OTLP metric batches allowed
	// in the queue at a given time. Ignored if Enabled is false.
	QueueSize int `mapstructure:"queue_size"`

	// NumWorkers configures the number of workers used by
	// the collector to fan out remote write requests.
	NumConsumers int `mapstructure:"num_consumers"`
}

// TODO(jbd): Add capacity, max_samples_per_send to QueueConfig.

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error {
	if cfg.RemoteWriteQueue.QueueSize < 0 {
		return fmt.Errorf("remote write queue size can't be negative")
	}

	if cfg.RemoteWriteQueue.Enabled && cfg.RemoteWriteQueue.QueueSize == 0 {
		return fmt.Errorf("a 0 size queue will drop all the data")
	}

	if cfg.RemoteWriteQueue.NumConsumers < 0 {
		return fmt.Errorf("remote write consumer number can't be negative")
	}

	if cfg.TargetInfo == nil {
		cfg.TargetInfo = &TargetInfo{
			Enabled: true,
		}
	}
	if cfg.CreatedMetric == nil {
		cfg.CreatedMetric = &CreatedMetric{
			Enabled: false,
		}
	}
	return nil
}
