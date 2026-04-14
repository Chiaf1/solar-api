package api

// energy contiene le vere e proprie funzioni handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
