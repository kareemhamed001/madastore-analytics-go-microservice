package worker

import (
	"context"
	"time"

	"madastore/analytics/internal/services"

	"github.com/rs/zerolog/log"
)

type CacheRefresher struct {
	service        *services.DashboardAnalysisService
	interval       time.Duration
	refreshTimeout time.Duration
	stopCh         chan struct{}
}

func NewCacheRefresher(service *services.DashboardAnalysisService, interval, refreshTimeout time.Duration) *CacheRefresher {
	return &CacheRefresher{
		service:        service,
		interval:       interval,
		refreshTimeout: refreshTimeout,
		stopCh:         make(chan struct{}),
	}
}

func (w *CacheRefresher) Start(ctx context.Context) {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-w.stopCh:
			return
		case <-ticker.C:
			log.Info().Msg("refreshing analytics cache")
			refreshCtx, cancel := context.WithTimeout(context.Background(), w.refreshTimeout)
			if _, err := w.service.GetTopProductsVisits(refreshCtx); err != nil {
				log.Warn().Err(err).Msg("refresh top products failed")
			}
			if _, err := w.service.GetVisitsPerDay(refreshCtx); err != nil {
				log.Warn().Err(err).Msg("refresh visits per day failed")
			}
			if _, err := w.service.GetVisitsPerMonth(refreshCtx); err != nil {
				log.Warn().Err(err).Msg("refresh visits per month failed")
			}
			if _, err := w.service.GetVisitsPerCountry(refreshCtx); err != nil {
				log.Warn().Err(err).Msg("refresh visits per country failed")
			}
			if _, err := w.service.GetVisitsFromEgyptPerDay(refreshCtx); err != nil {
				log.Warn().Err(err).Msg("refresh visits from Egypt failed")
			}
			if _, err := w.service.GetVisitsFromOtherCountriesPerDay(refreshCtx); err != nil {
				log.Warn().Err(err).Msg("refresh visits from other countries failed")
			}
			cancel()
		}
	}
}

func (w *CacheRefresher) Stop() {
	select {
	case <-w.stopCh:
		return
	default:
		close(w.stopCh)
	}
}
