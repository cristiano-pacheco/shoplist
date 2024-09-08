package hashservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashService_GenerateFromPassword(t *testing.T) {
	sut := New()
	password := []byte("test")
	result, err := sut.GenerateFromPassword(password)
	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestHashService_CompareHashAndPassword(t *testing.T) {
	sut := New()
	password := []byte("validpassword")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err, "bcrypt.GenerateFromPassword should not return an error")

	t.Run("password matches hash", func(t *testing.T) {
		// Act
		err := sut.CompareHashAndPassword(hashedPassword, password)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("password does not match hash", func(t *testing.T) {
		// Act
		err := sut.CompareHashAndPassword(hashedPassword, []byte("test"))

		// Assert
		assert.Error(t, err)
	})
}

func TestService_GenerateRandomBytes(t *testing.T) {
	s := New()

	t.Run("generate random bytes successfully", func(t *testing.T) {
		// Act
		bytes, err := s.GenerateRandomBytes()

		// Assert
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, defaultTotalRandomBytesSize, len(bytes))
	})
}
