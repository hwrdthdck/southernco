// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/logs/v1/logs.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v11 "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/common/v1"
	v1 "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/resource/v1"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	SeverityNumber_UNDEFINED_SEVERITY_NUMBER SeverityNumber = 0
	SeverityNumber_TRACE                     SeverityNumber = 1
	SeverityNumber_TRACE2                    SeverityNumber = 2
	SeverityNumber_TRACE3                    SeverityNumber = 3
	SeverityNumber_TRACE4                    SeverityNumber = 4
	SeverityNumber_DEBUG                     SeverityNumber = 5
	SeverityNumber_DEBUG2                    SeverityNumber = 6
	SeverityNumber_DEBUG3                    SeverityNumber = 7
	SeverityNumber_DEBUG4                    SeverityNumber = 8
	SeverityNumber_INFO                      SeverityNumber = 9
	SeverityNumber_INFO2                     SeverityNumber = 10
	SeverityNumber_INFO3                     SeverityNumber = 11
	SeverityNumber_INFO4                     SeverityNumber = 12
	SeverityNumber_WARN                      SeverityNumber = 13
	SeverityNumber_WARN2                     SeverityNumber = 14
	SeverityNumber_WARN3                     SeverityNumber = 15
	SeverityNumber_WARN4                     SeverityNumber = 16
	SeverityNumber_ERROR                     SeverityNumber = 17
	SeverityNumber_ERROR2                    SeverityNumber = 18
	SeverityNumber_ERROR3                    SeverityNumber = 19
	SeverityNumber_ERROR4                    SeverityNumber = 20
	SeverityNumber_FATAL                     SeverityNumber = 21
	SeverityNumber_FATAL2                    SeverityNumber = 22
	SeverityNumber_FATAL3                    SeverityNumber = 23
	SeverityNumber_FATAL4                    SeverityNumber = 24
)

var SeverityNumber_name = map[int32]string{
	0:  "UNDEFINED_SEVERITY_NUMBER",
	1:  "TRACE",
	2:  "TRACE2",
	3:  "TRACE3",
	4:  "TRACE4",
	5:  "DEBUG",
	6:  "DEBUG2",
	7:  "DEBUG3",
	8:  "DEBUG4",
	9:  "INFO",
	10: "INFO2",
	11: "INFO3",
	12: "INFO4",
	13: "WARN",
	14: "WARN2",
	15: "WARN3",
	16: "WARN4",
	17: "ERROR",
	18: "ERROR2",
	19: "ERROR3",
	20: "ERROR4",
	21: "FATAL",
	22: "FATAL2",
	23: "FATAL3",
	24: "FATAL4",
}

var SeverityNumber_value = map[string]int32{
	"UNDEFINED_SEVERITY_NUMBER": 0,
	"TRACE":                     1,
	"TRACE2":                    2,
	"TRACE3":                    3,
	"TRACE4":                    4,
	"DEBUG":                     5,
	"DEBUG2":                    6,
	"DEBUG3":                    7,
	"DEBUG4":                    8,
	"INFO":                      9,
	"INFO2":                     10,
	"INFO3":                     11,
	"INFO4":                     12,
	"WARN":                      13,
	"WARN2":                     14,
	"WARN3":                     15,
	"WARN4":                     16,
	"ERROR":                     17,
	"ERROR2":                    18,
	"ERROR3":                    19,
	"ERROR4":                    20,
	"FATAL":                     21,
	"FATAL2":                    22,
	"FATAL3":                    23,
	"FATAL4":                    24,
}

func (x SeverityNumber) String() string {
	return proto.EnumName(SeverityNumber_name, int32(x))
}

func (SeverityNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}

// Masks for LogRecord.flags field.
type LogRecordFlags int32

const (
	LogRecordFlags_UNDEFINED_LOG_RECORD_FLAG LogRecordFlags = 0
	LogRecordFlags_TRACE_FLAGS_MASK          LogRecordFlags = 255
)

