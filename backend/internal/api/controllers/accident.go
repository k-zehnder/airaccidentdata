// Package controllers handles HTTP requests and orchestrates responses
// by interacting with the underlying data model through the store.
package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/computers33333/airaccidentdata/internal/models"
	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetAircraftsHandler returns a handler for fetching all aircraft with pagination.
// @Summary Get a list of aircrafts
// @Description Retrieve a list of all aircrafts with pagination.
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

		aircrafts, totalCount, err := store.GetAircrafts(page, limit)
		if err != nil {
			log.WithError(err).Error("Failed to fetch aircrafts")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch aircrafts"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"aircrafts": aircrafts,
			"total":     totalCount,
			"page":      page,
			"limit":     limit,
		})
	}
}

// GetAccidentsHandler returns a handler for fetching a list of aviation accidents with pagination.
// @Summary Get a list of accidents
// @Description Get a list of all aviation accidents with pagination.
// @Tags Accidents
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of accidents per page"
// @Success 200 {object} models.AccidentPaginatedResponse "Accidents data with pagination details"
// @Failure 400 {object} models.ErrorResponse "Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents [get]
func GetAccidentsHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
			return
		}

		accidents, total, err := store.GetAccidents(page, limit)
		if err != nil {
			log.WithError(err).Error("Failed to get accidents")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get accidents"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accidents": accidents,
			"total":     total,
			"page":      page,
			"limit":     limit,
		})
	}
}

// GetAccidentByIdHandler returns a handler for fetching an accident by its ID.
// @Summary Get an accident by ID
// @Description Retrieve details of an accident by its ID
// @Tags Accidents
// @Produce json
// @Param id path int true "Accident ID"
// @Success 200 {object} models.Accident "Detailed accident data"
// @Failure 400 {object} models.ErrorResponse "Invalid accident ID"
// @Failure 404 {object} models.ErrorResponse "Accident not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents/{id} [get]
func GetAccidentByIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid accident ID"})
			return
		}

		accident, err := store.GetAccidentById(id)
		if err != nil {
			log.WithError(err).Error("Failed to fetch accident")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accident"})
			return
		}

		if accident == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Accident not found"})
			return
		}

		c.JSON(http.StatusOK, accident)
	}
}

// GetLocationByAccidentIdHandler returns a handler for fetching location details by accident ID.
// @Summary Get location by accident ID
// @Description Retrieve location details of an accident by its ID
// @Tags Accidents
// @Produce json
// @Param id path int true "Accident ID"
// @Success 200 {object} models.Location "Detailed location data"
// @Failure 400 {object} models.ErrorResponse "Invalid accident ID"
// @Failure 404 {object} models.ErrorResponse "Location not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents/{id}/location [get]
func GetLocationByAccidentIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		accidentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid accident ID"})
			return
		}

		location, err := store.GetLocationByAccidentId(accidentId)
		if err != nil {
			log.WithError(err).Error("Failed to fetch location")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to fetch location"})
			return
		}

		if location == nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Location not found"})
			return
		}

		c.JSON(http.StatusOK, location)
	}
}

// GetAircraftByIdHandler returns a handler for fetching an aircraft by its ID.
// @Summary Get details about an aircraft by ID
// @Description Retrieve details of an aircraft by its ID
// @Tags Aircrafts
// @Produce json
// @Param id path int true "Aircraft ID"
// @Success 200 {object} models.Aircraft "Detailed aircraft data"
// @Failure 400 {object} models.ErrorResponse "Invalid aircraft ID"
// @Failure 404 {object} models.ErrorResponse "Aircraft not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts/{id} [get]
func GetAircraftByIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid aircraft ID"})
			return
		}

		aircraft, err := store.GetAircraftById(id)
		if err != nil {
			log.WithError(err).Error("Failed to fetch aircraft")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch aircraft"})
			return
		}

		c.JSON(http.StatusOK, aircraft)
	}
}

// GetAllImagesForAircraftHandler returns a handler for fetching all images associated with a specific aircraft.
// @Summary Get all images for an aircraft
// @Description Retrieve all images associated with a specific aircraft.
// @Tags Aircrafts
// @Produce json
// @Param id path int true "Aircraft ID"
// @Success 200 {array} models.AircraftImage "List of aircraft images"
// @Failure 400 {object} models.ErrorResponse "Invalid aircraft ID"
// @Failure 404 {object} models.ErrorResponse "Aircraft not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /aircrafts/{id}/images [get]
func GetAllImagesForAircraftHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid aircraft ID"})
			return
		}

		images, err := store.GetAllImagesForAircraft(id)
		if err != nil {
			log.WithError(err).Error("Failed to fetch images for aircraft")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch images"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"images": images,
		})
	}
}

// GetInjuriesByAccidentIdHandler returns a handler for fetching injury details associated with a specific accident.
// @Summary Get injuries for an accident
// @Description Retrieve injury details for an accident based on the provided ID.
// @Tags Accidents
// @Produce json
// @Param id path int true "Accident ID"
// @Success 200 {array} models.Injury "List of injuries associated with the accident"
// @Failure 400 {object} models.ErrorResponse "Invalid accident ID provided"
// @Failure 404 {object} models.ErrorResponse "No injuries found for the specified accident ID"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accidents/{id}/injuries [get]
func GetInjuriesByAccidentIdHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.WithField("accidentID", c.Param("id")).WithError(err).Error("Invalid accident ID format")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid accident ID"})
			return
		}

		injuries, err := store.GetInjuriesByAccidentIdHandler(id)
		if err != nil {
			if err == sql.ErrNoRows {
				log.WithField("accidentID", id).Info("No injuries found for the accident ID")
				c.JSON(http.StatusNotFound, gin.H{"error": "No injuries found for the specified accident ID"})
			} else {
				log.WithField("accidentID", id).WithError(err).Error("Failed to fetch injuries")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch injuries"})
			}
			return
		}

		if len(injuries) == 0 {
			log.WithField("accidentID", id).Info("No injuries found for the accident ID")
			c.JSON(http.StatusNotFound, gin.H{"error": "No injuries found for the specified accident ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"injuries": injuries,
		})
	}
}
