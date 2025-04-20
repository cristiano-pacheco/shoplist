package validator

import (
	"unicode"
	"unicode/utf8"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/domain/err"
)

type PasswordValidator interface {
	Validate(password string) error
}

type passwordValidator struct {
}

func NewPasswordValidator() PasswordValidator {
	return &passwordValidator{}
}

func (s *passwordValidator) Validate(password string) error {
	// Track requirements
	var (
		hasUpper   bool // uppercase letter
		hasLower   bool // lowercase letter
		hasNumber  bool // number
		hasSpecial bool // special character
	)

	// Check minimum length (8 characters)
	if utf8.RuneCountInString(password) < 8 {
		return err.ErrPasswordTooShort
	}

	// Iterate through each rune (character) with UTF-8 support
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsNumber(r):
			hasNumber = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	// Verify all requirements are met
	if !hasUpper {
		return err.ErrPasswordNoUppercase
	}
	if !hasLower {
		return err.ErrPasswordNoLowercase
	}
	if !hasNumber {
		return err.ErrPasswordNoNumber
	}
	if !hasSpecial {
		return err.ErrPasswordNoSpecialChar
	}

	// Password meets all requirements
	return nil
}
