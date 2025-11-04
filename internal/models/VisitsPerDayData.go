package models

import "time"

type VisitsPerDayData struct {
	Date  time.Time `json:"date"`
	Total int       `json:"total"`
}
