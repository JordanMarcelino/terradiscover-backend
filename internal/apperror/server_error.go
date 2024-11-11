package apperror

import "github.com/jordanmarcelino/terradiscover-backend/internal/constant"

func NewServerError(err error) *AppError {
	msg := constant.InternalServerErrorMessage

	return NewAppError(err, DefaultServerErrorCode, msg)
}
