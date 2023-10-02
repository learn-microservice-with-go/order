//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	config "github.com/learn-microservice-with-go/order/internal/config"
	grpcServer "github.com/learn-microservice-with-go/order/internal/grpc"
	mongoClient "github.com/learn-microservice-with-go/order/internal/mongo"
	mysqlClient "github.com/learn-microservice-with-go/order/internal/mysql"
	redisClient "github.com/learn-microservice-with-go/order/internal/redis"
)

func InitializeApp() (*App, error) {
	panic(wire.Build(config.Provider, mysqlClient.Provider, mongoClient.Provider, redisClient.Provider, grpcServer.Provider, NewApp))
}
