package model

import "github.com/cristiano-pacheco/go-modulith/internal/shared/model"

type BillingModel struct {
	model.Base
	CategoryID uint64
	UserID     uint64
	Name       string
	Amount     float64
	Category   CategoryModel
	User       model.UserModel
}

func (BillingModel) TableName() string {
	return "billing"
}
