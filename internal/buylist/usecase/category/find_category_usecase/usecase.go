package find_category_usecase

import "context"

type UseCaseI interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
