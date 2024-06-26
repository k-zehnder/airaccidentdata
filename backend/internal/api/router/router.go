// Package router sets up the Gin router with routes and middleware.
package router

import (
	"github.com/computers33333/airaccidentdata/internal/api/controllers"
	"github.com/computers33333/airaccidentdata/internal/api/middleware"
	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter initializes a new Gin web server with custom logging and routing configured.
func NewRouter(store *store.Store) *gin.Engine {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	router := SetupRouter(store, log)

	return router
}

// SetupRouter configures a Gin router with necessary routes, middleware, and CORS policies.
func SetupRouter(store *store.Store, log *logrus.Logger) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080", "https://airaccidentdata.com", "https://www.airaccidentdata.com"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.LoggingMiddleware(log))

	v1 := router.Group("/api/v1")
	{
		aircrafts := v1.Group("/aircrafts")
		{
			aircrafts.GET("", controllers.GetAircraftsHandler(store, log))
			aircrafts.GET("/:id", controllers.GetAircraftByIdHandler(store, log))
			aircrafts.GET("/:id/accidents", controllers.GetAccidentByIdHandler(store, log))
			aircrafts.GET("/:id/images", controllers.GetAllImagesForAircraftHandler(store, log))
		}

		accidents := v1.Group("/accidents")
		{
			accidents.GET("", controllers.GetAccidentsHandler(store, log))
			accidents.GET("/:id", controllers.GetAccidentByIdHandler(store, log))
			accidents.GET("/:id/location", controllers.GetLocationByAccidentIdHandler(store, log))
			accidents.GET("/:id/injuries", controllers.GetInjuriesByAccidentIdHandler(store, log))
		}
	}

	return router
}
