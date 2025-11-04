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
func (s *DashboardAnalysisService) GetTopProductsVisits(ctx context.Context) (models.DashboardData, error) {
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
