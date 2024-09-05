package errormapper

type ErrorCode string

const (
	ValidationError ErrorCode = "VALIDATION_ERROR"
	InternalError   ErrorCode = "INTERNAL_ERROR"
)
