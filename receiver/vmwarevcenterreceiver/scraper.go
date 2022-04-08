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

package vmwarevcenterreceiver // import github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vmwarevcenterreceiver

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vsan/types"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vmwarevcenterreceiver/internal/metadata"
)

// example 2022-03-10 14:15:00
const timeFormat = "2006-01-02 15:04:05"

var _ component.Receiver = (*vcenterMetricScraper)(nil)

type vcenterMetricScraper struct {
	client      *VmwareVcenterClient
	mb          *metadata.MetricsBuilder
	vsanEnabled bool
}

func newVmwareVcenterScraper(
	logger *zap.Logger,
	config *Config,
) *vcenterMetricScraper {
	l := logger.Named("vcenter-client")
	client := newVmwarevcenterClient(config, l)
	return &vcenterMetricScraper{
		client:      client,
		mb:          metadata.NewMetricsBuilder(config.MetricsConfig.Metrics),
		vsanEnabled: true,
	}
}

func (v *vcenterMetricScraper) Start(ctx context.Context, _ component.Host) error {
	return v.client.Connect(ctx)
}

func (v *vcenterMetricScraper) Shutdown(ctx context.Context) error {
	return v.client.Disconnect(ctx)
}

func (v *vcenterMetricScraper) scrape(ctx context.Context) (pdata.Metrics, error) {
	if v.client == nil {
		return pdata.Metrics{}, errors.New("failed to connect to http client")
	}

	errs := &scrapererror.ScrapeErrors{}
	err := v.client.ConnectVSAN(ctx)
	if err != nil {
		// vsan is not required for a proper collection
		errs.AddPartial(1, err)
		v.vsanEnabled = false
	}
	v.collectClusters(ctx, errs)
	return v.mb.Emit(), errs.Combine()
}

func (v *vcenterMetricScraper) collectClusters(ctx context.Context, errs *scrapererror.ScrapeErrors) error {
	clusters, err := v.client.Clusters(ctx)
	if err != nil {
		return err
	}
	now := pdata.NewTimestampFromTime(time.Now())

	for _, c := range clusters {
		v.collectCluster(ctx, now, c, errs)
		v.collectHosts(ctx, now, c, errs)
		v.collectDatastores(ctx, now, c, errs)
		v.collectVMs(ctx, now, c, errs)
	}
	v.collectResourcePools(ctx, now, errs)

	return nil
}

func (v *vcenterMetricScraper) collectCluster(
	ctx context.Context,
	now pdata.Timestamp,
	c *object.ClusterComputeResource,
	errs *scrapererror.ScrapeErrors,
) {
	var moCluster mo.ClusterComputeResource
	c.Properties(ctx, c.Reference(), []string{"summary"}, &moCluster)
	s := moCluster.Summary.GetComputeResourceSummary()
	v.mb.RecordVcenterClusterCPUAvailableDataPoint(now, int64(s.TotalCpu))

	mor := c.Reference()
	csvs, err := v.client.CollectVSANCluster(ctx, &mor, time.Now().UTC(), time.Now().UTC())
	if err != nil {
		errs.AddPartial(1, err)
	}
	if csvs != nil {
		v.addVSANMetrics(*csvs, "*", clusterType, errs)
	}
	v.mb.EmitCluster(c.Name())
}

func (v *vcenterMetricScraper) collectDatastores(
	ctx context.Context,
	colTime pdata.Timestamp,
	cluster *object.ClusterComputeResource,
	errs *scrapererror.ScrapeErrors,
) {
	datastores, err := cluster.Datastores(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}

	for _, ds := range datastores {
		v.collectDatastore(ctx, colTime, ds)
	}
}

func (v *vcenterMetricScraper) collectDatastore(
	ctx context.Context,
	now pdata.Timestamp,
	ds *object.Datastore,
) {
	var moDS mo.Datastore
	ds.Properties(ctx, ds.Reference(), []string{"summary", "name"}, &moDS)

	v.recordDatastoreProperties(now, moDS)
	v.mb.EmitDatastore(moDS.Name)
}

func (v *vcenterMetricScraper) collectHosts(
	ctx context.Context,
	colTime pdata.Timestamp,
	cluster *object.ClusterComputeResource,
	errs *scrapererror.ScrapeErrors,
) {
	hosts, err := cluster.Hosts(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}

	clusterRef := cluster.Reference()
	var hostVsanCSVs *[]types.VsanPerfEntityMetricCSV
	if v.vsanEnabled {
		hostVsanCSVs, err = v.client.CollectVSANHosts(ctx, &clusterRef, time.Now(), time.Now())
		if err != nil {
			errs.AddPartial(1, err)
		}
	}

	for _, h := range hosts {
		v.collectHost(ctx, colTime, h, hostVsanCSVs, errs)
		v.mb.EmitHost(h.Name())
	}
}

