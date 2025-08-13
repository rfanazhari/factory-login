package http

import (
	"github.com/go-redis/redis/v8"
	strategylogin "github.com/rfanazhari/factory-login/internal/application/strategy/login"
	"github.com/rfanazhari/factory-login/internal/application/usecase"
	"github.com/rfanazhari/factory-login/internal/infrastructure/external/google"
	"github.com/rfanazhari/factory-login/internal/infrastructure/external/jwt"
	externalRedis "github.com/rfanazhari/factory-login/internal/infrastructure/external/redis"
	"github.com/rfanazhari/factory-login/internal/infrastructure/persistence/memory"
	"github.com/rfanazhari/factory-login/internal/interfaces/http/handler"
	"time"
)

// Container holds all dependencies
type Container struct {
	LoginHandler *handler.LoginHandler
}

// NewContainer creates and wires all dependencies
func NewContainer(secretCaptcha string, redisUrl string, maxRateLimit int, maxRateLimitDuration time.Duration, skipCaptcha bool) *Container {
	// Infrastructure
	rdb := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	// Services
	captchaService := google.NewGoogleCaptchaService(secretCaptcha)
	rateLimitService := externalRedis.NewRedisRateLimitService(rdb, maxRateLimit, maxRateLimitDuration)
	tokenService := jwt.NewJWTTokenService()
	googleOAuthService := google.NewGoogleOAuthService()

	// Repository
	userRepo := memory.NewInMemoryUserRepository()

	// Factory
	strategyFactory := strategylogin.NewLoginStrategyFactory(userRepo, googleOAuthService)

	// Use Case
	loginUseCase := usecase.NewLoginUseCase(strategyFactory, captchaService, rateLimitService, tokenService, skipCaptcha)

	// Handler
	loginHandler := handler.NewLoginHandler(loginUseCase)

	return &Container{
		LoginHandler: loginHandler,
	}
}
