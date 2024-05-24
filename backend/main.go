// Package main runs the AirAccidentData API server.
// @title AirAccidentData.com API
// @description API server for managing air accident data.
// @version 1
// @BasePath /api/v1
package main

import (
	"log"

	"github.com/computers33333/airaccidentdata/internal/api/server"
	"github.com/computers33333/airaccidentdata/internal/config"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// main sets up and starts the API server.
func main() {
	// Load configuration settings
	cfg := config.NewConfig()

	// Initialize the database store
	store, err := store.NewStore(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	// Create the router
	router := server.NewRouter(store)

	// Start the HTTP server
	srv := server.StartServer(cfg.ServerAddress, router)
	defer server.GracefulShutdown(srv)
}
