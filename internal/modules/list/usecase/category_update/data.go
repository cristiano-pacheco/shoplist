package category_update

type Input struct {
	UserID     uint64 `validate:"required"`
	CategoryID uint64 `validate:"required"`
	Name       string `validate:"required,min=1,max=255"`
}
