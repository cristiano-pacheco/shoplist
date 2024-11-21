package dto

type CreateCategoryRequestDTO struct {
	Name string `json:"name"`
}

type CreateCategoryResponseDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
