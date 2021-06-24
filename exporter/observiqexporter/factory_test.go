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

package observiqexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configcheck"
	"go.uber.org/zap"
)

func TestNewFactory(t *testing.T) {
	fact := NewFactory()
	require.NotNil(t, fact, "failed to create new factory")
}

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	require.NotNil(t, cfg, "failed to create default config")
	require.NoError(t, configcheck.ValidateConfig(cfg))
}

func TestCreateLogsExporter(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	params := component.ExporterCreateSettings{Logger: zap.NewNop()}
	_, err := createLogsExporter(context.Background(), params, cfg)
	require.NoError(t, err)
}

func TestCreateLogsExporterNilConfig(t *testing.T) {
	params := component.ExporterCreateSettings{Logger: zap.NewNop()}
	_, err := createLogsExporter(context.Background(), params, nil)
	require.Error(t, err)
}
