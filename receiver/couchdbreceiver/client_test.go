// Copyright  The OpenTelemetry Authors
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

package couchdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configtls"
	"go.uber.org/zap"
)

func defaultConfig(t *testing.T, endpoint string) CouchDBClient {
	couchdbClient, err := NewCouchDBClient(
		&Config{
			Username: "otelu",
			Password: "otelp",
			HTTPClientSettings: confighttp.HTTPClientSettings{
				Endpoint: endpoint,
			},
		},
		componenttest.NewNopHost(),
		zap.NewNop())
	require.Nil(t, err)
	require.NotNil(t, couchdbClient)
	return couchdbClient
}

func TestNewCouchDBClient(t *testing.T) {
	t.Run("Invalid config", func(t *testing.T) {
		couchdbClient, err := NewCouchDBClient(
			&Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: defaultEndpoint,
					TLSSetting: configtls.TLSClientSetting{
						TLSSetting: configtls.TLSSetting{
							CAFile: "/non/existent",
						},
					},
				}},
			componenttest.NewNopHost(),
			zap.NewNop())

		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to create HTTP Client: ")
		require.Nil(t, couchdbClient)
	})
	t.Run("no error", func(t *testing.T) {
		client, err := NewCouchDBClient(
			&Config{},
			componenttest.NewNopHost(),
			zap.NewNop(),
		)

		require.NoError(t, err)
		require.NotNil(t, client)
	})
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		if u == "" || p == "" {
			w.WriteHeader(401)
			return
		}
		if strings.Contains(r.URL.Path, "/_stats/couchdb") {
			w.WriteHeader(200)
			return
		}
		if strings.Contains(r.URL.Path, "/invalid_endpoint") {
			w.WriteHeader(404)
			return
		}
		if strings.Contains(r.URL.Path, "/invalid_body") {
			w.Header().Set("Content-Length", "1")
			return
		}

		w.WriteHeader(404)
	}))
	defer ts.Close()

	t.Run("invalid url request", func(t *testing.T) {
		url := ts.URL + " /space"
		couchdbClient := defaultConfig(t, url)

		result, err := couchdbClient.Get(url)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Contains(t, err.Error(), "invalid port ")
	})
	t.Run("invalid endpoint", func(t *testing.T) {
		url := ts.URL + "/invalid_endpoint"
		couchdbClient := defaultConfig(t, url)

		result, err := couchdbClient.Get(url)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Contains(t, err.Error(), "404 Not Found")
	})
	t.Run("invalid body", func(t *testing.T) {
		url := ts.URL + "/invalid_body"
		couchdbClient := defaultConfig(t, url)

		result, err := couchdbClient.Get(url)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Contains(t, err.Error(), "failed to read response body ")
	})
	t.Run("401 Unauthorized", func(t *testing.T) {
		url := ts.URL + "/_node/_local/_stats/couchdb"
		couchdbClient, err := NewCouchDBClient(
			&Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: url,
				},
			},
			componenttest.NewNopHost(),
			zap.NewNop())
		require.Nil(t, err)
		require.NotNil(t, couchdbClient)

		result, err := couchdbClient.Get(url)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Contains(t, err.Error(), "401 Unauthorized")
	})
	t.Run("no error", func(t *testing.T) {
		url := ts.URL + "/_node/_local/_stats/couchdb"
		couchdbClient := defaultConfig(t, url)

		result, err := couchdbClient.Get(url)
		require.Nil(t, err)
		require.NotNil(t, result)
	})
}

