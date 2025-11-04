package main

import (
	"log"
	"madastore/analytics/internal/config"
	"madastore/analytics/internal/handlers"
	repositories "madastore/analytics/internal/repository"
	"madastore/analytics/internal/services"
	"net/http"

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

	log.Printf("Starting server on port %s...", config.ServerPort)
	if err := router.Run("0.0.0.0:" + config.ServerPort); err != nil {
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

	api := router.Group("/api/v1")

	api.GET("/stats", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Stats endpoint"})
	})
	api.GET("analytics", analyticsHandlers.GetDashboardData)
	api.GET("analytics/top-products-visits", analyticsHandlers.GetTopProductsVisits)
	api.GET("analytics/visits-per-day", analyticsHandlers.GetVisitsPerDay)
	api.GET("analytics/visits-per-month", analyticsHandlers.GetVisitsPerMonth)
	api.GET("analytics/visits-per-country", analyticsHandlers.GetVisitsPerCountry)
	api.GET("analytics/visits-per-city", analyticsHandlers.GetVisitsPerCity)
	api.GET("analytics/visits-from-egypt-per-day", analyticsHandlers.GetVisitsFromEgyptPerDay)
	api.GET("analytics/visits-from-other-countries-per-day", analyticsHandlers.GetVisitsFromOtherCountriesPerDay)
	api.GET("analytics/visits-from-egypt-per-hour-past-month", analyticsHandlers.GetVisitsFromEgyptPerHourForPastMonth)
	api.GET("analytics/visits-from-egypt-per-hour-today", analyticsHandlers.GetVisitsFromEgyptPerHourForToday)

}
