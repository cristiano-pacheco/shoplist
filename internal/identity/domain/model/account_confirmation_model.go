package model

import (
	"errors"
	"strings"
	"time"
)

type AccountConfirmationModel struct {
	id        uint64
	userID    uint64
	token     string
	expiresAt time.Time
	createdAt time.Time
}

func CreateAccountConfirmationModel(
	userID uint64,
	token string,
	expiresAt time.Time,
) (AccountConfirmationModel, error) {
	if userID == 0 {
		return AccountConfirmationModel{}, errors.New("user ID cannot be empty")
	}

	if strings.TrimSpace(token) == "" {
		return AccountConfirmationModel{}, errors.New("token cannot be empty")
	}

	return AccountConfirmationModel{
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		createdAt: time.Now(),
	}, nil
}

func RestoreAccountConfirmationModel(
	id uint64,
	userID uint64,
	token string,
	expiresAt time.Time,
	createdAt time.Time,
) (AccountConfirmationModel, error) {
	if id == 0 {
		return AccountConfirmationModel{}, errors.New("ID cannot be empty")
	}

	if userID == 0 {
		return AccountConfirmationModel{}, errors.New("user ID cannot be empty")
	}

	if strings.TrimSpace(token) == "" {
		return AccountConfirmationModel{}, errors.New("token cannot be empty")
	}

	return AccountConfirmationModel{
		id:        id,
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		createdAt: createdAt,
	}, nil
}

func (c *AccountConfirmationModel) ID() uint64 {
	return c.id
}

func (c *AccountConfirmationModel) UserID() uint64 {
	return c.userID
}

func (c *AccountConfirmationModel) Token() string {
	return c.token
}

func (c *AccountConfirmationModel) ExpiresAt() time.Time {
	return c.expiresAt
}

func (c *AccountConfirmationModel) CreatedAt() time.Time {
	return c.createdAt
}

func (c *AccountConfirmationModel) IsExpired() bool {
	return time.Now().After(c.expiresAt)
}
