package mediator

import (
	"context"
	"errors"
)

type MediatorI interface {
	Register(event string, callback func(ctx context.Context, input any) (any, error))
	Execute(ctx context.Context, event string, input any) (any, error)
}

type mediator struct {
	handlers map[string]func(ctx context.Context, input any) (any, error)
}

func New() MediatorI {
	return &mediator{
		handlers: make(map[string]func(ctx context.Context, input any) (any, error)),
	}
}

func (m *mediator) Register(event string, callback func(ctx context.Context, input any) (any, error)) {
	m.handlers[event] = callback
}

func (m *mediator) Execute(ctx context.Context, event string, input any) (any, error) {
	if _, ok := m.handlers[event]; !ok {
		return nil, errors.New("handler not found")
	}

	callback := m.handlers[event]
	return callback(ctx, input)
}
