package logic

import (
	"context"

	"github.com/billbliu/lebron/apps/product/rpc/internal/svc"
	"github.com/billbliu/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationProductsLogic struct {
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	productListLogic *ProductListLogic
	logx.Logger
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		productListLogic: NewProductListLogic(ctx, svcCtx),
	}
}

const (
	validStatus          = 1
	operationProductsKey = "operation#products"
)

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductsRequest) (*product.OperationProductsResponse, error) {
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)
	if ok {
		return &product.OperationProductsResponse{Products: opProducts.([]*product.ProductItem)}, nil
	}

	// 取出所有上线的运营位产品
	pos, err := l.svcCtx.OperationModel.OperationProducts(l.ctx, validStatus)
	if err != nil {
		return nil, err
	}

	var pids []int64
	for _, p := range pos {
		pids = append(pids, int64(p.ProductId))
	}

	// 根据上线的运营位产品id列表查出所有商品信息
	products, err := l.productListLogic.productsByIds(l.ctx, pids)
	if err != nil {
		return nil, err
	}

	var pItems []*product.ProductItem
	for _, p := range products {
		pItems = append(pItems, &product.ProductItem{
			ProductId: int64(p.Id),
			Name:      p.Name,
		})
	}

	l.svcCtx.LocalCache.Set(operationProductsKey, pItems)
	return &product.OperationProductsResponse{Products: pItems}, nil
}
