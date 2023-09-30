package main

import (
	"log"
	"net"

	proto "github.com/learn-frame/learn-micro-service/api/service"
	service "github.com/learn-frame/learn-micro-service/internal/service"
	"google.golang.org/grpc"
)

func main() {

	// 监听 12345 端口, 返回一个 listener 和 error
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s := service.Server{}

	grpcServer := grpc.NewServer()

	//注册一个server
	proto.RegisterUserServiceServer(grpcServer, &s)

	// server 开始监听
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v", err)
	}

}
