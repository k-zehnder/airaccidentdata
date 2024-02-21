package controllers

import (
	"net/http"
	"strconv"

	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetAccidentsHandler creates a gin.HandlerFunc that handles requests to fetch a list of aviation accidents.
// It utilizes pagination to efficiently return a subset of accidents based on the provided query parameters.
// @Summary Get list of Accidents
// @Description Get list of all aviation accidents with pagination
// @Tags Accidents
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of accidents per page"
// @Success 200 {object} models.AccidentResponse
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
