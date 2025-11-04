package models

type DashboardData struct {
	TopProducts                  []ProductVisitStat         `json:"top_products"`
	PerDay                       []VisitsPerDayData         `json:"per_day"`
	PerMonth                     []VisitsPerMonthData       `json:"per_month"`
	PerCountry                   []VisitsPerCountryData     `json:"per_country"`
	PerCity                      []VisitsPerCityData        `json:"per_city"`
	FromEgyptPerDay              []VisitsFromEgyptData      `json:"from_egypt_per_day"`
	FromOtherCountriesPerDay     []VisitsFromCountriesData  `json:"from_other_countries_per_day"`
	FromEgyptPerHourForPastMonth []VisitsFromEgyptHoursData `json:"from_egypt_per_hour_past_month"`
	FromEgyptPerHourForToday     []VisitsFromEgyptHoursData `json:"from_egypt_per_hour_today"`
}
