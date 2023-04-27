// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awscloudwatchlogsexporter

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/multierr"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()

	tests := []struct {
		id           component.ID
		expected     component.Config
		errorMessage string
	}{
		{
			id: component.NewIDWithName(typeStr, "e1-defaults"),
			expected: &Config{
				RetrySettings:      defaultRetrySettings,
				LogGroupName:       "test-1",
				LogStreamName:      "testing",
				Endpoint:           "",
				AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
				QueueSettings: QueueSettings{
					QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
				},
			},
		},
		{
			id: component.NewIDWithName(typeStr, "e2-no-retries-short-queue"),
			expected: &Config{
				RetrySettings: exporterhelper.RetrySettings{
					Enabled:             false,
					InitialInterval:     defaultRetrySettings.InitialInterval,
					MaxInterval:         defaultRetrySettings.MaxInterval,
					MaxElapsedTime:      defaultRetrySettings.MaxElapsedTime,
					RandomizationFactor: backoff.DefaultRandomizationFactor,
					Multiplier:          backoff.DefaultMultiplier,
				},
				AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
				LogGroupName:       "test-2",
				LogStreamName:      "testing",
				QueueSettings: QueueSettings{
					QueueSize: 2,
				},
			},
		},
		{
			id:           component.NewIDWithName(typeStr, "invalid_queue_size"),
			errorMessage: "'sending_queue.queue_size' must be 1 or greater",
		},
		{
			id:           component.NewIDWithName(typeStr, "invalid_required_field_stream"),
			errorMessage: "'log_stream_name' must be set",
		},
		{
			id:           component.NewIDWithName(typeStr, "invalid_required_field_group"),
			errorMessage: "'log_group_name' must be set",
		},
		{
			id:           component.NewIDWithName(typeStr, "invalid_queue_setting"),
			errorMessage: `'sending_queue' has invalid keys: enabled, num_consumers`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			err = component.UnmarshalConfig(sub, cfg)

			if tt.expected == nil {
				err = multierr.Append(err, component.ValidateConfig(cfg))
				assert.ErrorContains(t, err, tt.errorMessage)
				return
			}
			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestRetentionValidateCorrect(t *testing.T) {
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	cfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		LogRetention:       365,
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.NoError(t, component.ValidateConfig(cfg))

}

func TestRetentionValidateWrong(t *testing.T) {
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		LogRetention:       366,
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.Error(t, component.ValidateConfig(wrongcfg))

}

func TestTagsValidateCorrect(t *testing.T) {
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	avalue := "avalue"
	cfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{"akey": &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.NoError(t, component.ValidateConfig(cfg))

}

func TestTagsValidateTooLittleTags(t *testing.T) {
	m := make(map[string]*string)
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               m,
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.EqualError(t, component.ValidateConfig(wrongcfg), "invalid amount of items. Please input at least 1 tag or remove the tag field")
}

func TestTagsValidateTooManyTags(t *testing.T) {
	m := make(map[string]*string)
	avalue := "avalue"
	for i := 0; i < 51; i++ {
		m[strconv.Itoa(i)] = &avalue
	}
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               m,
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.EqualError(t, component.ValidateConfig(wrongcfg), "invalid amount of items. Please input at most 50 tags")
}

func TestTagsValidateWrongKeyRegex(t *testing.T) {
	avalue := "avalue"
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{"***": &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.EqualError(t, component.ValidateConfig(wrongcfg), "key - *** does not follow the regex pattern"+`^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`)
}

func TestTagsValidateWrongValueRegex(t *testing.T) {
	avalue := "***"
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{"akey": &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.EqualError(t, component.ValidateConfig(wrongcfg), "value - "+avalue+" does not follow the regex pattern"+`^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`)
}

func TestTagsValidateKeyTooShort(t *testing.T) {
	avalue := "avalue"
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{"": &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.EqualError(t, component.ValidateConfig(wrongcfg), "key -  has an invalid length. Please use keys with a length of 1 to 128 characters")
}

func TestTagsValidateKeyTooLong(t *testing.T) {
	avalue := "avalue"
	akey := strings.Repeat("a", 129)
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{akey: &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.Error(t, component.ValidateConfig(wrongcfg), "key - "+akey+" has an invalid length. Please use keys with a length of 1 to 128 characters")
}

func TestTagsValidateValueTooShort(t *testing.T) {
	avalue := ""
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{"akey": &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.Error(t, component.ValidateConfig(wrongcfg), "value - "+avalue+" has an invalid length. Please use keys with a length of 1 to 256 characters")
}

func TestTagsValidateValueTooLong(t *testing.T) {
	avalue := strings.Repeat("a", 257)
	defaultRetrySettings := exporterhelper.NewDefaultRetrySettings()
	wrongcfg := &Config{
		RetrySettings:      defaultRetrySettings,
		LogGroupName:       "test-1",
		LogStreamName:      "testing",
		Endpoint:           "",
		Tags:               map[string]*string{"akey": &avalue},
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
	assert.Error(t, component.ValidateConfig(wrongcfg), "value - "+avalue+" has an invalid length. Please use keys with a length of 1 to 256 characters")
}
