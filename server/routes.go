package server

import (
	"github.com/dela-dels/go-projee/handlers"
	"github.com/dela-dels/go-projee/handlers/projects"
)

func (s *Server) setupRoutes() {
	s.Mux.GET("/ping", handlers.Ping)

	s.Mux.POST("login", handlers.Login)
	s.Mux.POST("register", handlers.Register)

	s.Mux.POST("/projects", projects.Create)
}
