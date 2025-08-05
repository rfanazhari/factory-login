package memory

import (
	"context"
	"github.com/rfanazhari/factory-login/internal/domain"
	"github.com/rfanazhari/factory-login/internal/domain/entity"
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// InMemoryUserRepository implements UserRepository interface
type InMemoryUserRepository struct {
	users map[string]*entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	users := make(map[string]*entity.User)

	// Sample users
	user1 := entity.NewUser(valueobject.NewUserID("1"), valueobject.NewPassword("hashed_password_1"), true)
	msisdn1, _ := valueobject.NewMSISDN("628123456789")
	user1.SetMSISDN(msisdn1)
	users["628123456789"] = user1

	user2 := entity.NewUser(valueobject.NewUserID("2"), valueobject.NewPassword("hashed_password_2"), true)
	email2, _ := valueobject.NewEmail("user@example.com")
	user2.SetEmail(email2)
	users["user@example.com"] = user2

	user3 := entity.NewUser(valueobject.NewUserID("3"), valueobject.NewPassword("hashed_password_3"), false)
	msisdn3, _ := valueobject.NewMSISDN("628987654321")
	user3.SetMSISDN(msisdn3)
	users["628987654321"] = user3

	return &InMemoryUserRepository{users: users}
}

func (r *InMemoryUserRepository) FindByMSISDN(ctx context.Context, msisdn valueobject.MSISDN) (*entity.User, error) {
	if foundUser, exists := r.users[msisdn.String()]; exists {
		return foundUser, nil
	}
	return nil, domain.ErrUserNotFound
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error) {
	if foundUser, exists := r.users[email.String()]; exists {
		return foundUser, nil
	}
	return nil, domain.ErrUserNotFound
}
