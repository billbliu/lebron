package svc

import (
	"github.com/billbliu/lebron/apps/order/rpc/internal/model"
	"github.com/billbliu/lebron/apps/product/rpc/productclient"
	"github.com/billbliu/lebron/apps/user/rpc/userclient"

	"github.com/billbliu/lebron/apps/order/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	OrderModel     model.OrdersModel
	OrderitemModel model.OrderitemModel
	ShippingModel  model.ShippingModel
	UserRpc        userclient.User
	ProductRpc     productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
