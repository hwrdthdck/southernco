package ecs

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"go.opentelemetry.io/collector/consumer/pdata"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockMetaDataProvider struct {
	isV4 bool
}

var _ ecsMetadataProvider = (*mockMetaDataProvider)(nil)

func (md *mockMetaDataProvider) fetchTaskMetaData(tmde string) (*TaskMetaData, error) {
	c := createTestContainer(md.isV4)
	c.DockerId = "05281997" // Simulate one "application" and one "collector" container
	cs := []Container{createTestContainer(md.isV4), c}
	tmd := &TaskMetaData{
		Cluster:          "cluster",
		TaskARN:          "arn:aws:ecs:us-west-2:123456789123:task/123",
		Family:           "family",
		AvailabilityZone: "us-west-2a",
		Containers:       cs,
	}

	if md.isV4 {
		tmd.LaunchType = "EC2"
	}

	return tmd, nil
}

func (md *mockMetaDataProvider) fetchContainerMetaData(tmde string) (*Container, error) {
	c := createTestContainer(md.isV4)
	return &c, nil
}

func Test_ecsNewDetector(t *testing.T) {
	d, err := NewDetector()

	assert.NotNil(t, d)
	assert.Nil(t, err)
}

func Test_detectorReturnsIfNoEnvVars(t *testing.T) {
	os.Clearenv()
	d, _ := NewDetector()
	res, err := d.Detect(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, 0, res.Attributes().Len())
}

func Test_ecsPrefersLatestTmde(t *testing.T) {
	os.Clearenv()
	os.Setenv(tmde3EnvVar, "3")
	os.Setenv(tmde4EnvVar, "4")

	tmde := getTmdeFromEnv()

	assert.Equal(t, "4", tmde)
}

func Test_ecsFiltersInvalidContainers(t *testing.T) {
	// Should ignore empty container
	c1 := Container{}

	// Should ignore non-normal container
	c2 := createTestContainer(true)
	c2.Type = "INTERNAL"

	// Should ignore stopped containers
	c3 := createTestContainer(true)
	c3.KnownStatus = "STOPPED"

	// Should ignore its own container
	c4 := createTestContainer(true)

	containers := []Container{c1, c2, c3, c4}

	ld := getValidLogData(containers, &c4, "123")

	for _, attrib := range ld {
		assert.Equal(t, 0, attrib.Len())
	}
}

func Test_ecsDetectV4(t *testing.T) {
	os.Clearenv()
	os.Setenv(tmde4EnvVar, "endpoint")

	want := pdata.NewResource()
	want.InitEmpty()
	attr := want.Attributes()
	attr.InsertString("cloud.provider", "aws")
	attr.InsertString("cloud.infrastructure_service", "ECS")
	attr.InsertString("aws.ecs.cluster", "cluster")
	attr.InsertString("aws.ecs.task.arn", "arn:aws:ecs:us-west-2:123456789123:task/123")
	attr.InsertString("aws.ecs.task.family", "family")
	attr.InsertString("cloud.region", "us-west-2")
	attr.InsertString("cloud.zone", "us-west-2a")
	attr.InsertString("cloud.account.id", "123456789123")
	attr.InsertString("aws.ecs.launchtype", "EC2")

	attribFields := []string{"aws.log.group.names", "aws.log.group.arns", "aws.log.stream.names", "aws.log.stream.arns"}
	attribVals := []string{"group", "arn:aws:logs:us-east-1:123456789123:log-group:group", "stream", "arn:aws:logs:us-east-1:123456789123:log-group:group:log-stream:stream"}

	for i, field := range attribFields {
		av := pdata.NewAnyValueArray()
		av.Append(pdata.NewAttributeValueString(attribVals[i]))
		ava := pdata.NewAttributeValueArray()
		ava.SetArrayVal(av)
		attr.Insert(field, ava)
	}

	d := Detector{provider: &mockMetaDataProvider{isV4: true}}
	got, err := d.Detect(context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, internal.AttributesToMap(want.Attributes()), internal.AttributesToMap(got.Attributes()))
}

func Test_ecsDetectV3(t *testing.T) {
	os.Clearenv()
	os.Setenv(tmde3EnvVar, "endpoint")

	want := pdata.NewResource()
	want.InitEmpty()
	attr := want.Attributes()
	attr.InsertString("cloud.provider", "aws")
	attr.InsertString("cloud.infrastructure_service", "ECS")
	attr.InsertString("aws.ecs.cluster", "cluster")
	attr.InsertString("aws.ecs.task.arn", "arn:aws:ecs:us-west-2:123456789123:task/123")
	attr.InsertString("aws.ecs.task.family", "family")
	attr.InsertString("cloud.region", "us-west-2")
	attr.InsertString("cloud.zone", "us-west-2a")
	attr.InsertString("cloud.account.id", "123456789123")

	d := Detector{provider: &mockMetaDataProvider{isV4: false}}
	got, err := d.Detect(context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, internal.AttributesToMap(want.Attributes()), internal.AttributesToMap(got.Attributes()))
}

func createTestContainer(isV4 bool) Container {
	c := Container{
		DockerId:    "123",
		Type:        "NORMAL",
		KnownStatus: "RUNNING",
	}

	if isV4 {
		c.LogDriver = "awslogs"
		c.ContainerARN = "arn:aws:ecs"
		c.LogOptions = LogData{LogGroup: "group", Region: "us-east-1", Stream: "stream"}
	}

	return c
}
