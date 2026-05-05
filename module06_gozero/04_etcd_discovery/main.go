// 04 Etcd 服务注册与发现
//
// 启动 Etcd（如果本地没有运行）：
//   docker run -d --name etcd-demo -p 2379:2379 quay.io/coreos/etcd:v3.5.9
//   --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://localhost:2379
//
// 启动本程序：go run main.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func main() {
	fmt.Println("=== Etcd 服务注册与发现 ===")
	fmt.Println()

	// 1. 连接 Etcd
	fmt.Println("连接 Etcd (localhost:2379)...")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		fmt.Printf("无法连接 Etcd: %v\n", err)
		fmt.Println()
		fmt.Println("请先启动 Etcd：")
		fmt.Println("  docker run -d --name etcd-demo -p 2379:2379 \\")
		fmt.Println("    quay.io/coreos/etcd:v3.5.9 --listen-client-urls http://0.0.0.0:2379 \\")
		fmt.Println("    --advertise-client-urls http://localhost:2379")
		fmt.Println()
		fmt.Println("本程序将以演示模式运行，展示 Etcd 的核心概念。")
		showDemoMode()
		return
	}
	defer cli.Close()

	ctx := context.Background()
	fmt.Println("✓ 已连接 Etcd")
	fmt.Println()

	// 2. Put：写入 KV
	fmt.Println("--- 1. Put（服务注册）---")
	key := "/services/user.rpc/instance-1"
	value := "192.168.1.100:9091"
	_, err = cli.Put(ctx, key, value)
	if err != nil {
		log.Printf("Put 失败: %v", err)
	} else {
		fmt.Printf("  注册服务: %s = %s\n", key, value)
	}

	// 3. Get：读取 KV
	fmt.Println()
	fmt.Println("--- 2. Get（服务发现）---")
	resp, err := cli.Get(ctx, "/services/user.rpc/", clientv3.WithPrefix())
	if err != nil {
		log.Printf("Get 失败: %v", err)
	} else {
		fmt.Printf("  发现 %d 个 user.rpc 实例:\n", len(resp.Kvs))
		for _, kv := range resp.Kvs {
			fmt.Printf("    %s = %s\n", kv.Key, kv.Value)
		}
	}

	// 4. Lease + KeepAlive（服务自动续约）
	fmt.Println()
	fmt.Println("--- 3. Lease + KeepAlive（健康检查）---")
	lease, err := cli.Grant(ctx, 5) // 5 秒 TTL
	if err != nil {
		log.Printf("Grant 失败: %v", err)
	} else {
		fmt.Printf("  创建 Lease: ID=%d, TTL=5s\n", lease.ID)

		// 绑定 Lease 到 Key
		leaseKey := "/services/user.rpc/instance-2"
		_, err = cli.Put(ctx, leaseKey, "192.168.1.101:9091", clientv3.WithLease(lease.ID))
		if err != nil {
			log.Printf("Put with lease 失败: %v", err)
		} else {
			fmt.Printf("  注册带 Lease 的服务: %s = 192.168.1.101:9091\n", leaseKey)

			// KeepAlive 续约
			keepCh, err := cli.KeepAlive(ctx, lease.ID)
			if err != nil {
				log.Printf("KeepAlive 失败: %v", err)
			} else {
				go func() {
					for ka := range keepCh {
						if ka != nil {
							fmt.Printf("  Lease %d 续约成功 (TTL=%d)\n", ka.ID, ka.TTL)
						}
					}
				}()
			}
		}
	}

	// 5. Watch（监听变更）
	fmt.Println()
	fmt.Println("--- 4. Watch（服务上下线监听）---")
	watchCh := cli.Watch(ctx, "/services/user.rpc/", clientv3.WithPrefix())
	go func() {
		fmt.Println("  开始监听 /services/user.rpc/ 变更...")
		for wr := range watchCh {
			for _, ev := range wr.Events {
				fmt.Printf("  [Watch] %s %s = %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()

	// 模拟一次变更
	time.Sleep(500 * time.Millisecond)
	cli.Put(ctx, "/services/user.rpc/instance-3", "192.168.1.102:9091")
	time.Sleep(500 * time.Millisecond)
	cli.Delete(ctx, "/services/user.rpc/instance-3")

	// 6. 清理
	fmt.Println()
	fmt.Println("--- 5. 清理 ---")
	cli.Delete(ctx, "/services/user.rpc/", clientv3.WithPrefix())
	fmt.Println("  已删除所有测试数据")

	fmt.Println()
	fmt.Println("=== go-zero 中的 Etcd ===")
	fmt.Println("  zRPC 服务启动时自动通过 Etcd Lease + KeepAlive 注册")
	fmt.Println("  zRPC 客户端通过 Etcd Watch 自动发现服务实例")
	fmt.Println("  内置负载均衡：Round Robin / Weighted Random")
	fmt.Println()
	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Eureka: AP（可用性优先）, 自我保护模式")
	fmt.Println("  Nacos:  AP + CP 可切换")
	fmt.Println("  Etcd:   CP（一致性优先）, Raft 算法, 强一致性")
	time.Sleep(1 * time.Second)
}

func showDemoMode() {
	fmt.Println("=== Etcd 概念演示（无 Etcd 连接） ===")
	fmt.Println()
	fmt.Println("Etcd 是一个分布式 KV 存储，基于 Raft 共识算法。")
	fmt.Println()
	fmt.Println("核心概念：")
	fmt.Println("  1. Key-Value: 存储服务地址，如 /services/user.rpc/instance-1 = 192.168.1.100:9091")
	fmt.Println("  2. Lease: 租约 + TTL，服务停止后 Key 自动过期删除（默认 5 秒）")
	fmt.Println("  3. KeepAlive: 定期续约，服务存活期间不断刷新 Lease")
	fmt.Println("  4. Watch: 监听 Key 变更，服务上下线实时通知")
	fmt.Println("  5. 强一致性: Raft 算法，保证集群数据一致")
	fmt.Println()
	fmt.Println("go-zero 集成方式：")
	fmt.Println("  - zRPC 服务通过配置文件指定 Etcd 地址 → 启动时自动注册")
	fmt.Println("  - zRPC 客户端通过 Etcd Key 发现服务 → 自动负载均衡")
}
