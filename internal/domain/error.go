package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserInactive = errors.New("user account is inactive")
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRateLimitExceeded  = errors.New("rate limit exceeded")
	ErrInvalidCaptcha     = errors.New("invalid captcha code")
)
