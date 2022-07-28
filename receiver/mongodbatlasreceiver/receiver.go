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

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

type receiver struct {
	log             *zap.Logger
	cfg             *Config
	client          *internal.MongoDBAtlasClient
	lastRun         time.Time
	mb              *metadata.MetricsBuilder
	consumer        consumer.Logs
	stopperChanList []chan struct{}
}

type timeconstraints struct {
	start      string
	end        string
	resolution string
}

func newMongoDBAtlasReciever(settings component.ReceiverCreateSettings, cfg *Config) (*receiver, error) {
	client, err := internal.NewMongoDBAtlasClient(cfg.PublicKey, cfg.PrivateKey, cfg.RetrySettings, settings.Logger)
	if err != nil {
		return nil, err
	}
	recv := &receiver{log: settings.Logger, cfg: cfg, client: client, mb: metadata.NewMetricsBuilder(cfg.Metrics, settings.BuildInfo)}
	return recv, nil
}

func newMongoDBAtlasScraper(recv *receiver) (scraperhelper.Scraper, error) {
	return scraperhelper.NewScraper(typeStr, recv.scrape, scraperhelper.WithShutdown(recv.shutdown))
}

func (s *receiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	now := time.Now()
	if err := s.poll(ctx, s.timeConstraints(now)); err != nil {
		return pmetric.Metrics{}, err
	}
	s.lastRun = now
	return s.mb.Emit(), nil
}

func (s *receiver) timeConstraints(now time.Time) timeconstraints {
	var start time.Time
	if s.lastRun.IsZero() {
		start = now.Add(s.cfg.CollectionInterval * -1)
	} else {
		start = s.lastRun
	}
	return timeconstraints{
		start.UTC().Format(time.RFC3339),
		now.UTC().Format(time.RFC3339),
		s.cfg.Granularity,
	}
}

func (s *receiver) shutdown(context.Context) error {
	return s.client.Shutdown()
}

func (s *receiver) poll(ctx context.Context, time timeconstraints) error {
	orgs, err := s.client.Organizations(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving organizations: %w", err)
	}
	for _, org := range orgs {
		projects, err := s.client.Projects(ctx, org.ID)
		if err != nil {
			return fmt.Errorf("error retrieving projects: %w", err)
		}
		for _, project := range projects {
			processes, err := s.client.Processes(ctx, project.ID)
			if err != nil {
				return fmt.Errorf("error retrieving MongoDB Atlas processes: %w", err)
			}
			for _, process := range processes {
				if err := s.extractProcessMetrics(
					ctx,
					time,
					org.Name,
					project,
					process,
				); err != nil {
					return err
				}
				s.mb.EmitForResource(
					metadata.WithMongodbAtlasOrgName(org.Name),
					metadata.WithMongodbAtlasProjectName(project.Name),
					metadata.WithMongodbAtlasProjectID(project.ID),
					metadata.WithMongodbAtlasHostName(process.Hostname),
					metadata.WithMongodbAtlasProcessPort(strconv.Itoa(process.Port)),
					metadata.WithMongodbAtlasProcessTypeName(process.TypeName),
					metadata.WithMongodbAtlasProcessID(process.ID),
				)
			}
		}
	}
	return nil
}

func (s *receiver) extractProcessMetrics(
	ctx context.Context,
	time timeconstraints,
	orgName string,
	project *mongodbatlas.Project,
	process *mongodbatlas.Process,
) error {
	// This receiver will support both logs and metrics- if one pipeline
	//  or the other is not configured, it will be nil.
	if err := s.client.ProcessMetrics(
		ctx,
		s.mb,
		project.ID,
		process.Hostname,
		process.Port,
		time.start,
		time.end,
		time.resolution,
	); err != nil {
		return fmt.Errorf("error when polling process metrics from MongoDB Atlas: %w", err)
	}

	if err := s.extractProcessDatabaseMetrics(ctx, time, orgName, project, process); err != nil {
		return fmt.Errorf("error when polling process database metrics from MongoDB Atlas: %w", err)
	}

	if err := s.extractProcessDiskMetrics(ctx, time, orgName, project, process); err != nil {
		return fmt.Errorf("error when polling process disk metrics from MongoDB Atlas: %w", err)
	}
	return nil
}

