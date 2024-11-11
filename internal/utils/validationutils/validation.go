package validationutils

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	uppercaseRegexp   = regexp.MustCompile(`[A-Z]`)
	numberRegexp      = regexp.MustCompile(`[0-9]`)
	specialCharRegexp = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':",\.<>\/\\\|` + "`" + `~]`)
	phoneRegexp       = regexp.MustCompile(`^(\+62|62|08)[0-9]{8,12}$`)
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok {
		if strings.Contains(password, " ") {
			return false
		}

		if !uppercaseRegexp.MatchString(password) {
			return false
		}
		if !numberRegexp.MatchString(password) {
			return false
		}
		if !specialCharRegexp.MatchString(password) {
			return false
		}

		if len(password) < 8 || len(password) > 255 {
			return false
		}
		return true
	}

	return false
}

func PhoneNumberValidator(fl validator.FieldLevel) bool {
	phoneNumber, ok := fl.Field().Interface().(string)
	if ok {
		phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
		phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")

		if phoneRegexp.MatchString(phoneNumber) {
			return true
		}
	}

	return false
}
