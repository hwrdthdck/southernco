// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aks // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure/aks"

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/azure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure/aks/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "aks"

	// Environment variable that is set when running on Kubernetes
	kubernetesServiceHostEnvVar = "KUBERNETES_SERVICE_HOST"
)

type Detector struct {
	provider           azure.Provider
	resourceAttributes metadata.ResourceAttributesConfig
}

// NewDetector creates a new AKS detector
func NewDetector(_ processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	return &Detector{provider: azure.NewProvider(), resourceAttributes: cfg.ResourceAttributes}, nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	res := pcommon.NewResource()

	if !onK8s() {
		return res, "", nil
	}

	// If we can't get a response from the metadata endpoint, we're not running in Azure
	if !azureMetadataAvailable(ctx, d.provider) {
		return res, "", nil
	}

	attrs := res.Attributes()
	if d.resourceAttributes.CloudProvider.Enabled {
		attrs.PutStr(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderAzure)
	}
	if d.resourceAttributes.CloudPlatform.Enabled {
		attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformAzureAKS)

		if d.resourceAttributes.K8sClusterName.Enabled {
			m, err := d.provider.Metadata(ctx)
			if err != nil {
				return res, "", fmt.Errorf("failed to get IMDS metadata: %w", err)
			}
			attrs.PutStr(conventions.AttributeK8SClusterName, parseClusterName(m.ResourceGroupName))
		}
	}

	return res, conventions.SchemaURL, nil
}

func onK8s() bool {
	return os.Getenv(kubernetesServiceHostEnvVar) != ""
}

func azureMetadataAvailable(ctx context.Context, p azure.Provider) bool {
	_, err := p.Metadata(ctx)
	return err == nil
}

// parseClusterName parses the cluster name from the infrastructure
// resource group name. AKS IMDS returns the resource group name in
// the following formats:
//
// 1. Generated group: MC_<resource group>_<cluster name>_<location>
//   - Example:
//   - Resource group: my-resource-group
//   - Cluster name:   my-cluster
//   - Location:       eastus
//   - Generated name: MC_my-resource-group_my-cluster_eastus
//
// 2. Custom group: custom-infra-resource-group-name
//
// When using the generated infrastructure resource group, the resource
// group will include the cluster name. If the cluster's resource group
// or cluster name contains underscores, parsing will fall back on the
// unparsed infrastructure resource group name.
//
// When using a custom infrastructure resource group, the resource group name
// does not contain the cluster name. The custom infrastructure resource group
// name is returned instead.
//
// It is safe to use the infrastructure resource group name as a unique identifier
// because Azure will not allow the user to create multiple AKS clusters with the same
// infrastructure resource group name.
func parseClusterName(resourceGroup string) string {
	// Code inspired by https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/exporter/datadogexporter/internal/hostmetadata/internal/azure/provider.go#L36
	splitAll := strings.Split(resourceGroup, "_")

	if len(splitAll) == 4 && strings.ToLower(splitAll[0]) == "mc" {
		return splitAll[len(splitAll)-2]
	}

	return resourceGroup
}
