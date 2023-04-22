package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jxeldotdev/url-backend/api/v1/controllers"
	"github.com/jxeldotdev/url-backend/config"
	"github.com/jxeldotdev/url-backend/internal/db"
	"github.com/jxeldotdev/url-backend/middleware"
	"github.com/jxeldotdev/url-backend/models"
)

func main() {
	// Get config
	var cfg config.Config
	config.PopulateConfigFromEnv(&cfg)

	// Spin up database
	db.Connect(&cfg)

	db.Database.AutoMigrate(&models.User{})
	db.Database.AutoMigrate(&models.Url{})
	serveApplication()
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/url", controllers.AddUrl)
	protectedRoutes.GET("/url", controllers.GetAllUrls)

	router.Run(":8080")
	fmt.Println("Running on 8080")

}
