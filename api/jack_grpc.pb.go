// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: api/jack.proto

package server

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
	Jack_GetStudent_FullMethodName = "/server.Jack/GetStudent"
)

// JackClient is the client API for Jack service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JackClient interface {
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error)
}

type jackClient struct {
	cc grpc.ClientConnInterface
}

func NewJackClient(cc grpc.ClientConnInterface) JackClient {
	return &jackClient{cc}
}

func (c *jackClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStudentResponse)
	err := c.cc.Invoke(ctx, Jack_GetStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JackServer is the server API for Jack service.
// All implementations must embed UnimplementedJackServer
// for forward compatibility
type JackServer interface {
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error)
	mustEmbedUnimplementedJackServer()
}

// UnimplementedJackServer must be embedded to have forward compatible implementations.
type UnimplementedJackServer struct {
}

func (UnimplementedJackServer) GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedJackServer) mustEmbedUnimplementedJackServer() {}

// UnsafeJackServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JackServer will
// result in compilation errors.
type UnsafeJackServer interface {
	mustEmbedUnimplementedJackServer()
}

func RegisterJackServer(s grpc.ServiceRegistrar, srv JackServer) {
	s.RegisterService(&Jack_ServiceDesc, srv)
}

func _Jack_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JackServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jack_GetStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JackServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jack_ServiceDesc is the grpc.ServiceDesc for Jack service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jack_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Jack",
	HandlerType: (*JackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudent",
			Handler:    _Jack_GetStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/jack.proto",
}