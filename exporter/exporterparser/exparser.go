// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This package provides support for parsing and creating the
// respective exporters given a YAML configuration payload.
// For now it currently only provides statically imported OpenCensus
// exporters like:
//  * Stackdriver Tracing and Monitoring
//  * DataDog
//  * Zipkin
package exporterparser

import (
	"fmt"

	"go.opencensus.io/trace"
	yaml "gopkg.in/yaml.v2"
)

func exportSpans(te trace.Exporter, spandata []*trace.SpanData) error {
	for _, sd := range spandata {
		te.ExportSpan(sd)
	}
	return nil
}

func yamlUnmarshal(yamlBlob []byte, dest interface{}) error {
	if err := yaml.Unmarshal(yamlBlob, dest); err != nil {
		return fmt.Errorf("Cannot YAML unmarshal data: %v", err)
	}
	return nil
}
