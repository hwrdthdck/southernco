module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter

go 1.21.0

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.100.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.100.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.100.0
	github.com/openzipkin/zipkin-go v0.4.3
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/config/confighttp v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/config/configretry v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/config/configtls v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/confmap v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/consumer v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/exporter v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/collector/pdata v1.7.1-0.20240517133416-ede9e304314d
	go.opentelemetry.io/collector/receiver v0.100.1-0.20240509190532-c555005fcc80
	go.opentelemetry.io/otel/metric v1.26.0
	go.opentelemetry.io/otel/trace v1.26.0
	go.uber.org/goleak v1.3.0
)

require (
	github.com/apache/thrift v0.20.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/jaegertracing/jaeger v1.57.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.100.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	go.opentelemetry.io/collector v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/collector/config/configauth v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.7.1-0.20240517133416-ede9e304314d // indirect
	go.opentelemetry.io/collector/config/configopaque v1.7.1-0.20240517133416-ede9e304314d // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/collector/config/internal v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/collector/extension v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/collector/extension/auth v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/collector/featuregate v1.7.1-0.20240517133416-ede9e304314d // indirect
	go.opentelemetry.io/collector/semconv v0.100.1-0.20240509190532-c555005fcc80 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.48.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.26.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin => ../../pkg/translator/zipkin

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver => ../../receiver/zipkinreceiver

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
