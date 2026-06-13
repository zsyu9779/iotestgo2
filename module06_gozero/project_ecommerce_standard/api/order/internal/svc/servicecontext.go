// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"iotestgo/module06_gozero/project_ecommerce_standard/api/order/internal/config"
	orderclient "iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/order"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
