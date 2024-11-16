package model

import "time"

type CategoryModel struct {
	ID        string    `json:"id"`
	UserID    int64     `gorm:"type:bigint;not null; column:user_id"`
	Name      string    `gorm:"type:varchar(255);not null; column:name"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
