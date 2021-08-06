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

package splunkhecexporter

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"
)

// client sends the data to the splunk backend.
type client struct {
	config  *Config
	url     *url.URL
	client  *http.Client
	logger  *zap.Logger
	zippers sync.Pool
	wg      sync.WaitGroup
	headers map[string]string
}

// Minimum number of bytes to compress. 1500 is the MTU of an ethernet frame.
const minCompressionLen = 1500

func (c *client) pushMetricsData(
	ctx context.Context,
	md pdata.Metrics,
) error {
	c.wg.Add(1)
	defer c.wg.Done()

	splunkDataPoints, _ := metricDataToSplunk(c.logger, md, c.config)
	if len(splunkDataPoints) == 0 {
		return nil
	}

	body, compressed, err := encodeBodyEvents(&c.zippers, splunkDataPoints, c.config.DisableCompression)
	if err != nil {
		return consumererror.Permanent(err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.url.String(), body)
	if err != nil {
		return consumererror.Permanent(err)
	}

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	if compressed {
		req.Header.Set("Content-Encoding", "gzip")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	return splunk.HandleHTTPCode(resp)
}

func (c *client) pushTraceData(
	ctx context.Context,
	td pdata.Traces,
) error {
	c.wg.Add(1)
	defer c.wg.Done()

	splunkEvents, _ := traceDataToSplunk(c.logger, td, c.config)
	if len(splunkEvents) == 0 {
		return nil
	}

	return c.sendSplunkEvents(ctx, splunkEvents)
}

func (c *client) sendSplunkEvents(ctx context.Context, splunkEvents []*splunk.Event) error {
	body, compressed, err := encodeBodyEvents(&c.zippers, splunkEvents, c.config.DisableCompression)
	if err != nil {
		return consumererror.Permanent(err)
	}
	return c.postEvents(ctx, body, nil, compressed)
}

func (c *client) pushLogData(ctx context.Context, ld pdata.Logs) error {
	c.wg.Add(1)
	defer c.wg.Done()

	gzipWriter := c.zippers.Get().(*gzip.Writer)
	defer c.zippers.Put(gzipWriter)

	gzipBuffer := bytes.NewBuffer(make([]byte, 0, c.config.MaxContentLengthLogs))
	gzipWriter.Reset(gzipBuffer)

	// Callback when each batch is to be sent.
	send := func(ctx context.Context, buf *bytes.Buffer, headers map[string]string) (err error) {
		shouldCompress := buf.Len() >= minCompressionLen && !c.config.DisableCompression

		if shouldCompress {
			gzipBuffer.Reset()
			gzipWriter.Reset(gzipBuffer)

			if _, err = io.Copy(gzipWriter, buf); err != nil {
				return fmt.Errorf("failed copying buffer to gzip writer: %v", err)
			}

			if err = gzipWriter.Close(); err != nil {
				return fmt.Errorf("failed flushing compressed data to gzip writer: %v", err)
			}

			return c.postEvents(ctx, gzipBuffer, headers, shouldCompress)
		}

		return c.postEvents(ctx, buf, headers, shouldCompress)
	}

	return c.pushLogDataInBatches(ctx, ld, send)
}

// A guesstimated value > length of bytes of a single event.
// Added to buffer capacity so that buffer is likely to grow by reslicing when buf.Len() > bufCap.
const bufCapPadding = uint(4096)
const libraryHeaderName = "X-Splunk-Instrumentation-Library"
const profilingLibraryName = "otel.profiling"

var profilingHeaders = map[string]string{
	libraryHeaderName: profilingLibraryName,
}

func isProfilingData(ill pdata.InstrumentationLibraryLogs) bool {
	return ill.InstrumentationLibrary().Name() == profilingLibraryName
}

// pushLogDataInBatches sends batches of Splunk events in JSON format.
// The batch content length is restricted to MaxContentLengthLogs.
// ld log records are parsed to Splunk events.
// TODO: gofmt?
func (c *client) pushLogDataInBatches(ctx context.Context, ld pdata.Logs, send func(context.Context, *bytes.Buffer, map[string]string) error) error {
	// Buffer capacity.
	var bufCap = c.config.MaxContentLengthLogs

	// Buffer of JSON encoded Splunk events.
	// Expected to grow more than bufCap then truncated to bufLen.
	var buf = bytes.NewBuffer(make([]byte, 0, bufCap+bufCapPadding))
	var encoder = json.NewEncoder(buf)
	var tmpBuf = bytes.NewBuffer(make([]byte, 0, bufCapPadding))

	var profilingBuf = bytes.NewBuffer(make([]byte, 0, bufCap+bufCapPadding))
	var profilingEncoder = json.NewEncoder(profilingBuf)
	var profilingTmpBuf = bytes.NewBuffer(make([]byte, 0, bufCapPadding))

	// Length of retained bytes in buffer after truncation.
	var bufLen int
	var profilingBufLen int

	// Index of the log record of the first unsent event in buffer.
	var bufFront *logIndex
	var profilingBufFront *logIndex

	var permanentErrors []error

	var rls = ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		res := rls.At(i).Resource()
		ills := rls.At(i).InstrumentationLibraryLogs()
		for j := 0; j < ills.Len(); j++ {
			logs := ills.At(j).Logs()

			var err error

			//c.logger.Info("ILLS: ", zap.Int("j", j), zap.String("libraryName", ills.At(j).InstrumentationLibrary().Name()))

			if isProfilingData(ills.At(j)) {
				err, profilingBufLen, profilingBufFront, permanentErrors = c.pushLogRecords(ctx, logs, res, i, j, profilingBuf, profilingEncoder, profilingTmpBuf, bufCap, profilingBufLen, profilingBufFront, permanentErrors, profilingHeaders, send)
			} else {
				err, bufLen, bufFront, permanentErrors = c.pushLogRecords(ctx, logs, res, i, j, buf, encoder, tmpBuf, bufCap, bufLen, bufFront, permanentErrors, nil, send)
			}

			if err != nil {
				//c.logger.Error("Error in the loop!", zap.Any("err", err))
				return consumererror.NewLogs(err, *subLogs(&ld, bufFront, profilingBufFront))
			}
		}
	}

	if buf.Len() > 0 {
		if err := send(ctx, buf, nil); err != nil {
			//c.logger.Info("Error sending buf from the end!", zap.Any("err", err))
			return consumererror.NewLogs(err, *subLogs(&ld, bufFront, profilingBufFront))
		}
		//c.logger.Info("Sent buf from the end with no error!")
	}

	if profilingBuf.Len() > 0 {
		if err := send(ctx, profilingBuf, profilingHeaders); err != nil {
			//c.logger.Info("Error sending profilingBuf from the end!", zap.Any("err", err))
			return consumererror.NewLogs(err, *subLogs(&ld, nil, profilingBufFront))
		}
		//c.logger.Info("Sent profilingBuf from the end with no error!")
	}

	return consumererror.Combine(permanentErrors)
}

func (c *client) pushLogRecords(ctx context.Context, logs pdata.LogSlice, res pdata.Resource, i int, j int, buf *bytes.Buffer, encoder *json.Encoder, tmpBuf *bytes.Buffer, bufCap uint, bufLen int, bufFront *logIndex, permanentErrors []error, headers map[string]string, send func(context.Context, *bytes.Buffer, map[string]string) error) (error, int, *logIndex, []error) {
	for k := 0; k < logs.Len(); k++ {
		if bufFront == nil {
			bufFront = &logIndex{resource: i, library: j, record: k}
			//c.logger.Info("Updating in loop!", zap.Int("bufFront.resource", bufFront.resource), zap.Int("bufFront.library", bufFront.library), zap.Int("bufFront.record", bufFront.record))
		}

		// Parsing log record to Splunk event.
		event := mapLogRecordToSplunkEvent(res, logs.At(k), c.config, c.logger)
		// JSON encoding event and writing to buffer.
		//c.logger.Info("Before encoding", zap.Int("buf_len", buf.Len()))
		if err := encoder.Encode(event); err != nil {
			permanentErrors = append(permanentErrors, consumererror.Permanent(fmt.Errorf("dropped log event: %v, error: %v", event, err)))
			continue
		}

		//c.logger.Info("After encoding:", zap.Int("buf_len", buf.Len()))

		// Continue adding events to buffer up to capacity.
		// 0 capacity is interpreted as unknown/unbound consistent with ContentLength in http.Request.
		if buf.Len() <= int(bufCap) || bufCap == 0 {
			// Tracking length of event bytes below capacity in buffer.
			bufLen = buf.Len()
			//c.logger.Info("Continuing since buffer is not full")
			continue
		}

		//c.logger.Info("Time to flush buffer!", zap.Uint("bufCap", bufCap))

		tmpBuf.Reset()
		// Storing event bytes over capacity in buffer before truncating.
		if bufCap > 0 {
			if over := buf.Len() - bufLen; over <= int(bufCap) {
				tmpBuf.Write(buf.Bytes()[bufLen:buf.Len()])
				//c.logger.Info("Copying to tmpBuf",
				//	zap.Int("bufLen", bufLen), zap.Int("buf.Len()", buf.Len()))
			} else {
				permanentErrors = append(permanentErrors, consumererror.Permanent(
					fmt.Errorf("dropped log event: %s, error: event size %d bytes larger than configured max content length %d bytes", string(buf.Bytes()[bufLen:buf.Len()]), over, bufCap)))
			}
		}

		// Truncating buffer at tracked length below capacity and sending.
		buf.Truncate(bufLen)
		if buf.Len() > 0 {
			if err := send(ctx, buf, headers); err != nil {
				//c.logger.Info("Error flushing buffer from the loop!", zap.Any("err", err))
				return err, bufLen, bufFront, permanentErrors
			}
			//c.logger.Info("Flushed buffer from the loop with no error!")
		}
		buf.Reset()

		// Writing truncated bytes back to buffer.
		tmpBuf.WriteTo(buf)

		if buf.Len() > 0 {
			// This means that the current record has overflown the buffer and was not sent
			bufFront = &logIndex{resource: i, library: j, record: k}
			//c.logger.Info("Updating in loop!", zap.Int("bufFront.resource", bufFront.resource), zap.Int("bufFront.library", bufFront.library), zap.Int("bufFront.record", bufFront.record))
		} else {
			// This means that the entire buffer was sent, including the current record
			bufFront = nil
			//c.logger.Info("Reset at iteration end!", zap.Any("bufFront", bufFront))
		}

		bufLen = buf.Len()
	}

	return nil, bufLen, bufFront, permanentErrors
}

func (c *client) postEvents(ctx context.Context, events io.Reader, headers map[string]string, compressed bool) error {
	req, err := http.NewRequestWithContext(ctx, "POST", c.url.String(), events)
	if err != nil {
		return consumererror.Permanent(err)
	}

	// Set the headers configured for the client
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	// Set extra headers passed by the caller
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if compressed {
		req.Header.Set("Content-Encoding", "gzip")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = splunk.HandleHTTPCode(resp)

	io.Copy(ioutil.Discard, resp.Body)

	return err
}

// subLogs returns a subset of `ld` starting from `profilingBufFront` for profiling data
// plus starting from `bufFront` for non-profiling data. Both can be nil, in which case they are ignored
func subLogs(ld *pdata.Logs, bufFront *logIndex, profilingBufFront *logIndex) *pdata.Logs {
	if ld == nil {
		return ld
	}

	subset := pdata.NewLogs()
	subLogsByType(ld, bufFront, &subset, false)
	subLogsByType(ld, profilingBufFront, &subset, true)

	return &subset
}

func subLogsByType(src *pdata.Logs, from *logIndex, dst *pdata.Logs, profiling bool) {
	if from == nil {
		return // All the data of this type was sent successfully
	}

	resources := src.ResourceLogs()
	resourcesSub := dst.ResourceLogs()

	for i := from.resource; i < resources.Len(); i++ {
		resourcesSub.AppendEmpty()
		resources.At(i).Resource().CopyTo(resourcesSub.At(i - from.resource).Resource())

		libraries := resources.At(i).InstrumentationLibraryLogs()
		librariesSub := resourcesSub.At(i - from.resource).InstrumentationLibraryLogs()

		j := 0
		if i == from.resource {
			j = from.library
		}
		for jSub := 0; j < libraries.Len(); j++ {
			lib := libraries.At(j)

			// Only copy profiling data if requested. If not requested, only copy non-profiling data
			if profiling != isProfilingData(lib) {
				continue
			}

			librariesSub.AppendEmpty()
			lib.InstrumentationLibrary().CopyTo(librariesSub.At(jSub).InstrumentationLibrary())

			logs := lib.Logs()
			logsSub := librariesSub.At(jSub).Logs()
			jSub++

			k := 0
			if i == from.resource && j == from.library {
				k = from.record
			}

			for kSub := 0; k < logs.Len(); k++ { //revive:disable-line:var-naming
				logsSub.AppendEmpty()
				logs.At(k).CopyTo(logsSub.At(kSub))
				kSub++
			}
		}
	}
}

func encodeBodyEvents(zippers *sync.Pool, evs []*splunk.Event, disableCompression bool) (bodyReader io.Reader, compressed bool, err error) {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	for _, e := range evs {
		err := encoder.Encode(e)
		if err != nil {
			return nil, false, err
		}
	}
	return getReader(zippers, buf, disableCompression)
}

// avoid attempting to compress things that fit into a single ethernet frame
func getReader(zippers *sync.Pool, b *bytes.Buffer, disableCompression bool) (io.Reader, bool, error) {
	var err error
	if !disableCompression && b.Len() > minCompressionLen {
		buf := new(bytes.Buffer)
		w := zippers.Get().(*gzip.Writer)
		defer zippers.Put(w)
		w.Reset(buf)
		_, err = w.Write(b.Bytes())
		if err == nil {
			err = w.Close()
			if err == nil {
				return buf, true, nil
			}
		}
	}
	return b, false, err
}

func (c *client) stop(context.Context) error {
	c.wg.Wait()
	return nil
}

func (c *client) start(context.Context, component.Host) (err error) {
	return nil
}
