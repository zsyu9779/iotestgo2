package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required,min=2"`
}

var users = []User{{ID: 1, Name: "Alice"}}

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) { c.JSON(200, users) })
		v1.POST("/users", func(c *gin.Context) {
			var u User
			if err := c.ShouldBindJSON(&u); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			u.ID = len(users) + 1
			users = append(users, u)
			c.JSON(201, u)
		})
	}
	r.Run(":8085")
}
