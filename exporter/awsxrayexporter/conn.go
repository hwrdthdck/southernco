// Copyright 2019, OpenTelemetry Authors
// Portions of this file Copyright 2018-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package awsxrayexporter

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/version"

	"go.uber.org/zap"
	"golang.org/x/net/http2"
)

type connAttr interface {
	newAWSSession(logger *zap.Logger, roleArn string, region string) (*session.Session, error)
	getEC2Region(s *session.Session) (string, error)
}

// Conn implements connAttr interface.
type Conn struct{}

func (c *Conn) getEC2Region(s *session.Session) (string, error) {
	return ec2metadata.New(s).Region()
}

// AWS STS endpoint constants
const (
	STSEndpointPrefix         = "https://sts."
	STSEndpointSuffix         = ".amazonaws.com"
	STSAwsCnPartitionIDSuffix = ".amazonaws.com.cn" // AWS China partition.
)

var collectorUserAgentHandler = request.NamedHandler{
	Name: "otel.collector.UserAgentHandler",
	Fn:   request.MakeAddToUserAgentHandler("opentelemetry-collector-contrib", version.Version, version.GitHash),
}

// newHTTPClient returns new HTTP client instance with provided configuration.
func newHTTPClient(logger *zap.Logger, maxIdle int, requestTimeout int, noVerify bool,
	proxyAddress string) (*http.Client, error) {
	logger.Debug("Using proxy address: ",
		zap.String("proxyAddr", proxyAddress),
	)
	tls := &tls.Config{
		InsecureSkipVerify: noVerify,
	}

	finalProxyAddress := getProxyAddress(proxyAddress)
	proxyURL, err := getProxyURL(finalProxyAddress)
	if err != nil {
		logger.Error("unable to obtain proxy URL", zap.Error(err))
		return nil, err
	}
	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdle,
		TLSClientConfig:     tls,
		Proxy:               http.ProxyURL(proxyURL),
	}

	// is not enabled by default as we configure TLSClientConfig for supporting SSL to data plane.
	// http2.ConfigureTransport will setup transport layer to use HTTP2
	http2.ConfigureTransport(transport)
	http := &http.Client{
		Transport: transport,
		Timeout:   time.Second * time.Duration(requestTimeout),
	}
	return http, err
}

func getProxyAddress(proxyAddress string) string {
	var finalProxyAddress string
	if proxyAddress != "" {
		finalProxyAddress = proxyAddress
	} else if proxyAddress == "" && os.Getenv("HTTPS_PROXY") != "" {
		finalProxyAddress = os.Getenv("HTTPS_PROXY")
	} else {
		finalProxyAddress = ""
	}
	return finalProxyAddress
}

func getProxyURL(finalProxyAddress string) (*url.URL, error) {
	var proxyURL *url.URL
	var err error
	if finalProxyAddress != "" {
		proxyURL, err = url.Parse(finalProxyAddress)
	} else {
		proxyURL = nil
		err = nil
	}
	return proxyURL, err
}

// GetAWSConfigSession returns AWS config and session instances.
func GetAWSConfigSession(logger *zap.Logger, cn connAttr, cfg *Config) (*aws.Config, *session.Session, error) {
	var s *session.Session
	var err error
	var awsRegion string
	http, err := newHTTPClient(logger, cfg.NumberOfWorkers, cfg.RequestTimeoutSeconds, cfg.NoVerifySSL, cfg.ProxyAddress)
	if err != nil {
		logger.Error("unable to obtain proxy URL", zap.Error(err))
		return nil, nil, err
	}
	regionEnv := os.Getenv("AWS_REGION")
	if cfg.Region == "" && regionEnv != "" {
		awsRegion = regionEnv
		logger.Debug("Fetch region from environment variables", zap.String("region", awsRegion))
	} else if cfg.Region != "" {
		awsRegion = cfg.Region
		logger.Debug("Fetch region from commandline/config file", zap.String("region", awsRegion))
	} else if !cfg.NoVerifySSL {
		var es *session.Session
		es, err = getDefaultSession(logger)
		if err != nil {
			logger.Error("Unable to retrieve default session", zap.Error(err))
		} else {
			awsRegion, err = cn.getEC2Region(es)
			if err != nil {
				logger.Error("Unable to retrieve the region from the EC2 instance", zap.Error(err))
			} else {
				logger.Debug("Fetch region from ec2 metadata", zap.String("region", awsRegion))
			}
		}
	}
	if awsRegion == "" {
		msg := "Cannot fetch region variable from config file, environment variables and ec2 metadata."
		logger.Error(msg)
		return nil, nil, awserr.New("NoAwsRegion", msg, nil)
	}
	s, err = cn.newAWSSession(logger, cfg.RoleARN, awsRegion)
	if err != nil {
		return nil, nil, err
	}

	s.Handlers.Build.PushBackNamed(collectorUserAgentHandler)

	config := &aws.Config{
		Region:                 aws.String(awsRegion),
		DisableParamValidation: aws.Bool(true),
		MaxRetries:             aws.Int(cfg.MaxRetries),
		Endpoint:               aws.String(cfg.Endpoint),
		HTTPClient:             http,
	}
	return config, s, nil
}

