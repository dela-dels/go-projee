package projects

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Project struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartDate   time.Time `json:"start_date" time_format:"2006-01-02" binding:"required"`
	EndDate     time.Time `json:"end_date" time_format:"2006-01-01-02" binding:"required"`
	IsCompleted bool      `json:"is_completed"`

	ProjectLead int `json:"project_lead"`
}

func Create(context *gin.Context) {
	var project Project

	if err := context.BindJSON(&project); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to process your request. Please try again",
			"error":   err.Error(),
		})
	}

	// Here, if everything goes well with the request, we will persist the project
	// into storage and return the project details to the client
	context.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"data":   project,
	})
}

func Index(context *gin.Context) {
	email := context.Param("email")
	context.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"projects": email,
	})
}
