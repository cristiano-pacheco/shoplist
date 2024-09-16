package errs

import "errors"

var (
	ErrUserIsNotActivated = errors.New("the user is not activated")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
)
