package api

// energy contiene le vere e proprie funzioni handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler function for today endpoint
func (h *Handler) GetTodayEnergy(c *gin.Context) {
	data, err := h.metrics.GetTodayEnergy(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch today energy",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Handler function for Yesterday endpoint
func (h *Handler) GetYesterdayEnergy(c *gin.Context) {
	data, err := h.metrics.GetYesterdayEnergy(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch yesterday energy",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Handler function for daily endpoint
func (h *Handler) GetDailyEnergyRange(c *gin.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")
	window := c.DefaultQuery("window", "10m")

	if fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "from and to  parameters are required",
		})
		return
	}

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid from date format (expected YYYY-MM-DD)",
		})
		return
	}
	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid to date format (expected YYYY-MM-DD)",
		})
		return
	}

	data, err := h.metrics.GetRangeEnergyByDay(c.Request.Context(), from, to, window)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch daily energy in range",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
