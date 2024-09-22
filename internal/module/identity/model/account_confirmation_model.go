package model

import "time"

type AccountConfirmationModel struct {
	UserID    uint64    `gorm:"not null;unique;column:user_id"`
	Token     string    `gorm:"not null;column:token"`
	CreatedAt time.Time `gorm:"not null;column:created_at"`
}

func (AccountConfirmationModel) TableName() string {
	return "account_confirmation"
}
