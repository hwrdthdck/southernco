// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stefexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stefexporter"

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	stefgrpc "github.com/splunk/stef/go/grpc"
	"github.com/splunk/stef/go/grpc/stef_proto"
	"github.com/splunk/stef/go/otel/oteltef"
	stefpdatametrics "github.com/splunk/stef/go/pdata/metrics"
	"github.com/splunk/stef/go/pkg"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// stefExporter implements sending metrics over STEF/gRPC stream.
// The exporter uses a single stream and accepts concurrent exportMetrics calls,
// sequencing the metric data as needed over a single stream.
// The exporter will block exportMetrics call until an acknowledgement is
// received from destination. To achieve high throughput, appropriate level of
// concurrency is necessary, so make sure to configure the exporter helper Queue
// with sufficiently high number of consumers.
// The exporter relies on a preceding Retry helper to retry sending data that is
// not acknowledged or otherwise fails to be sent. The exporter will not retry
// sending the data itself.
type stefExporter struct {
	set         component.TelemetrySettings
	host        component.Host
	logger      *zap.Logger
	cfg         *Config
	compression pkg.Compression

	// How long to wait for ack to be received from server. Configurable for testing purposes.
	maxAckWaitTime time.Duration

	// connMutex is taken when connecting, disconnecting or checking connection status.
	connMutex   sync.Mutex
	isConnected bool
	grpcConn    *grpc.ClientConn

	// Maximum number of concurrent write operations we expect.
	maxWritesInProgress int64

	// Number of write operations currently in progress.
	writesInProgress atomic.Int64

	// The STEF writer we write metrics to and which in turns sends them over gRPC.
	stefWriter      *oteltef.MetricsWriter
	stefWriterMutex sync.Mutex // protects stefWriter

	// Ack-related fields follow and are protected by ackDataMutex.
	ackDataMutex sync.Mutex
	// Last ack ID received from the server.
	lastAckedID uint64
	// Flag to set to interrupt all waiting for acks. This is set when disconnecting.
	interruptAckWaits bool
	// A slice of deadlines waiting for acks.
	ackExpectedBy []time.Time
	// Condition used to notify ack waiters. Works in pair with ackDataMutex.
	ackCond *sync.Cond

	// Indicates ack wait operations that are ongoing.
	inProgressWaits sync.WaitGroup
	// Timer that fires when ack waiting time exceeds maxAckWaitTime.
	ackTimer *time.Timer

	// Channel to stop ticker() goroutine.
	stopTicker chan struct{}
}

type loggerWrapper struct {
	logger *zap.Logger
}

func (w *loggerWrapper) Debugf(_ context.Context, format string, v ...any) {
	w.logger.Debug(fmt.Sprintf(format, v...))
}

func (w *loggerWrapper) Errorf(_ context.Context, format string, v ...any) {
	w.logger.Error(fmt.Sprintf(format, v...))
}

// By default, wait up to 30 seconds for acks to arrive from the server.
const defMaxAckWaitTime = 30 * time.Second

var errAckTimeout = errors.New("ack timed out")

func newStefExporter(set component.TelemetrySettings, cfg *Config) *stefExporter {
	exp := &stefExporter{
		set:                 set,
		logger:              set.Logger,
		cfg:                 cfg,
		maxWritesInProgress: int64(cfg.NumConsumers),
		ackTimer:            time.NewTimer(time.Duration(1<<63 - 1)),
		maxAckWaitTime:      cfg.maxAckWaitTime,
	}

	exp.ackCond = sync.NewCond(&exp.ackDataMutex)
	if exp.maxAckWaitTime == 0 {
		exp.maxAckWaitTime = defMaxAckWaitTime
	}

	exp.compression = pkg.CompressionNone
	if cfg.Compression == "zstd" {
		exp.compression = pkg.CompressionZstd
	}
	return exp
}

func (s *stefExporter) Start(ctx context.Context, host component.Host) error {
	s.host = host

	// No need to block Start(), we will begin connection attempt in a goroutine.
	go func() {
		if err := s.ensureConnected(ctx, host); err != nil {
			s.logger.Error("Error connecting to destination", zap.Error(err))
			// exportMetrics() will try to connect again as needed.
		}
	}()
	return nil
}

