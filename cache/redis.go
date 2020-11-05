package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

// redis缓存客户端
var RedisClient *redis.Client

func Redis(addr string, pwd string, db int) {
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	_, err := redis.Ping().Result()
	if err != nil {
		fmt.Printf("redis connect failed, err: %v", err)
	}

	RedisClient = redis
}
