// Package main is the entry point for the AirAccidentData API server. This program sets up and runs an HTTP server
// with endpoints for accessing air accident data. It includes features like graceful shutdown, Swagger documentation
// support, and a configured API router. The server reads its configuration, sets up the necessary data store,
// and listens on a specified port for incoming HTTP requests.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/computers33333/airaccidentdata/docs" // Used for creating the Swagger documentation.

	"github.com/computers33333/airaccidentdata/internal/api"
	"github.com/computers33333/airaccidentdata/internal/config"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// Starting point of the AirAccidentData API server.
func main() {
	// Set up the application settings.
	cfg := config.NewConfig()

	// Create a new place to store our data.
	store, err := store.NewStore(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Couldn't make a data store: %v", err)
	}

	// Get our API server ready using the data store.
	router := api.NewServer(store)

	// Set up the web server.
	httpServer := &http.Server{
		Addr:    ":8080", // The address to listen on.
		Handler: router,  // The handler to use, in this case, our router.
	}

	// Run the server in the background so it doesn't stop our program.
	go func() {
		log.Println("Server is getting ready...")
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// If the server can't start, show an error and stop.
			log.Fatalf("The server couldn't start: %v", err)
		}
	}()

	// Prepare for a smooth shutdown.
	quit := make(chan os.Signal, 1)
	// Listen for signals to stop the server (like pressing Ctrl+C).
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit  // Wait here until we get a signal to stop.

	// Start shutting down and give it 5 seconds to finish everything.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		// If there's a problem shutting down, show an error.
		log.Fatalf("Had to force the server to stop: %v", err)
	}

	// Say that the server has stop
