// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/sqlquery"

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type SqlOpenerFunc func(driverName, dataSourceName string) (*sql.DB, error)

type DbProviderFunc func() (*sql.DB, error)

type ClientProviderFunc func(Db, string, *zap.Logger, TelemetryConfig) DbClient

type Scraper struct {
	Id                 component.ID
	Query              Query
	ScrapeCfg          scraperhelper.ScraperControllerSettings
	StartTime          pcommon.Timestamp
	ClientProviderFunc ClientProviderFunc
	DbProviderFunc     DbProviderFunc
	Logger             *zap.Logger
	Telemetry          TelemetryConfig
	Client             DbClient
	Db                 *sql.DB
}

var _ scraperhelper.Scraper = (*Scraper)(nil)

func (s *Scraper) ID() component.ID {
	return s.Id
}

func (s *Scraper) Start(context.Context, component.Host) error {
	var err error
	s.Db, err = s.DbProviderFunc()
	if err != nil {
		return fmt.Errorf("failed to open Db connection: %w", err)
	}
	s.Client = s.ClientProviderFunc(DbWrapper{s.Db}, s.Query.SQL, s.Logger, s.Telemetry)
	s.StartTime = pcommon.NewTimestampFromTime(time.Now())

	return nil
}

func (s *Scraper) Scrape(ctx context.Context) (pmetric.Metrics, error) {
	out := pmetric.NewMetrics()
	rows, err := s.Client.QueryRows(ctx)
	if err != nil {
		if errors.Is(err, errNullValueWarning) {
			s.Logger.Warn("problems encountered getting metric rows", zap.Error(err))
		} else {
			return out, fmt.Errorf("Scraper: %w", err)
		}
	}
	ts := pcommon.NewTimestampFromTime(time.Now())
	rms := out.ResourceMetrics()
	rm := rms.AppendEmpty()
	sms := rm.ScopeMetrics()
	sm := sms.AppendEmpty()
	ms := sm.Metrics()
	var errs error
	for _, metricCfg := range s.Query.Metrics {
		for i, row := range rows {
			if err = rowToMetric(row, metricCfg, ms.AppendEmpty(), s.StartTime, ts, s.ScrapeCfg); err != nil {
				err = fmt.Errorf("row %d: %w", i, err)
				errs = multierr.Append(errs, err)
			}
		}
	}
	if errs != nil {
		return out, scrapererror.NewPartialScrapeError(errs, len(multierr.Errors(errs)))
	}
	return out, nil
}

func (s *Scraper) Shutdown(_ context.Context) error {
	if s.Db != nil {
		return s.Db.Close()
	}
	return nil
}
