// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubeadm // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/kubeadm"

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

type Provider interface {
	// ClusterName returns the current K8S Node Name
	ClusterName(ctx context.Context) (string, error)
}

type kubeadmProvider struct {
	kubeadmClient      kubernetes.Interface
	configMapName      string
	configMapNamespace string
}

func NewProvider(configMapName string, configMapNamespace string, apiConf k8sconfig.APIConfig) (Provider, error) {
	if configMapName == "" || configMapNamespace == "" {
		return nil, fmt.Errorf("configMapName and configMapNamespace can't be empty")
	}
	k8sAPIClient, err := k8sconfig.MakeClient(apiConf)
	if err != nil {
		return nil, fmt.Errorf("failed to create K8s API client: %w", err)
	}
	return &kubeadmProvider{
		kubeadmClient:      k8sAPIClient,
		configMapName:      configMapName,
		configMapNamespace: configMapNamespace,
	}, nil
}

func (k *kubeadmProvider) ClusterName(ctx context.Context) (string, error) {
	cluster, err := k.kubeadmClient.CoreV1().ConfigMaps(k.configMapNamespace).Get(ctx, k.configMapName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to fetch ConfigMap with name %s and namespace %s from K8s API: %w", k.configMapName, k.configMapNamespace, err)
	}
	return cluster.Name, nil
}
