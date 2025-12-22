package main

import (
	"iotestgo/module04_gorm/project_blog_api/internal/handler"
	"iotestgo/module04_gorm/project_blog_api/internal/model"
	"iotestgo/module04_gorm/project_blog_api/internal/repository"
	"iotestgo/module04_gorm/project_blog_api/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	db.AutoMigrate(&model.Post{}, &model.Comment{})
	return db
}

func main() {
	// 1. Setup DB
	db := setupDB()

	// 2. DI
	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(repo)
	h := handler.NewPostHandler(svc)

	// 3. Router
	r := gin.Default()

	r.GET("/posts", h.List)
	r.POST("/posts", h.Create)
	r.POST("/posts/with-comment", h.CreateWithComment) // Transaction example

	// 4. Run
	r.Run(":8091")
}
