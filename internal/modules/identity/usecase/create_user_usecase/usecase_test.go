package create_user_usecase

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/cristiano-pacheco/shoplist/internal/module/identity/repository/mocks"
// 	hashservice_mocks "github.com/cristiano-pacheco/shoplist/internal/module/identity/service/hashservice/mocks"
// 	"github.com/cristiano-pacheco/shoplist/internal/shared/model"
// 	validator_mocks "github.com/cristiano-pacheco/shoplist/internal/shared/validator/mocks"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateUserUseCase_Execute_Success(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	input := Input{
// 		Name:     "John Doe",
// 		Email:    "test@example.com",
// 		Password: "password",
// 	}
// 	hashedPassword := []byte("hashed password")

// 	validatorMock := validator_mocks.MockValidateI{}
// 	userRepoMock := mocks.MockUserRepositoryI{}
// 	hashServiceMock := hashservice_mocks.MockHashServiceI{}

// 	useCase := New(&userRepoMock, &validatorMock, &hashServiceMock)

// 	now := time.Now()
// 	userModelInput := model.UserModel{
// 		Name:         input.Name,
// 		Email:        input.Email,
// 		PasswordHash: string(hashedPassword),
// 	}

// 	userModelOutput := model.UserModel{
// 		Base: model.Base{
// 			ID:        uint64(1),
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		Name:         input.Name,
// 		Email:        input.Email,
// 		PasswordHash: string(hashedPassword),
// 		IsActivated:  true,
// 	}

// 	validatorMock.On("Struct", input).Once().Return(nil)
// 	userRepoMock.On("Create", ctx, userModelInput).Once().Return(&userModelOutput, nil)
// 	hashServiceMock.On("GenerateFromPassword", []byte(input.Password)).Once().Return(hashedPassword, nil)

// 	// Act
// 	result, err := useCase.Execute(ctx, input)

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, uint64(1), result.UserID)
// 	assert.Equal(t, "John Doe", result.Name)
// 	assert.Equal(t, "test@example.com", result.Email)
// }

// func TestCreateUserUseCase_Execute_ValidationError(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	input := Input{
// 		Name:     "John Doe",
// 		Email:    "test@example.com",
// 		Password: "password",
// 	}

// 	validatorMock := validator_mocks.MockValidateI{}
// 	userRepoMock := mocks.MockUserRepositoryI{}
// 	hashServiceMock := hashservice_mocks.MockHashServiceI{}

// 	useCase := New(&userRepoMock, &validatorMock, &hashServiceMock)

// 	userModelInput := model.UserModel{}

// 	validatorMock.On("Struct", input).Return(fmt.Errorf("error"))
// 	userRepoMock.AssertNotCalled(t, "Create", ctx, userModelInput)

// 	// Act
// 	result, err := useCase.Execute(ctx, input)

// 	// Assert
// 	assert.Error(t, err)
// 	assert.Equal(t, Output{}, result)
// }

// func TestCreateUserUseCase_Execute_RepositoryError(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	input := Input{
// 		Name:     "John Doe",
// 		Email:    "test@example.com",
// 		Password: "password",
// 	}
// 	hashedPassword := []byte("hashed password")

// 	validatorMock := validator_mocks.MockValidateI{}
// 	userRepoMock := mocks.MockUserRepositoryI{}
// 	hashServiceMock := hashservice_mocks.MockHashServiceI{}

// 	useCase := New(&userRepoMock, &validatorMock, &hashServiceMock)

// 	userModelInput := model.UserModel{
// 		Name:         input.Name,
// 		Email:        input.Email,
// 		PasswordHash: string(hashedPassword),
// 	}

// 	validatorMock.On("Struct", input).Return(nil)
// 	hashServiceMock.On("GenerateFromPassword", []byte(input.Password)).Once().Return(hashedPassword, nil)
// 	userRepoMock.On("Create", ctx, userModelInput).Return(nil, fmt.Errorf("error"))

// 	// Act
// 	result, err := useCase.Execute(ctx, input)

// 	// Assert
// 	assert.Error(t, err)
// 	assert.Equal(t, Output{}, result)
// }
