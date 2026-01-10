package layoutbuilder

import (
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	ChildFunc func(constraints constraints.Constraints) widget.Widget
	mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{parentWidget: m,
		ID:      random.ID(),
		Manager: m.Manager,
	}
}

func New(childFunc func(constraints constraints.Constraints) widget.Widget, options ...Option) *Model {
	m := &Model{
		ChildFunc: childFunc,
	}
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