func (v *vcenterMetricScraper) collectHost(
	ctx context.Context,
	now pdata.Timestamp,
	host *object.HostSystem,
	vsanCsvs *[]types.VsanPerfEntityMetricCSV,
	errs *scrapererror.ScrapeErrors,
) {
	var hwSum mo.HostSystem
	err := host.Properties(ctx, host.Reference(),
		[]string{
			"config",
			"summary.hardware",
			"summary.quickStats",
		}, &hwSum)

	if err != nil {
		errs.AddPartial(1, err)
		return
	}
	v.recordHostSystemMemoryUsage(now, hwSum)
	v.recordHostPerformanceMetrics(ctx, hwSum, time.Now().Add(-15*time.Minute).UTC(), time.Now().Add(-5*time.Minute).UTC(), errs)
	if vsanCsvs != nil {
		entityRef := fmt.Sprintf("host-domclient:%v",
			hwSum.Config.VsanHostConfig.ClusterInfo.NodeUuid,
		)
		v.addVSANMetrics(*vsanCsvs, entityRef, hostType, errs)
	}
}

func (v *vcenterMetricScraper) collectResourcePools(
	ctx context.Context,
	ts pdata.Timestamp,
	errs *scrapererror.ScrapeErrors,
) {
	rps, err := v.client.ResourcePools(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}
	for _, rp := range rps {
		var moRP mo.ResourcePool
		rp.Properties(ctx, rp.Reference(), []string{
			"summary",
			"summary.quickStats",
			"name",
		}, &moRP)
		v.recordResourcePool(ts, moRP)
		v.mb.EmitResourcePool(rp.Name())
	}
}

func (v *vcenterMetricScraper) collectVMs(
	ctx context.Context,
	colTime pdata.Timestamp,
	cluster *object.ClusterComputeResource,
	errs *scrapererror.ScrapeErrors,
) {
	vms, err := v.client.VMs(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}

	var vsanCsvs *[]types.VsanPerfEntityMetricCSV
	if v.vsanEnabled {
		clusterRef := cluster.Reference()
		vsanCsvs, err = v.client.CollectVSANVirtualMachine(ctx, &clusterRef, time.Now().UTC(), time.Now().UTC())
		if err != nil {
			errs.AddPartial(1, err)
		}
	}

	for _, vm := range vms {
		var moVM mo.VirtualMachine
		err := vm.Properties(ctx, vm.Reference(), []string{
			"config",
			"runtime",
			"summary",
		}, &moVM)

		if err != nil {
			errs.AddPartial(1, err)
			continue
		}

		vmUUID := moVM.Config.InstanceUuid
		entityRefID := fmt.Sprintf("virtual-machine:%s", vmUUID)
		ps := string(moVM.Runtime.PowerState)

		v.collectVM(ctx, colTime, moVM, entityRefID, vsanCsvs, errs)
		v.mb.EmitVM(vm.Name(), vmUUID, ps)
	}

}

func (v *vcenterMetricScraper) collectVM(
	ctx context.Context,
	colTime pdata.Timestamp,
	vm mo.VirtualMachine,
	entityRefID string,
	vsanCsvs *[]types.VsanPerfEntityMetricCSV,
	errs *scrapererror.ScrapeErrors,
) {
	v.recordVMUsages(colTime, vm)
	v.recordVMPerformance(ctx,
		vm,
		// start time
		time.Now().Add(-15*time.Minute).UTC(),
		// end time
		time.Now().Add(-5*time.Minute).UTC(),
		errs,
	)

	if vsanCsvs != nil {
		v.addVSANMetrics(*vsanCsvs, entityRefID, vmType, errs)
	}
}

type vsanType int

const (
	clusterType vsanType = iota
	hostType
	vmType
)

func (v *vcenterMetricScraper) addVSANMetrics(
	csvs []types.VsanPerfEntityMetricCSV,
	entityID string,
	vsanType vsanType,
	errs *scrapererror.ScrapeErrors,
) {
	for _, r := range csvs {
		// can't correlate this point to a timestamp so just skip it
		if r.SampleInfo == "" {
			continue
		}
		// If not this entity ID, then skip it
		if vsanType != clusterType && r.EntityRefId != entityID {
			continue
		}

		time, err := time.Parse(timeFormat, r.SampleInfo)
		if err != nil {
			errs.AddPartial(1, err)
			continue
		}

		ts := pdata.NewTimestampFromTime(time)
		for _, val := range r.Value {
			values := strings.Split(val.Values, ",")
			for _, value := range values {
				switch vsanType {
				case clusterType:
					v.recordClusterVsanMetric(ts, val.MetricId.Label, value, errs)
				case hostType:
					v.recordHostVsanMetric(ts, val.MetricId.Label, value, errs)
				case vmType:
					v.recordVMVsanMetric(ts, val.MetricId.Label, value, errs)
				}
			}
		}
	}
}
