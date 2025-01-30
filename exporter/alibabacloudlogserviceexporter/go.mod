module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter

go 1.22.0

require (
	github.com/aliyun/aliyun-log-go-sdk v0.1.83
	github.com/gogo/protobuf v1.3.2
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.118.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.118.1-0.20250130224751-50b76b95bf35
	go.opentelemetry.io/collector/component/componenttest v0.118.1-0.20250130224751-50b76b95bf35
	go.opentelemetry.io/collector/config/configopaque v1.24.1-0.20250130224751-50b76b95bf35
	go.opentelemetry.io/collector/confmap v1.24.1-0.20250130224751-50b76b95bf35
	go.opentelemetry.io/collector/exporter v0.118.1-0.20250130220845-0276d781be89
	go.opentelemetry.io/collector/exporter/exportertest v0.118.1-0.20250130220845-0276d781be89
	go.opentelemetry.io/collector/pdata v1.24.1-0.20250130220845-0276d781be89
	go.opentelemetry.io/collector/pdata/testdata v0.118.1-0.20250130220845-0276d781be89
	go.opentelemetry.io/collector/semconv v0.118.1-0.20250130220845-0276d781be89
	go.uber.org/zap v1.27.0
)

require (
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.24.1-0.20250130224751-50b76b95bf35 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/consumer v1.24.1-0.20250130224751-50b76b95bf35 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/extension v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/featuregate v1.24.1-0.20250130224751-50b76b95bf35 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/pipeline v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/receiver v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.118.1-0.20250130220845-0276d781be89 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
