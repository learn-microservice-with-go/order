package grpc

import (
	"log"
	"net"

	proto "github.com/learn-frame/learn-micro-service/api/service"
	service "github.com/learn-frame/learn-micro-service/internal/service"
	"google.golang.org/grpc"
)

func NewGrpc() {

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s := service.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v", err)
	}
}
