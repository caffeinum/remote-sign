// Code generated by protoc-gen-go. DO NOT EDIT.
// source: browser_sign.proto

/*
Package browser is a generated protocol buffer package.

It is generated from these files:
	browser_sign.proto

It has these top-level messages:
	SignMessageRequest
	SignMessageResponse
*/
package browser

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// from lnd's rpc.proto:
type SignMessageRequest struct {
	// / The message to be signed
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *SignMessageRequest) Reset()                    { *m = SignMessageRequest{} }
func (m *SignMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SignMessageRequest) ProtoMessage()               {}
func (*SignMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignMessageRequest) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type SignMessageResponse struct {
	// / The signature for the given message
	Signature string `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignMessageResponse) Reset()                    { *m = SignMessageResponse{} }
func (m *SignMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*SignMessageResponse) ProtoMessage()               {}
func (*SignMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignMessageResponse) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*SignMessageRequest)(nil), "browser.SignMessageRequest")
	proto.RegisterType((*SignMessageResponse)(nil), "browser.SignMessageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BrowserSign service

type BrowserSignClient interface {
	// * lncli: `signmessage`
	// SignMessage signs a message with this node's private key. The returned
	// signature string is `zbase32` encoded and pubkey recoverable, meaning that
	// only the message digest and signature are needed for verification.
	SignMessage(ctx context.Context, in *SignMessageRequest, opts ...grpc.CallOption) (*SignMessageResponse, error)
}

type browserSignClient struct {
	cc *grpc.ClientConn
}

func NewBrowserSignClient(cc *grpc.ClientConn) BrowserSignClient {
	return &browserSignClient{cc}
}

func (c *browserSignClient) SignMessage(ctx context.Context, in *SignMessageRequest, opts ...grpc.CallOption) (*SignMessageResponse, error) {
	out := new(SignMessageResponse)
	err := grpc.Invoke(ctx, "/browser.BrowserSign/SignMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BrowserSign service

type BrowserSignServer interface {
	// * lncli: `signmessage`
	// SignMessage signs a message with this node's private key. The returned
	// signature string is `zbase32` encoded and pubkey recoverable, meaning that
	// only the message digest and signature are needed for verification.
	SignMessage(context.Context, *SignMessageRequest) (*SignMessageResponse, error)
}

func RegisterBrowserSignServer(s *grpc.Server, srv BrowserSignServer) {
	s.RegisterService(&_BrowserSign_serviceDesc, srv)
}

func _BrowserSign_SignMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrowserSignServer).SignMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/browser.BrowserSign/SignMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrowserSignServer).SignMessage(ctx, req.(*SignMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BrowserSign_serviceDesc = grpc.ServiceDesc{
	ServiceName: "browser.BrowserSign",
	HandlerType: (*BrowserSignServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignMessage",
			Handler:    _BrowserSign_SignMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "browser_sign.proto",
}

func init() { proto.RegisterFile("browser_sign.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2a, 0xca, 0x2f,
	0x2f, 0x4e, 0x2d, 0x8a, 0x2f, 0xce, 0x4c, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x87, 0x8a, 0x29, 0xa9, 0x71, 0x09, 0x05, 0x67, 0xa6, 0xe7, 0xf9, 0xa6, 0x16, 0x17, 0x27, 0xa6,
	0xa7, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0x98, 0x4a, 0xc6, 0x5c, 0xc2, 0x28, 0xea, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x64, 0xb8, 0x38, 0x41, 0xa6, 0x26, 0x96, 0x94, 0x16, 0xa5, 0x82,
	0x95, 0x73, 0x06, 0x21, 0x04, 0x8c, 0xc2, 0xb9, 0xb8, 0x9d, 0x20, 0xf6, 0x80, 0xf4, 0x0a, 0x79,
	0x70, 0x71, 0x23, 0x99, 0x21, 0x24, 0xad, 0x07, 0x75, 0x84, 0x1e, 0xa6, 0x0b, 0xa4, 0x64, 0xb0,
	0x4b, 0x42, 0xac, 0x4d, 0x62, 0x03, 0xfb, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x02,
	0x28, 0xb3, 0xdb, 0x00, 0x00, 0x00,
}
