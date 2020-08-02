// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package myservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	DoSomething(ctx context.Context, in *DoSomethingRequest, opts ...grpc.CallOption) (*DoSomethingResponse, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) DoSomething(ctx context.Context, in *DoSomethingRequest, opts ...grpc.CallOption) (*DoSomethingResponse, error) {
	out := new(DoSomethingResponse)
	err := c.cc.Invoke(ctx, "/test.MyService/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	DoSomething(context.Context, *DoSomethingRequest) (*DoSomethingResponse, error)
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (*UnimplementedMyServiceServer) DoSomething(context.Context, *DoSomethingRequest) (*DoSomethingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSomething not implemented")
}
func (*UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

func RegisterMyServiceServer(s *grpc.Server, srv MyServiceServer) {
	s.RegisterService(&_MyService_serviceDesc, srv)
}

func _MyService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoSomethingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.MyService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).DoSomething(ctx, req.(*DoSomethingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _MyService_DoSomething_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
