package middleware

import (
	"iotestgo/module03_web_gin/project_user_center/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
	sugar := logger.Sugar()
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		sugar.Infow("req",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"ms", time.Since(start).Milliseconds(),
		)
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		// Support "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 {
			authHeader = parts[1]
		}

		token, err := utils.ParseToken(authHeader)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		
		// Set username in context if needed
		// claims := token.Claims.(jwt.MapClaims)
		// c.Set("username", claims["sub"])
		
		c.Next()
	}
}
