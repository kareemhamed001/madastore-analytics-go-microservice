package main

import (
	"context"
	"log"
	"madastore/analytics/internal/config"
	"madastore/analytics/internal/handlers"
	"madastore/analytics/internal/middleware"
	repositories "madastore/analytics/internal/repository"
	"madastore/analytics/internal/services"
	"madastore/analytics/internal/types"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := config.Load()

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.DatabaseDSN == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("mysql", config.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	registerRoutes(router, db)

	grpcServer := types.NewGRPCServer(":" + config.GRPCPort)
	go func() {
		if err := grpcServer.Run(db); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	log.Printf("Starting server on port %s...", config.ServerPort)

	srv := &http.Server{
		Addr:              "0.0.0.0:" + config.ServerPort,
		Handler:           router,
		ReadTimeout:       120 * time.Second, // max time to read request
		WriteTimeout:      120 * time.Second, // max time to write response
		IdleTimeout:       60 * time.Second,  // max keep-alive time
		ReadHeaderTimeout: 5 * time.Second,   // header read timeout
	}

	// Start the server

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}

func registerRoutes(router *gin.Engine, db *sql.DB) {

	analyticsRepo := repositories.NewDashboardAnalysisRepository(db)
	analyticsService := services.NewDashboardAnalysisService(analyticsRepo)
	analyticsHandlers := handlers.NewAnalyticsHandlers(analyticsService)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	api := router.Group("/api/v1", middleware.CheckApiKeyMiddleware)

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

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for {
			<-ticker.C
			log.Println("⏱  Refreshing analytics cache...")

			ctx := context.Background()

			// call all the service methods to refresh caches
			if _, err := analyticsService.GetTopProductsVisits(ctx); err != nil {
				log.Println("Error refreshing top products:", err)
			}
			if _, err := analyticsService.GetVisitsPerDay(ctx); err != nil {
				log.Println("Error refreshing visits per day:", err)
			}
			if _, err := analyticsService.GetVisitsPerMonth(ctx); err != nil {
				log.Println("Error refreshing visits per month:", err)
			}
			if _, err := analyticsService.GetVisitsPerCountry(ctx); err != nil {
				log.Println("Error refreshing visits per country:", err)
			}
			if _, err := analyticsService.GetVisitsFromEgyptPerDay(ctx); err != nil {
				log.Println("Error refreshing visits from Egypt:", err)
			}
			if _, err := analyticsService.GetVisitsFromOtherCountriesPerDay(ctx); err != nil {
				log.Println("Error refreshing visits from other countries:", err)
			}
		}
	}()

}
