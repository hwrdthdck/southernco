module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter

go 1.20

require (
	github.com/census-instrumentation/opencensus-proto v0.4.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.87.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.87.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.87.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver v0.87.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.87.1-0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/config/configgrpc v0.87.1-0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/config/configopaque v0.88.0
	go.opentelemetry.io/collector/config/configtls v0.87.1-0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/confmap v0.87.1-0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/consumer v0.87.1-0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/exporter v0.87.1-0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0016.0.20231023033326-37116a25be8d
	go.opentelemetry.io/collector/receiver v0.87.1-0.20231023033326-37116a25be8d
	google.golang.org/grpc v1.59.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.1 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.87.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/config/configauth v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/config/configcompression v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/config/confignet v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/config/internal v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/extension v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/extension/auth v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/featuregate v1.0.0-rcv0016.0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/collector/semconv v0.87.1-0.20231023033326-37116a25be8d // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.45.0 // indirect
	go.opentelemetry.io/otel v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent => ../../internal/sharedcomponent

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../../pkg/translator/opencensus

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver => ../../receiver/opencensusreceiver

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

// ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules
replace cloud.google.com/go => cloud.google.com/go v0.110.7
