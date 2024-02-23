package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/computers33333/airaccidentdata/docs" // Blank identifier.This import is for Swagger documentation generation.

	"github.com/computers33333/airaccidentdata/internal/api"
	"github.com/computers33333/airaccidentdata/internal/config"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// @title AirAccidentData API
// @version 1.0
// @description API server for airaccidentdata.com
// @BasePath /api/v1
// Starting point of the AirAccidentData API server.

// Main function entry point for the AirAccidentData API server.
func main() {
	// Configure the application.
	cfg := config.NewConfig()

	// Initialize the data store.
	store, err := store.NewStore(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	// Set up the API server with the initialized store.
	router := api.NewServer(store)

	// Configure the HTTP server.
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server in a new goroutine for non-blocking operation.
	go func() {
		log.Println("Server is starting...")
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// Log fatal error if the server fails to start.
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Setup for graceful shutdown.
	quit := make(chan os.Signal, 1)
	// Notify the quit channel on SIGINT (Ctrl+C) or SIGTERM signals.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Block until a signal is received.

	// Initiate graceful shutdown with a timeout context.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		// Log fatal error if the server shutdown fails.
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Log server exiting.
	log.Println("Server exiting")
}
