package main

import (
	"log"
	"os"

	"study-app/internal/database"
	"study-app/internal/server"
)

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Ensure database connection is closed when the program exits
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	if err := server.StartServer(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
