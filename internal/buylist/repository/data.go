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
