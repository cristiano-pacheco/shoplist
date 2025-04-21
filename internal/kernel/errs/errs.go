package errs

import (
	"github.com/cristiano-pacheco/shoplist/internal/kernel/validator"
	ut "github.com/go-playground/universal-translator"
)

type ErrorMapper interface {
	Map(err error) error
	MapCustomError(status int, message string) error
}

type errorMapper struct {
	validate   validator.Validate
	translator ut.Translator
}

func New(validate validator.Validate, translator ut.Translator) ErrorMapper {
	return &errorMapper{validate, translator}
}

func (em *errorMapper) Map(err error) error {
	return em.mapError(err)
}

func (em *errorMapper) MapCustomError(status int, message string) error {
	return em.mapCustomError(status, message)
}
