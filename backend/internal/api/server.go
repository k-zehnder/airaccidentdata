package api

import (
	"github.com/computers33333/airaccidentdata/internal/api/router"
	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewServer(store *store.Store) *gin.Engine {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// Create a Gin router with the custom logger and server
	router := router.SetupRouter(store, log)

	return router
}
