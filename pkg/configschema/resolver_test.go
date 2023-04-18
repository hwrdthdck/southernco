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

package configschema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const gcpCollectorPath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/collector"

func TestTokensToPartialPath(t *testing.T) {
	path, err := requireTokensToPartialPath([]string{gcpCollectorPath, "42"})
	require.NoError(t, err)
	assert.Equal(t, "github.com/!google!cloud!platform/opentelemetry-operations-go/exporter/collector@42", path)
}

func TestDirResolver_PackagePathToProjectPath(t *testing.T) {
	dr := dirResolver{}
	fmt.Printf(": %v\n", dr)
	dr.packagePathToProjectPath("")
}
