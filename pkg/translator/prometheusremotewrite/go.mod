module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite

go 1.20

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.85.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.85.0
	github.com/prometheus/common v0.44.0
	github.com/prometheus/prometheus v0.47.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0014.0.20230922175119-921b6125f017
	go.opentelemetry.io/collector/semconv v0.85.1-0.20230922175119-921b6125f017
	go.uber.org/multierr v1.11.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	go.opentelemetry.io/collector/featuregate v1.0.0-rcv0014.0.20230922175119-921b6125f017 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230717213848-3f92550aa753 // indirect
	google.golang.org/grpc v1.58.1 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
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
