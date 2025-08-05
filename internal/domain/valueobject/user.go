package valueobject

import (
	"errors"
	"regexp"
)

// UserID represents a unique user identifier
type UserID struct {
	value string
}

func NewUserID(value string) UserID {
	return UserID{value: value}
}

func (u UserID) String() string {
	return u.value
}

// MSISDN represents a mobile phone number
type MSISDN struct {
	value string
}

func NewMSISDN(value string) (*MSISDN, error) {
	msisdnRegex := regexp.MustCompile(`^628\d{8,11}$`)
	if !msisdnRegex.MatchString(value) {
		return nil, errors.New("invalid MSISDN format")
	}
	return &MSISDN{value: value}, nil
}

func (m MSISDN) String() string {
	return m.value
}

// Email represents an email address
type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		return nil, errors.New("invalid email format")
	}
	return &Email{value: value}, nil
}

func (e Email) String() string {
	return e.value
}

// Password represents a hashed password
type Password struct {
	hashedValue string
}

func NewPassword(hashedValue string) Password {
	return Password{hashedValue: hashedValue}
}

func (p Password) Verify(plainPassword string) bool {
	// In real implementation, use bcrypt.CompareHashAndPassword
	return plainPassword == "password123"
}
