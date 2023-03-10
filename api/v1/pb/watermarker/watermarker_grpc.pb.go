// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: watermarker.proto

package watermarker

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

// WatermarkerClient is the client API for Watermarker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatermarkerClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusResponse, error)
	AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDocumentResponse, error)
}

type watermarkerClient struct {
	cc grpc.ClientConnInterface
}

func NewWatermarkerClient(cc grpc.ClientConnInterface) WatermarkerClient {
	return &watermarkerClient{cc}
}

func (c *watermarkerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/watermarker.Watermarker/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkerClient) Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkResponse, error) {
	out := new(WatermarkResponse)
	err := c.cc.Invoke(ctx, "/watermarker.Watermarker/Watermark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkerClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/watermarker.Watermarker/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkerClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusResponse, error) {
	out := new(ServiceStatusResponse)
	err := c.cc.Invoke(ctx, "/watermarker.Watermarker/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkerClient) AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDocumentResponse, error) {
	out := new(AddDocumentResponse)
	err := c.cc.Invoke(ctx, "/watermarker.Watermarker/AddDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatermarkerServer is the server API for Watermarker service.
// All implementations must embed UnimplementedWatermarkerServer
// for forward compatibility
type WatermarkerServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Watermark(context.Context, *WatermarkRequest) (*WatermarkResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusResponse, error)
	AddDocument(context.Context, *AddDocumentRequest) (*AddDocumentResponse, error)
	mustEmbedUnimplementedWatermarkerServer()
}

// UnimplementedWatermarkerServer must be embedded to have forward compatible implementations.
type UnimplementedWatermarkerServer struct {
}

func (UnimplementedWatermarkerServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWatermarkerServer) Watermark(context.Context, *WatermarkRequest) (*WatermarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watermark not implemented")
}
func (UnimplementedWatermarkerServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedWatermarkerServer) ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}
func (UnimplementedWatermarkerServer) AddDocument(context.Context, *AddDocumentRequest) (*AddDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDocument not implemented")
}
func (UnimplementedWatermarkerServer) mustEmbedUnimplementedWatermarkerServer() {}

// UnsafeWatermarkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatermarkerServer will
// result in compilation errors.
type UnsafeWatermarkerServer interface {
	mustEmbedUnimplementedWatermarkerServer()
}

func RegisterWatermarkerServer(s grpc.ServiceRegistrar, srv WatermarkerServer) {
	s.RegisterService(&Watermarker_ServiceDesc, srv)
}

func _Watermarker_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarker.Watermarker/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarker_Watermark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatermarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkerServer).Watermark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarker.Watermarker/Watermark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkerServer).Watermark(ctx, req.(*WatermarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarker.Watermarker/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkerServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarker_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkerServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarker.Watermarker/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkerServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermarker_AddDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkerServer).AddDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watermarker.Watermarker/AddDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkerServer).AddDocument(ctx, req.(*AddDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Watermarker_ServiceDesc is the grpc.ServiceDesc for Watermarker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watermarker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "watermarker.Watermarker",
	HandlerType: (*WatermarkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Watermarker_Get_Handler,
		},
		{
			MethodName: "Watermark",
			Handler:    _Watermarker_Watermark_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Watermarker_Status_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Watermarker_ServiceStatus_Handler,
		},
		{
			MethodName: "AddDocument",
			Handler:    _Watermarker_AddDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watermarker.proto",
}
