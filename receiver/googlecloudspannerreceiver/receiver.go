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

package googlecloudspannerreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver"

import (
	"context"
	_ "embed"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filterfactory"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadataparser"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"
)

//go:embed "internal/metadataconfig/metadata.yaml"
var metadataYaml []byte

var _ component.MetricsReceiver = (*googleCloudSpannerReceiver)(nil)

type googleCloudSpannerReceiver struct {
	logger         *zap.Logger
	config         *Config
	cancel         context.CancelFunc
	projectReaders []statsreader.CompositeReader
	metricsBuilder metadata.MetricsBuilder
}

func newGoogleCloudSpannerReceiver(logger *zap.Logger, config *Config) *googleCloudSpannerReceiver {
	return &googleCloudSpannerReceiver{
		logger: logger,
		config: config,
	}
}

func (r *googleCloudSpannerReceiver) Scrape(ctx context.Context) (pmetric.Metrics, error) {
	var allMetricsDataPoints []*metadata.MetricsDataPoint

	for _, projectReader := range r.projectReaders {
		dataPoints, err := projectReader.Read(ctx)
		if err != nil {
			return pmetric.Metrics{}, err
		}

		allMetricsDataPoints = append(allMetricsDataPoints, dataPoints...)
	}

	return r.metricsBuilder.Build(allMetricsDataPoints)
}

func (r *googleCloudSpannerReceiver) Start(ctx context.Context, _ component.Host) error {
	ctx, r.cancel = context.WithCancel(ctx)
	err := r.initialize(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *googleCloudSpannerReceiver) Shutdown(context.Context) error {
	for _, projectReader := range r.projectReaders {
		projectReader.Shutdown()
	}

	err := r.metricsBuilder.Shutdown()
	if err != nil {
		return err
	}

	r.cancel()

	return nil
}

func (r *googleCloudSpannerReceiver) initialize(ctx context.Context) error {
	parsedMetadata, err := metadataparser.ParseMetadataConfig(metadataYaml)
	if err != nil {
		return fmt.Errorf("error occurred during parsing of metadata: %w", err)
	}

	err = r.initializeProjectReaders(ctx, parsedMetadata)
	if err != nil {
		return err
	}

	return r.initializeMetricsBuilder(parsedMetadata)
}

func (r *googleCloudSpannerReceiver) initializeProjectReaders(ctx context.Context,
	parsedMetadata []*metadata.MetricsMetadata) error {

	readerConfig := statsreader.ReaderConfig{
		BackfillEnabled:        r.config.BackfillEnabled,
		TopMetricsQueryMaxRows: r.config.TopMetricsQueryMaxRows,
	}

	for _, project := range r.config.Projects {
		projectReader, err := newProjectReader(ctx, r.logger, project, parsedMetadata, readerConfig)
		if err != nil {
			return err
		}

		r.projectReaders = append(r.projectReaders, projectReader)
	}

	return nil
}

func (r *googleCloudSpannerReceiver) initializeMetricsBuilder(parsedMetadata []*metadata.MetricsMetadata) error {
	r.logger.Debug("Constructing metrics builder")

	projectAmount := len(r.config.Projects)
	instanceAmount := 0
	databaseAmount := 0

	for _, project := range r.config.Projects {
		instanceAmount += len(project.Instances)

		for _, instance := range project.Instances {
			databaseAmount += len(instance.Databases)
		}
	}

	factoryConfig := &filterfactory.ItemFilterFactoryConfig{
		MetadataItems:  parsedMetadata,
		TotalLimit:     r.config.CardinalityTotalLimit,
		ProjectAmount:  projectAmount,
		InstanceAmount: instanceAmount,
		DatabaseAmount: databaseAmount,
	}
	itemFilterResolver, err := filterfactory.NewItemFilterResolver(r.logger, factoryConfig)
	if err != nil {
		return err
	}

	r.metricsBuilder = metadata.NewMetricsFromDataPointBuilder(itemFilterResolver)

	return nil
}

func newProjectReader(ctx context.Context, logger *zap.Logger, project Project, parsedMetadata []*metadata.MetricsMetadata,
	readerConfig statsreader.ReaderConfig) (*statsreader.ProjectReader, error) {
	logger.Debug("Constructing project reader for project", zap.String("project id", project.ID))

	databaseReadersCount := 0
	for _, instance := range project.Instances {
		databaseReadersCount += len(instance.Databases)
	}

	databaseReaders := make([]statsreader.CompositeReader, databaseReadersCount)
	databaseReaderIndex := 0
	for _, instance := range project.Instances {
		for _, database := range instance.Databases {
			logger.Debug("Constructing database reader for combination of project, instance, database",
				zap.String("project id", project.ID), zap.String("instance id", instance.ID), zap.String("database", database))

			databaseID := datasource.NewDatabaseID(project.ID, instance.ID, database)

			databaseReader, err := statsreader.NewDatabaseReader(ctx, parsedMetadata, databaseID,
				project.ServiceAccountKey, readerConfig, logger)
			if err != nil {
				return nil, err
			}

			databaseReaders[databaseReaderIndex] = databaseReader
			databaseReaderIndex++
		}
	}

	return statsreader.NewProjectReader(databaseReaders, logger), nil
}
