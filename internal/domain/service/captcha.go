package service

import (
	"context"

	"github.com/rfanazhari/factory-login/internal/domain/valueobject"
)

type CaptchaService interface {
	Validate(ctx context.Context, code valueobject.CaptchaCode) error
}
