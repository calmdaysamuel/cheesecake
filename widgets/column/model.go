package column

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Style             lipgloss.Style
	Children          []widget.Widget
	MainAxisAlignment lipgloss.Position
}

func (m *Model) Element() widget.Element {
	return &Element{
		parentWidget: m,
		ID:           random.ID(),
	}
}

func New(children []widget.Widget, options ...Option) *Model {
	m := &Model{
		Children: children,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func WithMainAxisAlignment(position lipgloss.Position) Option {
	return func(model *Model) {
		model.MainAxisAlignment = position
	}
}
