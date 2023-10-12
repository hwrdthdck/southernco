// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/zipkinencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	extensionName  = "zipkinencoding"
	stabilityLevel = component.StabilityLevelDevelopment
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		extensionName,
		createDefaultConfig,
		createExtension,
		stabilityLevel,
	)
}

func createExtension(_ context.Context, _ extension.CreateSettings, config component.Config) (extension.Extension, error) {
	return newExtension(config.(*Config))
}

func createDefaultConfig() component.Config {
	return &Config{Protocol: zipkinProtobufEncoding, Version: v2}
}
