package services

import (
	"context"
	"log"
	"madastore/analytics/internal/models"
	repositories "madastore/analytics/internal/repository/interfaces"
	"madastore/analytics/internal/utils"
	"time"

	"golang.org/x/sync/errgroup"
)

// --------------------
// Service Implementation
// --------------------

type DashboardAnalysisService struct {
	analytticsRepo repositories.DashboardAnalysisRepositoryInterface
}

// Constructor
func NewDashboardAnalysisService(repo repositories.DashboardAnalysisRepositoryInterface) *DashboardAnalysisService {
	return &DashboardAnalysisService{analytticsRepo: repo}
}

// Example method to get top products visits
func (s *DashboardAnalysisService) GetDashboardAnalytics(ctx context.Context) (models.DashboardData, error) {
	var data models.DashboardData
	cacheKey := "dashboard_data"

	// 🧠 Step 1: Try getting cached result
	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	// 🧩 Step 2: Not cached — fetch fresh data from DB
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		topProducts, err := s.analytticsRepo.GetTopProductsVisits(ctx)
		if err != nil {
			return err
		}
		data.TopProducts = topProducts
		return nil
	})

	g.Go(func() error {
		perDay, err := s.analytticsRepo.GetVisitsPerDay(ctx)
		if err != nil {
			return err
		}
		data.PerDay = perDay
		return nil
	})

	g.Go(func() error {
		perMonth, err := s.analytticsRepo.GetVisitsPerMonth(ctx)
		if err != nil {
			return err
		}
		data.PerMonth = perMonth
		return nil
	})

	g.Go(func() error {
		perCountry, err := s.analytticsRepo.GetVisitsPerCountry(ctx)
		if err != nil {
			return err
		}
		data.PerCountry = perCountry
		return nil
	})

	g.Go(func() error {
		perCity, err := s.analytticsRepo.GetVisitsPerCity(ctx)
		if err != nil {
			return err
		}
		data.PerCity = perCity
		return nil
	})

	g.Go(func() error {
		fromEgypt, err := s.analytticsRepo.GetVisitsFromEgyptPerDay(ctx)
		if err != nil {
			return err
		}
		data.FromEgyptPerDay = fromEgypt
		return nil
	})

	g.Go(func() error {
		fromOthers, err := s.analytticsRepo.GetVisitsFromOtherCountriesPerDay(ctx)
		if err != nil {
			return err
		}
		data.FromOtherCountriesPerDay = fromOthers
		return nil
	})

	g.Go(func() error {
		fromEgyptPastMonth, err := s.analytticsRepo.GetVisitsFromEgyptPerHourForPastMonth(ctx)
		if err != nil {
			return err
		}
		data.FromEgyptPerHourForPastMonth = fromEgyptPastMonth
		return nil
	})

	g.Go(func() error {
		fromEgyptToday, err := s.analytticsRepo.GetVisitsFromEgyptPerHourForToday(ctx)
		if err != nil {
			return err
		}
		data.FromEgyptPerHourForToday = fromEgyptToday
		return nil
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return data, err
	}

	// 💾 Step 3: Cache result for 5 minutes
	if err := utils.SetCache(cacheKey, data, 5*time.Minute); err != nil {
		// Don't crash if cache write fails — just log or ignore
		log.Printf("Failed to set cache: %v", err)
	}

	return data, nil
}
func (s *DashboardAnalysisService) GetTopProductsVisits(ctx context.Context) ([]models.ProductVisitStat, error) {
	var data []models.ProductVisitStat
	cacheKey := "analytics:top_products_visits"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetTopProductsVisits(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsPerDay(ctx context.Context) ([]models.VisitsPerDayData, error) {
	var data []models.VisitsPerDayData
	cacheKey := "analytics:visits_per_day"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsPerDay(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsPerMonth(ctx context.Context) ([]models.VisitsPerMonthData, error) {
	var data []models.VisitsPerMonthData
	cacheKey := "analytics:visits_per_month"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsPerMonth(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsPerCountry(ctx context.Context) ([]models.VisitsPerCountryData, error) {
	var data []models.VisitsPerCountryData
	cacheKey := "analytics:visits_per_country"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsPerCountry(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsPerCity(ctx context.Context) ([]models.VisitsPerCityData, error) {
	var data []models.VisitsPerCityData
	cacheKey := "analytics:visits_per_city"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsPerCity(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}


func (s *DashboardAnalysisService) GetVisitsPerCityForToday(ctx context.Context) ([]models.VisitsPerCityData, error) {
	var data []models.VisitsPerCityData
	cacheKey := "analytics:visits_per_city_for_today"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsPerCityForToday(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}


func (s *DashboardAnalysisService) GetVisitsFromEgyptPerDay(ctx context.Context) ([]models.VisitsFromEgyptData, error) {
	var data []models.VisitsFromEgyptData
	cacheKey := "analytics:visits_from_egypt_per_day"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsFromEgyptPerDay(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsFromOtherCountriesPerDay(ctx context.Context) ([]models.VisitsFromCountriesData, error) {
	var data []models.VisitsFromCountriesData
	cacheKey := "analytics:visits_from_other_countries_per_day"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsFromOtherCountriesPerDay(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsFromEgyptPerHourForPastMonth(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error) {
	var data []models.VisitsFromEgyptHoursData
	cacheKey := "analytics:visits_from_egypt_per_hour_past_month"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsFromEgyptPerHourForPastMonth(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}

func (s *DashboardAnalysisService) GetVisitsFromEgyptPerHourForToday(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error) {
	var data []models.VisitsFromEgyptHoursData
	cacheKey := "analytics:visits_from_egypt_per_hour_today"

	if ok, err := utils.GetCache(cacheKey, &data); err == nil && ok {
		return data, nil
	}

	data, err := s.analytticsRepo.GetVisitsFromEgyptPerHourForToday(ctx)
	if err != nil {
		return nil, err
	}

	_ = utils.SetCache(cacheKey, data, 5*time.Minute)
	return data, nil
}
