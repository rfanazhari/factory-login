package strategy_login

import (
	"context"
	"fmt"

	"github.com/rfanazhari/factory-login/internal/application/dto"
	"github.com/rfanazhari/factory-login/internal/domain/entity"
	"github.com/rfanazhari/factory-login/internal/domain/repository"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// MSISDNLoginStrategy handles MSISDN-based login
type MSISDNLoginStrategy struct {
	userRepo repository.UserRepository
}

func NewMSISDNLoginStrategy(userRepo repository.UserRepository) *MSISDNLoginStrategy {
	return &MSISDNLoginStrategy{userRepo: userRepo}
}

func (s *MSISDNLoginStrategy) ValidateInput(req *dto.LoginRequest) error {
	_, err := valueobject.NewMSISDN(req.Identifier)
	return err
}

func (s *MSISDNLoginStrategy) Authenticate(ctx context.Context, req *dto.LoginRequest) (*entity.User, error) {
	msisdn, err := valueobject.NewMSISDN(req.Identifier)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.FindByMSISDN(ctx, *msisdn)
	if err != nil {
		return nil, err
	}
	if !user.VerifyPassword(req.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

func (s *MSISDNLoginStrategy) GetRateLimitKey(req *dto.LoginRequest) string {
	return fmt.Sprintf("login:msisdn:%s", req.Identifier)
}
