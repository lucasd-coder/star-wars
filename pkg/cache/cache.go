package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/lucasd-coder/star-wars/config"
)

var client *redis.Client

func SetUpRedis(cfg *config.Config) {
	url := fmt.Sprintf("%s:%d", cfg.RedisUrl, cfg.RedisPort)

	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
		DB:   cfg.RedisDb,
		Password: cfg.RedisPassword,
	})

	client = redisClient
}

func GetClient() *redis.Client {
	return client
}
