package main

import (
	"log"

	config "github.com/learn-microservice-with-go/user_microservice/internal/config"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type App struct {
	config      *config.Config
	mysqlClient *gorm.DB
	redisClient *redis.Client
	grpcServer  *grpc.Server
}

func NewApp(config *config.Config, mysqlClient *gorm.DB, redisClient *redis.Client, grpcServer *grpc.Server) *App {
	return &App{config: config, mysqlClient: mysqlClient, redisClient: redisClient, grpcServer: grpcServer}
}

func main() {
	app, err := InitializeApp()

	if err != nil {
		log.Fatal(err)
	}

	println(app)
}
