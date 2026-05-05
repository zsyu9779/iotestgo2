// 07 可观测性：Prometheus 指标暴露 + Grafana 看板 + Jaeger 链路追踪
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== 07 可观测性（Observability） ===")
	fmt.Println()

	fmt.Println("三大支柱：Metrics（指标）+ Tracing（追踪）+ Logging（日志）")
	fmt.Println()

	fmt.Println("--- 1. Metrics（Prometheus） ---")
	fmt.Println("  go-zero 内置 Prometheus 指标暴露：")
	fmt.Println()
	fmt.Println("  配置 etc/user-api.yaml：")
	fmt.Println(`  Prometheus:
    Host: 0.0.0.0
    Port: 9091
    Path: /metrics`)
	fmt.Println()
	fmt.Println("  自动收集的指标：")
	fmt.Println("  - http_requests_total:         请求总数（按 method/path/status 分组）")
	fmt.Println("  - http_request_duration_ms:    请求耗时分布（P50/P90/P99）")
	fmt.Println("  - rpc_server_handled_total:    RPC 调用次数")
	fmt.Println("  - rpc_server_handled_seconds:  RPC 处理耗时")
	fmt.Println("  - breker_circuit_breaker:      熔断器状态")
	fmt.Println("  - limiter_dropped_total:       限流丢弃数量")
	fmt.Println()

	fmt.Println("--- 2. Tracing（Jaeger / OpenTelemetry） ---")
	fmt.Println("  链路追踪配置：")
	fmt.Println(`  Telemetry:
    Name: user-api
    Endpoint: http://jaeger:14268/api/traces
    Sampler: 1.0           # 采样率 100%（开发环境）
    Batcher: jaeger`)
	fmt.Println()
	fmt.Println("  追踪范围示例：")
	fmt.Println("  Browser → API Gateway → user-api → user-rpc → MySQL")
	fmt.Println("  每一步的耗时、状态码、错误信息都会被记录")
	fmt.Println()

	fmt.Println("--- 3. Logging（结构化日志 + 链路关联） ---")
	fmt.Println("  go-zero 使用 logx 日志库：")
	fmt.Println(`  import "github.com/zeromicro/go-zero/core/logx"

  // 带 trace_id 的日志（自动从 context 提取）
  logx.WithContext(ctx).Infof("user registered: userId=%d", userId)

  // 结构化日志：
  logx.Infow("order created",
      logx.Field("orderId", orderId),
      logx.Field("amount", amount),
      logx.Field("traceId", traceId),
  )`)
	fmt.Println()

	fmt.Println("--- Grafana 看板 ---")
	fmt.Println("  go-zero 提供预置 Grafana Dashboard JSON：")
	fmt.Println("  1. API 服务 QPS / 延迟 / 错误率")
	fmt.Println("  2. RPC 服务调用量 / 成功率")
	fmt.Println("  3. 数据库连接池状态")
	fmt.Println("  4. Redis 缓存命中率")
	fmt.Println("  5. 熔断器/限流器状态")
	fmt.Println("  6. JVM 指标 → Go 对应为 goroutine 数量、内存使用、GC 暂停")
	fmt.Println()

	fmt.Println("--- Docker-Compose 启动监控栈 ---")
	fmt.Println(`  prometheus:
    image: prom/prometheus
    ports: ["9090:9090"]
    volumes: ["./prometheus.yml:/etc/prometheus/prometheus.yml"]

  grafana:
    image: grafana/grafana
    ports: ["3000:3000"]
    environment: ["GF_SECURITY_ADMIN_PASSWORD=admin"]

  jaeger:
    image: jaegertracing/all-in-one
    ports: ["16686:16686", "14268:14268"]`)
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: Micrometer + Prometheus + Zipkin/SkyWalking")
	fmt.Println("  go-zero: 内置 metrics + OpenTelemetry + 预置 Grafana 模板")
	fmt.Println("  go-zero 的可观测性更开箱即用，零配置即可暴露指标")

	_ = fmt.Sprint
}
