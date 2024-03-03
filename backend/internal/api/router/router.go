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
		v1.GET("/aircrafts", controllers.GetAllAircraftsHandler(store, log))
		v1.GET("/aircrafts/:registration_number/accidents", controllers.GetAccidentsByRegistrationHandler(store, log))
		v1.GET("/aircrafts/byId/:id", controllers.GetAircraftByIdHandler(store, log))
		v1.GET("/accidents", controllers.GetAllAccidentsHandler(store, log))
	}

	// Return configured router.
	return router
}
