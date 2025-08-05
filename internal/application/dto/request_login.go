package dto

import "github.com/rfanazhari/factory-login/internal/domain/valueobject"

// LoginRequest represents the login request data
type LoginRequest struct {
	Identifier  string                `json:"identifier"`
	Password    string                `json:"password"`
	CaptchaCode string                `json:"captcha_code"`
	Type        valueobject.LoginType `json:"type"`
}
