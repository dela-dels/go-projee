package server

import (
	"github.com/dela-dels/go-projee/handlers"
	"github.com/dela-dels/go-projee/handlers/projects"
	"github.com/dela-dels/go-projee/middlewares"
)

func (s *Server) setupRoutes() {
	s.Mux.GET("/ping", handlers.Ping)

	s.Mux.POST("login", handlers.Login)
	s.Mux.POST("register", handlers.Register)

	s.Mux.POST("/projects", projects.Create)
	s.Mux.GET("/projects/user/:email", middlewares.Authenticated(), projects.Index)
}
