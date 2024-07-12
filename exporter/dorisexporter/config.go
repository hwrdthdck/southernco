// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"errors"
	"regexp"
	"time"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"`
	configretry.BackOffConfig      `mapstructure:"retry_on_failure"`
	exporterhelper.QueueSettings   `mapstructure:"sending_queue"`

	// TableNames is the table name for logs, traces and metrics.
	Table `mapstructure:"table"`

	// Endpoint is the http stream load address.
	Endpoint string `mapstructure:"endpoint"`
	// Database is the database name.
	Database string `mapstructure:"database"`
	// Username is the authentication username.
	Username string `mapstructure:"username"`
	// Password is the authentication password.
	Password configopaque.String `mapstructure:"password"`
	// CreateSchema is whether databases and tables are created automatically.
	CreateSchema bool `mapstructure:"create_schema"`
	// MySQLEndpoint is the mysql protocol address to create the schema; ignored if create_schema is false.
	MySQLEndpoint string `mapstructure:"mysql_endpoint"`
	// Data older than these days will be deleted; ignored if create_schema is false. If set to 0, historical data will not be deleted.
	HistoryDays int32 `mapstructure:"history_days"`
	// The number of days in the history partition that was created when the table was created. ignored if create_schema is false.
	// If history_days is not 0, create_history_days needs to be less than or equal to history_days.
	CreateHistoryDays int32 `mapstructure:"create_history_days"`
	// Timezone is the timezone of the doris.
	TimeZone string `mapstructure:"timezone"`
}

type Table struct {
	// Logs is the table name for logs.
	Logs string `mapstructure:"logs"`
	// Traces is the table name for traces.
	Traces string `mapstructure:"traces"`
	// Metrics is the table name for metrics.
	Metrics string `mapstructure:"metrics"`
}

func (cfg *Config) Validate() (err error) {
	if cfg.Endpoint == "" {
		err = errors.Join(err, errors.New("endpoint must be specified"))
	}
	if cfg.CreateSchema {
		if cfg.MySQLEndpoint == "" {
			err = errors.Join(err, errors.New("mysql_endpoint must be specified"))
		}

		if cfg.HistoryDays < 0 {
			err = errors.Join(err, errors.New("history_days must be greater than or equal to 0"))
		}

		if cfg.CreateHistoryDays < 0 {
			err = errors.Join(err, errors.New("create_history_days must be greater than or equal to 0"))
		}

		if cfg.HistoryDays > 0 && cfg.CreateHistoryDays > cfg.HistoryDays {
			err = errors.Join(err, errors.New("create_history_days must be less than or equal to history_days"))
		}
	}

	// Preventing SQL Injection Attacks
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !re.MatchString(cfg.Database) {
		err = errors.Join(err, errors.New("database name must be alphanumeric and underscore"))
	}
	if !re.MatchString(cfg.Table.Logs) {
		err = errors.Join(err, errors.New("logs table name must be alphanumeric and underscore"))
	}
	if !re.MatchString(cfg.Table.Traces) {
		err = errors.Join(err, errors.New("traces table name must be alphanumeric and underscore"))
	}
	if !re.MatchString(cfg.Table.Metrics) {
		err = errors.Join(err, errors.New("metrics table name must be alphanumeric and underscore"))
	}

	return err
}

const (
	defaultStart = -2147483648 // IntMin
)

func (cfg *Config) start() int32 {
	if cfg.HistoryDays == 0 {
		return defaultStart
	}
	return -cfg.HistoryDays
}

func (cfg *Config) timeZone() (*time.Location, error) {
	return time.LoadLocation(cfg.TimeZone)
}
