// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/traceutil"
)

type logsExporter struct {
	client    *sql.DB
	insertSQL string

	logger *zap.Logger
	cfg    *Config
}

func newLogsExporter(logger *zap.Logger, cfg *Config) (*logsExporter, error) {
	client, err := newClickHouseConn(cfg)
	if err != nil {
		return nil, err
	}

	return &logsExporter{
		client:    client,
		insertSQL: renderInsertLogsSQL(cfg),
		logger:    logger,
		cfg:       cfg,
	}, nil
}

func (e *logsExporter) start(ctx context.Context, _ component.Host) error {
	if err := createDatabase(ctx, e.cfg); err != nil {
		return err
	}

	return createLogsTable(ctx, e.cfg, e.client)
}

// shutdown will shut down the exporter.
func (e *logsExporter) shutdown(_ context.Context) error {
	if e.client != nil {
		return e.client.Close()
	}
	return nil
}

func (e *logsExporter) pushLogsData(ctx context.Context, ld plog.Logs) error {
	start := time.Now()
	err := func() error {
		scope, err := e.client.Begin()
		if err != nil {
			return fmt.Errorf("Begin:%w", err)
		}

		batch, err := scope.Prepare(e.insertSQL)
		if err != nil {
			return fmt.Errorf("Prepare:%w", err)
		}

		var serviceName string
		var podName string
		var containerName string
		var region string
		var cloudProvider string
		var cell string

		resAttr := make(map[string]string)

		resourceLogs := ld.ResourceLogs()
		for i := 0; i < resourceLogs.Len(); i++ {
			logs := resourceLogs.At(i)
			res := logs.Resource()
<<<<<<< HEAD

			attrs := res.Attributes()
			attributesToMap(attrs, resAttr)

			attrs.Range(func(key string, value pcommon.Value) bool {
				switch key {
				case conventions.AttributeServiceName:
					serviceName = value.Str()
				case conventions.AttributeK8SPodName:
					podName = value.AsString()
				case conventions.AttributeK8SContainerName:
					containerName = value.AsString()
				// TODO use AttributeCloudRegion 'cloud.region'
				// https://github.com/ClickHouse/data-plane-application/issues/4155
				case "region":
					fallthrough
				case conventions.AttributeCloudRegion:
					region = value.AsString()
				case conventions.AttributeCloudProvider:
					cloudProvider = value.AsString()
				case "cell":
					cell = value.AsString()
				}
				return true
			})
=======
			resURL := logs.SchemaUrl()
			resAttr := attributesToMap(res.Attributes())
			if v, ok := res.Attributes().Get(conventions.AttributeServiceName); ok {
				serviceName = v.Str()
			}
>>>>>>> upstream/main
			for j := 0; j < logs.ScopeLogs().Len(); j++ {
				rs := logs.ScopeLogs().At(j).LogRecords()
				scopeURL := logs.ScopeLogs().At(j).SchemaUrl()
				scopeName := logs.ScopeLogs().At(j).Scope().Name()
				scopeVersion := logs.ScopeLogs().At(j).Scope().Version()
				scopeAttr := attributesToMap(logs.ScopeLogs().At(j).Scope().Attributes())
				for k := 0; k < rs.Len(); k++ {
					r := rs.At(k)

					logAttr := make(map[string]string, attrs.Len())
					attributesToMap(r.Attributes(), logAttr)

					_, err = batch.Exec(
						r.Timestamp().AsTime(),
						traceutil.TraceIDToHexOrEmptyString(r.TraceID()),
						traceutil.SpanIDToHexOrEmptyString(r.SpanID()),
						uint32(r.Flags()),
						r.SeverityText(),
						int32(r.SeverityNumber()),
						serviceName,
						r.Body().AsString(),
<<<<<<< HEAD
						podName,
						containerName,
						region,
						cloudProvider,
						cell,
=======
						resURL,
>>>>>>> upstream/main
						resAttr,
						scopeURL,
						scopeName,
						scopeVersion,
						scopeAttr,
						logAttr,
					)
					if err != nil {
						return fmt.Errorf("Append:%w", err)
					}
				}
			}

			// clear map for reuse
			for k := range resAttr {
				delete(resAttr, k)
			}
		}

		return scope.Commit()
	}()

	duration := time.Since(start)
	e.logger.Debug("insert logs", zap.Int("records", ld.LogRecordCount()),
		zap.String("cost", duration.String()))
	return err
}

func attributesToMap(attributes pcommon.Map, dest map[string]string) {
	attributes.Range(func(k string, v pcommon.Value) bool {
		dest[k] = v.AsString()
		return true
	})
}

