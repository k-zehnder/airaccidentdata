// Package api sets up the web server and routing logic for the AirAccidentData application.
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/computers33333/airaccidentdata/internal/api/router"
	"github.com/computers33333/airaccidentdata/internal/store"
)

// NewServer initializes a new Gin web server with custom logging and routing configured.
func NewServer(store *store.Store) *gin.Engine {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	router := router.SetupRouter(store, log)

	return router
}
