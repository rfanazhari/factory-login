package entity

import (
	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

// Authentication represents an authenticated session
type Authentication struct {
	userID      valueobject.UserID
	accessToken string
	expiresAt   time.Time
}

func NewAuthentication(userID valueobject.UserID, accessToken string, expiresAt time.Time) *Authentication {
	return &Authentication{
		userID:      userID,
		accessToken: accessToken,
		expiresAt:   expiresAt,
	}
}

func (a *Authentication) UserID() valueobject.UserID {
	return a.userID
}

func (a *Authentication) AccessToken() string {
	return a.accessToken
}

func (a *Authentication) ExpiresAt() time.Time {
	return a.expiresAt
}
