// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: semrel/v1/plugin.proto

package semrelpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SemrelPlugin_Release_FullMethodName = "/semrel.v1.SemrelPlugin/Release"
)

// SemrelPluginClient is the client API for SemrelPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SemrelPluginClient interface {
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
}

type semrelPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSemrelPluginClient(cc grpc.ClientConnInterface) SemrelPluginClient {
	return &semrelPluginClient{cc}
}

func (c *semrelPluginClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, SemrelPlugin_Release_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SemrelPluginServer is the server API for SemrelPlugin service.
// All implementations must embed UnimplementedSemrelPluginServer
// for forward compatibility
type SemrelPluginServer interface {
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	mustEmbedUnimplementedSemrelPluginServer()
}

// UnimplementedSemrelPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSemrelPluginServer struct {
}

func (UnimplementedSemrelPluginServer) Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (UnimplementedSemrelPluginServer) mustEmbedUnimplementedSemrelPluginServer() {}

// UnsafeSemrelPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SemrelPluginServer will
// result in compilation errors.
type UnsafeSemrelPluginServer interface {
	mustEmbedUnimplementedSemrelPluginServer()
}

func RegisterSemrelPluginServer(s grpc.ServiceRegistrar, srv SemrelPluginServer) {
	s.RegisterService(&SemrelPlugin_ServiceDesc, srv)
}

func _SemrelPlugin_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SemrelPluginServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SemrelPlugin_Release_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SemrelPluginServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SemrelPlugin_ServiceDesc is the grpc.ServiceDesc for SemrelPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SemrelPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "semrel.v1.SemrelPlugin",
	HandlerType: (*SemrelPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Release",
			Handler:    _SemrelPlugin_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "semrel/v1/plugin.proto",
}
