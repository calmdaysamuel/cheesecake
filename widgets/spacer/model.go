package spacer

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

type Option func(*Model)

var _ widget.RenderWidget = &Model{}

type Model struct {
	flex int
	mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:      random.ID(),
		flex:    m.flex,
		Manager: m.Manager,
	}
}

func New(flex int, options ...Option) *Model {
	m := &Model{flex: flex}
	for _, option := range options {
		option(m)
	}
	return m
}

func WithMouseActions(m mouseactions.Manager) Option {
	return func(model *Model) {
		model.Manager = m
	}
}
