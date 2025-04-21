package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cristiano-pacheco/shoplist/internal/kernel/errs"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/response"
)

type ErrorHandlerMiddleware struct {
	errorMapper errs.ErrorMapper
}

func NewErrorHandlerMiddleware(errorMapper errs.ErrorMapper) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{errorMapper}
}

// contextKey type for context values
type contextKey string

// ErrorCtxKey is the key used to store errors in the request context
const ErrorCtxKey contextKey = "request-error"

// Middleware returns a middleware that handles errors for Chi router
func (h *ErrorHandlerMiddleware) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a custom response writer that can capture the status code
			ww := NewResponseWriter(w)

			// Create a context that can be used to store errors
			ctx := r.Context()

			// Call the next handler with our wrapped response writer
			next.ServeHTTP(ww, r.WithContext(ctx))

			// Check if an error was stored in the context
			if err, ok := ctx.Value(ErrorCtxKey).(error); ok && err != nil {
				mappedError := h.errorMapper.Map(err)
				response.Error(w, mappedError)
				return
			}

			// If status code indicates an error and no error was explicitly set
			if ww.Status() >= 400 && ctx.Value(ErrorCtxKey) == nil {
				// Create a generic error for the status code
				err := fmt.Errorf("HTTP error: %s", http.StatusText(ww.Status()))
				mappedErr := h.errorMapper.MapCustomError(ww.Status(), err.Error())
				response.Error(w, mappedErr)
			}
		})
	}
}

// StoreError stores an error in the request context for later handling
// This should be called by handlers when they encounter an error
func StoreError(r *http.Request, err error) *http.Request {
	ctx := context.WithValue(r.Context(), ErrorCtxKey, err)
	return r.WithContext(ctx)
}

// GetError retrieves an error from the request context
func GetError(ctx context.Context) error {
	if err, ok := ctx.Value(ErrorCtxKey).(error); ok {
		return err
	}
	return nil
}

// ResponseWriter is a wrapper around http.ResponseWriter that captures status code
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewResponseWriter creates a new ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader captures the status code before writing it
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Status returns the HTTP status code
func (rw *ResponseWriter) Status() int {
	return rw.statusCode
}
