package find_user_usecase

type Input struct {
	UserID uint64
}

type Output struct {
	UserID   uint64
	Name     string
	Email    string
	Password string
}
