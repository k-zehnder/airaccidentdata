// Package api sets up the web server and routing logic for the AirAccidentData application.
package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/computers33333/airaccidentdata/internal/api/router"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// NewRouter initializes a new Gin router with custom logging and routing configured.
func NewRouter(store *store.Store) *gin.Engine {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	router := router.SetupRouter(store, log)

	return router
}

// StartServer initializes and starts the HTTP server.
func StartServer(addr string, handler http.Handler) *http.Server {
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

// GracefulShutdown handles the clean shutdown of the server upon receiving a signal.
func GracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Blocks

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

	log.Println("Server exited cleanly")
}
