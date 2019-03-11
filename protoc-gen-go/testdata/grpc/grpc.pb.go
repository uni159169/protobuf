// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package testing

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SimpleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return xxx_File_grpc_grpc_proto_rawdesc_gzipped, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

type SimpleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return xxx_File_grpc_grpc_proto_rawdesc_gzipped, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

type StreamMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMsg) Reset()         { *m = StreamMsg{} }
func (m *StreamMsg) String() string { return proto.CompactTextString(m) }
func (*StreamMsg) ProtoMessage()    {}
func (*StreamMsg) Descriptor() ([]byte, []int) {
	return xxx_File_grpc_grpc_proto_rawdesc_gzipped, []int{2}
}

func (m *StreamMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg.Unmarshal(m, b)
}
func (m *StreamMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg.Marshal(b, m, deterministic)
}
func (m *StreamMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg.Merge(m, src)
}
func (m *StreamMsg) XXX_Size() int {
	return xxx_messageInfo_StreamMsg.Size(m)
}
func (m *StreamMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg proto.InternalMessageInfo

type StreamMsg2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMsg2) Reset()         { *m = StreamMsg2{} }
func (m *StreamMsg2) String() string { return proto.CompactTextString(m) }
func (*StreamMsg2) ProtoMessage()    {}
func (*StreamMsg2) Descriptor() ([]byte, []int) {
	return xxx_File_grpc_grpc_proto_rawdesc_gzipped, []int{3}
}

func (m *StreamMsg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg2.Unmarshal(m, b)
}
func (m *StreamMsg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg2.Marshal(b, m, deterministic)
}
func (m *StreamMsg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg2.Merge(m, src)
}
func (m *StreamMsg2) XXX_Size() int {
	return xxx_messageInfo_StreamMsg2.Size(m)
}
func (m *StreamMsg2) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg2.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg2 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("grpc/grpc.proto", xxx_File_grpc_grpc_proto_rawdesc_gzipped)
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
	proto.RegisterType((*StreamMsg)(nil), "grpc.testing.StreamMsg")
	proto.RegisterType((*StreamMsg2)(nil), "grpc.testing.StreamMsg2")
}

var xxx_File_grpc_grpc_proto_rawdesc = []byte{
	// 450 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x0f, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x22,
	0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x32, 0x32, 0x98, 0x02,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d,
	0x73, 0x67, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x08, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x3d, 0x0a, 0x04, 0x42, 0x69, 0x64,
	0x69, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4d, 0x73, 0x67, 0x32, 0x28, 0x01, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var xxx_File_grpc_grpc_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_grpc_grpc_proto_rawdesc)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// This RPC streams from the server only.
	Downstream(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Test_DownstreamClient, error)
	// This RPC streams from the client.
	Upstream(ctx context.Context, opts ...grpc.CallOption) (Test_UpstreamClient, error)
	// This one streams in both directions.
	Bidi(ctx context.Context, opts ...grpc.CallOption) (Test_BidiClient, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.Test/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Downstream(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Test_DownstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/grpc.testing.Test/Downstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testDownstreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Test_DownstreamClient interface {
	Recv() (*StreamMsg, error)
	grpc.ClientStream
}

type testDownstreamClient struct {
	grpc.ClientStream
}

func (x *testDownstreamClient) Recv() (*StreamMsg, error) {
	m := new(StreamMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) Upstream(ctx context.Context, opts ...grpc.CallOption) (Test_UpstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[1], "/grpc.testing.Test/Upstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testUpstreamClient{stream}
	return x, nil
}

type Test_UpstreamClient interface {
	Send(*StreamMsg) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testUpstreamClient struct {
	grpc.ClientStream
}

func (x *testUpstreamClient) Send(m *StreamMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testUpstreamClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) Bidi(ctx context.Context, opts ...grpc.CallOption) (Test_BidiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[2], "/grpc.testing.Test/Bidi", opts...)
	if err != nil {
		return nil, err
	}
	x := &testBidiClient{stream}
	return x, nil
}

type Test_BidiClient interface {
	Send(*StreamMsg) error
	Recv() (*StreamMsg2, error)
	grpc.ClientStream
}

type testBidiClient struct {
	grpc.ClientStream
}

func (x *testBidiClient) Send(m *StreamMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testBidiClient) Recv() (*StreamMsg2, error) {
	m := new(StreamMsg2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// This RPC streams from the server only.
	Downstream(*SimpleRequest, Test_DownstreamServer) error
	// This RPC streams from the client.
	Upstream(Test_UpstreamServer) error
	// This one streams in both directions.
	Bidi(Test_BidiServer) error
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.Test/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Downstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServer).Downstream(m, &testDownstreamServer{stream})
}

type Test_DownstreamServer interface {
	Send(*StreamMsg) error
	grpc.ServerStream
}

type testDownstreamServer struct {
	grpc.ServerStream
}

func (x *testDownstreamServer) Send(m *StreamMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _Test_Upstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Upstream(&testUpstreamServer{stream})
}

type Test_UpstreamServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamMsg, error)
	grpc.ServerStream
}

type testUpstreamServer struct {
	grpc.ServerStream
}

func (x *testUpstreamServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testUpstreamServer) Recv() (*StreamMsg, error) {
	m := new(StreamMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Test_Bidi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Bidi(&testBidiServer{stream})
}

type Test_BidiServer interface {
	Send(*StreamMsg2) error
	Recv() (*StreamMsg, error)
	grpc.ServerStream
}

type testBidiServer struct {
	grpc.ServerStream
}

func (x *testBidiServer) Send(m *StreamMsg2) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testBidiServer) Recv() (*StreamMsg, error) {
	m := new(StreamMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _Test_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Downstream",
			Handler:       _Test_Downstream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upstream",
			Handler:       _Test_Upstream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Bidi",
			Handler:       _Test_Bidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/grpc.proto",
}
