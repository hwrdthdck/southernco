module github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver

go 1.16

require (
	github.com/armon/go-metrics v0.3.3 // indirect
	github.com/gogo/googleapis v1.3.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.27.1-0.20210527142130-1f972bbd7997
	go.uber.org/zap v1.17.0
	gopkg.in/ini.v1 v1.57.0 // indirect
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer => ../

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ../../../internal/k8sconfig
