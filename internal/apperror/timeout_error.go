package apperror

import (
	"errors"

	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
)

func NewTimeoutError() *AppError {
	msg := constant.RequestTimeoutErrorMessage

	err := errors.New(msg)

	return NewAppError(err, RequestTimeoutErrorCode, msg)
}
