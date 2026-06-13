package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/userpb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, ok := l.svcCtx.GetUser(in.GetUserId())
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return user, nil
}
