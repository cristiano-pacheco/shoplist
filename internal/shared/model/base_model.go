package model

import "time"

type Base struct {
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now();column:updated_at"`
	ID        uint64    `gorm:"primarykey;autoIncrement;column:id"`
}
