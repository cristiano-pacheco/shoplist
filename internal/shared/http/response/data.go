package response

// Data is a generic response wrapper that can hold any type T
type Data[T any] struct {
	Data T `json:"data"`
}

// NewData creates a new Data instance with the provided value
func NewData[T any](data T) Data[T] {
	return Data[T]{
		Data: data,
	}
}
