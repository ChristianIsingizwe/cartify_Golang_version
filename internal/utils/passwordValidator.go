package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var strongPasswordRegex = regexp.MustCompile(`^(?=.*[A-Z])(?=.*\d)(?=.*[!@#?$])[A-Za-z\d!@#?$]{8,}$`)

func StrongPasswordFunction(fl validator.FieldLevel) bool {
	return strongPasswordRegex.MatchString(fl.Field().String())
}
