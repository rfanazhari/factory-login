package repository

import (
	"context"

	"github.com/rfanazhari/factory-login/internal/domain/entity"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	FindByMSISDN(ctx context.Context, msisdn valueobject.MSISDN) (*entity.User, error)
	FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error)
}
