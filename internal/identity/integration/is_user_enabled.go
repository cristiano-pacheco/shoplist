package integration

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
)

type IsUserActivated struct {
	userRepository repository.UserRepositoryI
}

func NewIsUserActivated(userRepository repository.UserRepositoryI) *IsUserActivated {
	return &IsUserActivated{userRepository: userRepository}
}

func (i *IsUserActivated) Execute(userID uint64) bool {
	return i.userRepository.IsActivated(context.Background(), userID)
}
