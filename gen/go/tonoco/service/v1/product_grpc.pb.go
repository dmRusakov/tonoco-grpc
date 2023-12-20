// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tonoco/service/v1/product.proto

package tonoco_service_v1

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

const (
	TonocoProductService_GetProduct_FullMethodName    = "/tonoco.service.v1.TonocoProductService/GetProduct"
	TonocoProductService_ListProducts_FullMethodName  = "/tonoco.service.v1.TonocoProductService/ListProducts"
	TonocoProductService_CreateProduct_FullMethodName = "/tonoco.service.v1.TonocoProductService/CreateProduct"
	TonocoProductService_UpdateProduct_FullMethodName = "/tonoco.service.v1.TonocoProductService/UpdateProduct"
	TonocoProductService_PatchProduct_FullMethodName  = "/tonoco.service.v1.TonocoProductService/PatchProduct"
	TonocoProductService_DeleteProduct_FullMethodName = "/tonoco.service.v1.TonocoProductService/DeleteProduct"
)

// TonocoProductServiceClient is the client API for TonocoProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TonocoProductServiceClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	PatchProduct(ctx context.Context, in *PatchProductRequest, opts ...grpc.CallOption) (*PatchProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
}

type tonocoProductServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTonocoProductServiceClient(cc grpc.ClientConnInterface) TonocoProductServiceClient {
	return &tonocoProductServiceClient{cc}
}

func (c *tonocoProductServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, TonocoProductService_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonocoProductServiceClient) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, TonocoProductService_ListProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonocoProductServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, TonocoProductService_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonocoProductServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, TonocoProductService_UpdateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonocoProductServiceClient) PatchProduct(ctx context.Context, in *PatchProductRequest, opts ...grpc.CallOption) (*PatchProductResponse, error) {
	out := new(PatchProductResponse)
	err := c.cc.Invoke(ctx, TonocoProductService_PatchProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonocoProductServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, TonocoProductService_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TonocoProductServiceServer is the server API for TonocoProductService service.
// All implementations must embed UnimplementedTonocoProductServiceServer
// for forward compatibility
type TonocoProductServiceServer interface {
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	PatchProduct(context.Context, *PatchProductRequest) (*PatchProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	mustEmbedUnimplementedTonocoProductServiceServer()
}

// UnimplementedTonocoProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTonocoProductServiceServer struct {
}

func (UnimplementedTonocoProductServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedTonocoProductServiceServer) ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedTonocoProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedTonocoProductServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedTonocoProductServiceServer) PatchProduct(context.Context, *PatchProductRequest) (*PatchProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchProduct not implemented")
}
func (UnimplementedTonocoProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedTonocoProductServiceServer) mustEmbedUnimplementedTonocoProductServiceServer() {}

// UnsafeTonocoProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TonocoProductServiceServer will
// result in compilation errors.
type UnsafeTonocoProductServiceServer interface {
	mustEmbedUnimplementedTonocoProductServiceServer()
}

func RegisterTonocoProductServiceServer(s grpc.ServiceRegistrar, srv TonocoProductServiceServer) {
	s.RegisterService(&TonocoProductService_ServiceDesc, srv)
}

func _TonocoProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonocoProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TonocoProductService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonocoProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TonocoProductService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonocoProductServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TonocoProductService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonocoProductServiceServer).ListProducts(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TonocoProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonocoProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TonocoProductService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonocoProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TonocoProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonocoProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TonocoProductService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonocoProductServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TonocoProductService_PatchProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonocoProductServiceServer).PatchProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TonocoProductService_PatchProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonocoProductServiceServer).PatchProduct(ctx, req.(*PatchProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TonocoProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonocoProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TonocoProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonocoProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TonocoProductService_ServiceDesc is the grpc.ServiceDesc for TonocoProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TonocoProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tonoco.service.v1.TonocoProductService",
	HandlerType: (*TonocoProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _TonocoProductService_GetProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _TonocoProductService_ListProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _TonocoProductService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _TonocoProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "PatchProduct",
			Handler:    _TonocoProductService_PatchProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _TonocoProductService_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tonoco/service/v1/product.proto",
}
