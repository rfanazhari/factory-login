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

func (s *MSISDNLoginStrategy) FindUser(ctx context.Context, identifier string) (*entity.User, error) {
	msisdn, err := valueobject.NewMSISDN(identifier)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindByMSISDN(ctx, *msisdn)
}

func (s *MSISDNLoginStrategy) GetRateLimitKey(identifier string) string {
	return fmt.Sprintf("login:msisdn:%s", identifier)
}
