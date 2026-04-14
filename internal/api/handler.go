package api

// Handler crea la struttura handler che racchiude service per la raccolta dati

import "github.com/chiaf1/solar-api/internal/metrics"

type Handler struct {
	metrics *metrics.Service
}

func NewHandler(metricService *metrics.Service) *Handler {
	return &Handler{
		metrics: metricService,
	}
}
