// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gelfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

var (
	defaultTCPTimeout      = 2000
	defaultChunkSize       = 1420
	magicChunked           = []byte{0x1e, 0x0f}
	magicZlib              = []byte{0x78}
	magicGzip              = []byte{0x1f, 0x8b}
	chunkedHeaderLen       = 12
	defaultCompressionType = "gzip"
)

type gelfexporter struct {
	chunkedDataLen int
	config         *Config
	logger         *zap.Logger
	conn           net.Conn
}

type GELFMessage struct {
	Version   string                 `json:"version"`
	Host      string                 `json:"host"`
	Short     string                 `json:"short_message"`
	Full      string                 `json:"full_message,omitempty"`
	Timestamp float64                `json:"timestamp"`
	Level     int32                  `json:"level,omitempty"`
	Facility  string                 `json:"facility,omitempty"`
	Extra     map[string]interface{} `json:"-"`
	RawExtra  json.RawMessage        `json:"-"`
}

func logLevelToSeverity(logLevel string) int32 {
	switch logLevel {
	case "TRACE", "TRACE2", "TRACE3", "TRACE4", "DEBUG", "DEBUG2", "DEBUG3", "DEBUG4":
		return int32(7)
	case "INFO", "INFO2", "INFO3", "INFO4":
		return int32(6)
	case "WARN", "WARN2", "WARN3", "WARN4":
		return int32(4)
	case "ERROR", "ERROR2", "ERROR3", "ERROR4":
		return int32(3)
	case "FATAL", "FATAL2", "FATAL3", "FATAL4":
		return int32(2)
	default:
		return int32(6) // Default or unknown severity
	}
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

func newBuffer() *bytes.Buffer {
	b := bufPool.Get().(*bytes.Buffer)
	if b != nil {
		b.Reset()
		return b
	}
	return bytes.NewBuffer(nil)
}

func makeConn(protocol string, endpoint string) (net.Conn, error) {
	var err error
	var conn net.Conn
	switch protocol {
	case "udp":
		udpAddr, err := net.ResolveUDPAddr(protocol, endpoint)
		if err != nil {
			fmt.Println("Error resolving UDP address:", err)
			os.Exit(1)
		}
		if conn, err = net.DialUDP(protocol, nil, udpAddr); err != nil {
			return nil, err
		}
		return conn, nil
	case "tcp":
		if conn, err = net.Dial(protocol, endpoint); err != nil {
			return nil, err
		}
		return conn, nil
	default:
		return nil, fmt.Errorf("Could not create connection! Invalid protocol: %s", protocol)
	}

}

func initExporter(cfg *Config, createSettings exporter.Settings) (*gelfexporter, error) {

	conn, err := makeConn(cfg.Protocol, cfg.Endpoint)

	if err != nil {
		return nil, fmt.Errorf("Error while creating connection: %w", err)
	}

	g := &gelfexporter{
		config:         cfg,
		logger:         createSettings.Logger,
		chunkedDataLen: cfg.ChunkSize - chunkedHeaderLen,
		conn:           conn,
	}

	g.logger.Info("Gelf Exporter configured",
		zap.String("endpoint", cfg.Endpoint),
		zap.String("protocol", cfg.Protocol),
		zap.String("compression", cfg.CompressionType),
		zap.Bool("send_chunks_with_overflow", cfg.ChunksOverflow),
	)

	return g, nil
}

func newLogsExporter(ctx context.Context,
	params exporter.Settings,
	cfg *Config) (exporter.Logs, error) {

	g, err := initExporter(cfg, params)

	if err != nil {
		return nil, fmt.Errorf("failed to create the logs exporter: %w", err)
	}

	g.logger.Debug("Called newLogsExporter")

	return exporterhelper.NewLogsExporter(
		ctx,
		params,
		cfg,
		g.pushLogsData,
	)

}

// numChunks returns the number of GELF chunks necessary to transmit
// the given compressed buffer.
func (g *gelfexporter) numChunks(b []byte) int {
	lenB := len(b)
	if lenB <= g.config.ChunkSize {
		return 1
	}
	return len(b)/g.chunkedDataLen + 1
}

func (m *GELFMessage) MarshalJSONBuf(buf *bytes.Buffer) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	// write up until the final }
	if _, err = buf.Write(b[:len(b)-1]); err != nil {
		return err
	}
	if len(m.Extra) > 0 {
		eb, err := json.Marshal(m.Extra)
		if err != nil {
			return err
		}
		// merge serialized message + serialized extra map
		if err = buf.WriteByte(','); err != nil {
			return err
		}
		// write serialized extra bytes, without enclosing quotes
		if _, err = buf.Write(eb[1 : len(eb)-1]); err != nil {
			return err
		}
	}

	if len(m.RawExtra) > 0 {
		if err := buf.WriteByte(','); err != nil {
			return err
		}

		// write serialized extra bytes, without enclosing quotes
		if _, err = buf.Write(m.RawExtra[1 : len(m.RawExtra)-1]); err != nil {
			return err
		}
	}

	// write final closing quotes
	return buf.WriteByte('}')
}

