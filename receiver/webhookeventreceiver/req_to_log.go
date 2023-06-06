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

package webhookeventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"

import (
	"bufio"
	"net/url"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/plog"
)

func reqToLog(sc *bufio.Scanner,
	query url.Values,
	_config *Config,
	idName string) (plog.Logs, int) {
	log := plog.NewLogs()
	resourceLog := log.ResourceLogs().AppendEmpty()
	appendMetadata(resourceLog, query)
	scopeLog := resourceLog.ScopeLogs().AppendEmpty()

	scopeLog.Scope().Attributes().PutStr("source", idName)
	scopeLog.Scope().Attributes().PutStr("receiver", metadata.Type)

	for sc.Scan() {
		logRecord := scopeLog.LogRecords().AppendEmpty()
		line := sc.Text()
		logRecord.Body().SetStr(line)
	}

	return log, scopeLog.LogRecords().Len()
}

// append query parameters and webhook source as resource attributes
func appendMetadata(resourceLog plog.ResourceLogs, query url.Values) {
	for k := range query {
		if query.Get(k) != "" {
			resourceLog.Resource().Attributes().PutStr(k, query.Get(k))
		}
	}

}
