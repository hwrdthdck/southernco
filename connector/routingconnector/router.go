// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingconnector"

import (
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var errPipelineNotFound = errors.New("pipeline not found")

// consumerProvider is a function with a type parameter T (expected to be one
// of consumer.Traces, consumer.Metrics, or Consumer.Logs). returns a
// consumer for the given component ID(s).
type consumerProvider[T any] func(...component.ID) (T, error)

// router registers consumers and default consumers for a pipeline. the type
// parameter C is expected to be one of: consumer.Traces, consumer.Metrics, or
// consumer.Logs. K is expected to be one of: ottlspan.TransformContext,
// ottlmetrics.TransformContext, or ottllog.TransformContext
type router[C any, K any] struct {
	logger *zap.Logger
	parser ottl.Parser[K]

	table  []RoutingTableItem
	routes map[string]routingItem[C, K]

	defaultConsumer  C
	consumerProvider consumerProvider[C]
}

// newRouter creates a new router instance with based on type parameters C and K.
// see router struct definition for the allowed types.
func newRouter[C any, K any](
	table []RoutingTableItem,
	defaultPipelineIDs []string,
	provider consumerProvider[C],
	settings component.TelemetrySettings,
	parser ottl.Parser[K],
) (*router[C, K], error) {
	r := &router[C, K]{
		logger:           settings.Logger,
		parser:           parser,
		table:            table,
		routes:           make(map[string]routingItem[C, K]),
		consumerProvider: provider,
	}

	if err := r.registerConsumers(defaultPipelineIDs); err != nil {
		return nil, err
	}

	return r, nil
}

type routingItem[C any, K any] struct {
	consumer  C
	statement *ottl.Statement[K]
}

func (r *router[C, K]) registerConsumers(defaultPipelineIDs []string) error {
	// register default pipelines
	err := r.registerDefaultConsumer(defaultPipelineIDs)
	if err != nil {
		return err
	}

	// register pipelines for each route
	err = r.registerRouteConsumers()
	if err != nil {
		return err
	}

	return nil
}

// registerDefaultConsumer registers a consumer for the default
// pipelines configured
func (r *router[C, K]) registerDefaultConsumer(pipelineIDs []string) error {
	if len(pipelineIDs) == 0 {
		return nil
	}

	consumer, err := r.consumerForPipelines(pipelineIDs)
	if err != nil {
		return fmt.Errorf("%w: %s", errPipelineNotFound, err)
	}

	r.defaultConsumer = consumer

	return nil
}

// registerRouteConsumers registers a consumer for the pipelines configured
// for each route
func (r *router[C, K]) registerRouteConsumers() error {
	for _, item := range r.table {
		statement, err := r.getStatementFrom(item)
		if err != nil {
			return err
		}

		route, ok := r.routes[key(item)]
		if !ok {
			route.statement = statement
		}

		consumer, err := r.consumerForPipelines(item.Pipelines)
		if err != nil {
			return fmt.Errorf("%w: %s", errPipelineNotFound, err)
		}
		route.consumer = consumer

		r.routes[key(item)] = route
	}
	return nil
}

func (r *router[C, K]) consumerForPipelines(pipelines []string) (C, error) {
	pipelineIDs, err := toComponentIDs(pipelines)
	if err != nil {
		var c C
		return c, err
	}

	return r.consumerProvider(pipelineIDs...)
}

// getStatementFrom builds a routing OTTL statement from the provided
// routing table entry configuration. If the routing table entry configuration
// does not contain a valid OTTL statement then nil is returned.
func (r *router[C, K]) getStatementFrom(item RoutingTableItem) (*ottl.Statement[K], error) {
	var statement *ottl.Statement[K]
	if item.Statement != "" {
		var err error
		statement, err = r.parser.ParseStatement(item.Statement)
		if err != nil {
			return statement, err
		}
	}
	return statement, nil
}

func (r *router[C, K]) getConsumer(key string) C {
	item, ok := r.routes[key]
	if !ok {
		return r.defaultConsumer
	}
	return item.consumer
}

func key(entry RoutingTableItem) string {
	return entry.Statement
}

func toComponentIDs(names []string) ([]component.ID, error) {
	componentIDs := make([]component.ID, 0, len(names))

	for _, name := range names {
		id := component.ID{}
		if err := id.UnmarshalText([]byte(name)); err != nil {
			return nil, err
		}
		componentIDs = append(componentIDs, id)
	}

	return componentIDs, nil
}
