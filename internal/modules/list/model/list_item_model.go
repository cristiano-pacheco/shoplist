package model

import "time"

type ListItemModel struct {
	ID              uint64    `gorm:"primarykey;autoIncrement;column:id"`
	UserID          uint64    `gorm:"type:bigint;not null; column:user_id"`
	ListID          uint64    `gorm:"type:bigint;not null; column:list_id"`
	CategoryID      uint64    `gorm:"type:bigint;not null; column:category_id"`
	Name            string    `gorm:"type:varchar(255);not null; column:name"`
	CurrentQuantity int       `gorm:"type:int;not null; column:current_quantity"`
	MinimumQuantity int       `gorm:"type:int;not null; column:minimum_quantity"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:now();column:created_at"`
}

func (ListItemModel) TableName() string {
	return "list_items"
}
