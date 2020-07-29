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
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer/pdata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

type mockMetadata struct {
	ret       ec2metadata.EC2InstanceIdentityDocument
	returnErr error
}

func (mm mockMetadata) get(ctx context.Context) (ec2metadata.EC2InstanceIdentityDocument, error) {
	if mm.returnErr != nil {
		return ec2metadata.EC2InstanceIdentityDocument{}, mm.returnErr
	}
	return mm.ret, nil
}

func TestDetector_Detect(t *testing.T) {
	type fields struct {
		provider ec2MetadataProvider
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    pdata.Resource
		wantErr bool
	}{
		{"success", fields{provider: &mockMetadata{ec2metadata.EC2InstanceIdentityDocument{
			Region:           "us-west-2",
			AccountID:        "account1234",
			AvailabilityZone: "us-west-2a",
			InstanceID:       "i-abcd1234",
			ImageID:          "abcdef",
			InstanceType:     "c4.xlarge",
		}, nil}}, args{ctx: context.Background()}, func() pdata.Resource {
			res := pdata.NewResource()
			res.InitEmpty()
			attr := res.Attributes()
			attr.InsertString("cloud.account.id", "account1234")
			attr.InsertString("cloud.provider", "aws")
			attr.InsertString("cloud.region", "us-west-2")
			attr.InsertString("cloud.zone", "us-west-2a")
			attr.InsertString("host.id", "i-abcd1234")
			attr.InsertString("host.image.id", "abcdef")
			attr.InsertString("host.type", "c4.xlarge")
			return res
		}(), false},
		{"endpoint not available", fields{provider: &mockMetadata{
			ret:       ec2metadata.EC2InstanceIdentityDocument{},
			returnErr: errors.New("metadata lookup failed"),
		}}, args{ctx: context.Background()}, func() pdata.Resource {
			res := pdata.NewResource()
			res.InitEmpty()
			attr := res.Attributes()
			attr.InsertString("cloud.provider", "aws")
			return res
		}(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Detector{
				provider: tt.fields.provider,
			}
			got, err := d.Detect(tt.args.ctx)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				require.False(t, got.IsNil())
				assert.Equal(t, internal.AttributesToMap(tt.want.Attributes()), internal.AttributesToMap(got.Attributes()))
			}
		})
	}
}
