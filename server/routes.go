package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) setupRoutes() {
	s.Mux.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name": "Dela Akakpo",
		})
	})
}
