package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/database"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/errs"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/otel"
)

type LoginTokenRepository interface {
	repository.LoginTokenRepository
}

type loginTokenRepository struct {
	db     *database.ShoplistDB
	mapper mapper.LoginTokenMapper
}

func NewLoginTokenRepository(db *database.ShoplistDB, mapper mapper.LoginTokenMapper) LoginTokenRepository {
	return &loginTokenRepository{db, mapper}
}

func (r *loginTokenRepository) Create(ctx context.Context, loginTokenModel model.LoginTokenModel) (model.LoginTokenModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "LoginTokenRepository.Create")
	defer span.End()

	loginTokenEntity := r.mapper.ToEntity(loginTokenModel)
	result := r.db.WithContext(ctx).Create(&loginTokenEntity)
	if result.Error != nil {
		return model.LoginTokenModel{}, result.Error
	}

	loginTokenModel, err := r.mapper.ToModel(loginTokenEntity)
	if err != nil {
		return model.LoginTokenModel{}, err
	}

	return loginTokenModel, nil
}

func (r *loginTokenRepository) Update(ctx context.Context, loginTokenModel model.LoginTokenModel) error {
	ctx, span := otel.Trace().StartSpan(ctx, "LoginTokenRepository.Update")
	defer span.End()

	loginTokenEntity := r.mapper.ToEntity(loginTokenModel)
	result := r.db.WithContext(ctx).Save(&loginTokenEntity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *loginTokenRepository) Delete(ctx context.Context, id uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "LoginTokenRepository.Delete")
	defer span.End()

	result := r.db.WithContext(ctx).Delete(&entity.LoginTokenEntity{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *loginTokenRepository) FindByToken(ctx context.Context, token string) (model.LoginTokenModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "LoginTokenRepository.FindByToken")
	defer span.End()

	var loginTokenEntity entity.LoginTokenEntity
	result := r.db.WithContext(ctx).Where("token = ?", token).First(&loginTokenEntity)

	if result.Error != nil {
		return model.LoginTokenModel{}, result.Error
	}

	if loginTokenEntity.ID == 0 {
		return model.LoginTokenModel{}, errs.ErrNotFound
	}

	loginTokenModel, err := r.mapper.ToModel(loginTokenEntity)
	if err != nil {
		return model.LoginTokenModel{}, err
	}

	return loginTokenModel, nil
}
