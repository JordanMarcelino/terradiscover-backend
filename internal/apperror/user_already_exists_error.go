package apperror

import (
	"errors"

	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
)

func NewUserAlreadyExistsError() *AppError {
	msg := constant.UserAlreadyExistsErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg)
}
