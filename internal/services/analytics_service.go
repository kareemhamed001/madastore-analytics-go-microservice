package services

import (
	"context"
	"madastore/analytics/internal/models"
	repositories "madastore/analytics/internal/repository/interfaces"

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

	// create an errgroup with context
	g, ctx := errgroup.WithContext(ctx)

	// each call runs in its own goroutine
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

	// wait for all goroutines
	if err := g.Wait(); err != nil {
		return data, err
	}

	return data, nil
}
func (s *DashboardAnalysisService) GetTopProductsVisits(ctx context.Context) ([]models.ProductVisitStat, error) {
	return s.analytticsRepo.GetTopProductsVisits(ctx)
}

func (s *DashboardAnalysisService) GetVisitsPerDay(ctx context.Context) ([]models.VisitsPerDayData, error) {
	return s.analytticsRepo.GetVisitsPerDay(ctx)
}

func (s *DashboardAnalysisService) GetVisitsPerMonth(ctx context.Context) ([]models.VisitsPerMonthData, error) {
	return s.analytticsRepo.GetVisitsPerMonth(ctx)
}

func (s *DashboardAnalysisService) GetVisitsPerCountry(ctx context.Context) ([]models.VisitsPerCountryData, error) {
	return s.analytticsRepo.GetVisitsPerCountry(ctx)
}

func (s *DashboardAnalysisService) GetVisitsPerCity(ctx context.Context) ([]models.VisitsPerCityData, error) {
	return s.analytticsRepo.GetVisitsPerCity(ctx)
}

func (s *DashboardAnalysisService) GetVisitsFromEgyptPerDay(ctx context.Context) ([]models.VisitsFromEgyptData, error) {
	return s.analytticsRepo.GetVisitsFromEgyptPerDay(ctx)
}

func (s *DashboardAnalysisService) GetVisitsFromOtherCountriesPerDay(ctx context.Context) ([]models.VisitsFromCountriesData, error) {
	return s.analytticsRepo.GetVisitsFromOtherCountriesPerDay(ctx)
}

func (s *DashboardAnalysisService) GetVisitsFromEgyptPerHourForPastMonth(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error) {
	return s.analytticsRepo.GetVisitsFromEgyptPerHourForPastMonth(ctx)
}
func (s *DashboardAnalysisService) GetVisitsFromEgyptPerHourForToday(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error) {
	return s.analytticsRepo.GetVisitsFromEgyptPerHourForToday(ctx)
}
