package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository/mocks"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: "password",
	}

	v := validator.New()
	userRepoMock := mocks.MockUserRepositoryI{}

	useCase := NewCreateUserUseCaseUseCase(&userRepoMock, v)

	now := time.Now()
	userModelInput := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.Password,
	}

	userModelOutput := model.UserModel{
		Base: model.Base{
			ID:        uint64(1),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.Password,
	}

	userRepoMock.On("Create", ctx, userModelInput).Return(&userModelOutput, nil)

	// Act
	result, err := useCase.Execute(ctx, input)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result.UserID)
	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, "test@example.com", result.Email)
}
