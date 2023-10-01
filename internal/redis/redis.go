package redis

import (
	"context"

	"github.com/learn-frame/learn-micro-service/internal/env"
	"github.com/redis/go-redis/v9"
)

var REDIS_PASSWORD = env.GetEnv("REDIS_PASSWORD")

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: REDIS_PASSWORD,
	DB:       0,
})

var ctx = context.Background()
