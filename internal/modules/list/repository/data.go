package repository

import (
	"time"
)

type FindCategoriesCriteria struct {
	UserID     uint64
	CategoryID *uint64
	Name       *string
}

type DeleteCategoryCriteria struct {
	UserID     uint64
	CategoryID uint64
}

type FindListsCriteria struct {
	UserID    uint64
	Name      string
	CreatedAt time.Time
}

type DeleteListCriteria struct {
	ID     uint64
	UserID uint64
}

type FindListItemsCriteria struct {
	UserID          uint64
	ListID          uint64
	ListItemID      *uint64
	CategoryID      *uint64
	MinimunQuantity *uint64
	CurrentQuantity *uint64
	Name            *string
}

type DeleteListItemCriteria struct {
	UserID     uint64
	ListItemID uint64
}
