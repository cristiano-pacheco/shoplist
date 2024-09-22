package hash_service

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

type ServiceI interface {
	GenerateFromPassword(password []byte) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
	GenerateRandomBytes() ([]byte, error)
}

type service struct {
}

func New() ServiceI {
	return service{}
}

const defaultTotalRandomBytesSize = 128

func (s service) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (s service) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (s service) GenerateRandomBytes() ([]byte, error) {
	buffer := make([]byte, defaultTotalRandomBytesSize)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
