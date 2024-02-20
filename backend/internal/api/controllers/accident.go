package controllers

import (
	"net/http"

	"github.com/computers33333/airaccidentdata/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Get list of Accidents
// @Description Get list of all aviation accidents
// @Tags Accidents
// @Produce  json
// @Success 200 {array} models.AircraftAccident
// @Failure 500 {object} models.ErrorResponse
// @Router /accidents [get]
func GetAccidentsHandler(store *store.Store, log *logrus.Logger) gin.HandlerFunc {
	// Closure
	// Things here only get called once on server start.
	return func(c *gin.Context) {
		// Things here get called on every request.
		incidents, err := store.GetAccidents()
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to get incidents")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, incidents)
	}
}
