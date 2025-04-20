package model

import (
	"errors"
	"time"
)

type UserModel struct {
	id           uint
	name         NameModel
	email        EmailModel
	passwordHash string
	isActivated  bool
	rpToken      string // Reset password token
	createdAt    time.Time
	updatedAt    time.Time
}

func CreateUserModel(name string, email string, passwordHash string) (UserModel, error) {
	// Validate and create name model
	nameModel, err := CreateNameModel(name)
	if err != nil {
		return UserModel{}, err
	}

	// Validate and create email model
	emailModel, err := CreateEmailModel(email)
	if err != nil {
		return UserModel{}, err
	}

	// Validate password hash
	if passwordHash == "" {
		return UserModel{}, errors.New("password hash is required")
	}

	// Create user model
	return UserModel{
		name:         *nameModel,
		email:        *emailModel,
		passwordHash: passwordHash,
		isActivated:  false,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

func RestoreUserModel(
	id uint,
	name string,
	email string,
	passwordHash string,
	isActivated bool,
	rpToken string,
	createdAt time.Time,
	updatedAt time.Time,
) (UserModel, error) {
	// Validate and create name model
	nameModel, err := CreateNameModel(name)
	if err != nil {
		return UserModel{}, err
	}

	// Validate and create email model
	emailModel, err := CreateEmailModel(email)
	if err != nil {
		return UserModel{}, err
	}

	// Validate password hash
	if passwordHash == "" {
		return UserModel{}, errors.New("password hash is required")
	}

	// Create user model with all fields
	return UserModel{
		id:           id,
		name:         *nameModel,
		email:        *emailModel,
		passwordHash: passwordHash,
		isActivated:  isActivated,
		rpToken:      rpToken,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}, nil
}

func (u *UserModel) ID() uint {
	return u.id
}

func (u *UserModel) Name() string {
	return u.name.String()
}

func (u *UserModel) Email() string {
	return u.email.String()
}

func (u *UserModel) PasswordHash() string {
	return u.passwordHash
}

func (u *UserModel) IsActivated() bool {
	return u.isActivated
}

func (u *UserModel) RpToken() string {
	return u.rpToken
}

func (u *UserModel) CreatedAt() time.Time {
	return u.createdAt
}

func (u *UserModel) UpdatedAt() time.Time {
	return u.updatedAt
}
