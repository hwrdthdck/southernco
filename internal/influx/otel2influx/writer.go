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

package otel2influx // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/influx/otel2influx"

import (
	"context"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/influx/common"
)

type InfluxWriter interface {
	NewBatch() InfluxWriterBatch
}

type InfluxWriterBatch interface {
	WritePoint(ctx context.Context, measurement string, tags map[string]string, fields map[string]interface{}, ts time.Time, vType common.InfluxMetricValueType) error
	FlushBatch(ctx context.Context) error
}
