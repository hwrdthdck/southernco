// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package zipkincore

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	thrift "github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

const CLIENT_SEND = "cs"
const CLIENT_RECV = "cr"
const SERVER_SEND = "ss"
const SERVER_RECV = "sr"
const MESSAGE_SEND = "ms"
const MESSAGE_RECV = "mr"
const WIRE_SEND = "ws"
const WIRE_RECV = "wr"
const CLIENT_SEND_FRAGMENT = "csf"
const CLIENT_RECV_FRAGMENT = "crf"
const SERVER_SEND_FRAGMENT = "ssf"
const SERVER_RECV_FRAGMENT = "srf"
const LOCAL_COMPONENT = "lc"
const CLIENT_ADDR = "ca"
const SERVER_ADDR = "sa"
const MESSAGE_ADDR = "ma"

func init() {
}
