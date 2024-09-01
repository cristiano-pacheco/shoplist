package model

import "time"

type Base struct {
	ID        uint64    `gorm:"primarykey;autoIncrement;column:id"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now();column:updated_at"`
}
