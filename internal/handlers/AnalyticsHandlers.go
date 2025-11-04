package handlers

import (
	"madastore/analytics/internal/services"
	"madastore/analytics/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandlers struct {
	analyticsService *services.DashboardAnalysisService
}

// NewAnalyticsHandlers creates a new instance of AnalyticsHandlers
func NewAnalyticsHandlers(service *services.DashboardAnalysisService) *AnalyticsHandlers {
	return &AnalyticsHandlers{
		analyticsService: service,
	}
}

func (h *AnalyticsHandlers) GetDashboardData(c *gin.Context) {
	data, err := h.analyticsService.GetTopProductsVisits(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)

}
