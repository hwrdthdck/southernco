module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite

go 1.21

require (
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.98.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.98.0
	github.com/prometheus/common v0.52.3
	github.com/prometheus/prometheus v0.51.2-0.20240405174432-b4a973753c6e
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/pdata v1.5.1-0.20240416174005-d0f15e2463f8
	go.opentelemetry.io/collector/semconv v0.98.1-0.20240416174005-d0f15e2463f8
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	go.opentelemetry.io/collector/featuregate v1.5.1-0.20240416174005-d0f15e2463f8 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus => ../prometheus

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../../internal/common

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../golden
