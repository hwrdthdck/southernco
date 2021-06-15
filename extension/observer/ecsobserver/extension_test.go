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

package ecsobserver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver/internal/ecsmock"
)

// Simply start and stop, the actual test logic is in sd_test.go until we implement the ListWatcher interface.
// In that case sd itself does not use timer and relies on caller to trigger List.
func TestExtensionStartStop(t *testing.T) {
	c := ecsmock.NewCluster()
	f := newTestTaskFetcher(t, c)
	ctx := context.WithValue(context.TODO(), ctxFetcherOverrideKey, f)
	ext, err := createExtension(ctx, component.ExtensionCreateSettings{Logger: zap.NewExample()}, createDefaultConfig())
	require.NoError(t, err)
	require.IsType(t, &ecsObserver{}, ext)
	require.NoError(t, ext.Start(context.TODO(), componenttest.NewNopHost()))
	require.NoError(t, ext.Shutdown(context.TODO()))
}
