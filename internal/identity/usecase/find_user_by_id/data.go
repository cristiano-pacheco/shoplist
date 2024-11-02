package find_user_by_id_usecase

type Input struct {
	UserID uint64
}

type Output struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	UserID   uint64 `json:"user_id"`
}
