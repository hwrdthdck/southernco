module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter

go 1.14

require (
	github.com/aws/aws-sdk-go v1.34.9
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/golang/protobuf v1.4.2
	github.com/shirou/gopsutil v2.20.9+incompatible // indirect
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.11.1-0.20201006165100-07236c11fb27
	go.uber.org/zap v1.16.0
)
