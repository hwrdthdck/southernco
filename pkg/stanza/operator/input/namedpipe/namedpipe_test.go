package namedpipe

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// filename attempts to get an unused filename.
func filename(t testing.TB) string {
	t.Helper()

	file, err := os.CreateTemp("", "")
	require.NoError(t, err)

	name := file.Name()
	require.NoError(t, file.Close())
	require.NoError(t, os.Remove(name))

	return name
}

// TestCreatePipe tests that the named pipe is created as a pipe with the correct permissions.
func TestCreatePipe(t *testing.T) {
	conf := NewConfig()
	conf.Path = filename(t)
	conf.Permissions = 0666

	op, err := conf.Build(zaptest.NewLogger(t).Sugar())
	require.NoError(t, err)

	require.NoError(t, op.Start(testutil.NewUnscopedMockPersister()))
	defer func() {
		require.NoError(t, op.Stop())
	}()

	stat, err := os.Stat(conf.Path)
	require.NoError(t, err)

	isPipe := stat.Mode()&os.ModeNamedPipe != 0
	require.True(t, isPipe, "file is not a named pipe")
	require.Equal(t, conf.Permissions, uint32(stat.Mode().Perm()))
}

// TestPipeWrites writes a few logs to the pipe over a few different connections and verifies that they are received.
func TestPipeWrites(t *testing.T) {
	fake := testutil.NewFakeOutput(t)

	conf := NewConfig()
	conf.Path = filename(t)
	conf.Permissions = 0666
	conf.OutputIDs = []string{fake.ID()}

	op, err := conf.Build(zaptest.NewLogger(t).Sugar())
	require.NoError(t, err)
	ops := []operator.Operator{op, fake}

	p, err := pipeline.NewDirectedPipeline(ops)
	require.NoError(t, err)

	require.NoError(t, p.Start(testutil.NewUnscopedMockPersister()))
	defer p.Stop()

	logs := [][]string{
		{"log1\n", "log2\n"},
		{"log3\n", "log4\n"},
		{"log5\n"},
	}

	for _, toSend := range logs {
		pipe, err := os.OpenFile(conf.Path, os.O_WRONLY, 0)
		require.NoError(t, err)
		defer pipe.Close()

		for _, log := range toSend {
			_, err = pipe.WriteString(log)
			require.NoError(t, err)
		}

		for _, log := range toSend {
			expect := &entry.Entry{
				Body: strings.TrimSpace(log),
			}

			select {
			case e := <-fake.Received:
				obs := time.Now()
				expect.ObservedTimestamp = obs
				e.ObservedTimestamp = obs
				require.Equal(t, expect, e)
			case <-time.After(time.Second):
				t.Fatal("timed out waiting for entry")
			}
		}
	}
}
