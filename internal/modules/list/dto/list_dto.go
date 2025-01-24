package dto

type ListCreateRequest struct {
	Name string `json:"name"`
}

type ListCreateResponse struct {
	List
}

type ListFindRequest struct {
	ListID   uint64 `json:"list_id"`
	ListName string `json:"list_name"`
}

type ListFindResponse struct {
	Lists []List `json:"lists"`
}

type List struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ListUpdateRequest struct {
	Name string `json:"name"`
}

type ListUpdateResponse struct {
	List
}
