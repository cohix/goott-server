// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	model "github.com/cohix/goott-server/model"
	simplcrypto "github.com/cohix/simplcrypto"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "goott.server.service.Empty")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x11, 0x3c, 0x85, 0x80, 0x70, 0x04, 0xb5, 0x88, 0xc5, 0x89, 0x20, 0x56, 0x4e, 0x40,
	0xff, 0x80, 0x0a, 0x7a, 0x58, 0xd8, 0x78, 0x9d, 0xdd, 0x6d, 0x1c, 0x36, 0x81, 0x8d, 0x13, 0x93,
	0x89, 0xb8, 0xff, 0xd5, 0x1f, 0x23, 0x9b, 0xa4, 0x11, 0x56, 0xb8, 0xf6, 0xbd, 0x8f, 0xf7, 0x66,
	0x9e, 0x38, 0x4a, 0x18, 0xbf, 0x9c, 0x41, 0x08, 0x91, 0x98, 0xe4, 0x71, 0x4f, 0xc4, 0x0c, 0x93,
	0x88, 0x11, 0x9a, 0xa7, 0x4e, 0x3d, 0xbd, 0xe3, 0xa0, 0x0b, 0xa2, 0xb7, 0x99, 0x6d, 0xa5, 0xd5,
	0x2a, 0x39, 0x1f, 0x06, 0x13, 0xc7, 0xc0, 0xd4, 0x5c, 0x8f, 0x29, 0x6d, 0xfb, 0x16, 0x77, 0x71,
	0x28, 0x16, 0x8f, 0x3e, 0xf0, 0x78, 0xf3, 0xb3, 0x27, 0x16, 0xeb, 0x29, 0x5a, 0x3e, 0x8b, 0xfd,
	0xfb, 0xcc, 0x56, 0xae, 0xe0, 0x4f, 0x55, 0x69, 0x80, 0xc9, 0x79, 0xc5, 0xcf, 0x8c, 0x89, 0xd5,
	0xf9, 0xff, 0x40, 0x0a, 0xf4, 0x91, 0x50, 0xde, 0x89, 0xe5, 0x1a, 0x79, 0x83, 0x26, 0x22, 0xbf,
	0xd4, 0x5e, 0x79, 0x02, 0xe5, 0x26, 0xa8, 0x47, 0x41, 0x93, 0xd5, 0xbc, 0x2c, 0x9f, 0xc4, 0x72,
	0xb3, 0x63, 0xc2, 0x19, 0xcc, 0x4d, 0x03, 0xe5, 0xbd, 0x87, 0xab, 0xb7, 0xcb, 0xde, 0xb1, 0xcd,
	0x1d, 0x18, 0xf2, 0xda, 0x90, 0x75, 0xdf, 0xba, 0xe0, 0xd7, 0x15, 0xd7, 0x0d, 0xef, 0x0e, 0xca,
	0x2e, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x5b, 0x90, 0x41, 0x77, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoottClient is the client API for Goott service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoottClient interface {
	Auth(ctx context.Context, in *model.AuthRequest, opts ...grpc.CallOption) (*model.AuthResponse, error)
	GetSecretMessage(ctx context.Context, in *simplcrypto.Message, opts ...grpc.CallOption) (*simplcrypto.Message, error)
	SetSecretMessage(ctx context.Context, in *simplcrypto.Message, opts ...grpc.CallOption) (*Empty, error)
}

type goottClient struct {
	cc *grpc.ClientConn
}

func NewGoottClient(cc *grpc.ClientConn) GoottClient {
	return &goottClient{cc}
}

func (c *goottClient) Auth(ctx context.Context, in *model.AuthRequest, opts ...grpc.CallOption) (*model.AuthResponse, error) {
	out := new(model.AuthResponse)
	err := c.cc.Invoke(ctx, "/goott.server.service.Goott/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goottClient) GetSecretMessage(ctx context.Context, in *simplcrypto.Message, opts ...grpc.CallOption) (*simplcrypto.Message, error) {
	out := new(simplcrypto.Message)
	err := c.cc.Invoke(ctx, "/goott.server.service.Goott/GetSecretMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goottClient) SetSecretMessage(ctx context.Context, in *simplcrypto.Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/goott.server.service.Goott/SetSecretMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoottServer is the server API for Goott service.
type GoottServer interface {
	Auth(context.Context, *model.AuthRequest) (*model.AuthResponse, error)
	GetSecretMessage(context.Context, *simplcrypto.Message) (*simplcrypto.Message, error)
	SetSecretMessage(context.Context, *simplcrypto.Message) (*Empty, error)
}

func RegisterGoottServer(s *grpc.Server, srv GoottServer) {
	s.RegisterService(&_Goott_serviceDesc, srv)
}

func _Goott_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoottServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goott.server.service.Goott/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoottServer).Auth(ctx, req.(*model.AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goott_GetSecretMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(simplcrypto.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoottServer).GetSecretMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goott.server.service.Goott/GetSecretMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoottServer).GetSecretMessage(ctx, req.(*simplcrypto.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goott_SetSecretMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(simplcrypto.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoottServer).SetSecretMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goott.server.service.Goott/SetSecretMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoottServer).SetSecretMessage(ctx, req.(*simplcrypto.Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goott_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goott.server.service.Goott",
	HandlerType: (*GoottServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Goott_Auth_Handler,
		},
		{
			MethodName: "GetSecretMessage",
			Handler:    _Goott_GetSecretMessage_Handler,
		},
		{
			MethodName: "SetSecretMessage",
			Handler:    _Goott_SetSecretMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
