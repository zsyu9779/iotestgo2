// 07 可观测性：Prometheus 指标暴露 + 结构化日志
//
// 启动：go run main.go
// 测试：
//   curl http://localhost:8883/ping
//   curl http://localhost:8883/slow
//   curl http://localhost:8883/error
//   curl http://localhost:8883/metrics
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/prometheus"
	"github.com/zeromicro/go-zero/rest"
)

var (
	// 自定义指标
	pingCounter   int64
	errorCounter  int64
	totalRequests int64
)

func main() {
	// 启动 Prometheus 指标端点（自动暴露在 /metrics）
	// go-zero 的 rest server 内置了 metrics handler

	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8883,
	})

	// 注册自定义 Prometheus 指标
	prometheus.StartAgent(prometheus.Config{
		Host: "0.0.0.0",
		Port: 9100,
		Path: "/metrics",
	})

	// /ping：正常请求
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ping",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&pingCounter, 1)
			atomic.AddInt64(&totalRequests, 1)
			logx.WithContext(r.Context()).Info("ping request received")
			w.Write([]byte(`{"status":"ok","service":"observability-demo"}`))
		},
	})

	// /slow：模拟慢查询
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/slow",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&totalRequests, 1)
			start := time.Now()
			delay := time.Duration(200+rand.Intn(800)) * time.Millisecond
			time.Sleep(delay)
			elapsed := time.Since(start)
			logx.WithContext(r.Context()).Infow("slow request completed",
				logx.Field("latency_ms", elapsed.Milliseconds()),
			)
			w.Write([]byte(fmt.Sprintf(`{"status":"ok","latency":"%v"}`, elapsed)))
		},
	})

	// /error：返回错误
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/error",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&errorCounter, 1)
			atomic.AddInt64(&totalRequests, 1)
			logx.WithContext(r.Context()).Error("simulated error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"internal server error"}`))
		},
	})

	// /stats：自定义统计
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/stats",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(fmt.Sprintf(`{
  "total_requests": %d,
  "ping_count": %d,
  "error_count": %d
}`, atomic.LoadInt64(&totalRequests),
				atomic.LoadInt64(&pingCounter),
				atomic.LoadInt64(&errorCounter))))
		},
	})

	fmt.Println("=== 可观测性演示 ===")
	fmt.Println("  HTTP 服务: :8883")
	fmt.Println("  Prometheus 指标: :9100/metrics")
	fmt.Println()
	fmt.Println("  测试：")
	fmt.Println("    curl http://localhost:8883/ping")
	fmt.Println("    curl http://localhost:8883/slow")
	fmt.Println("    curl http://localhost:8883/error")
	fmt.Println("    curl http://localhost:8883/stats")
	fmt.Println("    curl http://localhost:9100/metrics | head -20")
	fmt.Println()
	fmt.Println("  三大支柱：")
	fmt.Println("    1. Metrics（指标）: Prometheus → 请求数、延迟、错误率")
	fmt.Println("    2. Tracing（追踪）: Jaeger/OpenTelemetry → 分布式调用链")
	fmt.Println("    3. Logging（日志）: go-zero logx → 结构化日志 + trace_id 关联")
	fmt.Println()
	fmt.Println("  go-zero 默认集成：")
	fmt.Println("    - 自动暴露 /metrics 到 Prometheus")
	fmt.Println("    - 请求级别的结构化日志")
	fmt.Println("    - 内置熔断/限流状态的指标")
	fmt.Println()

	server.Start()
}
