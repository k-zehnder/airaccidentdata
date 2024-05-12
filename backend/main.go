// Package main runs the AirAccidentData API server.
// @title Air Accident Data API
// @description This is the server for managing air accident data.
// @version 1
// @BasePath /api/v1
package main

import (
	"log"

	_ "github.com/computers33333/airaccidentdata/docs" // Import for Swagger documentation generation
	"github.com/joho/godotenv"

	"github.com/computers33333/airaccidentdata/internal/api/server"
	"github.com/computers33333/airaccidentdata/internal/config"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// main sets up and starts the API server.
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Load configuration settings
	cfg := config.NewConfig()

	// Initialize the data store
	store, err := store.NewStore(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	// Initialize and start the HTTP server
	router := server.NewServer(store)
	srv := server.StartServer(":8080", router)
	defer server.GracefulShutdown(srv)
}
