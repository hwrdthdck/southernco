// Code generated by mdatagen. DO NOT EDIT.

package dockerstatsreceiver

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/testcontainers/testcontainers-go.(*Reaper).Connect.func1"), goleak.IgnoreTopFunction("net/http.(*persistConn).writeLoop"), goleak.IgnoreTopFunction("internal/poll.runtime_pollWait"))
}
