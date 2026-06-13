package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/productpb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReserveStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReserveStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReserveStockLogic {
	return &ReserveStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReserveStockLogic) ReserveStock(in *productpb.ReserveStockRequest) (*productpb.ReserveStockResponse, error) {
	products := map[int64]*productpb.GetProductResponse{
		101: {ProductId: 101, Name: "Go Backend Book", Stock: 10, Price: 59.9},
		102: {ProductId: 102, Name: "Cloud Native Notebook", Stock: 5, Price: 29.9},
	}
	product, ok := products[in.GetProductId()]
	if !ok {
		return &productpb.ReserveStockResponse{Success: false, Reason: "product not found"}, nil
	}
	if in.GetQuantity() <= 0 {
		return &productpb.ReserveStockResponse{Success: false, Reason: "quantity must be positive"}, nil
	}
	if in.GetQuantity() > product.GetStock() {
		return &productpb.ReserveStockResponse{Success: false, Reason: "insufficient stock"}, nil
	}
	return &productpb.ReserveStockResponse{Success: true}, nil
}
