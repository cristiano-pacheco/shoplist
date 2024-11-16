package repository

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
	UserID uint64
	ListID *uint64
	Name   *string
}

type DeleteListCriteria struct {
	UserID uint64
	ListID uint64
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
