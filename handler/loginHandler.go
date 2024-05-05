package handler

import (
	db "LibraryManagementSystem/database"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Age      int    `json:"age"`
}

func Login(ctx *gin.Context) {
	var loginInput Credentials
	err := ctx.ShouldBindJSON(&loginInput)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	getUserRowQuery := fmt.Sprintf("select * from users where email = '%s' limit 1;", loginInput.Email)
	fmt.Println(getUserRowQuery)
	userDetailQuery, err := db.GetDatabase().Query(getUserRowQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "There was some error processing your request", "error": err.Error()})
	}
	var user User
	if userDetailQuery.Next() {
		if err := userDetailQuery.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Mobile, &user.Age); err != nil {
			panic(err)
		}
	}
	if user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password combination"})
		return
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": tokenString})
}

func Logout(ctx *gin.Context) {
	fmt.Println("Logging out")
}
