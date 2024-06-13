// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productclient

import (
	"context"

	"github.com/billbliu/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ProductItem         = product.ProductItem
	ProductItemRequest  = product.ProductItemRequest
	ProductListRequest  = product.ProductListRequest
	ProductListResponse = product.ProductListResponse
	ProductRequest      = product.ProductRequest
	ProductResponse     = product.ProductResponse

	Product interface {
		Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error)
		Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
		ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Product(ctx, in, opts...)
}

func (m *defaultProduct) Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Products(ctx, in, opts...)
}

func (m *defaultProduct) ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}
