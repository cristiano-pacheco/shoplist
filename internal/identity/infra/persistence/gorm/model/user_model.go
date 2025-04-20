package model

import "time"

type UserModel struct {
	ID           uint64    `gorm:"primarykey;autoIncrement;column:id"`
	Name         string    `gorm:"type:varchar;not null; column:name"`
	Email        string    `gorm:"type:varchar;not null; column:email"`
	PasswordHash string    `gorm:"type:varchar;not null; column:password_hash"`
	IsActivated  bool      `gorm:"type:boolean;not null;default:false;column:is_activated"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:now();column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:now();column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
