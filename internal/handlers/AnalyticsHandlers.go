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
	data, err := h.analyticsService.GetDashboardAnalytics(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)

}

func (h *AnalyticsHandlers) GetTopProductsVisits(c *gin.Context) {
	data, err := h.analyticsService.GetTopProductsVisits(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsPerDay(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsPerDay(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsPerMonth(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsPerMonth(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsPerCountry(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsPerCountry(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}

func (h *AnalyticsHandlers) GetVisitsPerCity(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsPerCity(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsFromEgyptPerDay(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsFromEgyptPerDay(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsFromOtherCountriesPerDay(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsFromOtherCountriesPerDay(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsFromEgyptPerHourForPastMonth(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsFromEgyptPerHourForPastMonth(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
func (h *AnalyticsHandlers) GetVisitsFromEgyptPerHourForToday(c *gin.Context) {
	data, err := h.analyticsService.GetVisitsFromEgyptPerHourForToday(c.Request.Context())
	if err != nil {
		utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondWithJSON(c.Writer, http.StatusOK, data)
}
