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

func (s *UserTestSuite) TestCreateUserValidations() {
	testCases := []struct {
		name         string
		payload      map[string]string
		expectedErrs []map[string]string
	}{
		{
			name: "should return error when name is empty",
			payload: map[string]string{
				"email":    "user@email.com",
				"password": "12345678",
			},
			expectedErrs: []map[string]string{
				{
					"field":   "name",
					"message": "Name is a required field",
				},
			},
		},
		{
			name: "should return error when name is less than 3 characters",
			payload: map[string]string{
				"name":     "ab",
				"email":    "user@email.com",
				"password": "12345678",
			},
			expectedErrs: []map[string]string{
				{
					"field":   "name",
					"message": "Name must be at least 3 characters in length",
				},
			},
		},
		{
			name: "should return error when email is empty",
			payload: map[string]string{
				"name":     "User Name",
				"password": "12345678",
			},
			expectedErrs: []map[string]string{
				{
					"field":   "email",
					"message": "Email is a required field",
				},
			},
		},
		{
			name: "should return error when email is invalid",
			payload: map[string]string{
				"name":     "User Name",
				"email":    "invalid-email",
				"password": "12345678",
			},
			expectedErrs: []map[string]string{
				{
					"field":   "email",
					"message": "Email must be a valid email address",
				},
			},
		},
		{
			name: "should return error when password is empty",
			payload: map[string]string{
				"name":  "User Name",
				"email": "user@email.com",
			},
			expectedErrs: []map[string]string{
				{
					"field":   "password",
					"message": "Password is a required field",
				},
			},
		},
		{
			name: "should return error when password is less than 8 characters",
			payload: map[string]string{
				"name":     "User Name",
				"email":    "user@email.com",
				"password": "123",
			},
			expectedErrs: []map[string]string{
				{
					"field":   "password",
					"message": "Password must be at least 8 characters in length",
				},
			},
		},
		{
			name:    "should return multiple errors when multiple fields are invalid",
			payload: map[string]string{},
			expectedErrs: []map[string]string{
				{
					"field":   "name",
					"message": "Name is a required field",
				},
				{
					"field":   "email",
					"message": "Email is a required field",
				},
				{
					"field":   "password",
					"message": "Password is a required field",
				},
			},
		},
	}

	createUserURL := fmt.Sprintf("%s/api/v1/users", s.itest.BaseURL)

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			jsonPayload, err := json.Marshal(tc.payload)
			s.NoError(err)

			req, err := http.NewRequest(http.MethodPost, createUserURL, bytes.NewBuffer(jsonPayload))
			s.NoError(err)

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			s.NoError(err)
			defer resp.Body.Close()

			s.Equal(http.StatusUnprocessableEntity, resp.StatusCode)

			responseBody, err := io.ReadAll(resp.Body)
			s.NoError(err)

			var response struct {
				Error struct {
					Code    string              `json:"code"`
					Message string              `json:"message"`
					Details []map[string]string `json:"details"`
				} `json:"error"`
			}

			err = json.Unmarshal(responseBody, &response)
			s.NoError(err)

			s.Equal("INVALID_ARGUMENT", response.Error.Code)
			s.Equal("Invalid input provided", response.Error.Message)
			s.Equal(tc.expectedErrs, response.Error.Details)
		})
	}
}

func (s *UserTestSuite) TestEmailIsAlreadyInUse() {
	// Create user directly in database
	existingUser := model.UserModel{
		Name:         "Existing User",
		Email:        "existing@email.com",
		PasswordHash: "some-hash",
	}

	err := s.itest.DB.Create(&existingUser).Error
	s.NoError(err)

	// Try to create user with same email via API
	createUserURL := fmt.Sprintf("%s/api/v1/users", s.itest.BaseURL)

	payload := map[string]string{
		"name":     "Another User",
		"email":    "existing@email.com", // Same email as existing user
		"password": "12345678",
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

	s.Equal(http.StatusBadRequest, resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	s.NoError(err)

	var response struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	err = json.Unmarshal(responseBody, &response)
	s.NoError(err)

	s.Equal("UNKNOWN", response.Error.Code)
	s.Equal("email already in use", response.Error.Message)

	// Clean up
	err = s.itest.DB.Unscoped().Delete(&existingUser).Error
	s.NoError(err)
}
