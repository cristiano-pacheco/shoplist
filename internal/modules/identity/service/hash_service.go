package service

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

type HashServiceI interface {
	GenerateFromPassword(password []byte) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
	GenerateRandomBytes() ([]byte, error)
}

type HashService struct {
}

func NewHashService() HashServiceI {
	return &HashService{}
}

const defaultTotalRandomBytesSize = 128

func (s HashService) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (s HashService) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (s HashService) GenerateRandomBytes() ([]byte, error) {
	buffer := make([]byte, defaultTotalRandomBytesSize)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
