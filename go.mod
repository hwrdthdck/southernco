module github.com/open-telemetry/opentelemetry-collector-contrib

// NOTE:
// This go.mod is NOT used to build any official binary.
// To see the builder manifests used for official binaries,
// check https://github.com/open-telemetry/opentelemetry-collector-releases
//
// For the OpenTelemetry Collector Contrib distribution specifically, see
// https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib

go 1.21.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlemanagedprometheusexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/instanaexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/bigipreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otlpjsonfilereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcplogreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.103.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver v0.103.0
	go.opentelemetry.io/collector/connector v0.103.0
	go.opentelemetry.io/collector/connector/forwardconnector v0.103.0
	go.opentelemetry.io/collector/exporter v0.103.0
	go.opentelemetry.io/collector/exporter/debugexporter v0.103.0
	go.opentelemetry.io/collector/exporter/loggingexporter v0.103.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.103.0
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.103.0
	go.opentelemetry.io/collector/extension v0.103.0
	go.opentelemetry.io/collector/extension/ballastextension v0.103.0
	go.opentelemetry.io/collector/extension/zpagesextension v0.103.0
	go.opentelemetry.io/collector/otelcol v0.103.0
	go.opentelemetry.io/collector/processor v0.103.0
	go.opentelemetry.io/collector/processor/batchprocessor v0.103.0
	go.opentelemetry.io/collector/processor/memorylimiterprocessor v0.103.0
	go.opentelemetry.io/collector/receiver v0.103.0
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.103.0
)

require (
	bitbucket.org/atlassian/go-asap/v2 v2.8.0 // indirect
	cloud.google.com/go v0.115.0 // indirect
	cloud.google.com/go/auth v0.5.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	cloud.google.com/go/iam v1.1.8 // indirect
	cloud.google.com/go/logging v1.10.0 // indirect
	cloud.google.com/go/longrunning v0.5.7 // indirect
	cloud.google.com/go/monitoring v1.19.0 // indirect
	cloud.google.com/go/pubsub v1.38.0 // indirect
	cloud.google.com/go/spanner v1.63.0 // indirect
	cloud.google.com/go/trace v1.10.7 // indirect
	code.cloudfoundry.org/clock v1.0.0 // indirect
	code.cloudfoundry.org/go-diodes v0.0.0-20211115184647-b584dd5df32c // indirect
	code.cloudfoundry.org/go-loggregator v7.4.0+incompatible // indirect
	code.cloudfoundry.org/rfc5424 v0.0.0-20201103192249-000122071b78 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.2 // indirect
	github.com/AthenZ/athenz v1.10.39 // indirect
	github.com/Azure/azure-amqp-common-go/v4 v4.2.0 // indirect
	github.com/Azure/azure-event-hubs-go/v3 v3.6.2 // indirect
	github.com/Azure/azure-kusto-go v0.15.3 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.12.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.9.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5 v5.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor v0.11.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4 v4.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.3.2 // indirect
	github.com/Azure/azure-storage-queue-go v0.0.0-20230531184854-c06a8eff66fe // indirect
	github.com/Azure/go-amqp v1.0.5 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.29 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.12 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/ClickHouse/ch-go v0.61.5 // indirect
	github.com/ClickHouse/clickhouse-go/v2 v2.23.0 // indirect
	github.com/Code-Hex/go-generics-cache v1.3.1 // indirect
	github.com/DataDog/agent-payload/v5 v5.0.122 // indirect
	github.com/DataDog/datadog-agent/comp/core/config v0.54.0 // indirect
	github.com/DataDog/datadog-agent/comp/core/hostname/hostnameinterface v0.54.0 // indirect
	github.com/DataDog/datadog-agent/comp/core/log v0.54.0 // indirect
	github.com/DataDog/datadog-agent/comp/core/secrets v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/core/telemetry v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/def v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/logs/agent/config v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/logsagentpipeline v0.54.0 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/logsagentpipeline/logsagentpipelineimpl v0.54.0 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/exporter/logsagentexporter v0.54.0 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/metricsclient v0.0.0-20240525065430-d0b647bcb646 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/statsprocessor v0.0.0-20240525065430-d0b647bcb646 // indirect
	github.com/DataDog/datadog-agent/pkg/collector/check/defaults v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/env v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/model v0.54.0 // indirect
	github.com/DataDog/datadog-agent/pkg/config/setup v0.54.0 // indirect
	github.com/DataDog/datadog-agent/pkg/config/utils v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/auditor v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/client v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/diagnostic v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/message v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/metrics v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/pipeline v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/processor v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/sds v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/sender v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/sources v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/status/statusinterface v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/status/utils v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/proto v0.54.0 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/status/health v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/telemetry v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/trace v0.54.0 // indirect
	github.com/DataDog/datadog-agent/pkg/util/backoff v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/cgroups v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/executable v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/filesystem v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/fxutil v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/hostname/validate v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/http v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/optional v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/pointer v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/startstop v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/statstracker v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/system v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/system/socket v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/winutil v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/version v0.54.0-rc.5 // indirect
	github.com/DataDog/datadog-api-client-go/v2 v2.26.0 // indirect
	github.com/DataDog/datadog-go/v5 v5.5.0 // indirect
	github.com/DataDog/dd-sensitive-data-scanner/sds-go/go v0.0.0-20240419161837-f1b2f553edfe // indirect
	github.com/DataDog/go-sqllexer v0.0.9 // indirect
	github.com/DataDog/go-tuf v1.1.0-0.5.2 // indirect
	github.com/DataDog/gohai v0.0.0-20230524154621-4316413895ee // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata v0.16.1 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes v0.16.1 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/logs v0.16.1 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics v0.16.1 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/quantile v0.16.1 // indirect
	github.com/DataDog/sketches-go v1.4.5 // indirect
	github.com/DataDog/viper v1.13.3 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/GehirnInc/crypt v0.0.0-20200316065508-bb7000b8a962 // indirect
	github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp v1.5.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.24.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/collector v0.48.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/collector/googlemanagedprometheus v0.48.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.24.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.48.0 // indirect
	github.com/IBM/sarama v1.43.2 // indirect
	github.com/JohnCGriffin/overflow v0.0.0-20211019200055-46fa312c352c // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ReneKroon/ttlcache/v2 v2.11.0 // indirect
	github.com/SAP/go-hdb v1.9.8 // indirect
	github.com/SermoDigital/jose v0.9.2-0.20180104203859-803625baeddc // indirect
	github.com/Showmax/go-fqdn v1.0.0 // indirect
	github.com/aerospike/aerospike-client-go/v6 v6.13.0 // indirect
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/alecthomas/units v0.0.0-20231202071711-9a357b53e9c9 // indirect
	github.com/aliyun/aliyun-log-go-sdk v0.1.80 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/apache/arrow/go/v15 v15.0.0 // indirect
	github.com/apache/pulsar-client-go v0.8.1 // indirect
	github.com/apache/pulsar-client-go/oauth2 v0.0.0-20220120090717-25e59572242e // indirect
	github.com/apache/thrift v0.20.0 // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.53.11 // indirect
	github.com/aws/aws-sdk-go-v2 v1.27.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.16 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.16 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.16.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.27.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.53.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.29.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.24.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.10 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/briandowns/spinner v1.23.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/cilium/ebpf v0.11.0 // indirect
	github.com/cloudfoundry-incubator/uaago v0.0.0-20190307164349-8136b7bbe76e // indirect
	github.com/cncf/xds/go v0.0.0-20240318125728-8a4994d93e50 // indirect
	github.com/containerd/cgroups/v3 v3.0.3 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/containerd/ttrpc v1.2.2 // indirect
	github.com/coreos/go-oidc/v3 v3.10.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cskr/pubsub v1.0.2 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/devigned/tab v0.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/digitalocean/godo v1.109.0 // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/docker v25.0.5+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v1.6.0 // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.6.0 // indirect
	github.com/elastic/go-docappender/v2 v2.1.4 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.17.10 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.14.0 // indirect
	github.com/elastic/go-structform v0.0.10 // indirect
	github.com/elastic/go-sysinfo v1.7.1 // indirect
	github.com/elastic/go-windows v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/envoyproxy/go-control-plane v0.12.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/facebook/time v0.0.0-20240510113249-fa89cc575891 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/getsentry/sentry-go v0.28.1 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.7.1 // indirect
	github.com/go-jose/go-jose/v4 v4.0.1 // indirect
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.20.2 // indirect
	github.com/go-openapi/jsonreference v0.20.4 // indirect
	github.com/go-openapi/swag v0.22.9 // indirect
	github.com/go-resty/resty/v2 v2.12.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/go-zookeeper/zk v1.0.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gocql/gocql v1.6.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/cadvisor v0.49.0 // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
	github.com/gophercloud/gophercloud v1.8.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gosnmp/gosnmp v1.37.0 // indirect
	github.com/grafana/loki/pkg/push v0.0.0-20240514112848-a1b1eeb09583 // indirect
	github.com/grafana/regexp v0.0.0-20221122212121-6b5c0a4cb7fd // indirect
	github.com/grobie/gomemcache v0.0.0-20230213081705-239240bbc445 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/consul/api v1.29.1 // indirect
	github.com/hashicorp/cronexpr v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20240306004928-3e7191ccb702 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hectane/go-acl v0.0.0-20190604041725-da78bae5fc95 // indirect
	github.com/hetznercloud/hcloud-go/v2 v2.6.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/influxdata/influxdb-observability/common v0.5.12 // indirect
	github.com/influxdata/influxdb-observability/influx2otel v0.5.12 // indirect
	github.com/influxdata/influxdb-observability/otel2influx v0.5.12 // indirect
	github.com/influxdata/line-protocol/v2 v2.2.1 // indirect
	github.com/ionos-cloud/sdk-go/v6 v6.1.11 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jaegertracing/jaeger v1.58.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/kolo/xmlrpc v0.0.0-20220921171641-a4b6fa1dd06b // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-syslog/v4 v4.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20190525184631-5f46317e436b // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/lightstep/go-expohisto v1.0.0 // indirect
	github.com/linkedin/goavro/v2 v2.9.8 // indirect
	github.com/linode/linodego v1.33.0 // indirect
	github.com/logicmonitor/lm-data-sdk-go v1.3.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20220913051719-115f729f3c8c // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-ieproxy v0.0.11 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/microsoft/ApplicationInsights-Go v0.4.4 // indirect
	github.com/microsoft/go-mssqldb v1.7.2 // indirect
	github.com/miekg/dns v1.1.58 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mongodb-forks/digest v1.1.0 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/mrunalp/fileutils v0.5.1 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/nginxinc/nginx-prometheus-exporter v0.11.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/collectd v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/skywalking v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.103.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters v0.103.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opencontainers/runc v1.1.12 // indirect
	github.com/opencontainers/runtime-spec v1.1.0-rc.3 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/opensearch-project/opensearch-go/v2 v2.3.0 // indirect
	github.com/openshift/api v3.9.0+incompatible // indirect
	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/outcaste-io/ristretto v0.2.1 // indirect
	github.com/ovh/go-ovh v1.4.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paulmach/orb v0.11.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/sftp v1.13.6 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20220216144756-c35f1ee13d7c // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/prometheus-community/windows_exporter v0.25.1 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.54.0 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/prometheus/prometheus v0.51.2-0.20240405174432-b4a973753c6e // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/redis/go-redis/v9 v9.5.3 // indirect
	github.com/relvacode/iso8601 v1.4.0 // indirect
	github.com/rs/cors v1.11.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.25 // indirect
	github.com/scalyr/dataset-go v0.18.0 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20220502022130-f33da4d89646 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.7.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shirou/gopsutil/v4 v4.24.5 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3 // indirect
	github.com/signalfx/sapm-proto v0.14.0 // indirect
	github.com/sijms/go-ora/v2 v2.8.19 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/snowflakedb/gosnowflake v1.10.1 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	github.com/stormcat24/protodep v0.1.8 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.944 // indirect
	github.com/tg123/go-htpasswd v1.2.2 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/tinylru v1.1.0 // indirect
	github.com/tidwall/wal v1.1.7 // indirect
	github.com/tilinna/clock v1.1.0 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	github.com/vincent-petithory/dataurl v1.0.0 // indirect
	github.com/vishvananda/netlink v1.2.1-beta.2 // indirect
	github.com/vishvananda/netns v0.0.0-20210104183010-2eb08e3e575f // indirect
	github.com/vmware/go-vmware-nsxt v0.0.0-20230223012718-d31b8a1ca05e // indirect
	github.com/vmware/govmomi v0.36.3 // indirect
	github.com/vultr/govultr/v2 v2.17.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/yuin/gopher-lua v0.0.0-20220504180219-658193537a64 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.elastic.co/apm/module/apmzap/v2 v2.6.0 // indirect
	go.elastic.co/apm/v2 v2.6.0 // indirect
	go.elastic.co/fastjson v1.3.0 // indirect
	go.etcd.io/bbolt v1.3.10 // indirect
	go.mongodb.org/atlas v0.36.0 // indirect
	go.mongodb.org/mongo-driver v1.15.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.103.0 // indirect
	go.opentelemetry.io/collector/component v0.103.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.103.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.10.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.103.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.103.0 // indirect
	go.opentelemetry.io/collector/config/confignet v0.103.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.10.0 // indirect
	go.opentelemetry.io/collector/config/configretry v0.103.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.103.0 // indirect
	go.opentelemetry.io/collector/config/configtls v0.103.0 // indirect
	go.opentelemetry.io/collector/config/internal v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap/converter/expandconverter v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap/provider/envprovider v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap/provider/fileprovider v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap/provider/httpprovider v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap/provider/httpsprovider v0.103.0 // indirect
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v0.103.0 // indirect
	go.opentelemetry.io/collector/consumer v0.103.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.103.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.10.0 // indirect
	go.opentelemetry.io/collector/filter v0.103.0 // indirect
	go.opentelemetry.io/collector/pdata v1.10.0 // indirect
	go.opentelemetry.io/collector/semconv v0.103.0 // indirect
	go.opentelemetry.io/collector/service v0.103.0 // indirect
	go.opentelemetry.io/contrib/config v0.7.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.52.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.27.0 // indirect
	go.opentelemetry.io/contrib/zpages v0.52.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/bridge/opencensus v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.27.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	go.opentelemetry.io/proto/otlp v1.2.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/fx v1.18.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/term v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
	google.golang.org/api v0.185.0 // indirect
	google.golang.org/genproto v0.0.0-20240617180043-68d350f18fd4 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240610135401-a8a62080eff3 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240617180043-68d350f18fd4 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gopkg.in/zorkian/go-datadog-api.v2 v2.30.0 // indirect
	gotest.tools/v3 v3.5.0 // indirect
	howett.net/plist v1.0.0 // indirect
	k8s.io/api v0.29.3 // indirect
	k8s.io/apimachinery v0.29.3 // indirect
	k8s.io/client-go v0.29.3 // indirect
	k8s.io/klog/v2 v2.130.0 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	k8s.io/kubelet v0.29.3 // indirect
	k8s.io/utils v0.0.0-20240502163921-fe8a2dddb1d0 // indirect
	sigs.k8s.io/controller-runtime v0.17.3 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
	skywalking.apache.org/repo/goapi v0.0.0-20240104145220-ba7202308dd4 // indirect
)

// Replace references to modules that are in this repository with their relateive paths
// so that we always build with current (latest) version of the source code.

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => ./internal/aws/awsutil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight => ./internal/aws/containerinsight

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs => ./internal/aws/cwlogs

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil => ./internal/aws/ecsutil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s => ./internal/aws/k8s

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => ./internal/aws/metrics

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy => ./internal/aws/proxy

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray => ./internal/aws/xray

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ./internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ./internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker => ./internal/docker

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter => ./internal/filter

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ./internal/k8sconfig

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8stest => ./internal/k8stest

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka => ./internal/kafka

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet => ./internal/kubelet

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders => ./internal/metadataproviders

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent => ./internal/sharedcomponent

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ./internal/splunk

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter => ./exporter/alibabacloudlogserviceexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter => ./exporter/awscloudwatchlogsexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter => ./exporter/awsemfexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter => ./exporter/awskinesisexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter => ./exporter/awsxrayexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter => ./exporter/azuredataexplorerexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter => ./exporter/azuremonitorexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter => ./exporter/carbonexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter => ./exporter/clickhouseexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter => ./exporter/cassandraexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter => ./exporter/coralogixexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter => ./exporter/datadogexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter => ./exporter/datasetexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter => ./exporter/fileexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter => ./exporter/googlecloudexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlemanagedprometheusexporter => ./exporter/googlemanagedprometheusexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter => ./exporter/googlecloudpubsubexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter => ./exporter/influxdbexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/instanaexporter => ./exporter/instanaexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter => ./exporter/kafkaexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter => ./exporter/loadbalancingexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter => ./exporter/logicmonitorexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter => ./exporter/logzioexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter => ./exporter/lokiexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter => ./exporter/mezmoexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter => ./exporter/opencensusexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter => ./exporter/opensearchexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter => ./exporter/prometheusexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter => ./exporter/prometheusremotewriteexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter => ./exporter/pulsarexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter => ./exporter/sapmexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter => ./exporter/sentryexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter => ./exporter/signalfxexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter => ./exporter/skywalkingexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter => ./exporter/splunkhecexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter => ./exporter/sumologicexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter => ./exporter/tencentcloudlogserviceexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter => ./exporter/elasticsearchexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter => ./exporter/zipkinexporter

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension => ./extension/asapauthextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy => ./extension/awsproxy

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension => ./extension/basicauthextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension => ./extension/bearertokenauthextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension => ./extension/headerssetterextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension => ./extension/healthcheckextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension => ./extension/oauth2clientauthextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer => ./extension/observer

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver => ./extension/observer/ecstaskobserver

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver => ./extension/observer/hostobserver

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver => ./extension/observer/k8sobserver

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension => ./extension/oidcauthextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension => ./extension/pprofextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension => ./extension/sigv4authextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ./extension/storage

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage => ./extension/storage/dbstorage

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage => ./extension/storage/filestorage

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension => ./extension/sumologicextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => ./pkg/batchperresourceattr

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal => ./pkg/batchpersignal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters => ./pkg/winperfcounters

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => ./pkg/experimentalmetricmetadata

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ./pkg/ottl

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ./pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ./pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry => ./pkg/resourcetotelemetry

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ./pkg/stanza

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure => ./pkg/translator/azure

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger => ./pkg/translator/jaeger

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki => ./pkg/translator/loki

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ./pkg/translator/opencensus

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus => ./pkg/translator/prometheus

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite => ./pkg/translator/prometheusremotewrite

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx => ./pkg/translator/signalfx

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/skywalking => ./pkg/translator/skywalking

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin => ./pkg/translator/zipkin

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor => ./processor/attributesprocessor

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor => ./processor/cumulativetodeltaprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor => ./processor/deltatorateprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor => ./processor/filterprocessor

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor => ./processor/groupbyattrsprocessor

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor => ./processor/groupbytraceprocessor

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor => ./processor/k8sattributesprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor => ./processor/metricsgenerationprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor => ./processor/metricstransformprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor => ./processor/probabilisticsamplerprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor => ./processor/resourcedetectionprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor => ./processor/resourceprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor => ./processor/routingprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor => ./processor/spanprocessor/

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor => ./processor/tailsamplingprocessor

replace github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor => ./processor/transformprocessor

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver => ./receiver/activedirectorydsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver => ./receiver/aerospikereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver => ./receiver/apachereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver => ./receiver/apachesparkreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver => ./receiver/awscloudwatchreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver => ./receiver/awscontainerinsightreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver => ./receiver/awsecscontainermetricsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver => ./receiver/awsfirehosereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver => ./receiver/awsxrayreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver => ./receiver/azureblobreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver => ./receiver/azuremonitorreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/bigipreceiver => ./receiver/bigipreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver => ./receiver/carbonreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver => ./receiver/chronyreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver => ./receiver/cloudfoundryreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver => ./receiver/collectdreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver => ./receiver/couchdbreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver => ./receiver/datadogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver => ./receiver/flinkmetricsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver => ./receiver/dockerstatsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver => ./receiver/elasticsearchreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver => ./receiver/expvarreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver => ./receiver/filelogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver => ./receiver/googlecloudspannerreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver => ./receiver/googlecloudpubsubreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver => ./receiver/haproxyreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver => ./receiver/hostmetricsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver => ./receiver/httpcheckreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver => ./receiver/influxdbreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver => ./receiver/iisreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver => ./receiver/jaegerreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver => ./receiver/jmxreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver => ./receiver/journaldreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver => ./receiver/k8sclusterreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver => ./receiver/k8seventsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver => ./receiver/k8sobjectsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver => ./receiver/fluentforwardreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver => ./receiver/kafkametricsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver => ./receiver/kafkareceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver => ./receiver/kubeletstatsreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver => ./receiver/lokireceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver => ./receiver/memcachedreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver => ./receiver/mongodbreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver => ./receiver/mongodbatlasreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver => ./receiver/mysqlreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver => ./receiver/nginxreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver => ./receiver/nsxtreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver => ./receiver/opencensusreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver => ./receiver/oracledbreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otlpjsonfilereceiver => ./receiver/otlpjsonfilereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver => ./receiver/podmanreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver => ./receiver/postgresqlreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver => ./receiver/prometheusreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver => ./receiver/pulsarreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver => ./receiver/purefareceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver => ./receiver/purefbreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver => ./receiver/rabbitmqreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator => ./receiver/receivercreator

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver => ./receiver/redisreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver => ./receiver/riakreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver => ./receiver/saphanareceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver => ./receiver/sapmreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver => ./receiver/signalfxreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver => ./receiver/simpleprometheusreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver => ./receiver/skywalkingreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver => ./receiver/snmpreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver => ./receiver/solacereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver => ./receiver/splunkhecreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver => ./receiver/sqlqueryreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver => ./receiver/sqlserverreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver => ./receiver/sshcheckreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver => ./receiver/statsdreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver => ./receiver/syslogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcplogreceiver => ./receiver/tcplogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver => ./receiver/udplogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver => ./receiver/vcenterreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver => ./receiver/wavefrontreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver => ./receiver/windowseventlogreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver => ./receiver/webhookeventreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver => ./receiver/windowsperfcountersreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver => ./receiver/zipkinreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver => ./receiver/zookeeperreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver => ./extension/observer/dockerobserver

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling => ./extension/jaegerremotesampling

