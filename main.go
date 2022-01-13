package main

import (
	"fmt"
	"github.com/dela-dels/go-projee/server"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Print("this is a message")

	godotenv.Load()

	webServer := server.New()
	if err := webServer.Start(); err != nil {
		fmt.Errorf("unable to start server: %w", err)
	}
}