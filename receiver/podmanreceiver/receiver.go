// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows
// +build !windows

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver/internal/metadata"
)

type metricsReceiver struct {
	config        *Config
	set           receiver.CreateSettings
	clientFactory clientFactory
	scraper       *ContainerScraper
	mb            *metadata.MetricsBuilder
}

func newMetricsReceiver(
	_ context.Context,
	set receiver.CreateSettings,
	config *Config,
	nextConsumer consumer.Metrics,
	clientFactory clientFactory,
) (receiver.Metrics, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	if clientFactory == nil {
		clientFactory = newLibpodClient
	}

	recv := &metricsReceiver{
		config:        config,
		clientFactory: clientFactory,
		set:           set,
		mb:            metadata.NewMetricsBuilder(config.MetricsBuilderConfig, set),
	}

	scrp, err := scraperhelper.NewScraper(metadata.Type, recv.scrape, scraperhelper.WithStart(recv.start))
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(&recv.config.ScraperControllerSettings, set, nextConsumer, scraperhelper.AddScraper(scrp))
}

func (r *metricsReceiver) start(ctx context.Context, _ component.Host) error {
	var err error
	podmanClient, err := r.clientFactory(r.set.Logger, r.config)
	if err != nil {
		return err
	}

	r.scraper = newContainerScraper(podmanClient, r.set.Logger, r.config)
	if err = r.scraper.loadContainerList(ctx); err != nil {
		return err
	}
	go r.scraper.containerEventLoop(ctx)
	return nil
}

type result struct {
	container      container
	containerStats containerStats
	err            error
}

func (r *metricsReceiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	containers := r.scraper.getContainers()
	results := make(chan result, len(containers))

	wg := &sync.WaitGroup{}
	wg.Add(len(containers))
	for _, c := range containers {
		go func(c container) {
			defer wg.Done()
			stats, err := r.scraper.fetchContainerStats(ctx, c)
			results <- result{container: c, containerStats: stats, err: err}
		}(c)
	}

	wg.Wait()
	close(results)

	var errs error
	now := pcommon.NewTimestampFromTime(time.Now())

	for res := range results {
		if res.err != nil {
			// Don't know the number of failed metrics, but one container fetch is a partial error.
			errs = multierr.Append(errs, scrapererror.NewPartialScrapeError(res.err, 0))
			continue
		}
		r.recordContainerStats(now, res.container, &res.containerStats)
	}
	return r.mb.Emit(), errs
}

func (r *metricsReceiver) recordContainerStats(now pcommon.Timestamp, container container, stats *containerStats) {
	r.recordCPUMetrics(now, stats)
	r.recordNetworkMetrics(now, stats)
	r.recordMemoryMetrics(now, stats)
	r.recordIOMetrics(now, stats)

	rb := r.mb.NewResourceBuilder()
	rb.SetContainerRuntime("podman")
	rb.SetContainerName(stats.Name)
	rb.SetContainerID(stats.ContainerID)
	rb.SetContainerImageName(container.Image)

	r.mb.EmitForResource(metadata.WithResource(rb.Emit()))
}

func (r *metricsReceiver) recordCPUMetrics(now pcommon.Timestamp, stats *containerStats) {
	r.mb.RecordContainerCPUUsageSystemDataPoint(now, int64(stats.CPUSystemNano))
	r.mb.RecordContainerCPUUsageTotalDataPoint(now, int64(stats.CPUNano))
	r.mb.RecordContainerCPUPercentDataPoint(now, stats.CPU)

	for i, cpu := range stats.PerCPU {
		r.mb.RecordContainerCPUUsagePercpuDataPoint(now, int64(cpu), fmt.Sprintf("cpu%d", i))
	}

}

func (r *metricsReceiver) recordNetworkMetrics(now pcommon.Timestamp, stats *containerStats) {
	r.mb.RecordContainerNetworkIoUsageRxBytesDataPoint(now, int64(stats.NetInput))
	r.mb.RecordContainerNetworkIoUsageTxBytesDataPoint(now, int64(stats.NetOutput))
}

func (r *metricsReceiver) recordMemoryMetrics(now pcommon.Timestamp, stats *containerStats) {
	r.mb.RecordContainerMemoryUsageTotalDataPoint(now, int64(stats.MemUsage))
	r.mb.RecordContainerMemoryUsageLimitDataPoint(now, int64(stats.MemLimit))
	r.mb.RecordContainerMemoryPercentDataPoint(now, stats.MemPerc)
}

func (r *metricsReceiver) recordIOMetrics(now pcommon.Timestamp, stats *containerStats) {
	r.mb.RecordContainerBlockioIoServiceBytesRecursiveReadDataPoint(now, int64(stats.BlockInput))
	r.mb.RecordContainerBlockioIoServiceBytesRecursiveWriteDataPoint(now, int64(stats.BlockOutput))
}
