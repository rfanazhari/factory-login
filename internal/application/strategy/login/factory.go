package strategy_login

import (
	"errors"

	"github.com/rfanazhari/factory-login/internal/domain/repository"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// LoginStrategyFactory creates appropriate login strategy based on type
type LoginStrategyFactory struct {
	userRepo repository.UserRepository
}

func NewLoginStrategyFactory(userRepo repository.UserRepository) *LoginStrategyFactory {
	return &LoginStrategyFactory{userRepo: userRepo}
}

func (f *LoginStrategyFactory) CreateStrategy(loginType valueobject.LoginType) (LoginStrategy, error) {
	switch loginType {
	case valueobject.LoginTypeMSISDN:
		return NewMSISDNLoginStrategy(f.userRepo), nil
	case valueobject.LoginTypeEmail:
		return NewEmailLoginStrategy(f.userRepo), nil
	default:
		return nil, errors.New("unsupported login type")
	}
}
