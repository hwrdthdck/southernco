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

package k8sapiserver

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"

	ci "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"
)

const (
	lockName = "aoc-clusterleader"
)

// eventBroadcaster is adpated from record.EventBroadcaster
type eventBroadcaster interface {
	// StartRecordingToSink starts sending events received from this EventBroadcaster to the given
	// sink. The return value can be ignored or used to stop recording, if desired.
	StartRecordingToSink(sink record.EventSink) watch.Interface
	// StartLogging starts sending events received from this EventBroadcaster to the given logging
	// function. The return value can be ignored or used to stop recording, if desired.
	StartLogging(logf func(format string, args ...interface{})) watch.Interface
	// NewRecorder returns an EventRecorder that can be used to send events to this EventBroadcaster
	// with the event source set to the given event source.
	NewRecorder(scheme *runtime.Scheme, source v1.EventSource) record.EventRecorder
}

// K8sAPIServer is a struct that produces metrics from kubernetes api server
type K8sAPIServer struct {
	nodeName            string //get the value from downward API
	logger              *zap.Logger
	clusterNameProvider clusterNameProvider
	cancel              context.CancelFunc

	leading   bool
	k8sClient *k8sclient.K8sClient

	// the following can be set to mocks in testing
	broadcaster eventBroadcaster
	// the close of isLeadingC indicates the leader election is done. This is used in testing
	isLeadingC chan bool
}

type clusterNameProvider interface {
	GetClusterName() string
}

type k8sAPIServerOption func(*K8sAPIServer)

// New creates a k8sApiServer which can generate cluster-level metrics
func New(clusterNameProvider clusterNameProvider, logger *zap.Logger, options ...k8sAPIServerOption) (*K8sAPIServer, error) {
	_, cancel := context.WithCancel(context.Background())
	k := &K8sAPIServer{
		logger:              logger,
		clusterNameProvider: clusterNameProvider,
		k8sClient:           k8sclient.Get(logger),
		broadcaster:         record.NewBroadcaster(),
		cancel:              cancel,
	}

	for _, opt := range options {
		opt(k)
	}

	if k.k8sClient == nil {
		return nil, errors.New("failed to start k8sapiserver because k8sclient is nil")
	}

	if err := k.init(); err != nil {
		return nil, fmt.Errorf("fail to initialize k8sapiserver, err: %v", err)
	}

	return k, nil
}

// GetMetrics returns an array of metrics
func (k *K8sAPIServer) GetMetrics() []pdata.Metrics {
	var result []pdata.Metrics

	//don't emit metrics if the cluster name is not detected
	clusterName := k.clusterNameProvider.GetClusterName()
	if clusterName == "" {
		k.logger.Warn("Failed to detect cluster name. Drop all metrics")
		return result
	}

	if k.leading {
		k.logger.Info("collect data from K8s API Server...")
		timestampNs := strconv.FormatInt(time.Now().UnixNano(), 10)
		client := k.k8sClient

		fields := map[string]interface{}{
			"cluster_failed_node_count": client.Node.ClusterFailedNodeCount(),
			"cluster_node_count":        client.Node.ClusterNodeCount(),
		}
		attributes := map[string]string{
			ci.ClusterNameKey: clusterName,
			ci.MetricType:     ci.TypeCluster,
			ci.Timestamp:      timestampNs,
			ci.Version:        "0",
		}
		if k.nodeName != "" {
			attributes["NodeName"] = k.nodeName
		}
		md := ci.ConvertToOTLPMetrics(fields, attributes, k.logger)
		result = append(result, md)

		for service, podNum := range client.Ep.ServiceToPodNum() {
			fields := map[string]interface{}{
				"service_number_of_running_pods": podNum,
			}
			attributes := map[string]string{
				ci.ClusterNameKey: clusterName,
				ci.MetricType:     ci.TypeClusterService,
				ci.Timestamp:      timestampNs,
				ci.TypeService:    service.ServiceName,
				ci.K8sNamespace:   service.Namespace,
				ci.Version:        "0",
			}
			if k.nodeName != "" {
				attributes["NodeName"] = k.nodeName
			}
			md := ci.ConvertToOTLPMetrics(fields, attributes, k.logger)
			result = append(result, md)
		}

		for namespace, podNum := range client.Pod.NamespaceToRunningPodNum() {
			fields := map[string]interface{}{
				"namespace_number_of_running_pods": podNum,
			}
			attributes := map[string]string{
				ci.ClusterNameKey: clusterName,
				ci.MetricType:     ci.TypeClusterNamespace,
				ci.Timestamp:      timestampNs,
				ci.K8sNamespace:   namespace,
				ci.Version:        "0",
			}
			if k.nodeName != "" {
				attributes["NodeName"] = k.nodeName
			}
			md := ci.ConvertToOTLPMetrics(fields, attributes, k.logger)
			result = append(result, md)
		}
	}
	return result
}

