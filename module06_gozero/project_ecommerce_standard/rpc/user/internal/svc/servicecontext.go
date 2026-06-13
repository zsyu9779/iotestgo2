package svc

import (
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/internal/config"
	"iotestgo/module06_gozero/project_ecommerce_standard/rpc/user/userpb"
)

type ServiceContext struct {
	Config config.Config
	users  map[int64]*userpb.GetUserResponse
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		users: map[int64]*userpb.GetUserResponse{
			1: {UserId: 1, Username: "gopher", Email: "gopher@example.com", Status: 1},
			2: {UserId: 2, Username: "alice", Email: "alice@example.com", Status: 1},
			3: {UserId: 3, Username: "disabled", Email: "disabled@example.com", Status: 2},
		},
	}
}

func (s *ServiceContext) GetUser(userID int64) (*userpb.GetUserResponse, bool) {
	user, ok := s.users[userID]
	if !ok {
		return nil, false
	}
	return &userpb.GetUserResponse{
		UserId:   user.GetUserId(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		Status:   user.GetStatus(),
	}, true
}
