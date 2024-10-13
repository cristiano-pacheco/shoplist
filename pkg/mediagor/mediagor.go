package mediagor

import (
	"context"
	"fmt"
)

type MediagorI interface {
	Register(event string, callback func(ctx context.Context, input any) (any, error))
	Execute(ctx context.Context, event string, input any) (any, error)
}

type mediagor struct {
	handlers map[string]func(ctx context.Context, input any) (any, error)
}

func New() MediagorI {
	return &mediagor{
		handlers: make(map[string]func(ctx context.Context, input any) (any, error)),
	}
}

func (m *mediagor) Register(event string, callback func(ctx context.Context, input any) (any, error)) {
	m.handlers[event] = callback
}

func (m *mediagor) Execute(ctx context.Context, event string, input any) (any, error) {
	if _, ok := m.handlers[event]; !ok {
		return nil, fmt.Errorf("handler not found for event: %s", event)
	}

	callback := m.handlers[event]
	return callback(ctx, input)
}
