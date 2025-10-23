package routes

import (
	"LibraryManagementSystem/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
	bookRoute := router.Group("/v1")
	bookRoute.Use(handler.Authenticate)
	{
		bookRoute.GET("/get-all-books", handler.GetAllBooks)
	}
}
