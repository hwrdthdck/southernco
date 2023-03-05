module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver

go 1.19

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.72.0
	github.com/stretchr/testify v1.8.2
	go.opentelemetry.io/collector/component v0.72.1-0.20230303235035-7318c14f1a2b
	go.opentelemetry.io/collector/confmap v0.72.1-0.20230303235035-7318c14f1a2b
	go.opentelemetry.io/collector/consumer v0.72.1-0.20230303235035-7318c14f1a2b
	go.opentelemetry.io/collector/receiver v0.0.0-20230302200458-4071a47d0ee3
)

require (
	github.com/antonmedv/expr v1.12.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/observiq/ctimefmt v1.0.0 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.72.1-0.20230303235035-7318c14f1a2b // indirect
	go.opentelemetry.io/collector/exporter v0.0.0-20230303211526-ec5d71fec2da // indirect
	go.opentelemetry.io/collector/featuregate v0.72.1-0.20230303235035-7318c14f1a2b // indirect
	go.opentelemetry.io/collector/pdata v1.0.0-rc6.0.20230303235035-7318c14f1a2b // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.13.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gonum.org/v1/gonum v0.12.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ../../pkg/stanza

retract v0.65.0
