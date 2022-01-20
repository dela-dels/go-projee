package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Login Route",
	})
}

func Register(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Registration Route",
	})
}
