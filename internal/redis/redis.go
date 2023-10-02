package redis

import (
	"fmt"

	config "github.com/learn-microservice-with-go/user_microservice/internal/config"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var Provider = wire.NewSet(NewRedis)

func NewRedis(config *config.Config) (*redis.Client, error) {
	redisAddr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)
	redisPasswd := config.RedisPassword

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPasswd,
		DB:       0,
	})

	return client, nil
}
