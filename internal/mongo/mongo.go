package mongo

import (
	"context"
	"fmt"

	config "github.com/learn-microservice-with-go/order/internal/config"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Provider = wire.NewSet(NewMongo)

func NewMongo(config *config.Config) (*mongo.Client, error) {
	dbAddr := fmt.Sprintf("%s:%s", config.MongoHost, config.MongoPort)

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", dbAddr)))
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return mongoClient, err
}
