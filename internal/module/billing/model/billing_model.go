package model

import "github.com/cristiano-pacheco/go-modulith/internal/shared/model"

type BillingModel struct {
	model.Base
	UserID     uint64
	Category   CategoryModel
	Name       string
	CategoryID uint64
	Amount     float64
}

func (BillingModel) TableName() string {
	return "billing"
}
