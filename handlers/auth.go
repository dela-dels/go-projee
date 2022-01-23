package handlers

import (
	"net/http"

	"github.com/dela-dels/go-projee/models"
	"github.com/dela-dels/go-projee/storage/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	//PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDetails struct {
	Email     string
	FirstName string
	Lastname  string
	Password  string
}

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Login Route",
	})
}

var db, _ = database.New().Connect()

func Register(context *gin.Context) {
	var registrationRequest RegistrationRequest

	if err := context.BindJSON(&registrationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to process your request. Please try again",
			"error":   err.Error(),
		})
	}

	password, _ := hashPassword(context.Param("password"))

	userDetails := UserDetails{
		registrationRequest.Email,
		registrationRequest.Firstname,
		registrationRequest.Lastname,
		password,
	}

	err := db.Create(&models.User{
		FirstName: userDetails.FirstName,
		LastName:  userDetails.Lastname,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}).Error

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "Failed",
			"error":  "Operation Failed. Please contact support for assistance",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"user":   userDetails,
	})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}
