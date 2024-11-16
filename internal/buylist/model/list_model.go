package model

import (
	"time"
)

type ListModel struct {
	ID        uint64    `gorm:"primarykey;autoIncrement;column:id"`
	UserID    uint64    `gorm:"type:bigint;not null; column:user_id"`
	Name      string    `gorm:"type:varchar(255);not null; column:name"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at"`
}

func (ListModel) TableName() string {
	return "lists"
}
