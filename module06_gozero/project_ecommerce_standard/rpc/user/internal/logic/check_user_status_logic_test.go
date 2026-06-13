package logic

import (
	"context"
	"testing"

	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/internal/config"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/internal/svc"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/userpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCheckUserStatusReturnsNotFoundForMissingUser(t *testing.T) {
	logic := NewCheckUserStatusLogic(context.Background(), svc.NewServiceContext(config.Config{}))

	_, err := logic.CheckUserStatus(&userpb.CheckUserStatusRequest{UserId: 404})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v from %v", status.Code(err), err)
	}
}
