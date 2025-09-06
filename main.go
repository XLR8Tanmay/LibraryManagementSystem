package main

import (
	"fmt"
	"os"

	db "LibraryManagementSystem/database"
	"LibraryManagementSystem/handler"
	logger "LibraryManagementSystem/log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	logger.Init()
	envError := godotenv.Load()
	if envError != nil {
		logger.Log("Error Loading .env file", envError)
	}
	db.Connect()
	db.Migrate()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Library Management System!",
		})
	})
	router.GET("/register", handler.Register)
	router.GET("/login", handler.Login)

	authRoutes := router.Group("/")
	authRoutes.Use(handler.Authenticate)
	{
		authRoutes.GET("/get-all-books", handler.GetAllBooks)
		authRoutes.POST("/logout", handler.Logout)
	}

	router.Run(":" + getAppServerPort())
	fmt.Println("Library Management System running on port: ", getAppServerPort())
}

func getAppServerPort() string {
	serverPort := os.Getenv("APP_PORT")
	if len(serverPort) > 0 {
		return serverPort
	}
	return "8000"
}
