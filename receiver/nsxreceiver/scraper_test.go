// Copyright  The OpenTelemetry Authors
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

package nsxreceiver

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/scrapertest/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxreceiver/internal/metadata"
	dm "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxreceiver/internal/model"
)

func TestScrape(t *testing.T) {
	mockClient := NewMockClient(t)

	mockClient.On("ClusterNodes", mock.Anything).Return(loadTestClusterNodes())
	mockClient.On("TransportNodes", mock.Anything).Return(loadTestTransportNodes())

	mockClient.On("NodeStatus", mock.Anything, transportNode1, transportClass).Return(loadTestNodeStatus(t, transportNode1, transportClass))
	mockClient.On("NodeStatus", mock.Anything, transportNode2, transportClass).Return(loadTestNodeStatus(t, transportNode2, transportClass))
	mockClient.On("NodeStatus", mock.Anything, transportNode2, transportClass).Return(loadTestNodeStatus(t, transportNode2, transportClass))
	mockClient.On("NodeStatus", mock.Anything, managerNode1, managerClass).Return(loadTestNodeStatus(t, managerNode1, managerClass))

	mockClient.On("Interfaces", mock.Anything, managerNode1, managerClass).Return(loadTestNodeInterfaces(t, managerNode1, managerClass))
	mockClient.On("Interfaces", mock.Anything, transportNode1, transportClass).Return(loadTestNodeInterfaces(t, transportNode1, transportClass))
	mockClient.On("Interfaces", mock.Anything, transportNode2, transportClass).Return(loadTestNodeInterfaces(t, transportNode2, transportClass))

	mockClient.On("InterfaceStatus", mock.Anything, transportNode1, transportNodeNic1, transportClass).Return(loadInterfaceStats(t, transportNode1, transportNodeNic1, transportClass))
	mockClient.On("InterfaceStatus", mock.Anything, transportNode1, transportNodeNic2, transportClass).Return(loadInterfaceStats(t, transportNode1, transportNodeNic2, transportClass))
	mockClient.On("InterfaceStatus", mock.Anything, transportNode2, transportNodeNic1, transportClass).Return(loadInterfaceStats(t, transportNode2, transportNodeNic1, transportClass))
	mockClient.On("InterfaceStatus", mock.Anything, transportNode2, transportNodeNic2, transportClass).Return(loadInterfaceStats(t, transportNode2, transportNodeNic2, transportClass))
	mockClient.On("InterfaceStatus", mock.Anything, managerNode1, managerNodeNic1, managerClass).Return(loadInterfaceStats(t, managerNode1, managerNodeNic1, managerClass))
	mockClient.On("InterfaceStatus", mock.Anything, managerNode1, managerNodeNic2, managerClass).Return(loadInterfaceStats(t, managerNode1, managerNodeNic2, managerClass))

	scraper := newScraper(
		&Config{
			MetricsConfig: &MetricsConfig{Settings: metadata.DefaultMetricsSettings()},
		},
		componenttest.NewNopTelemetrySettings(),
	)
	scraper.client = mockClient

	metrics, err := scraper.scrape(context.Background())
	require.NoError(t, err)

	expectedMetrics, err := golden.ReadMetrics(filepath.Join("testdata", "metrics", "expected_metrics.json"))
	require.NoError(t, err)

	err = scrapertest.CompareMetrics(metrics, expectedMetrics)
	require.NoError(t, err)
}

func loadTestNodeStatus(t *testing.T, nodeID string, class nodeClass) (*dm.NodeStatus, error) {
	var classType string
	switch class {
	case transportClass:
		classType = "transport"
	default:
		classType = "cluster"
	}
	testFile, err := ioutil.ReadFile(filepath.Join("testdata", "metrics", "nodes", classType, nodeID, "status.json"))
	require.NoError(t, err)
	var stats dm.NodeStatus
	err = json.Unmarshal(testFile, &stats)
	require.NoError(t, err)
	return &stats, nil
}

func loadTestNodeInterfaces(t *testing.T, nodeID string, class nodeClass) ([]dm.NetworkInterface, error) {
	var classType string
	switch class {
	case transportClass:
		classType = "transport"
	default:
		classType = "cluster"
	}
	testFile, err := ioutil.ReadFile(filepath.Join("testdata", "metrics", "nodes", classType, nodeID, "interfaces", "index.json"))
	require.NoError(t, err)
	var interfaces dm.NodeNetworkInterfacePropertiesListResult
	err = json.Unmarshal(testFile, &interfaces)
	require.NoError(t, err)
	return interfaces.Results, nil
}

func loadInterfaceStats(t *testing.T, nodeID, interfaceID string, class nodeClass) (*dm.NetworkInterfaceStats, error) {
	var classType string
	switch class {
	case transportClass:
		classType = "transport"
	default:
		classType = "cluster"
	}
	testFile, err := ioutil.ReadFile(filepath.Join("testdata", "metrics", "nodes", classType, nodeID, "interfaces", interfaceID, "stats.json"))
	require.NoError(t, err)
	var stats dm.NetworkInterfaceStats
	err = json.Unmarshal(testFile, &stats)
	require.NoError(t, err)
	return &stats, nil
}

func loadTestClusterNodes() ([]dm.ClusterNode, error) {
	testFile, err := ioutil.ReadFile(filepath.Join("testdata", "metrics", "cluster_nodes.json"))
	if err != nil {
		return nil, err
	}
	var nodes dm.ClusterNodeList
	err = json.Unmarshal(testFile, &nodes)
	return nodes.Results, err
}

func loadTestTransportNodes() ([]dm.TransportNode, error) {
	testFile, err := ioutil.ReadFile(filepath.Join("testdata", "metrics", "transport_nodes.json"))
	if err != nil {
		return nil, err
	}
	var nodes dm.TransportNodeList
	err = json.Unmarshal(testFile, &nodes)
	return nodes.Results, err
}
