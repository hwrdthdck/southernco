// Copyright 2020, OpenTelemetry Authors
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

package httpClient

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
)

type fakeClient struct {
	response *http.Response
	err      error
}

func (f *fakeClient) Do(req *retryablehttp.Request) (*http.Response, error) {
	return f.response, f.err
}
func TestRequestSuccessWithKnownLength(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: 5 * 1024,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      nil,
	}

	httpFake := New(withClientOption(fakeClient))

	ctx := context.Background()

	body, err := httpFake.Request(ctx, "0.0.0.0")

	assert.Nil(t, err)

	assert.NotNil(t, body)

}

func TestRequestSuccessWithUnknownLength(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: -1,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      nil,
	}

	httpFake := New(withClientOption(fakeClient))

	ctx := context.Background()

	body, err := httpFake.Request(ctx, "0.0.0.0")

	assert.Nil(t, err)

	assert.NotNil(t, body)

}

func TestRequestWithFailedStatus(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		Status:        "Bad Request",
		StatusCode:    400,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: 5 * 1024,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      errors.New(""),
	}

	httpFake := New(withClientOption(fakeClient))

	ctx := context.Background()

	body, err := httpFake.Request(ctx, "0.0.0.0")

	assert.Nil(t, body)

	assert.NotNil(t, err)

}

func TestRequestWithLargeContentLength(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: 5 * 1024 * 1024,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      nil,
	}

	httpFake := New(withClientOption(fakeClient))

	ctx := context.Background()

	body, err := httpFake.Request(ctx, "0.0.0.0")

	assert.Nil(t, body)

	assert.NotNil(t, err)

}
