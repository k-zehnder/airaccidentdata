// Package router sets up the Gin router with routes and middleware.
package router

import (
	"net/http"

	"github.com/computers33333/airaccidentdata/internal/api/controllers"
	"github.com/computers33333/airaccidentdata/internal/api/middleware"
	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures a Gin router with necessary routes, middleware, and CORS policies.
func SetupRouter(store *store.Store, log *logrus.Logger) *gin.Engine {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080", "https://airaccidentdata.com", "https://www.airaccidentdata.com"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	// Swagger API documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply custom logging middleware
	router.Use(middleware.LoggingMiddleware(log))

	// Simple root GET
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "API is running"})
	})

	// API version 1 routes
	v1 := router.Group("/api/v1")
	{
		setupAircraftRoutes(v1, store, log)
		setupAccidentRoutes(v1, store, log)
		setupInjuryRoutes(v1, store, log)
	}

	return router
}

func setupAircraftRoutes(r *gin.RouterGroup, store *store.Store, log *logrus.Logger) {
	aircrafts := r.Group("/aircrafts")
	{
		aircrafts.GET("/", controllers.GetAircraftsHandler(store, log))
		aircrafts.GET("/:id", controllers.GetAircraftByIdHandler(store, log))
		aircrafts.GET("/:id/accidents", controllers.GetAccidentByIdHandler(store, log))
		aircrafts.GET("/:id/images", controllers.GetAllImagesForAircraftHandler(store, log))
	}
}

func setupAccidentRoutes(r *gin.RouterGroup, store *store.Store, log *logrus.Logger) {
	accidents := r.Group("/accidents")
	{
		accidents.GET("/", controllers.GetAccidentsHandler(store, log))
		accidents.GET("/:id", controllers.GetAccidentByIdHandler(store, log))
	}
}

func setupInjuryRoutes(r *gin.RouterGroup, store *store.Store, log *logrus.Logger) {
	injuries := r.Group("/injuries")
	{
		injuries.GET("/:id", controllers.GetInjuriesByAccidentIdHandler(store, log))
	}
}
