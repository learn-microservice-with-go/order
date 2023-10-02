package main

import (
	"log"

	config "github.com/learn-microservice-with-go/order/internal/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type App struct {
	config      *config.Config
	mysqlClient *gorm.DB
	mongoClint  *mongo.Client
	redisClient *redis.Client
	grpcServer  *grpc.Server
}

func NewApp(config *config.Config, mysqlClient *gorm.DB, mongoClint *mongo.Client, redisClient *redis.Client, grpcServer *grpc.Server) *App {
	return &App{config: config, mysqlClient: mysqlClient, mongoClint: mongoClint, redisClient: redisClient, grpcServer: grpcServer}
}

func main() {
	app, err := InitializeApp()

	if err != nil {
		log.Fatal(err)
	}

	println(app)
}
