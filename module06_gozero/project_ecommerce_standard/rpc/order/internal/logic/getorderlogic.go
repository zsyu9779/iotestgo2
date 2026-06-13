package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/orderpb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	order, ok := l.svcCtx.GetOrder(in.GetOrderId())
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &orderpb.GetOrderResponse{
		OrderId:     order.OrderID,
		UserId:      order.UserID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
	}, nil
}
