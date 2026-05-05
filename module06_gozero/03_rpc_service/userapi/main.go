// API 调用 RPC 的客户端侧代码（通常在 API 服务中）
package main

import "fmt"

func main() {
	fmt.Println("=== API 调用 RPC - 客户端视角 ===")
	fmt.Println()
	fmt.Println("在 go-zero 微服务拆分中：")
	fmt.Println()
	fmt.Println("  用户请求:")
	fmt.Println("   Browser → API Gateway → user-api (HTTP)")
	fmt.Println("                              │")
	fmt.Println("                              └─→ user-rpc (gRPC + Etcd)")
	fmt.Println("                                     │")
	fmt.Println("                                     └─→ MySQL")
	fmt.Println()
	fmt.Println("  API 服务职责：")
	fmt.Println("  1. 接收 HTTP 请求，参数校验")
	fmt.Println("  2. 调用 RPC 服务完成业务")
	fmt.Println("  3. 组装 HTTP 响应返回")
	fmt.Println()
	fmt.Println("  RPC 服务职责：")
	fmt.Println("  1. 数据访问（CRUD）")
	fmt.Println("  2. 核心业务逻辑")
	fmt.Println("  3. 缓存储存管理")
	fmt.Println()
	fmt.Println("  分层的好处：")
	fmt.Println("  - API 层可以水平扩展，不涉及数据库")
	fmt.Println("  - RPC 层可以独立扩缩容")
	fmt.Println("  - 各服务独立部署，互不阻塞")
	fmt.Println()
	fmt.Println("配置文件 etc/user-api.yaml：")
	fmt.Println("  UserRpc:")
	fmt.Println("    Etcd:")
	fmt.Println("      Hosts:")
	fmt.Println("        - localhost:2379")
	fmt.Println("      Key: user.rpc    # 通过此 Key 在 Etcd 中找到 RPC 服务地址")

	_ = fmt.Sprint
}
