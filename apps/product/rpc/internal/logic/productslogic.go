package logic

import (
	"context"
	"strconv"
	"strings"

	"github.com/billbliu/lebron/apps/product/rpc/internal/model"
	"github.com/billbliu/lebron/apps/product/rpc/internal/svc"
	"github.com/billbliu/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	products := make(map[int64]*product.ProductItem)
	productIds := strings.Split(in.ProductIds, ",")
	productList, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, productId := range productIds {
			source <- productId
		}
	}, func(item interface{}, writer mr.Writer[*model.Product], cancel func(error)) {
		productIdStr := item.(string)
		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			return
		}

		product, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(productId))
		if err != nil {
			return
		}
		writer.Write(product)
	}, func(pipe <-chan *model.Product, writer mr.Writer[[]*model.Product], cancel func(error)) {
		var products []*model.Product
		for item := range pipe {
			products = append(products, item)
		}
		writer.Write(products)
	})

	if err != nil {
		return nil, err
	}

	for _, p := range productList {
		pID := int64(p.Id)
		products[pID] = &product.ProductItem{
			ProductId: pID,
			Name:      p.Name,
		}
	}

	return &product.ProductResponse{Products: products}, nil
}
