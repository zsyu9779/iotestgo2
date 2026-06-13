package logic

import (
	"context"
	"fmt"
	"time"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/order/orderpb"
	productclient "iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/product"
	userclient "iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	userStatus, err := l.svcCtx.UserRpc.CheckUserStatus(l.ctx, &userclient.CheckUserStatusRequest{UserId: in.GetUserId()})
	if err != nil {
		return nil, err
	}
	if !userStatus.GetValid() {
		return nil, status.Error(codes.PermissionDenied, userStatus.GetReason())
	}

	var totalAmount float64
	for _, item := range in.GetItems() {
		productResp, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &productclient.GetProductRequest{ProductId: item.GetProductId()})
		if err != nil {
			return nil, err
		}

		reserveResp, err := l.svcCtx.ProductRpc.ReserveStock(l.ctx, &productclient.ReserveStockRequest{
			ProductId: item.GetProductId(),
			Quantity:  item.GetQuantity(),
		})
		if err != nil {
			return nil, err
		}
		if !reserveResp.GetSuccess() {
			return nil, status.Error(codes.FailedPrecondition, reserveResp.GetReason())
		}

		totalAmount += productResp.GetPrice() * float64(item.GetQuantity())
	}

	orderID := fmt.Sprintf("ORD-%d", time.Now().UnixNano())
	l.svcCtx.SaveOrder(&svc.StoredOrder{
		OrderID:     orderID,
		UserID:      in.GetUserId(),
		Status:      "created",
		TotalAmount: totalAmount,
	})

	return &orderpb.CreateOrderResponse{
		OrderId:     orderID,
		Status:      "created",
		TotalAmount: totalAmount,
	}, nil
}
