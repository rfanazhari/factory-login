package usecase

import (
	"context"
	"github.com/rfanazhari/factory-login/internal/application/dto"
	strategy_login "github.com/rfanazhari/factory-login/internal/application/strategy/login"
	"github.com/rfanazhari/factory-login/internal/domain/service"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// LoginUseCase orchestrates the login process
type LoginUseCase struct {
	strategyFactory  *strategy_login.LoginStrategyFactory
	captchaService   service.CaptchaService
	rateLimitService service.RateLimitService
	tokenService     service.TokenService
	skipCaptcha      bool
}

func NewLoginUseCase(
	strategyFactory *strategy_login.LoginStrategyFactory,
	captchaService service.CaptchaService,
	rateLimitService service.RateLimitService,
	tokenService service.TokenService,
	skipCaptcha bool,
) *LoginUseCase {
	return &LoginUseCase{
		strategyFactory:  strategyFactory,
		captchaService:   captchaService,
		rateLimitService: rateLimitService,
		tokenService:     tokenService,
		skipCaptcha:      skipCaptcha,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. Validate captcha
	if !uc.skipCaptcha {
		captchaCode, err := valueobject.NewCaptchaCode(req.CaptchaCode)
		if err != nil {
			return &dto.LoginResponse{Success: false, Message: err.Error()}, nil
		}

		if err := uc.captchaService.Validate(ctx, *captchaCode); err != nil {
			return &dto.LoginResponse{Success: false, Message: "captcha validation failed"}, nil
		}
	}

	// 2. Create login strategy using factory
	loginStrategy, err := uc.strategyFactory.CreateStrategy(req.Type)
	if err != nil {
		return &dto.LoginResponse{Success: false, Message: "invalid login type"}, nil
	}

	// 3. Validate input format
	if err := loginStrategy.ValidateInput(req); err != nil {
		return &dto.LoginResponse{Success: false, Message: "invalid input format"}, nil
	}

	// 4. Check rate limit
	rateLimitKey := loginStrategy.GetRateLimitKey(req.Identifier)
	if err := uc.rateLimitService.CheckLimit(ctx, rateLimitKey); err != nil {
		return &dto.LoginResponse{Success: false, Message: "too many attempts, please try again later"}, nil
	}

	// 5. Find user
	foundUser, err := loginStrategy.FindUser(ctx, req.Identifier)
	if err != nil {
		uc.rateLimitService.IncrementAttempt(ctx, rateLimitKey)
		return &dto.LoginResponse{Success: false, Message: "invalid credentials"}, nil
	}

	// 6. Check if user is active
	if !foundUser.IsActive() {
		uc.rateLimitService.IncrementAttempt(ctx, rateLimitKey)
		return &dto.LoginResponse{Success: false, Message: "account is inactive"}, nil
	}

	// 7. Verify password
	if !foundUser.VerifyPassword(req.Password) {
		uc.rateLimitService.IncrementAttempt(ctx, rateLimitKey)
		return &dto.LoginResponse{Success: false, Message: "invalid credentials"}, nil
	}

	// 8. Generate authentication token
	authentication, err := uc.tokenService.GenerateAccessToken(foundUser.ID())
	if err != nil {
		return &dto.LoginResponse{Success: false, Message: "failed to generate token"}, nil
	}

	return &dto.LoginResponse{
		Success:     true,
		Message:     "login successful",
		AccessToken: authentication.AccessToken(),
		UserID:      authentication.UserID().String(),
		ExpiresAt:   authentication.ExpiresAt().Unix(),
	}, nil
}
