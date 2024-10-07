package errormapper

import (
	"errors"
	"regexp"
	"strings"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/logger"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
	ut "github.com/go-playground/universal-translator"
	lib_validator "github.com/go-playground/validator/v10"
)

type Mapper struct {
	validate   validator.ValidateI
	translator ut.Translator
	logger     logger.LoggerI
}

func New(
	validate validator.ValidateI,
	translator ut.Translator,
	logger logger.LoggerI,
) *Mapper {
	return &Mapper{validate, translator, logger}
}

type ResponseError struct {
	ErrorCode ErrorCode
	Errors    []Error `json:"errors"`
}

type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

const ServerErrorMessage = "internal server error"

func (m *Mapper) MapErrorToResponseError(err error) ResponseError {
	var responseError ResponseError
	fieldErrors, ok := err.(lib_validator.ValidationErrors)

	// validation error flow
	if ok {
		var errs []Error
		for _, e := range fieldErrors {
			errs = append(errs, Error{
				Field:   camelToSnake(e.Field()),
				Message: e.Translate(m.translator),
			})
		}

		responseError.ErrorCode = ValidationError
		responseError.Errors = errs

		return responseError
	}

	switch {
	// Authentication
	case errors.Is(err, errs.ErrInvalidCredentials),
		errors.Is(err, errs.ErrUserIsNotActivated),
		errors.Is(err, errs.ErrInvalidToken):
		responseError.ErrorCode = AuthenticationError
	default:
		// Defaut: internal server error
		responseError.ErrorCode = InternalError
	}

	responseError.Errors = []Error{
		{
			Message: err.Error(),
		},
	}

	m.logger.Error(err.Error())

	return responseError
}

func camelToSnake(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
