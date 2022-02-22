// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crashy.proto

package proto

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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76afcfc1dbadb4c3, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76afcfc1dbadb4c3, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "proto.HelloResponse")
}

func init() { proto.RegisterFile("crashy.proto", fileDescriptor_76afcfc1dbadb4c3) }

var fileDescriptor_76afcfc1dbadb4c3 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2e, 0x4a, 0x2c,
	0xce, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x4a, 0x5c, 0x3c,
	0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x26, 0x17,
	0x2f, 0x54, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71,
	0x71, 0x62, 0x3a, 0x4c, 0x1d, 0x8c, 0x6b, 0xe4, 0xcc, 0xc5, 0xe6, 0x0c, 0xb6, 0x45, 0xc8, 0x92,
	0x8b, 0x0b, 0xac, 0x29, 0x24, 0x23, 0xb5, 0x28, 0x55, 0x48, 0x18, 0x62, 0xab, 0x1e, 0xb2, 0x5d,
	0x52, 0x22, 0xa8, 0x82, 0x10, 0xc3, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xc2, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1e, 0x37, 0xc3, 0x19, 0xb1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CrashyClient is the client API for Crashy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrashyClient interface {
	// Sends a greeting
	HelloThere(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type crashyClient struct {
	cc *grpc.ClientConn
}

func NewCrashyClient(cc *grpc.ClientConn) CrashyClient {
	return &crashyClient{cc}
}

func (c *crashyClient) HelloThere(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.Crashy/HelloThere", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrashyServer is the server API for Crashy service.
type CrashyServer interface {
	// Sends a greeting
	HelloThere(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedCrashyServer can be embedded to have forward compatible implementations.
type UnimplementedCrashyServer struct {
}

func (*UnimplementedCrashyServer) HelloThere(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloThere not implemented")
}

func RegisterCrashyServer(s *grpc.Server, srv CrashyServer) {
	s.RegisterService(&_Crashy_serviceDesc, srv)
}

func _Crashy_HelloThere_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrashyServer).HelloThere(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Crashy/HelloThere",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrashyServer).HelloThere(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crashy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Crashy",
	HandlerType: (*CrashyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloThere",
			Handler:    _Crashy_HelloThere_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crashy.proto",
}
