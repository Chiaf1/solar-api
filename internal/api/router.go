package api

// Gestisce le rout dell'api e basta

import "github.com/gin-gonic/gin"

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	energy := r.Group("/energy")
	{
		energy.GET("/today")
		energy.GET("/yesterday")
		energy.GET("/daily")
	}

	return r
}
