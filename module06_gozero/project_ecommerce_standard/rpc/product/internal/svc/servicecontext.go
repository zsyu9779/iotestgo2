package svc

import "iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
