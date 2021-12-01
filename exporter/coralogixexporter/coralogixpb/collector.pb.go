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

package coralogixpb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	model "github.com/jaegertracing/jaeger/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PostSpansRequest struct {
	Batch                model.Batch `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch"`
	Metadata             *Metadata   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PostSpansRequest) Reset()         { *m = PostSpansRequest{} }
func (m *PostSpansRequest) String() string { return proto.CompactTextString(m) }
func (*PostSpansRequest) ProtoMessage()    {}
func (*PostSpansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9305884a292fdf82, []int{0}
}
func (m *PostSpansRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostSpansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostSpansRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostSpansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostSpansRequest.Merge(m, src)
}
func (m *PostSpansRequest) XXX_Size() int {
	return m.Size()
}
func (m *PostSpansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostSpansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostSpansRequest proto.InternalMessageInfo

func (m *PostSpansRequest) GetBatch() model.Batch {
	if m != nil {
		return m.Batch
	}
	return model.Batch{}
}

func (m *PostSpansRequest) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	SubsystemName        string   `protobuf:"bytes,2,opt,name=subsystemName,proto3" json:"subsystemName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9305884a292fdf82, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *Metadata) GetSubsystemName() string {
	if m != nil {
		return m.SubsystemName
	}
	return ""
}

type PostSpansResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostSpansResponse) Reset()         { *m = PostSpansResponse{} }
func (m *PostSpansResponse) String() string { return proto.CompactTextString(m) }
func (*PostSpansResponse) ProtoMessage()    {}
func (*PostSpansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9305884a292fdf82, []int{2}
}
func (m *PostSpansResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostSpansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostSpansResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostSpansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostSpansResponse.Merge(m, src)
}
func (m *PostSpansResponse) XXX_Size() int {
	return m.Size()
}
func (m *PostSpansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostSpansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostSpansResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PostSpansRequest)(nil), "jaeger.api_v2.PostSpansRequest")
	proto.RegisterType((*Metadata)(nil), "jaeger.api_v2.Metadata")
	proto.RegisterType((*PostSpansResponse)(nil), "jaeger.api_v2.PostSpansResponse")
}

func init() { proto.RegisterFile("collector.proto", fileDescriptor_9305884a292fdf82) }

var fileDescriptor_9305884a292fdf82 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x5b, 0xa9, 0xd0, 0x11, 0xd3, 0x26, 0x41, 0x91, 0x50, 0x59, 0x3a, 0x48, 0x87, 0xdd,
	0x58, 0xe9, 0xd2, 0xd1, 0xce, 0x45, 0xac, 0x37, 0x2f, 0x31, 0xae, 0x2f, 0xeb, 0xc4, 0xee, 0xbe,
	0xd3, 0xcc, 0x28, 0x09, 0x9d, 0xfa, 0x0a, 0x7d, 0x29, 0x8f, 0x41, 0xf7, 0x08, 0xe9, 0x83, 0xc4,
	0xce, 0xae, 0x92, 0x0b, 0xdd, 0x5e, 0x9e, 0xe7, 0xf7, 0xfe, 0x99, 0x67, 0x48, 0x3d, 0xc0, 0x28,
	0x82, 0x40, 0xa3, 0x74, 0x84, 0x44, 0x8d, 0xb4, 0xf6, 0xc4, 0x20, 0x04, 0xe9, 0x30, 0xc1, 0x1f,
	0x97, 0x5e, 0xa7, 0x1a, 0xe3, 0x0c, 0xa2, 0xcc, 0xeb, 0x34, 0x43, 0x0c, 0xd1, 0x94, 0x6e, 0x5a,
	0xe5, 0xea, 0x79, 0x88, 0x18, 0x46, 0xe0, 0x32, 0xc1, 0x5d, 0x96, 0x24, 0xa8, 0x99, 0xe6, 0x98,
	0xa8, 0xcc, 0xb5, 0x57, 0xa4, 0xf1, 0x80, 0x4a, 0x8f, 0x05, 0x4b, 0x94, 0x0f, 0xcf, 0x0b, 0x50,
	0x9a, 0x5e, 0x91, 0xa3, 0x29, 0xd3, 0xc1, 0xbc, 0x6d, 0xf5, 0xad, 0x41, 0xd5, 0x6b, 0x3a, 0x7b,
	0x3b, 0x9d, 0x51, 0xea, 0x8d, 0x0e, 0xd7, 0x5f, 0xbd, 0x03, 0x3f, 0x03, 0xe9, 0x90, 0x94, 0x63,
	0xd0, 0x6c, 0xc6, 0x34, 0x6b, 0x97, 0x4c, 0x53, 0xab, 0xd0, 0x74, 0x97, 0xdb, 0xfe, 0x0e, 0xb4,
	0x27, 0xa4, 0xbc, 0x55, 0xe9, 0x80, 0xd4, 0x99, 0x10, 0x11, 0x0f, 0xcc, 0x71, 0xf7, 0x2c, 0x06,
	0xb3, 0xbc, 0xe2, 0x17, 0x65, 0x7a, 0x41, 0x6a, 0x6a, 0x31, 0x55, 0x2b, 0xa5, 0x21, 0x36, 0x5c,
	0xc9, 0x70, 0xfb, 0xa2, 0x7d, 0x46, 0x4e, 0xff, 0x3c, 0x4b, 0x09, 0x4c, 0x14, 0x78, 0xaf, 0xa4,
	0x71, 0xbb, 0x8d, 0x73, 0x0c, 0x72, 0xc9, 0x03, 0xa0, 0x73, 0x52, 0xd9, 0x81, 0xb4, 0x57, 0x38,
	0xba, 0x98, 0x4c, 0xa7, 0xff, 0x3f, 0x90, 0xed, 0xb0, 0xdb, 0x6f, 0x9f, 0x3f, 0xef, 0x25, 0x6a,
	0xd7, 0x4c, 0xde, 0x4b, 0xcf, 0x55, 0xa9, 0x7d, 0x63, 0x5d, 0x8e, 0xae, 0xd7, 0x9b, 0xae, 0xf5,
	0xb1, 0xe9, 0x5a, 0xdf, 0x9b, 0xae, 0x45, 0x5a, 0x1c, 0xf3, 0x59, 0x5a, 0xb2, 0x80, 0x27, 0x61,
	0x3e, 0x72, 0x72, 0x92, 0xa9, 0xf0, 0x22, 0x50, 0x6a, 0x90, 0xd3, 0x63, 0xf3, 0x4f, 0xc3, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x87, 0x99, 0x56, 0x0a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CollectorServiceClient is the client API for CollectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorServiceClient interface {
	PostSpans(ctx context.Context, in *PostSpansRequest, opts ...grpc.CallOption) (*PostSpansResponse, error)
}

type collectorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCollectorServiceClient(cc *grpc.ClientConn) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) PostSpans(ctx context.Context, in *PostSpansRequest, opts ...grpc.CallOption) (*PostSpansResponse, error) {
	out := new(PostSpansResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.CollectorService/PostSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServiceServer is the server API for CollectorService service.
type CollectorServiceServer interface {
	PostSpans(context.Context, *PostSpansRequest) (*PostSpansResponse, error)
}

// UnimplementedCollectorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCollectorServiceServer struct {
}

func (*UnimplementedCollectorServiceServer) PostSpans(ctx context.Context, req *PostSpansRequest) (*PostSpansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSpans not implemented")
}

func RegisterCollectorServiceServer(s *grpc.Server, srv CollectorServiceServer) {
	s.RegisterService(&_CollectorService_serviceDesc, srv)
}

func _CollectorService_PostSpans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostSpansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).PostSpans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.CollectorService/PostSpans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).PostSpans(ctx, req.(*PostSpansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.api_v2.CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSpans",
			Handler:    _CollectorService_PostSpans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}

func (m *PostSpansRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostSpansRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostSpansRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCollector(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Batch.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCollector(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SubsystemName) > 0 {
		i -= len(m.SubsystemName)
		copy(dAtA[i:], m.SubsystemName)
		i = encodeVarintCollector(dAtA, i, uint64(len(m.SubsystemName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ApplicationName) > 0 {
		i -= len(m.ApplicationName)
		copy(dAtA[i:], m.ApplicationName)
		i = encodeVarintCollector(dAtA, i, uint64(len(m.ApplicationName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostSpansResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostSpansResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostSpansResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollector(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollector(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PostSpansRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Batch.Size()
	n += 1 + l + sovCollector(uint64(l))
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCollector(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ApplicationName)
	if l > 0 {
		n += 1 + l + sovCollector(uint64(l))
	}
	l = len(m.SubsystemName)
	if l > 0 {
		n += 1 + l + sovCollector(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostSpansResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCollector(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollector(x uint64) (n int) {
	return sovCollector(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PostSpansRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
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
			return fmt.Errorf("proto: PostSpansRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostSpansRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
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
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollector
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Batch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
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
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollector
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollector
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
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
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollector
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubsystemName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
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
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollector
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubsystemName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollector
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PostSpansResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
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
			return fmt.Errorf("proto: PostSpansResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostSpansResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollector
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCollector(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollector
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
					return 0, ErrIntOverflowCollector
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
					return 0, ErrIntOverflowCollector
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
				return 0, ErrInvalidLengthCollector
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollector
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollector
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollector        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollector          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollector = fmt.Errorf("proto: unexpected end of group")
)
