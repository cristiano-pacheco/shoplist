package find_activated_user_by_id_usecase

type Input struct {
	UserID uint64
}

type Output struct {
	IsUserActivated bool
}
