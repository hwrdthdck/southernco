module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter

go 1.14

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/awsxray => ../../internal/common/awsxray

require (
	github.com/aws/aws-sdk-go v1.34.10
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/awsxray v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-proto v0.4.0
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.8.1-0.20200820203435-961c48b75778
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
)

// Yet another hack that we need until kubernetes client moves to the new github.com/googleapis/gnostic
replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