var LogRecordFlags_name = map[int32]string{
	0:   "UNDEFINED_LOG_RECORD_FLAG",
	255: "TRACE_FLAGS_MASK",
}

var LogRecordFlags_value = map[string]int32{
	"UNDEFINED_LOG_RECORD_FLAG": 0,
	"TRACE_FLAGS_MASK":          255,
}

func (x LogRecordFlags) String() string {
	return proto.EnumName(LogRecordFlags_name, int32(x))
}

func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}

// A collection of InstrumentationLibraryLogs from a Resource.
type ResourceLogs struct {
	// The resource for the logs in this message.
	// If this field is not set then no resource info is known.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of InstrumentationLibraryLogs that originate from a resource.
	InstrumentationLibraryLogs []*InstrumentationLibraryLogs `protobuf:"bytes,2,rep,name=instrumentation_library_logs,json=instrumentationLibraryLogs,proto3" json:"instrumentation_library_logs,omitempty"`
}

func (m *ResourceLogs) Reset()         { *m = ResourceLogs{} }
func (m *ResourceLogs) String() string { return proto.CompactTextString(m) }
func (*ResourceLogs) ProtoMessage()    {}
func (*ResourceLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}
func (m *ResourceLogs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceLogs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLogs.Merge(m, src)
}
func (m *ResourceLogs) XXX_Size() int {
	return m.Size()
}
func (m *ResourceLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLogs.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLogs proto.InternalMessageInfo

func (m *ResourceLogs) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceLogs) GetInstrumentationLibraryLogs() []*InstrumentationLibraryLogs {
	if m != nil {
		return m.InstrumentationLibraryLogs
	}
	return nil
}

