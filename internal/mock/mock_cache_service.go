package mock

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockCacheService struct {
	mock.Mock
}

func (mock *MockCacheService) Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	args := mock.Called(ctx, key, value, ttl)

	var r0 error
	if rf, ok := args.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (mock *MockCacheService) Get(ctx context.Context, key string) (string, error) {
	args := mock.Called(ctx, key)

	var r0 string
	if rf, ok := args.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
