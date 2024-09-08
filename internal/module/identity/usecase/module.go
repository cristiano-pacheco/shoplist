package usecase

import "go.uber.org/fx"

var Module = fx.Module(
	"identity/usecase",
	fx.Provide(NewCreateUserUseCaseUseCase),
	fx.Provide(NewFindUserByIDUseCase),
	fx.Provide(NewUpdateUserUseCase),
)
