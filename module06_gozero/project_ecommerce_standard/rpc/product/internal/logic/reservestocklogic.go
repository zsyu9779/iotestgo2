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
	reason, ok := l.svcCtx.ReserveStock(in.GetProductId(), in.GetQuantity())
	if !ok {
		return &productpb.ReserveStockResponse{Success: false, Reason: reason}, nil
	}
	return &productpb.ReserveStockResponse{Success: true}, nil
}
