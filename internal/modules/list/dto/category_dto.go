package dto

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

type CategoryCreateResponse struct {
	Category
}

type CategoryFindRequest struct {
	CategoryID   uint64 `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type CategoryFindResponse struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
