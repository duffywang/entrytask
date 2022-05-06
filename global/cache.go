package global

import "github.com/go-redis/redis/v8"


//全局变量，redis客户端
var(
	RedisClient *redis.Client
)