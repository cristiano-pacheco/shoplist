package dto

type CreateUserAccountInput struct {
	Name     string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type CreateUserAccountOutput struct {
	UserID uint64
	Name   string
	Email  string
}
