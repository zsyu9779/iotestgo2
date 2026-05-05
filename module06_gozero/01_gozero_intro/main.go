// 01 go-zero 架构全景：用 go-zero rest 包构建最小 HTTP 服务
//
// 启动：go run main.go
// 测试：curl http://localhost:8888/ping
//       curl http://localhost:8888/hello?name=Gopher
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	// 1. 创建 go-zero HTTP Server（无需 goctl，直接使用 rest 包）
	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8888,
	})

	// 2. 注册路由
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ping",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message":"pong","framework":"go-zero"}`))
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/hello",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "Gopher"
			}
			logx.WithContext(r.Context()).Infof("received hello request for: %s", name)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"message":"Hello, %s!","framework":"go-zero"}`, name)))
		},
	})

	// 3. 添加 go-zero 内置中间件（自动记录请求日志、指标、超时控制）
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Infof("[%s] %s", r.Method, r.URL.Path)
			next(w, r)
		}
	})

	// 4. 优雅退出
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logx.Info("Shutting down...")
		server.Stop()
	}()

	fmt.Println("=== go-zero HTTP Server 已启动 ===")
	fmt.Println("  架构：API Gateway → API Service(HTTP) → RPC Service(gRPC) → MySQL/Redis/Etcd")
	fmt.Println()
	fmt.Println("  测试：")
	fmt.Println("    curl http://localhost:8888/ping")
	fmt.Println("    curl http://localhost:8888/hello?name=Gopher")
	fmt.Println()
	fmt.Println("  go-zero 核心概念：")
	fmt.Println("    1. .api 文件定义 HTTP 服务 → goctl 自动生成代码")
	fmt.Println("    2. .proto 文件定义 RPC 服务 → goctl 生成 zRPC 代码")
	fmt.Println("    3. Etcd 做服务注册与发现（替代 Eureka/Nacos）")
	fmt.Println("    4. 内置熔断(breaker)、限流(limiter)、超时控制")
	fmt.Println("    5. 默认集成 Prometheus 指标暴露")
	fmt.Println()

	server.Start()
}
