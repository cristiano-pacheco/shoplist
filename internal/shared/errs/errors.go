package errs

import "errors"

var (
	// Authentication & Authorization
	ErrUserIsNotActivated = errors.New("the user is not activated")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")

	// Database
	ErrNotFound = errors.New("not found")

	// private key
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
)
