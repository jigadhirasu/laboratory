// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloTask struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            int32    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Response             string   `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloTask) Reset()         { *m = HelloTask{} }
func (m *HelloTask) String() string { return proto.CompactTextString(m) }
func (*HelloTask) ProtoMessage()    {}
func (*HelloTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloTask.Unmarshal(m, b)
}
func (m *HelloTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloTask.Marshal(b, m, deterministic)
}
func (m *HelloTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloTask.Merge(m, src)
}
func (m *HelloTask) XXX_Size() int {
	return xxx_messageInfo_HelloTask.Size(m)
}
func (m *HelloTask) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloTask.DiscardUnknown(m)
}

var xxx_messageInfo_HelloTask proto.InternalMessageInfo

func (m *HelloTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloTask) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloTask) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HelloTask) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloTask)(nil), "hello.HelloTask")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x8a, 0xb9, 0x38, 0x3d,
	0x40, 0x8c, 0x90, 0xc4, 0xe2, 0x6c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x3d, 0x55, 0x82, 0x09, 0x2c, 0x0c, 0xe3, 0x0a, 0xc9, 0x70, 0x71, 0x96, 0x64, 0xe6, 0xa6, 0x16,
	0x97, 0x24, 0xe6, 0x16, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xb0, 0x06, 0x21, 0x04, 0x84, 0xa4, 0xb8,
	0x38, 0x8a, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x25, 0x58, 0xc0, 0x1a, 0xe1, 0x7c, 0xa3,
	0x3c, 0x2e, 0x1e, 0xb0, 0xa5, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x06, 0x5c, 0x1c,
	0xc1, 0x89, 0x95, 0x60, 0x21, 0x21, 0x01, 0x3d, 0x88, 0x2b, 0xe1, 0xae, 0x92, 0xc2, 0x10, 0x31,
	0x60, 0x14, 0xd2, 0xe1, 0x62, 0x8b, 0x20, 0x5a, 0x7d, 0x12, 0x1b, 0xd8, 0xcb, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0xce, 0xdc, 0xc6, 0x01, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloTask, opts ...grpc.CallOption) (HelloService_SayHelloClient, error)
	XHello(ctx context.Context, in *HelloTask, opts ...grpc.CallOption) (*HelloTask, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloTask, opts ...grpc.CallOption) (HelloService_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/hello.HelloService/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_SayHelloClient interface {
	Recv() (*HelloTask, error)
	grpc.ClientStream
}

type helloServiceSayHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloClient) Recv() (*HelloTask, error) {
	m := new(HelloTask)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) XHello(ctx context.Context, in *HelloTask, opts ...grpc.CallOption) (*HelloTask, error) {
	out := new(HelloTask)
	err := c.cc.Invoke(ctx, "/hello.HelloService/XHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(*HelloTask, HelloService_SayHelloServer) error
	XHello(context.Context, *HelloTask) (*HelloTask, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) SayHello(req *HelloTask, srv HelloService_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServiceServer) XHello(ctx context.Context, req *HelloTask) (*HelloTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XHello not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloTask)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).SayHello(m, &helloServiceSayHelloServer{stream})
}

type HelloService_SayHelloServer interface {
	Send(*HelloTask) error
	grpc.ServerStream
}

type helloServiceSayHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloServer) Send(m *HelloTask) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_XHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).XHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/XHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).XHello(ctx, req.(*HelloTask))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "XHello",
			Handler:    _HelloService_XHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _HelloService_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello.proto",
}
