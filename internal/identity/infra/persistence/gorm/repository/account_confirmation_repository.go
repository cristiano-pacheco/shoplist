package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/database"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/otel"
)

type AccountConfirmationRepository interface {
	repository.AccountConfirmationRepository
}

type accountConfirmationRepository struct {
	db     *database.ShoplistDB
	mapper mapper.AccountConfirmationMapper
}

func NewAccountConfirmationRepository(
	db *database.ShoplistDB,
	mapper mapper.AccountConfirmationMapper,
) AccountConfirmationRepository {
	return &accountConfirmationRepository{db, mapper}
}

func (r *accountConfirmationRepository) Create(ctx context.Context, m model.AccountConfirmationModel) (model.AccountConfirmationModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "AccountConfirmationRepository.Create")
	defer span.End()

	entity := r.mapper.ToEntity(m)
	err := r.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return model.AccountConfirmationModel{}, err
	}

	return r.mapper.ToModel(entity)
}

func (r *accountConfirmationRepository) FindByUserID(
	ctx context.Context,
	userID uint64,
) (model.AccountConfirmationModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "AccountConfirmationRepository.FindByUserID")
	defer span.End()

	var m entity.AccountConfirmationEntity

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error
	if err != nil {
		return model.AccountConfirmationModel{}, err
	}

	return r.mapper.ToModel(m)
}

func (r *accountConfirmationRepository) Delete(ctx context.Context, m model.AccountConfirmationModel) error {
	ctx, span := otel.Trace().StartSpan(ctx, "AccountConfirmationRepository.Delete")
	defer span.End()
	return r.db.WithContext(ctx).Where("user_id = ?", m.UserID).Delete(&m).Error
}

func (r *accountConfirmationRepository) DeleteById(ctx context.Context, id uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "AccountConfirmationRepository.DeleteById")
	defer span.End()

	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.AccountConfirmationEntity{}).Error
}

func (r *accountConfirmationRepository) FindByToken(ctx context.Context, token string) (model.AccountConfirmationModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "AccountConfirmationRepository.FindByToken")
	defer span.End()

	var entity entity.AccountConfirmationEntity
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&entity).Error
	if err != nil {
		return model.AccountConfirmationModel{}, err
	}

	return r.mapper.ToModel(entity)
}
