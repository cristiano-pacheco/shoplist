package errs

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
	ut "github.com/go-playground/universal-translator"
)

type ErrorMapper interface {
	Map(err error) error
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
