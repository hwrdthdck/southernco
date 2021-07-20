module github.com/open-telemetry/opentelemetry-collector-contrib/tracegen

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/otel v1.0.0-RC1
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.0-RC1
	go.opentelemetry.io/otel/sdk v1.0.0-RC1
	go.opentelemetry.io/otel/trace v1.0.0-RC1
	go.uber.org/zap v1.18.1
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6
	google.golang.org/grpc v1.39.0
)

replace go.opentelemetry.io/collector => go.opentelemetry.io/collector v0.29.1-0.20210716020257-4d8e3082465d

replace go.opentelemetry.io/collector/model => go.opentelemetry.io/collector/model v0.0.0-20210716020257-4d8e3082465d
