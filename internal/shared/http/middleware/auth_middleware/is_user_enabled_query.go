package auth_middleware

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"gorm.io/gorm"
)

type isUserEnabledQuery struct {
	db *gorm.DB
}

func newIsUserEnabledQuery(db *database.ShoplistDB) *isUserEnabledQuery {
	return &isUserEnabledQuery{db.DB}
}

func (q *isUserEnabledQuery) Execute(ctx context.Context, userID uint64) (bool, error) {
	var exists bool
	err := q.db.WithContext(ctx).
		Select("1").
		Table("users").
		Where("id = ? AND is_activated = ?", userID, true).
		Limit(1).
		Scan(&exists).
		Error

	if err != nil {
		return false, err
	}

	return exists, nil
}
