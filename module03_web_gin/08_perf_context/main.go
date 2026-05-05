package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func gracefulShutdownExample() {
	// 优雅关机完整示例：signal.NotifyContext + srv.Shutdown()
	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8086",
		Handler: r,
	}

	// 创建监听信号的 context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 在 goroutine 中启动服务
	go func() {
		log.Println("Server starting on :8086")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号
	<-ctx.Done()
	stop()
	log.Println("Shutting down gracefully...")

	// 设置 5 秒超时来等待请求完成
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited")
}

func main() {
	r := gin.Default()
	r.GET("/slow", func(c *gin.Context) {
		select {
		case <-time.After(150 * time.Millisecond):
			c.JSON(200, gin.H{"done": true})
		case <-c.Request.Context().Done():
			c.JSON(499, gin.H{"error": "client canceled"})
		}
	})
	r.Run(":8086")
}
