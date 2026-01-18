package row

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
	Children          []widget.Widget
	MainAxisAlignment lipgloss.Position
	*mouseactions.Manager
	BgColor lipgloss.Color
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:           random.ID(),
		parentWidget: m,
		Manager:      m.Manager,
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
