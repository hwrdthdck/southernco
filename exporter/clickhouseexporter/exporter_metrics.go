// Copyright  The OpenTelemetry Authors
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

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type metricsExporter struct {
	client *sql.DB

	logger *zap.Logger
	cfg    *Config
}

func newMetricsExporter(logger *zap.Logger, cfg *Config) (*metricsExporter, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if err := createDatabase(cfg); err != nil {
		return nil, err
	}
	client, err := newClickhouseClient(cfg)
	if err != nil {
		return nil, err
	}

	if err = internal.CreateMetricsTable(cfg.MetricsTableName, cfg.TTLDays, client); err != nil {
		return nil, err
	}

	return &metricsExporter{
		client: client,
		logger: logger,
		cfg:    cfg,
	}, nil
}

// Shutdown will shutdown the exporter.
func (e *metricsExporter) shutdown(ctx context.Context) error {
	if e.client != nil {
		return e.client.Close()
	}
	return nil
}

func (e *metricsExporter) pushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	metricsMap := internal.CreateMetricsModel(e.cfg.MetricsTableName)
	for i := 0; i < md.ResourceMetrics().Len(); i++ {
		metaData := internal.MetricsMetaData{}
		metrics := md.ResourceMetrics().At(i)
		res := metrics.Resource()
		metaData.ResAttr = attributesToMap(res.Attributes())
		metaData.ResURL = metrics.SchemaUrl()
		for j := 0; j < metrics.ScopeMetrics().Len(); j++ {
			rs := metrics.ScopeMetrics().At(j).Metrics()
			metaData.ScopeURL = metrics.ScopeMetrics().At(j).SchemaUrl()
			metaData.ScopeInstr = metrics.ScopeMetrics().At(j).Scope()
			for k := 0; k < rs.Len(); k++ {
				r := rs.At(k)
				var errs []error
				switch r.Type() {
				case pmetric.MetricTypeGauge:
					errs = append(errs, metricsMap[pmetric.MetricTypeGauge].Add(r.Gauge(), &metaData, r.Name(), r.Description(), r.Unit()))
				case pmetric.MetricTypeSum:
					errs = append(errs, metricsMap[pmetric.MetricTypeSum].Add(r.Sum(), &metaData, r.Name(), r.Description(), r.Unit()))
				case pmetric.MetricTypeHistogram:
					errs = append(errs, metricsMap[pmetric.MetricTypeHistogram].Add(r.Histogram(), &metaData, r.Name(), r.Description(), r.Unit()))
				case pmetric.MetricTypeExponentialHistogram:
					errs = append(errs, metricsMap[pmetric.MetricTypeExponentialHistogram].Add(r.ExponentialHistogram(), &metaData, r.Name(), r.Description(), r.Unit()))
				case pmetric.MetricTypeSummary:
					errs = append(errs, metricsMap[pmetric.MetricTypeSummary].Add(r.Summary(), &metaData, r.Name(), r.Description(), r.Unit()))
				default:
					return fmt.Errorf("unsupported metrics type")
				}
				if multierr.Combine(errs...) != nil {
					return multierr.Combine(errs...)
				}
			}
		}
	}
	// batch insert https://clickhouse.com/docs/en/about-us/performance/#performance-when-inserting-data
	if err := internal.InsertMetrics(ctx, e.client, metricsMap, e.logger); err != nil {
		return err
	}
	return nil
}