func (s *stefExporter) ensureConnected(ctx context.Context, host component.Host) error {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	if s.isConnected {
		return nil
	}

	s.logger.Debug("Connecting to destination", zap.String("endpoint", s.cfg.Endpoint))

	// Initialize ack-related fields.
	s.lastAckedID = 0
	s.interruptAckWaits = false
	s.ackExpectedBy = s.ackExpectedBy[:0]

	// Connect to the server.
	var err error
	s.grpcConn, err = s.cfg.ClientConfig.ToClientConn(ctx, host, s.set)
	if err != nil {
		return err
	}

	// Open a STEF/gRPC stream to the server.
	grpcClient := stef_proto.NewSTEFDestinationClient(s.grpcConn)

	// Let server know about our schema.
	schema, err := oteltef.MetricsWireSchema()
	if err != nil {
		return err
	}

	settings := stefgrpc.ClientSettings{
		Logger:       &loggerWrapper{s.logger},
		GrpcClient:   grpcClient,
		ClientSchema: schema,
		Callbacks: stefgrpc.ClientCallbacks{
			OnAck: s.onGrpcAck,
		},
	}
	client := stefgrpc.NewClient(settings)

	grpcWriter, opts, err := client.Connect(context.Background())
	if err != nil {
		s.grpcConn.Close()
		return err
	}

	opts.Compression = s.compression

	// Create record writer over gRPC stream.
	s.stefWriter, err = oteltef.NewMetricsWriter(grpcWriter, opts)
	if err != nil {
		return err
	}

	// Start background timing tasks.
	s.stopTicker = make(chan struct{})
	go s.ticker()

	s.isConnected = true

	return nil
}

func (s *stefExporter) disconnect() {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	if !s.isConnected {
		return
	}

	s.logger.Debug("Disconnecting...")

	// Notify ticker() to stop.
	close(s.stopTicker)

	// Interrupt any pending waitForAck() call. They can't succeed anymore.
	s.ackDataMutex.Lock()
	s.interruptAckWaits = true
	s.ackExpectedBy = s.ackExpectedBy[:0]
	s.ackCond.Broadcast()
	s.ackDataMutex.Unlock()

	// If there is an existing connection close it.
	if s.grpcConn != nil {
		if err := s.grpcConn.Close(); err != nil {
			s.logger.Error("failed to close grpc connection", zap.Error(err))
		}
	}

	// Wait for all in progress wait in goroutines to get the notification that
	// we are disconnecting and exit.
	s.inProgressWaits.Wait()

	s.isConnected = false
}

func (s *stefExporter) Shutdown(_ context.Context) error {
	s.disconnect()
	return nil
}

func (s *stefExporter) exportMetrics(ctx context.Context, md pmetric.Metrics) error {
	// Keep count of how many concurrent writes are in progress.
	s.writesInProgress.Add(1)
	defer s.writesInProgress.Add(-1)

	// Write the metrics to STEF stream.
	expectedAckID, err := s.writeMetrics(ctx, md)
	if err != nil {
		return err
	}

	// Wait for acknowledgement from destination.
	if err := s.waitForAck(expectedAckID); err != nil {
		return err
	}

	return nil
}

func (s *stefExporter) writeMetrics(ctx context.Context, md pmetric.Metrics) (uint64, error) {
	if err := s.ensureConnected(ctx, s.host); err != nil {
		return 0, err
	}

	// stefWriter is not safe for concurrent writing, protect it.
	s.stefWriterMutex.Lock()
	defer s.stefWriterMutex.Unlock()

	converter := stefpdatametrics.OtlpToTEFUnsorted{}
	err := converter.WriteMetrics(md, s.stefWriter)
	if err != nil {
		// TODO: check if err is because STEF encoding failed. If so we must not
		// try to re-encode the same data. Return consumererror.NewPermanent(err)
		// to the caller.
		s.disconnect()

		// Return an error to retry sending these metrics again next time.
		return 0, err
	}

	// According to STEF gRPC spec the destination ack IDs match written record number.
	// When the data we have just written is received by destination it will send us
	// back and ack ID that numerically matches the last written record number.
	expectedAckID := s.stefWriter.RecordCount()

	if s.writesInProgress.Load() >= s.maxWritesInProgress {
		// We achieved max concurrency. No further exportMetrics calls will make
		// progress until at least some of the exportMetrics calls return. For
		// calls to return they need to receive an Ack. For Ack to be received
		// we need to make the data we wrote is actually sent, so we need to
		// issue a Flush call.
		if err = s.stefWriter.Flush(); err != nil {
			// Failure to write the gRPC stream normally means something is
			// wrong with the connection. We will reconnect.
			s.disconnect()

			// Return an error to retry sending these metrics again next time.
			return 0, err
		}
	}

	return expectedAckID, nil
}

