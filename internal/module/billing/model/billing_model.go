package model

import "github.com/cristiano-pacheco/go-modulith/internal/shared/model"

type BillingModel struct {
	model.Base
	Category   CategoryModel
	Name       string
	User       model.UserModel
	CategoryID uint64
	UserID     uint64
	Amount     float64
}

func (BillingModel) TableName() string {
	return "billing"
}
