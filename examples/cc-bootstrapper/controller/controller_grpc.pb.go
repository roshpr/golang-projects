// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: controller.proto

package controller

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ControllerRequest_Fetch_FullMethodName = "/ControllerRequest/fetch"
)

// ControllerRequestClient is the client API for ControllerRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerRequestClient interface {
	Fetch(ctx context.Context, in *ConnectorCtrlReq, opts ...grpc.CallOption) (*ConnectorCtrlResp, error)
}

type controllerRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerRequestClient(cc grpc.ClientConnInterface) ControllerRequestClient {
	return &controllerRequestClient{cc}
}

func (c *controllerRequestClient) Fetch(ctx context.Context, in *ConnectorCtrlReq, opts ...grpc.CallOption) (*ConnectorCtrlResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectorCtrlResp)
	err := c.cc.Invoke(ctx, ControllerRequest_Fetch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerRequestServer is the server API for ControllerRequest service.
// All implementations must embed UnimplementedControllerRequestServer
// for forward compatibility.
type ControllerRequestServer interface {
	Fetch(context.Context, *ConnectorCtrlReq) (*ConnectorCtrlResp, error)
	mustEmbedUnimplementedControllerRequestServer()
}

// UnimplementedControllerRequestServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedControllerRequestServer struct{}

func (UnimplementedControllerRequestServer) Fetch(context.Context, *ConnectorCtrlReq) (*ConnectorCtrlResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedControllerRequestServer) mustEmbedUnimplementedControllerRequestServer() {}
func (UnimplementedControllerRequestServer) testEmbeddedByValue()                           {}

// UnsafeControllerRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerRequestServer will
// result in compilation errors.
type UnsafeControllerRequestServer interface {
	mustEmbedUnimplementedControllerRequestServer()
}

func RegisterControllerRequestServer(s grpc.ServiceRegistrar, srv ControllerRequestServer) {
	// If the following call pancis, it indicates UnimplementedControllerRequestServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ControllerRequest_ServiceDesc, srv)
}

func _ControllerRequest_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectorCtrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerRequestServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerRequest_Fetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerRequestServer).Fetch(ctx, req.(*ConnectorCtrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerRequest_ServiceDesc is the grpc.ServiceDesc for ControllerRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ControllerRequest",
	HandlerType: (*ControllerRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "fetch",
			Handler:    _ControllerRequest_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}
