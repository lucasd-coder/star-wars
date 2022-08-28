package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
)

var redisServer *miniredis.Miniredis

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	return s
}

func TestSave(t *testing.T) {
	redisServer = mockRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	defer teardown()

	testService := service.CacheService{redisClient}

	err := testService.Save(context.TODO(), "data", "something here", time.Minute)

	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	redisServer = mockRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	defer teardown()

	redisClient.Set(context.TODO(), "data", "something here", time.Minute)

	testService := service.CacheService{redisClient}

	result, err := testService.Get(context.TODO(), "data")

	assert.Nil(t, err)
	assert.Equal(t, result, "something here")
}

func teardown() {
	redisServer.Close()
}
