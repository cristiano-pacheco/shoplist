package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository/mocks"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	validator_mocks "github.com/cristiano-pacheco/go-modulith/internal/shared/validator/mocks"
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

	validatorMock := validator_mocks.MockValidateI{}
	userRepoMock := mocks.MockUserRepositoryI{}

	useCase := NewCreateUserUseCaseUseCase(&userRepoMock, &validatorMock)

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
		IsActivated:  true,
	}

	validatorMock.On("Struct", input).Return(nil)
	userRepoMock.On("Create", ctx, userModelInput).Return(&userModelOutput, nil)

	// Act
	result, err := useCase.Execute(ctx, input)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), result.UserID)
	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, "test@example.com", result.Email)
}

func TestCreateUserUseCase_Execute_ValidationError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: "password",
	}

	validatorMock := validator_mocks.MockValidateI{}
	userRepoMock := mocks.MockUserRepositoryI{}

	useCase := NewCreateUserUseCaseUseCase(&userRepoMock, &validatorMock)

	userModelInput := model.UserModel{}

	validatorMock.On("Struct", input).Return(fmt.Errorf("error"))
	userRepoMock.AssertNotCalled(t, "Create", ctx, userModelInput)

	// Act
	result, err := useCase.Execute(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, dto.CreateUserOutput{UserID: 0, Name: "", Email: ""}, result)
}

func TestCreateUserUseCase_Execute_RepositoryError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: "password",
	}

	validatorMock := validator_mocks.MockValidateI{}
	userRepoMock := mocks.MockUserRepositoryI{}

	useCase := NewCreateUserUseCaseUseCase(&userRepoMock, &validatorMock)

	userModelInput := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.Password,
	}

	validatorMock.On("Struct", input).Return(nil)
	userRepoMock.On("Create", ctx, userModelInput).Return(nil, fmt.Errorf("error"))

	// Act
	result, err := useCase.Execute(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, dto.CreateUserOutput{}, result)
}
