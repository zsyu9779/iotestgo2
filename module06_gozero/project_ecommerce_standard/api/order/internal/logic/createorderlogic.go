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

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	items := make([]*orderclient.OrderItem, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, &orderclient.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	orderResp, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &orderclient.CreateOrderRequest{
		UserId: req.UserId,
		Items:  items,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateOrderResp{
		OrderId:     orderResp.GetOrderId(),
		Status:      orderResp.GetStatus(),
		TotalAmount: orderResp.GetTotalAmount(),
	}, nil
}
