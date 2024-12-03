package activate_user

type Input struct {
	UserID uint64 `validate:"required,number"`
	Token  string `validate:"required"`
}
