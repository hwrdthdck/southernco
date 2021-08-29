module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/aliyun/aliyun-log-go-sdk v0.1.21
	github.com/gogo/protobuf v1.3.2
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.33.1-0.20210827152330-09258f969908
	go.opentelemetry.io/collector/model v0.33.1-0.20210827152330-09258f969908
	go.uber.org/zap v1.19.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal
