// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sprocessor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.opentelemetry.io/collector/translator/conventions"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sprocessor/kube"
)

const (
	// The value of "type" key in configuration.
	typeStr = "k8s_tagger"
)

var kubeClientProvider = kube.ClientProvider(nil)
var processorCapabilities = component.ProcessorCapabilities{MutatesConsumedData: true}

// NewFactory returns a new factory for the k8s processor.
func NewFactory() component.ProcessorFactory {
	return processorhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		processorhelper.WithTraces(createTracesProcessor),
		processorhelper.WithMetrics(createMetricsProcessor),
		processorhelper.WithLogs(createLogsProcessor),
	)
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewID(typeStr)),
		APIConfig:         k8sconfig.APIConfig{AuthType: k8sconfig.AuthTypeServiceAccount},
	}
}

func createTracesProcessor(
	ctx context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	next consumer.Traces,
) (component.TracesProcessor, error) {
	return createTracesProcessorWithOptions(ctx, params, cfg, next)
}

func createLogsProcessor(
	ctx context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextLogsConsumer consumer.Logs,
) (component.LogsProcessor, error) {
	return createLogsProcessorWithOptions(ctx, params, cfg, nextLogsConsumer)
}

func createMetricsProcessor(
	ctx context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextMetricsConsumer consumer.Metrics,
) (component.MetricsProcessor, error) {
	return createMetricsProcessorWithOptions(ctx, params, cfg, nextMetricsConsumer)
}

func createTracesProcessorWithOptions(
	_ context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	next consumer.Traces,
	options ...Option,
) (component.TracesProcessor, error) {
	kp, err := createKubernetesProcessor(params, cfg, options...)
	if err != nil {
		return nil, err
	}

	return processorhelper.NewTracesProcessor(
		cfg,
		next,
		kp,
		processorhelper.WithCapabilities(processorCapabilities),
		processorhelper.WithStart(kp.Start),
		processorhelper.WithShutdown(kp.Shutdown))
}

func createMetricsProcessorWithOptions(
	_ context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextMetricsConsumer consumer.Metrics,
	options ...Option,
) (component.MetricsProcessor, error) {
	kp, err := createKubernetesProcessor(params, cfg, options...)
	if err != nil {
		return nil, err
	}

	return processorhelper.NewMetricsProcessor(
		cfg,
		nextMetricsConsumer,
		kp,
		processorhelper.WithCapabilities(processorCapabilities),
		processorhelper.WithStart(kp.Start),
		processorhelper.WithShutdown(kp.Shutdown))
}

func createLogsProcessorWithOptions(
	_ context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextLogsConsumer consumer.Logs,
	options ...Option,
) (component.LogsProcessor, error) {
	kp, err := createKubernetesProcessor(params, cfg, options...)
	if err != nil {
		return nil, err
	}

	return processorhelper.NewLogsProcessor(
		cfg,
		nextLogsConsumer,
		kp,
		processorhelper.WithCapabilities(processorCapabilities),
		processorhelper.WithStart(kp.Start),
		processorhelper.WithShutdown(kp.Shutdown))
}

func createKubernetesProcessor(
	params component.ProcessorCreateParams,
	cfg config.Processor,
	options ...Option,
) (*kubernetesprocessor, error) {
	kp := &kubernetesprocessor{logger: params.Logger}

	warnDeprecatedMetadataConfig(kp.logger, cfg)

	allOptions := append(createProcessorOpts(cfg), options...)

	for _, opt := range allOptions {
		if err := opt(kp); err != nil {
			return nil, err
		}
	}

	// This might have been set by an option already
	if kp.kc == nil {
		err := kp.initKubeClient(kp.logger, kubeClientProvider)
		if err != nil {
			return nil, err
		}
	}

	return kp, nil
}

func createProcessorOpts(cfg config.Processor) []Option {
	oCfg := cfg.(*Config)
	opts := []Option{}
	if oCfg.Passthrough {
		opts = append(opts, WithPassthrough())
	}

	// extraction rules
	opts = append(opts, WithExtractMetadata(oCfg.Extract.Metadata...))
	opts = append(opts, WithExtractLabels(oCfg.Extract.Labels...))
	opts = append(opts, WithExtractAnnotations(oCfg.Extract.Annotations...))

	// filters
	opts = append(opts, WithFilterNode(oCfg.Filter.Node, oCfg.Filter.NodeFromEnvVar))
	opts = append(opts, WithFilterNamespace(oCfg.Filter.Namespace))
	opts = append(opts, WithFilterLabels(oCfg.Filter.Labels...))
	opts = append(opts, WithFilterFields(oCfg.Filter.Fields...))
	opts = append(opts, WithAPIConfig(oCfg.APIConfig))

	opts = append(opts, WithExtractPodAssociations(oCfg.Association...))

	return opts
}

func warnDeprecatedMetadataConfig(logger *zap.Logger, cfg config.Processor) {
	oCfg := cfg.(*Config)
	var oldName, newName string

	for _, field := range oCfg.Extract.Metadata {
		switch field {
		case metdataNamespace:
			oldName = metdataNamespace
			newName = conventions.AttributeK8sNamespace
		case metadataPodName:
			oldName = metadataPodName
			newName = conventions.AttributeK8sPod
		case metadataPodUID:
			oldName = metadataPodUID
			newName = conventions.AttributeK8sPodUID
		case metadataStartTime:
			oldName = metadataStartTime
			newName = metadataPodStartTime
		case metadataDeployment:
			oldName = metadataDeployment
			newName = conventions.AttributeK8sDeployment
		case metadataCluster:
			oldName = metadataCluster
			newName = conventions.AttributeK8sCluster
		case metadataNode:
			oldName = metadataNode
			newName = conventions.AttributeK8sNodeName
		}
		if oldName != "" {
			logger.Warn(fmt.Sprintf("%s has been deprecated in favor of %s for k8s-tagger processor", oldName, newName))
		}
	}
}
