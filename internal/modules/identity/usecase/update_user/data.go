package update_user

type Input struct {
	Name     string `validate:"required,min=3,max=255"`
	Password string `validate:"required,min=8"`
	UserID   uint64 `validate:"required"`
}
