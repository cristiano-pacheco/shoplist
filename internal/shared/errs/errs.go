package errs

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
	ut "github.com/go-playground/universal-translator"
)

type ErrorMapperI interface {
	Map(err error) error
}

type errorMapper struct {
	validate   validator.ValidateI
	translator ut.Translator
}

func New(validate validator.ValidateI, translator ut.Translator) ErrorMapperI {
	return &errorMapper{validate, translator}
}

func (em *errorMapper) Map(err error) error {
	return em.mapError(err)
}
