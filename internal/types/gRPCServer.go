package types

import (
	"database/sql"
	"log"
	"madastore/analytics/common/genproto/analytics"
	handlers "madastore/analytics/internal/handlers/analytics"
	repositories "madastore/analytics/internal/repository"
	"madastore/analytics/internal/services"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run(db *sql.DB) error {

	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}
	server := grpc.NewServer()

	//register services
	analyticsRepo := repositories.NewDashboardAnalysisRepository(db)
	analyticsService := services.NewDashboardAnalysisService(analyticsRepo)
	analyticsHandlers := handlers.NewAnalyticsGrpcHandler(server, analyticsService)

	analytics.RegisterAnalyticsServiceServer(server, analyticsHandlers)

	log.Println("gRPC server is running on", s.addr)

	return server.Serve(lis)
}
