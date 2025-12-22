package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}

func main() {
	viper.SetDefault("app.port", 8082)
	viper.SetDefault("app.name", "user-service")

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var req RegisterReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"ok": true, "user": req.Username, "app": viper.GetString("app.name")})
	})
	r.Run(":8082")
}
