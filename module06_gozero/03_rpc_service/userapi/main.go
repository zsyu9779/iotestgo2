// 03 RPC 服务 - UserApi 客户端（HTTP API 调用 gRPC RPC）
//
// 启动（先启动 userrpc）：go run userapi/main.go
//
// 本示例演示第一个微服务拆分：
//   用户 → userapi (HTTP :8882) → userrpc (gRPC :9091) → 数据
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "iotestgo/module06_gozero/03_rpc_service/userpb"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// ServiceContext（装载 RPC 客户端等依赖）
type ServiceContext struct {
	UserRpc pb.UserRpcClient
}

// Handler
func UserInfoHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Query().Get("user_id")
		userID, _ := strconv.ParseInt(userIDStr, 10, 64)

		// 调用 UserRpc 服务（gRPC）
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		userResp, err := svcCtx.UserRpc.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
		if err != nil {
			st := status.Convert(err)
			httpx.Error(w, fmt.Errorf("[%s] %s", st.Code(), st.Message()))
			return
		}

		user := userResp.GetUser()
		httpx.OkJson(w, map[string]interface{}{
			"user_id":  user.GetId(),
			"username": user.GetUsername(),
			"email":    user.GetEmail(),
			"status":   user.GetStatus(),
			"source":   "from UserRpc (gRPC service)",
		})
	}
}

func main() {
	// 依赖注入：创建 RPC 客户端连接
	conn, err := grpc.NewClient("localhost:9091",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接 RPC 服务失败: %v", err)
	}
	defer conn.Close()

	svcCtx := &ServiceContext{
		UserRpc: pb.NewUserRpcClient(conn),
	}

	server := rest.MustNewServer(rest.RestConf{Host: "0.0.0.0", Port: 8882})

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/api/v1/user/info",
		Handler: UserInfoHandler(svcCtx),
	})

	fmt.Println("=== UserApi 服务已启动（API 调用 RPC） ===")
	fmt.Println("  HTTP 端口: 8882")
	fmt.Println("  调用链: curl → userapi (HTTP) → userrpc (gRPC :9091)")
	fmt.Println()
	fmt.Println("  测试:")
	fmt.Println("    curl http://localhost:8882/api/v1/user/info?user_id=1")
	fmt.Println("    curl http://localhost:8882/api/v1/user/info?user_id=2")
	fmt.Println("    curl http://localhost:8882/api/v1/user/info?user_id=99  # 不存在")
	fmt.Println()
	fmt.Println("  go-zero 微服务拆分原则：")
	fmt.Println("    API 服务：接收 HTTP，参数校验，调用 RPC，组装响应")
	fmt.Println("    RPC 服务：数据访问，核心业务逻辑，独立部署")
	fmt.Println()

	server.Start()
}
