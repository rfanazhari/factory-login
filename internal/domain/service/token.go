package service

import (
	"github.com/rfanazhari/factory-login/internal/domain/entity"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

type TokenService interface {
	GenerateAccessToken(userID valueobject.UserID) (*entity.Authentication, error)
}
