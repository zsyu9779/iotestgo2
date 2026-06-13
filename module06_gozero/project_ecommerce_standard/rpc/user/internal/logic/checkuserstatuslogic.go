package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/userpb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserStatusLogic {
	return &CheckUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserStatusLogic) CheckUserStatus(in *userpb.CheckUserStatusRequest) (*userpb.CheckUserStatusResponse, error) {
	user, ok := l.svcCtx.GetUser(in.GetUserId())
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if user.GetStatus() != 1 {
		return &userpb.CheckUserStatusResponse{Valid: false, Reason: "user disabled"}, nil
	}
	return &userpb.CheckUserStatusResponse{Valid: true}, nil
}
