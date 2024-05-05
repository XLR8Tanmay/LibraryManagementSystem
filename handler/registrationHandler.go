package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	db "LibraryManagementSystem/database"

	"golang.org/x/crypto/bcrypt"
)

type RegistrationInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8,max=50" regex:"^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)[a-zA-Z\\d]+$" binding_error:"The password field must contain atleast one lowercase letter, one uppercase letter and a number"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,min=1,max=100"`
	Mobile   string `json:"mobile" binding:"required"`
}

func Register(c *gin.Context) {
	var registrationDetails RegistrationInput
	err := c.ShouldBindJSON(&registrationDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	registrationDetails.Password = string(hashedPassword)
	dbRegisterQuery := "INSERT INTO users (username, password, email, mobile, age) VALUES ('%s','%s','%s','%s',%d);"
	dbRegisterQuery = fmt.Sprintf(dbRegisterQuery, registrationDetails.Username, registrationDetails.Password, registrationDetails.Email, registrationDetails.Mobile, registrationDetails.Age)

	database := db.GetDatabase()
	res, err := database.Exec(dbRegisterQuery)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Registration Unsuccessful", "error": err.Error()})
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Registration unsuccessful", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration Successful!", "rows_affected": rowsAffected})
}
