package apperror

import (
	"errors"

	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
)

func NewUnauthorizedError() *AppError {
	msg := constant.UnauthorizedErrorMessage

	err := errors.New(msg)

	return NewAppError(err, UnauthorizedErrorCode, msg)
}
