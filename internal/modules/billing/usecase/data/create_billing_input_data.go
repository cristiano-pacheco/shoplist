package data

type CreateBillingInput struct {
	CategoryID uint64  `validate:"required,gte=0"`
	UserID     uint64  `validate:"required,gte=0"`
	Name       string  `validate:"required,min=3,max=255"`
	Amount     float64 `validate:"required,gte=0"`
}

type CreateBillingOutput struct {
	ID         uint64
	CategoryID uint64
	UserID     uint64
	Name       string
	Amount     float64
	CreatedAt  string
	UpdatedAt  string
}
