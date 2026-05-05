// 08 K8s 部署入门：可部署的健康检查 HTTP 服务 + 部署说明
//
// 启动：go run main.go
//
// Docker 构建与 K8s 部署步骤（见输出）
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 简单的 HTTP Server
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ready"}`))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"service":"user-api","version":"1.0.0"}`))
	})

	// 优雅退出
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		fmt.Println("\nShutting down...")
		os.Exit(0)
	}()

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Println("=== K8s 部署示例服务 ===")
	fmt.Println("  端口:", port)
	fmt.Println()
	fmt.Println("  健康检查:")
	fmt.Println("    curl http://localhost:" + port + "/health")
	fmt.Println("    curl http://localhost:" + port + "/ready")
	fmt.Println()
	fmt.Println("--- 容器化部署流程 ---")
	fmt.Println()
	fmt.Println("1. 构建 Docker 镜像（使用已有 Dockerfile）：")
	fmt.Println("   docker build -t user-api:latest .")
	fmt.Println()
	fmt.Println("2. 本地测试容器：")
	fmt.Println("   docker run -p 8080:8080 user-api:latest")
	fmt.Println()
	fmt.Println("3. 推送镜像到仓库：")
	fmt.Println("   docker tag user-api:latest myrepo/user-api:latest")
	fmt.Println("   docker push myrepo/user-api:latest")
	fmt.Println()
	fmt.Println("4. 部署到 K8s：")
	fmt.Println("   kubectl apply -f deployment.yaml")
	fmt.Println("   kubectl get pods")
	fmt.Println("   kubectl logs -f <pod-name>")
	fmt.Println()
	fmt.Println("5. 扩缩容：")
	fmt.Println("   kubectl scale deployment user-api --replicas=5")
	fmt.Println()
	fmt.Println("--- K8s 核心概念 ---")
	fmt.Println("  Deployment:  管理 Pod 副本、滚动更新、回滚")
	fmt.Println("  Service:     Pod 的负载均衡器（ClusterIP/NodePort/LoadBalancer）")
	fmt.Println("  ConfigMap:   非敏感配置（环境变量、配置文件）")
	fmt.Println("  Secret:      敏感数据（密码、证书、Token）")
	fmt.Println("  Ingress:     HTTP 路由规则 → Service")
	fmt.Println()
	fmt.Println("--- Go 在 K8s 中的优势 ---")
	fmt.Println("  1. 单二进制 → 镜像体积 ~10MB（Java 200MB+）")
	fmt.Println("  2. 毫秒级启动（JVM 需要秒级预热）")
	fmt.Println("  3. 内存占用 ~50MB（JVM 256MB+）")
	fmt.Println("  4. 无需 JVM → Pod 资源请求更少 → 更高密度部署")

	http.ListenAndServe(":"+port, nil)
}
