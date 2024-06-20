// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: product.proto

package product

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
	Product_Product_FullMethodName              = "/product.Product/Product"
	Product_Products_FullMethodName             = "/product.Product/Products"
	Product_ProductList_FullMethodName          = "/product.Product/ProductList"
	Product_OperationProducts_FullMethodName    = "/product.Product/OperationProducts"
	Product_UpdateProductStock_FullMethodName   = "/product.Product/UpdateProductStock"
	Product_CheckAndUpdateStock_FullMethodName  = "/product.Product/CheckAndUpdateStock"
	Product_CheckProductStock_FullMethodName    = "/product.Product/CheckProductStock"
	Product_RollbackProductStock_FullMethodName = "/product.Product/RollbackProductStock"
	Product_DecrStock_FullMethodName            = "/product.Product/DecrStock"
	Product_DecrStockRevert_FullMethodName      = "/product.Product/DecrStockRevert"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error)
	Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error)
	OperationProducts(ctx context.Context, in *OperationProductsRequest, opts ...grpc.CallOption) (*OperationProductsResponse, error)
	UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error)
	CheckAndUpdateStock(ctx context.Context, in *CheckAndUpdateStockRequest, opts ...grpc.CallOption) (*CheckAndUpdateStockResponse, error)
	CheckProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error)
	RollbackProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error)
	DecrStock(ctx context.Context, in *DecrStockRequest, opts ...grpc.CallOption) (*DecrStockResponse, error)
	DecrStockRevert(ctx context.Context, in *DecrStockRequest, opts ...grpc.CallOption) (*DecrStockResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductItem)
	err := c.cc.Invoke(ctx, Product_Product_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_Products_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductListResponse)
	err := c.cc.Invoke(ctx, Product_ProductList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) OperationProducts(ctx context.Context, in *OperationProductsRequest, opts ...grpc.CallOption) (*OperationProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationProductsResponse)
	err := c.cc.Invoke(ctx, Product_OperationProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductStockResponse)
	err := c.cc.Invoke(ctx, Product_UpdateProductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CheckAndUpdateStock(ctx context.Context, in *CheckAndUpdateStockRequest, opts ...grpc.CallOption) (*CheckAndUpdateStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAndUpdateStockResponse)
	err := c.cc.Invoke(ctx, Product_CheckAndUpdateStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CheckProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductStockResponse)
	err := c.cc.Invoke(ctx, Product_CheckProductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) RollbackProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*UpdateProductStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductStockResponse)
	err := c.cc.Invoke(ctx, Product_RollbackProductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DecrStock(ctx context.Context, in *DecrStockRequest, opts ...grpc.CallOption) (*DecrStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecrStockResponse)
	err := c.cc.Invoke(ctx, Product_DecrStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DecrStockRevert(ctx context.Context, in *DecrStockRequest, opts ...grpc.CallOption) (*DecrStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecrStockResponse)
	err := c.cc.Invoke(ctx, Product_DecrStockRevert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	Product(context.Context, *ProductItemRequest) (*ProductItem, error)
	Products(context.Context, *ProductRequest) (*ProductResponse, error)
	ProductList(context.Context, *ProductListRequest) (*ProductListResponse, error)
	OperationProducts(context.Context, *OperationProductsRequest) (*OperationProductsResponse, error)
	UpdateProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error)
	CheckAndUpdateStock(context.Context, *CheckAndUpdateStockRequest) (*CheckAndUpdateStockResponse, error)
	CheckProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error)
	RollbackProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error)
	DecrStock(context.Context, *DecrStockRequest) (*DecrStockResponse, error)
	DecrStockRevert(context.Context, *DecrStockRequest) (*DecrStockResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) Product(context.Context, *ProductItemRequest) (*ProductItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Product not implemented")
}
func (UnimplementedProductServer) Products(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Products not implemented")
}
func (UnimplementedProductServer) ProductList(context.Context, *ProductListRequest) (*ProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (UnimplementedProductServer) OperationProducts(context.Context, *OperationProductsRequest) (*OperationProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperationProducts not implemented")
}
func (UnimplementedProductServer) UpdateProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductStock not implemented")
}
func (UnimplementedProductServer) CheckAndUpdateStock(context.Context, *CheckAndUpdateStockRequest) (*CheckAndUpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndUpdateStock not implemented")
}
func (UnimplementedProductServer) CheckProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProductStock not implemented")
}
func (UnimplementedProductServer) RollbackProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackProductStock not implemented")
}
func (UnimplementedProductServer) DecrStock(context.Context, *DecrStockRequest) (*DecrStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrStock not implemented")
}
func (UnimplementedProductServer) DecrStockRevert(context.Context, *DecrStockRequest) (*DecrStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrStockRevert not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_Product_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Product(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_Product_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Product(ctx, req.(*ProductItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Products_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Products(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_Products_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Products(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductList(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_OperationProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).OperationProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_OperationProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).OperationProducts(ctx, req.(*OperationProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_UpdateProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProductStock(ctx, req.(*UpdateProductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CheckAndUpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndUpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CheckAndUpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CheckAndUpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CheckAndUpdateStock(ctx, req.(*CheckAndUpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CheckProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CheckProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CheckProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CheckProductStock(ctx, req.(*UpdateProductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_RollbackProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).RollbackProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_RollbackProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).RollbackProductStock(ctx, req.(*UpdateProductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DecrStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DecrStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DecrStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DecrStock(ctx, req.(*DecrStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DecrStockRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DecrStockRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DecrStockRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DecrStockRevert(ctx, req.(*DecrStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Product",
			Handler:    _Product_Product_Handler,
		},
		{
			MethodName: "Products",
			Handler:    _Product_Products_Handler,
		},
		{
			MethodName: "ProductList",
			Handler:    _Product_ProductList_Handler,
		},
		{
			MethodName: "OperationProducts",
			Handler:    _Product_OperationProducts_Handler,
		},
		{
			MethodName: "UpdateProductStock",
			Handler:    _Product_UpdateProductStock_Handler,
		},
		{
			MethodName: "CheckAndUpdateStock",
			Handler:    _Product_CheckAndUpdateStock_Handler,
		},
		{
			MethodName: "CheckProductStock",
			Handler:    _Product_CheckProductStock_Handler,
		},
		{
			MethodName: "RollbackProductStock",
			Handler:    _Product_RollbackProductStock_Handler,
		},
		{
			MethodName: "DecrStock",
			Handler:    _Product_DecrStock_Handler,
		},
		{
			MethodName: "DecrStockRevert",
			Handler:    _Product_DecrStockRevert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
