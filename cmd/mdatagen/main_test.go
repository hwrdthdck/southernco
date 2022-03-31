// Copyright The OpenTelemetry Authors
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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	validMetadata = `
name: metricreceiver
attributes:
  cpu_type:
    value: type
    description: The type of CPU consumption
    enum:
    - user
    - io_wait
    - system
  host:
    description: The type of CPU consumption
metrics:
  system.cpu.time:
    enabled: true
    description: Total CPU seconds broken down by different states.
    extended_documentation: Additional information on CPU Time can be found [here](https://en.wikipedia.org/wiki/CPU_time).
    unit: s
    sum:
      aggregation: cumulative
      value_type: double
    attributes: [host, cpu_type]
`
)

func Test_runContents(t *testing.T) {
	type args struct {
		yml       string
		useExpGen bool
	}
	tests := []struct {
		name                   string
		args                   args
		expected_documentation string
		want                   string
		wantErr                string
	}{
		{
			name:                   "valid metadata",
			args:                   args{validMetadata, false},
			expected_documentation: "testdata/documentation_v1.md",
			want:                   "",
		},
		{
			name:                   "valid metadata v2",
			args:                   args{validMetadata, true},
			expected_documentation: "testdata/documentation_v2.md",
			want:                   "",
		},
		{
			name:    "invalid yaml",
			args:    args{"invalid", false},
			want:    "",
			wantErr: "cannot unmarshal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpdir, err := ioutil.TempDir("", "metadata-test-*")
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, os.RemoveAll(tmpdir))
			})

			metadataFile := filepath.Join(tmpdir, "metadata.yaml")
			require.NoError(t, ioutil.WriteFile(metadataFile, []byte(tt.args.yml), 0600))

			err = run(metadataFile, tt.args.useExpGen)

			if tt.wantErr != "" {
				require.Regexp(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)

				genFilePath := filepath.Join(tmpdir, "internal/metadata/generated_metrics.go")
				if tt.args.useExpGen {
					genFilePath = filepath.Join(tmpdir, "internal/metadata/generated_metrics_v2.go")
				}
				require.FileExists(t, genFilePath)

				actualDocumentation := filepath.Join(tmpdir, "documentation.md")
				require.FileExists(t, actualDocumentation)
				if tt.expected_documentation != "" {
					expectedFileBytes, err := ioutil.ReadFile(tt.expected_documentation)
					require.NoError(t, err)

					actualFileBytes, err := ioutil.ReadFile(actualDocumentation)
					require.NoError(t, err)

					require.Equal(t, expectedFileBytes, actualFileBytes)
				}
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		ymlPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no argument",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "no such file",
			args:    args{"/no/such/file"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.args.ymlPath, false); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
