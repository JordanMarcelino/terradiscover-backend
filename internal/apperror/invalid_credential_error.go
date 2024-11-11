package apperror

import (
	"errors"

	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
)

func NewInvalidCredentialError() *AppError {
	msg := constant.InvalidCredentialErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg)
}