// A collection of Logs produced by an InstrumentationLibrary.
type InstrumentationLibraryLogs struct {
	// The instrumentation library information for the logs in this message.
	// If this field is not set then no library info is known.
	InstrumentationLibrary *v11.InstrumentationLibrary `protobuf:"bytes,1,opt,name=instrumentation_library,json=instrumentationLibrary,proto3" json:"instrumentation_library,omitempty"`
	// A list of log records.
	Logs []*LogRecord `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (m *InstrumentationLibraryLogs) Reset()         { *m = InstrumentationLibraryLogs{} }
func (m *InstrumentationLibraryLogs) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibraryLogs) ProtoMessage()    {}
func (*InstrumentationLibraryLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}
func (m *InstrumentationLibraryLogs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstrumentationLibraryLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstrumentationLibraryLogs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstrumentationLibraryLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibraryLogs.Merge(m, src)
}
func (m *InstrumentationLibraryLogs) XXX_Size() int {
	return m.Size()
}
func (m *InstrumentationLibraryLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibraryLogs.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibraryLogs proto.InternalMessageInfo

func (m *InstrumentationLibraryLogs) GetInstrumentationLibrary() *v11.InstrumentationLibrary {
	if m != nil {
		return m.InstrumentationLibrary
	}
	return nil
}

func (m *InstrumentationLibraryLogs) GetLogs() []*LogRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/master/text/logs/0097-log-data-model.md
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3,enum=opentelemetry.proto.logs.v1.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// Short event identifier that does not contain varying parts. Name describes
	// what happened (e.g. "ProcessStarted"). Recommended to be no longer than 50
	// characters. Not guaranteed to be unique in any way. [Optional].
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *v11.AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	Attributes             []*v11.KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32          `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid. Can be set for logs that are part of request processing
	// and have an assigned trace id. [Optional].
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid. Can be set for logs that are part of a particular processing span.
	// If span_id is present trace_id SHOULD be also present. [Optional].
	SpanId []byte `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{2}
}
func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return m.Size()
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetTimeUnixNano() uint64 {
	if m != nil {
		return m.TimeUnixNano
	}
	return 0
}

func (m *LogRecord) GetSeverityNumber() SeverityNumber {
	if m != nil {
		return m.SeverityNumber
	}
	return SeverityNumber_UNDEFINED_SEVERITY_NUMBER
}

func (m *LogRecord) GetSeverityText() string {
	if m != nil {
		return m.SeverityText
	}
	return ""
}

func (m *LogRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogRecord) GetBody() *v11.AnyValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LogRecord) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LogRecord) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *LogRecord) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LogRecord) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *LogRecord) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.logs.v1.SeverityNumber", SeverityNumber_name, SeverityNumber_value)
	proto.RegisterEnum("opentelemetry.proto.logs.v1.LogRecordFlags", LogRecordFlags_name, LogRecordFlags_value)
	proto.RegisterType((*ResourceLogs)(nil), "opentelemetry.proto.logs.v1.ResourceLogs")
	proto.RegisterType((*InstrumentationLibraryLogs)(nil), "opentelemetry.proto.logs.v1.InstrumentationLibraryLogs")
	proto.RegisterType((*LogRecord)(nil), "opentelemetry.proto.logs.v1.LogRecord")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/logs/v1/logs.proto", fileDescriptor_d1c030a3ec7e961e)
}

var fileDescriptor_d1c030a3ec7e961e = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0x4f, 0x8f, 0x22, 0x45,
	0x18, 0xc6, 0x29, 0xfe, 0x53, 0xc3, 0xb0, 0x65, 0xb9, 0x3b, 0xd3, 0x8b, 0x4a, 0xc8, 0x68, 0x56,
	0x1c, 0xb3, 0x90, 0x69, 0x30, 0x1a, 0x3d, 0x35, 0x43, 0x43, 0xc8, 0xb0, 0xb0, 0x29, 0x60, 0x8d,
	0x5e, 0x3a, 0x0d, 0x5d, 0x92, 0x4e, 0x9a, 0x2a, 0xd2, 0x5d, 0x10, 0xf8, 0x16, 0xc6, 0x6f, 0xe4,
	0x6d, 0xe3, 0x69, 0x8e, 0x1e, 0xcd, 0xcc, 0x07, 0xd1, 0x54, 0x41, 0xb7, 0xcb, 0x06, 0x98, 0x13,
	0xbf, 0x7a, 0xde, 0xe7, 0x79, 0xfb, 0xad, 0x97, 0x34, 0xc0, 0x57, 0x7c, 0x41, 0x99, 0xa0, 0x1e,
	0x9d, 0x53, 0xe1, 0x6f, 0x6a, 0x0b, 0x9f, 0x0b, 0x5e, 0xf3, 0xf8, 0x2c, 0xa8, 0xad, 0x6e, 0xd4,
	0x67, 0x55, 0x49, 0xf8, 0xb3, 0x3d, 0xdf, 0x56, 0xac, 0xaa, 0xfa, 0xea, 0xa6, 0x78, 0x7d, 0xa8,
	0xc9, 0x94, 0xcf, 0xe7, 0x9c, 0xc9, 0x36, 0x5b, 0xda, 0x66, 0x8a, 0xd5, 0x43, 0x5e, 0x9f, 0x06,
	0x7c, 0xe9, 0x4f, 0xa9, 0x74, 0x87, 0xbc, 0xf5, 0x5f, 0xdd, 0x03, 0x98, 0x27, 0x3b, 0xa9, 0xc7,
	0x67, 0x01, 0x36, 0x61, 0x36, 0xb4, 0x68, 0xa0, 0x0c, 0x2a, 0x67, 0xfa, 0x37, 0xd5, 0x43, 0xc3,
	0x45, 0x7d, 0x56, 0x37, 0xd5, 0xb0, 0x01, 0x89, 0xa2, 0x78, 0x03, 0x3f, 0x77, 0x59, 0x20, 0xfc,
	0xe5, 0x9c, 0x32, 0x61, 0x0b, 0x97, 0x33, 0xcb, 0x73, 0x27, 0xbe, 0xed, 0x6f, 0x2c, 0x79, 0x2d,
	0x2d, 0x5e, 0x4e, 0x54, 0xce, 0xf4, 0xef, 0xab, 0x27, 0xee, 0x5d, 0xed, 0xee, 0x37, 0xe8, 0x6d,
	0xf3, 0x72, 0x4a, 0x52, 0x74, 0x8f, 0xd6, 0xae, 0xde, 0x03, 0x58, 0x3c, 0x1e, 0xc5, 0x0c, 0x5e,
	0x1e, 0x99, 0x6c, 0x77, 0xdf, 0xef, 0x0e, 0x0e, 0xb5, 0xdb, 0xf2, 0xd1, 0xb1, 0xc8, 0xc5, 0xe1,
	0x91, 0xf0, 0x8f, 0x30, 0xf9, 0xc1, 0x8d, 0x5f, 0x9d, 0xbc, 0x71, 0x8f, 0xcf, 0x08, 0x9d, 0x72,
	0xdf, 0x21, 0x2a, 0x73, 0xf5, 0x57, 0x02, 0xe6, 0x22, 0x0d, 0x7f, 0x05, 0x0b, 0xc2, 0x9d, 0x53,
	0x6b, 0xc9, 0xdc, 0xb5, 0xc5, 0x6c, 0xc6, 0xd5, 0xc0, 0x69, 0x92, 0x97, 0xea, 0x98, 0xb9, 0xeb,
	0xbe, 0xcd, 0x38, 0x1e, 0xc1, 0x67, 0x01, 0x5d, 0x51, 0xdf, 0x15, 0x1b, 0x8b, 0x2d, 0xe7, 0x13,
	0xea, 0x6b, 0xf1, 0x32, 0xa8, 0x14, 0xf4, 0x6f, 0x4f, 0x3e, 0x7a, 0xb8, 0xcb, 0xf4, 0x55, 0x84,
	0x14, 0x82, 0xbd, 0x33, 0xfe, 0x12, 0x9e, 0x47, 0x5d, 0x05, 0x5d, 0x0b, 0x2d, 0x51, 0x06, 0x95,
	0x1c, 0xc9, 0x87, 0xe2, 0x88, 0xae, 0x05, 0xc6, 0x30, 0xc9, 0xec, 0x39, 0xd5, 0x92, 0xaa, 0xa6,
	0x18, 0xff, 0x04, 0x93, 0x13, 0xee, 0x6c, 0xb4, 0x94, 0xda, 0xed, 0xd7, 0x4f, 0xec, 0xd6, 0x60,
	0x9b, 0x77, 0xb6, 0xb7, 0xa4, 0x44, 0x85, 0x70, 0x07, 0x42, 0x5b, 0x08, 0xdf, 0x9d, 0x2c, 0x05,
	0x0d, 0xb4, 0xb4, 0xda, 0xe0, 0x53, 0x2d, 0xee, 0xe8, 0xae, 0xc5, 0x07, 0x51, 0xfc, 0x03, 0xd4,
	0x1c, 0x9f, 0x2f, 0x16, 0xd4, 0xb1, 0xfe, 0x57, 0xad, 0x29, 0x5f, 0x32, 0xa1, 0x65, 0xca, 0xa0,
	0x72, 0x4e, 0x2e, 0x76, 0x75, 0x23, 0x2a, 0xdf, 0xca, 0x2a, 0x7e, 0x0e, 0x53, 0xbf, 0x79, 0xf6,
	0x2c, 0xd0, 0xb2, 0x65, 0x50, 0xc9, 0x90, 0xed, 0x01, 0xbf, 0x84, 0x59, 0xe1, 0xdb, 0x53, 0x6a,
	0xb9, 0x8e, 0x96, 0x2b, 0x83, 0x4a, 0x9e, 0x64, 0xd4, 0xb9, 0xeb, 0xe0, 0x4b, 0x98, 0x09, 0x16,
	0x36, 0x93, 0x15, 0xa8, 0x2a, 0x69, 0x79, 0xec, 0x3a, 0xd7, 0x7f, 0xc6, 0x61, 0x61, 0x7f, 0xcb,
	0xf8, 0x0b, 0xf8, 0x72, 0xdc, 0x6f, 0x99, 0xed, 0x6e, 0xdf, 0x6c, 0x59, 0x43, 0xf3, 0x9d, 0x49,
	0xba, 0xa3, 0x5f, 0xac, 0xfe, 0xf8, 0x4d, 0xd3, 0x24, 0x28, 0x86, 0x73, 0x30, 0x35, 0x22, 0xc6,
	0xad, 0x89, 0x00, 0x86, 0x30, 0xad, 0x50, 0x47, 0xf1, 0x88, 0xeb, 0x28, 0x11, 0x71, 0x03, 0x25,
	0xa5, 0xbd, 0x65, 0x36, 0xc7, 0x1d, 0x94, 0x92, 0xb2, 0x42, 0x1d, 0xa5, 0x23, 0xae, 0xa3, 0x4c,
	0xc4, 0x0d, 0x94, 0xc5, 0x59, 0x98, 0xec, 0xf6, 0xdb, 0x03, 0x94, 0x93, 0x41, 0x49, 0x3a, 0x82,
	0x21, 0xd6, 0xd1, 0x59, 0x88, 0x0d, 0x94, 0x97, 0xd6, 0x9f, 0x0d, 0xd2, 0x47, 0xe7, 0x52, 0x94,
	0xa4, 0xa3, 0x42, 0x88, 0x75, 0xf4, 0x2c, 0xc4, 0x06, 0x42, 0x12, 0x4d, 0x42, 0x06, 0x04, 0x7d,
	0x22, 0x1f, 0xa6, 0x50, 0x47, 0x38, 0xe2, 0x3a, 0xfa, 0x34, 0xe2, 0x06, 0x7a, 0x2e, 0xed, 0x6d,
	0x63, 0x64, 0xf4, 0xd0, 0x0b, 0x29, 0x2b, 0xd4, 0xd1, 0x45, 0xc4, 0x75, 0x74, 0x19, 0x71, 0x03,
	0x69, 0xd7, 0x6d, 0x58, 0x88, 0xde, 0x87, 0xb6, 0xfa, 0x26, 0xf6, 0x56, 0xd8, 0x1b, 0x74, 0x2c,
	0x62, 0xde, 0x0e, 0x48, 0xcb, 0x6a, 0xf7, 0x8c, 0x0e, 0x8a, 0xe1, 0x17, 0x10, 0xa9, 0xfd, 0xa8,
	0xf3, 0xd0, 0x7a, 0x63, 0x0c, 0xef, 0xd0, 0xbf, 0xa0, 0xf9, 0x07, 0x78, 0xff, 0x50, 0x02, 0xf7,
	0x0f, 0x25, 0xf0, 0xcf, 0x43, 0x09, 0xfc, 0xfe, 0x58, 0x8a, 0xdd, 0x3f, 0x96, 0x62, 0x7f, 0x3f,
	0x96, 0x62, 0xb0, 0xe4, 0xf2, 0x53, 0x2f, 0x4a, 0x53, 0xbe, 0x90, 0xc1, 0x5b, 0x29, 0xbd, 0x05,
	0xbf, 0xde, 0xcd, 0x3e, 0x36, 0xbb, 0xf2, 0xa7, 0xd9, 0xf3, 0xe8, 0x54, 0x70, 0xbf, 0xe6, 0x32,
	0x41, 0x7d, 0x66, 0x7b, 0x35, 0xc7, 0x16, 0x76, 0x6d, 0xcf, 0xf8, 0x5a, 0x75, 0x7d, 0x3d, 0xa3,
	0x2c, 0xfc, 0x3f, 0x98, 0xa4, 0x95, 0x54, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x66, 0x84,
	0x11, 0x35, 0x06, 0x00, 0x00,
}

func (m *ResourceLogs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceLogs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceLogs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InstrumentationLibraryLogs) > 0 {
		for iNdEx := len(m.InstrumentationLibraryLogs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InstrumentationLibraryLogs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLogs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLogs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InstrumentationLibraryLogs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstrumentationLibraryLogs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstrumentationLibraryLogs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLogs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.InstrumentationLibrary != nil {
		{
			size, err := m.InstrumentationLibrary.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLogs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpanId) > 0 {
		i -= len(m.SpanId)
		copy(dAtA[i:], m.SpanId)
		i = encodeVarintLogs(dAtA, i, uint64(len(m.SpanId)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.TraceId) > 0 {
		i -= len(m.TraceId)
		copy(dAtA[i:], m.TraceId)
		i = encodeVarintLogs(dAtA, i, uint64(len(m.TraceId)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Flags != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Flags))
		i--
		dAtA[i] = 0x45
	}
	if m.DroppedAttributesCount != 0 {
		i = encodeVarintLogs(dAtA, i, uint64(m.DroppedAttributesCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLogs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLogs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintLogs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SeverityText) > 0 {
		i -= len(m.SeverityText)
		copy(dAtA[i:], m.SeverityText)
		i = encodeVarintLogs(dAtA, i, uint64(len(m.SeverityText)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SeverityNumber != 0 {
		i = encodeVarintLogs(dAtA, i, uint64(m.SeverityNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.TimeUnixNano != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.TimeUnixNano))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogs(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceLogs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovLogs(uint64(l))
	}
	if len(m.InstrumentationLibraryLogs) > 0 {
		for _, e := range m.InstrumentationLibraryLogs {
			l = e.Size()
			n += 1 + l + sovLogs(uint64(l))
		}
	}
	return n
}

func (m *InstrumentationLibraryLogs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InstrumentationLibrary != nil {
		l = m.InstrumentationLibrary.Size()
		n += 1 + l + sovLogs(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovLogs(uint64(l))
		}
	}
	return n
}

func (m *LogRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeUnixNano != 0 {
		n += 9
	}
	if m.SeverityNumber != 0 {
		n += 1 + sovLogs(uint64(m.SeverityNumber))
	}
	l = len(m.SeverityText)
	if l > 0 {
		n += 1 + l + sovLogs(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLogs(uint64(l))
	}
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovLogs(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovLogs(uint64(l))
		}
	}
	if m.DroppedAttributesCount != 0 {
		n += 1 + sovLogs(uint64(m.DroppedAttributesCount))
	}
	if m.Flags != 0 {
		n += 5
	}
	l = len(m.TraceId)
	if l > 0 {
		n += 1 + l + sovLogs(uint64(l))
	}
	l = len(m.SpanId)
	if l > 0 {
		n += 1 + l + sovLogs(uint64(l))
	}
	return n
}

func sovLogs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogs(x uint64) (n int) {
	return sovLogs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceLogs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceLogs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceLogs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &v1.Resource{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstrumentationLibraryLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstrumentationLibraryLogs = append(m.InstrumentationLibraryLogs, &InstrumentationLibraryLogs{})
			if err := m.InstrumentationLibraryLogs[len(m.InstrumentationLibraryLogs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstrumentationLibraryLogs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstrumentationLibraryLogs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstrumentationLibraryLogs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstrumentationLibrary", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InstrumentationLibrary == nil {
				m.InstrumentationLibrary = &v11.InstrumentationLibrary{}
			}
			if err := m.InstrumentationLibrary.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &LogRecord{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeUnixNano", wireType)
			}
			m.TimeUnixNano = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeUnixNano = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeverityNumber", wireType)
			}
			m.SeverityNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeverityNumber |= SeverityNumber(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeverityText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SeverityText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &v11.AnyValue{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, &v11.KeyValue{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DroppedAttributesCount", wireType)
			}
			m.DroppedAttributesCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DroppedAttributesCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Flags = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = append(m.TraceId[:0], dAtA[iNdEx:postIndex]...)
			if m.TraceId == nil {
				m.TraceId = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLogs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpanId = append(m.SpanId[:0], dAtA[iNdEx:postIndex]...)
			if m.SpanId == nil {
				m.SpanId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLogs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogs = fmt.Errorf("proto: unexpected end of group")
)
