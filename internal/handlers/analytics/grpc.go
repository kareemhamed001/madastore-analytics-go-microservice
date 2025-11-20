package handlers

import (
	"context"
	"madastore/analytics/common/genproto/analytics"
	"madastore/analytics/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AnalyticsGrpcHandler struct {
	analyticsService *services.DashboardAnalysisService
	analytics.UnimplementedAnalyticsServiceServer
}

func NewAnalyticsGrpcHandler(server *grpc.Server, service *services.DashboardAnalysisService) *AnalyticsGrpcHandler {
	handler := &AnalyticsGrpcHandler{
		analyticsService: service,
	}

	return handler
}

func (handler *AnalyticsGrpcHandler) TopProductsVisits(ctx context.Context, req *analytics.TopProductsVisitsRequest) (*analytics.TopProductsVisitsResponse, error) {
	data, err := handler.analyticsService.GetTopProductsVisits(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get top products visits: %v", err)
	}

	response := &analytics.TopProductsVisitsResponse{}
	for _, item := range data {
		response.ProductVisits = append(response.ProductVisits, &analytics.ProductVisit{
			PageURL: item.PageURL,
			Total:   int32(item.Total),
		})
	}

	return response, nil
}
func (handler *AnalyticsGrpcHandler) VisitsPerDay(ctx context.Context, req *analytics.VisitsPerDayRequest) (*analytics.VisitsPerDayResponse, error) {
	data, err := handler.analyticsService.GetVisitsPerDay(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits per day: %v", err)
	}
	response := &analytics.VisitsPerDayResponse{}
	for _, item := range data {
		response.DayVisits = append(response.DayVisits, &analytics.DayVisit{
			Date:  item.Date.Format("2006-01-02"),
			Total: int32(item.Total),
		})
	}

	return response, nil
}

func (handler *AnalyticsGrpcHandler) VisitsPerMonth(ctx context.Context, req *analytics.VisitsPerMonthRequest) (*analytics.VisitsPerMonthResponse, error) {
	data, err := handler.analyticsService.GetVisitsPerMonth(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits per month: %v", err)
	}
	response := &analytics.VisitsPerMonthResponse{}
	for _, item := range data {
		response.MonthVisits = append(response.MonthVisits, &analytics.MonthVisit{
			Month: item.Month,
			Total: int32(item.Total),
		})
	}

	return response, nil
}
func (handler *AnalyticsGrpcHandler) VisitsPerCountry(ctx context.Context, req *analytics.VisitsPerCountryRequest) (*analytics.VisitsPerCountryResponse, error) {
	data, err := handler.analyticsService.GetVisitsPerCountry(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits per country: %v", err)
	}
	response := &analytics.VisitsPerCountryResponse{}
	for _, item := range data {
		response.CountryVisits = append(response.CountryVisits, &analytics.CountryVisit{
			Country: item.Country.String,
			Total:   int32(item.Total),
		})
	}

	return response, nil
}

func (handler *AnalyticsGrpcHandler) VisitsPerCity(ctx context.Context, req *analytics.VisitsPerCityRequest) (*analytics.VisitsPerCityResponse, error) {
	data, err := handler.analyticsService.GetVisitsPerCity(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits per city: %v", err)
	}
	response := &analytics.VisitsPerCityResponse{}
	for _, item := range data {
		response.CityVisits = append(response.CityVisits, &analytics.CityVisit{
			City:  item.City,
			Total: int32(item.Total),
		})
	}

	return response, nil
}
func (handler *AnalyticsGrpcHandler) VisitsFromEgyptPerDay(ctx context.Context, req *analytics.VisitsFromEgyptPerDayRequest) (*analytics.VisitsFromEgyptPerDayResponse, error) {
	data, err := handler.analyticsService.GetVisitsFromEgyptPerDay(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits from Egypt per day: %v", err)
	}
	response := &analytics.VisitsFromEgyptPerDayResponse{}
	for _, item := range data {
		response.DayVisits = append(response.DayVisits, &analytics.DayVisit{
			Date:  item.Date,
			Total: int32(item.Total),
		})
	}

	return response, nil
}
func (handler *AnalyticsGrpcHandler) VisitsFromOtherCountriesPerDay(ctx context.Context, req *analytics.VisitsFromOtherCountriesPerDayRequest) (*analytics.VisitsFromOtherCountriesPerDayResponse, error) {
	data, err := handler.analyticsService.GetVisitsFromOtherCountriesPerDay(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits from other countries per day: %v", err)
	}
	response := &analytics.VisitsFromOtherCountriesPerDayResponse{}
	for _, item := range data {
		response.DayVisits = append(response.DayVisits, &analytics.DayVisit{
			Date:  item.Date,
			Total: int32(item.Total),
		})
	}

	return response, nil
}
func (handler *AnalyticsGrpcHandler) VisitsFromEgyptPerHourPastMonth(ctx context.Context, req *analytics.VisitsFromEgyptPerHourPastMonthRequest) (*analytics.VisitsFromEgyptPerHourPastMonthResponse, error) {
	data, err := handler.analyticsService.GetVisitsFromEgyptPerHourForPastMonth(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits from Egypt per hour past month: %v", err)
	}
	response := &analytics.VisitsFromEgyptPerHourPastMonthResponse{}
	for _, item := range data {
		response.HourVisits = append(response.HourVisits, &analytics.HourVisit{
			Hour:  item.Hour,
			Total: int32(item.Total),
		})
	}

	return response, nil
}
func (handler *AnalyticsGrpcHandler) VisitsFromEgyptPerHourToday(ctx context.Context, req *analytics.VisitsFromEgyptPerHourTodayRequest) (*analytics.VisitsFromEgyptPerHourTodayResponse, error) {
	data, err := handler.analyticsService.GetVisitsFromEgyptPerHourForToday(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get visits from Egypt per hour today: %v", err)
	}
	response := &analytics.VisitsFromEgyptPerHourTodayResponse{}
	for _, item := range data {
		response.HourVisits = append(response.HourVisits, &analytics.HourVisit{
			Hour:  item.Hour,
			Total: int32(item.Total),
		})
	}

	return response, nil
}
