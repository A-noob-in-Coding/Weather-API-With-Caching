package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func redisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Docker exposes Redis here
	})
	return rdb
}

func CheckInCache(location string) string {
	rdb := redisClient()
	val, err := rdb.Get(Ctx, location).Result()
	if err != nil {
		fmt.Println("Cache Miss")
		return ""
	}
	fmt.Println("Cache Hit")
	return val
}

func StoreInCache(body []byte, location string) {
	rdb := redisClient()
	rdb.Set(Ctx, location, string(body), time.Hour)
}
