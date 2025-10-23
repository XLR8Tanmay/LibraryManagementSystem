package main

import (
	"fmt"
	"os"

	db "LibraryManagementSystem/database"
	logger "LibraryManagementSystem/log"
	"LibraryManagementSystem/routes"

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
	// Register all routes in one place
	routes.RegisterRoutes(router)

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
