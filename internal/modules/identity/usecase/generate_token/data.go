package generate_token

type Input struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type Output struct {
	Token string
}