// ProxyServerTransport configures HTTP transport for TCP Proxy Server.
func ProxyServerTransport(logger *zap.Logger, config *Config) (*http.Transport, error) {
	tls := &tls.Config{
		InsecureSkipVerify: config.NoVerifySSL,
	}

	proxyAddr := getProxyAddress(config.ProxyAddress)
	proxyURL, err := getProxyURL(proxyAddr)
	if err != nil {
		logger.Error("unable to obtain proxy URL", zap.Error(err))
		return nil, err
	}

	// Connection timeout in seconds
	idleConnTimeout := time.Duration(config.RequestTimeoutSeconds) * time.Second

	transport := &http.Transport{
		MaxIdleConns:        config.NumberOfWorkers,
		MaxIdleConnsPerHost: config.NumberOfWorkers,
		IdleConnTimeout:     idleConnTimeout,
		Proxy:               http.ProxyURL(proxyURL),
		TLSClientConfig:     tls,

		// If not disabled the transport will add a gzip encoding header
		// to requests with no `accept-encoding` header value. The header
		// is added after we sign the request which invalidates the
		// signature.
		DisableCompression: true,
	}

	return transport, nil
}

func (c *Conn) newAWSSession(logger *zap.Logger, roleArn string, region string) (*session.Session, error) {
	var s *session.Session
	var err error
	if roleArn == "" {
		s, err = getDefaultSession(logger)
		if err != nil {
			return s, err
		}
	} else {
		stsCreds, _ := getSTSCreds(logger, region, roleArn)

		s, err = session.NewSession(&aws.Config{
			Credentials: stsCreds,
		})

		if err != nil {
			logger.Error("Error in creating session object : ", zap.Error(err))
			return s, err
		}
	}
	return s, nil
}

// getSTSCreds gets STS credentials from regional endpoint. ErrCodeRegionDisabledException is received if the
// STS regional endpoint is disabled. In this case STS credentials are fetched from STS primary regional endpoint
// in the respective AWS partition.
func getSTSCreds(logger *zap.Logger, region string, roleArn string) (*credentials.Credentials, error) {
	t, err := getDefaultSession(logger)
	if err != nil {
		return nil, err
	}

	stsCred := getSTSCredsFromRegionEndpoint(logger, t, region, roleArn)
	// Make explicit call to fetch credentials.
	_, err = stsCred.Get()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			err = nil
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				logger.Error("Region ", zap.String("region", region), zap.String("error", aerr.Error()))
				stsCred = getSTSCredsFromPrimaryRegionEndpoint(logger, t, roleArn, region)
			}
		}
	}
	return stsCred, err
}

// getSTSCredsFromRegionEndpoint fetches STS credentials for provided roleARN from regional endpoint.
// AWS STS recommends that you provide both the Region and endpoint when you make calls to a Regional endpoint.
// Reference: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html#id_credentials_temp_enable-regions_writing_code
func getSTSCredsFromRegionEndpoint(logger *zap.Logger, sess *session.Session, region string,
	roleArn string) *credentials.Credentials {
	regionalEndpoint := getSTSRegionalEndpoint(region)
	// if regionalEndpoint is "", the STS endpoint is Global endpoint for classic regions except ap-east-1 - (HKG)
	// for other opt-in regions, region value will create STS regional endpoint.
	// This will be only in the case, if provided region is not present in aws_regions.go
	c := &aws.Config{Region: aws.String(region), Endpoint: &regionalEndpoint}
	st := sts.New(sess, c)
	logger.Info("STS Endpoint ", zap.String("endpoint", st.Endpoint))
	return stscreds.NewCredentialsWithClient(st, roleArn)
}

// getSTSCredsFromPrimaryRegionEndpoint fetches STS credentials for provided roleARN from primary region endpoint in
// the respective partition.
func getSTSCredsFromPrimaryRegionEndpoint(logger *zap.Logger, t *session.Session, roleArn string,
	region string) *credentials.Credentials {
	logger.Info("Credentials for provided RoleARN being fetched from STS primary region endpoint.")
	partitionID := getPartition(region)
	if partitionID == endpoints.AwsPartitionID {
		return getSTSCredsFromRegionEndpoint(logger, t, endpoints.UsEast1RegionID, roleArn)
	} else if partitionID == endpoints.AwsCnPartitionID {
		return getSTSCredsFromRegionEndpoint(logger, t, endpoints.CnNorth1RegionID, roleArn)
	} else if partitionID == endpoints.AwsUsGovPartitionID {
		return getSTSCredsFromRegionEndpoint(logger, t, endpoints.UsGovWest1RegionID, roleArn)
	}

	return nil
}

func getSTSRegionalEndpoint(r string) string {
	p := getPartition(r)

	var e string
	if p == endpoints.AwsPartitionID || p == endpoints.AwsUsGovPartitionID {
		e = STSEndpointPrefix + r + STSEndpointSuffix
	} else if p == endpoints.AwsCnPartitionID {
		e = STSEndpointPrefix + r + STSAwsCnPartitionIDSuffix
	}
	return e
}

func getDefaultSession(logger *zap.Logger) (*session.Session, error) {
	result, serr := session.NewSession()
	if serr != nil {
		logger.Error("Error in creating session object ", zap.Error(serr))
		return result, serr
	}
	return result, nil
}

// getPartition return AWS Partition for the provided region.
func getPartition(region string) string {
	p, _ := endpoints.PartitionForRegion(endpoints.DefaultPartitions(), region)
	return p.ID()
}
