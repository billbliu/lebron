package svc

import (
	"github.com/billbliu/lebron/apps/app/api/internal/config"
	"github.com/billbliu/lebron/apps/order/rpc/orderclient"
	"github.com/billbliu/lebron/apps/product/rpc/productclient"
	"github.com/billbliu/lebron/apps/reply/rpc/replyclient"
	"github.com/billbliu/lebron/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
	ReplyRPC   replyclient.Reply
	UserRPC    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		ReplyRPC:   replyclient.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
		UserRPC:    userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
