package strategy_login

import (
	"context"
	"fmt"

	"github.com/rfanazhari/factory-login/internal/application/dto"
	"github.com/rfanazhari/factory-login/internal/domain/entity"
	"github.com/rfanazhari/factory-login/internal/domain/repository"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// EmailLoginStrategy handles email-based login
type EmailLoginStrategy struct {
	userRepo repository.UserRepository
}

func NewEmailLoginStrategy(userRepo repository.UserRepository) *EmailLoginStrategy {
	return &EmailLoginStrategy{userRepo: userRepo}
}

func (s *EmailLoginStrategy) ValidateInput(req *dto.LoginRequest) error {
	_, err := valueobject.NewEmail(req.Identifier)
	return err
}

func (s *EmailLoginStrategy) Authenticate(ctx context.Context, req *dto.LoginRequest) (*entity.User, error) {
	email, err := valueobject.NewEmail(req.Identifier)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.FindByEmail(ctx, *email)
	if err != nil {
		return nil, err
	}
	if !user.VerifyPassword(req.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

func (s *EmailLoginStrategy) GetRateLimitKey(req *dto.LoginRequest) string {
	return fmt.Sprintf("login:email:%s", req.Identifier)
}
