// Copyright 2020, OpenTelemetry Authors
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

package cassandraexporter

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/traceutil"
	"go.opentelemetry.io/collector/pdata/ptrace"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"
	"time"
)

type tracesExporter struct {
	client *gocql.Session
	logger *zap.Logger
	cfg    *Config
}

func newTracesExporter(logger *zap.Logger, cfg *Config) (*tracesExporter, error) {
	initializeErr := initializeTraceKernel(cfg)
	if initializeErr != nil {
		return nil, initializeErr
	}

	cluster := gocql.NewCluster(cfg.DSN)
	session, err := cluster.CreateSession()
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.Quorum

	if err != nil {
		return nil, err
	}

	return &tracesExporter{logger: logger, client: session, cfg: cfg}, nil
}

func initializeTraceKernel(cfg *Config) error {
	ctx := context.Background()
	cluster := gocql.NewCluster(cfg.DSN)
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	session.Query(parseCreateDatabaseSql(cfg)).WithContext(ctx).Exec()
	session.Query(parseCreateLinksTypeSql(cfg)).WithContext(ctx).Exec()
	session.Query(parseCreateEventsTypeSql(cfg)).WithContext(ctx).Exec()
	session.Query(parseCreateSpanTableSql(cfg)).WithContext(ctx).Exec()

	defer session.Close()

	return nil
}

func parseCreateSpanTableSql(cfg *Config) string {
	return fmt.Sprintf(createSpanTableSQL, cfg.Keyspace, cfg.TraceTable, cfg.Compression.Algorithm)
}

func parseCreateEventsTypeSql(cfg *Config) string {
	return fmt.Sprintf(createEventTypeSql, cfg.Keyspace)
}

func parseCreateLinksTypeSql(cfg *Config) string {
	return fmt.Sprintf(createLinksTypeSql, cfg.Keyspace)
}

func parseCreateDatabaseSql(cfg *Config) string {
	return fmt.Sprintf(createDatabaseSQL, cfg.Keyspace, cfg.Replication.Class, cfg.Replication.ReplicationFactor)
}

func (e *tracesExporter) Shutdown(_ context.Context) error {
	if e.client != nil {
		e.client.Close()
	}

	return nil
}

func (e *tracesExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	start := time.Now()

	for i := 0; i < td.ResourceSpans().Len(); i++ {
		spans := td.ResourceSpans().At(i)
		res := spans.Resource()
		resAttr := attributesToMap(res.Attributes())
		var serviceName string
		if v, ok := res.Attributes().Get(conventions.AttributeServiceName); ok {
			serviceName = v.Str()
		}
		for j := 0; j < spans.ScopeSpans().Len(); j++ {
			rs := spans.ScopeSpans().At(j).Spans()
			for k := 0; k < rs.Len(); k++ {
				r := rs.At(k)
				spanAttr := attributesToMap(r.Attributes())
				status := r.Status()

				e.client.Query(fmt.Sprintf(insertSpanSQL, e.cfg.Keyspace, e.cfg.TraceTable), r.StartTimestamp().AsTime(),
					traceutil.TraceIDToHexOrEmptyString(r.TraceID()),
					traceutil.SpanIDToHexOrEmptyString(r.SpanID()),
					traceutil.SpanIDToHexOrEmptyString(r.ParentSpanID()),
					r.TraceState().AsRaw(),
					r.Name(),
					traceutil.SpanKindStr(r.Kind()),
					serviceName,
					resAttr,
					spanAttr,
					r.EndTimestamp().AsTime().Sub(r.StartTimestamp().AsTime()).Nanoseconds(),
					traceutil.StatusCodeStr(status.Code()),
					status.Message(),
				).WithContext(ctx).Exec()
			}
		}
	}

	duration := time.Since(start)
	e.logger.Debug("insert traces", zap.Int("records", td.SpanCount()),
		zap.String("cost", duration.String()))
	return nil
}
