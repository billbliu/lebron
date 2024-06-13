package svc

import (
	"github.com/billbliu/lebron/apps/product/rpc/internal/config"
	"github.com/billbliu/lebron/apps/product/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
)

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	OperationModel model.ProductOperationModel
	SingleGroup    singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		ProductModel:   model.NewProductModel(conn, c.CacheRedis),
		CategoryModel:  model.NewCategoryModel(conn, c.CacheRedis),
		OperationModel: model.NewProductOperationModel(conn, c.CacheRedis),
	}
}
