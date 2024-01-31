// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"testing"

	"go.uber.org/goleak"
)

// The IgnoreTopFunction call prevents catching the leak generated by opencensus
// defaultWorker.Start which at this time is part of the package's init call.
// See https://github.com/census-instrumentation/opencensus-go/issues/1191 for more information.
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start"))
}
