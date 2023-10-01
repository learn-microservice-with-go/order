package redis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func Insert(uid string, verification_code string) {
	val, err := rdb.Set(ctx, uid, verification_code, 10*time.Minute).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}

func Get(uid string) {
	val, err := rdb.Get(ctx, uid).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}

	Delete(uid)
	fmt.Println(val)
}

func Delete(uid string) {
	val, err := rdb.Del(ctx, uid).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
