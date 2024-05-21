package model

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
    ctx = context.Background()
    rdb *redis.Client
)

func init()  {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "redis",
		DB: 0,
	})

	// 连接redis数据库的错误判断
	_, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
}