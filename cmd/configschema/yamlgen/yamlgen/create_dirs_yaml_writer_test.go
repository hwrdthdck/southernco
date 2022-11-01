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

package yamlgen

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/configschema"
)

func TestTreeYAMLWriter(t *testing.T) {
	tempDir := t.TempDir()
	w := &createDirsYAMLWriter{
		dirsCreated: map[string]struct{}{},
		baseDir:     tempDir,
	}
	err := w.write(configschema.CfgInfo{Group: "mygroup", Type: "mytype"}, []byte("hello"))
	require.NoError(t, err)
	file, err := os.Open(filepath.Join(tempDir, "mygroup", "mytype.yaml"))
	require.NoError(t, err)
	bytes, err := io.ReadAll(file)
	require.NoError(t, err)
	assert.EqualValues(t, "hello", bytes)
}
