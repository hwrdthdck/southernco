// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("k8s_cluster")
	ScopeName = "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"
)

const (
	LogsStability    = component.StabilityLevelDevelopment
	MetricsStability = component.StabilityLevelBeta
)
