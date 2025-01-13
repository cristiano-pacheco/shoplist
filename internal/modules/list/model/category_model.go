package model

import "time"

type CategoryModel struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `gorm:"type:bigint;not null; column:user_id"`
	Name      string    `gorm:"type:varchar(255);not null; column:name"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at"`
}

func (CategoryModel) TableName() string {
	return "category"
}
