package jwt

import (
	"fmt"
	"github.com/rfanazhari/factory-login/internal/domain/entity"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
	"time"
)

// TokenService JWTTokenService implements TokenService interface
type TokenService struct{}

func NewJWTTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) GenerateAccessToken(userID valueobject.UserID) (*entity.Authentication, error) {
	// Mock implementation - in real scenario, use proper JWT library
	token := fmt.Sprintf("jwt_token_%s_%d", userID.String(), time.Now().Unix())
	expiresAt := time.Now().Add(24 * time.Hour)

	return entity.NewAuthentication(userID, token, expiresAt), nil
}
