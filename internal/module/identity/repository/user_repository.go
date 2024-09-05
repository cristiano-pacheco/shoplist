package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
)

type UserRepositoryI interface {
	Create(ctx context.Context, model model.UserModel) (*model.UserModel, error)
	Update(ctx context.Context, model model.UserModel) (*model.UserModel, error)
	FindOneByID(ctx context.Context, ID uint64) (*model.UserModel, error)
}

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) UserRepositoryI {
	return &UserRepository{db}
}

func (r *UserRepository) Create(
	ctx context.Context,
	userModel model.UserModel,
) (*model.UserModel, error) {
	result := r.db.Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userModel, nil
}

func (r *UserRepository) Update(
	ctx context.Context,
	model model.UserModel,
) (*model.UserModel, error) {
	return nil, nil
}

func (r *UserRepository) FindOneByID(ctx context.Context, ID uint64) (*model.UserModel, error) {
	return nil, nil
}
