package main

import (
	"iotestgo/module03_web_gin/project_user_center/internal/handler"
	"iotestgo/module03_web_gin/project_user_center/internal/middleware"
	"iotestgo/module03_web_gin/project_user_center/internal/repository"
	"iotestgo/module03_web_gin/project_user_center/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 1. Config
	viper.SetDefault("app.port", 8090)
	viper.SetDefault("app.name", "user-center")

	// 2. Logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	// 3. Dependency Injection
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	// 4. Router Setup
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(logger))

	// Public Routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	// Protected Routes
	auth := r.Group("/me")
	auth.Use(middleware.Auth())
	auth.GET("", h.Me)

	// 5. Run
	port := viper.GetString("app.port")
	logger.Info("Starting server", zap.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Fatal("Server failed", zap.Error(err))
	}
}
