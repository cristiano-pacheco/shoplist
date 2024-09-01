package persistence

type BillingModel struct {
	Base
	CategoryID uint64
	UserID     uint64
	Name       string
	Amount     float64
	Category   CategoryModel
	User       User
}

func (BillingModel) TableName() string {
	return "billing"
}
