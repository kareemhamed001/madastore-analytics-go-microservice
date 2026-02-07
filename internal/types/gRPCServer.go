package types

import (
	"madastore/analytics/common/genproto/analytics"
	handlers "madastore/analytics/internal/handlers/analytics"
	"madastore/analytics/internal/services"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr             string
	server           *grpc.Server
	lis              net.Listener
	analyticsService *services.DashboardAnalysisService
}

func NewGRPCServer(addr string, analyticsService *services.DashboardAnalysisService) *GRPCServer {
	return &GRPCServer{addr: addr, analyticsService: analyticsService}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.lis = lis
	s.server = grpc.NewServer()

	//register services
	analyticsHandlers := handlers.NewAnalyticsGrpcHandler(s.server, s.analyticsService)

	analytics.RegisterAnalyticsServiceServer(s.server, analyticsHandlers)

	log.Info().Str("addr", s.addr).Msg("grpc server is running")

	return s.server.Serve(lis)
}

func (s *GRPCServer) Stop() {
	if s.server == nil {
		return
	}
	s.server.GracefulStop()
}
