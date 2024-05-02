package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/computers33333/airaccidentdata/internal/models"
	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetAircraftsHandler creates a gin.HandlerFunc that handles requests to fetch all aircraft with pagination.
// @Summary Get a list of aircrafts
// @Description Retrieve a list of all aircrafts.
// @Tags Aircrafts
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of aircraft per page"
// @Success 200 {object} models.AircraftPaginatedResponse "Aircrafts data with pagination details"
// @Failure 400 {object} models.ErrorResponse "Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts [get]
func GetAircraftsHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract page and limit from query parameters with default values
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
			return
		}

		// Call the store method to fetch paginated aircraft data
		aircrafts, totalCount, err := store.GetAircrafts(page, limit)
		if err != nil {
			log.WithError(err).Error("Failed to fetch aircrafts")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch aircrafts"})
			return
		}

		// Respond with paginated aircraft data
		c.JSON(http.StatusOK, gin.H{
			"aircrafts": aircrafts,
			"total":     totalCount,
			"page":      page,
			"limit":     limit,
		})
	}
}

// GetAccidentsHandler creates a gin.HandlerFunc that handles requests to fetch a list of aviation accidents.
// It utilizes pagination to efficiently return a subset of accidents based on the provided query parameters.
// @Summary Get a list of accidents
// @Description Get a list of all aviation accidents
// @Tags Accidents
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of accidents per page"
// @Success 200 {object} models.AccidentPaginatedResponse
// @Failure 400 {object} models.ErrorResponse "Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents [get]
func GetAccidentsHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
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
// @Success 200 {object} models.Aircraft
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

// GetAllImagesForAircraftHandler creates a gin.HandlerFunc that handles requests to fetch all images for a specific aircraft.
// @Summary Get all images for an aircraft
// @Description Retrieve all images associated with a specific aircraft.
// @Tags Aircrafts
// @Produce json
// @Param id path int true "Aircraft ID"
// @Success 200 {object} models.ImagesForAircraftResponse "Image IDs, Image URLs, and S3 URLs"
// @Failure 400 {object} models.ErrorResponse "Invalid aircraft ID"
// @Failure 404 {object} models.ErrorResponse "Aircraft not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts/{id}/images [get]
func GetAllImagesForAircraftHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the aircraft ID from the URL path parameters.
		idStr := c.Param("id")
		aircraftID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid aircraft ID"})
			return
		}

		// Call the store method to fetch all images for the aircraft.
		images, err := store.GetAllImagesForAircraft(aircraftID)
		if err != nil {
			log.WithError(err).Error("Failed to fetch images for aircraft")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to fetch images for aircraft"})
			return
		}

		// Construct the response data.
		response := make([]gin.H, len(images))
		for i, image := range images {
			response[i] = gin.H{
				"id":          image.ID,
				"aircraft_id": image.AircraftID,
				"image_url":   image.ImageURL,
				"s3_url":      image.S3URL,
				"description": image.Description,
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"aircraft_id": aircraftID,
			"images":      response,
		})
	}
}

// GetInjuriesByAccidentIdHandler creates a gin.HandlerFunc that handles requests to fetch injury details for an accident.
// @Summary Get injuries for an accident
// @Description Retrieve injuries for an accident.
// @Tags Injuries
// @Produce json
// @Param id path int true "Aircraft ID"
// @Success 200 {object} models.Injury
// @Failure 400 {object} models.ErrorResponse "Invalid accident ID"
// @Failure 404 {object} models.ErrorResponse "Accident not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /injuries/{id} [get]
func GetInjuriesByAccidentIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the aircraft ID from the URL path parameters.
		idStr := c.Param("id")
		accidentID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid aircraft ID"})
			return
		}

		// Call the store method to fetch paginated aircraft data
		injuries, err := store.GetInjuriesByAccidentIdHandler(accidentID)
		fmt.Println(injuries)
		if err != nil {
			log.WithError(err).Error("Failed to fetch injuries")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch injuries"})
			return
		}

		// Respond with paginated aircraft data
		c.JSON(http.StatusOK, gin.H{
			"injuries": injuries,
		})
	}
}
