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

package logs

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap/zaptest"
)

func TestSubmitLogs(t *testing.T) {
	logger := zaptest.NewLogger(t)
	var counter int

	tests := []struct {
		name    string
		payload []datadogV2.HTTPLogItem
		handler func(ctx context.Context, s *Sender, batch []datadogV2.HTTPLogItem, tags string) error
	}{
		{
			name: "batches same tags",
			payload: []datadogV2.HTTPLogItem{{
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag1:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}, {
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag1:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}},
			handler: func(ctx context.Context, s *Sender, batch []datadogV2.HTTPLogItem, tags string) error {
				switch counter {
				case 0:
					assert.Equal(t, tags, "tag1:true")
					assert.Len(t, batch, 2)
				default:
					t.Fail()
				}
				counter++
				return nil
			},
		},
		{
			name: "does not batch same tags",
			payload: []datadogV2.HTTPLogItem{{
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag1:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}, {
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag2:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}},
			handler: func(ctx context.Context, s *Sender, batch []datadogV2.HTTPLogItem, tags string) error {
				switch counter {
				case 0:
					assert.Equal(t, tags, "tag1:true")
					assert.Len(t, batch, 1)
				case 1:
					assert.Equal(t, tags, "tag2:true")
					assert.Len(t, batch, 1)
				default:
					t.Fail()
				}
				counter++
				return nil
			},
		},
		{
			name: "does two batches",
			payload: []datadogV2.HTTPLogItem{{
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag1:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}, {
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag1:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}, {
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag2:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}, {
				Ddsource:             nil,
				Ddtags:               datadog.PtrString("tag2:true"),
				Hostname:             nil,
				Message:              "",
				Service:              nil,
				UnparsedObject:       nil,
				AdditionalProperties: nil,
			}},
			handler: func(ctx context.Context, s *Sender, batch []datadogV2.HTTPLogItem, tags string) error {
				switch counter {
				case 0:
					assert.Equal(t, tags, "tag1:true")
					assert.Len(t, batch, 2)
				case 1:
					assert.Equal(t, tags, "tag2:true")
					assert.Len(t, batch, 2)
				default:
					t.Fail()
				}
				counter++
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter = 0
			s := NewSender("", logger, exporterhelper.TimeoutSettings{Timeout: 60}, false, false, "", tt.handler)
			_ = s.SubmitLogs(context.Background(), tt.payload)
		})
	}
}
