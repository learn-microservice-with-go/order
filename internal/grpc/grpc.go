package grpc

import (
	"log"
	"net"

	proto "github.com/learn-microservice-with-go/user_microservice/api/service"
	service "github.com/learn-microservice-with-go/user_microservice/internal/service"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(NewGrpc)

func NewGrpc(mysqlClient *gorm.DB, mongoClient *mongo.Client, redisClient *redis.Client) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := service.Server{
		Db:    mysqlClient,
		Mongo: mongoClient,
		Rds:   redisClient,
	}

	proto.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v", err)
	}

	return grpcServer, nil
}
