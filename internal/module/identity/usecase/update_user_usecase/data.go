package update_user_usecase

type Input struct {
	UserID   uint64 `validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}
