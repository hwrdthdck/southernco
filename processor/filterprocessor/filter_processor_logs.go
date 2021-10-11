// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filterprocessor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterlog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filtermatcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterset"
)

type filterLogProcessor struct {
	cfg              *Config
	exclude          filterlog.Matcher
	excludeResources filtermatcher.AttributesMatcher
	excludeRecords   filtermatcher.AttributesMatcher
	include          filterlog.Matcher
	includeResources filtermatcher.AttributesMatcher
	includeRecords   filtermatcher.AttributesMatcher
	logger           *zap.Logger
}

func newFilterLogsProcessor(logger *zap.Logger, cfg *Config) (*filterLogProcessor, error) {

	inc, includeResources, includeRecords, err := createLogsMatcher(cfg.Logs.Include)
	if err != nil {
		logger.Error(
			"filterlog: Error creating include logs matcher", zap.Error(err),
		)
		return nil, err
	}

	exc, excludeResources, excludeRecords, err := createLogsMatcher(cfg.Logs.Exclude)
	if err != nil {
		logger.Error(
			"filterlog: Error creating exclude logs matcher", zap.Error(err),
		)
		return nil, err
	}

	return &filterLogProcessor{
		cfg:              cfg,
		include:          inc,
		includeResources: includeResources,
		includeRecords:   includeRecords,
		exclude:          exc,
		excludeResources: excludeResources,
		excludeRecords:   excludeRecords,
		logger:           logger,
	}, nil
}

func createLogsMatcher(lp *filterlog.LogMatchProperties) (logMatcher filterlog.Matcher, resourceMatcher filtermatcher.AttributesMatcher, recordMatcher filtermatcher.AttributesMatcher, err error) {
	// Nothing specified in configuration
	if lp == nil {
		return nil, nil, nil, nil
	}

	recordMatcher, err = filtermatcher.NewAttributesMatcher(
		filterset.Config{
			MatchType: filterset.MatchType(lp.MatchType),
		},
		lp.RecordAttributes,
	)
	if err != nil {
		return nil, nil, recordMatcher, fmt.Errorf("record attributes: %v", err)
	}

	resourceMatcher, err = filtermatcher.NewAttributesMatcher(
		filterset.Config{
			MatchType: filterset.MatchType(lp.MatchType),
		},
		lp.ResourceAttributes,
	)
	if err != nil {
		return nil, resourceMatcher, recordMatcher, fmt.Errorf("resource attributes: %v", err)
	}

	if lp.MatchType != filterlog.Expr {
		return nil, resourceMatcher, recordMatcher, nil
	}

	logMatcher, err = filterlog.NewMatcher(lp)
	if err != nil {
		return logMatcher, resourceMatcher, recordMatcher, fmt.Errorf("expr matcher: %v", err)
	}

	return logMatcher, resourceMatcher, recordMatcher, nil
}

func (flp *filterLogProcessor) ProcessLogs(ctx context.Context, logs pdata.Logs) (pdata.Logs, error) {
	rLogs := logs.ResourceLogs()

	// Filter logs by resource level attributes
	rLogs.RemoveIf(func(rm pdata.ResourceLogs) bool {
		return flp.shouldSkipLogsForResource(rm.Resource())
	})

	// Filter logs by record level attributes
	flp.filterByRecordAttributes(rLogs)

	if rLogs.Len() == 0 {
		return logs, processorhelper.ErrSkipProcessingData
	}

	return logs, nil
}

func (flp *filterLogProcessor) filterByRecordAttributes(rLogs pdata.ResourceLogsSlice) {
	for i := 0; i < rLogs.Len(); i++ {
		ills := rLogs.At(i).InstrumentationLibraryLogs()

		for j := 0; j < ills.Len(); j++ {
			ls := ills.At(j).Logs()

			ls.RemoveIf(func(lr pdata.LogRecord) bool {
				return flp.shouldSkipLogsForRecord(lr)
			})
		}

		ills.RemoveIf(func(ill pdata.InstrumentationLibraryLogs) bool {
			return ill.Logs().Len() == 0
		})
	}

	rLogs.RemoveIf(func(rl pdata.ResourceLogs) bool {
		return rl.InstrumentationLibraryLogs().Len() == 0
	})
}

// shouldSkipLogsForRecord determines if a log record should be processed.
// True is returned when a log record should be skipped.
// False is returned when a log record should not be skipped.
// The logic determining if a log record should be skipped is set in the
// record attribute configuration.
func (flp *filterLogProcessor) shouldSkipLogsForRecord(lr pdata.LogRecord) bool {
	if flp.includeRecords != nil {
		matches := flp.includeRecords.Match(lr.Attributes())
		if !matches {
			return true
		}
	}

	if flp.include != nil {
		matches := flp.include.MatchLogRecord(lr)
		if !matches {
			return true
		}
	}

	if flp.excludeRecords != nil {
		matches := flp.excludeRecords.Match(lr.Attributes())
		if matches {
			return true
		}
	}

	return false
}

// shouldSkipLogsForResource determines if a log should be processed.
// True is returned when a log should be skipped.
// False is returned when a log should not be skipped.
// The logic determining if a log should be skipped is set in the resource attribute configuration.
func (flp *filterLogProcessor) shouldSkipLogsForResource(resource pdata.Resource) bool {
	resourceAttributes := resource.Attributes()
	if flp.includeResources != nil {
		matches := flp.includeResources.Match(resourceAttributes)
		if !matches {
			return true
		}
	}

	if flp.excludeResources != nil {
		matches := flp.excludeResources.Match(resourceAttributes)
		if matches {
			return true
		}
	}

	return false
}
