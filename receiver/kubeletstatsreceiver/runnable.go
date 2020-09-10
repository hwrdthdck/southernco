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

package kubeletstatsreceiver

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdatautil"
	"go.opentelemetry.io/collector/obsreport"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/kubelet"
	// todo replace with scraping lib when it's ready
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver/interval"
)

var _ interval.Runnable = (*runnable)(nil)

type runnable struct {
	ctx                   context.Context
	receiverName          string
	statsProvider         *kubelet.StatsProvider
	metadataProvider      *kubelet.MetadataProvider
	consumer              consumer.MetricsConsumer
	logger                *zap.Logger
	restClient            kubelet.RestClient
	extraMetadataLabels   []kubelet.MetadataLabel
	metricGroupsToCollect map[kubelet.MetricGroup]bool
	k8sAPIClient          kubernetes.Interface
}

func newRunnable(
	ctx context.Context,
	consumer consumer.MetricsConsumer,
	restClient kubelet.RestClient,
	logger *zap.Logger,
	rOptions *receiverOptions,
) *runnable {
	return &runnable{
		ctx:                   ctx,
		receiverName:          rOptions.name,
		consumer:              consumer,
		restClient:            restClient,
		logger:                logger,
		extraMetadataLabels:   rOptions.extraMetadataLabels,
		metricGroupsToCollect: rOptions.metricGroupsToCollect,
		k8sAPIClient:          rOptions.k8sAPIClient,
	}
}

// Sets up the kubelet connection at startup time.
func (r *runnable) Setup() error {
	r.statsProvider = kubelet.NewStatsProvider(r.restClient)
	r.metadataProvider = kubelet.NewMetadataProvider(r.restClient)
	return nil
}

func (r *runnable) Run() error {
	const transport = "http"
	summary, err := r.statsProvider.StatsSummary()
	if err != nil {
		r.logger.Error("call to /stats/summary endpoint failed", zap.Error(err))
		return nil
	}

	var podsMetadata *v1.PodList
	// fetch metadata only when extra metadata labels are needed
	if len(r.extraMetadataLabels) > 0 {
		podsMetadata, err = r.metadataProvider.Pods()
		if err != nil {
			r.logger.Error("call to /pods endpoint failed", zap.Error(err))
			return nil
		}
	}

	metadata := kubelet.NewMetadata(r.extraMetadataLabels, podsMetadata, detailedPVCLabelsSetter(r.k8sAPIClient))
	mds := kubelet.MetricsData(r.logger, summary, metadata, typeStr, r.metricGroupsToCollect)
	metrics := pdatautil.MetricsFromMetricsData(mds)

	var numTimeSeries, numPoints int
	ctx := obsreport.ReceiverContext(r.ctx, typeStr, transport, r.receiverName)
	ctx = obsreport.StartMetricsReceiveOp(ctx, typeStr, transport)
	err = r.consumer.ConsumeMetrics(ctx, metrics)
	if err != nil {
		r.logger.Error("ConsumeMetricsData failed", zap.Error(err))
	} else {
		numTimeSeries, numPoints = pdatautil.MetricAndDataPointCount(metrics)
	}
	obsreport.EndMetricsReceiveOp(ctx, typeStr, numTimeSeries, numPoints, err)

	return nil
}

func detailedPVCLabelsSetter(k8sAPIClient kubernetes.Interface) func(volumeClaim, namespace string, labels map[string]string) error {
	return func(volumeClaim, namespace string, labels map[string]string) error {
		if k8sAPIClient == nil {
			return nil
		}

		ctx := context.Background()
		pvc, err := k8sAPIClient.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, volumeClaim, metav1.GetOptions{})
		if err != nil {
			return err
		}

		volName := pvc.Spec.VolumeName
		if volName == "" {
			return fmt.Errorf("PersistentVolumeClaim %s does not have a volume name", pvc.Name)
		}

		pv, err := k8sAPIClient.CoreV1().PersistentVolumes().Get(ctx, volName, metav1.GetOptions{})
		if err != nil {
			return err
		}

		kubelet.GetPersistentVolumeLabels(pv.Spec.PersistentVolumeSource, labels)
		return nil
	}
}
