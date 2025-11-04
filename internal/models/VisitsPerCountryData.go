package models

import "database/sql"

type VisitsPerCountryData struct {
	Country sql.NullString `json:"country"`
	Total   int            `json:"total"`
}
