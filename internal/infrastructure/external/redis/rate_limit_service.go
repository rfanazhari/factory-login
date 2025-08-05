package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rfanazhari/factory-login/internal/domain"
	"time"
)

// RateLimitService RedisRateLimitService implements RateLimitService interface
type RateLimitService struct {
	client   *redis.Client
	maxTries int
	window   time.Duration
}

func NewRedisRateLimitService(client *redis.Client, maxTries int, window time.Duration) *RateLimitService {
	return &RateLimitService{
		client:   client,
		maxTries: maxTries,
		window:   window,
	}
}

func (s *RateLimitService) CheckLimit(ctx context.Context, key string) error {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if val != "" {
		var attempts int
		fmt.Sscanf(val, "%d", &attempts)
		if attempts >= s.maxTries {
			return domain.ErrRateLimitExceeded
		}
	}
	return nil
}

func (s *RateLimitService) IncrementAttempt(ctx context.Context, key string) error {
	pipe := s.client.Pipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, s.window)
	_, err := pipe.Exec(ctx)
	return err
}
