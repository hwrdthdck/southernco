// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/multierr"
)

const HeaderRetryAfter = "Retry-After"

// HandleHTTPCode handles an http response and returns the right type of error in case of a failure.
func HandleHTTPCode(resp *http.Response) error {
	// Splunk accepts all 2XX codes.
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	err := fmt.Errorf(
		"HTTP %d %q",
		resp.StatusCode,
		http.StatusText(resp.StatusCode))

	// Check if there is any error text returned by Splunk that we can append to error message
	var jsonResponse map[string]any
	bodyString, _ := io.ReadAll(resp.Body)
	e := json.Unmarshal(bodyString, &jsonResponse)
	
	if e == nil {
		errorValue, ok := jsonResponse["text"]
		if ok {
			err = multierr.Append(err, fmt.Errorf("Reason: %q", errorValue))
		}
	}
	
	switch resp.StatusCode {
	// Check for responses that may include "Retry-After" header.
	case http.StatusTooManyRequests, http.StatusServiceUnavailable:
		// Fallback to 0 if the Retry-After header is not present. This will trigger the
		// default backoff policy by our caller (retry handler).
		retryAfter := 0
		if val := resp.Header.Get(HeaderRetryAfter); val != "" {
			if seconds, err2 := strconv.Atoi(val); err2 == nil {
				retryAfter = seconds
			}
		}
		// Indicate to our caller to pause for the specified number of seconds.
		err = exporterhelper.NewThrottleRetry(err, time.Duration(retryAfter)*time.Second)
	// Check for permanent errors.
	case http.StatusBadRequest, http.StatusUnauthorized:
		dump, err2 := httputil.DumpResponse(resp, true)
		if err2 == nil {
			err = consumererror.NewPermanent(fmt.Errorf("%w", fmt.Errorf("%q", dump)))
		} else {
			err = multierr.Append(err, err2)
		}
	}

	return err
}
