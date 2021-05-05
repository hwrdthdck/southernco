// Copyright OpenTelemetry Authors
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

package azuremonitorexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.uber.org/zap"
)

// An inappropriate config
type badConfig struct {
	config.ExporterSettings `mapstructure:",squash"`
}

func TestCreateTracesExporterUsingSpecificTransportChannel(t *testing.T) {
	// mock transport channel creation
	f := factory{tChannel: &mockTransportChannel{}}
	ctx := context.Background()
	componentSettings := component.ComponentSettings{Logger: zap.NewNop()}
	exporter, err := f.createTracesExporter(ctx, componentSettings, createDefaultConfig())
	assert.NotNil(t, exporter)
	assert.Nil(t, err)
}

func TestCreateTracesExporterUsingDefaultTransportChannel(t *testing.T) {
	// We get the default transport channel creation, if we don't specify one during f creation
	f := factory{}
	assert.Nil(t, f.tChannel)
	ctx := context.Background()
	logger, _ := zap.NewDevelopment()
	componentSettings := component.ComponentSettings{Logger: logger}
	exporter, err := f.createTracesExporter(ctx, componentSettings, createDefaultConfig())
	assert.NotNil(t, exporter)
	assert.Nil(t, err)
	assert.NotNil(t, f.tChannel)
}

func TestCreateTracesExporterUsingBadConfig(t *testing.T) {
	// We get the default transport channel creation, if we don't specify one during factory creation
	f := factory{}
	assert.Nil(t, f.tChannel)
	ctx := context.Background()
	componentSettings := component.ComponentSettings{Logger: zap.NewNop()}

	badConfig := &badConfig{}

	exporter, err := f.createTracesExporter(ctx, componentSettings, badConfig)
	assert.Nil(t, exporter)
	assert.NotNil(t, err)
}
