package main

import (
	"fmt"
	"os"

	db "LibraryManagementSystem/database"
	logger "LibraryManagementSystem/log"

	"github.com/joho/godotenv"
)

func init() {
	logger.Init()
	envError := godotenv.Load()
	if envError != nil {
		logger.Log("Error Loading .env file", envError)
	}
	db.Connect()
}

func main() {
	appEnvironment := os.Getenv("APP_ENV")
	fmt.Println("Hey", appEnvironment)
}
