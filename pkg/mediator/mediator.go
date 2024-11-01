package mediator

type MediatorI interface {
	Send(input Input) (Output, error)
}

type mediator struct {
}

func New() MediatorI {
	return &mediator{}
}

func (m *mediator) Send(input Input) (Output, error) {
	return Output{}, nil
}
