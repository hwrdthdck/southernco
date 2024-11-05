// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cgroupruntimeextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/cgroupruntimeextension"

import (
	"context"
	"fmt"
	"runtime/debug"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/KimMachineGun/automemlimit/memlimit"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/cgroupruntimeextension/internal/metadata"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		GoMaxProcs: GoMaxProcsConfig{
			Enabled: true,
		},
		GoMemLimit: GoMemLimitConfig{
			Enabled: true,
			// By default, it sets `GOMEMLIMIT` to 90% of cgroup's memory limit.
			Ratio: 0.9,
		},
	}
}

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	cgroupConfig := cfg.(*Config)
	return newCgroupRuntime(cgroupConfig, set.Logger,
		func() (undoFunc, error) {
			undo, err := maxprocs.Set(maxprocs.Logger(func(str string, params ...interface{}) {
				set.Logger.Debug(fmt.Sprintf(str, params))
			}))
			return undoFunc(undo), err
		},
		func(ratio float64) (undoFunc, error) {
			initial, err := memlimit.SetGoMemLimitWithOpts(memlimit.WithRatio(ratio))
			return func() { debug.SetMemoryLimit(initial) }, err
		}), nil
}
