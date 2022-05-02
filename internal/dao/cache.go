package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redisClient *redis.Client
}

func NewRedisClient(redisClient *redis.Client) *RedisClient {
	return &RedisClient{redisClient: redisClient}
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	//TODO:ctx context.Context 出现好多次
	//err 如何更好的处理，打上日志，go日志组件
	value, err := rc.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rc *RedisClient) Set(ctx context.Context, key string, value any, expireTime time.Duration) error {
	err := rc.redisClient.Set(ctx, key, value, expireTime).Err()
	if err != nil {
		fmt.Printf("Redis Set Fail: %v",err)
	}
	return err
}
