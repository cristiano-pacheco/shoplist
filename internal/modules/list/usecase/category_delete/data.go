package category_delete

type Input struct {
	UserID     uint64 `validate:"required"`
	CategoryID uint64 `validate:"required"`
}
