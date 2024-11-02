package activate_user_usecase

type Input struct {
	UserID uint64 `json:"user_id" validate:"required,number"`
	Token  string `json:"token" validate:"required"`
}
