package repositories

import (
	"context"
	"database/sql"
	"madastore/analytics/internal/models"
)

type DashboardAnalysisRepository struct {
	db *sql.DB
}

func NewDashboardAnalysisRepository(db *sql.DB) *DashboardAnalysisRepository {
	return &DashboardAnalysisRepository{db: db}
}

func (r *DashboardAnalysisRepository) GetTopProductsVisits(ctx context.Context) ([]models.ProductVisitStat, error) {
	query := `
	SELECT 
    page_url, 
    COUNT(DISTINCT ip) AS total
FROM visits
WHERE page_url IS NOT NULL
  AND page_url LIKE '%product/%'
GROUP BY page_url
ORDER BY total DESC
LIMIT 20;

	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visitStats []models.ProductVisitStat

	for rows.Next() {
		var stat models.ProductVisitStat
		err := rows.Scan(&stat.PageURL, &stat.Total)
		if err != nil {
			return nil, err
		}
		visitStats = append(visitStats, stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return visitStats, nil

}

func (r *DashboardAnalysisRepository) GetVisitsPerDay(ctx context.Context) ([]models.VisitsPerDayData, error) {
	query := `
	SELECT 
			DATE(created_at) AS date,
			COUNT(DISTINCT ip) AS total
		FROM visits
	
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at) ASC
		LIMIT 30;
	
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visitsData []models.VisitsPerDayData
	for rows.Next() {
		var data models.VisitsPerDayData
		err := rows.Scan(&data.Date, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil

}
func (r *DashboardAnalysisRepository) GetVisitsPerMonth(ctx context.Context) ([]models.VisitsPerMonthData, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT 
    DATE_FORMAT(created_at, '%Y-%m') AS month,
    COUNT(DISTINCT ip) AS total
FROM visits
WHERE created_at >= CURDATE() - INTERVAL 12 MONTH
GROUP BY DATE_FORMAT(created_at, '%Y-%m')
ORDER BY month ASC;
`)
	if err != nil {
		return []models.VisitsPerMonthData{}, err
	}
	defer rows.Close()

	var data []models.VisitsPerMonthData

	for rows.Next() {
		var visit models.VisitsPerMonthData
		if err := rows.Scan(&visit.Month, &visit.Total); err != nil {
			return data, err
		}
		data = append(data, visit)
	}
	return data, nil
}

func (r *DashboardAnalysisRepository) GetVisitsPerCountry(ctx context.Context) ([]models.VisitsPerCountryData, error) {
	query := `
	SELECT 
			country,
			COUNT(DISTINCT ip) AS total
		FROM visits
		WHERE country IS NOT NULL AND country != ''
		GROUP BY country
		ORDER BY total DESC
		LIMIT 40;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visitsData []models.VisitsPerCountryData
	for rows.Next() {
		var data models.VisitsPerCountryData
		err := rows.Scan(&data.Country, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil
}

func (r *DashboardAnalysisRepository) GetVisitsPerCity(ctx context.Context) ([]models.VisitsPerCityData, error) {
	query := `	
	SELECT city, COUNT(DISTINCT ip) AS total
		FROM visits
		GROUP BY city
		ORDER BY total DESC
		LIMIT 50
		`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visitsData []models.VisitsPerCityData
	for rows.Next() {
		var data models.VisitsPerCityData
		err := rows.Scan(&data.City, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil
}

func (r *DashboardAnalysisRepository) GetVisitsFromEgyptPerDay(ctx context.Context) ([]models.VisitsFromEgyptData, error) {
	query := `
	SELECT 
			DATE(created_at) AS date,
			COUNT(DISTINCT ip) AS total
		FROM visits
		WHERE country = 'Egypt'
		  
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at) ASC;
		
		`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visitsData []models.VisitsFromEgyptData
	for rows.Next() {
		var data models.VisitsFromEgyptData
		err := rows.Scan(&data.Date, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil
}

func (r *DashboardAnalysisRepository) GetVisitsFromOtherCountriesPerDay(ctx context.Context) ([]models.VisitsFromCountriesData, error) {
	query := `
	SELECT 
			DATE(created_at) AS date,
			COUNT(DISTINCT ip) AS total
		FROM visits
		WHERE (country IS NULL OR country != 'Egypt')
		  
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at) ASC;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visitsData []models.VisitsFromCountriesData
	for rows.Next() {
		var data models.VisitsFromCountriesData
		err := rows.Scan(&data.Date, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil
}

func (r *DashboardAnalysisRepository) GetVisitsFromEgyptPerHourForPastMonth(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error) {
	query := `
	SELECT 
    HOUR(created_at) AS hour, 
    COUNT(DISTINCT ip) AS total
FROM visits
WHERE  country_code = 'EG'
GROUP BY HOUR(created_at)
ORDER BY HOUR(created_at) ASC;

	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var visitsData []models.VisitsFromEgyptHoursData
	for rows.Next() {
		var data models.VisitsFromEgyptHoursData
		err := rows.Scan(&data.Hour, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil
}
func (r *DashboardAnalysisRepository) GetVisitsFromEgyptPerHourForToday(ctx context.Context) ([]models.VisitsFromEgyptHoursData, error) {
	query := `
	SELECT 
			HOUR(created_at) AS hour,
			COUNT(DISTINCT ip) AS total
		FROM visits
		WHERE created_at >= CURDATE()
			AND country_code = 'EG'
		GROUP BY HOUR(created_at)
		ORDER BY HOUR(created_at) ASC;`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var visitsData []models.VisitsFromEgyptHoursData
	for rows.Next() {
		var data models.VisitsFromEgyptHoursData
		err := rows.Scan(&data.Hour, &data.Total)
		if err != nil {
			return nil, err
		}
		visitsData = append(visitsData, data)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return visitsData, nil
}
