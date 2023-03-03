module github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector

go 1.18

require github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor v0.72.0

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.72.1-0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/component v0.72.1-0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/confmap v0.72.1-0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/consumer v0.72.1-0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/featuregate v0.72.1-0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/pdata v1.0.0-rc6.0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/receiver v0.0.0-20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/collector/semconv v0.72.1-0.20230303004618-4a6ebc82b8e3 // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.13.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

retract v0.65.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor => ../../processor/servicegraphprocessor/
