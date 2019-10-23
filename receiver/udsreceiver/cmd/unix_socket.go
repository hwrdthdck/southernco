// Copyright 2019, OpenTelementry Authors
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

// +build !windows

package main

import (
	"flag"

	"github.com/oklog/run"

        "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udsreceiver"
        "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udsreceiver/transport/unixsocket"
)

const defaultUnixSocketPath = "/tmp/ot-daemon.sock"

func addTransportPath(fs *flag.FlagSet) *string {
	return fs.String("socket.path", defaultUnixSocketPath, "Unix socket path to listen on")
}

func svcTransport(g *run.Group, hnd udsreceiver.Handler, socketPath string) {
	uss := unixsocket.New(socketPath, hnd)

	g.Add(func() error {
		return uss.ListenAndServe()
	}, func(error) {
		_ = uss.Close()
	})

}
