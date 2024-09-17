package create_user_usecase

type Input struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Output struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID uint64 `json:"user_id"`
}
