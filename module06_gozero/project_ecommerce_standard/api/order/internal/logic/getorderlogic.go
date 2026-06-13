// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/api/order/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/api/order/internal/types"
	orderclient "iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.GetOrderReq) (resp *types.GetOrderResp, err error) {
	orderResp, err := l.svcCtx.OrderRpc.GetOrder(l.ctx, &orderclient.GetOrderRequest{OrderId: req.OrderId})
	if err != nil {
		return nil, err
	}

	return &types.GetOrderResp{
		OrderId:     orderResp.GetOrderId(),
		UserId:      orderResp.GetUserId(),
		Status:      orderResp.GetStatus(),
		TotalAmount: orderResp.GetTotalAmount(),
	}, nil
}
