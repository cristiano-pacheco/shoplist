package dto

type CreateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutputDTO struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type ActivateUserInputDTO struct {
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

type FindUserOutputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserInputDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
