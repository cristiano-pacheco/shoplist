package jwtparser

import "github.com/golang-jwt/jwt/v5"

func New() *jwt.Parser {
	return jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))
}
