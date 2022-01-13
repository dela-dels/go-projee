package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

//type Server struct {
//	Address string
//	Mux gin.Context
//	Server *http.Server
//}

type Server struct {
	Mux *gin.Engine
	Port string
	Server *http.Server
}

func New() *Server {
	mux := gin.Default()
	port := os.Getenv("APP_PORT")

	return &Server{
		Mux:     mux,
		Port: port,
		Server:  &http.Server{
			Addr: port,
			Handler: mux,
			ReadTimeout: 5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout: 5 * time.Second,
			TLSConfig: &tls.Config{
				CurvePreferences: []tls.CurveID{
					tls.CurveP256,
					tls.X25519,
				},
			},
		},
	}
}

func (s *Server) Start() error {
	s.setupRoutes()

	fmt.Println("Starting server on port", s.Port)
	if err := s.Server.ListenAndServe(); err != nil {
		return fmt.Errorf("error starting the server. error: %w", err)
	}

	return nil
}

func (s *Server) Stop() error {
	fmt.Println("Attempting to stop server")

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		return  fmt.Errorf("error shutting down the server. Error: %w", err)
	}

	return nil
}

