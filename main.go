package main

import (
	"fmt"
	"github.com/dela-dels/go-projee/models"
	"github.com/dela-dels/go-projee/server"
	"github.com/dela-dels/go-projee/storage/database"
	"github.com/joho/godotenv"
	"log"
)

func main(){
	godotenv.Load()

	setupDatabaseSystem()

	webServer := server.New()
	if err := webServer.Start(); err != nil {
		fmt.Errorf("unable to start server: %w", err)
	}
}

func setupDatabaseSystem() {
	database := database.New()
	connection, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to start and connect to database server. Error: %v", err)
	}

	database.RunMigrations(connection, &models.Project{}, &models.Sprint{}, &models.Task{}, &models.User{}, &models.UserTasks{})
}