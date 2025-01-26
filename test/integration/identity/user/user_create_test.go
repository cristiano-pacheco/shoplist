package identity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/cristiano-pacheco/shoplist/test"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	itest *test.IntegrationTest
}

func (s *UserTestSuite) SetupSuite() {
	s.itest = test.Setup()
	s.itest.SetupDB()
}

func (s *UserTestSuite) TearDownSuite() {
	s.itest.Close()
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestCreateUserSuccess() {
	createUserURL := fmt.Sprintf("%s/api/v1/users", s.itest.BaseURL)

	payload := map[string]string{
		"name":     "User Name",
		"email":    "user@email.com",
		"password": "123123123",
	}

	jsonPayload, err := json.Marshal(payload)
	s.NoError(err)

	req, err := http.NewRequest(http.MethodPost, createUserURL, bytes.NewBuffer(jsonPayload))
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	s.NoError(err)

	var response response.Data[dto.CreateUserResponse]
	err = json.Unmarshal(responseBody, &response)
	s.NoError(err)

	// Verify user in database
	var userModel model.UserModel

	s.Equal(http.StatusCreated, resp.StatusCode)
	s.Equal(payload["name"], response.Data.Name)
	s.Equal(payload["email"], response.Data.Email)

	query := "SELECT id, name, password_hash, email, created_at, updated_at, is_activated FROM users WHERE email = ?"
	err = s.itest.DB.Raw(query, payload["email"]).Scan(&userModel).Error
	s.NoError(err)

	s.Equal(payload["name"], userModel.Name)
	s.Equal(payload["email"], userModel.Email)
	s.Equal(response.Data.UserID, userModel.ID)
	s.False(userModel.IsActivated)
	s.NotZero(userModel.CreatedAt)
	s.NotZero(userModel.UpdatedAt)
}

func (s *UserTestSuite) TestUpdateUser() {
	// Create user directly in database
	user := model.UserModel{
		Name:         "Original Name",
		Email:        "update-test@email.com",
		PasswordHash: "some-hash",
		IsActivated:  true,
	}

	err := s.itest.DB.Create(&user).Error
	s.NoError(err)

	// Prepare update request
	updateUserURL := fmt.Sprintf("%s/api/v1/users/%d", s.itest.BaseURL, user.ID)

	payload := map[string]string{
		"name": "Updated Name",
	}

	jsonPayload, err := json.Marshal(payload)
	s.NoError(err)

	req, err := http.NewRequest(http.MethodPut, updateUserURL, bytes.NewBuffer(jsonPayload))
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)
	defer resp.Body.Close()

	// Verify response
	s.Equal(http.StatusNoContent, resp.StatusCode)

	// Verify user was updated in database
	var updatedUser model.UserModel
	err = s.itest.DB.First(&updatedUser, user.ID).Error
	s.NoError(err)

	s.Equal(payload["name"], updatedUser.Name)
	s.Equal(user.Email, updatedUser.Email)
	s.Equal(user.IsActivated, updatedUser.IsActivated)
	s.Equal(user.PasswordHash, updatedUser.PasswordHash)
}
