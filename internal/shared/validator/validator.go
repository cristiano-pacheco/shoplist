package validator

import "github.com/go-playground/validator/v10"

type ValidateI interface {
	Struct(s interface{}) error
}

func New() ValidateI {
	return validator.New(validator.WithRequiredStructEnabled())
}
