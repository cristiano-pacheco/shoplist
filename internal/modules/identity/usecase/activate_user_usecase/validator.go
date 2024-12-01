package activate_user_usecase

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/sdk/empty"
)

func validateInput(input Input) error {
	if empty.IsEmpty(input.UserID) {
		return errs.NewBadRequestError("user_id is required")
	}

	if empty.IsEmpty(input.Token) {
		return errs.NewBadRequestError("token is required")
	}

	return nil
}
