package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserModel(t *testing.T) {
	tests := []struct {
		name          string
		userName      string
		email         string
		passwordHash  string
		expectError   bool
		errorContains string
	}{
		{
			name:         "Valid user",
			userName:     "John Doe",
			email:        "john@example.com",
			passwordHash: "hashed_password",
			expectError:  false,
		},
		{
			name:          "Empty name",
			userName:      "",
			email:         "john@example.com",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "name is required",
		},
		{
			name:          "Name too short",
			userName:      "J",
			email:         "john@example.com",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "name must be at least 2 characters long",
		},
		{
			name:          "Name too long",
			userName:      string(make([]byte, 256)),
			email:         "john@example.com",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "name cannot exceed 255 characters",
		},
		{
			name:          "Invalid email",
			userName:      "John Doe",
			email:         "invalid-email",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "invalid email format",
		},
		{
			name:          "Empty email",
			userName:      "John Doe",
			email:         "",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "email is required",
		},
		{
			name:          "Empty password hash",
			userName:      "John Doe",
			email:         "john@example.com",
			passwordHash:  "",
			expectError:   true,
			errorContains: "password hash is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := CreateUserModel(tt.userName, tt.email, tt.passwordHash)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
				if err != nil {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				require.NotNil(t, user)
				
				// Verify name
				nameModel, err := CreateNameModel(tt.userName)
				require.NoError(t, err)
				assert.Equal(t, nameModel.String(), user.Name())
				
				// Verify email
				emailModel, err := CreateEmailModel(tt.email)
				require.NoError(t, err)
				assert.Equal(t, emailModel.String(), user.Email())
				
				// Verify other fields
				assert.Equal(t, tt.passwordHash, user.PasswordHash())
				assert.False(t, user.IsActivated())
				assert.Empty(t, user.RpToken())
				
				// Verify timestamps are set
				assert.False(t, user.CreatedAt().IsZero())
				assert.False(t, user.UpdatedAt().IsZero())
				
				// CreatedAt and UpdatedAt should be very close to each other
				timeDiff := user.UpdatedAt().Sub(user.CreatedAt())
				assert.LessOrEqual(t, timeDiff.Milliseconds(), int64(100))
			}
		})
	}
}

func TestUserModel_Fields(t *testing.T) {
	// Create test data
	now := time.Now()
	nameModel, err := CreateNameModel("John Doe")
	require.NoError(t, err)
	emailModel, err := CreateEmailModel("john@example.com")
	require.NoError(t, err)
	
	// Create a user model manually for field testing
	user := UserModel{
		id:           123,
		name:         *nameModel,
		email:        *emailModel,
		passwordHash: "hashed_password",
		isActivated:  true,
		rpToken:      "reset_token",
		createdAt:    now,
		updatedAt:    now,
	}
	
	// Test field values
	assert.Equal(t, uint(123), user.ID())
	assert.Equal(t, "John Doe", user.Name())
	assert.Equal(t, "john@example.com", user.Email())
	assert.Equal(t, "hashed_password", user.PasswordHash())
	assert.True(t, user.IsActivated())
	assert.Equal(t, "reset_token", user.RpToken())
	assert.Equal(t, now, user.CreatedAt())
	assert.Equal(t, now, user.UpdatedAt())
}

func TestRestoreUserModel(t *testing.T) {
	tests := []struct {
		name        string
		id          uint
		userName    string
		email       string
		password    string
		isActivated bool
		rpToken     string
		createdAt   time.Time
		updatedAt   time.Time
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid user restoration",
			id:          1,
			userName:    "John Doe",
			email:       "test@example.com",
			password:    "hashedpassword",
			isActivated: true,
			rpToken:     "token123",
			createdAt:   time.Now().Add(-24 * time.Hour),
			updatedAt:   time.Now(),
			expectError: false,
		},
		{
			name:        "Invalid email",
			id:          2,
			userName:    "John Doe",
			email:       "invalid-email",
			password:    "hashedpassword",
			isActivated: true,
			rpToken:     "token123",
			createdAt:   time.Now().Add(-24 * time.Hour),
			updatedAt:   time.Now(),
			expectError: true,
			errorMsg:    "invalid email format",
		},
		{
			name:        "Invalid name",
			id:          3,
			userName:    "J", // Too short
			email:       "test@example.com",
			password:    "hashedpassword",
			isActivated: true,
			rpToken:     "token123",
			createdAt:   time.Now().Add(-24 * time.Hour),
			updatedAt:   time.Now(),
			expectError: true,
			errorMsg:    "name must be at least 2 characters",
		},
		{
			name:        "Empty password hash",
			id:          4,
			userName:    "John Doe",
			email:       "test@example.com",
			password:    "",
			isActivated: true,
			rpToken:     "token123",
			createdAt:   time.Now().Add(-24 * time.Hour),
			updatedAt:   time.Now(),
			expectError: true,
			errorMsg:    "password hash is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			user, err := RestoreUserModel(
				tc.id,
				tc.userName,
				tc.email,
				tc.password,
				tc.isActivated,
				tc.rpToken,
				tc.createdAt,
				tc.updatedAt,
			)

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorMsg != "" {
					assert.Contains(t, err.Error(), tc.errorMsg)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.id, user.ID())
				assert.Equal(t, tc.email, user.Email())
				assert.Equal(t, tc.userName, user.Name())
				assert.Equal(t, tc.password, user.PasswordHash())
				assert.Equal(t, tc.isActivated, user.IsActivated())
				assert.Equal(t, tc.rpToken, user.RpToken())
				assert.Equal(t, tc.createdAt.Unix(), user.CreatedAt().Unix())
				assert.Equal(t, tc.updatedAt.Unix(), user.UpdatedAt().Unix())
			}
		})
	}
}
