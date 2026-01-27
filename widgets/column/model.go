package column

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Style             lipgloss.Style
	children          []widget.Widget
	mainAxisAlignment lipgloss.Position
	*mouseactions.Manager
	BgColor lipgloss.Color
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		parentWidget: m,
		ID:           random.ID(),
		Manager:      m.Manager,
	}
}

func New(children []widget.Widget, options ...Option) *Model {
	m := &Model{
		children: children,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func WithMainAxisAlignment(position lipgloss.Position) Option {
	return func(model *Model) {
		model.mainAxisAlignment = position
	}
}

func WithMouseActions(m *mouseactions.Manager) Option {
	return func(model *Model) {
		model.Manager = m
	}
}

func WithBackgroundColor(color lipgloss.Color) Option {
	return func(model *Model) {
		model.BgColor = color
	}
}
