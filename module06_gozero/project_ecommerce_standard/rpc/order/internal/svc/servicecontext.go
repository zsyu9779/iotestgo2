package svc

import (
	"sync"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/internal/config"
	productclient "iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/product"
	userclient "iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type StoredOrder struct {
	OrderID     string
	UserID      int64
	Status      string
	TotalAmount float64
}

type ServiceContext struct {
	Config     config.Config
	UserRpc    userclient.User
	ProductRpc productclient.Product

	mu     sync.RWMutex
	orders map[string]*StoredOrder
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		orders:     make(map[string]*StoredOrder),
	}
}

func (s *ServiceContext) SaveOrder(order *StoredOrder) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders[order.OrderID] = order
}

func (s *ServiceContext) GetOrder(orderID string) (*StoredOrder, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.orders[orderID]
	return order, ok
}
