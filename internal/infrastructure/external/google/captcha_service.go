package google

import (
	"context"
	"github.com/rfanazhari/factory-login/internal/domain"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// CaptchaService GoogleCaptchaService implements CaptchaService interface
type CaptchaService struct {
	secretKey string
}

func NewGoogleCaptchaService(secretKey string) *CaptchaService {
	return &CaptchaService{secretKey: secretKey}
}

func (s *CaptchaService) Validate(ctx context.Context, code valueobject.CaptchaCode) error {
	// Mock implementation - in real scenario, call Google reCAPTCHA API
	if code.String() == "invalid" {
		return domain.ErrInvalidCaptcha
	}
	return nil
}
