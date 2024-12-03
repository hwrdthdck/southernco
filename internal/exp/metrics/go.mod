module github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.114.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.114.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.114.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/pdata v1.20.1-0.20241202231142-b9ff1bc54c99
	go.opentelemetry.io/collector/semconv v0.114.1-0.20241202231142-b9ff1bc54c99
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../../pkg/golden

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../../pkg/pdatatest
