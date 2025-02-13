// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: weather/weather.proto

package proto

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
	WeatherService_GetWeather_FullMethodName    = "/rpc.proto.WeatherService/GetWeather"
	WeatherService_CreateWeather_FullMethodName = "/rpc.proto.WeatherService/CreateWeather"
	WeatherService_UpdateWeather_FullMethodName = "/rpc.proto.WeatherService/UpdateWeather"
	WeatherService_DeleteWeather_FullMethodName = "/rpc.proto.WeatherService/DeleteWeather"
)

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The WeatherService exposes gRPC methods
type WeatherServiceClient interface {
	GetWeather(ctx context.Context, in *GetWeatherRequest, opts ...grpc.CallOption) (*GetWeatherResponse, error)
	CreateWeather(ctx context.Context, in *CreateWeatherRequest, opts ...grpc.CallOption) (*CreateWeatherResponse, error)
	UpdateWeather(ctx context.Context, in *UpdateWeatherRequest, opts ...grpc.CallOption) (*UpdateWeatherResponse, error)
	DeleteWeather(ctx context.Context, in *DeleteWeatherRequest, opts ...grpc.CallOption) (*DeleteWeatherResponse, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetWeather(ctx context.Context, in *GetWeatherRequest, opts ...grpc.CallOption) (*GetWeatherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeatherResponse)
	err := c.cc.Invoke(ctx, WeatherService_GetWeather_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) CreateWeather(ctx context.Context, in *CreateWeatherRequest, opts ...grpc.CallOption) (*CreateWeatherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWeatherResponse)
	err := c.cc.Invoke(ctx, WeatherService_CreateWeather_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) UpdateWeather(ctx context.Context, in *UpdateWeatherRequest, opts ...grpc.CallOption) (*UpdateWeatherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateWeatherResponse)
	err := c.cc.Invoke(ctx, WeatherService_UpdateWeather_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) DeleteWeather(ctx context.Context, in *DeleteWeatherRequest, opts ...grpc.CallOption) (*DeleteWeatherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteWeatherResponse)
	err := c.cc.Invoke(ctx, WeatherService_DeleteWeather_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility.
//
// The WeatherService exposes gRPC methods
type WeatherServiceServer interface {
	GetWeather(context.Context, *GetWeatherRequest) (*GetWeatherResponse, error)
	CreateWeather(context.Context, *CreateWeatherRequest) (*CreateWeatherResponse, error)
	UpdateWeather(context.Context, *UpdateWeatherRequest) (*UpdateWeatherResponse, error)
	DeleteWeather(context.Context, *DeleteWeatherRequest) (*DeleteWeatherResponse, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeatherServiceServer struct{}

func (UnimplementedWeatherServiceServer) GetWeather(context.Context, *GetWeatherRequest) (*GetWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}
func (UnimplementedWeatherServiceServer) CreateWeather(context.Context, *CreateWeatherRequest) (*CreateWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWeather not implemented")
}
func (UnimplementedWeatherServiceServer) UpdateWeather(context.Context, *UpdateWeatherRequest) (*UpdateWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWeather not implemented")
}
func (UnimplementedWeatherServiceServer) DeleteWeather(context.Context, *DeleteWeatherRequest) (*DeleteWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWeather not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}
func (UnimplementedWeatherServiceServer) testEmbeddedByValue()                        {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	// If the following call pancis, it indicates UnimplementedWeatherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_GetWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeather(ctx, req.(*GetWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_CreateWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).CreateWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_CreateWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).CreateWeather(ctx, req.(*CreateWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_UpdateWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).UpdateWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_UpdateWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).UpdateWeather(ctx, req.(*UpdateWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_DeleteWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).DeleteWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherService_DeleteWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).DeleteWeather(ctx, req.(*DeleteWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.proto.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _WeatherService_GetWeather_Handler,
		},
		{
			MethodName: "CreateWeather",
			Handler:    _WeatherService_CreateWeather_Handler,
		},
		{
			MethodName: "UpdateWeather",
			Handler:    _WeatherService_UpdateWeather_Handler,
		},
		{
			MethodName: "DeleteWeather",
			Handler:    _WeatherService_DeleteWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather/weather.proto",
}
