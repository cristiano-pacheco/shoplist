package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/identity/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
)

type UserRepositoryI interface {
	Create(ctx context.Context, model model.UserModel) (*model.UserModel, error)
	Update(ctx context.Context, model model.UserModel) error
	FindByID(ctx context.Context, id uint64) (*model.UserModel, error)
	FindByEmail(ctx context.Context, email string) (*model.UserModel, error)
	IsActivated(ctx context.Context, userID uint64) bool
}

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) UserRepositoryI {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, userModel model.UserModel) (*model.UserModel, error) {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "user_repository.create")
	defer span.End()
	result := r.db.WithContext(ctx).Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userModel, nil
}

func (r *userRepository) Update(ctx context.Context, model model.UserModel) error {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "user_repository.update")
	defer span.End()
	result := r.db.WithContext(ctx).Omit("created_at", "updated_at").Save(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint64) (*model.UserModel, error) {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "user_repository.find_by_id")
	defer span.End()
	var userModel model.UserModel
	r.db.WithContext(ctx).First(&userModel, id)
	if userModel.ID == 0 {
		return nil, errs.ErrNotFound
	}
	return &userModel, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "user_repository.find_by_email")
	defer span.End()
	var userModel model.UserModel
	r.db.WithContext(ctx).Where("email = ?", email).First(&userModel)
	if userModel.ID == 0 {
		return nil, errs.ErrNotFound
	}
	return &userModel, nil
}

func (r *userRepository) IsActivated(ctx context.Context, userID uint64) bool {
	var userModel model.UserModel
	r.db.WithContext(ctx).Where("id = ?", userID).First(&userModel)
	if userModel.ID == 0 {
		return false
	}
	return userModel.IsActivated
}
