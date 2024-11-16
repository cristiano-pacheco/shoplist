package generate_token_usecase

type Input struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type Output struct {
	Token string
}
