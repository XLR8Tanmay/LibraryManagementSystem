package routes

import (
	"LibraryManagementSystem/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/")
	router.GET("/register", handler.Register)
	router.GET("/login", handler.Login)
	authRoutes.Use(handler.Authenticate)
	{
		authRoutes.POST("/logout", handler.Logout)
	}

}