func (g *gelfexporter) send(payload []byte) error {
	var err error
	var n int
	if g.conn == nil {
		g.conn, err = makeConn(g.config.Protocol, g.config.Endpoint)
		if err != nil {
			return fmt.Errorf("Error while creating connection: %w", err)
		}
	}

	if g.config.Protocol == "tcp" {
		timeout := time.Duration(g.config.TCPTimeout) * time.Millisecond
		err := g.conn.SetWriteDeadline(time.Now().Add(timeout))
		if err != nil {
			return fmt.Errorf("failed to set write deadline(timeout): %w", err)
		}
	}

	n, err = g.conn.Write(payload)

	if err != nil {
		g.logger.Error("Error while sending response", zap.Error(err))

		// FOR TCP ONLY:
		// Trap "broken pipe" which might happen due to stale broken connections.
		// One case for this is when the upstream server was down momentarily and the connection was closed.
		if errors.Is(err, syscall.EPIPE) || errors.Is(err, syscall.ECONNRESET) {
			g.logger.Debug("Gelf-exporter connection closed, re-creating connection")
			g.conn, err = makeConn(g.config.Protocol, g.config.Endpoint)

			if err != nil {
				return fmt.Errorf("Error while re-creating connection: %w", err)
			}

			n, err = g.conn.Write(payload)

			if err != nil {
				return fmt.Errorf("Write timeout: %w", err)
			}
			g.logger.Debug("Re-sent message")
		}
		return fmt.Errorf("Error while sending response: %w", err)
	}

	if n != len(payload) {
		return fmt.Errorf("Error while validating UDP response, (received length/payload length): (%d/%d)", n, len(payload))
	}

	return nil
}

func (g *gelfexporter) writeChunked(ctx context.Context, gelfBytes []byte) (err error) {
	b := make([]byte, 0, g.config.ChunkSize)
	buf := bytes.NewBuffer(b)
	nChunksI := g.numChunks(gelfBytes)
	if nChunksI > 128 {
		if g.config.ChunksOverflow {
			nChunksI = 128
		} else {
			return fmt.Errorf("msg too large, would need %d chunks", nChunksI)
		}
	}
	nChunks := uint8(nChunksI)
	// use urandom to get a unique message id
	msgId := make([]byte, 8)
	n, err := io.ReadFull(rand.Reader, msgId)
	if err != nil || n != 8 {
		return fmt.Errorf("rand.Reader: %d/%s", n, err)
	}

	bytesLeft := len(gelfBytes)
	for i := uint8(0); i < nChunks; i++ {
		buf.Reset()
		// manually write header.  Don't care about
		// host/network byte order, because the spec only
		// deals in individual bytes.
		buf.Write(magicChunked) //magic
		buf.Write(msgId)
		buf.WriteByte(i)
		buf.WriteByte(nChunks)
		// slice out our chunk from gelfBytes
		chunkLen := g.chunkedDataLen
		if chunkLen > bytesLeft {
			chunkLen = bytesLeft
		}
		off := int(i) * g.chunkedDataLen
		chunk := gelfBytes[off : off+chunkLen]
		buf.Write(chunk)

		fmt.Println("Output bytes:%X", buf.Bytes())

		// write this chunk, and make sure the write was good
		err := g.send(buf.Bytes())
		if err != nil {
			return fmt.Errorf("Write (chunk %d/%d): %s", i,
				nChunks, err)
		}

		bytesLeft -= chunkLen
	}

	if bytesLeft != 0 {
		return fmt.Errorf("error: %d bytes left after sending", bytesLeft)
	}
	return nil
}

