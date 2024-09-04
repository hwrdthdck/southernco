// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/huaweicloudcesreceiver/internal"

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
)

// This interface should have all the function defined inside https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/v0.1.110/services/ces/v1/ces_client.go
// Check https://github.com/vektra/mockery on how to install it on your machine.
//
//go:generate mockery --name CesClient --case=underscore --output=./mocks
type CesClient interface {
	ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error)
	ShowMetricData(request *model.ShowMetricDataRequest) (*model.ShowMetricDataResponse, error)
}
