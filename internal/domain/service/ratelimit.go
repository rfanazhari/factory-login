package service

import "context"

// RateLimitService defines the interface for rate limiting
type RateLimitService interface {
	CheckLimit(ctx context.Context, key string) error
	IncrementAttempt(ctx context.Context, key string) error
}
