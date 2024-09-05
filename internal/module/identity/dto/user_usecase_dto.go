package dto

type CreateUserInput struct {
	Name     string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type CreateUserOutput struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
