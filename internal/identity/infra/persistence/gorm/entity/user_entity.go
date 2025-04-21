package entity

import "time"

type UserEntity struct {
	ID           uint64    `gorm:"primarykey;autoIncrement;column:id"`
	Name         string    `gorm:"type:varchar;not null; column:name"`
	Email        string    `gorm:"type:varchar;not null; column:email"`
	PasswordHash string    `gorm:"type:varchar;not null; column:password_hash"`
	IsActivated  bool      `gorm:"type:boolean;not null;default:false;column:is_activated"`
	RpToken      string    `gorm:"type:varchar;not null; column:rp_token"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:now();column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:now();column:updated_at"`
}

func (UserEntity) TableName() string {
	return "users"
}
