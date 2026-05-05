// 电商项目 - User RPC 服务（用户服务）
//
// 职责：用户信息查询、校验
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== 电商项目：User RPC 服务 ===")
	fmt.Println()

	fmt.Println("--- 服务职责 ---")
	fmt.Println("  1. 提供用户信息查询（GetUser）")
	fmt.Println("  2. 用户状态校验（是否激活、是否被禁用）")
	fmt.Println("  3. 被 Order API 调用：校验下单用户是否合法")
	fmt.Println()

	fmt.Println("--- .proto 定义 ---")
	fmt.Println(`  service UserRpc {
      rpc GetUser(GetUserRequest) returns (GetUserResponse);
      rpc CheckUserStatus(CheckUserStatusRequest) returns (CheckUserStatusResponse);
  }`)
	fmt.Println()

	fmt.Println("--- Logic 层实现要点 ---")
	fmt.Println(`  func (l *GetUserLogic) GetUser(in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
      user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
      if err != nil {
          if err == model.ErrNotFound {
              return nil, status.Error(codes.NotFound, "user not found")
          }
          return nil, status.Error(codes.Internal, err.Error())
      }

      return &userpb.GetUserResponse{
          User: &userpb.User{
              Id:       user.Id,
              Username: user.Username,
              Email:    user.Email,
              Status:   int32(user.Status),
          },
      }, nil
  }`)
	fmt.Println()

	fmt.Println("--- 微服务调用关系 ---")
	fmt.Println()
	fmt.Println("  ┌─────────────────────────────────────────────┐")
	fmt.Println("  │  用户请求 POST /api/v1/orders               │")
	fmt.Println("  │       │                                     │")
	fmt.Println("  │       ▼                                     │")
	fmt.Println("  │  [Order API]  (:8888 HTTP)                  │")
	fmt.Println("  │       │              │                      │")
	fmt.Println("  │       ▼              ▼                      │")
	fmt.Println("  │  [Order RPC]     [User RPC]                  │")
	fmt.Println("  │  (:9090 gRPC)    (:9091 gRPC)               │")
	fmt.Println("  │       │              │                      │")
	fmt.Println("  │       ▼              ▼                      │")
	fmt.Println("  │    MySQL           MySQL                    │")
	fmt.Println("  │  (orders 表)     (users 表)                  │")
	fmt.Println("  │                                              │")
	fmt.Println("  │  注册中心: Etcd (:2379)                      │")
	fmt.Println("  │  缓存:     Redis (:6379)                     │")
	fmt.Println("  │  消息队列: Kafka (:9092)                     │")
	fmt.Println("  └─────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("--- 电商系统服务拆分总结 ---")
	fmt.Println("  Order API   → 对外 HTTP 接口，编排调用")
	fmt.Println("  Order RPC   → 订单核心业务 + 数据持久化")
	fmt.Println("  User RPC    → 用户信息查询")
	fmt.Println("  Product RPC → 商品信息查询（可扩展）")
	fmt.Println("  Inventory RPC → 库存管理（可扩展）")

	_ = fmt.Sprint
}
