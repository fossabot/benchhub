// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
	Package grpc is a generated protocol buffer package.

	It is generated from these files:
		rpc.proto

	It has these top-level messages:
*/
package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import bh "github.com/benchhub/benchhub/pkg/bhpb"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for BenchHubAgent service

type BenchHubAgentClient interface {
	Ping(ctx context.Context, in *bh.Ping, opts ...grpc1.CallOption) (*bh.Pong, error)
	NodeInfo(ctx context.Context, in *bh.NodeInfoReq, opts ...grpc1.CallOption) (*bh.NodeInfoRes, error)
}

type benchHubAgentClient struct {
	cc *grpc1.ClientConn
}

func NewBenchHubAgentClient(cc *grpc1.ClientConn) BenchHubAgentClient {
	return &benchHubAgentClient{cc}
}

func (c *benchHubAgentClient) Ping(ctx context.Context, in *bh.Ping, opts ...grpc1.CallOption) (*bh.Pong, error) {
	out := new(bh.Pong)
	err := grpc1.Invoke(ctx, "/benchubcentralrpc.BenchHubAgent/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benchHubAgentClient) NodeInfo(ctx context.Context, in *bh.NodeInfoReq, opts ...grpc1.CallOption) (*bh.NodeInfoRes, error) {
	out := new(bh.NodeInfoRes)
	err := grpc1.Invoke(ctx, "/benchubcentralrpc.BenchHubAgent/NodeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BenchHubAgent service

type BenchHubAgentServer interface {
	Ping(context.Context, *bh.Ping) (*bh.Pong, error)
	NodeInfo(context.Context, *bh.NodeInfoReq) (*bh.NodeInfoRes, error)
}

func RegisterBenchHubAgentServer(s *grpc1.Server, srv BenchHubAgentServer) {
	s.RegisterService(&_BenchHubAgent_serviceDesc, srv)
}

func _BenchHubAgent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(bh.Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchHubAgentServer).Ping(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/benchubcentralrpc.BenchHubAgent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchHubAgentServer).Ping(ctx, req.(*bh.Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _BenchHubAgent_NodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(bh.NodeInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchHubAgentServer).NodeInfo(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/benchubcentralrpc.BenchHubAgent/NodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchHubAgentServer).NodeInfo(ctx, req.(*bh.NodeInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BenchHubAgent_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "benchubcentralrpc.BenchHubAgent",
	HandlerType: (*BenchHubAgentServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BenchHubAgent_Ping_Handler,
		},
		{
			MethodName: "NodeInfo",
			Handler:    _BenchHubAgent_NodeInfo_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0x4a, 0xcd, 0x4b, 0xce, 0x28, 0x4d, 0x4a, 0x4e,
	0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0x29, 0x2a, 0x48, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4c, 0x2a, 0x4d, 0x03,
	0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0x62, 0x82, 0x94, 0x21, 0x92, 0x72, 0xb0, 0x61, 0x19, 0xa5, 0x49,
	0x08, 0x46, 0x41, 0x76, 0xba, 0x7e, 0x52, 0x46, 0x41, 0x92, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0x31,
	0x44, 0x8b, 0x51, 0x2c, 0x17, 0xaf, 0x13, 0x48, 0x81, 0x47, 0x69, 0x92, 0x63, 0x7a, 0x6a, 0x5e,
	0x89, 0x90, 0x0c, 0x17, 0x4b, 0x40, 0x66, 0x5e, 0xba, 0x10, 0x87, 0x5e, 0x52, 0x86, 0x1e, 0x88,
	0x25, 0x05, 0x61, 0xe5, 0xe7, 0xa5, 0x2b, 0x31, 0x08, 0xe9, 0x71, 0x71, 0xf8, 0xe5, 0xa7, 0xa4,
	0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0xf1, 0x83, 0xc4, 0x61, 0xbc, 0xa0, 0xd4, 0x42, 0x29, 0x34, 0x81,
	0x62, 0x25, 0x06, 0x27, 0xb1, 0x13, 0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x28, 0x96, 0xf4, 0xa2, 0x82, 0xe4, 0x24, 0x36, 0xb0, 0xed,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x8e, 0x20, 0x3e, 0xff, 0x00, 0x00, 0x00,
}
