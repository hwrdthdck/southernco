module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger

go 1.17

require (
	github.com/jaegertracing/jaeger v1.33.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.49.0
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/collector/pdata v0.49.0
	go.opentelemetry.io/collector/semconv v0.0.0-20220421154122-427f7dde5a8f
)

require (
	github.com/apache/thrift v0.16.0 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../../internal/coreinternal
