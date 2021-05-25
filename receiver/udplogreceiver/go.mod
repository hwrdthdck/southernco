module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver

go 1.15

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/stanza v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-log-collection v0.18.1-0.20210517203758-2ccc3070b8b6
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.26.1-0.20210517223805-ce829467ebea
	go.uber.org/zap v1.16.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/stanza => ../../internal/stanza

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage
