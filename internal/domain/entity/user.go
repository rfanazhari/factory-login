package entity

import "github.com/rfanazhari/factory-login/internal/domain/valueobject"

// User represents the user aggregate root
type User struct {
	id       valueobject.UserID
	msisdn   *valueobject.MSISDN
	email    *valueobject.Email
	password valueobject.Password
	isActive bool
}

func NewUser(id valueobject.UserID, password valueobject.Password, isActive bool) *User {
	return &User{
		id:       id,
		password: password,
		isActive: isActive,
	}
}

func (u *User) SetMSISDN(msisdn *valueobject.MSISDN) {
	u.msisdn = msisdn
}

func (u *User) SetEmail(email *valueobject.Email) {
	u.email = email
}

func (u *User) ID() valueobject.UserID {
	return u.id
}

func (u *User) MSISDN() *valueobject.MSISDN {
	return u.msisdn
}

func (u *User) Email() *valueobject.Email {
	return u.email
}

func (u *User) VerifyPassword(plainPassword string) bool {
	return u.password.Verify(plainPassword)
}

func (u *User) IsActive() bool {
	return u.isActive
}