func (g *gelfexporter) convertLogsToGELF(ctx context.Context, incomingTimestamp float64, severity int32, body string) error {

	var err error

	short := []byte(body)
	full := []byte("")
	if i := bytes.IndexRune(short, '\n'); i > 0 {
		full = short
		short = short[:i]

	}

	message := GELFMessage{
		Version:   "1.1",
		Host:      g.config.Hostname,
		Short:     string(short),
		Full:      string(full),
		Timestamp: incomingTimestamp,
		Level:     severity,
	}

	// messageBytes, err := json.Marshal(message)

	messageBuffer := newBuffer()
	defer bufPool.Put(messageBuffer)
	if err = message.MarshalJSONBuf(messageBuffer); err != nil {
		return err
	}
	messageBytes := messageBuffer.Bytes()

	var (
		compressedBuf *bytes.Buffer
		gelfBytes     []byte
	)

	var zw io.WriteCloser

	switch g.config.CompressionType {
	case "zlib":
		compressedBuf = newBuffer()
		defer bufPool.Put(compressedBuf)
		zw, err = zlib.NewWriterLevel(compressedBuf, flate.BestSpeed)
	case "gzip":
		compressedBuf = newBuffer()
		defer bufPool.Put(compressedBuf)
		zw, err = gzip.NewWriterLevel(compressedBuf, flate.BestSpeed)
	case "none":
		gelfBytes = messageBytes
	default:
		return fmt.Errorf("Invalid compression type: %s", g.config.CompressionType)
	}

	// use urandom to get a unique message id
	gelfMessageID := make([]byte, 8)
	n, err := io.ReadFull(rand.Reader, gelfMessageID)
	if err != nil || n != 8 {
		return fmt.Errorf("rand.Reader: %d/%s", n, err)
	}

	if zw != nil {
		if err != nil {
			return err
		}
		if _, err = zw.Write(messageBytes); err != nil {
			zw.Close()
			return err
		}
		zw.Close()
		gelfBytes = compressedBuf.Bytes()
	}

	if g.numChunks(gelfBytes) > 1 {
		return g.writeChunked(ctx, gelfBytes)
	}

	return g.send(gelfBytes)
}

func (g *gelfexporter) pushLogsData(ctx context.Context, logs plog.Logs) error {

	var err error

	// sender, err := connect(ctx, g.logger, g.config)

	if err != nil {
		return err
	}

	rls := logs.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rl := rls.At(i)

		ills := rl.ScopeLogs()
		for j := 0; j < ills.Len(); j++ {
			ils := ills.At(j)
			logs := ils.LogRecords()
			for k := 0; k < logs.Len(); k++ {
				lr := logs.At(k)
				body := lr.Body()
				var outputTimestamp float64
				var outputSeverity int32
				timestamp := lr.Timestamp()
				severity := int32(lr.SeverityNumber())
				severityText := lr.SeverityText()

				// g.logger.Debug("Timestamp: %s, Severity: %s, Body: %s, SeverityText: %s", timestamp, severity, body.AsString(), severityText)

				// If timestamp is not set, set it to the current time
				if timestamp == 0 {
					outputTimestamp = float64(time.Now().UnixNano())
				} else {
					outputTimestamp = float64(timestamp)
				}

				if severityText != "" {
					outputSeverity = logLevelToSeverity(severityText)
				} else if severity != 0 {
					outputSeverity = severity
				} else {
					outputSeverity = logLevelToSeverity("INFO")
				}

				// Convert logs to GELF format
				err := g.convertLogsToGELF(ctx, outputTimestamp, outputSeverity, body.AsString())
				if err != nil {
					return err
				}

			}
		}
	}

	return err

}
