package dto

type UpdateUserInput struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8"`
	UserID   uint64 `validate:"required"`
}
