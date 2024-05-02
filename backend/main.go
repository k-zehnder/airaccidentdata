// Package main sets up and runs the AirAccidentData API server.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/computers33333/airaccidentdata/docs" // Import for Swagger doc generation

	"github.com/computers33333/airaccidentdata/internal/api"
	"github.com/computers33333/airaccidentdata/internal/config"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// main initializes and runs the API server.
func main() {
	cfg := config.NewConfig()

	store, err := store.NewStore(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	router := api.NewServer(store)

	server := startServer(":8080", router)
	defer gracefulShutdown(server)
}

// startServer initializes and starts the HTTP server.
func startServer(addr string, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return server
}

// gracefulShutdown handles the clean shutdown of the server upon receiving a signal.
func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

	log.Println("Server exited")
}
