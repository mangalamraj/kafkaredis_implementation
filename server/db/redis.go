package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func SetCache(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, json, expiration).Err()
}

func GetCache(key string) (string, error) {
	ctx := context.Background()
	return redisClient.Get(ctx, key).Result()
} 