// Copyright The OpenTelemetry Authors
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

package gcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/internal/gcp"

import (
	"context"
	"fmt"
	"strings"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata/provider"

	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"
)

var _ provider.HostnameProvider = (*Provider)(nil)

var _ gcpDetector = gcp.NewDetector()

type gcpDetector interface {
	ProjectID() (string, error)
	CloudPlatform() gcp.Platform
	GCEHostName() (string, error)
}

type Provider struct {
	detector gcpDetector
}

// Hostname returns the GCP cloud integration hostname.
func (p *Provider) Hostname(context.Context) (string, error) {
	if p.detector.CloudPlatform() != gcp.GCE {
		return "", fmt.Errorf("not on Google Cloud Engine")
	}

	name, err := p.detector.GCEHostName()
	if err != nil {
		return "", fmt.Errorf("failed to get instance name: %w", err)
	}

	// Use the same logic as in the metadata from attributes logic.
	if strings.Count(name, ".") >= 3 {
		name = strings.SplitN(name, ".", 2)[0]
	}

	cloudAccount, err := p.detector.ProjectID()
	if err != nil {
		return "", fmt.Errorf("failed to get project ID: %w", err)
	}

	return fmt.Sprintf("%s.%s", name, cloudAccount), nil
}

// NewProvider creates a new GCP hostname provider.
func NewProvider() *Provider {
	return &Provider{detector: gcp.NewDetector()}
}
