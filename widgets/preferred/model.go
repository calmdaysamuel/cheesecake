package preferred

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Child           widget.Widget
	PreferredHeight int
	PreferredWidth  int
	mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		parent:  m,
		ID:      random.ID(),
		Manager: m.Manager,
	}
}

func Height(child widget.Widget, height int, options ...Option) *Model {
	m := &Model{
		Child:           child,
		PreferredHeight: height,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func Width(child widget.Widget, width int, options ...Option) *Model {
	m := &Model{
		Child:          child,
		PreferredWidth: width,
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
