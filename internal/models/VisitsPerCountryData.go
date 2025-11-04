package models

import "madastore/analytics/internal/utils"

type VisitsPerCountryData struct {
	Country utils.NullString `json:"country"`
	Total   int              `json:"total"`
}