func TestGetNodeNamesSingle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/invalid_json") {
			w.WriteHeader(200)
			_, _ = w.Write([]byte(`{"}`))
			return
		}

		if r.URL.Path == nodeNamesPath {
			w.WriteHeader(200)
			w.Write([]byte(`{"all_nodes":["nonode@nohost"],"cluster_nodes":["nonode@nohost"]}`))
			return
		}
		w.WriteHeader(404)
	}))
	defer ts.Close()

	t.Run("invalid endpoint", func(t *testing.T) {
		couchdbClient := defaultConfig(t, "invalid")

		actualNodeNames, err := couchdbClient.GetNodeNames()
		require.NotNil(t, err)
		require.Nil(t, actualNodeNames)
	})
	t.Run("invalid json", func(t *testing.T) {
		couchdbClient := defaultConfig(t, ts.URL+"/invalid_json")

		actualNodeNames, err := couchdbClient.GetNodeNames()
		require.NotNil(t, err)
		require.Nil(t, actualNodeNames)
	})
	t.Run("no error", func(t *testing.T) {
		expectedNodeNames := []string{"nonode@nohost"}
		couchdbClient := defaultConfig(t, ts.URL)

		actualNodeNames, err := couchdbClient.GetNodeNames()
		require.Nil(t, err)
		require.EqualValues(t, expectedNodeNames, actualNodeNames)
	})
}

func TestGetNodeNamesMultiple(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == nodeNamesPath {
			w.WriteHeader(200)
			w.Write([]byte(`{"all_nodes":["couchdb@couchdb0.otel.com","couchdb@couchdb1.otel.com","couchdb@couchdb2.otel.com"],"cluster_nodes":["couchdb@couchdb0.otel.com","couchdb@couchdb1.otel.com","couchdb@couchdb2.otel.com"]}`))
			return
		}
		w.WriteHeader(404)
	}))
	defer ts.Close()

	expectedNodeNames := []string{"couchdb@couchdb0.otel.com", "couchdb@couchdb1.otel.com", "couchdb@couchdb2.otel.com"}
	couchdbClient := defaultConfig(t, ts.URL)

	actualNodeNames, err := couchdbClient.GetNodeNames()
	require.Nil(t, err)
	require.EqualValues(t, expectedNodeNames, actualNodeNames)
}

func TestGetNodeStats(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.Path, "/invalid_json") {
			w.WriteHeader(200)
			_, _ = w.Write([]byte(`{"}`))
			return
		}
		if strings.Contains(r.URL.Path, "/_stats/couchdb") {
			w.WriteHeader(200)
			w.Write([]byte(`{"key":["value"]}`))
			return
		}
		w.WriteHeader(404)
	}))
	defer ts.Close()

	t.Run("invalid endpoint", func(t *testing.T) {
		couchdbClient := defaultConfig(t, "invalid")

		actualStats, err := couchdbClient.GetStats("_local")
		require.NotNil(t, err)
		require.Nil(t, actualStats)
	})
	t.Run("invalid json", func(t *testing.T) {
		couchdbClient := defaultConfig(t, ts.URL+"/invalid_json")

		actualStats, err := couchdbClient.GetStats("_local")
		require.NotNil(t, err)
		require.Nil(t, actualStats)
	})
	t.Run("no error", func(t *testing.T) {
		expectedStats := map[string]interface{}{"key": []interface{}{"value"}}
		couchdbClient := defaultConfig(t, ts.URL)

		actualStats, err := couchdbClient.GetStats("_local")
		require.Nil(t, err)
		require.EqualValues(t, expectedStats, actualStats)
	})
}

func TestBuildReq(t *testing.T) {
	couchdbClient := couchDBClient{
		client: &http.Client{},
		cfg: &Config{
			HTTPClientSettings: confighttp.HTTPClientSettings{
				Endpoint: defaultEndpoint,
			},
			Username: "otelu",
			Password: "otelp",
		},
		logger: zap.NewNop(),
	}

	req, err := couchdbClient.buildReq(nodeNamesPath)
	require.NoError(t, err)
	require.NotNil(t, req)
	require.Equal(t, "application/json", req.Header["Content-Type"][0])
	require.Equal(t, []string{"Basic b3RlbHU6b3RlbHA="}, req.Header["Authorization"])
	require.Equal(t, defaultEndpoint+nodeNamesPath, req.URL.String())
}

func TestBuildBadReq(t *testing.T) {
	couchdbClient := couchDBClient{
		client: &http.Client{},
		cfg: &Config{
			HTTPClientSettings: confighttp.HTTPClientSettings{
				Endpoint: defaultEndpoint,
			},
		},
		logger: zap.NewNop(),
	}

	_, err := couchdbClient.buildReq(" ")
	require.Error(t, err)
}
