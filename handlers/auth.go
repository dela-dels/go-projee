package handlers

import (
	"log"
	"net/http"

	"github.com/dela-dels/go-projee/models"
	"github.com/dela-dels/go-projee/storage/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db, _ = database.New().Connect()

//TODO: will probably use the go-validator package to create a custom unique-email validation where we throw a custom error for duplicate errors
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
	//get the user's credentials from the request
	//retrieve the users's data from storage matching the credentials sent
	//generate a jwt token for the user trying to authenticate
	var loginRequest LoginRequest

	if err := context.BindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to process your request. Please try again and make sure you are sending the right data",
		})
		return
	}

	user := models.User{}
	err := db.Where("email = ?", loginRequest.Email).First(&user).Error

	if err != nil {
		context.JSON(404, gin.H{
			"status": "Failed",
			"error":  err.Error(),
		})
		return
	}

	log.Printf("user password %s and request password %s", user.Password, loginRequest.Password)

	err = hashMatchesPassword(user.Password, loginRequest.Password)

	log.Printf("hash pass check error is: %s", err)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "Failed",
			"error":  "Invalid credentials. Please check and try again.",
			"res": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": user,
	})
	return
}

func Register(context *gin.Context) {
	var registrationRequest RegistrationRequest

	if err := context.BindJSON(&registrationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to process your request. Please try again",
			"error":   err.Error(),
		})
	}

	password, _ := hashPassword(registrationRequest.Password)

	err := db.Create(&models.User{
		FirstName: registrationRequest.Firstname,
		LastName:  registrationRequest.Lastname,
		Email:     registrationRequest.Email,
		Password:  password,
	}).Error

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "Failed",
			"error":  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"user":   err,
	})
	return
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func hashMatchesPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
