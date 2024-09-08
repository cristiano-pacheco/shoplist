package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/err"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
)

type UserRepositoryI interface {
	Create(ctx context.Context, model model.UserModel) (*model.UserModel, error)
	Update(ctx context.Context, model model.UserModel) error
	FindOneByID(ctx context.Context, ID uint64) (*model.UserModel, error)
}

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) UserRepositoryI {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, userModel model.UserModel) (*model.UserModel, error) {
	result := r.db.WithContext(ctx).Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userModel, nil
}

func (r *userRepository) Update(ctx context.Context, model model.UserModel) error {
	result := r.db.WithContext(ctx).Omit("created_at", "updated_at").Save(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) FindOneByID(ctx context.Context, ID uint64) (*model.UserModel, error) {
	var userModel model.UserModel
	r.db.WithContext(ctx).First(&userModel, ID)
	if userModel.ID == 0 {
		return nil, err.ErrNotFound
	}
	return &userModel, nil
}
