package logic

import (
	"context"
	"testing"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/internal/config"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/productpb"
)

func TestReserveStockReducesAvailableStock(t *testing.T) {
	svcCtx := svc.NewServiceContext(config.Config{})
	reserve := NewReserveStockLogic(context.Background(), svcCtx)
	get := NewGetProductLogic(context.Background(), svcCtx)

	resp, err := reserve.ReserveStock(&productpb.ReserveStockRequest{ProductId: 101, Quantity: 3})
	if err != nil {
		t.Fatalf("reserve stock returned error: %v", err)
	}
	if !resp.GetSuccess() {
		t.Fatalf("reserve stock failed: %s", resp.GetReason())
	}

	product, err := get.GetProduct(&productpb.GetProductRequest{ProductId: 101})
	if err != nil {
		t.Fatalf("get product returned error: %v", err)
	}
	if product.GetStock() != 7 {
		t.Fatalf("expected stock 7 after reservation, got %d", product.GetStock())
	}

	resp, err = reserve.ReserveStock(&productpb.ReserveStockRequest{ProductId: 101, Quantity: 8})
	if err != nil {
		t.Fatalf("reserve stock returned error: %v", err)
	}
	if resp.GetSuccess() {
		t.Fatal("expected reservation beyond remaining stock to fail")
	}
}
