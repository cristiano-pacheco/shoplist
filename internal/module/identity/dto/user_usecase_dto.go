package dto

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserOutput struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type FindUserByIDInput struct {
	UserID uint64
}

type FindUserByIDOutput struct {
	UserID   uint64 `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UpdateUserInput struct {
	UserID   uint64 `validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}
