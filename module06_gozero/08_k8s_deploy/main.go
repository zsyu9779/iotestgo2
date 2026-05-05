// 08 K8s 部署入门：Dockerfile 编写 + K8s Deployment/Service
package main

import "fmt"

func main() {
	fmt.Println("=== 08 K8s 部署入门 ===")
	fmt.Println()

	fmt.Println("--- goctl 生成部署文件 ---")
	fmt.Println("  goctl docker -go user-api.go          # 生成 Dockerfile")
	fmt.Println("  goctl kube deploy -name user-api \\")
	fmt.Println("       -namespace default \\")
	fmt.Println("       -image user-api:latest \\")
	fmt.Println("       -o deployment.yaml              # 生成 K8s 部署清单")
	fmt.Println()

	fmt.Println("--- 容器化流程 ---")
	fmt.Println("  1. goctl docker → 生成多阶段构建 Dockerfile")
	fmt.Println("  2. docker build -t user-api:latest .")
	fmt.Println("  3. docker push registry.example.com/user-api:latest")
	fmt.Println()

	fmt.Println("--- K8s 部署流程 ---")
	fmt.Println("  1. goctl kube deploy → 生成 deployment.yaml")
	fmt.Println("  2. kubectl apply -f deployment.yaml")
	fmt.Println("  3. kubectl get pods    # 查看 Pod 运行状态")
	fmt.Println("  4. kubectl logs -f <pod-name>  # 查看日志")
	fmt.Println("  5. kubectl scale deployment user-api --replicas=5  # 扩缩容")
	fmt.Println()

	fmt.Println("--- K8s 核心概念与 Java 对比 ---")
	fmt.Println("  Deployment:  管理 Pod 副本数、滚动更新（类似 Spring Cloud + K8s Operator）")
	fmt.Println("  Service:     负载均衡到 Pod（类似 Spring Cloud LoadBalancer）")
	fmt.Println("  ConfigMap:   配置管理（类似 Spring Cloud Config）")
	fmt.Println("  Secret:      密钥管理（类似 Vault / 加密的 Config）")
	fmt.Println("  Ingress:     HTTP 路由（类似 Spring Cloud Gateway）")
	fmt.Println()

	fmt.Println("--- Go 应用在 K8s 中的优势 ---")
	fmt.Println("  1. 单二进制文件 → 镜像体积小（~10MB vs Java 200MB+）")
	fmt.Println("  2. 启动快（毫秒级 vs JVM 秒级）")
	fmt.Println("  3. 内存占用低（~50MB vs JVM 256MB+）")
	fmt.Println("  4. 无需单独 JVM → 更少的 Pod 资源请求")
	fmt.Println("  5. 天然适合 K8s 的不可变基础设施理念")
}
