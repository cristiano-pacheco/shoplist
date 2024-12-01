package dto

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CreateCategoryResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
