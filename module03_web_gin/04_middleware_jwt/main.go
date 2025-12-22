package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret-key-demo")

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}
		tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !tkn.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		token, _ := generateToken("demo-user")
		c.JSON(200, gin.H{"token": token})
	})

	auth := r.Group("/secure").Use(AuthMiddleware())
	auth.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{"user": "demo-user"})
	})

	r.Run(":8083")
}