func (s *receiver) extractProcessDatabaseMetrics(
	ctx context.Context,
	time timeconstraints,
	orgName string,
	project *mongodbatlas.Project,
	process *mongodbatlas.Process,
) error {
	processDatabases, err := s.client.ProcessDatabases(
		ctx,
		project.ID,
		process.Hostname,
		process.Port,
	)
	if err != nil {
		return fmt.Errorf("error retrieving process databases: %w", err)
	}

	for _, db := range processDatabases {
		if err := s.client.ProcessDatabaseMetrics(
			ctx,
			s.mb,
			project.ID,
			process.Hostname,
			process.Port,
			db.DatabaseName,
			time.start,
			time.end,
			time.resolution,
		); err != nil {
			return fmt.Errorf("error when polling database metrics from MongoDB Atlas: %w", err)
		}
		s.mb.EmitForResource(
			metadata.WithMongodbAtlasOrgName(orgName),
			metadata.WithMongodbAtlasProjectName(project.Name),
			metadata.WithMongodbAtlasProjectID(project.ID),
			metadata.WithMongodbAtlasHostName(process.Hostname),
			metadata.WithMongodbAtlasProcessPort(strconv.Itoa(process.Port)),
			metadata.WithMongodbAtlasProcessTypeName(process.TypeName),
			metadata.WithMongodbAtlasProcessID(process.ID),
			metadata.WithMongodbAtlasDbName(db.DatabaseName),
		)
	}
	return nil
}

func (s *receiver) extractProcessDiskMetrics(
	ctx context.Context,
	time timeconstraints,
	orgName string,
	project *mongodbatlas.Project,
	process *mongodbatlas.Process,
) error {
	for _, disk := range s.client.ProcessDisks(ctx, project.ID, process.Hostname, process.Port) {
		if err := s.client.ProcessDiskMetrics(
			ctx,
			s.mb,
			project.ID,
			process.Hostname,
			process.Port,
			disk.PartitionName,
			time.start,
			time.end,
			time.resolution,
		); err != nil {
			return fmt.Errorf("error when polling from MongoDB Atlas: %w", err)
		}
		s.mb.EmitForResource(
			metadata.WithMongodbAtlasOrgName(orgName),
			metadata.WithMongodbAtlasProjectName(project.Name),
			metadata.WithMongodbAtlasProjectID(project.ID),
			metadata.WithMongodbAtlasHostName(process.Hostname),
			metadata.WithMongodbAtlasProcessPort(strconv.Itoa(process.Port)),
			metadata.WithMongodbAtlasProcessTypeName(process.TypeName),
			metadata.WithMongodbAtlasProcessID(process.ID),
			metadata.WithMongodbAtlasDiskPartition(disk.PartitionName),
		)
	}
	return nil
}

// Log receiver logic
func (s *receiver) Start(ctx context.Context, host component.Host) error {
	go func() {
		s.KickoffReceiver(ctx)
	}()
	return nil
}

func (s *receiver) Shutdown(ctx context.Context) error {
	for _, stopperChan := range s.stopperChanList {
		close(stopperChan)
	}
	ctx.Done()
	return nil
}

func (s *receiver) getHostLogs(groupID, hostname, logName string) ([]model.LogEntry, error) {
	// Get gzip bytes buffer from API
	buf, err := s.client.GetLogs(context.Background(), groupID, hostname, logName, collection_interval)
	if err != nil {
		return nil, err
	}
	// Pass this into a gzip reader for decoding
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	// Logs are in JSON format so create a JSON decoder to process them
	dec := json.NewDecoder(reader)

	entries := make([]model.LogEntry, 0)
	for {
		var entry model.LogEntry
		err := dec.Decode(&entry)
		if errors.Is(err, io.EOF) {
			return entries, nil
		}
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}
}

func (s *receiver) getHostAuditLogs(groupID, hostname, logName string) ([]model.AuditLog, error) {
	// Get gzip bytes buffer from API
	buf, err := s.client.GetLogs(context.Background(), groupID, hostname, logName, collection_interval)
	if err != nil {
		return nil, err
	}
	// Pass this into a gzip reader for decoding
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	// Logs are in JSON format so create a JSON decoder to process them
	dec := json.NewDecoder(reader)

	entries := make([]model.AuditLog, 0)
	for {
		var entry model.AuditLog
		err := dec.Decode(&entry)
		if errors.Is(err, io.EOF) {
			return entries, nil
		}
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}
}

func (s *receiver) sendLogs(r resourceInfo, logName string) {
	logs, err := s.getHostLogs(r.Project.ID, r.Hostname, logName)
	if err != nil && err != io.EOF {
		s.log.Warn("Failed to retreive logs", zap.Error(err))
	}

	for _, log := range logs {
		r.LogName = logName
		plog := mongodbEventToLogData(s.log, &log, r)
		s.consumer.ConsumeLogs(context.Background(), plog)
	}
}

func (s *receiver) sendAuditLogs(r resourceInfo, logName string) {
	logs, err := s.getHostAuditLogs(r.Project.ID, r.Hostname, logName)
	if err != nil && err != io.EOF {
		s.log.Warn("Failed to retreive logs", zap.Error(err))
	}

	for _, log := range logs {
		r.LogName = logName
		plog := mongodbAuditEventToLogData(s.log, &log, r)
		s.consumer.ConsumeLogs(context.Background(), plog)
	}
}
