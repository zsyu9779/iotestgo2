package svc

import (
	"sync"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/internal/config"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/productpb"
)

type ServiceContext struct {
	Config   config.Config
	mu       sync.Mutex
	products map[int64]*productpb.GetProductResponse
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		products: map[int64]*productpb.GetProductResponse{
			101: {ProductId: 101, Name: "Go Backend Book", Stock: 10, Price: 59.9},
			102: {ProductId: 102, Name: "Cloud Native Notebook", Stock: 5, Price: 29.9},
		},
	}
}

func (s *ServiceContext) GetProduct(productID int64) (*productpb.GetProductResponse, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, ok := s.products[productID]
	if !ok {
		return nil, false
	}
	return cloneProduct(product), true
}

func (s *ServiceContext) ReserveStock(productID, quantity int64) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, ok := s.products[productID]
	if !ok {
		return "product not found", false
	}
	if quantity <= 0 {
		return "quantity must be positive", false
	}
	if quantity > product.GetStock() {
		return "insufficient stock", false
	}
	product.Stock -= quantity
	return "", true
}

func cloneProduct(product *productpb.GetProductResponse) *productpb.GetProductResponse {
	return &productpb.GetProductResponse{
		ProductId: product.GetProductId(),
		Name:      product.GetName(),
		Stock:     product.GetStock(),
		Price:     product.GetPrice(),
	}
}
