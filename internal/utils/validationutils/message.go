package validationutils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TagToMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "password":
		return fmt.Sprintf("%s must contain at least 8 characters including 1 number, 1 special character, and 1 capital letter excluding whitespaces", fe.Field())
	case "len":
		return fmt.Sprintf("%s length or value must be exactly %v", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s length or value %v must be at most", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be lower than or equal to %v", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s has invalid email format", fe.Field())
	case "eq":
		return fmt.Sprintf("%s must be equal to %v", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s length or value must be at least %v", fe.Field(), fe.Param())
	case "numeric":
		return fmt.Sprintf("%s must be a number", fe.Field())
	case "phone_number":
		return fmt.Sprintf("%s has an invalid phone number format", fe.Field())
	default:
		return "invalid input"
	}
}
