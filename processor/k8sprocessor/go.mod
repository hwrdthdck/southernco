module github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sprocessor

go 1.16

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.33.1-0.20210827152330-09258f969908
	go.opentelemetry.io/collector/model v0.33.1-0.20210827152330-09258f969908
	go.uber.org/zap v1.19.0
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.3 // indirect
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ./../../internal/k8sconfig