// see https://github.com/google/gnostic/issues/262
replace github.com/googleapis/gnostic v0.5.6 => github.com/googleapis/gnostic v0.5.5

// see https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/12322#issuecomment-1185029670
replace github.com/docker/go-connections v0.4.1-0.20210727194412-58542c764a11 => github.com/docker/go-connections v0.4.0

retract (
	v0.76.2
	v0.76.1
	v0.65.0
	v0.37.0 // Contains dependencies on v0.36.0 components, which should have been updated to v0.37.0.
)

// see https://github.com/distribution/distribution/issues/3590
exclude github.com/docker/distribution v2.8.0+incompatible

// see https://github.com/DataDog/agent-payload/issues/218
exclude github.com/DataDog/agent-payload/v5 v5.0.59

// see https://github.com/mattn/go-ieproxy/issues/45
replace github.com/mattn/go-ieproxy => github.com/mattn/go-ieproxy v0.0.1

// openshift removed all tags from their repo, use the pseudoversion from the release-3.9 branch HEAD
replace github.com/openshift/api v3.9.0+incompatible => github.com/openshift/api v0.0.0-20180801171038-322a19404e37

// It appears that the v0.2.0 tag was modified.  Replacing with v0.2.1
replace github.com/outcaste-io/ristretto v0.2.0 => github.com/outcaste-io/ristretto v0.2.1

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver => ./receiver/snowflakereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver => ./receiver/azureeventhubreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ./pkg/golden

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/collectd => ./internal/collectd

replace (
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector => ./connector/countconnector
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector => ./connector/datadogconnector
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector => ./connector/exceptionsconnector
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector => ./connector/grafanacloudconnector
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector => ./connector/routingconnector
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector => ./connector/servicegraphconnector
	github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector => ./connector/spanmetricsconnector

)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery => ./internal/sqlquery

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding => ./extension/encoding

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/otlpencodingextension => ./extension/encoding/otlpencodingextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension => ./extension/ackextension

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver => ./receiver/splunkenterprisereceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling => ./pkg/sampling

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics => ./internal/exp/metrics

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil => ./internal/pdatautil
