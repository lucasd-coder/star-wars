package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheService struct {
	Client *redis.Client
}

func NewCacheService(redisClient *redis.Client) *CacheService {
	return &CacheService{
		Client: redisClient,
	}
}

func (cacheService *CacheService) Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	r := cacheService.Client.Set(ctx, key, val, ttl)
	_, err = r.Result()

	if err != nil {
		return err
	}

	return nil
}

func (cacheService *CacheService) Get(ctx context.Context, key string) (string, error) {
	return cacheService.Client.Get(ctx, key).Result()
}
