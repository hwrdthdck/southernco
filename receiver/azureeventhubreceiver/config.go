// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"errors"
	"fmt"
	"time"

	"github.com/Azure/azure-amqp-common-go/v4/conn"
	"go.opentelemetry.io/collector/component"
)

type logFormat string

const (
	defaultLogFormat logFormat = ""
	rawLogFormat     logFormat = "raw"
	azureLogFormat   logFormat = "azure"
)

var (
	validFormats         = []logFormat{defaultLogFormat, rawLogFormat, azureLogFormat}
	errMissingConnection = errors.New("missing connection")
)

type Config struct {
	Connection               string        `mapstructure:"connection"`
	Partition                string        `mapstructure:"partition"`
	Offset                   string        `mapstructure:"offset"`
	StorageID                *component.ID `mapstructure:"storage"`
	Format                   string        `mapstructure:"format"`
	ConsumerGroup            string        `mapstructure:"group"`
	ApplySemanticConventions bool          `mapstructure:"apply_semantic_conventions"`
	TimeFormat               TimeFormat    `mapstructure:"time_format"`
	TimeOffset               TimeOffset    `mapstructure:"time_offset"`
}

type TimeFormat struct {
	Logs    []string `mapstructure:"logs"`
	Metrics []string `mapstructure:"metrics"`
	Traces  []string `mapstructure:"traces"`
}

type TimeOffset struct {
	Logs    time.Duration `mapstructure:"logs"`
	Metrics time.Duration `mapstructure:"metrics"`
	Traces  time.Duration `mapstructure:"traces"`
}

func isValidFormat(format string) bool {
	for _, validFormat := range validFormats {
		if logFormat(format) == validFormat {
			return true
		}
	}
	return false
}

// Validate config
func (config *Config) Validate() error {
	if config.Connection == "" {
		return errMissingConnection
	}
	if _, err := conn.ParsedConnectionFromStr(config.Connection); err != nil {
		return err
	}
	if !isValidFormat(config.Format) {
		return fmt.Errorf("invalid format; must be one of %#v", validFormats)
	}
	return nil
}
