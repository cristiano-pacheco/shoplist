package create_user_usecase

type Input struct {
	Name     string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type Output struct {
	Name   string
	Email  string
	UserID uint64
}
