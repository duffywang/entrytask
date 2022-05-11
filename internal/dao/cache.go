package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redisClient *redis.Client
}

func NewRedisClient(redisClient *redis.Client) *RedisClient {
	return &RedisClient{redisClient: redisClient}
}

//缓存获取
func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	value, err := rc.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

//缓存存储
func (rc *RedisClient) Set(ctx context.Context, key string, value any, expireTime time.Duration) error {
	start := time.Now()
	err := rc.redisClient.Set(ctx, key, value, expireTime).Err()
	cost := time.Since(start)
	log.Printf("set cost time : %d\n", cost)
	if err != nil {
		fmt.Printf("Redis Set Fail: %v", err)
	}
	return err
}