func (s *stefExporter) waitForAck(ackID uint64) error {
	s.inProgressWaits.Add(1)
	defer s.inProgressWaits.Done()

	// Calculate deadline when wait for this ack expires.
	ackDeadline := time.Now().Add(s.maxAckWaitTime)

	s.ackCond.L.Lock()

	// Add deadline to the list of deadline times.
	s.ackExpectedBy = append(s.ackExpectedBy, ackDeadline)

	if len(s.ackExpectedBy) == 1 {
		// The list was empty, we were not waiting for any acks.
		// Start the timer to wait for acks.
		s.ackTimer.Reset(s.maxAckWaitTime)
	}

	for s.lastAckedID < ackID && // Stop if ack ID we wait for is received
		!s.interruptAckWaits && // Stop if all ack waits are interrupted
		time.Now().Before(ackDeadline) { // Stop if this particular ack waiting expires
		// Wait to be notified when one of the above 3 conditions change.
		s.ackCond.Wait()
	}

	// After s.ackCond.Wait() returns we are holding the lock on s.ackCond.L so it
	// is safe to work with the ack-related fields.

	ackReceived := s.lastAckedID >= ackID

	if i := slices.Index(s.ackExpectedBy, ackDeadline); i >= 0 {
		// Remove the deadline time for this ack from the slice.
		s.ackExpectedBy = slices.Delete(s.ackExpectedBy, i, i+1)
	}

	if len(s.ackExpectedBy) > 0 {
		// Start the timer again till the deadline of the next nearest ack timeout.
		intervalToNextDeadline := time.Until(s.ackExpectedBy[0])
		s.ackTimer.Reset(intervalToNextDeadline)
	}

	s.ackCond.L.Unlock()

	if ackReceived {
		// We received the ack we were waiting for.
		return nil
	}

	s.logger.Debug("Waiting for ack is interrupted.")

	return errAckTimeout
}

func (s *stefExporter) onGrpcAck(ackID uint64) error {
	s.ackDataMutex.Lock()
	defer s.ackDataMutex.Unlock()
	if s.lastAckedID < ackID {
		s.lastAckedID = ackID

		// Let all waitForAck() calls know we received an ack.
		s.ackCond.Broadcast()
	}
	return nil
}

func (s *stefExporter) ticker() {
	// This goroutine monitors 2 timeouts:
	// 1. To flush accumulated data.
	// 2. To notify waiters when ack waiting time expires.

	flushTimer := time.NewTicker(100 * time.Millisecond)
	defer flushTimer.Stop()

	for {
		select {
		case <-flushTimer.C:
			s.stefWriterMutex.Lock()
			err := s.stefWriter.Flush()
			s.stefWriterMutex.Unlock()
			if err != nil {
				// If Flush() fails something is wrong with the connection.
				// We need to reconnect.
				s.logger.Error("Error flushing data. Will disconnect to reconnect.", zap.Error(err))
				s.disconnect()
				return
			}

		case <-s.ackTimer.C:
			// Timer waiting for ack expired and we did not receive the ack.
			// Server is unresponsive or connection is broken.
			// We need to reconnect.
			s.logger.Error("Timeout waiting for server to ack. Will disconnect to reconnect.")
			s.disconnect()
			return

		case <-s.stopTicker:
			return
		}
	}
}
