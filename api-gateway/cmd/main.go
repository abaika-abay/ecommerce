package main

import (
	"log"

	"api-gateway/internal/app"
)

func main() {
	server := app.NewApp()

	log.Println("API Gateway started on :8080")
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
