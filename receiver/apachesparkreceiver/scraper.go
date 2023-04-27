// Copyright 2020 OpenTelemetry Authors
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

package apachesparkreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/models"
)

var (
	errClientNotInit = errors.New("client not initialized")
)

type sparkScraper struct {
	client   client
	logger   *zap.Logger
	config   *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

func newSparkScraper(logger *zap.Logger, cfg *Config, settings receiver.CreateSettings) *sparkScraper {
	return &sparkScraper{
		logger:   logger,
		config:   cfg,
		settings: settings.TelemetrySettings,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
	}
}

func (s *sparkScraper) start(_ context.Context, host component.Host) (err error) {
	httpClient, err := newApacheSparkClient(s.config, host, s.settings)
	if err != nil {
		return fmt.Errorf("failed to start: %w", err)
	}
	s.client = httpClient
	return nil
}

func (s *sparkScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	now := pcommon.NewTimestampFromTime(time.Now())
	var scrapeErrors scrapererror.ScrapeErrors

	if s.client == nil {
		return pmetric.NewMetrics(), errClientNotInit
	}

	// call applications endpoint
	// not getting app name for now, just ids
	appIds, err := s.client.GetApplicationIDs()
	if err != nil {
		scrapeErrors.AddPartial(1, err)
		s.logger.Warn("Failed to scrape application ids", zap.Error(err))
	}

	// get stats from the 'metrics' endpoint
	clusterStats, err := s.client.GetClusterStats()
	if err != nil {
		scrapeErrors.AddPartial(1, err)
		s.logger.Warn("Failed to scrape cluster stats", zap.Error(err))
	}

	for _, appID := range appIds {
		s.collectCluster(clusterStats, now, appID, "TODO-appName") // TODO
	}

	// for each application id, get stats from stages & executors endpoints
	for _, appID := range appIds {

		stageStats, err := s.client.GetStageStats(appID)
		if err != nil {
			scrapeErrors.AddPartial(1, err)
			s.logger.Warn("Failed to scrape stage stats", zap.Error(err))
		}
		s.collectStage(*stageStats, now, appID, "TODO-appName") // TODO

		executorStats, err := s.client.GetExecutorStats(appID)
		if err != nil {
			scrapeErrors.AddPartial(1, err)
			s.logger.Warn("Failed to scrape executor stats", zap.Error(err))
		}
		s.collectExecutor(*executorStats, now, appID, "TODO-appName") // TODO

		jobStats, err := s.client.GetJobStats(appID)
		if err != nil {
			scrapeErrors.AddPartial(1, err)
			s.logger.Warn("Failed to scrape job stats", zap.Error(err))
		}
		s.collectJob(*jobStats, now, appID, "TODO-appName") // TODO
	}

	return s.mb.Emit(), scrapeErrors.Combine()
}

// TODO: Add app name as attribute to everything
func (s *sparkScraper) collectCluster(clusterStats *models.ClusterProperties, now pcommon.Timestamp, appID string, appName string) {
	key := fmt.Sprintf("%s.driver.BlockManager.disk.diskSpaceUsed", appID)
	s.mb.RecordSparkDriverBlockManagerDiskDiskSpaceUsedDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)
 
	key = fmt.Sprintf("%s.driver.BlockManager.memory.offHeapMemUsed_MB", appID)
	s.mb.RecordSparkDriverBlockManagerMemoryUsedDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName, metadata.AttributeLocationOffHeap)
	key = fmt.Sprintf("%s.driver.BlockManager.memory.onHeapMemUsed_MB", appID)
	s.mb.RecordSparkDriverBlockManagerMemoryUsedDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName, metadata.AttributeLocationOnHeap)

	key = fmt.Sprintf("%s.driver.BlockManager.memory.remainingOffHeapMem_MB", appID)
	s.mb.RecordSparkDriverBlockManagerMemoryRemainingDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName, metadata.AttributeLocationOffHeap)
	key = fmt.Sprintf("%s.driver.BlockManager.memory.remainingOnHeapMem_MB", appID)
	s.mb.RecordSparkDriverBlockManagerMemoryRemainingDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName, metadata.AttributeLocationOnHeap)

	key = fmt.Sprintf("%s.driver.HiveExternalCatalog.fileCacheHits", appID)
	s.mb.RecordSparkDriverHiveExternalCatalogFileCacheHitsDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.HiveExternalCatalog.filesDiscovered", appID)
	s.mb.RecordSparkDriverHiveExternalCatalogFilesDiscoveredDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.HiveExternalCatalog.hiveClientCalls", appID)
	s.mb.RecordSparkDriverHiveExternalCatalogHiveClientCallsDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.HiveExternalCatalog.parallelListingJobCount", appID)
	s.mb.RecordSparkDriverHiveExternalCatalogParallelListingJobsDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.HiveExternalCatalog.partitionsFetched", appID)
	s.mb.RecordSparkDriverHiveExternalCatalogPartitionsFetchedDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.compilationTime", appID)
	s.mb.RecordSparkDriverCodeGeneratorCompilationCountDataPoint(now, int64(clusterStats.Histograms[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.compilationTime", appID)
	s.mb.RecordSparkDriverCodeGeneratorCompilationAverageTimeDataPoint(now, clusterStats.Histograms[key].Mean, appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.generatedClassSize", appID)
	s.mb.RecordSparkDriverCodeGeneratorGeneratedClassCountDataPoint(now, int64(clusterStats.Histograms[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.generatedClassSize", appID)
	s.mb.RecordSparkDriverCodeGeneratorGeneratedClassAverageSizeDataPoint(now, clusterStats.Histograms[key].Mean, appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.generatedMethodSize", appID)
	s.mb.RecordSparkDriverCodeGeneratorGeneratedMethodCountDataPoint(now, int64(clusterStats.Histograms[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.generatedMethodSize", appID)
	s.mb.RecordSparkDriverCodeGeneratorGeneratedMethodAverageSizeDataPoint(now, clusterStats.Histograms[key].Mean, appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.sourceCodeSize", appID)
	s.mb.RecordSparkDriverCodeGeneratorSourceCodeCountDataPoint(now, int64(clusterStats.Histograms[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.CodeGenerator.sourceCodeSize", appID)
	s.mb.RecordSparkDriverCodeGeneratorSourceCodeAverageSizeDataPoint(now, clusterStats.Histograms[key].Mean, appID, appName)

	key = fmt.Sprintf("%s.driver.DAGScheduler.job.activeJobs", appID)
	s.mb.RecordSparkDriverDagSchedulerJobActiveJobsDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	key = fmt.Sprintf("%s.driver.DAGScheduler.job.allJobs", appID)
	s.mb.RecordSparkDriverDagSchedulerJobAllJobsDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	key = fmt.Sprintf("%s.driver.DAGScheduler.stage.failedStages", appID)
	s.mb.RecordSparkDriverDagSchedulerStageFailedStagesDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	key = fmt.Sprintf("%s.driver.DAGScheduler.stage.runningStages", appID)
	s.mb.RecordSparkDriverDagSchedulerStageRunningStagesDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	key = fmt.Sprintf("%s.driver.DAGScheduler.stage.waitingStages", appID)
	s.mb.RecordSparkDriverDagSchedulerStageWaitingStagesDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	key = fmt.Sprintf("%s.driver.LiveListenerBus.numEventsPosted", appID)
	s.mb.RecordSparkDriverLiveListenerBusEventsPostedDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.LiveListenerBus.queue.appStatus.listenerProcessingTime", appID)
	s.mb.RecordSparkDriverLiveListenerBusListenerProcessingTimeAverageDataPoint(now, clusterStats.Histograms[key].Mean, appID, appName)

	key = fmt.Sprintf("%s.driver.LiveListenerBus.queue.appStatus.numDroppedEvents", appID)
	s.mb.RecordSparkDriverLiveListenerBusEventsDroppedDataPoint(now, int64(clusterStats.Counters[key].Count), appID, appName)

	key = fmt.Sprintf("%s.driver.LiveListenerBus.queue.appStatus.size", appID)
	s.mb.RecordSparkDriverLiveListenerBusQueueSizeDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	key = fmt.Sprintf("%s.driver.JVMCPU.jvmCpuTime", appID)
	s.mb.RecordSparkDriverJvmCPUTimeDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, appName)

	// TODO: revisit after spec review
	// key = fmt.Sprintf("%s.driver.ExecutorMetrics.JVMOffHeapMemory", appID)
	// s.mb.RecordSparkDriverExecutorMetricsJvmMemoryDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, metadata.AttributeLocationOffHeap)
	// key = fmt.Sprintf("%s.driver.ExecutorMetrics.JVMHeapMemory", appID)
	// s.mb.RecordSparkDriverExecutorMetricsJvmMemoryDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, metadata.AttributeLocationOnHeap)

	// key = fmt.Sprintf("%s.driver.ExecutorMetrics.OffHeapExecutionMemory", appID)
	// s.mb.RecordSparkDriverExecutorMetricsExecutionMemoryDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, metadata.AttributeLocationOffHeap)
	// key = fmt.Sprintf("%s.driver.ExecutorMetrics.OnHeapExecutionMemory", appID)
	// s.mb.RecordSparkDriverExecutorMetricsJvmMemoryDataPoint(now, int64(clusterStats.Gauges[key].Value), appID, metadata.AttributeLocationOnHeap)

}

func (s *sparkScraper) collectStage(stageStats models.Stages, now pcommon.Timestamp, appID string, appName string) {
	for i := range stageStats {
		var stageStatus metadata.AttributeStageStatus
		switch stageStats[i].Status {
		case "ACTIVE":
			stageStatus = metadata.AttributeStageStatusACTIVE
		case "COMPLETE":
			stageStatus = metadata.AttributeStageStatusCOMPLETE
		case "PENDING":
			stageStatus = metadata.AttributeStageStatusPENDING
		case "FAILED":
			stageStatus = metadata.AttributeStageStatusFAILED
		} 

		s.mb.RecordSparkStageActiveTasksDataPoint(now, int64(stageStats[i].ExecutorRunTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageCompleteTasksDataPoint(now, int64(stageStats[i].ExecutorRunTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageFailedTasksDataPoint(now, int64(stageStats[i].ExecutorRunTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageKilledTasksDataPoint(now, int64(stageStats[i].ExecutorRunTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageExecutorRunTimeDataPoint(now, int64(stageStats[i].ExecutorRunTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageExecutorCPUTimeDataPoint(now, int64(stageStats[i].ExecutorCpuTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageResultSizeDataPoint(now, int64(stageStats[i].ResultSize), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageJvmGcTimeDataPoint(now, int64(stageStats[i].JvmGcTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageMemorySpilledDataPoint(now, int64(stageStats[i].MemoryBytesSpilled), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageDiskSpaceSpilledDataPoint(now, int64(stageStats[i].DiskBytesSpilled), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStagePeakExecutionMemoryDataPoint(now, int64(stageStats[i].PeakExecutionMemory), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageInputBytesDataPoint(now, int64(stageStats[i].InputBytes), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageInputRecordsDataPoint(now, int64(stageStats[i].InputRecords), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageOutputBytesDataPoint(now, int64(stageStats[i].OutputBytes), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageOutputRecordsDataPoint(now, int64(stageStats[i].OutputRecords), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleBlocksFetchedDataPoint(now, int64(stageStats[i].ShuffleRemoteBlocksFetched), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus, metadata.AttributeSourceRemote)
		s.mb.RecordSparkStageShuffleBlocksFetchedDataPoint(now, int64(stageStats[i].ShuffleLocalBlocksFetched), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus, metadata.AttributeSourceLocal)
		s.mb.RecordSparkStageShuffleFetchWaitTimeDataPoint(now, int64(stageStats[i].ShuffleFetchWaitTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleBytesReadDataPoint(now, int64(stageStats[i].ShuffleRemoteBytesRead), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus, metadata.AttributeSourceRemote)
		s.mb.RecordSparkStageShuffleBytesReadDataPoint(now, int64(stageStats[i].ShuffleLocalBytesRead), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus, metadata.AttributeSourceLocal)
		s.mb.RecordSparkStageShuffleRemoteBytesReadToDiskDataPoint(now, int64(stageStats[i].ShuffleRemoteBytesReadToDisk), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleReadBytesDataPoint(now, int64(stageStats[i].ShuffleReadBytes), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleReadRecordsDataPoint(now, int64(stageStats[i].ShuffleReadRecords), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleWriteBytesDataPoint(now, int64(stageStats[i].ShuffleWriteBytes), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleWriteRecordsDataPoint(now, int64(stageStats[i].ShuffleWriteRecords), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
		s.mb.RecordSparkStageShuffleWriteTimeDataPoint(now, int64(stageStats[i].ShuffleWriteTime), appID, appName, stageStats[i].StageId, stageStats[i].AttemptId, stageStatus)
	}
}

func (s *sparkScraper) collectExecutor(executorStats models.Executors, now pcommon.Timestamp, appID string, appName string) {
	for i := range executorStats {
		s.mb.RecordSparkExecutorMemoryUsedDataPoint(now, executorStats[i].MemoryUsed, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorDiskUsedDataPoint(now, executorStats[i].DiskUsed, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorMaxTasksDataPoint(now, executorStats[i].MaxTasks, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorActiveTasksDataPoint(now, executorStats[i].ActiveTasks, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorFailedTasksDataPoint(now, executorStats[i].FailedTasks, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorCompletedTasksDataPoint(now, executorStats[i].CompletedTasks, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorDurationDataPoint(now, executorStats[i].TotalDuration, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorGcTimeDataPoint(now, executorStats[i].TotalGCTime, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorInputBytesDataPoint(now, executorStats[i].TotalInputBytes, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorShuffleReadBytesDataPoint(now, executorStats[i].TotalShuffleRead, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorShuffleWriteBytesDataPoint(now, executorStats[i].TotalShuffleWrite, appID, appName, executorStats[i].Id)
		s.mb.RecordSparkExecutorUsedStorageMemoryDataPoint(now, executorStats[i].UsedOnHeapStorageMemory, appID, appName, executorStats[i].Id, metadata.AttributeLocationOnHeap)
		s.mb.RecordSparkExecutorUsedStorageMemoryDataPoint(now, executorStats[i].UsedOffHeapStorageMemory, appID, appName, executorStats[i].Id, metadata.AttributeLocationOffHeap)
		s.mb.RecordSparkExecutorTotalStorageMemoryDataPoint(now, executorStats[i].TotalOnHeapStorageMemory, appID, appName, executorStats[i].Id, metadata.AttributeLocationOnHeap)
		s.mb.RecordSparkExecutorTotalStorageMemoryDataPoint(now, executorStats[i].TotalOffHeapStorageMemory, appID, appName, executorStats[i].Id, metadata.AttributeLocationOffHeap)
	}
}

func (s *sparkScraper) collectJob(jobStats models.Jobs, now pcommon.Timestamp, appID string, appName string) {
	for i := range jobStats {
		s.mb.RecordSparkJobActiveTasksDataPoint(now, int64(jobStats[i].NumActiveTasks), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobCompletedTasksDataPoint(now, int64(jobStats[i].NumCompletedTasks), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobSkippedTasksDataPoint(now, int64(jobStats[i].NumSkippedTasks), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobFailedTasksDataPoint(now, int64(jobStats[i].NumFailedTasks), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobActiveStagesDataPoint(now, int64(jobStats[i].NumActiveStages), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobCompletedStagesDataPoint(now, int64(jobStats[i].NumCompletedStages), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobSkippedStagesDataPoint(now, int64(jobStats[i].NumSkippedStages), appID, appName, jobStats[i].JobId)
		s.mb.RecordSparkJobFailedStagesDataPoint(now, int64(jobStats[i].NumFailedStages), appID, appName, jobStats[i].JobId)
	}
}
