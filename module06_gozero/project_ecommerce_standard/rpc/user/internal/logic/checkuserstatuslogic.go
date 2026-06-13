package logic

import (
	"context"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/userpb"

	"github.com/zeromicro/go-zero/core/logx"
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
	users := map[int64]*userpb.GetUserResponse{
		1: {UserId: 1, Username: "gopher", Email: "gopher@example.com", Status: 1},
		2: {UserId: 2, Username: "alice", Email: "alice@example.com", Status: 1},
		3: {UserId: 3, Username: "disabled", Email: "disabled@example.com", Status: 2},
	}
	user, ok := users[in.GetUserId()]
	if !ok {
		return &userpb.CheckUserStatusResponse{Valid: false, Reason: "user not found"}, nil
	}
	if user.GetStatus() != 1 {
		return &userpb.CheckUserStatusResponse{Valid: false, Reason: "user disabled"}, nil
	}
	return &userpb.CheckUserStatusResponse{Valid: true}, nil
}
