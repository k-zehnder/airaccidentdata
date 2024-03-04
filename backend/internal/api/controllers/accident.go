package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// Importing models package
)

// GetAllAccidentsHandler creates a gin.HandlerFunc that handles requests to fetch a list of aviation accidents.
// It utilizes pagination to efficiently return a subset of accidents based on the provided query parameters.
// @Summary Get a list of accidents
// @Description Get list of all aviation accidents with pagination
// @Tags Accidents
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of accidents per page"
// @Success 200 {object} models.AccidentResponse
// @Failure 400 {object} models.ErrorResponse "Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents [get]
func GetAllAccidentsHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting 'page' and 'limit' from the query parameters.
		// Default values are used if they are not provided or invalid.
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
			return
		}

		// Calling the GetAccidents method of the store to retrieve the accidents.
		incidents, total, err := store.GetAccidents(page, limit)
		if err != nil {
			// Logging the error and sending an internal server error response.
			log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to get incidents")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Responding with the accidents and pagination details.
		c.JSON(http.StatusOK, gin.H{
			"accidents": incidents,
			"total":     total,
			"page":      page,
			"limit":     limit,
		})
	}
}

// GetAccidentByIdHandler creates a gin.HandlerFunc that handles requests to fetch an accident by its ID.
// @Summary Get an accident by ID
// @Description Retrieve details of an accident by its ID
// @Tags Accidents
// @Produce json
// @Param id path int true "Accident ID"
// @Success 200 {object} models.AircraftAccidentResponse
// @Failure 400 {object} models.ErrorResponse "Invalid accident ID"
// @Failure 404 {object} models.ErrorResponse "Accident not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents/{id} [get]
func GetAccidentByIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the ID parameter from the URL path.
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid accident ID"})
			return
		}

		// Call the store method to fetch the accident by its ID.
		accident, err := store.GetAccidentById(id)
		if err != nil {
			log.WithError(err).Error("Failed to fetch accident by ID")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accident"})
			return
		}

		// Return the accident in the response.
		c.JSON(http.StatusOK, accident)
	}
}

// GetAllAircraftsHandler creates a gin.HandlerFunc that handles requests to fetch all aircraft.
// @Summary Get a list of aircrafts
// @Description Retrieve a list of all aircraft.
// @Tags Aircrafts
// @Produce json
// @Success 200 {array} models.AircraftListResponse
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts [get]
func GetAllAircraftsHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Calling the GetAllAircrafts method of the store to retrieve all aircraft.
		aircrafts, err := store.GetAllAircrafts()
		if err != nil {
			// Logging the error and sending an internal server error response.
			log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to get aircrafts")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Responding with the aircrafts.
		c.JSON(http.StatusOK, aircrafts)
	}
}

// GetAccidentsByIdnHandler creates a gin.HandlerFunc that handles requests to fetch aviation accidents by ID.
// @Summary Get a list of accidents by aircraft ID
// @Description Get details of an aviation accident by its ID
// @Tags Aircrafts
// @Produce json
// @Param id path string true "ID of the aircraft"
// @Success 200 {object} models.AccidentDetailResponse "Accident details"
// @Failure 400 {object} models.ErrorResponse "Invalid ID number"
// @Failure 404 {object} models.ErrorResponse "Accident not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts/{id}/accidents [get]
func GetAccidentsByIdnHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting the ID from the URL path parameters.
		idStr := c.Param("id")

		// Converting the ID from string to integer.
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
			return
		}

		// Calling the GetAccidentByReg method of the store to retrieve the accident.
		accident, err := store.GetAccidentByIdHandler(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Responding with the accident details.
		c.JSON(http.StatusOK, accident)
	}
}

// GetAircraftByIdHandler creates a gin.HandlerFunc that handles requests to fetch an aircraft by its ID.
// @Summary Get details about an aircraft by ID
// @Description Retrieve details of an aircraft by its ID
// @Tags Aircrafts
// @Produce json
// @Param id path int true "Aircraft ID"
// @Success 200 {object} models.Aircraft
// @Failure 400 {object} models.ErrorResponse "Invalid aircraft ID"
// @Failure 404 {object} models.ErrorResponse "Aircraft not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts/{id} [get]
func GetAircraftByIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the ID parameter from the URL path.
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid aircraft ID"})
			return
		}

		// Call the store method to fetch the aircraft by its ID.
		aircraft, err := store.GetAircraftById(id)
		if err != nil {
			log.WithError(err).Error("Failed to fetch aircraft by ID")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch aircraft"})
			return
		}

		// Return the aircraft in the response.
		c.JSON(http.StatusOK, aircraft)
	}
}
