package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		sugar.Infow("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})

	r.GET("/work", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 200*time.Millisecond)
		defer cancel()
		select {
		case <-time.After(100 * time.Millisecond):
			c.JSON(200, gin.H{"ok": true})
		case <-ctx.Done():
			c.JSON(504, gin.H{"error": "timeout"})
		}
	})

	r.Run(":8084")
}
