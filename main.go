package main

import (
	"log"

	"github.com/dela-dels/go-projee/models"
	"github.com/dela-dels/go-projee/server"
	"github.com/dela-dels/go-projee/storage/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	setupDatabaseSystem()
	setupWebServer()
}

func setupWebServer() {
	webServer := server.New()

	if err := webServer.Start(); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}

func setupDatabaseSystem() {
	database := database.New()
	connection, err := database.Connect()

	if err != nil {
		log.Fatalf("Failed to start and connect to database server. Error: %v", err)
	}

	if err := database.RunMigrations(connection, &models.Project{}, &models.Sprint{}, &models.Task{}, &models.User{}, &models.UserTasks{}); err != nil {
		log.Fatalf("Failed to run migrations. Error: %v", err)
	}

}
