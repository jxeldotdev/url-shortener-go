package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jxeldotdev/url-backend/api/v1/controllers"
	"github.com/jxeldotdev/url-backend/config"
	"github.com/jxeldotdev/url-backend/internal/db"
	"github.com/jxeldotdev/url-backend/middleware"
	"github.com/jxeldotdev/url-backend/models"
)

var AppLogger *log.Logger

func main() {
	AppLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	// Get config
	var cfg config.Config
	config.PopulateConfigFromEnv(&cfg)

	// Spin up database
	AppLogger.Printf("Connecting to Database at %s:%v", cfg.Database.Host.Value, cfg.Database.Port.Value)
	db.Connect(&cfg)

	AppLogger.Println("Running migrations")
	db.Database.AutoMigrate(&models.User{})
	db.Database.AutoMigrate(&models.Url{})

	//db.Database.AutoMigrate(&models.ApiPermission{})
	//db.Database.AutoMigrate(&models.Role{})
	//models.CreateRoles()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/url", controllers.AddUrl)
	protectedRoutes.GET("/url", controllers.GetAllUrls)
	protectedRoutes.GET("/url/:id", controllers.GetSingleUrl)
	protectedRoutes.DELETE("/url/:id", controllers.DeleteUrl)
	protectedRoutes.GET("/user", controllers.GetAllUsers)
	protectedRoutes.GET("/user/:id/urls")

	router.GET("/r/:shortUrl", controllers.RedirectToLongUrl)
	AppLogger.Println("Initializing Server")
	router.Run(":8080")
}
