package valueobject

import "errors"

// CaptchaCode represents a captcha validation code
type CaptchaCode struct {
	value string
}

func NewCaptchaCode(value string) (*CaptchaCode, error) {
	if value == "" {
		return nil, errors.New("captcha code is required")
	}
	return &CaptchaCode{value: value}, nil
}

func (c CaptchaCode) String() string {
	return c.value
}

// LoginType represents the type of login method
type LoginType int

const (
	LoginTypeMSISDN LoginType = iota
	LoginTypeEmail
	LoginTypeGoogle
)

func (lt LoginType) String() string {
	switch lt {
	case LoginTypeMSISDN:
		return "msisdn"
	case LoginTypeEmail:
		return "email"
	case LoginTypeGoogle:
		return "google"
	default:
		return "unknown"
	}
}
