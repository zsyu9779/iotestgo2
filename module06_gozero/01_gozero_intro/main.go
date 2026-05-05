// 01 go-zero 全景：架构设计理念、goctl 安装与项目初始化
package main

import "fmt"

func main() {
	fmt.Println("=== 01 go-zero 架构全景 ===")
	fmt.Println()

	fmt.Println("--- go-zero 核心理念 ---")
	fmt.Println("1. 极简依赖：不依赖 Spring Cloud 那样的庞大生态")
	fmt.Println("2. 代码生成优先：goctl 生成 80% 模板代码，聚焦 20% 业务逻辑")
	fmt.Println("3. 并发安全：内置熔断、限流、降级，保护微服务雪崩")
	fmt.Println("4. 可观测性：默认集成 Prometheus + 链路追踪")
	fmt.Println()

	fmt.Println("--- 架构分层 ---")
	fmt.Println("  ┌─────────────────────────────────────┐")
	fmt.Println("  │            API Gateway              │  ← 网关层（路由、限流、鉴权）")
	fmt.Println("  ├─────────────────────────────────────┤")
	fmt.Println("  │          API Service (HTTP)         │  ← 对外接口层（.api 定义）")
	fmt.Println("  ├─────────────────────────────────────┤")
	fmt.Println("  │          RPC Service (gRPC)         │  ← 内部服务层（.proto 定义）")
	fmt.Println("  ├──────────┬──────────┬───────────────┤")
	fmt.Println("  │  MySQL   │  Redis   │     Etcd       │  ← 数据与基础设施层")
	fmt.Println("  └──────────┴──────────┴───────────────┘")
	fmt.Println()

	fmt.Println("--- goctl 安装 ---")
	fmt.Println("  go install github.com/zeromicro/go-zero/tools/goctl@latest")
	fmt.Println()

	fmt.Println("--- goctl 常用命令 ---")
	fmt.Println("  goctl api new <name>          # 创建 API 项目")
	fmt.Println("  goctl api go -api <file>.api  # 从 .api 文件生成 Go 代码")
	fmt.Println("  goctl rpc new <name>          # 创建 RPC 项目")
	fmt.Println("  goctl rpc protoc <file>.proto # 从 .proto 生成 RPC 代码")
	fmt.Println("  goctl model mysql ddl -src <sql> -dir <dir> # 从 DDL 生成 Model")
	fmt.Println("  goctl docker -go <go-file>    # 生成 Dockerfile")
	fmt.Println("  goctl kube deploy -name <name> # 生成 K8s 部署 YAML")
	fmt.Println()

	fmt.Println("--- 项目结构（goctl 生成） ---")
	fmt.Println("  service/")
	fmt.Println("  ├── etc/")
	fmt.Println("  │   └── service.yaml       # 配置文件")
	fmt.Println("  ├── internal/")
	fmt.Println("  │   ├── config/            # 配置结构体")
	fmt.Println("  │   ├── handler/           # HTTP Handler（路由层）")
	fmt.Println("  │   ├── logic/             # 业务逻辑（核心）")
	fmt.Println("  │   ├── svc/               # ServiceContext（依赖注入容器）")
	fmt.Println("  │   └── types/             # 请求/响应类型")
	fmt.Println("  └── service.go             # 入口文件")

	_ = fmt.Sprint
}
