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

func (s *EmailLoginStrategy) FindUser(ctx context.Context, identifier string) (*entity.User, error) {
	email, err := valueobject.NewEmail(identifier)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindByEmail(ctx, *email)
}

func (s *EmailLoginStrategy) GetRateLimitKey(identifier string) string {
	return fmt.Sprintf("login:email:%s", identifier)
}
