package repositories

import (
	"context"
	"madastore/analytics/internal/models"
)

type DashboardAnalysisRepositoryInterface interface {
	GetTopProductsVisits(ctx context.Context) ([]models.ProductVisitStat, error)
	GetVisitsPerDay(ctx context.Context) ([]models.VisitsPerDayData, error)
	GetVisitsPerMonth(ctx context.Context) ([]models.VisitsPerMonthData, error)
	GetVisitsPerCountry(ctx context.Context) ([]models.VisitsPerCountryData, error)
	GetVisitsPerCity(ctx context.Context) ([]models.VisitsPerCityData, error)
	GetVisitsFromEgyptPerDay(ctx context.Context) ([]models.VisitsFromEgyptData, error)
	GetVisitsFromOtherCountriesPerDay(ctx context.Context) ([]models.VisitsFromCountriesData, error)
	GetVisitsFromEgyptPerHourForPastMonth(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error)
	GetVisitsFromEgyptPerHourForToday(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error)
}
