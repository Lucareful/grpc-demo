// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleClient interface {
	GetSimpleInfo(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type simpleClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleClient(cc grpc.ClientConnInterface) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) GetSimpleInfo(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.Simple/GetSimpleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
// All implementations must embed UnimplementedSimpleServer
// for forward compatibility
type SimpleServer interface {
	GetSimpleInfo(context.Context, *SimpleRequest) (*SimpleResponse, error)
	mustEmbedUnimplementedSimpleServer()
}

// UnimplementedSimpleServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (UnimplementedSimpleServer) GetSimpleInfo(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimpleInfo not implemented")
}
func (UnimplementedSimpleServer) mustEmbedUnimplementedSimpleServer() {}

// UnsafeSimpleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServer will
// result in compilation errors.
type UnsafeSimpleServer interface {
	mustEmbedUnimplementedSimpleServer()
}

func RegisterSimpleServer(s grpc.ServiceRegistrar, srv SimpleServer) {
	s.RegisterService(&Simple_ServiceDesc, srv)
}

func _Simple_GetSimpleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).GetSimpleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Simple/GetSimpleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).GetSimpleInfo(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Simple_ServiceDesc is the grpc.ServiceDesc for Simple service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Simple_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSimpleInfo",
			Handler:    _Simple_GetSimpleInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple.proto",
}