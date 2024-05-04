package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	fmt.Println("Getting books")
}
