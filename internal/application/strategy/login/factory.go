package strategy_login

import (
	"errors"

	"github.com/rfanazhari/factory-login/internal/domain/repository"
	"github.com/rfanazhari/factory-login/internal/domain/service"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// LoginStrategyFactory creates appropriate login strategy based on type
type LoginStrategyFactory struct {
	userRepo       repository.UserRepository
	googleOAuthSvc service.GoogleOAuthService
}

func NewLoginStrategyFactory(userRepo repository.UserRepository, googleSvc service.GoogleOAuthService) *LoginStrategyFactory {
	return &LoginStrategyFactory{userRepo: userRepo, googleOAuthSvc: googleSvc}
}

func (f *LoginStrategyFactory) CreateStrategy(loginType valueobject.LoginType) (LoginStrategy, error) {
	switch loginType {
	case valueobject.LoginTypeMSISDN:
		return NewMSISDNLoginStrategy(f.userRepo), nil
	case valueobject.LoginTypeEmail:
		return NewEmailLoginStrategy(f.userRepo), nil
	case valueobject.LoginTypeGoogle:
		if f.googleOAuthSvc == nil {
			return nil, errors.New("google oauth not configured")
		}
		return NewGoogleOAuthLoginStrategy(f.userRepo, f.googleOAuthSvc), nil
	default:
		return nil, errors.New("unsupported login type")
	}
}
