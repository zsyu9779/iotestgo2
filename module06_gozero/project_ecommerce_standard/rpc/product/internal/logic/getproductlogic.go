package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/product/productpb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	product, ok := l.svcCtx.GetProduct(in.GetProductId())
	if !ok {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	return product, nil
}
