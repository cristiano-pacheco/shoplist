package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountConfirmationModel(t *testing.T) {
	tests := []struct {
		name          string
		userID        uint
		token         string
		expiresAt     time.Time
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid confirmation",
			userID:      1,
			token:       "valid_token_123",
			expiresAt:   time.Now().Add(24 * time.Hour),
			expectError: false,
		},
		{
			name:          "Empty user ID",
			userID:        0,
			token:         "valid_token_123",
			expiresAt:     time.Now().Add(24 * time.Hour),
			expectError:   true,
			errorContains: "user ID cannot be empty",
		},
		{
			name:          "Empty token",
			userID:        1,
			token:         "",
			expiresAt:     time.Now().Add(24 * time.Hour),
			expectError:   true,
			errorContains: "token cannot be empty",
		},
		{
			name:          "Whitespace token",
			userID:        1,
			token:         "   ",
			expiresAt:     time.Now().Add(24 * time.Hour),
			expectError:   true,
			errorContains: "token cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			confirmation, err := CreateAccountConfirmationModel(tt.userID, tt.token, tt.expiresAt)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, confirmation)
				if err != nil {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				require.NotNil(t, confirmation)
				
				// Verify fields
				assert.Equal(t, tt.userID, confirmation.UserID())
				assert.Equal(t, tt.token, confirmation.Token())
				assert.Equal(t, tt.expiresAt.Unix(), confirmation.ExpiresAt().Unix())
				
				// Verify createdAt is set
				assert.False(t, confirmation.CreatedAt().IsZero())
				
				// ID should be 0 (unset) for a newly created confirmation
				assert.Equal(t, uint(0), confirmation.ID())
			}
		})
	}
}

func TestRestoreAccountConfirmationModel(t *testing.T) {
	now := time.Now()
	expiry := now.Add(24 * time.Hour)
	
	tests := []struct {
		name          string
		id            uint
		userID        uint
		token         string
		expiresAt     time.Time
		createdAt     time.Time
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid restoration",
			id:          1,
			userID:      2,
			token:       "valid_token_123",
			expiresAt:   expiry,
			createdAt:   now,
			expectError: false,
		},
		{
			name:          "Empty ID",
			id:            0,
			userID:        2,
			token:         "valid_token_123",
			expiresAt:     expiry,
			createdAt:     now,
			expectError:   true,
			errorContains: "ID cannot be empty",
		},
		{
			name:          "Empty user ID",
			id:            1,
			userID:        0,
			token:         "valid_token_123",
			expiresAt:     expiry,
			createdAt:     now,
			expectError:   true,
			errorContains: "user ID cannot be empty",
		},
		{
			name:          "Empty token",
			id:            1,
			userID:        2,
			token:         "",
			expiresAt:     expiry,
			createdAt:     now,
			expectError:   true,
			errorContains: "token cannot be empty",
		},
		{
			name:          "Whitespace token",
			id:            1,
			userID:        2,
			token:         "   ",
			expiresAt:     expiry,
			createdAt:     now,
			expectError:   true,
			errorContains: "token cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			confirmation, err := RestoreAccountConfirmationModel(
				tt.id,
				tt.userID,
				tt.token,
				tt.expiresAt,
				tt.createdAt,
			)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, confirmation)
				if err != nil {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				require.NotNil(t, confirmation)
				
				// Verify fields
				assert.Equal(t, tt.id, confirmation.ID())
				assert.Equal(t, tt.userID, confirmation.UserID())
				assert.Equal(t, tt.token, confirmation.Token())
				assert.Equal(t, tt.expiresAt.Unix(), confirmation.ExpiresAt().Unix())
				assert.Equal(t, tt.createdAt.Unix(), confirmation.CreatedAt().Unix())
			}
		})
	}
}

func TestAccountConfirmationModel_IsExpired(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name      string
		expiresAt time.Time
		expected  bool
	}{
		{
			name:      "Not expired",
			expiresAt: now.Add(1 * time.Hour),
			expected:  false,
		},
		{
			name:      "Expired",
			expiresAt: now.Add(-1 * time.Hour),
			expected:  true,
		},
		{
			name:      "Expires now",
			expiresAt: now,
			expected:  true, // time.Now().After(now) will be true because of tiny time differences
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			confirmation, err := CreateAccountConfirmationModel(1, "token", tt.expiresAt)
			require.NoError(t, err)
			
			// Test the IsExpired method
			assert.Equal(t, tt.expected, confirmation.IsExpired())
		})
	}
}
