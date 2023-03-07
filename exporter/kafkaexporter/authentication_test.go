// Copyright The OpenTelemetry Authors
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

package kafkaexporter

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configtls"
)

func TestAuthentication(t *testing.T) {
	saramaPlaintext := &sarama.Config{}
	saramaPlaintext.Net.SASL.Enable = true
	saramaPlaintext.Net.SASL.User = "jdoe"
	saramaPlaintext.Net.SASL.Password = "pass"

	saramaSASLSCRAM256Config := &sarama.Config{}
	saramaSASLSCRAM256Config.Net.SASL.Enable = true
	saramaSASLSCRAM256Config.Net.SASL.User = "jdoe"
	saramaSASLSCRAM256Config.Net.SASL.Password = "pass"
	saramaSASLSCRAM256Config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

	saramaSASLSCRAM512Config := &sarama.Config{}
	saramaSASLSCRAM512Config.Net.SASL.Enable = true
	saramaSASLSCRAM512Config.Net.SASL.User = "jdoe"
	saramaSASLSCRAM512Config.Net.SASL.Password = "pass"
	saramaSASLSCRAM512Config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512

	saramaSASLPLAINConfig := &sarama.Config{}
	saramaSASLPLAINConfig.Net.SASL.Enable = true
	saramaSASLPLAINConfig.Net.SASL.User = "jdoe"
	saramaSASLPLAINConfig.Net.SASL.Password = "pass"

	saramaSASLPLAINConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	saramaTLSCfg := &sarama.Config{}
	saramaTLSCfg.Net.TLS.Enable = true
	tlsClient := configtls.TLSClientSetting{}
	tlscfg, err := tlsClient.LoadTLSConfig()
	require.NoError(t, err)
	saramaTLSCfg.Net.TLS.Config = tlscfg

	saramaKerberosCfg := &sarama.Config{}
	saramaKerberosCfg.Net.SASL.Mechanism = sarama.SASLTypeGSSAPI
	saramaKerberosCfg.Net.SASL.Enable = true
	saramaKerberosCfg.Net.SASL.GSSAPI.ServiceName = "foobar"
	saramaKerberosCfg.Net.SASL.GSSAPI.AuthType = sarama.KRB5_USER_AUTH

	saramaKerberosKeyTabCfg := &sarama.Config{}
	saramaKerberosKeyTabCfg.Net.SASL.Mechanism = sarama.SASLTypeGSSAPI
	saramaKerberosKeyTabCfg.Net.SASL.Enable = true
	saramaKerberosKeyTabCfg.Net.SASL.GSSAPI.KeyTabPath = "/path"
	saramaKerberosKeyTabCfg.Net.SASL.GSSAPI.AuthType = sarama.KRB5_KEYTAB_AUTH

	tests := []struct {
		auth         Authentication
		saramaConfig *sarama.Config
		err          string
	}{
		{
			auth:         Authentication{PlainText: &PlainTextConfig{Username: "jdoe", Password: "pass"}},
			saramaConfig: saramaPlaintext,
		},
		{
			auth:         Authentication{TLS: &configtls.TLSClientSetting{}},
			saramaConfig: saramaTLSCfg,
		},
		{
			auth: Authentication{TLS: &configtls.TLSClientSetting{
				TLSSetting: configtls.TLSSetting{CAFile: "/doesnotexists"},
			}},
			saramaConfig: saramaTLSCfg,
			err:          "failed to load TLS config",
		},
		{
			auth:         Authentication{Kerberos: &KerberosConfig{ServiceName: "foobar"}},
			saramaConfig: saramaKerberosCfg,
		},
		{
			auth:         Authentication{Kerberos: &KerberosConfig{UseKeyTab: true, KeyTabPath: "/path"}},
			saramaConfig: saramaKerberosKeyTabCfg,
		},
		{
			auth:         Authentication{SASL: &SASLConfig{Username: "jdoe", Password: "pass", Mechanism: "SCRAM-SHA-256"}},
			saramaConfig: saramaSASLSCRAM256Config,
		},
		{
			auth:         Authentication{SASL: &SASLConfig{Username: "jdoe", Password: "pass", Mechanism: "SCRAM-SHA-512"}},
			saramaConfig: saramaSASLSCRAM512Config,
		},

		{
			auth:         Authentication{SASL: &SASLConfig{Username: "jdoe", Password: "pass", Mechanism: "PLAIN"}},
			saramaConfig: saramaSASLPLAINConfig,
		},
		{
			auth:         Authentication{SASL: &SASLConfig{Username: "jdoe", Password: "pass", Mechanism: "SCRAM-SHA-222"}},
			saramaConfig: saramaSASLSCRAM512Config,
			err:          "invalid SASL Mechanism",
		},
		{
			auth:         Authentication{SASL: &SASLConfig{Username: "", Password: "pass", Mechanism: "SCRAM-SHA-512"}},
			saramaConfig: saramaSASLSCRAM512Config,
			err:          "username have to be provided",
		},
		{
			auth:         Authentication{SASL: &SASLConfig{Username: "jdoe", Password: "", Mechanism: "SCRAM-SHA-512"}},
			saramaConfig: saramaSASLSCRAM512Config,
			err:          "password have to be provided",
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			config := &sarama.Config{}
			err := ConfigureAuthentication(test.auth, config)
			if test.err != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.err)
			} else {
				// equalizes SCRAMClientGeneratorFunc to do assertion with the same reference.
				config.Net.SASL.SCRAMClientGeneratorFunc = test.saramaConfig.Net.SASL.SCRAMClientGeneratorFunc
				assert.Equal(t, test.saramaConfig, config)
			}
		})
	}
}
