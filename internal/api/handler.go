package api

// Handler crea la struttura handler che racchiude service per la raccolta dati

import "github.com/chiaf1/solar-api/internal/metrics"

type Handler struct {
	metrics *metrics.Service
}

// Creates a new handler to handle http responses
func NewHandler(metricService *metrics.Service) *Handler {
	return &Handler{
		metrics: metricService,
	}
}
