package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	RegisterAuthRoutes(router)
	RegisterBookRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Library Management System!",
		})
	})
}