func (k *K8sAPIServer) init() error {
	var ctx context.Context
	ctx, k.cancel = context.WithCancel(context.Background())

	k.nodeName = os.Getenv("HOST_NAME")
	if k.nodeName == "" {
		return errors.New("missing environment variable HOST_NAME. Please check your deployment YAML config")
	}

	lockNamespace := os.Getenv("K8S_NAMESPACE")
	if lockNamespace == "" {
		return errors.New("missing environment variable K8S_NAMESPACE. Please check your deployment YAML config")
	}

	configMapInterface := k.k8sClient.ClientSet.CoreV1().ConfigMaps(lockNamespace)
	if configMap, err := configMapInterface.Get(ctx, lockName, metav1.GetOptions{}); configMap == nil || err != nil {
		k.logger.Info(fmt.Sprintf("Cannot get the leader config map: %v, try to create the config map...", err))
		configMap, err = configMapInterface.Create(ctx,
			&v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: lockNamespace,
					Name:      lockName,
				},
			}, metav1.CreateOptions{})
		k.logger.Info(fmt.Sprintf("configMap: %v, err: %v", configMap, err))
	}

	lock, err := resourcelock.New(
		resourcelock.ConfigMapsResourceLock,
		lockNamespace, lockName,
		k.k8sClient.ClientSet.CoreV1(),
		k.k8sClient.ClientSet.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity:      k.nodeName,
			EventRecorder: k.createRecorder(lockName, lockNamespace),
		})
	if err != nil {
		k.logger.Warn("Failed to create resource lock", zap.Error(err))
		return err
	}

	go k.startLeaderElection(ctx, lock)

	return nil
}

// Shutdown stops the k8sApiServer
func (k *K8sAPIServer) Shutdown() {
	if k.cancel != nil {
		k.cancel()
	}
}

func (k *K8sAPIServer) startLeaderElection(ctx context.Context, lock resourcelock.Interface) {

	for {
		leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
			Lock: lock,
			// IMPORTANT: you MUST ensure that any code you have that
			// is protected by the lease must terminate **before**
			// you call cancel. Otherwise, you could have a background
			// loop still running and another process could
			// get elected before your background loop finished, violating
			// the stated goal of the lease.
			LeaseDuration: 60 * time.Second,
			RenewDeadline: 15 * time.Second,
			RetryPeriod:   5 * time.Second,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStartedLeading: func(ctx context.Context) {
					k.logger.Info(fmt.Sprintf("k8sapiserver OnStartedLeading: %s", k.nodeName))
					// we're notified when we start
					k.leading = true

					if k.isLeadingC != nil {
						// this executes only in testing
						close(k.isLeadingC)
					}
				},
				OnStoppedLeading: func() {
					k.logger.Info(fmt.Sprintf("k8sapiserver OnStoppedLeading: %s", k.nodeName))
					// we can do cleanup here, or after the RunOrDie method returns
					k.leading = false
					//node and pod are only used for cluster level metrics, endpoint is used for decorator too.
					k.k8sClient.Node.Shutdown()
					k.k8sClient.Pod.Shutdown()
				},
				OnNewLeader: func(identity string) {
					k.logger.Info(fmt.Sprintf("k8sapiserver Switch New Leader: %s", identity))
				},
			},
		})

		select {
		case <-ctx.Done(): //when leader election ends, the channel ctx.Done() will be closed
			k.logger.Info(fmt.Sprintf("k8sapiserver shutdown Leader Election: %s", k.nodeName))
			return
		default:
		}
	}
}

func (k *K8sAPIServer) createRecorder(name, namespace string) record.EventRecorder {
	k.broadcaster.StartLogging(klog.Infof)
	clientSet := k.k8sClient.ClientSet
	k.broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(clientSet.CoreV1().RESTClient()).Events(namespace)})
	return k.broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: name})
}
