// 电商项目 - Order API 服务（对外 HTTP 接口）
//
// 职责：接收用户订单请求，调用 Order RPC 服务创建订单
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== 电商项目：Order API 服务 ===")
	fmt.Println()

	fmt.Println("架构角色：")
	fmt.Println("  用户 → [Order API] → [Order RPC] → MySQL")
	fmt.Println("                    ↘ [User RPC]  → MySQL")
	fmt.Println()

	fmt.Println("--- .api 定义（order.api 片段）---")
	fmt.Println(`  type CreateOrderRequest {
      UserId  int64   ` + "`" + `json:"user_id"` + "`" + `
      Items   []OrderItem ` + "`" + `json:"items"` + "`" + `
      Amount  float64 ` + "`" + `json:"amount"` + "`" + `
  }

  type OrderItem {
      ProductId int64  ` + "`" + `json:"product_id"` + "`" + `
      Quantity  int32  ` + "`" + `json:"quantity"` + "`" + `
  }

  @server(
      prefix: /api/v1
      group:  order
      jwt:    JwtAuth
  )
  service order-api {
      @handler CreateOrder
      post /orders (CreateOrderRequest) returns (CreateOrderResponse)

      @handler GetOrder
      get /orders/:id (GetOrderRequest) returns (GetOrderResponse)
  }`)
	fmt.Println()

	fmt.Println("--- Logic 层核心流程 ---")
	fmt.Println(`  func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (*types.CreateOrderResponse, error) {
      // 1. 调用 UserRpc 校验用户是否存在
      userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userpb.GetUserRequest{UserId: req.UserId})
      if err != nil {
          return nil, status.Error(codes.NotFound, "user not found")
      }

      // 2. 调用 OrderRpc 创建订单
      orderResp, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &orderpb.CreateOrderRequest{
          UserId: req.UserId,
          Items:  convertItems(req.Items),
          Amount: req.Amount,
      })
      if err != nil {
          return nil, err
      }

      // 3. 返回响应
      return &types.CreateOrderResponse{
          OrderId: orderResp.OrderId,
          Status:  "created",
      }, nil
  }`)
	fmt.Println()

	fmt.Println("--- ServiceContext 依赖注入 ---")
	fmt.Println(`  type ServiceContext struct {
      Config   config.Config
      OrderRpc orderpb.OrderRpcClient   // 调用 Order RPC
      UserRpc  userpb.UserRpcClient     // 调用 User RPC（校验用户）
  }`)
	fmt.Println()

	fmt.Println("微服务间的调用链：")
	fmt.Println("  Order API  → [通过 Etcd 发现] → Order RPC → MySQL (订单表)")
	fmt.Println("  Order API  → [通过 Etcd 发现] → User RPC  → MySQL (用户表)")
	fmt.Println()
	fmt.Println("  注意：Order API 不直接访问数据库！所有数据操作通过 RPC 完成。")

	_ = fmt.Sprint
}
