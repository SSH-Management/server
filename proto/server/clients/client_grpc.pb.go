// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package clients

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

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
	Delete(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, "/clients.ClientService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Delete(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error) {
	out := new(DeleteClientResponse)
	err := c.cc.Invoke(ctx, "/clients.ClientService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	Create(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
	Delete(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) Create(context.Context, *CreateClientRequest) (*CreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClientServiceServer) Delete(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Create(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Delete(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clients.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClientService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clients/client.proto",
}
