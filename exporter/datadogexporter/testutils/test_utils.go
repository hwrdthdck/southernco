// Copyright The OpenTelemetry Authors
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

package testutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"go.opentelemetry.io/collector/consumer/pdata"
)

type DatadogServer struct {
	*httptest.Server
	MetadataChan chan []byte
}

// DatadogServerMock mocks a Datadog backend server
func DatadogServerMock() *DatadogServer {
	metadataChan := make(chan []byte)
	handler := http.NewServeMux()
	handler.HandleFunc("/api/v1/validate", validateAPIKeyEndpoint)
	handler.HandleFunc("/api/v1/series", metricsEndpoint)
	handler.HandleFunc("/intake", newMetadataEndpoint(metadataChan))

	srv := httptest.NewServer(handler)

	return &DatadogServer{
		srv,
		metadataChan,
	}
}

type validateAPIKeyResponse struct {
	Valid bool `json:"valid"`
}

func validateAPIKeyEndpoint(w http.ResponseWriter, r *http.Request) {
	res := validateAPIKeyResponse{Valid: true}
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

type metricsResponse struct {
	Status string `json:"status"`
}

func metricsEndpoint(w http.ResponseWriter, r *http.Request) {
	res := metricsResponse{Status: "ok"}
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resJSON)
}

func newMetadataEndpoint(c chan []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		c <- body
	}
}

// NewAttributeMap creates a new attribute map (string only)
// from a Go map
func NewAttributeMap(mp map[string]string) pdata.AttributeMap {
	attrs := pdata.NewAttributeMap()
	attrs.InitEmptyWithCapacity(len(mp))

	for k, v := range mp {
		attrs.Insert(k, pdata.NewAttributeValueString(v))
	}

	return attrs
}
