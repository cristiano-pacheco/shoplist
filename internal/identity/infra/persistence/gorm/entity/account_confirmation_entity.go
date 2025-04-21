package entity

import "time"

type AccountConfirmationEntity struct {
	ID        uint64    `gorm:"primary_key;auto_increment;column:id"`
	UserID    uint64    `gorm:"not null;unique;column:user_id"`
	Token     string    `gorm:"not null;column:token"`
	ExpiresAt time.Time `gorm:"not null;column:expires_at"`
	CreatedAt time.Time `gorm:"not null;column:created_at"`
	UpdatedAt time.Time `gorm:"not null;column:updated_at"`
}

func (AccountConfirmationEntity) TableName() string {
	return "account_confirmation"
}
