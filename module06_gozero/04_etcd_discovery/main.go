// 04 Etcd 服务注册与发现
//
// Etcd 是 go-zero 默认的服务注册中心（类似 Java 生态的 Eureka/Nacos）
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 04 Etcd 服务注册与发现 ===")
	fmt.Println()

	fmt.Println("--- Etcd 基础 ---")
	fmt.Println("  Etcd 是一个分布式 KV 存储，基于 Raft 共识算法")
	fmt.Println("  核心特性：强一致性、Watch 监听、Lease 租约、TTL 过期")
	fmt.Println()

	fmt.Println("--- go-zero 中的 Etcd ---")
	fmt.Println()
	fmt.Println("服务注册（RPC 服务启动时自动注册）：")
	fmt.Println(`  // etc/user-rpc.yaml 配置文件
  Name: user.rpc
  ListenOn: 0.0.0.0:9090
  Etcd:
    Hosts:
      - localhost:2379
    Key: user.rpc    // 注册到 Etcd 的 Key`)
	fmt.Println()
	fmt.Println("  go-zero 内部机制：")
	fmt.Println("  1. 服务启动 → 创建 Lease（租约，默认 5 秒 TTL）")
	fmt.Println("  2. 将服务地址写入 Etcd：/user.rpc/<instance-id> = 192.168.1.100:9090")
	fmt.Println("  3. 定期续约（KeepAlive）→ 保持 Key 存活")
	fmt.Println("  4. 服务停止 → Lease 过期 → Key 自动删除")
	fmt.Println()

	fmt.Println("服务发现（客户端自动发现）：")
	fmt.Println(`  // API 服务中配置 RPC Client
  UserRpc:
    Etcd:
      Hosts:
        - localhost:2379
      Key: user.rpc    // 监听此 Key 下的所有实例
      Timeout: 5000    // 连接超时 ms`)
	fmt.Println()
	fmt.Println("  go-zero 内部机制：")
	fmt.Println("  1. 客户端通过 Etcd Key 获取所有可用实例列表")
	fmt.Println("  2. Watch 机制：实例上下线实时更新本地列表")
	fmt.Println("  3. 内置负载均衡：Round Robin / Weighted Random")
	fmt.Println()

	fmt.Println("--- 观察服务上下线 ---")
	fmt.Println("  实验步骤：")
	fmt.Println("  1. docker-compose up -d etcd        # 启动 Etcd")
	fmt.Println("  2. 启动 user-rpc 实例 1 (端口 9090)")
	fmt.Println("  3. 启动 user-rpc 实例 2 (端口 9091)")
	fmt.Println("  4. 启动 user-api → 自动发现两个 RPC 实例")
	fmt.Println("  5. 关掉实例 1 → user-api 自动切换到实例 2（无感知）")
	fmt.Println()

	fmt.Println("--- Etcd 与 Java 对比 ---")
	fmt.Println("  Eureka: AP 系统（可用性优先），自我保护模式")
	fmt.Println("  Nacos:  AP + CP 可切换")
	fmt.Println("  Etcd:   CP 系统（一致性优先），Raft 算法")
	fmt.Println("  go-zero 选 Etcd 因为更轻量，一致性强，适合核心服务发现")

	_ = time.Now
}
