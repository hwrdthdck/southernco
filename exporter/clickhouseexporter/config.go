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

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/multierr"
)

// Config defines configuration for Elastic exporter.
type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"`
	exporterhelper.RetrySettings   `mapstructure:"retry_on_failure"`
	// QueueSettings is a subset of exporterhelper.QueueSettings,
	// because only QueueSize is user-settable.
	QueueSettings QueueSettings `mapstructure:"sending_queue"`

	// Endpoint is the clickhouse endpoint.
	Endpoint string `mapstructure:"endpoint"`
	// Database is the database name to export.
	Database string `mapstructure:"database"`
	// Username is the authentication username.
	Username string `mapstructure:"username"`
	// Username is the authentication password.
	Password string `mapstructure:"password"`
	// LogsTableName is the table name for logs. default is `otel_logs`.
	LogsTableName string `mapstructure:"logs_table_name"`
	// TracesTableName is the table name for logs. default is `otel_traces`.
	TracesTableName string `mapstructure:"traces_table_name"`
	// MetricsTableName is the table name for metrics. default is `otel_metrics`.
	MetricsTableName string `mapstructure:"metrics_table_name"`
	// TTLDays is The data time-to-live in days, 0 means no ttl.
	TTLDays uint `mapstructure:"ttl_days"`
}

// QueueSettings is a subset of exporterhelper.QueueSettings.
type QueueSettings struct {
	// QueueSize set the length of the sending queue
	QueueSize int `mapstructure:"queue_size"`
}

const defaultDatabase = "default"

var (
	errConfigNoHost     = errors.New("host must be specified")
	errConfigInvalidDSN = errors.New("DSN is invalid")
)

// Validate the clickhouse server configuration.
func (cfg *Config) Validate() (err error) {
	if cfg.Endpoint == "" {
		err = multierr.Append(err, errConfigNoHost)
	}
	_, e := cfg.buildDB(cfg.Database)
	if e != nil {
		err = multierr.Append(err, e)
	}
	return err
}

func (cfg *Config) enforcedQueueSettings() exporterhelper.QueueSettings {
	return exporterhelper.QueueSettings{
		Enabled:      true,
		NumConsumers: 1,
		QueueSize:    cfg.QueueSettings.QueueSize,
	}
}

func (cfg *Config) buildDSN(database string) (string, error) {
	parsedDSN, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errConfigInvalidDSN, err)
	}

	dsnCopy := *parsedDSN

	queryParams := dsnCopy.Query()

	// Enable TLS if scheme is https. This flag is necessary to support https connections.
	if dsnCopy.Scheme == "https" {
		queryParams.Set("secure", "true")
	}

	// Override database if specified in config.
	if cfg.Database != "" {
		dsnCopy.Path = cfg.Database
	} else if database == "" && cfg.Database == "" {
		// Use default database if not specified.
		dsnCopy.Path = defaultDatabase
	}

	// Override username and password if specified in config.
	if cfg.Username != "" {
		dsnCopy.User = url.UserPassword(cfg.Username, cfg.Password)
	}

	dsnCopy.RawQuery = queryParams.Encode()

	return dsnCopy.String(), nil
}

func (cfg *Config) buildDB(database string) (*sql.DB, error) {
	dsn, err := cfg.buildDSN(database)
	if err != nil {
		return nil, err
	}

	// ClickHouse sql driver will read settings from the DSN string.
	// It also ensures good defaults.
	// See https://github.com/ClickHouse/clickhouse-go/blob/08b27884b899f587eb5c509769cd2bdf74a9e2a1/clickhouse_std.go#L189
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
