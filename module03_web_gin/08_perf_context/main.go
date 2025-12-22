package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

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
