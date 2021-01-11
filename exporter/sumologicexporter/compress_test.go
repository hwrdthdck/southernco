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

package sumologicexporter

import (
	"compress/flate"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockedEncrypter struct {
	writeError error
	closeError error
}

func (e mockedEncrypter) Reset(dst io.Writer) {
}

func (e mockedEncrypter) Write(p []byte) (n int, err error) {
	return 0, e.writeError
}

func (e mockedEncrypter) Close() error {
	return e.closeError
}

func getTestCompressor(w error, c error) compressor {
	return compressor{
		format: GZIPCompression,
		writer: mockedEncrypter{
			writeError: w,
			closeError: c,
		},
	}
}

type mockedReader struct{}

func (r mockedReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func TestCompressGzip(t *testing.T) {
	const message = "This is an example log"

	c, err := newCompressor(GZIPCompression)
	require.NoError(t, err)

	body := strings.NewReader(message)

	data, err := c.compress(body)
	require.NoError(t, err)

	assert.Equal(t, message, decodeGzip(t, data))
}

func TestCompressTwice(t *testing.T) {
	const (
		message       = "This is an example log"
		secondMessage = "This is an another example log"
	)

	c, err := newCompressor(GZIPCompression)
	require.NoError(t, err)

	body := strings.NewReader(message)

	data, err := c.compress(body)
	require.NoError(t, err)
	assert.Equal(t, message, decodeGzip(t, data))

	body = strings.NewReader(secondMessage)
	data, err = c.compress(body)
	require.NoError(t, err)
	assert.Equal(t, secondMessage, decodeGzip(t, data))
}

func decodeGzip(t *testing.T, data io.Reader) string {
	r, err := gzip.NewReader(data)
	require.NoError(t, err)

	var buf []byte
	buf, err = ioutil.ReadAll(r)
	require.NoError(t, err)

	return string(buf)
}

func TestCompressDeflate(t *testing.T) {
	const message = "This is an example log"

	c, err := newCompressor(DeflateCompression)
	require.NoError(t, err)

	body := strings.NewReader(message)

	data, err := c.compress(body)
	require.NoError(t, err)

	assert.Equal(t, message, decodeDeflate(t, data))
}

func decodeDeflate(t *testing.T, data io.Reader) string {
	r := flate.NewReader(data)

	var buf []byte
	buf, err := ioutil.ReadAll(r)
	require.NoError(t, err)

	return string(buf)
}

func TestCompressReadError(t *testing.T) {
	c := getTestCompressor(nil, nil)
	r := mockedReader{}
	_, err := c.compress(r)

	assert.EqualError(t, err, "read error")
}

func TestCompressWriteError(t *testing.T) {
	c := getTestCompressor(errors.New("write error"), nil)
	r := strings.NewReader("test string")
	_, err := c.compress(r)

	assert.EqualError(t, err, "write error")
}

func TestCompressCloseError(t *testing.T) {
	c := getTestCompressor(nil, errors.New("close error"))
	r := strings.NewReader("test string")
	_, err := c.compress(r)

	assert.EqualError(t, err, "close error")
}
