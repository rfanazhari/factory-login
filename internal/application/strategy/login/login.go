package strategy_login

import (
	"context"

	"github.com/rfanazhari/factory-login/internal/application/dto"
	"github.com/rfanazhari/factory-login/internal/domain/entity"
)

// LoginStrategy defines the interface for login strategies
type LoginStrategy interface {
	ValidateInput(req *dto.LoginRequest) error
	FindUser(ctx context.Context, identifier string) (*entity.User, error)
	GetRateLimitKey(identifier string) string
}
