package api

// Gestisce le rout dell'api e basta

import "github.com/gin-gonic/gin"

// Creates the router of the http endpoints for the api
func NewRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	energy := r.Group("/energy")
	{
		energy.GET("/today", h.GetTodayEnergy)
		energy.GET("/yesterday", h.GetYesterdayEnergy)
		energy.GET("/daily", h.GetDailyEnergyRange)
	}

	return r
}
