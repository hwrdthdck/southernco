// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8stest"

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/util/net"
	"k8s.io/client-go/discovery"
	memory "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	DynamicClient   *dynamic.DynamicClient
	DiscoveryClient *discovery.DiscoveryClient
	Mapper          *restmapper.DeferredDiscoveryRESTMapper

	ctx                 context.Context
	cancel              context.CancelFunc
	discoveryHttpClient *http.Client
	dynamicHttpClient   *http.Client
}

func NewK8sClient(ctx context.Context, kubeconfigPath string) (*K8sClient, error) {
	if kubeconfigPath == "" {
		return nil, errors.New("Please provide file path to load kubeconfig")
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load kubeconfig from %s: %w", kubeconfigPath, err)
	}

	restConfig.Proxy = net.NewProxierWithNoProxyCIDR(http.ProxyFromEnvironment)
	dynamicHttpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, err
	}
	discoveryHttpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, err
	}

	dynamicClient, err := dynamic.NewForConfigAndClient(restConfig, dynamicHttpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating dynamic client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfigAndClient(restConfig, discoveryHttpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating discovery client: %w", err)
	}

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))

	k8sClient := &K8sClient{
		DynamicClient:       dynamicClient,
		DiscoveryClient:     discoveryClient,
		Mapper:              mapper,
		dynamicHttpClient:   dynamicHttpClient,
		discoveryHttpClient: discoveryHttpClient,
	}

	cctx, cancel := context.WithCancel(ctx)
	k8sClient.ctx = cctx
	k8sClient.cancel = cancel

	return k8sClient, nil
}

func (k *K8sClient) Shutdown() {
	if k.dynamicHttpClient != nil {
		net.CloseIdleConnectionsFor(k.dynamicHttpClient.Transport)
	}

	if k.discoveryHttpClient != nil {
		net.CloseIdleConnectionsFor(k.discoveryHttpClient.Transport)
	}

	if k.cancel != nil {
		k.cancel()
	}
}
