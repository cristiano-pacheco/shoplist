package activate_user_usecase

import "context"

type UseCase struct {
}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Execute(ctx context.Context, input Input) error {
	return nil
}
