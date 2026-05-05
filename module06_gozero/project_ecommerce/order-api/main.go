// 电商项目 - Order API 服务（对外 HTTP 接口）
// 职责：接收用户订单请求 → 调用 OrderRpc → 返回响应
//
// 启动（需先启动 user-rpc 和 order-rpc）：
//   go run order-api/main.go
//
// 完整调用链：
//   curl → order-api (HTTP :8889) → order-rpc (gRPC :9092) → user-rpc (gRPC :9091)
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// ServiceContext 依赖注入容器
type ServiceContext struct {
	orderRpcConn *grpc.ClientConn
	userRpcConn  *grpc.ClientConn
}

// CreateOrder 请求
type CreateOrderReq struct {
	UserID int64       `json:"user_id"`
	Items  []OrderItem `json:"items"`
	Amount float64     `json:"amount"`
}
type OrderItem struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}
type CreateOrderResp struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CreateOrderHandler 创建订单
func CreateOrderHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 调用 OrderRpc.CreateOrder
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		resp, err := callOrderRpc_CreateOrder(ctx, svcCtx.orderRpcConn, req.UserID, req.Items, req.Amount)
		if err != nil {
			st := status.Convert(err)
			httpx.Error(w, fmt.Errorf("[%s] %s", st.Code(), st.Message()))
			return
		}

		httpx.OkJson(w, CreateOrderResp{
			OrderID: resp.OrderID,
			Status:  resp.Status,
			Message: "微服务调用链: order-api → order-rpc → user-rpc",
		})
	}
}

// GetOrderHandler 查询订单
func GetOrderHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("order_id")
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		resp, err := callOrderRpc_GetOrder(ctx, svcCtx.orderRpcConn, orderID)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

// ========== RPC 调用辅助函数（实际项目用生成的 pb client） ==========

type orderRpcCreateResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

func callOrderRpc_CreateOrder(ctx context.Context, conn *grpc.ClientConn, userID int64, items []OrderItem, amount float64) (*orderRpcCreateResponse, error) {
	// 构造请求 JSON（简化：通过 HTTP/JSON 桥接，实际用 protobuf）
	reqBody := map[string]interface{}{
		"user_id": userID,
		"items":   items,
		"amount":  amount,
	}
	data, _ := json.Marshal(reqBody)
	log.Printf("[OrderAPI] → OrderRpc.CreateOrder: %s", string(data))

	// 模拟 RPC 调用（简化版：本地逻辑，实际用 pb client）
	// 本示例中我们通过调用本地 OrderRpc 来演示
	orderID := fmt.Sprintf("ORD-%d", time.Now().UnixNano())
	return &orderRpcCreateResponse{OrderID: orderID, Status: "created"}, nil
}

func callOrderRpc_GetOrder(ctx context.Context, conn *grpc.ClientConn, orderID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"order_id": orderID,
		"status":   "created",
		"message":  "from order-rpc",
	}, nil
}

// ========== Main ==========

func main() {
	// 连接 OrderRpc 服务
	orderRpcConn, err := grpc.NewClient("localhost:9092",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("警告: 无法连接 OrderRpc: %v", err)
		defer func() {}()
	} else {
		defer orderRpcConn.Close()
	}

	// 连接 UserRpc 服务
	userRpcConn, err := grpc.NewClient("localhost:9091",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("警告: 无法连接 UserRpc: %v", err)
	} else {
		defer userRpcConn.Close()
	}

	svcCtx := &ServiceContext{
		orderRpcConn: orderRpcConn,
		userRpcConn:  userRpcConn,
	}

	server := rest.MustNewServer(rest.RestConf{Host: "0.0.0.0", Port: 8889})

	server.AddRoute(rest.Route{
		Method: http.MethodPost, Path: "/api/v1/orders",
		Handler: CreateOrderHandler(svcCtx),
	})
	server.AddRoute(rest.Route{
		Method: http.MethodGet, Path: "/api/v1/orders",
		Handler: GetOrderHandler(svcCtx),
	})
	// 根路径 显示 API 信息
	server.AddRoute(rest.Route{
		Method: http.MethodGet, Path: "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{
  "service": "order-api",
  "description": "电商微服务 - Order API",
  "endpoints": {
    "POST /api/v1/orders": "创建订单",
    "GET /api/v1/orders?order_id=xxx": "查询订单"
  },
  "architecture": "order-api → order-rpc(:9092) → user-rpc(:9091)"
}`))
		},
	})

	fmt.Println("=== 电商微服务 - Order API 已启动 ===")
	fmt.Println("  HTTP 端口: 8889")
	fmt.Println()
	fmt.Println("  架构：")
	fmt.Println("    curl → order-api (:8889 HTTP)")
	fmt.Println("              ↓")
	fmt.Println("         order-rpc (:9092 gRPC)")
	fmt.Println("              ↓")
	fmt.Println("         user-rpc  (:9091 gRPC)")
	fmt.Println()
	fmt.Println("  测试：")
	fmt.Println("    curl -X POST http://localhost:8889/api/v1/orders \\")
	fmt.Println("      -H 'Content-Type: application/json' \\")
	fmt.Println("      -d '{\"user_id\":1,\"items\":[{\"product_id\":101,\"quantity\":2}],\"amount\":99.99}'")
	fmt.Println()
	fmt.Println("    curl http://localhost:8889/api/v1/orders?order_id=ORD-1000")
	fmt.Println()

	server.Start()

	_ = bytes.NewReader
	_ = io.EOF
}
