package mapper

import (
	"testing"
	"time"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/entity"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginTokenMapper_ToModel(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	consumedAt := now.Add(-1 * time.Hour)
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token entity
	loginTokenEntity := entity.LoginTokenEntity{
		ID:         123,
		UserID:     456,
		Token:      "token123",
		ExpiresAt:  expiresAt,
		ConsumedAt: &consumedAt,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// Create mapper
	mapper := NewLoginTokenMapper()

	// Test ToModel
	loginTokenModel, err := mapper.ToModel(loginTokenEntity)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, uint64(123), loginTokenModel.ID())
	assert.Equal(t, uint64(456), loginTokenModel.UserID())
	assert.Equal(t, "token123", loginTokenModel.Token())
	assert.Equal(t, expiresAt.Unix(), loginTokenModel.ExpiresAt().Unix())

	require.NotNil(t, loginTokenModel.ConsumedAt())
	assert.Equal(t, consumedAt.Unix(), loginTokenModel.ConsumedAt().Unix())

	assert.Equal(t, now.Unix(), loginTokenModel.CreatedAt().Unix())
	assert.Equal(t, now.Unix(), loginTokenModel.UpdatedAt().Unix())

	// Test derived properties
	assert.True(t, loginTokenModel.IsConsumed())
	assert.False(t, loginTokenModel.IsExpired()) // Shouldn't be expired yet
	assert.False(t, loginTokenModel.IsValid())   // Not valid because it's consumed
}

func TestLoginTokenMapper_ToModel_WithNilConsumedAt(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token entity with nil ConsumedAt
	loginTokenEntity := entity.LoginTokenEntity{
		ID:         123,
		UserID:     456,
		Token:      "token123",
		ExpiresAt:  expiresAt,
		ConsumedAt: nil,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// Create mapper
	mapper := NewLoginTokenMapper()

	// Test ToModel
	loginTokenModel, err := mapper.ToModel(loginTokenEntity)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, uint64(123), loginTokenModel.ID())
	assert.Equal(t, uint64(456), loginTokenModel.UserID())
	assert.Equal(t, "token123", loginTokenModel.Token())
	assert.Equal(t, expiresAt.Unix(), loginTokenModel.ExpiresAt().Unix())
	assert.Nil(t, loginTokenModel.ConsumedAt())
	assert.Equal(t, now.Unix(), loginTokenModel.CreatedAt().Unix())
	assert.Equal(t, now.Unix(), loginTokenModel.UpdatedAt().Unix())

	// Test derived properties
	assert.False(t, loginTokenModel.IsConsumed())
	assert.False(t, loginTokenModel.IsExpired()) // Shouldn't be expired yet
	assert.True(t, loginTokenModel.IsValid())    // Valid because it's not consumed and not expired
}

func TestLoginTokenMapper_ToEntity(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	consumedAt := now.Add(-1 * time.Hour)
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token model
	loginTokenModel, err := model.RestoreLoginTokenModel(
		123,
		456,
		"token123",
		expiresAt,
		lo.ToPtr(consumedAt),
		now,
		now,
	)
	require.NoError(t, err)

	// Create mapper
	mapper := NewLoginTokenMapper()

	// Test ToEntity
	loginTokenEntity := mapper.ToEntity(loginTokenModel)

	// Verify fields
	assert.Equal(t, uint64(123), loginTokenEntity.ID)
	assert.Equal(t, uint64(456), loginTokenEntity.UserID)
	assert.Equal(t, "token123", loginTokenEntity.Token)
	assert.Equal(t, expiresAt.Unix(), loginTokenEntity.ExpiresAt.Unix())

	require.NotNil(t, loginTokenEntity.ConsumedAt)
	assert.Equal(t, consumedAt.Unix(), loginTokenEntity.ConsumedAt.Unix())

	assert.Equal(t, now.Unix(), loginTokenEntity.CreatedAt.Unix())
	assert.Equal(t, now.Unix(), loginTokenEntity.UpdatedAt.Unix())
}

func TestLoginTokenMapper_ToEntity_WithNilConsumedAt(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token model with nil ConsumedAt
	loginTokenModel, err := model.RestoreLoginTokenModel(
		123,
		456,
		"token123",
		expiresAt,
		nil,
		now,
		now,
	)
	require.NoError(t, err)

	// Create mapper
	mapper := NewLoginTokenMapper()

	// Test ToEntity
	loginTokenEntity := mapper.ToEntity(loginTokenModel)

	// Verify fields
	assert.Equal(t, uint64(123), loginTokenEntity.ID)
	assert.Equal(t, uint64(456), loginTokenEntity.UserID)
	assert.Equal(t, "token123", loginTokenEntity.Token)
	assert.Equal(t, expiresAt.Unix(), loginTokenEntity.ExpiresAt.Unix())
	assert.Nil(t, loginTokenEntity.ConsumedAt)
	assert.Equal(t, now.Unix(), loginTokenEntity.CreatedAt.Unix())
	assert.Equal(t, now.Unix(), loginTokenEntity.UpdatedAt.Unix())
}

func TestLoginTokenMapper_InvalidData(t *testing.T) {
	// Create test data with invalid values
	now := time.Now().UTC()

	// Test cases with invalid data
	testCases := []struct {
		name        string
		entity      entity.LoginTokenEntity
		errorString string
	}{
		{
			name: "Zero ID",
			entity: entity.LoginTokenEntity{
				ID:         0, // Invalid ID
				UserID:     456,
				Token:      "token123",
				ExpiresAt:  now.Add(24 * time.Hour),
				ConsumedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			errorString: "ID is required",
		},
		{
			name: "Zero UserID",
			entity: entity.LoginTokenEntity{
				ID:         123,
				UserID:     0, // Invalid UserID
				Token:      "token123",
				ExpiresAt:  now.Add(24 * time.Hour),
				ConsumedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			errorString: "user ID is required",
		},
		{
			name: "Empty Token",
			entity: entity.LoginTokenEntity{
				ID:         123,
				UserID:     456,
				Token:      "", // Invalid Token
				ExpiresAt:  now.Add(24 * time.Hour),
				ConsumedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			errorString: "token is required",
		},
		{
			name: "Zero ExpiresAt",
			entity: entity.LoginTokenEntity{
				ID:         123,
				UserID:     456,
				Token:      "token123",
				ExpiresAt:  time.Time{}, // Invalid ExpiresAt
				ConsumedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			errorString: "expiration time is required",
		},
	}

	// Create mapper
	mapper := NewLoginTokenMapper()

	// Test ToModel with invalid data
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := mapper.ToModel(tc.entity)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.errorString)
		})
	}
}

func TestLoginTokenModel_Consume(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token model
	loginTokenModel, err := model.CreateLoginTokenModel(456, "token123", expiresAt)
	require.NoError(t, err)

	// Verify initial state
	assert.False(t, loginTokenModel.IsConsumed())
	assert.Nil(t, loginTokenModel.ConsumedAt())
	assert.True(t, loginTokenModel.IsValid())

	// Consume the token
	loginTokenModel.Consume()

	// Verify the token is now consumed
	assert.True(t, loginTokenModel.IsConsumed())
	require.NotNil(t, loginTokenModel.ConsumedAt())
	assert.False(t, loginTokenModel.IsValid())

	// The consumed timestamp should be approximately now
	assert.WithinDuration(t, time.Now().UTC(), *loginTokenModel.ConsumedAt(), 1*time.Second)

	// Updates should reflect in updatedAt
	assert.True(t, loginTokenModel.UpdatedAt().After(loginTokenModel.CreatedAt()))
}
