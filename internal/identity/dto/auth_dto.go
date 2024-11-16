package dto

type GenerateTokenInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenOutputDTO struct {
	Token string `json:"token"`
}
