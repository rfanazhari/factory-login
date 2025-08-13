package strategy_login

import (
	"context"

	"github.com/rfanazhari/factory-login/internal/application/dto"
	"github.com/rfanazhari/factory-login/internal/domain/entity"
)

// LoginStrategy defines the interface for login strategies
type LoginStrategy interface {
	ValidateInput(req *dto.LoginRequest) error
	Authenticate(ctx context.Context, req *dto.LoginRequest) (*entity.User, error)
	GetRateLimitKey(req *dto.LoginRequest) string
}
