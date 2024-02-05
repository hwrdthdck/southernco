module github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension

go 1.20

require (
	bitbucket.org/atlassian/go-asap/v2 v2.7.0
	github.com/SermoDigital/jose v0.9.2-0.20180104203859-803625baeddc
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.93.1-0.20240205121451-f5a7315cf88e
	go.opentelemetry.io/collector/config/configopaque v0.93.1-0.20240205121451-f5a7315cf88e
	go.opentelemetry.io/collector/confmap v0.93.1-0.20240205121451-f5a7315cf88e
	go.opentelemetry.io/collector/extension v0.93.1-0.20240205121451-f5a7315cf88e
	go.opentelemetry.io/collector/extension/auth v0.93.1-0.20240205121451-f5a7315cf88e
	go.opentelemetry.io/otel/metric v1.22.0
	go.opentelemetry.io/otel/trace v1.22.0
	go.uber.org/multierr v1.11.0
	google.golang.org/grpc v1.61.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/vincent-petithory/dataurl v1.0.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.93.1-0.20240205121451-f5a7315cf88e // indirect
	go.opentelemetry.io/collector/pdata v1.0.2-0.20240205121451-f5a7315cf88e // indirect
	go.opentelemetry.io/otel v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.45.0 // indirect
	go.opentelemetry.io/otel/sdk v1.22.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.22.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