const (
	// language=ClickHouse SQL
	createLogsTableSQL = `
CREATE TABLE IF NOT EXISTS %s (
     Timestamp DateTime64(9) CODEC(Delta, ZSTD(1)),
     TraceId String CODEC(ZSTD(1)),
     SpanId String CODEC(ZSTD(1)),
     TraceFlags UInt32 CODEC(ZSTD(1)),
     SeverityText LowCardinality(String) CODEC(ZSTD(1)),
     SeverityNumber Int32 CODEC(ZSTD(1)),
     ServiceName LowCardinality(String) CODEC(ZSTD(1)),
<<<<<<< HEAD
     Body LowCardinality(String) CODEC(ZSTD(1)),
     PodName LowCardinality(String),
     ContainerName LowCardinality(String),
     Region LowCardinality(String),
     CloudProvider LowCardinality(String),
     Cell LowCardinality(String),
=======
     Body String CODEC(ZSTD(1)),
     ResourceSchemaUrl String CODEC(ZSTD(1)),
>>>>>>> upstream/main
     ResourceAttributes Map(LowCardinality(String), String) CODEC(ZSTD(1)),
     ScopeSchemaUrl String CODEC(ZSTD(1)),
     ScopeName String CODEC(ZSTD(1)),
     ScopeVersion String CODEC(ZSTD(1)),
     ScopeAttributes Map(LowCardinality(String), String) CODEC(ZSTD(1)),
     LogAttributes Map(LowCardinality(String), String) CODEC(ZSTD(1)),
     INDEX idx_trace_id TraceId TYPE bloom_filter(0.001) GRANULARITY 1,
     INDEX idx_res_attr_key mapKeys(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
     INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
     INDEX idx_scope_attr_key mapKeys(ScopeAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
     INDEX idx_scope_attr_value mapValues(ScopeAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
     INDEX idx_log_attr_key mapKeys(LogAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
     INDEX idx_log_attr_value mapValues(LogAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
     INDEX idx_body Body TYPE tokenbf_v1(32768, 3, 0) GRANULARITY 1
) ENGINE MergeTree()
%s
PARTITION BY toYYYYMM(Timestamp)
ORDER BY (PodName, ContainerName, SeverityText, Timestamp)
SETTINGS index_granularity=8192, ttl_only_drop_parts = 1;
`

	// language=ClickHouse SQL
	insertLogsSQLTemplate = `INSERT INTO %s (
                        Timestamp,
                        TraceId,
                        SpanId,
                        TraceFlags,
                        SeverityText,
                        SeverityNumber,
                        ServiceName,
                        Body,
<<<<<<< HEAD
                        PodName,
						ContainerName,
						Region,
						CloudProvider,
						Cell,
=======
                        ResourceSchemaUrl,
>>>>>>> upstream/main
                        ResourceAttributes,
                        ScopeSchemaUrl,
                        ScopeName,
                        ScopeVersion,
                        ScopeAttributes,
                        LogAttributes
<<<<<<< HEAD
                        )`
	inlineinsertLogsSQLTemplate = `INSERT INTO %s SETTINGS async_insert=1, wait_for_async_insert=0 (
                        Timestamp,
                        TraceId,
                        SpanId,
                        TraceFlags,
                        SeverityText,
                        SeverityNumber,
                        ServiceName,
                        Body,
                        PodName,
						ContainerName,
						Region,
						CloudProvider,
						Cell,
                        ResourceAttributes,
                        LogAttributes
                        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
=======
                        ) VALUES (
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?,
                                  ?
                                  )`
>>>>>>> upstream/main
)

var driverName = "clickhouse" // for testing

// newClickHouseClient create a clickhouse client.
// used by metrics and traces:
func newClickHouseClient(cfg *Config) (*sql.DB, error) {
	db, err := cfg.buildDB(cfg.Database)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// used by logs:
func newClickHouseConn(cfg *Config) (*sql.DB, error) {
	endpoint := cfg.Endpoint

	if len(cfg.ConnectionParams) > 0 {
		values := make(url.Values, len(cfg.ConnectionParams))
		for k, v := range cfg.ConnectionParams {
			values.Add(k, v)
		}

		if !strings.Contains(endpoint, "?") {
			endpoint += "?"
		} else if !strings.HasSuffix(endpoint, "&") {
			endpoint += "&"
		}

		endpoint += values.Encode()
	}

	opts, err := clickhouse.ParseDSN(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to parse endpoint: %w", err)
	}

	opts.Auth = clickhouse.Auth{
		Database: cfg.Database,
		Username: cfg.Username,
		Password: cfg.Password,
	}

	// can return a "bad" connection if misconfigured, we won't know
	// until a Ping, Exec, etc.. is done
	return clickhouse.OpenDB(opts), nil
}

func createDatabase(ctx context.Context, cfg *Config) error {
	// use default database to create new database
	if cfg.Database == defaultDatabase {
		return nil
	}

	db, err := cfg.buildDB(defaultDatabase)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.Database)
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("create database:%w", err)
	}
	return nil
}

func createLogsTable(ctx context.Context, cfg *Config, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, renderCreateLogsTableSQL(cfg)); err != nil {
		return fmt.Errorf("exec create logs table sql: %w", err)
	}
	return nil
}

func renderCreateLogsTableSQL(cfg *Config) string {
	var ttlExpr string
	if cfg.TTLDays > 0 {
		ttlExpr = fmt.Sprintf(`TTL toDateTime(Timestamp) + toIntervalDay(%d)`, cfg.TTLDays)
	}
	return fmt.Sprintf(createLogsTableSQL, cfg.LogsTableName, ttlExpr)
}

func renderInsertLogsSQL(cfg *Config) string {
	if strings.HasPrefix(cfg.Endpoint, "tcp") && cfg.ConnectionParams["async_insert"] == "1" {
		return fmt.Sprintf(inlineinsertLogsSQLTemplate, cfg.LogsTableName)
	}
	return fmt.Sprintf(insertLogsSQLTemplate, cfg.LogsTableName)
}
