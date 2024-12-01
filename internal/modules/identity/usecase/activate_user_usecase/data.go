package activate_user_usecase

type Input struct {
	UserID uint64 `validate:"required,number"`
	Token  string `validate:"required"`
}
