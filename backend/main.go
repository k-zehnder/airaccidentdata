// Package main runs the AirAccidentData API server.
package main

import (
	"log"

	_ "github.com/computers33333/airaccidentdata/docs" // Enables Swagger documentation generation.
	"github.com/joho/godotenv"

	"github.com/computers33333/airaccidentdata/internal/api/server"
	"github.com/computers33333/airaccidentdata/internal/config"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// main sets up and starts the API server.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config.NewConfig()

	store, err := store.NewStore(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	router := server.NewServer(store)
	srv := server.StartServer(":8080", router)
	defer server.GracefulShutdown(srv)
}
