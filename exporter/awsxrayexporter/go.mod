module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter

go 1.22.0

require (
	github.com/aws/aws-sdk-go v1.55.5
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.115.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.115.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.115.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/component/componenttest v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/confmap v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/consumer/consumererror v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/exporter v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/exporter/exportertest v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/featuregate v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/pdata v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/semconv v0.115.1-0.20241216091623-8ac40a01a5ff
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.21.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/consumer v1.21.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/extension v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/pipeline v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/receiver v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/grpc v1.68.1 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray => ./../../internal/aws/xray

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => ./../../internal/aws/awsutil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
