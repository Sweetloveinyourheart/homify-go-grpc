package authentication_service

import (
	"fmt"
	"log"
	"net"

	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/authentication-service/configs"
	"homify-go-grpc/internal/authentication-service/database"
	"homify-go-grpc/internal/authentication-service/server"

	"google.golang.org/grpc"
)

func RunGRPCAuthenticationServer() {
	configurations := configs.GetConfig()

	db := database.InitPostgresConnection()

	lis, err := net.Listen("tcp", configurations.TCPAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.NewGRPCAuthenticationServer(db)
	proto.RegisterAuthenticationServer(s, srv)

	fmt.Printf("🚀 Authentication Server launched on port %s ... \n", configurations.TCPAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
