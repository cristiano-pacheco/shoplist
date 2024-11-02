package generate_token_usecase

type Input struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Output struct {
	Token string
}
