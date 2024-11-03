package errs

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	lib_validator "github.com/go-playground/validator/v10"
)

func (em *errorMapper) mapError(err error) error {
	validationErrors, ok := err.(lib_validator.ValidationErrors)

	// validation error flow
	if ok {
		var details []detail
		for _, e := range validationErrors {
			details = append(details, detail{
				Field:   camelToSnake(e.Field()),
				Message: e.Translate(em.translator),
			})
		}

		e := Error{
			Status:        http.StatusUnprocessableEntity,
			OriginalError: err,
			Err: er{
				Code:    codeInvalidArgument,
				Message: mapMessage(codeInvalidArgument),
				Details: details,
			},
		}

		return &e
	}

	var (
		status = http.StatusInternalServerError
		code   = codeUnknown
	)

	switch {
	// Authentication
	case errors.Is(err, ErrInvalidCredentials),
		errors.Is(err, ErrUserIsNotActivated),
		errors.Is(err, ErrInvalidToken):
		status = http.StatusUnauthorized
		code = codeUnauthorized
	default:
		// Defaut: internal server error
		code = codeUnknown
	}

	e := Error{
		Status:        status,
		OriginalError: err,
		Err: er{
			Code:    code,
			Message: mapMessage(code),
		},
	}

	return &e
}

func camelToSnake(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
