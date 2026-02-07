package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"madastore/analytics/internal/config"
	"madastore/analytics/internal/handlers"
	"madastore/analytics/internal/metrics"
	"madastore/analytics/internal/middleware"
	repositories "madastore/analytics/internal/repository"
	"madastore/analytics/internal/services"
	"madastore/analytics/internal/types"
	"madastore/analytics/internal/utils"
	"madastore/analytics/internal/worker"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

type App struct {
	cfg        *config.Config
	db         *sql.DB
	httpServer *http.Server
	grpcServer *types.GRPCServer
	workers    []worker.Worker
}

func New(cfg *config.Config) (*App, error) {
	InitLogger(cfg.Environment)

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	if cfg.DatabaseDSN == "" {
		return nil, errors.New("DATABASE_DSN is not set")
	}

	db, err := sql.Open("mysql", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	if err := utils.InitCache(
		cfg.RedisAddr,
		cfg.RedisPassword,
		cfg.RedisDB,
		2*time.Second,
		2*time.Second,
		2*time.Second,
	); err != nil {
		log.Warn().Err(err).Msg("cache disabled")
	}

	repo := repositories.NewDashboardAnalysisRepository(db, cfg.RepoQueryTimeout)
	analyticsService := services.NewDashboardAnalysisService(repo)
	analyticsHandlers := handlers.NewAnalyticsHandlers(analyticsService)

	router := gin.New()
	router.Use(
		gin.Recovery(),
		middleware.RequestIDMiddleware(),
		middleware.RequestLoggerMiddleware(),
		middleware.RequestTimeoutMiddleware(cfg.RequestTimeout),
		metrics.HTTPMetricsMiddleware(),
	)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api := router.Group("/api/v1", middleware.CheckApiKeyMiddleware(cfg.ApiKey))
	api.GET("/stats", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Stats endpoint"})
	})
	api.GET("analytics", analyticsHandlers.GetDashboardData)
	api.GET("analytics/top-products-visits", analyticsHandlers.GetTopProductsVisits)
	api.GET("analytics/visits-per-day", analyticsHandlers.GetVisitsPerDay)
	api.GET("analytics/visits-per-month", analyticsHandlers.GetVisitsPerMonth)
	api.GET("analytics/visits-per-country", analyticsHandlers.GetVisitsPerCountry)
	api.GET("analytics/visits-per-city", analyticsHandlers.GetVisitsPerCity)
	api.GET("analytics/visits-per-city-today", analyticsHandlers.GetVisitsPerCityForToday)
	api.GET("analytics/visits-from-egypt-per-day", analyticsHandlers.GetVisitsFromEgyptPerDay)
	api.GET("analytics/visits-from-other-countries-per-day", analyticsHandlers.GetVisitsFromOtherCountriesPerDay)
	api.GET("analytics/visits-from-egypt-per-hour-past-month", analyticsHandlers.GetVisitsFromEgyptPerHourForPastMonth)
	api.GET("analytics/visits-from-egypt-per-hour-today", analyticsHandlers.GetVisitsFromEgyptPerHourForToday)

	httpServer := &http.Server{
		Addr:              "0.0.0.0:" + cfg.ServerPort,
		Handler:           router,
		ReadTimeout:       cfg.HTTPReadTimeout,
		WriteTimeout:      cfg.HTTPWriteTimeout,
		IdleTimeout:       cfg.HTTPIdleTimeout,
		ReadHeaderTimeout: cfg.HTTPReadHeaderTimeout,
	}

	grpcServer := types.NewGRPCServer(":"+cfg.GRPCPort, analyticsService)

	workers := []worker.Worker{
		worker.NewCacheRefresher(analyticsService, cfg.CacheRefreshInterval, 15*time.Second),
	}

	return &App{
		cfg:        cfg,
		db:         db,
		httpServer: httpServer,
		grpcServer: grpcServer,
		workers:    workers,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	for _, w := range a.workers {
		go w.Start(ctx)
	}

	go func() {
		if err := a.grpcServer.Run(); err != nil {
			log.Error().Err(err).Msg("grpc server failed")
		}
	}()

	log.Info().Str("port", a.cfg.ServerPort).Msg("http server starting")
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	for _, w := range a.workers {
		w.Stop()
	}

	a.grpcServer.Stop()
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("http shutdown error")
	}

	if err := utils.CloseCache(); err != nil {
		log.Error().Err(err).Msg("failed to close cache")
	}

	if err := a.db.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close db")
	}

	return nil
}
