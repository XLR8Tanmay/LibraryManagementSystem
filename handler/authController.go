package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("APP_KEY"))

func Authenticate(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing auth token"})
		ctx.Abort()
		return
	}

	// Parse and validate JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token supplied"})
		ctx.Abort()
		return
	}

	// Pass the user information to the next handler
	claims := token.Claims.(jwt.MapClaims)
	ctx.Set("user", claims["email"]) // Set user context
	ctx.Next()
	fmt.Println("Authenticating User")
}
