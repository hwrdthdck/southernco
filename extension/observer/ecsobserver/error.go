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

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver/internal/errctx"
)

// error.go defines common error interfaces and util methods for generating reports
// for log and metrics that can be used for debugging.

const (
	errKeyTask   = "task"
	errKeyTarget = "target"
)

type errWithAttributes interface {
	// message does not include attributes like task arn etc.
	// and expect the caller extract them using getters.
	message() string
	// zapFields will be logged as json attribute and allows searching and filter backend like cloudwatch.
	// For example { $.ErrScope == "Target" } list all the error whose scope is a (scrape) target.
	zapFields() []zap.Field
}

// hasCriticalError returns first critical error.
// Currently only access error and cluster not found are treated as critical.
func hasCriticalError(logger *zap.Logger, err error) error {
	merr := multierr.Errors(err)
	if merr == nil {
		merr = []error{err} // fake a multi error
	}
	for _, err := range merr {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) {
			// NOTE: we don't use zap.Error because the stack trace is quite useless here
			// We print the error after entire fetch and match loop is done, and source
			// of these error are from user config, wrong IAM, typo in cluster name etc.
			switch awsErr.Code() {
			case ecs.ErrCodeAccessDeniedException:
				logger.Error("AccessDenied", zap.String("ErrMessage", awsErr.Message()))
				return awsErr
			case ecs.ErrCodeClusterNotFoundException:
				logger.Error("Cluster NotFound", zap.String("ErrMessage", awsErr.Message()))
				return awsErr
			}
		}
	}
	return nil
}

func printErrors(logger *zap.Logger, err error) {
	merr := multierr.Errors(err)
	if merr == nil {
		return
	}

	for _, err := range merr {
		m := err.Error()
		// Use the short message, this makes searching the code via error message easier
		// as additional info are flushed as fields.
		var errAttr errWithAttributes
		if errors.As(err, &errAttr) {
			m = errAttr.message()
		}
		fields, scope := extractErrorFields(err)
		fields = append(fields, zap.String("ErrScope", scope))
		logger.Error(m, fields...)
	}
}

func extractErrorFields(err error) ([]zap.Field, string) {
	var fields []zap.Field
	scope := "Unknown"
	var errAttr errWithAttributes
	// Stop early because we are only attaching value for our internal errors.
	if !errors.As(err, &errAttr) {
		return fields, scope
	}
	fields = errAttr.zapFields()
	v, ok := errctx.ValueFrom(err, errKeyTask)
	if ok {
		// Rename ok to tok because linter says it shadows outer ok.
		// Though the linter seems to allow the similar block to shadow...
		if task, tok := v.(*taskAnnotated); tok {
			fields = append(fields, zap.String("TaskArn", aws.StringValue(task.Task.TaskArn)))
			scope = "taskAnnotated"
		}
	}
	v, ok = errctx.ValueFrom(err, errKeyTarget)
	if ok {
		if target, ok := v.(matchedTarget); ok {
			fields = append(fields, zap.String("matcherType", target.MatcherType.String()))
			scope = "Target"
		}
	}
	return fields, scope
}
