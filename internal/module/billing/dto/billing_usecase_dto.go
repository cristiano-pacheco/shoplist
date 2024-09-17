package dto

type CreateBillingInput struct {
	Name       string  `validate:"required,min=3,max=255"`
	CategoryID uint64  `validate:"required,gte=0"`
	UserID     uint64  `validate:"required,gte=0"`
	Amount     float64 `validate:"required,gte=0"`
}

type CreateBillingOutput struct {
	Name       string
	CreatedAt  string
	UpdatedAt  string
	ID         uint64
	CategoryID uint64
	UserID     uint64
	Amount     float64
}
