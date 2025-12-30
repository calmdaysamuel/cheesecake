package border

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Child  widget.Widget
	Border lipgloss.Border
	Label  string
	Sides  []bool
	Style  lipgloss.Style
}

func (m *Model) Element() widget.Element {
	return &Element{
		parent: m,
		ID:     random.ID(),
	}
}

func New(child widget.Widget, border lipgloss.Border, options ...Option) *Model {
	m := &Model{
		Child:  child,
		Border: border,
		Sides:  []bool{true},
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func WithBorderTopLabel(label string) Option {
	return func(model *Model) {
		model.Label = label
	}
}

func WithBorderStyle(style lipgloss.Style) Option {
	return func(model *Model) {
		model.Style = style
	}
}
func WithSides(sides ...bool) Option {
	return func(model *Model) {
		model.Sides = sides
	}
}
