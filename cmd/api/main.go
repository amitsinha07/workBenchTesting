package main

import (
	"log"

	"ondc-buyer-app/internal/api/server"
	"ondc-buyer-app/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create and setup server
	srv := server.NewServer(cfg)
	srv.SetupRoutes()

	// Start server
	log.Printf("Starting ONDC Mobility App server on port %s", cfg.Port)
	if err := srv.Start(cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
