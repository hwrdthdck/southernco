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

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.opentelemetry.io/collector/component"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/translator/conventions"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

const (
	TypeStr = "ec2"
	tagPrefix = "ec2.tag."
)

var _ internal.Detector = (*Detector)(nil)

type Detector struct {
	metadataProvider metadataProvider
	cfg Config
}

func NewDetector(_ component.ProcessorCreateParams, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return &Detector{metadataProvider: newMetadataClient(sess), cfg: cfg}, nil
}

func (d *Detector) Detect(ctx context.Context) (pdata.Resource, error) {
	res := pdata.NewResource()
	if !d.metadataProvider.available(ctx) {
		return res, nil
	}

	meta, err := d.metadataProvider.get(ctx)
	if err != nil {
		return res, fmt.Errorf("failed getting identity document: %w", err)
	}

	hostname, err := d.metadataProvider.hostname(ctx)
	if err != nil {
		return res, fmt.Errorf("failed getting hostname: %w", err)
	}

	attr := res.Attributes()
	attr.InsertString(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderAWS)
	attr.InsertString("cloud.infrastructure_service", "EC2")
	attr.InsertString(conventions.AttributeCloudRegion, meta.Region)
	attr.InsertString(conventions.AttributeCloudAccount, meta.AccountID)
	attr.InsertString(conventions.AttributeCloudZone, meta.AvailabilityZone)
	attr.InsertString(conventions.AttributeHostID, meta.InstanceID)
	attr.InsertString(conventions.AttributeHostImageID, meta.ImageID)
	attr.InsertString(conventions.AttributeHostType, meta.InstanceType)
	attr.InsertString(conventions.AttributeHostName, hostname)

	tags, err := fetchEc2Tags(meta.Region, meta.InstanceID, d.cfg)
	if err != nil {
		return res, err
	}

	for key, val := range tags {
		attr.InsertString(tagPrefix + key, val)
	}

	return res, nil
}

func fetchEc2Tags(region string, instanceID string, cfg Config) (map[string]string, error) {
	if(!cfg.AddAllTags && len(cfg.TagsToAdd) == 0){
		fmt.Println("skipped ec2 tags")
		return nil, nil
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}

	e := ec2.New(sess)
	if _, err = sess.Config.Credentials.Get(); err != nil {
		return nil, err
	}

	ec2Tags, err := e.DescribeTags(&ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{{
			Name: aws.String("resource-id"),
			Values: []*string{
				aws.String(instanceID),
			},
		}},
	})
	if err != nil {
		return nil, err
	}
	tags := make(map[string]string)
	for _, tag := range ec2Tags.Tags {
		if cfg.AddAllTags || contains(cfg.TagsToAdd, *tag.Key) {
			tags[*tag.Key]= *tag.Value
		}
	}
	return tags, nil
}

func contains(arr []string, val string) bool {
	for _, elem := range arr {
		if val == elem {
			return true
		}
	}
	return false
}