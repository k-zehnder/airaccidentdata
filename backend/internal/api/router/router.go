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

// SetupRouter sets up the Gin router with routes and middleware
func SetupRouter(store *store.Store, log *logrus.Logger) *gin.Engine {
	// Initialize a new Gin router.
	router := gin.Default()

	// Configure CORS.
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080", "https://airaccidentdata.com", "https://www.airaccidentdata.com"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	// Serve Swagger documentation.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply the custom logging middleware.
	router.Use(middleware.LoggingMiddleware(log))

	// Serve a debug route at :8080.
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "meow")
	})

	// Set up the v1 routes group
	v1 := router.Group("/api/v1")
	{
		// Set up aircrafts route
		aircrafts := v1.Group("/aircrafts")
		{
			aircrafts.GET("", controllers.GetAllAircraftsHandler(store, log))
			aircrafts.GET("/:id", controllers.GetAircraftByIdHandler(store, log))
			aircrafts.GET("/:id/accidents", controllers.GetAccidentByIdHandler(store, log))
			aircrafts.GET("/:id/images", controllers.GetAllImagesForAircraftHandler(store, log))
			aircrafts.GET("/:id/images/:imageID", controllers.GetImageForAircraftHandler(store, log))
		}

		// Set up accidents route
		accidents := v1.Group("/accidents")
		{
			accidents.GET("", controllers.GetAllAccidentsHandler(store, log))
			accidents.GET("/:id", controllers.GetAccidentByIdHandler(store, log))
		}
	}

	// Return configured router.
	return router
}
